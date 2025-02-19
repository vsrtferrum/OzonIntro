package model

type PostList struct {
	Id       uint64
	Comments bool
}

type Post struct {
	Id       uint64
	Text     string
	Comments bool
}
type WritePost struct {
	Text     string
	Comments bool
}
type WriteComment struct {
	PostId, Comment_id   uint64
	IsReferenceOnComment bool
	Text                 string
}

type Comments struct {
	Comment_id, Post_id uint64
	Text                string
}
type CommentRefs struct {
	Comment_id, Reference_id uint64
}
