package service

import (
	"context"
	"fmt"
	"time"

	"github.com/MneachDev/LinkhedIn-backend/authentication"
	"github.com/MneachDev/LinkhedIn-backend/graph/model"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

func CreatePost(db *gorm.DB, ctx context.Context, input model.InputPost) (*model.Post, error) {

	modelPost := &model.Post{
		ID:        uuid.NewString(),
		Text:      input.Text,
		PhotoUrl:  input.PhotoURL,
		VideoUrl:  input.VideoURL,
		CreatedAt: time.Now(),
		SenderId:  input.SenderID,
	}

	// CONNECTED USER POST
	var userIdList []string
	var connections1 []*model.Connection
	var connections2 []*model.Connection

	if err := db.Find(&connections1, "user1_id", input.SenderID).Error; err != nil {
		return nil, err
	}

	if err := db.Find(&connections2, "user2_id", input.SenderID).Error; err != nil {
		return nil, err
	}

	connetions1Ids := lo.Map(connections1, func(connectionData *model.Connection, _ int) string {
		return connectionData.User2ID
	})

	connetions2Ids := lo.Map(connections2, func(connectionData *model.Connection, _ int) string {
		return connectionData.User1ID
	})

	userIdList = append(userIdList, connetions1Ids...)
	userIdList = append(userIdList, connetions2Ids...)
	userIdList = lo.Uniq(userIdList)

	fmt.Println("=====================")
	for _, userId := range userIdList {
		fmt.Println("=====================")
		fmt.Println(userId)
		fmt.Println("=====================")
		AddNotification(db, ctx, userId, input.SenderID, "Create A New Post")
	}

	return modelPost, db.Create(modelPost).Error
}

func GetPosts(db *gorm.DB, ctx context.Context, limit int, offset int) ([]*model.Post, error) {

	var userIdList []string
	userID := authentication.GetJwtValueData(ctx).Userid
	userIdList = append(userIdList, userID)

	// FOLLOWED USER POST
	var follows []*model.Follow

	if err := db.Table("user_follows").Find(&follows, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}

	followIds := lo.Map(follows, func(x *model.Follow, _ int) string {
		return x.FollowID
	})

	userIdList = append(userIdList, followIds...)

	// CONNECTED USER POST
	var connections1 []*model.Connection
	var connections2 []*model.Connection

	if err := db.Find(&connections1, "user1_id", userID).Error; err != nil {
		return nil, err
	}

	if err := db.Find(&connections2, "user2_id", userID).Error; err != nil {
		return nil, err
	}

	connetions1Ids := lo.Map(connections1, func(connectionData *model.Connection, _ int) string {
		return connectionData.User2ID
	})

	connetions2Ids := lo.Map(connections2, func(connectionData *model.Connection, _ int) string {
		return connectionData.User1ID
	})

	userIdList = append(userIdList, connetions1Ids...)
	userIdList = append(userIdList, connetions2Ids...)
	userIdList = lo.Uniq(userIdList)

	var posts []*model.Post
	if err := db.Limit(limit).Offset(offset).Order("created_at desc").Find(&posts, "sender_id IN ?", userIdList).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func LikePost(db *gorm.DB, ctx context.Context, postID string, userID string) (*model.LikePosts, error) {
	modelLikePost := &model.LikePosts{
		PostId: postID,
		UserId: userID,
	}

	return modelLikePost, db.Create(modelLikePost).Error
}

func UnLikePost(db *gorm.DB, ctx context.Context, postID string, userID string) (*model.LikePosts, error) {
	modelLikePost := new(model.LikePosts)

	if err := db.Find(modelLikePost, "post_id = ? AND user_id = ?", postID, userID).Error; err != nil {
		return nil, err
	}

	return modelLikePost, db.Delete(modelLikePost, "post_id = ? AND user_id = ?", postID, userID).Error
}

func GetLikes(db *gorm.DB, ctx context.Context, obj *model.Post) ([]*model.LikePosts, error) {
	var modelLikePost []*model.LikePosts

	if err := db.Find(&modelLikePost, "post_id", obj.ID).Error; err != nil {
		return nil, err
	}

	return modelLikePost, nil
}

func GetComments(db *gorm.DB, ctx context.Context, obj *model.Post) ([]*model.Comment, error) {
	var modelComment []*model.Comment

	if err := db.Find(&modelComment, "post_id = ?", obj.ID).Error; err != nil {
		return nil, err
	}

	return modelComment, nil
}
