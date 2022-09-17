package service

import (
	"bytes"
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/MneachDev/LinkhedIn-backend/authentication"
	"github.com/MneachDev/LinkhedIn-backend/graph/model"
	"github.com/samber/lo"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

func SendEmailActivation(ToMail string, linkActivation string) {

	link := linkActivation

	msg := gomail.NewMessage()
	msg.SetHeader("From", "MneachDev@gmail.com")
	msg.SetHeader("To", ToMail)
	msg.SetHeader("Subject", "Confirm your account on LinkhedIn")
	msg.SetBody("text/html", `Activate your account via this link : <br> http://localhost:5173/activationAccount/`+link)

	n := gomail.NewDialer("smtp.gmail.com", 587, "MneachDev@gmail.com", "qpdwvwkkyxovxtue")

	if err := n.DialAndSend(msg); err != nil {
		panic(err)
	}
}

func SendEmailResetPassword(ToMail string, linkActivation string) {

	link := linkActivation

	log.Print(linkActivation)

	msg := gomail.NewMessage()
	msg.SetHeader("From", "MneachDev@gmail.com")
	msg.SetHeader("To", ToMail)
	msg.SetHeader("Subject", "Reset Password LinkhedIn")
	msg.SetBody("text/html", `You recently took steps to reset the password for your LinkhedIn Account. 
	Click on the link below to reset your password.
	: <br> http://localhost:5173/resetPassword/`+link)

	n := gomail.NewDialer("smtp.gmail.com", 587, "MneachDev@gmail.com", "qpdwvwkkyxovxtue")

	if err := n.DialAndSend(msg); err != nil {
		panic(err)
	}
}

func GenerateRandomLinkActivation() string {

	var link bytes.Buffer
	var charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 35; i++ {
		randomIndex := rand.Intn(len(charset))
		link.WriteByte(charset[randomIndex])
	}

	return link.String()
}

func CreateActiveLink(db *gorm.DB, ctx context.Context, userID string, linkActivation string) (*model.ActivateAccount, error) {
	modelActivationLink := &model.ActivateAccount{
		ID:     linkActivation,
		UserID: userID,
	}

	return modelActivationLink, db.Create(modelActivationLink).Error
}

func Search(db *gorm.DB, ctx context.Context, keyword string, limit int, offset int) (*model.Search, error) {
	search := new(model.Search)

	userID := authentication.GetJwtValueData(ctx).Userid
	var modelUsers []*model.User
	var modelPosts []*model.Post

	// SEARCH USER BY KEYWORD
	if err := db.Limit(limit).Offset(offset).Not("id = ?", userID).Find(&modelUsers, "concat(first_name, last_name) like ?", "%"+keyword+"%").Error; err != nil {
		return nil, err
	}

	// SEARCH POSTS BY KEYWOARD

	if err := db.Limit(limit).Offset(offset).Find(&modelPosts, "text like ? ", "%"+keyword+"%").Error; err != nil {
		return nil, err
	}

	search.Users = modelUsers
	search.Posts = modelPosts

	return search, nil
}

func SearchHastag(db *gorm.DB, ctx context.Context, keyword string, limit int, offset int) (*model.Search, error) {
	search := new(model.Search)

	var modelPosts []*model.Post

	// SEARCH POSTS BY KEYWOARD

	if err := db.Limit(limit).Offset(offset).Find(&modelPosts, "text like ? ", "%#"+keyword+"%").Error; err != nil {
		return nil, err
	}

	search.Posts = modelPosts

	return search, nil
}

func GetUserSearch(db *gorm.DB, ctx context.Context, obj *model.Search) ([]*model.User, error) {
	var users []*model.User

	userIds := lo.Map(obj.Users, func(user *model.User, _ int) string {
		return user.ID
	})
	if len(userIds) == 0 {
		return users, nil
	}

	if err := db.Find(&users, userIds).Error; err != nil {
		return nil, err
	}

	// log.Print(users)
	return users, nil
}

func GetPostSearch(db *gorm.DB, ctx context.Context, obj *model.Search) ([]*model.Post, error) {
	var posts []*model.Post

	postIds := lo.Map(obj.Posts, func(post *model.Post, _ int) string {
		return post.ID
	})

	if len(postIds) == 0 {
		return posts, nil
	}

	if err := db.Find(&posts, postIds).Error; err != nil {
		return nil, err
	}

	return posts, nil
}
