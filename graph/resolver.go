package graph

// THIS CODE WILL BE UPDATED WITH SCHEMA CHANGES. PREVIOUS IMPLEMENTATION FOR SCHEMA CHANGES WILL BE KEPT IN THE COMMENT SECTION. IMPLEMENTATION FOR UNCHANGED SCHEMA WILL BE KEPT.

import (
	"context"
	"strconv"

	"github.com/vsrtferrum/OzonIntro/graph/model"
	"github.com/vsrtferrum/OzonIntro/internal/transform"
	"github.com/vsrtferrum/OzonIntro/internal/workers"
)

type Resolver struct{
	Workers *workers.ConcurrentModule
}

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, data model.WritePost) (string, error) {
	id, err := r.Workers.AddPost(transform.ToInternalWritePost(&data))
	return strconv.FormatUint(id, 10), err
}

// AddComment is the resolver for the addComment field.
func (r *mutationResolver) AddComment(ctx context.Context, data model.WriteComment) (string, error) {
	internalData, err := transform.ToInternalWriteComment(&data)
	if err !=nil{
		return "", err 
	} 
	id, err := r.Workers.AddComment(internalData)
	return strconv.FormatUint(id, 10), err
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*model.PostList, error) {
	data, err := 	r.Workers.GetPosts()
	if err != nil{
		return nil , err 
	}
	res  := make([]*model.PostList, 0, len(*data))
	for _, val := range *data{
		res = append(res,  transform.ToGQLPostList(&val))
	}
	return res, nil
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	idNum, err  := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil , err 
	}
	postInternal, commentsInternal, err :=   r.Workers.GetPost(idNum)
	if err != nil {
		return nil , err 
	}
	post := transform.ToGQLPost(postInternal, commentsInternal)
	return post, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
