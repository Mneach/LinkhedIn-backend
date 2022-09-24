package service

import (
	"context"
	"time"

	"github.com/MneachDev/LinkhedIn-backend/graph/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func AddNotification(db *gorm.DB, ctx context.Context, toUserID string, fromUserID string, message string) (*model.Notification, error) {
	modelNotification := &model.Notification{
		ID:         uuid.NewString(),
		FromUserID: fromUserID,
		ToUserID:   toUserID,
		Message:    message,
		CreatedAt:  time.Now(),
	}

	return modelNotification, db.Create(modelNotification).Error
}

// FromUser is the resolver for the fromUser field.
func GetFromUserNotification(db *gorm.DB, ctx context.Context, obj *model.Notification) (*model.User, error) {
	return GetUser(db, ctx, obj.FromUserID)
}

// ToUser is the resolver for the toUser field.
func GetToUserNotification(db *gorm.DB, ctx context.Context, obj *model.Notification) (*model.User, error) {
	return GetUser(db, ctx, obj.ToUserID)
}

// UserNotification is the resolver for the userNotification field.
func GetUserNotification(db *gorm.DB, ctx context.Context, toUserID string) ([]*model.Notification, error) {
	var modelNotifications []*model.Notification

	if err := db.Order("created_at desc").Find(&modelNotifications, "to_user_id = ?", toUserID).Error; err != nil {
		return nil, err
	}

	return modelNotifications, nil
}
