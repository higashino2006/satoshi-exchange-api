package models

import (
	"time"
)

type User struct {
	ID             string
	JPYBalance     float32
	SatoshiBalance float32
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
