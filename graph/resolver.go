package graph

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//rep "main/repository"

import (
	"context"
	"fmt"
	"main/graph/model"
	rep "main/repository"
)

type Resolver struct {
	Store rep.StoreInterface
}

type queryResolver struct{ *Resolver }

type mutationResolver struct{ *Resolver }

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func (r *queryResolver) Posts(ctx context.Context, limit *int, offset *int) ([]*model.Post, error) {
	return r.Store.GetPosts()
	//return nil, nil
}

func (r *mutationResolver) CreatePost(ctx context.Context, in model.NewPost) (*model.Post, error) {
	id, err := r.Store.CreatePost(in)
	ans := model.Post{ID: id, Title: in.Title + fmt.Sprint(id), Text: in.Text, AuthorID: in.AuthorID, IsCommentsUnabled: in.IsCommentsUnabled}
	if err != nil {
		return nil, err
	}
	return &ans, nil
}
