package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/MneachDev/LinkhedIn-backend/graph/model"
	"github.com/MneachDev/LinkhedIn-backend/service"
)

// AddJob is the resolver for the addJob field.
func (r *mutationResolver) AddJob(ctx context.Context, title string, companyName string, workplace string, city string, country string, employmentType string, description string) (*model.Job, error) {
	return service.AddJob(r.DB, ctx, title, companyName, workplace, city, country, employmentType, description)
}

// Jobs is the resolver for the Jobs field.
func (r *queryResolver) Jobs(ctx context.Context) ([]*model.Job, error) {
	return service.GetJobs(r.DB, ctx)
}
