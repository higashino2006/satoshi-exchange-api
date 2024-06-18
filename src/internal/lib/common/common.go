package common

import (
	"fmt"
	"math"

	"github.com/gin-gonic/gin"
)

type KeyValue map[string]interface{}

func GetUserID(c *gin.Context) string {
	userId, _ := c.Get("user_id")
	return fmt.Sprint(userId)
}

func RoundDownFrom4DecimalPlaces(value float64) float64 {
	truncated := math.Trunc(value * 1000)
	return truncated / 1000
}

func ConvertSatoshiToJPY(Satoshi float32) float32 {
	return Satoshi * float32(0.3)
}
