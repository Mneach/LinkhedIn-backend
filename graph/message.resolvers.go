package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/MneachDev/LinkhedIn-backend/graph/generated"
	"github.com/MneachDev/LinkhedIn-backend/graph/model"
	"github.com/MneachDev/LinkhedIn-backend/service"
)

// Sender is the resolver for the sender field.
func (r *messageResolver) Sender(ctx context.Context, obj *model.Message) (*model.User, error) {
	return service.GetUser(r.DB, ctx, obj.SenderID)
}

// SharePost is the resolver for the SharePost field.
func (r *messageResolver) SharePost(ctx context.Context, obj *model.Message) (*model.Post, error) {
	return service.GetSharePost(r.DB, ctx, obj)
}

// ShareProfile is the resolver for the ShareProfile field.
func (r *messageResolver) ShareProfile(ctx context.Context, obj *model.Message) (*model.User, error) {
	return service.GetShareUser(r.DB, ctx, obj)
}

// VideoCall is the resolver for the VideoCall field.
func (r *messageResolver) VideoCall(ctx context.Context, obj *model.Message) (*model.VideoCall, error) {
	return service.GetVideoCall(r.DB, ctx, obj)
}

// AddRoom is the resolver for the addRoom field.
func (r *mutationResolver) AddRoom(ctx context.Context, userID1 string, userID2 string) (*model.Room, error) {
	return service.AddRoom(r.DB, ctx, userID1, userID2)
}

// AddMessage is the resolver for the addMessage field.
func (r *mutationResolver) AddMessage(ctx context.Context, senderID string, text string, imageURL string, roomID string) (*model.Message, error) {
	return service.AddMessage(r.DB, ctx, senderID, text, imageURL, roomID)
}

// AddMessageSharePost is the resolver for the addMessageSharePost field.
func (r *mutationResolver) AddMessageSharePost(ctx context.Context, senderID string, roomID string, sharePostID string) (*model.Message, error) {
	return service.AddMessageSharePost(r.DB, ctx, senderID, roomID, sharePostID)
}

// AddMessageShareProfile is the resolver for the addMessageShareProfile field.
func (r *mutationResolver) AddMessageShareProfile(ctx context.Context, senderID string, roomID string, shareProfileID string) (*model.Message, error) {
	return service.AddMessageShareProfile(r.DB, ctx, senderID, roomID, shareProfileID)
}

// AddMessageVideoCall is the resolver for the addMessageVideoCall field.
func (r *mutationResolver) AddMessageVideoCall(ctx context.Context, senderID string, roomID string, videoCallID string) (*model.Message, error) {
	return service.AddMessageVideoCall(r.DB, ctx, senderID, roomID, videoCallID)
}

// AddVideoCall is the resolver for the addVideoCall field.
func (r *mutationResolver) AddVideoCall(ctx context.Context, time string, title string, date string, duration string, userID1 string, userID2 string) (*model.VideoCall, error) {
	return service.AddVideoCall(r.DB, ctx, time, title, date, duration, userID1, userID2)
}

// Room is the resolver for the room field.
func (r *queryResolver) Room(ctx context.Context, roomID string) (*model.Room, error) {
	return service.Room(r.DB, ctx, roomID)
}

// Rooms is the resolver for the rooms field.
func (r *queryResolver) Rooms(ctx context.Context, userID string) ([]*model.Room, error) {
	return service.Rooms(r.DB, ctx, userID)
}

// User1 is the resolver for the user1 field.
func (r *roomResolver) User1(ctx context.Context, obj *model.Room) (*model.User, error) {
	return service.GetUser(r.DB, ctx, obj.User1ID)
}

// User2 is the resolver for the user2 field.
func (r *roomResolver) User2(ctx context.Context, obj *model.Room) (*model.User, error) {
	return service.GetUser(r.DB, ctx, obj.User2ID)
}

// LastMessage is the resolver for the lastMessage field.
func (r *roomResolver) LastMessage(ctx context.Context, obj *model.Room) (*model.Message, error) {
	return service.LastMessage(r.DB, ctx, obj)
}

// Messages is the resolver for the messages field.
func (r *roomResolver) Messages(ctx context.Context, obj *model.Room) ([]*model.Message, error) {
	return service.Messages(r.DB, ctx, obj)
}

// User1 is the resolver for the user1 field.
func (r *videoCallResolver) User1(ctx context.Context, obj *model.VideoCall) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// User2 is the resolver for the user2 field.
func (r *videoCallResolver) User2(ctx context.Context, obj *model.VideoCall) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Message returns generated.MessageResolver implementation.
func (r *Resolver) Message() generated.MessageResolver { return &messageResolver{r} }

// Room returns generated.RoomResolver implementation.
func (r *Resolver) Room() generated.RoomResolver { return &roomResolver{r} }

// VideoCall returns generated.VideoCallResolver implementation.
func (r *Resolver) VideoCall() generated.VideoCallResolver { return &videoCallResolver{r} }

type messageResolver struct{ *Resolver }
type roomResolver struct{ *Resolver }
type videoCallResolver struct{ *Resolver }
