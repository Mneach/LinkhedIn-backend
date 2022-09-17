package service

import (
	"context"

	"github.com/MneachDev/LinkhedIn-backend/graph/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetExperienceUser(db *gorm.DB, ctx context.Context, obj *model.User) ([]*model.Experience, error) {
	var modelExperiences []*model.Experience

	return modelExperiences, db.Where("user_id = ?", obj.ID).Find(&modelExperiences).Error
}

func AddExperience(db *gorm.DB, ctx context.Context, input model.InputExperience) (*model.Experience, error) {
	modelExperience := &model.Experience{
		ID:             uuid.NewString(),
		Title:          input.Title,
		UserID:         input.UserID,
		EmploymentType: input.EmploymentType,
		CompanyName:    input.CompanyName,
		Country:        input.Country,
		City:           input.City,
		IsActive:       input.IsActive,
		Industry:       input.Industry,
		MonthStartDate: input.MonthStartDate,
		MonthEndDate:   input.MonthEndDate,
		YearStartDate:  input.YearStartDate,
		YearEndDate:    input.YearEndDate,
	}

	return modelExperience, db.Save(modelExperience).Error
}

// UpdateExperience is the resolver for the updateExperience field.
func UpdateExperience(db *gorm.DB, ctx context.Context, id string, input model.InputExperience) (*model.Experience, error) {
	modelExperience := new(model.Experience)

	if err := db.First(modelExperience, "id = ?", id).Error; err != nil {
		return nil, err
	}

	modelExperience.EmploymentType = input.EmploymentType
	modelExperience.CompanyName = input.CompanyName
	modelExperience.Country = input.Country
	modelExperience.City = input.City
	modelExperience.IsActive = input.IsActive
	modelExperience.Industry = input.Industry
	modelExperience.MonthStartDate = input.MonthStartDate
	modelExperience.MonthEndDate = input.MonthEndDate
	modelExperience.YearStartDate = input.YearStartDate
	modelExperience.YearEndDate = input.YearEndDate

	return modelExperience, db.Save(modelExperience).Error
}

// DeleteExperience is the resolver for the deleteExperience field.
func DeleteExperience(db *gorm.DB, ctx context.Context, id string) (*model.Experience, error) {
	modelExperience := new(model.Experience)

	if err := db.First(modelExperience, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return modelExperience, db.Delete(modelExperience).Error
}

// Experiences is the resolver for the Experiences field.
func Experiences(db *gorm.DB, ctx context.Context) ([]*model.Experience, error) {
	var modelExperiences []*model.Experience
	return modelExperiences, db.Find(&modelExperiences).Error
}
