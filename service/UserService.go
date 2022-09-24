package service

import (
	"context"
	"fmt"
	"log"

	"github.com/MneachDev/LinkhedIn-backend/authentication"
	"github.com/MneachDev/LinkhedIn-backend/graph/model"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/gorm"
)

func RegisterUser(db *gorm.DB, ctx context.Context, input model.InputRegisterUser) (*model.User, error) {

	modelUsers := new(model.User)

	if err := db.Find(&modelUsers, "email = ?", input.Email).Error; err != nil {
		return nil, err
	}

	if modelUsers.ID != "" && !modelUsers.IsActive {
		return nil, gqlerror.Errorf("Email Already Registered And The Account Still Not Active")
	}

	if modelUsers.ID != "" {
		return nil, gqlerror.Errorf("Email Already Registered")
	}

	modelUser := &model.User{
		ID:                 uuid.NewString(),
		Email:              input.Email,
		Password:           input.Password,
		IsActive:           false,
		FirstName:          input.FirstName,
		LastName:           input.LastName,
		ProfileImageURL:    input.ProfileImageURL,
		BackgroundImageURL: "",
		Pronouns:           "",
		Headline:           input.Headline,
		About:              "",
		Country:            input.Country,
		City:               input.City,
		ProfileLink:        "",
	}

	linkActivation := GenerateRandomLinkActivation()
	SendEmailActivation(input.Email, linkActivation)

	if err := db.Create(modelUser).Error; err != nil {
		return nil, err
	}

	CreateActiveLink(db, ctx, modelUser.ID, linkActivation)

	return modelUser, nil
}

