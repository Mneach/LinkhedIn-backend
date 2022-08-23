package service

import (
	"context"

	"github.com/MneachDev/LinkhedIn-backend/graph/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetEducations(db *gorm.DB, ctx context.Context, obj *model.User) ([]*model.Education, error) {
	var modelEducations []*model.Education

	return modelEducations, db.Where("user_id = ? ", obj.ID).Find(&modelEducations).Error
}

// AddEducation is the resolver for the addEducation field.
func AddEducation(db *gorm.DB, ctx context.Context, input model.InputEducation) (*model.Education, error) {
	modelEducation := &model.Education{
		ID:             uuid.NewString(),
		UserID:         input.UserID,
		School:         input.School,
		Degree:         input.Degree,
		FieldStudy:     input.FieldStudy,
		Grade:          input.Grade,
		Activities:     input.Activities,
		Description:    input.Description,
		MonthStartDate: input.MonthStartDate,
		MonthEndDate:   input.MonthEndDate,
		YearStartDate:  input.YearStartDate,
		YearEndDate:    input.YearEndDate,
	}

	return modelEducation, db.Save(modelEducation).Error
}

// UpdateEducation is the resolver for the updateEducation field.
func UpdateEducation(db *gorm.DB, ctx context.Context, id string, input model.InputEducation) (*model.Education, error) {
	modelEducation := new(model.Education)

	if err := db.First(modelEducation, "id = ?", id).Error; err != nil {
		return nil, err
	}

	modelEducation.School = input.School
	modelEducation.Degree = input.Degree
	modelEducation.FieldStudy = input.FieldStudy
	modelEducation.Grade = input.Grade
	modelEducation.Activities = input.Activities
	modelEducation.Description = input.Description
	modelEducation.MonthStartDate = input.MonthStartDate
	modelEducation.MonthEndDate = input.MonthEndDate
	modelEducation.YearStartDate = input.YearStartDate
	modelEducation.YearEndDate = input.YearEndDate

	return modelEducation, db.Save(modelEducation).Error
}

// DeleteEducation is the resolver for the deleteEducation field.
func DeleteEducation(db *gorm.DB, ctx context.Context, id string) (*model.Education, error) {
	modelEducation := new(model.Education)

	if err := db.First(modelEducation, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return modelEducation, db.Delete(modelEducation).Error
}

// Educations is the resolver for the Educations field.
func Educations(db *gorm.DB, ctx context.Context) ([]*model.Education, error) {
	var modelEducations []*model.Education
	return modelEducations, db.Find(&modelEducations).Error
}
