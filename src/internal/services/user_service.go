package services

import (
	"errors"
	"se-api/src/internal/lib/common"
	"se-api/src/internal/models"
	"se-api/src/internal/repositories"

	"gorm.io/gorm"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService() *UserService {
	return &UserService{repo: repositories.NewUserRepository()}
}

func NewUserServiceWithTx(tx *gorm.DB) *UserService {
	return &UserService{repo: repositories.NewUserRepositoryWithTx(tx)}
}

func (s *UserService) GetUserByID(userID string) (*models.User, error) {
	return s.repo.GetUserByID(userID)
}

func (s *UserService) CreateUserFromID(userID string) (*models.User, error) {
	return s.repo.CreateUser(&models.User{ID: userID})
}

func (s *UserService) UpdateUser(userID string, m *common.KeyValue) error {
	if len(*m) == 0 {
		return errors.New("at least one key is required")
	}
	return s.repo.UpdateUserByID(userID, m)
}

func (s *UserService) UpdateUserBalance(userID string, satoshi, jpy float32) error {
	return s.UpdateUser(userID, &common.KeyValue{
		"jpy_balance":     jpy,
		"satoshi_balance": satoshi,
	})
}
