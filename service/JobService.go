package service

import (
	"context"
	"time"

	"github.com/MneachDev/LinkhedIn-backend/graph/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AddJob is the resolver for the addJob field.
func AddJob(db *gorm.DB, ctx context.Context, title string, companyName string, workplace string, city string, country string, employmentType string, description string) (*model.Job, error) {
	modelJobs := &model.Job{
		ID:             uuid.NewString(),
		Title:          title,
		CompanyName:    companyName,
		Workplace:      workplace,
		City:           city,
		Country:        country,
		EmploymentType: employmentType,
		Description:    description,
		CreatedAt:      time.Now(),
	}

	return modelJobs, db.Create(modelJobs).Error
}

// Jobs is the resolver for the Jobs field.
func GetJobs(db *gorm.DB, ctx context.Context) ([]*model.Job, error) {
	var modelJobs []*model.Job

	if err := db.Order("created_at desc").Find(&modelJobs).Error; err != nil {
		return nil, err
	}

	return modelJobs, nil
}
