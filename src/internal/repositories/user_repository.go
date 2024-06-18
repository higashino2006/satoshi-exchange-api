package repositories

import (
	"log"
	"se-api/src/internal/db"
	"se-api/src/internal/lib/common"
	"se-api/src/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{db: db.GetDB()}
}

func NewUserRepositoryWithTx(tx *gorm.DB) *UserRepository {
	return &UserRepository{db: tx}
}

func (r *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	err := r.db.Create(user).Error
	if err != nil {
		log.Printf("Error creating user on db: %v", err)
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByID(userID string) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, "id = ?", userID).Error
	if err != nil {
		log.Printf("Error getting user on db: %v", err)
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUserByID(userID string, m *common.KeyValue) error {
	err := r.db.Model(&models.User{}).Where("id = ?", userID).Updates(map[string]interface{}(*m)).Error
	if err != nil {
		log.Printf("Error updating user on db: %v", err)
		return err
	}
	return nil
}
