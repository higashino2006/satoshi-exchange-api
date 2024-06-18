package controllers

import (
	"log"
	"net/http"
	"se-api/src/internal/lib/common"
	"se-api/src/internal/models"
	"se-api/src/internal/services"

	"github.com/gin-gonic/gin"
)

type TradeController struct {
}

func NewTradeController() *TradeController {
	return &TradeController{}
}

func (tc *TradeController) BuyCrypto(c *gin.Context) {
	var jsonBody struct {
		Satoshi float32 `json:"satoshi" binding:"required,gt=0"`
	}

	if err := c.ShouldBindJSON(&jsonBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	satoshiToBuy := common.RoundDownFrom4DecimalPlaces(jsonBody.Satoshi)
	uid := common.GetUserID(c)
	userService := services.NewUserService()

	// check if the user has enough jpy
	requiredJPY := common.ConvertSatoshiToJPY(satoshiToBuy)
	user, err := userService.GetUserByID(uid)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Fatalf("Error getting user by id: %v", err)
	}
	if user.JPYBalance < requiredJPY {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "not enough jpy",
		})
		return
	}

	txService := services.NewTxService()

	// create trade record and update balance
	tradeRecord := &models.TradeRecord{
		UserID:  user.ID,
		Type:    "buy",
		JPY:     requiredJPY,
		Satoshi: satoshiToBuy,
	}
	newJPYBalance := user.JPYBalance - requiredJPY
	newSatoshiBalance := user.SatoshiBalance + satoshiToBuy
	err = txService.CreateTradeRecordAndUpdateBalance(
		tradeRecord,
		uid,
		newJPYBalance,
		newSatoshiBalance,
	)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Fatalf("Error transaction: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"jpy_balance":     newJPYBalance,
		"satoshi_balance": newSatoshiBalance,
	})
}

func (tc *TradeController) SellCrypto(c *gin.Context) {
	var jsonBody struct {
		Satoshi float32 `json:"satoshi" binding:"required,gt=0"`
	}

	if err := c.ShouldBindJSON(&jsonBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	satoshiToSell := common.RoundDownFrom4DecimalPlaces(jsonBody.Satoshi)
	uid := common.GetUserID(c)
	userService := services.NewUserService()

	// check if the user has enough satoshi
	user, err := userService.GetUserByID(uid)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Fatalf("Error getting user by id: %v", err)
	}
	if user.SatoshiBalance < satoshiToSell {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "not enough satoshi",
		})
		return
	}

	txService := services.NewTxService()

	// create trade record and update balance
	JPYObtained := common.ConvertSatoshiToJPY(satoshiToSell)
	tradeRecord := &models.TradeRecord{
		UserID:  user.ID,
		Type:    "sell",
		JPY:     JPYObtained,
		Satoshi: satoshiToSell,
	}
	newJPYBalance := user.JPYBalance + JPYObtained
	newSatoshiBalance := user.SatoshiBalance - satoshiToSell
	err = txService.CreateTradeRecordAndUpdateBalance(
		tradeRecord,
		uid,
		newJPYBalance,
		newSatoshiBalance,
	)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Fatalf("Error transaction: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"jpy_balance":     newJPYBalance,
		"satoshi_balance": newSatoshiBalance,
	})
}
