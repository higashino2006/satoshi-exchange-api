package common

import (
	"fmt"
	"math"
	"strings"

	"github.com/gin-gonic/gin"
)

type KeyValue map[string]interface{}

func GetUserID(c *gin.Context) string {
	userId, _ := c.Get("user_id")
	return fmt.Sprint(userId)
}

func RoundDownFrom4DecimalPlaces(value float32) float32 {
	truncated := math.Trunc(float64(value * 1000))
	return float32(truncated / 1000)
}

func JoinPaths(baseURL, subPath string) string {
	baseURL = strings.TrimSuffix(baseURL, "/")
	subPath = strings.TrimPrefix(subPath, "/")
	return fmt.Sprintf("%s/%s", baseURL, subPath)
}

func ConvertSatoshiToJPY(Satoshi float32) float32 {
	return Satoshi * float32(0.3)
}
