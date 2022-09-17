package service

import (
	"context"

	"github.com/MneachDev/LinkhedIn-backend/graph/model"
	"github.com/google/uuid"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/gorm"
)

func AddHastag(db *gorm.DB, ctx context.Context, hastag string) (*model.Hastag, error) {
	modelHastag := new(model.Hastag)
	modelInputHastag := &model.Hastag{
		ID:     uuid.NewString(),
		Hastag: hastag,
	}

	if err := db.Find(modelHastag, "hastag = ?", hastag).Error; err != nil {
		return nil, err
	}

	if modelHastag.ID != "" {
		return nil, gqlerror.Errorf("Hastag already exists")
	}

	return modelInputHastag, db.Create(modelInputHastag).Error
}

// Hastags is the resolver for the Hastags field.
func GetHastags(db *gorm.DB, ctx context.Context) ([]*model.Hastag, error) {
	var modelHastags []*model.Hastag

	return modelHastags, db.Find(&modelHastags).Error
}
