package service

import (
	"context"
	"time"

	"github.com/MneachDev/LinkhedIn-backend/graph/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func AddMessage(db *gorm.DB, ctx context.Context, senderID string, text string, imageURL string, roomID string) (*model.Message, error) {

	modelMessage := &model.Message{
		ID:        uuid.NewString(),
		Text:      text,
		ImageURL:  imageURL,
		SenderID:  senderID,
		RoomID:    roomID,
		CreatedAt: time.Now(),
	}

	return modelMessage, db.Create(modelMessage).Error
}

func AddMessageSharePost(db *gorm.DB, ctx context.Context, senderID string, roomID string, sharePostId string) (*model.Message, error) {

	modelMessage := &model.Message{
		ID:          uuid.NewString(),
		SenderID:    senderID,
		RoomID:      roomID,
		SharePostID: &sharePostId,
		CreatedAt:   time.Now(),
	}

	return modelMessage, db.Create(modelMessage).Error
}

func AddMessageShareProfile(db *gorm.DB, ctx context.Context, senderID string, roomID string, shareProfileId string) (*model.Message, error) {

	modelMessage := &model.Message{
		ID:             uuid.NewString(),
		SenderID:       senderID,
		RoomID:         roomID,
		ShareProfileID: &shareProfileId,
		CreatedAt:      time.Now(),
	}

	return modelMessage, db.Create(modelMessage).Error
}

func AddMessageVideoCall(db *gorm.DB, ctx context.Context, senderID string, roomID string, videoCallID string) (*model.Message, error) {

	modelMessage := &model.Message{
		ID:          uuid.NewString(),
		SenderID:    senderID,
		RoomID:      roomID,
		VideoCallID: &videoCallID,
		CreatedAt:   time.Now(),
	}

	return modelMessage, db.Create(modelMessage).Error
}

// AddRoom is the resolver for the addRoom field.
func AddRoom(db *gorm.DB, ctx context.Context, userID1 string, userID2 string) (*model.Room, error) {
	modelRoom := &model.Room{
		ID:        uuid.NewString(),
		User1ID:   userID1,
		User2ID:   userID2,
		CreatedAt: time.Now(),
	}

	return modelRoom, db.Create(modelRoom).Error
}

// Room is the resolver for the room field.
func Room(db *gorm.DB, ctx context.Context, roomID string) (*model.Room, error) {
	modelRoom := new(model.Room)

	return modelRoom, db.Order("created_at desc").Find(modelRoom, "id = ?", roomID).Error
}

// Rooms is the resolver for the rooms field.
func Rooms(db *gorm.DB, ctx context.Context, userID string) ([]*model.Room, error) {
	var modelRooms []*model.Room

	if err := db.Order("created_at desc").Where("user1_id = ?", userID).Or("user2_id = ?", userID).Find(&modelRooms).Error; err != nil {
		return nil, err
	}

	return modelRooms, nil
}

// LastMessage is the resolver for the lastMessage field.
func LastMessage(db *gorm.DB, ctx context.Context, obj *model.Room) (*model.Message, error) {
	modelMessage := new(model.Message)

	if err := db.Order("created_at desc").Limit(1).Find(&modelMessage, "room_id = ?", obj.ID).Error; err != nil {
		return nil, err
	}

	return modelMessage, nil
}

// Messages is the resolver for the messages field.
func Messages(db *gorm.DB, ctx context.Context, obj *model.Room) ([]*model.Message, error) {
	var modelMessages []*model.Message

	if err := db.Order("created_at asc").Find(&modelMessages, "room_id = ?", obj.ID).Error; err != nil {
		return nil, err
	}

	return modelMessages, nil
}

func GetSharePost(db *gorm.DB, ctx context.Context, obj *model.Message) (*model.Post, error) {
	modelPost := new(model.Post)

	if err := db.Find(&modelPost, "id = ?", obj.SharePostID).Error; err != nil {
		return nil, err
	}

	return modelPost, nil
}

func GetShareUser(db *gorm.DB, ctx context.Context, obj *model.Message) (*model.User, error) {
	modelUser := new(model.User)

	if obj.ShareProfileID == nil {
		return modelUser, nil
	}

	if err := db.Find(&modelUser, "id = ?", obj.ShareProfileID).Error; err != nil {
		return nil, err
	}

	return modelUser, nil
}

func AddVideoCall(db *gorm.DB, ctx context.Context, time string, title string, date string, duration string, userID1 string, userID2 string) (*model.VideoCall, error) {
	modelVideoCall := &model.VideoCall{
		ID:       uuid.NewString(),
		Date:     date,
		Title:    title,
		Time:     time,
		Duration: duration,
		User1ID:  userID1,
		User2ID:  userID2,
	}

	return modelVideoCall, db.Create(modelVideoCall).Error
}

func GetVideoCall(db *gorm.DB, ctx context.Context, obj *model.Message) (*model.VideoCall, error) {
	modelVideocall := new(model.VideoCall)

	if err := db.Find(&modelVideocall, "id = ?", obj.VideoCallID).Error; err != nil {
		return nil, err
	}

	return modelVideocall, nil
}
