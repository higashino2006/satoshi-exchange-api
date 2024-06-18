package models

import (
	"time"
)

type TradeRecord struct {
	ID        uint
	UserID    string
	Type      string
	JPY       float32
	Satoshi   float32
	CreatedAt time.Time
}
