package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"

	"meigo/gqlgen-todos/graph/generated"
	"meigo/gqlgen-todos/graph/model"
)

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {

	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	var t *model.Todo

	todo, err := t.CreateTodo(gc, input)
	/*
		todo := &model.Todo{
			Text:   input.Text,
			ID:     fmt.Sprintf("T%d", rand.Int()),
			UserID: input.UserID, // fix this line
		}
		r.todos = append(r.todos, todo)
	*/
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	var t *model.Todo

	todos, err := t.FindTodo(gc)

	return todos, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
