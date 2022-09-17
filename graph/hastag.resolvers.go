package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/MneachDev/LinkhedIn-backend/graph/model"
	"github.com/MneachDev/LinkhedIn-backend/service"
)

// AddHastag is the resolver for the addHastag field.
func (r *mutationResolver) AddHastag(ctx context.Context, hastag string) (*model.Hastag, error) {
	return service.AddHastag(r.DB, ctx, hastag)
}

// Hastags is the resolver for the Hastags field.
func (r *queryResolver) Hastags(ctx context.Context) ([]*model.Hastag, error) {
	return service.GetHastags(r.DB, ctx)
}
