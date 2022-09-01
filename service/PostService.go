package service

import (
	"context"
	"fmt"

	"github.com/MneachDev/LinkhedIn-backend/authentication"
	"github.com/MneachDev/LinkhedIn-backend/graph/model"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

func CreatePost(db *gorm.DB, ctx context.Context, input model.InputPost) (*model.Post, error) {
	modelPost := &model.Post{
		ID:       uuid.NewString(),
		Text:     input.Text,
		PhotoUrl: input.PhotoURL,
		VideoUrl: input.VideoURL,
		SenderId: input.SenderID,
	}

	return modelPost, db.Create(modelPost).Error
}

func GetPosts(db *gorm.DB, ctx context.Context, limit int, offset int) ([]*model.Post, error) {

	var userIdList []string
	userID := authentication.GetJwtValueData(ctx).Userid
	userIdList = append(userIdList, userID)

	var follows []*model.Follow

	if err := db.Table("user_follows").Find(&follows, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}

	followIds := lo.Map(follows, func(x *model.Follow, _ int) string {
		return x.FollowID
	})

	userIdList = append(userIdList, followIds...)
	userIdList = lo.Uniq(userIdList)

	fmt.Println("----------------")
	fmt.Println(userIdList)

	var posts []*model.Post
	if err := db.Limit(limit).Offset(offset).Find(&posts, "sender_id IN ?", userIdList).Error; err != nil {
		return nil, err
	}

	fmt.Println("----------------")
	fmt.Println(posts)
	fmt.Println("----------------")
	fmt.Println(userID)

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
