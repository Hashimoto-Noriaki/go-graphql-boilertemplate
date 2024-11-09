package graph

import (
	"context"
	"github.com/Hashimoto-Noriaki/go-graphql-boilertemplate/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	return &model.Todo{
		ID:   "TODO-3",
		Text: input.Text,
		User: &model.User{
			ID:   input.UserID,
			Name: "Mike",
		},
	}, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return []*model.Todo{
		{
			ID:   "TODO-1",
			Text: "やること1",
			User: &model.User{
				ID:   "User-1",
				Name: "Mookie",
			},
			Done: true,
		},
		{
			ID:   "TODO-2",
			Text: "やること2",
			User: &model.User{
				ID:   "User-1",
				Name: "Harry",
			},
			Done: false,
		},
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
