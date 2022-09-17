package service

import (
	"context"
	"time"

	"github.com/MneachDev/LinkhedIn-backend/graph/model"
	"github.com/google/uuid"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/gorm"
)

// AddComment is the resolver for the addComment field.
func AddComment(db *gorm.DB, ctx context.Context, postID string, commenterID string, comment string) (*model.Comment, error) {
	modelComment := &model.Comment{
		ID:          uuid.NewString(),
		PostID:      postID,
		CommenterID: commenterID,
		Comment:     comment,
		CreatedAt:   time.Now(),
	}

	return modelComment, db.Create(modelComment).Error
}

func AddReply(db *gorm.DB, ctx context.Context, commenterID string, postID string, reply_to_comment_id string, comment string) (*model.Comment, error) {
	modelComment := &model.Comment{
		ID:               uuid.NewString(),
		PostID:           postID,
		CommenterID:      commenterID,
		Comment:          comment,
		ReplyToCommentID: &reply_to_comment_id,
		CreatedAt:        time.Now(),
	}

	return modelComment, db.Create(modelComment).Error
}

func GetPostComments(db *gorm.DB, ctx context.Context, limit int, offset int, postID string) ([]*model.Comment, error) {
	var modelComments []*model.Comment

	if err := db.Limit(limit).Offset(offset).Order("created_at desc").Find(&modelComments, "post_id = ? AND reply_to_comment_id IS NULL", postID).Error; err != nil {
		return nil, err
	}

	if len(modelComments) == 0 {
		return nil, gqlerror.Errorf("No More Data")
	}

	return modelComments, nil
}

// // Replies is the resolver for the Replies field.
func GetRepliedToComments(db *gorm.DB, ctx context.Context, limit int, offset int, commentId string) ([]*model.Comment, error) {
	var modelComments []*model.Comment

	if err := db.Limit(limit).Offset(offset).Order("created_at desc").Find(&modelComments, "reply_to_comment_id = ?", commentId).Error; err != nil {
		return nil, err
	}

	return modelComments, nil
}

func GetReplied(db *gorm.DB, ctx context.Context, obj *model.Comment) ([]*model.Comment, error) {
	var modelComments []*model.Comment

	if err := db.Find(&modelComments, "reply_to_comment_id = ?", obj.ID).Error; err != nil {
		return nil, err
	}

	return modelComments, nil
}

func GetLikesComment(db *gorm.DB, ctx context.Context, obj *model.Comment) ([]*model.LikeComment, error) {
	var modelLikes []*model.LikeComment

	if err := db.Find(&modelLikes, "comment_id", obj.ID).Error; err != nil {
		return nil, err
	}

	return modelLikes, nil
}

func AddLikeComment(db *gorm.DB, ctx context.Context, commentID string, userID string) (*model.LikeComment, error) {
	modelLikeComment := &model.LikeComment{
		ID:        uuid.NewString(),
		CommentID: commentID,
		UserID:    userID,
	}

	modelLikeCommentDb := new(model.LikeComment)
	if err := db.Find(modelLikeComment, "user_id = ? AND comment_id = ?", userID, commentID).Error; err != nil {
		return nil, err
	}

	if modelLikeCommentDb.ID != "" {
		return nil, gqlerror.Errorf("You cannot like twice")
	}

	return modelLikeComment, db.Create(modelLikeComment).Error
}

// DeleteLikeComment is the resolver for the deleteLikeComment field.
func DeleteLikeComment(db *gorm.DB, ctx context.Context, commentId string, userId string) (*model.LikeComment, error) {

	modelLikeComment := new(model.LikeComment)

	if err := db.Find(modelLikeComment, "comment_id = ? AND user_id = ?", commentId, userId).Error; err != nil {
		return nil, err
	}

	return modelLikeComment, db.Delete(modelLikeComment).Error
}

func GetComment(db *gorm.DB, ctx context.Context, id string) (*model.Comment, error) {
	modelComment := new(model.Comment)

	if err := db.Find(modelComment, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return modelComment, nil
}
