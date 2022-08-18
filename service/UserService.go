package service

import (
	"context"
	"log"

	"github.com/MneachDev/LinkhedIn-backend/authentication"
	"github.com/MneachDev/LinkhedIn-backend/graph/model"
	"github.com/google/uuid"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func RegisterUser(ctx context.Context, input model.InputRegisterUser) (*model.User, error) {
	modelUser := &model.User{
		ID:                 uuid.NewString(),
		Email:              input.Email,
		Password:           input.Password,
		IsActive:           false,
		FirstName:          "",
		LastName:           "",
		AdditionalName:     "",
		ProfileImageURL:    "",
		BackgroundImageURL: "",
		Pronouns:           "",
		Headline:           "",
		About:              "",
		Country:            "",
		City:               "",
		ProfileLink:        "",
	}

	SendEmail(input.Email)

	return modelUser, useDB().Create(modelUser).Error
}

func Login(ctx context.Context, input model.InputLogin) (interface{}, error) {
	modelUser := new(model.User)

	if err := useDB().Where("email = ? AND password = ?", input.Email, input.Password).Find(modelUser).Error; err != nil {
		return nil, err
	}

	if !modelUser.IsActive {
		return nil, gqlerror.Errorf("your account is still not active")
	}

	token, err := authentication.JwtGenerate(ctx, modelUser.ID)
	if err != nil {
		return nil, err
	}

	log.Print(token)
	return map[string]interface{}{
		"token": token,
	}, nil
}

func UpdateUser(ctx context.Context, id string, input model.InputUpdateUser) (*model.User, error) {
	modelUser := new(model.User)

	if err := useDB().First(modelUser, "id = ?", id).Error; err != nil {
		return nil, err
	}

	modelUser.Email = input.Email
	modelUser.Password = input.Password
	modelUser.IsActive = input.IsActive
	modelUser.FirstName = input.FirstName
	modelUser.LastName = input.LastName
	modelUser.AdditionalName = input.AdditionalName
	modelUser.ProfileImageURL = input.ProfileImageURL
	modelUser.BackgroundImageURL = input.BackgroundImageURL
	modelUser.Pronouns = input.Pronouns
	modelUser.Headline = input.Headline
	modelUser.About = input.About
	modelUser.Country = input.Country
	modelUser.City = input.City
	modelUser.ProfileLink = input.ProfileLink

	return modelUser, useDB().Save(modelUser).Error
}

func DeleteUser(ctx context.Context, id string) (*model.User, error) {
	modelUser := new(model.User)

	if err := useDB().First(modelUser, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return modelUser, useDB().Delete(modelUser).Error
}

func GetUser(ctx context.Context, id string) (*model.User, error) {
	modelUser := new(model.User)
	return modelUser, useDB().First(modelUser, "id = ?", id).Error
}

func GetUsers(ctx context.Context) ([]*model.User, error) {
	var modelUsers []*model.User
	return modelUsers, useDB().Find(&modelUsers).Error
}

func ActivateUser(ctx context.Context, id string) (*model.User, error) {
	modelUser := new(model.User)

	if err := useDB().First(modelUser, "id = ?", id).Error; err != nil {
		return nil, err
	}

	modelUser.IsActive = true

	return modelUser, useDB().Save(modelUser).Error
}
