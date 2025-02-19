package errors

import "errors"

var (
	ErrNotInitPostList            = errors.New("error post list is nil")
	ErrNotInitCommentList         = errors.New("error comment list is nil")
	ErrPostNotExisted             = errors.New("error post is not existed")
	ErrReferenceCommentNotExisted = errors.New("error reference comment is not existed")
)
