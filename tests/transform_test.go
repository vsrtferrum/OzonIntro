package transform

import (
	"testing"

	"github.com/vsrtferrum/OzonIntro/graph/model"
	internalmodel "github.com/vsrtferrum/OzonIntro/internal/model"
	"github.com/vsrtferrum/OzonIntro/internal/transform"
)

func TestPostListTransform(t *testing.T) {
	postList := transform.PostList{
		Id:       1,
		Comments: true,
	}

	expected := internalmodel.PostList{
		Id:       1,
		Comments: true,
	}

	result := postList.PostListTransform()

	if result != expected {
		t.Errorf("PostListTransform() = %v, want %v", result, expected)
	}
}

func TestPostTransform(t *testing.T) {
	post := transform.Post{
		Id:       1,
		Text:     "Test Post",
		Comments: true,
	}

	expected := internalmodel.Post{
		Id:       1,
		Text:     "Test Post",
		Comments: true,
	}

	result := post.PostTransform()

	if result != expected {
		t.Errorf("PostTransform() = %v, want %v", result, expected)
	}
}

func TestCommentTransform(t *testing.T) {
	comment := transform.Comments{
		Comment_id: 1,
		Post_id:    1,
		Text:       "Test Comment",
	}

	expected := internalmodel.Comments{
		Comment_id: 1,
		Post_id:    1,
		Text:       "Test Comment",
	}

	result := comment.CommentTransfrom()

	if result != expected {
		t.Errorf("CommentTransfrom() = %v, want %v", result, expected)
	}
}

func TestToGQLPost(t *testing.T) {
	internalPost := &internalmodel.Post{
		Id:       1,
		Text:     "Test Post",
		Comments: true,
	}

	internalComments := []internalmodel.Comments{
		{
			Comment_id: 1,
			Post_id:    1,
			Text:       "Test Comment",
		},
	}

	expected := &model.Post{
		ID:      "1",
		Content: "Test Post",
		Comments: &model.CommentConnection{
			Edges: []*model.CommentEdge{
				{
					Node: &model.Comment{
						ID:     "1",
						PostID: "1",
						Text:   "Test Comment",
					},
					Cursor: "1",
				},
			},
		},
	}

	result := transform.ToGQLPost(internalPost, &internalComments)

	if result.ID != expected.ID || result.Content != expected.Content || len(result.Comments.Edges) != len(expected.Comments.Edges) {
		t.Errorf("ToGQLPost() = %v, want %v", result, expected)
	}
}

func TestToGQLPostList(t *testing.T) {
	internalPost := &internalmodel.PostList{
		Id:       1,
		Comments: true,
	}

	expected := &model.PostList{
		ID:       "1",
		Comments: true,
	}

	result := transform.ToGQLPostList(internalPost)

	if result.ID != expected.ID || result.Comments != expected.Comments {
		t.Errorf("ToGQLPostList() = %v, want %v", result, expected)
	}
}

func TestToInternalPostList(t *testing.T) {
	gqlPost := &model.PostList{
		ID:       "1",
		Comments: true,
	}

	expected := &internalmodel.PostList{
		Id:       1,
		Comments: true,
	}

	result, err := transform.ToInternalPostList(gqlPost)
	if err != nil {
		t.Fatalf("ToInternalPostList() error = %v", err)
	}

	if result.Id != expected.Id || result.Comments != expected.Comments {
		t.Errorf("ToInternalPostList() = %v, want %v", result, expected)
	}
}

func TestToInternalPost(t *testing.T) {
	gqlPost := &model.Post{
		ID:      "1",
		Content: "Test Post",
	}

	expected := &internalmodel.Post{
		Id:       1,
		Text:     "Test Post",
		Comments: false,
	}

	result, err := transform.ToInternalPost(gqlPost)
	if err != nil {
		t.Fatalf("ToInternalPost() error = %v", err)
	}

	if result.Id != expected.Id || result.Text != expected.Text || result.Comments != expected.Comments {
		t.Errorf("ToInternalPost() = %v, want %v", result, expected)
	}
}

