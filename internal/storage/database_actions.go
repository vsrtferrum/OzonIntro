package storage

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/vsrtferrum/OzonIntro/internal/errors"
	"github.com/vsrtferrum/OzonIntro/internal/model"
	"github.com/vsrtferrum/OzonIntro/internal/transform"
)

func (db *Database) GetPostsList() (*[]model.PostList, error) {
	rows, err := db.pool.Query(context.Background(), getPostsList)
	if err != nil {
		return nil, errors.ErrResultQuery
	}
	postList := new([]model.PostList)
	var temp transform.PostList
	for rows.Next() {
		err := rows.Scan(&temp.Id, &temp.Comments)
		if err != nil {
			return nil, errors.ErrResultQuery
		}
		*postList = append(*postList, temp.PostListTransform())
	}
	return postList, nil
}

func (db *Database) GetPostById(id uint64) (*model.Post, *[]model.Comments, error) {
	rows, err := db.pool.Query(context.Background(), getPostById, id)
	if err != nil {
		return nil, nil, errors.ErrResultQuery
	}
	post := []model.Post{}
	temp := transform.Post{}
	for rows.Next() {
		err := rows.Scan(&temp.Id, &temp.Comments, &temp.Text)
		if err != nil {
			return nil, nil, errors.ErrResultQuery
		}
		post = append(post, temp.PostTransform())
		if len(post) > 1 {
			return nil, nil, errors.ErrNonDeterministicId
		}
	}
	rows, err = db.pool.Query(context.Background(), getCommentsByPost, id)
	if err != nil {
		return nil, nil, errors.ErrResultQuery
	}
	comments := []model.Comments{}
	tempComments := transform.Comments{}
	for rows.Next() {
		err := rows.Scan(&tempComments.Comment_id, &tempComments.Post_id, &tempComments.Text)
		if err != nil {
			return nil, nil, errors.ErrResultQuery
		}
		comments = append(comments, tempComments.CommentTransfrom())
	}
	return &post[0], &comments, nil

}

func (db *Database) WriteComment(data *model.WriteComment) (uint64, error) {
	rows, err := db.pool.Query(context.Background(), getCommentsStatus, data.PostId)
	if err != nil {
		return 0, errors.ErrSendQuery
	}
	var temp transform.Post
	res := make([]transform.Post, 0, 1)
	for rows.Next(){
		err := rows.Scan(&temp.Comments)
		if err != nil {
			return 0, errors.ErrResultQuery
		}
		res = append(res, temp)
		if len(res) > 1{
			return 0, errors.ErrNonDeterministicId
		}
	}
	
	if !res[0].Comments{
		return 0 , errors.ErrClosedComments
	}
	
	tx, err := db.pool.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return 0, errors.ErrCreateTransaction
	}
	defer tx.Rollback(context.Background())

	var id uint64
	err = tx.QueryRow(context.Background(), writeComment, data.PostId, data.Text).Scan(&id)
	if err != nil {
		return 0, errors.ErrExecTransaction
	}

	if data.IsReferenceOnComment {
		_, err = tx.Exec(context.Background(), writeReference, data.PostId, data.Text)
		if err != nil {
			return 0, errors.ErrExecTransaction
		}
	}
	err = tx.Commit(context.Background())
	if err != nil {
		return 0, errors.ErrCommitTransaction
	}
	return id, nil
}

func (db *Database) WritePost(data *model.WritePost) (uint64, error) {
	tx, err := db.pool.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return 0, errors.ErrCreateTransaction
	}
	defer tx.Rollback(context.Background())

	var id uint64
	err = tx.QueryRow(context.Background(), writePost, data.Text, data.Comments).Scan(&id)
	if err != nil {
		return 0, errors.ErrExecTransaction
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return 0, errors.ErrCommitTransaction
	}

	return id, nil
}
