package auth_router

import (
	"github.com/gin-gonic/gin"

	authController "server/pkgs/auth/controllers"
)

func SetupRoutes(router *gin.Engine) {
	// Define routes for app1
	router.POST("/login", authController.Login)
	router.POST("/sign_up", authController.Register)

}
