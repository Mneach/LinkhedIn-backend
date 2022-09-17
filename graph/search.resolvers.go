package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/MneachDev/LinkhedIn-backend/graph/generated"
	"github.com/MneachDev/LinkhedIn-backend/graph/model"
	"github.com/MneachDev/LinkhedIn-backend/service"
)

// Search is the resolver for the Search field.
func (r *queryResolver) Search(ctx context.Context, keyword string, limit int, offset int) (*model.Search, error) {
	return service.Search(r.DB, ctx, keyword, limit, offset)
}

// SearchHastag is the resolver for the SearchHastag field.
func (r *queryResolver) SearchHastag(ctx context.Context, keyword string, limit int, offset int) (*model.Search, error) {
	return service.SearchHastag(r.DB, ctx, keyword, limit, offset)
}

// Users is the resolver for the Users field.
func (r *searchResolver) Users(ctx context.Context, obj *model.Search) ([]*model.User, error) {
	return service.GetUserSearch(r.DB, ctx, obj)
}

// Posts is the resolver for the Posts field.
func (r *searchResolver) Posts(ctx context.Context, obj *model.Search) ([]*model.Post, error) {
	return service.GetPostSearch(r.DB, ctx, obj)
}

// Search returns generated.SearchResolver implementation.
func (r *Resolver) Search() generated.SearchResolver { return &searchResolver{r} }

type searchResolver struct{ *Resolver }
