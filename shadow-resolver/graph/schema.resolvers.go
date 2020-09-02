package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/matiasanaya/gqlgen-scratch/shadow-resolver/graph/generated"
	"github.com/matiasanaya/gqlgen-scratch/shadow-resolver/graph/model"
)

func (r *queryResolver) Todo(ctx context.Context) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *todoResolver) Working(ctx context.Context, obj *model.Todo, in *int) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
