package transform

import "github.com/vsrtferrum/OzonIntro/internal/model"

type PostList struct {
	Id       uint64 `db:"id"`
	Comments bool   `db:"comments"`
}

type Post struct {
	Id       uint64 `db:"id"`
	Text     string `db:"text"`
	Comments bool   `db:"comments"`
}

type Comments struct {
	Comment_id uint64 `db:"comment_id"`
	Post_id    uint64 `db:"post_id"`
	Text       string `db:"text"`
}

func (postList *PostList) PostListTransform() model.PostList {
	return model.PostList{Id: postList.Id, Comments: postList.Comments}
}

func (post *Post) PostTransform() model.Post {
	return model.Post{Id: post.Id, Text: post.Text, Comments: post.Comments}
}

func (comments *Comments) CommentTransfrom() model.Comments {
	return model.Comments{Comment_id: comments.Comment_id,
		Post_id: comments.Post_id, Text: comments.Text}
}
