package service

import (
	"context"

	"github.com/MneachDev/LinkhedIn-backend/graph/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func User1(db *gorm.DB, ctx context.Context, obj *model.Connection) (*model.User, error) {
	modelUser, err := GetUser(db, ctx, obj.User1ID)

	if err != nil {
		return nil, err
	}

	return modelUser, nil
}

// User2 is the resolver for the user2 field.
func User2(db *gorm.DB, ctx context.Context, obj *model.Connection) (*model.User, error) {
	modelUser, err := GetUser(db, ctx, obj.User2ID)

	if err != nil {
		return nil, err
	}

	return modelUser, nil
}

func FromUser(db *gorm.DB, ctx context.Context, obj *model.ConnectRequest) (*model.User, error) {
	modelUser, err := GetUser(db, ctx, obj.FromUserID)

	if err != nil {
		return nil, err
	}

	return modelUser, nil
}

// ToUser is the resolver for the toUser field.
func ToUser(db *gorm.DB, ctx context.Context, obj *model.ConnectRequest) (*model.User, error) {
	modelUser, err := GetUser(db, ctx, obj.ToUserID)

	if err != nil {
		return nil, err
	}

	return modelUser, nil
}

func AddConnection(db *gorm.DB, ctx context.Context, user1id string, user2id string) (*model.Connection, error) {
	modelConnection := &model.Connection{
		ID:      uuid.NewString(),
		User1ID: user1id,
		User2ID: user2id,
	}

	return modelConnection, db.Create(modelConnection).Error
}

func AddConnectRequest(db *gorm.DB, ctx context.Context, fromUserID string, toUserID string, message string) (*model.ConnectRequest, error) {
	modelConnectRequest := &model.ConnectRequest{
		ID:         uuid.NewString(),
		FromUserID: fromUserID,
		ToUserID:   toUserID,
		Message:    message,
	}

	return modelConnectRequest, db.Create(modelConnectRequest).Error
}

func DeleteConnectRequest(db *gorm.DB, ctx context.Context, fromUserID string, toUserID string) (*model.ConnectRequest, error) {

	modelConnectRequest := new(model.ConnectRequest)

	if err := db.Find(&modelConnectRequest, "from_user_id = ? AND to_user_id = ?", fromUserID, toUserID).Error; err != nil {
		return nil, err
	}

	return modelConnectRequest, db.Delete(modelConnectRequest).Error
}
