package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.64

import (
	"context"
	"fmt"
	"workout-training-api/internal/graphql/graph/model"
)

// GetPing is the resolver for the getPing field.
func (r *mutationResolver) GetPing(ctx context.Context) (*model.Ping, error) {
	panic(fmt.Errorf("not implemented: GetPing - getPing"))
}

// GetPing is the resolver for the getPing field.
func (r *queryResolver) GetPing(ctx context.Context) (*model.Ping, error) {
	panic(fmt.Errorf("not implemented: GetPing - getPing"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
