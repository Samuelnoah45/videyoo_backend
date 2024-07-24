package auth_router

import (
	"server/middlewares"
	authController "server/pkgs/auth/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Define routes for app1
	router.POST("/login", authController.Login)
	router.POST("/register", middlewares.AuthMiddleware(), authController.Register)
	router.POST("/verifyEmailOrForgotPassword", authController.VerifyEmailOrForgotPassword)
	router.POST("/resetPasswordByEmail", authController.ResetPasswordByEmail)
}
