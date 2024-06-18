package controllers

import (
	"log"
	"net/http"
	"se-api/src/internal/lib/common"
	"se-api/src/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (ac *AuthController) Signup(c *gin.Context) {
	uid := common.GetUserID(c)
	userService := services.NewUserService()

	_, err := userService.GetUserByID(uid)
	if err != nil && err != gorm.ErrRecordNotFound {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Fatalf("Error getting user by id: %v", err)
	}
	if err == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "duplicate uid",
		})
		return
	}

	_, err = userService.CreateUserFromID(uid)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Fatalf("Error creating user: %v", err)
	}

	c.Status(http.StatusOK)
}