func TestToGQLWritePost(t *testing.T) {
	internalPost := &internalmodel.WritePost{
		Text:     "Test Post",
		Comments: true,
	}

	expected := &model.WritePost{
		Text:     "Test Post",
		Comments: true,
	}

	result := transform.ToGQLWritePost(internalPost)

	if result.Text != expected.Text || result.Comments != expected.Comments {
		t.Errorf("ToGQLWritePost() = %v, want %v", result, expected)
	}
}

func TestToInternalWritePost(t *testing.T) {
	gqlPost := &model.WritePost{
		Text:     "Test Post",
		Comments: true,
	}

	expected := &internalmodel.WritePost{
		Text:     "Test Post",
		Comments: true,
	}

	result := transform.ToInternalWritePost(gqlPost)

	if result.Text != expected.Text || result.Comments != expected.Comments {
		t.Errorf("ToInternalWritePost() = %v, want %v", result, expected)
	}
}

func TestToGQLWriteComment(t *testing.T) {
	internalComment := &internalmodel.WriteComment{
		PostId:               1,
		Comment_id:           2,
		IsReferenceOnComment: true,
		Text:                 "Test Comment",
	}

	expected := &model.WriteComment{
		PostID:               "1",
		CommentID:            stringPtr("2"),
		IsReferenceOnComment: true,
		Text:                 "Test Comment",
	}

	result := transform.ToGQLWriteComment(internalComment)

	if result.PostID != expected.PostID || *result.CommentID != *expected.CommentID || result.IsReferenceOnComment != expected.IsReferenceOnComment || result.Text != expected.Text {
		t.Errorf("ToGQLWriteComment() = %v, want %v", result, expected)
	}
}

func TestToInternalWriteComment(t *testing.T) {
	gqlComment := &model.WriteComment{
		PostID:               "1",
		CommentID:            stringPtr("2"),
		IsReferenceOnComment: true,
		Text:                 "Test Comment",
	}

	expected := &internalmodel.WriteComment{
		PostId:               1,
		Comment_id:           2,
		IsReferenceOnComment: true,
		Text:                 "Test Comment",
	}

	result, err := transform.ToInternalWriteComment(gqlComment)
	if err != nil {
		t.Fatalf("ToInternalWriteComment() error = %v", err)
	}

	if result.PostId != expected.PostId || result.Comment_id != expected.Comment_id || result.IsReferenceOnComment != expected.IsReferenceOnComment || result.Text != expected.Text {
		t.Errorf("ToInternalWriteComment() = %v, want %v", result, expected)
	}
}

func TestToGQLComments(t *testing.T) {
	internalComment := &internalmodel.Comments{
		Comment_id: 1,
		Post_id:    1,
		Text:       "Test Comment",
	}

	expected := &model.Comment{
		ID:     "1",
		PostID: "1",
		Text:   "Test Comment",
	}

	result := transform.ToGQLComments(internalComment)

	if result.ID != expected.ID || result.PostID != expected.PostID || result.Text != expected.Text {
		t.Errorf("ToGQLComments() = %v, want %v", result, expected)
	}
}

func TestToInternalComments(t *testing.T) {
	gqlComment := &model.Comment{
		ID:     "1",
		PostID: "1",
		Text:   "Test Comment",
	}

	expected := &internalmodel.Comments{
		Comment_id: 1,
		Post_id:    1,
		Text:       "Test Comment",
	}

	result, err := transform.ToInternalComments(gqlComment)
	if err != nil {
		t.Fatalf("ToInternalComments() error = %v", err)
	}

	if result.Comment_id != expected.Comment_id || result.Post_id != expected.Post_id || result.Text != expected.Text {
		t.Errorf("ToInternalComments() = %v, want %v", result, expected)
	}
}

func stringPtr(s string) *string {
	return &s
}