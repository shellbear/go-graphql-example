package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/shellbear/go-graphql-example/ent"
	"github.com/shellbear/go-graphql-example/ent/user"
	"github.com/shellbear/go-graphql-example/graph/generated"
	"github.com/shellbear/go-graphql-example/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*ent.Todo, error) {
	// Fetch User based on UserName input parameter.
	usr, err := r.Client.User.Query().Where(
		user.NameEQ(input.UserName),
	).First(ctx)

	// Check errors
	if err != nil {
		// If error is not not found, return the error.
		if !ent.IsNotFound(err) {
			return nil, err
		}

		// Otherwise create the User.
		usr, err = r.Client.User.Create().
			SetName(input.UserName).
			Save(ctx)
		if err != nil {
			return nil, err
		}
	}

	// Create the new Todo.
	return r.Client.Todo.Create().
		SetText(input.Text). // Set text
		SetUser(usr).        // Set User
		Save(ctx)            // Save it
}

func (r *queryResolver) Todos(ctx context.Context) ([]*ent.Todo, error) {
	// Returns all the Todo.
	return r.Client.Todo.Query().All(ctx)
}

func (r *queryResolver) UserTodos(ctx context.Context, name string) ([]*ent.Todo, error) {
	// Query all Users in database.
	return r.Client.User.Query().
		Where(
			// Filter only User that exactly match the name parameter.
			user.NameEQ(name),
		).
		QueryTodos(). // Query all Todos of the found User.
		All(ctx)      // Returns all Todos.
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
