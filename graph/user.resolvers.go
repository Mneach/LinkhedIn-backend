package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/MneachDev/LinkhedIn-backend/graph/generated"
	"github.com/MneachDev/LinkhedIn-backend/graph/model"
	"github.com/MneachDev/LinkhedIn-backend/service"
)

// RegisterUser is the resolver for the registerUser field.
func (r *mutationResolver) RegisterUser(ctx context.Context, input model.InputRegisterUser) (*model.User, error) {
	return service.RegisterUser(r.DB, ctx, input)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input model.InputUpdateUser) (*model.User, error) {
	return service.UpdateUser(r.DB, ctx, id, input)
}

// UpdatePasswordUser is the resolver for the updatePasswordUser field.
func (r *mutationResolver) UpdatePasswordUser(ctx context.Context, id string, password string) (*model.User, error) {
	return service.UpdatePasswordUser(r.DB, ctx, id, password)
}

// RegisterResetPassword is the resolver for the registerResetPassword field.
func (r *mutationResolver) RegisterResetPassword(ctx context.Context, email string) (*model.ResetPasswordAccount, error) {
	return service.RegisterResetPassword(r.DB, ctx, email)
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*model.User, error) {
	return service.DeleteUser(r.DB, ctx, id)
}

// FollowUser is the resolver for the FollowUser field.
func (r *mutationResolver) FollowUser(ctx context.Context, id1 string, id2 string) (interface{}, error) {
	return service.FollowUser(r.DB, ctx, id1, id2)
}

// UnFollowUser is the resolver for the UnFollowUser field.
func (r *mutationResolver) UnFollowUser(ctx context.Context, id1 string, id2 string) (interface{}, error) {
	return service.UnFollowUser(r.DB, ctx, id1, id2)
}

// VisitUser is the resolver for the VisitUser field.
func (r *mutationResolver) VisitUser(ctx context.Context, id1 string, id2 string) (interface{}, error) {
	return service.VisitUser(r.DB, ctx, id1, id2)
}

// User is the resolver for the User field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return service.GetUser(r.DB, ctx, id)
}

// Users is the resolver for the Users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return service.GetUsers(r.DB, ctx)
}

// UserSuggestion is the resolver for the UserSuggestion field.
func (r *queryResolver) UserSuggestion(ctx context.Context, userID string) ([]*model.User, error) {
	return service.GetUserSuggestion(r.DB, ctx, userID)
}

// UserByEmail is the resolver for the UserByEmail field.
func (r *queryResolver) UserByEmail(ctx context.Context, email string) (*model.User, error) {
	return service.GetUserByEmail(r.DB, ctx, email)
}

// CheckEmailUser is the resolver for the CheckEmailUser field.
func (r *queryResolver) CheckEmailUser(ctx context.Context, email string) (*model.User, error) {
	return service.CheckEmailUser(r.DB, ctx, email)
}

// UserByActivationID is the resolver for the UserByActivationId field.
func (r *queryResolver) UserByActivationID(ctx context.Context, activationID string) (*model.User, error) {
	return service.GetUserByActivationID(r.DB, ctx, activationID)
}

// UserByResetPasswordID is the resolver for the UserByResetPasswordId field.
func (r *queryResolver) UserByResetPasswordID(ctx context.Context, resetPasswordID string) (*model.User, error) {
	return service.GetUserByResetPasswordID(r.DB, ctx, resetPasswordID)
}

// Login is the resolver for the Login field.
func (r *queryResolver) Login(ctx context.Context, input model.InputLogin) (interface{}, error) {
	return service.Login(r.DB, ctx, input)
}

// Protected is the resolver for the protected field.
func (r *queryResolver) Protected(ctx context.Context) (string, error) {
	return "SUCCESS", nil
}

// Visits is the resolver for the Visits field.
func (r *userResolver) Visits(ctx context.Context, obj *model.User) ([]*model.Visit, error) {
	return service.GetVisits(r.DB, ctx, obj)
}

// Follows is the resolver for the Follows field.
func (r *userResolver) Follows(ctx context.Context, obj *model.User) ([]*model.Follow, error) {
	return service.GetFollows(r.DB, ctx, obj)
}

// Blocks is the resolver for the Blocks field.
func (r *userResolver) Blocks(ctx context.Context, obj *model.User) ([]*model.Block, error) {
	return service.GetBlocks(r.DB, ctx, obj)
}

// Connections is the resolver for the Connections field.
func (r *userResolver) Connections(ctx context.Context, obj *model.User) ([]*model.Connection, error) {
	return service.GetConnections(r.DB, ctx, obj)
}

// ConnectRequests is the resolver for the ConnectRequests field.
func (r *userResolver) ConnectRequests(ctx context.Context, obj *model.User) ([]*model.ConnectRequest, error) {
	return service.GetConnectRequests(r.DB, ctx, obj)
}

// Experiences is the resolver for the Experiences field.
func (r *userResolver) Experiences(ctx context.Context, obj *model.User) ([]*model.Experience, error) {
	return service.GetExperienceUser(r.DB, ctx, obj)
}

// Educations is the resolver for the Educations field.
func (r *userResolver) Educations(ctx context.Context, obj *model.User) ([]*model.Education, error) {
	return service.GetEducations(r.DB, ctx, obj)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
