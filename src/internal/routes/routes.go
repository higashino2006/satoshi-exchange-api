package routes

import (
	"se-api/src/internal/config"
	"se-api/src/internal/controllers"
	"se-api/src/internal/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	configureCORS(r)
	registerRoutes(r)
	r.Run()
}

func configureCORS(r *gin.Engine) {
	appUrl := config.AppConfig.FRONTEND_URL
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{appUrl},
		AllowMethods: []string{
			"GET",
			"POST",
			"PATCH",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{"*"},
		MaxAge:       24 * time.Hour,
	}))
}

func registerRoutes(r *gin.Engine) {
	accountController := controllers.NewAccountController()
	authController := controllers.NewAuthController()
	tradeController := controllers.NewTradeController()

	r.GET("/health", healthCheckHandler)

	v1 := r.Group("/v1")

	authGroup := v1.Group("")
	authGroup.Use(middleware.CheckAuthentication())
	authGroup.POST("/signup", authController.Signup)
	authGroup.GET("/balance", accountController.Balance)
	authGroup.POST("/buy_crypto", tradeController.BuyCrypto)
	authGroup.POST("/sell_crypto", tradeController.SellCrypto)
}

func healthCheckHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "healthy",
	})
}
