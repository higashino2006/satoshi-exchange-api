package repositories

import (
	"log"
	"se-api/src/internal/db"
	"se-api/src/internal/models"

	"gorm.io/gorm"
)

type TradeRecordRepository struct {
	db *gorm.DB
}

func NewTradeRecordRepository() *TradeRecordRepository {
	return &TradeRecordRepository{db: db.GetDB()}
}

func NewTradeRecordRepositoryWithTx(tx *gorm.DB) *TradeRecordRepository {
	return &TradeRecordRepository{db: tx}
}

func (r *TradeRecordRepository) CreateTradeRecord(record *models.TradeRecord) (*models.TradeRecord, error) {
	err := r.db.Create(record).Error
	if err != nil {
		log.Printf("Error creating trade record on db: %v", err)
		return nil, err
	}
	return record, nil
}
