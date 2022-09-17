package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/MneachDev/LinkhedIn-backend/graph/generated"
	"github.com/MneachDev/LinkhedIn-backend/graph/model"
	"github.com/MneachDev/LinkhedIn-backend/service"
)

// Commenter is the resolver for the Commenter field.
func (r *commentResolver) Commenter(ctx context.Context, obj *model.Comment) (*model.User, error) {
	return service.GetUser(r.DB, ctx, obj.CommenterID)
}

// Replies is the resolver for the Replies field.
func (r *commentResolver) Replies(ctx context.Context, obj *model.Comment) ([]*model.Comment, error) {
	return service.GetReplied(r.DB, ctx, obj)
}

// Likes is the resolver for the Likes field.
func (r *commentResolver) Likes(ctx context.Context, obj *model.Comment) ([]*model.LikeComment, error) {
	return service.GetLikesComment(r.DB, ctx, obj)
}

// User is the resolver for the User field.
func (r *likeCommentResolver) User(ctx context.Context, obj *model.LikeComment) (*model.User, error) {
	return service.GetUser(r.DB, ctx, obj.UserID)
}

// AddComment is the resolver for the addComment field.
func (r *mutationResolver) AddComment(ctx context.Context, postID string, commenterID string, comment string) (*model.Comment, error) {
	return service.AddComment(r.DB, ctx, postID, commenterID, comment)
}

// AddLikeComment is the resolver for the addLikeComment field.
func (r *mutationResolver) AddLikeComment(ctx context.Context, commentID string, userID string) (*model.LikeComment, error) {
	return service.AddLikeComment(r.DB, ctx, commentID, userID)
}

// DeleteLikeComment is the resolver for the deleteLikeComment field.
func (r *mutationResolver) DeleteLikeComment(ctx context.Context, commentID string, userID string) (*model.LikeComment, error) {
	return service.DeleteLikeComment(r.DB, ctx, commentID, userID)
}

// AddReply is the resolver for the addReply field.
func (r *mutationResolver) AddReply(ctx context.Context, commenterID string, postID string, replyToCommentID string, comment string) (*model.Comment, error) {
	return service.AddReply(r.DB, ctx, commenterID, postID, replyToCommentID, comment)
}

// PostComment is the resolver for the postComment field.
func (r *queryResolver) PostComment(ctx context.Context, id string) (*model.Comment, error) {
	return service.GetComment(r.DB, ctx, id)
}

// RepliedToComments is the resolver for the repliedToComments field.
func (r *queryResolver) RepliedToComments(ctx context.Context, limit int, offset int, commentID string) ([]*model.Comment, error) {
	return service.GetRepliedToComments(r.DB, ctx, limit, offset, commentID)
}

// PostComments is the resolver for the postComments field.
func (r *queryResolver) PostComments(ctx context.Context, limit int, offset int, postID string) ([]*model.Comment, error) {
	return service.GetPostComments(r.DB, ctx, limit, offset, postID)
}

// Comment returns generated.CommentResolver implementation.
func (r *Resolver) Comment() generated.CommentResolver { return &commentResolver{r} }

// LikeComment returns generated.LikeCommentResolver implementation.
func (r *Resolver) LikeComment() generated.LikeCommentResolver { return &likeCommentResolver{r} }

type commentResolver struct{ *Resolver }
type likeCommentResolver struct{ *Resolver }
