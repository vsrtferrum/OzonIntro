package transform

import (
	"strconv"

	"github.com/vsrtferrum/OzonIntro/graph/model"
	internalmodel "github.com/vsrtferrum/OzonIntro/internal/model"
)


func ToGQLPost(internalPost *internalmodel.Post, internalComments *[]internalmodel.Comments) *model.Post {
    gqlPost := &model.Post{
        ID:      strconv.FormatUint(internalPost.Id, 10),
        Content: internalPost.Text,
    }

    if internalComments == nil {
        gqlPost.Comments = &model.CommentConnection{
            Edges: []*model.CommentEdge{},
        }
        return gqlPost
    }

    gqlComments := make([]*model.Comment, 0, len(*internalComments))
    for _, internalComment := range *internalComments {
        gqlComments = append(gqlComments, ToGQLComments(&internalComment))
    }

    gqlPost.Comments = &model.CommentConnection{
        Edges: make([]*model.CommentEdge, len(gqlComments)),
    }

    for i, comment := range gqlComments {
        gqlPost.Comments.Edges[i] = &model.CommentEdge{
            Node:   comment,
            Cursor: comment.ID, 
        }
    }

    return gqlPost
}

func ToGQLPostList(internalPost *internalmodel.PostList) *model.PostList {
    return &model.PostList{
        ID:       strconv.FormatUint(internalPost.Id, 10), 
        Comments: internalPost.Comments,
    }
}

func ToInternalPostList(gqlPost *model.PostList) (*internalmodel.PostList, error) {
    id, err := strconv.ParseUint(gqlPost.ID, 10, 64)
    if err != nil {
        return nil, err
    }

    return &internalmodel.PostList{
        Id:       id,
        Comments: gqlPost.Comments,
    }, nil
}




func ToInternalPost(gqlPost *model.Post) (*internalmodel.Post, error) {
    id, err := strconv.ParseUint(gqlPost.ID, 10, 64)
    if err != nil {
        return nil, err
    }

    return &internalmodel.Post{
        Id:       id,
        Text:     gqlPost.Content,
        Comments: false, 
    }, nil
}


func ToGQLWritePost(internalPost *internalmodel.WritePost) *model.WritePost {
    return &model.WritePost{
        Text:     internalPost.Text,
        Comments: internalPost.Comments,
    }
}


func ToInternalWritePost(gqlPost *model.WritePost) *internalmodel.WritePost {
    return &internalmodel.WritePost{
        Text:     gqlPost.Text,
        Comments: gqlPost.Comments,
    }
}

func ToGQLWriteComment(internalComment *internalmodel.WriteComment) *model.WriteComment {
    return &model.WriteComment{
        PostID:               strconv.FormatUint(internalComment.PostId, 10),
        CommentID:            uint64ToStringPtr(internalComment.Comment_id),
        IsReferenceOnComment: internalComment.IsReferenceOnComment,
        Text:                 internalComment.Text,
    }
}

func ToInternalWriteComment(gqlComment *model.WriteComment) (*internalmodel.WriteComment, error) {
    postID, err := strconv.ParseUint(gqlComment.PostID, 10, 64)
    if err != nil {
        return nil, err
    }

    var commentID uint64
    if gqlComment.CommentID != nil {
        commentID, err = strconv.ParseUint(*gqlComment.CommentID, 10, 64)
        if err != nil {
            return nil, err
        }
    }

    return &internalmodel.WriteComment{
        PostId:               postID,
        Comment_id:           commentID,
        IsReferenceOnComment: gqlComment.IsReferenceOnComment,
        Text:                 gqlComment.Text,
    }, nil
}


func ToGQLComments(internalComment *internalmodel.Comments) *model.Comment {
    return &model.Comment{
        ID:     strconv.FormatUint(internalComment.Comment_id, 10),
        PostID: strconv.FormatUint(internalComment.Post_id, 10),
        Text:   internalComment.Text,
    }
}


func ToInternalComments(gqlComment *model.Comment) (*internalmodel.Comments, error) {
    commentID, err := strconv.ParseUint(gqlComment.ID, 10, 64)
    if err != nil {
        return nil, err
    }

    postID, err := strconv.ParseUint(gqlComment.PostID, 10, 64)
    if err != nil {
        return nil, err
    }

    return &internalmodel.Comments{
        Comment_id: commentID,
        Post_id:    postID,
        Text:       gqlComment.Text,
    }, nil
}


func uint64ToStringPtr(value uint64) *string {
    if value == 0 {
        return nil
    }
    str := strconv.FormatUint(value, 10)
    return &str
}