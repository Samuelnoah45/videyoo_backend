package auth_router

import (
	"github.com/gin-gonic/gin"

	"server/middlewares"
	authController "server/pkgs/auth/controllers"
)

func SetupRoutes(router *gin.Engine) {
	// Define routes for app1
	router.POST("/login", authController.Login)
	router.POST("/register", middlewares.AuthMiddleware(), authController.Register)

}