func Login(db *gorm.DB, ctx context.Context, input model.InputLogin) (interface{}, error) {
	modelUser := new(model.User)

	if err := db.Where("email = ? AND password = ?", input.Email, input.Password).Find(modelUser).Error; err != nil {
		return nil, err
	}

	if modelUser.ID == "" {
		return nil, gqlerror.Errorf("Wrong Credential!")
	}

	if !modelUser.IsActive {
		return nil, gqlerror.Errorf("Your Account Is Still Not Active")
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

func UpdateUser(db *gorm.DB, ctx context.Context, id string, input model.InputUpdateUser) (*model.User, error) {
	modelUser := new(model.User)

	if err := db.First(modelUser, "id = ?", id).Error; err != nil {
		return nil, err
	}

	log.Print(input.Email)
	log.Print(input.BackgroundImageURL)

	modelUser.Email = input.Email
	modelUser.Password = input.Password
	modelUser.IsActive = input.IsActive
	modelUser.FirstName = input.FirstName
	modelUser.LastName = input.LastName
	modelUser.ProfileImageURL = input.ProfileImageURL
	modelUser.BackgroundImageURL = input.BackgroundImageURL
	modelUser.Pronouns = input.Pronouns
	modelUser.Headline = input.Headline
	modelUser.About = input.About
	modelUser.Country = input.Country
	modelUser.City = input.City
	modelUser.ProfileLink = input.ProfileLink

	log.Print(modelUser)

	return modelUser, db.Save(modelUser).Error
}

func DeleteUser(db *gorm.DB, ctx context.Context, id string) (*model.User, error) {
	modelUser := new(model.User)

	if err := db.First(modelUser, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return modelUser, db.Delete(modelUser).Error
}

func GetUser(db *gorm.DB, ctx context.Context, id string) (*model.User, error) {
	modelUser := new(model.User)
	return modelUser, db.Find(&modelUser, "id = ?", id).Error
}

func GetUsers(db *gorm.DB, ctx context.Context) ([]*model.User, error) {
	var modelUsers []*model.User
	return modelUsers, db.Find(&modelUsers).Error
}

func GetUserByActivationID(db *gorm.DB, ctx context.Context, activationID string) (*model.User, error) {
	modelUser := new(model.User)
	modelActivationAccount := new(model.ActivateAccount)

	if err := db.First(modelActivationAccount, "id = ?", activationID).Error; err != nil {
		return nil, err
	}

	if err := db.First(modelUser, "id = ?", modelActivationAccount.UserID).Error; err != nil {
		return nil, err
	}

	if modelUser.ID == "" {
		return nil, gqlerror.Errorf("User Not Found")
	}

	ActivateUser(modelUser)

	return modelUser, db.Save(modelUser).Error
}

func ActivateUser(modelUser *model.User) *model.User {

	modelUser.IsActive = true

	return modelUser
}

func RegisterResetPassword(db *gorm.DB, ctx context.Context, email string) (*model.ResetPasswordAccount, error) {
	modelUser := new(model.User)
	linkResetPassword := GenerateRandomLinkActivation()

	if err := db.First(modelUser, "email = ?", email).Error; err != nil {
		return nil, gqlerror.Errorf("Email address is not linked to any account")
	}

	modelResetPassword := &model.ResetPasswordAccount{
		ID:     linkResetPassword,
		UserID: modelUser.ID,
	}

	SendEmailResetPassword(email, linkResetPassword)

	return modelResetPassword, db.Create(modelResetPassword).Error
}

func GetUserByResetPasswordID(db *gorm.DB, ctx context.Context, resetPasswordID string) (*model.User, error) {
	modelUser := new(model.User)
	modelResetPasswordAccount := new(model.ResetPasswordAccount)

	db.First(modelResetPasswordAccount, "id = ?", resetPasswordID)
	db.First(modelUser, "id = ?", modelResetPasswordAccount.UserID)

	if modelResetPasswordAccount.UserID == "" {
		return nil, gqlerror.Errorf("Link Invalid")
	}

	if modelUser.ID == "" {
		return nil, gqlerror.Errorf("User Not Found")
	}

	return modelUser, nil
}

func GetUserByEmail(db *gorm.DB, ctx context.Context, email string) (*model.User, error) {
	modelUser := new(model.User)

	if err := db.First(modelUser, "email = ?", email).Error; err != nil {
		return nil, err
	}

	if modelUser.ID != "" && !modelUser.IsActive {
		return nil, gqlerror.Errorf("Email Already Registered And The Account Still Not Active")
	}

	return modelUser, nil
}

func CheckEmailUser(db *gorm.DB, ctx context.Context, email string) (*model.User, error) {
	modelUser := new(model.User)

	db.First(modelUser, "email = ?", email)

	if modelUser.ID != "" && !modelUser.IsActive {
		return nil, gqlerror.Errorf("Email Already Registered And The Account Still Not Active")
	}

	if modelUser.ID != "" {
		return nil, gqlerror.Errorf("Email Already Registered")
	}

	return modelUser, nil

}

func UpdatePasswordUser(db *gorm.DB, ctx context.Context, id string, password string) (*model.User, error) {
	modelUser := new(model.User)
	modelResetPasswordAccount := new(model.ResetPasswordAccount)

	if err := db.First(modelUser, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if modelUser.ID == "" {
		return nil, gqlerror.Errorf("User Not Found")
	}

	modelUser.Password = password

	if err := db.First(modelResetPasswordAccount, "user_id = ?", id).Error; err != nil {
		return nil, gqlerror.Errorf("Reset Password Data Not Found!")
	}

	if err := db.Delete(modelResetPasswordAccount).Error; err != nil {
		return nil, err
	}

	return modelUser, db.Save(modelUser).Error
}

func GetVisits(db *gorm.DB, ctx context.Context, obj *model.User) ([]*model.Visit, error) {
	var modelVisits []*model.Visit

	return modelVisits, db.Table("user_visits").Find(&modelVisits, "visit_id = ?", obj.ID).Error
}

func GetFollows(db *gorm.DB, ctx context.Context, obj *model.User) ([]*model.Follow, error) {
	var modelFollow []*model.Follow

	return modelFollow, db.Table("user_follows").Find(&modelFollow, "follow_id = ? ", obj.ID).Error
}

func VisitUser(db *gorm.DB, ctx context.Context, id1 string, id2 string) (interface{}, error) {
	modelVisit := new(model.Visit)

	db.Table("user_visits").First(&modelVisit, "user_id = ? AND visit_id = ?", id1, id2)

	if modelVisit.UserID != "" {
		var modelVisits []*model.Visit
		db.Table("user_visits").Find(&modelVisits, "visit_id = ?", id2)

		return map[string]interface{}{
			"length": len(modelVisits),
		}, nil
	} else {

		modelVisit.UserID = id1
		modelVisit.VisitID = id2

		if err := db.Table("user_visits").Create(modelVisit).Error; err == nil {
			AddNotification(db, ctx, id2, id1, "Visit Your Profile")
		}

		var modelVisits []*model.Visit
		db.Table("user_visits").Find(&modelVisits, "visit_id = ?", id2)

		return map[string]interface{}{
			"length": len(modelVisits),
		}, nil
	}
}

func FollowUser(db *gorm.DB, ctx context.Context, id1 string, id2 string) (interface{}, error) {
	modelFollow := new(model.Follow)

	modelFollow.UserID = id1
	modelFollow.FollowID = id2

	db.Table("user_follows").Create(modelFollow)

	var modelFollows []*model.Follow
	db.Table("user_follows").Find(&modelFollows, "follow_id = ?", id2)

	return map[string]interface{}{
		"length": len(modelFollows),
	}, nil

}

func UnFollowUser(db *gorm.DB, ctx context.Context, id1 string, id2 string) (interface{}, error) {
	modelFollow := new(model.Follow)

	if err := db.Table("user_follows").First(&modelFollow, "user_id = ? AND follow_id = ?", id1, id2).Error; err != nil {
		return nil, err
	}

	if modelFollow.UserID == "" {
		var modelFollows []*model.Follow
		db.Table("user_follows").Find(&modelFollows, "follow_id = ?", id2)

		return map[string]interface{}{
			"length": len(modelFollows),
		}, nil
	} else {
		db.Table("user_follows").Delete(&modelFollow, "user_id = ? AND follow_id = ?", id1, id2)

		var modelFollows []*model.Follow
		db.Table("user_follows").Find(&modelFollows, "follow_id = ?", id2)

		return map[string]interface{}{
			"length": len(modelFollows),
		}, nil
	}
}

func GetConnections(db *gorm.DB, ctx context.Context, obj *model.User) ([]*model.Connection, error) {
	var modelConnections []*model.Connection

	if err := db.Where("user1_id = ?", obj.ID).Or("user2_id = ?", obj.ID).Find(&modelConnections).Error; err != nil {
		return nil, err
	}

	return modelConnections, nil
}

// ConnectRequests is the resolver for the ConnectRequests field.
func GetConnectRequests(db *gorm.DB, ctx context.Context, obj *model.User) ([]*model.ConnectRequest, error) {
	var modelConnectionRequests []*model.ConnectRequest

	if err := db.Find(&modelConnectionRequests, "to_user_id = ?", obj.ID).Error; err != nil {
		return nil, err
	}

	return modelConnectionRequests, nil
}

func GetBlockUsers(db *gorm.DB, ctx context.Context, obj *model.User) ([]*model.Block, error) {
	var modelBlocks []*model.Block

	if err := db.Table("user_blocks").Where("user_id = ?", obj.ID).Or("block_id = ?", obj.ID).Find(&modelBlocks).Error; err != nil {
		return nil, err
	}

	return modelBlocks, nil
}

func GetUserSuggestion(db *gorm.DB, ctx context.Context, userID string) ([]*model.User, error) {
	var modelUsers []*model.User
	var userIdList []string
	var userSuggestionId []string

	// CONNECTED ALL USER CONNECTION
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

	// GET ALL FRIEND CONNECTION , Exclude Current user
	var friendConnection1 []*model.Connection
	var friendConnection2 []*model.Connection

	if err := db.Where("user1_id IN ?", userIdList).Not("user2_id = ?", userID).Find(&friendConnection1).Error; err != nil {
		return nil, err
	}

	if err := db.Where("user2_id IN ?", userIdList).Not("user1_id = ?", userID).Find(&friendConnection2).Error; err != nil {
		return nil, err
	}

	fmt.Println(userIdList)

	userSuggestion1Ids := lo.Map(friendConnection1, func(connectionData *model.Connection, _ int) string {
		return connectionData.User2ID
	})

	userSuggestion2Ids := lo.Map(friendConnection2, func(connectionData *model.Connection, _ int) string {
		return connectionData.User1ID
	})

	userSuggestionId = append(userSuggestionId, userSuggestion1Ids...)
	userSuggestionId = append(userSuggestionId, userSuggestion2Ids...)
	userSuggestionId = lo.Uniq(userSuggestionId)
	fmt.Println(userSuggestionId)

	//REMOVE USER THAT CURRENT USER HAS CONNECTED
	var finalUserSuggestionId []string
	for _, suggestionIdUser := range userSuggestionId {
		checkSame := false
		for _, userConnectionId := range userIdList {
			if suggestionIdUser == userConnectionId {
				checkSame = true
			}
		}

		if !checkSame {
			finalUserSuggestionId = append(finalUserSuggestionId, suggestionIdUser)
		}
	}

	fmt.Println(finalUserSuggestionId)

	if len(finalUserSuggestionId) == 0 {
		return nil, gqlerror.Errorf("No Connection User Data")
	}

	if err := db.Find(&modelUsers, finalUserSuggestionId).Error; err != nil {
		return nil, err
	}

	return modelUsers, nil
}

func UserConnected(db *gorm.DB, ctx context.Context, userID string) ([]*model.User, error) {
	var modelUsers []*model.User

	// CONNECTED USER POST
	var userIdList []string
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

	if err := db.Find(&modelUsers, userIdList).Error; err != nil {
		return nil, err
	}

	return modelUsers, nil
}
