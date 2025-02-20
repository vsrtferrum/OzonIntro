package storage

import (
	"sync"

	"github.com/vsrtferrum/OzonIntro/internal/errors"
	"github.com/vsrtferrum/OzonIntro/internal/model"
)

type InMemoryStorage struct {
	post                                      *[]model.Post
	comment                                   *[]model.Comments
	commentsRefs                              *[]model.CommentRefs
	postMutex, commentMutex, commentRefsMutex sync.RWMutex
}

func NewInMemoryStorage() *InMemoryStorage {
	post, comment, commentsRefs := new([]model.Post), new([]model.Comments), new([]model.CommentRefs)
	return &InMemoryStorage{post: post, comment: comment, commentsRefs: commentsRefs}
}

func (inMemoryStorage *InMemoryStorage) GetPostsList() (*[]model.PostList, error) {
	if inMemoryStorage.post == nil {
		return nil, errors.ErrNotInitPostList
	}
	postList := make([]model.PostList, 0, len(*inMemoryStorage.post))

	inMemoryStorage.postMutex.RLock()
	defer inMemoryStorage.postMutex.RUnlock()

	for _, val := range *inMemoryStorage.post {
		postList = append(postList, model.PostList{Id: val.Id, Comments: val.Comments})
	}

	return &postList, nil
}
func (inMemoryStorage *InMemoryStorage) GetPostById(id uint64) (*model.Post, *[]model.Comments, error) {
	if inMemoryStorage.post == nil {
		return nil, nil, errors.ErrNotInitPostList
	}

	if inMemoryStorage.comment == nil {
		return nil, nil, errors.ErrNotInitCommentList
	}

	inMemoryStorage.postMutex.RLock()
	defer inMemoryStorage.postMutex.RUnlock()

	if uint64(len(*inMemoryStorage.post)) <= id {
		return nil, nil, errors.ErrPostNotExisted
	}

	inMemoryStorage.commentMutex.RLock()
	defer inMemoryStorage.commentMutex.RUnlock()

	commentsList := make([]model.Comments, 0)
	for _, val := range *inMemoryStorage.comment {
		if uint64(val.Post_id) == id {
			commentsList = append(commentsList, val)
		}

	}

	postCopy := (*(*inMemoryStorage).post)[int(id)]
	return &postCopy, &commentsList, nil

}
func (inMemoryStorage *InMemoryStorage) WriteComment(comment *model.WriteComment) (uint64, error) {
	inMemoryStorage.postMutex.Lock()
	defer inMemoryStorage.postMutex.Unlock()

	if uint64(len(*inMemoryStorage.post)) <= comment.PostId {
		return 0, errors.ErrPostNotExisted
	}

	inMemoryStorage.commentMutex.Lock()
	defer inMemoryStorage.commentMutex.Unlock()

	if comment.IsReferenceOnComment && uint64(len(*inMemoryStorage.comment)) < comment.Comment_id {
		return 0, errors.ErrReferenceCommentNotExisted
	}

	*inMemoryStorage.comment = append(*inMemoryStorage.comment,
		model.Comments{
			Comment_id: uint64(len(*inMemoryStorage.comment)),
			Post_id:    comment.PostId,
			Text:       comment.Text})

	inMemoryStorage.commentRefsMutex.Lock()
	defer inMemoryStorage.commentRefsMutex.Unlock()

	if comment.IsReferenceOnComment {
		*inMemoryStorage.commentsRefs = append(*inMemoryStorage.commentsRefs,
			model.CommentRefs{Comment_id: uint64(len(*inMemoryStorage.comment) - 1),
				Reference_id: comment.PostId})
	}

	return uint64(len(*inMemoryStorage.comment) - 1), nil
}
func (inMemoryStorage *InMemoryStorage) WritePost(post *model.WritePost) (uint64, error) {
	inMemoryStorage.postMutex.Lock()
	defer inMemoryStorage.postMutex.Unlock()

	if inMemoryStorage.post == nil {
		return 0, errors.ErrNotInitPostList
	}

	*inMemoryStorage.post = append(*inMemoryStorage.post,
		model.Post{Id: uint64(len(*inMemoryStorage.post)),
			Text:     post.Text,
			Comments: post.Comments})

	return uint64(len(*inMemoryStorage.post) - 1), nil
}
