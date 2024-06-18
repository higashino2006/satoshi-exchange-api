package services

import (
	"log"
	"se-api/src/internal/db"
	"se-api/src/internal/models"

	"gorm.io/gorm"
)

type TxService struct {
	db *gorm.DB
}

func NewTxService() *TxService {
	return &TxService{db: db.GetDB()}
}

func (ts *TxService) CreateTradeRecordAndUpdateBalance(
	tradeRecord *models.TradeRecord,
	uid string,
	jpyBalance float32,
	satoshiBalance float32,
) error {
	tx := ts.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	userServiceWithTx := NewUserServiceWithTx(tx)
	tradeRecordServiceWithTx := NewTradeRecordServiceWithTx(tx)

	_, err := tradeRecordServiceWithTx.CreateTradeRecord(tradeRecord)
	if err != nil {
		log.Printf("Error creating trade record: %v", err)
		tx.Rollback()
		return err
	}

	err = userServiceWithTx.UpdateUserBalance(uid, satoshiBalance, jpyBalance)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		log.Printf("Error commiting transaction: %v", err)
		tx.Rollback()
		return err
	}

	return nil
}
