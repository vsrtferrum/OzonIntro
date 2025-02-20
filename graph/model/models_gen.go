// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Comment struct {
	ID     string `json:"id"`
	PostID string `json:"postId"`
	Text   string `json:"text"`
}

type CommentConnection struct {
	Edges    []*CommentEdge `json:"edges"`
	PageInfo *PageInfo      `json:"pageInfo"`
}

type CommentEdge struct {
	Node   *Comment `json:"node"`
	Cursor string   `json:"cursor"`
}

type Mutation struct {
}

type PageInfo struct {
	HasNextPage bool    `json:"hasNextPage"`
	EndCursor   *string `json:"endCursor,omitempty"`
}

type Post struct {
	ID       string             `json:"id"`
	Content  string             `json:"content"`
	Comments *CommentConnection `json:"comments"`
}

type PostList struct {
	ID       string `json:"id"`
	Comments bool   `json:"comments"`
}

type Query struct {
}

type WriteComment struct {
	PostID               string  `json:"post_id"`
	CommentID            *string `json:"comment_id,omitempty"`
	IsReferenceOnComment bool    `json:"isReferenceOnComment"`
	Text                 string  `json:"text"`
}

type WritePost struct {
	Text     string `json:"text"`
	Comments bool   `json:"comments"`
}
