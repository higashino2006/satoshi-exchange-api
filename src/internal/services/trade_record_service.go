package services

import (
	"se-api/src/internal/models"
	"se-api/src/internal/repositories"

	"gorm.io/gorm"
)

type TradeRecordService struct {
	repo *repositories.TradeRecordRepository
}

func NewTradeRecordService() *TradeRecordService {
	return &TradeRecordService{repo: repositories.NewTradeRecordRepository()}
}

func NewTradeRecordServiceWithTx(tx *gorm.DB) *TradeRecordService {
	return &TradeRecordService{repo: repositories.NewTradeRecordRepositoryWithTx(tx)}
}

func (s *TradeRecordService) CreateTradeRecord(tradeRecord *models.TradeRecord) (*models.TradeRecord, error) {
	return s.repo.CreateTradeRecord(tradeRecord)
}
