package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"meigo/gqlgen-todos/graph/generated"
	"meigo/gqlgen-todos/graph/model"

	"github.com/gin-gonic/gin"
)

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

	/*
		// Print a formatted string
		graphql.AddErrorf(ctx, "Error %d", 1)

		// Pass an existing error out
		graphql.AddError(ctx, gqlerror.Errorf("zzzzzt"))

		// Or fully customize the error
		graphql.AddError(ctx, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: "A descriptive error message",
			Extensions: map[string]interface{}{
				"code": "10-4",
			},
		})

		// And you can still return an error if you need
		return todo, gqlerror.Errorf("BOOM! Headshot")
	*/
}

func (r *mutationResolver) DeleteUser(ctx context.Context, userID string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
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
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	var user *model.User
	user, err = user.FindUser(gc, obj)
	return user, err
	//return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
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
