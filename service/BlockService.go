package service

import (
	"context"

	"github.com/MneachDev/LinkhedIn-backend/graph/model"
	"gorm.io/gorm"
)

func AddBlock(db *gorm.DB, ctx context.Context, userID string, blockID string) (*model.Block, error) {
	modelBlock := &model.Block{
		UserID:  userID,
		BlockID: blockID,
	}

	return modelBlock, db.Table("user_blocks").Create(modelBlock).Error
}

// DeleteBlock is the resolver for the deleteBlock field.
func DeleteBlock(db *gorm.DB, ctx context.Context, userID string, blockID string) (*model.Block, error) {
	modelBlock := new(model.Block)

	return modelBlock, db.Table("user_blocks").Delete(modelBlock, "user_id = ? AND block_id = ?", userID, blockID).Error
}
