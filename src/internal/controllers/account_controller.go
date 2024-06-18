package controllers

import (
	"log"
	"net/http"
	"se-api/src/internal/lib/common"
	"se-api/src/internal/services"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
}

func NewAccountController() *AccountController {
	return &AccountController{}
}

func (ac *AccountController) Balance(c *gin.Context) {
	uid := common.GetUserID(c)
	userService := services.NewUserService()

	user, err := userService.GetUserByID(uid)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Fatalf("Error getting user by id: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"jpy_balance":     user.JPYBalance,
		"satoshi_balance": user.SatoshiBalance,
	})
}
