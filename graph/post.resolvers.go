package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/MneachDev/LinkhedIn-backend/graph/generated"
	"github.com/MneachDev/LinkhedIn-backend/graph/model"
	"github.com/MneachDev/LinkhedIn-backend/service"
)

// CreatePost is the resolver for the CreatePost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.InputPost) (*model.Post, error) {
	return service.CreatePost(r.DB, ctx, input)
}

// LikePost is the resolver for the LikePost field.
func (r *mutationResolver) LikePost(ctx context.Context, postID string, userID string) (*model.LikePosts, error) {
	return service.LikePost(r.DB, ctx, postID, userID)
}

// UnLikePost is the resolver for the UnLikePost field.
func (r *mutationResolver) UnLikePost(ctx context.Context, postID string, userID string) (*model.LikePosts, error) {
	return service.UnLikePost(r.DB, ctx, postID, userID)
}

// Sender is the resolver for the Sender field.
func (r *postResolver) Sender(ctx context.Context, obj *model.Post) (*model.User, error) {
	return service.GetUser(r.DB, ctx, obj.SenderId)
}

// Likes is the resolver for the Likes field.
func (r *postResolver) Likes(ctx context.Context, obj *model.Post) ([]*model.LikePosts, error) {
	return service.GetLikes(r.DB, ctx, obj)
}

// Comments is the resolver for the Comments field.
func (r *postResolver) Comments(ctx context.Context, obj *model.Post) ([]*model.Comment, error) {
	return service.GetComments(r.DB, ctx, obj)
}

// Shares is the resolver for the Shares field.
func (r *postResolver) Shares(ctx context.Context, obj *model.Post) (int, error) {
	return service.GetTotalShares(r.DB, ctx, obj)
}

// Posts is the resolver for the Posts field.
func (r *queryResolver) Posts(ctx context.Context, limit int, offset int) ([]*model.Post, error) {
	return service.GetPosts(r.DB, ctx, limit, offset)
}

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

type postResolver struct{ *Resolver }
