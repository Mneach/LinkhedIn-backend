package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/MneachDev/LinkhedIn-backend/graph/model"
	"github.com/MneachDev/LinkhedIn-backend/service"
)

// AddEducation is the resolver for the addEducation field.
func (r *mutationResolver) AddEducation(ctx context.Context, input model.InputEducation) (*model.Education, error) {
	return service.AddEducation(r.DB, ctx, input)
}

// UpdateEducation is the resolver for the updateEducation field.
func (r *mutationResolver) UpdateEducation(ctx context.Context, id string, input model.InputEducation) (*model.Education, error) {
	return service.UpdateEducation(r.DB, ctx, id, input)
}

// DeleteEducation is the resolver for the deleteEducation field.
func (r *mutationResolver) DeleteEducation(ctx context.Context, id string) (*model.Education, error) {
	return service.DeleteEducation(r.DB, ctx, id)
}

// Educations is the resolver for the Educations field.
func (r *queryResolver) Educations(ctx context.Context) ([]*model.Education, error) {
	return service.Educations(r.DB, ctx)
}
