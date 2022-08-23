package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/MneachDev/LinkhedIn-backend/graph/model"
	"github.com/MneachDev/LinkhedIn-backend/service"
)

// AddExperience is the resolver for the addExperience field.
func (r *mutationResolver) AddExperience(ctx context.Context, input model.InputExperience) (*model.Experience, error) {
	return service.AddExperience(r.DB, ctx, input)
}

// UpdateExperience is the resolver for the updateExperience field.
func (r *mutationResolver) UpdateExperience(ctx context.Context, id string, input model.InputExperience) (*model.Experience, error) {
	return service.UpdateExperience(r.DB, ctx, id, input)
}

// DeleteExperience is the resolver for the deleteExperience field.
func (r *mutationResolver) DeleteExperience(ctx context.Context, id string) (*model.Experience, error) {
	return service.DeleteExperience(r.DB, ctx, id)
}

// Experiences is the resolver for the Experiences field.
func (r *queryResolver) Experiences(ctx context.Context) ([]*model.Experience, error) {
	return service.Experiences(r.DB, ctx)
}
