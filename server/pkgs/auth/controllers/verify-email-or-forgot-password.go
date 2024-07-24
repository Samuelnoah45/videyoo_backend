package auth_controllers

// imports
import (
	"fmt"
	"net/http"

	authService "server/pkgs/auth/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func VerifyEmailOrForgotPassword(ctx *gin.Context) {
	var validate = validator.New()

	// 1. Get the user input from the request body
	var inputData struct {
		Email     string `json:"email" validate:"required,email,min=6"`
		Reset_url string `json:"reset_url"`
	}
	if dataBindError := ctx.ShouldBindJSON(&inputData); dataBindError != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": dataBindError.Error()})
		return
	}

	// validate email
	if err := validate.Struct(inputData); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var errorMsg string
		for _, e := range validationErrors {
			errorMsg = fmt.Sprintf("Field %s validation failed on the %s tag", e.Field(), e.Tag())
			break
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"message": errorMsg})
		return
	}
	message, sendError := authService.SendResetTokenByEmail(inputData.Email, inputData.Reset_url)
	if sendError != nil {
		ctx.JSON(400, gin.H{"message": sendError.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": message, "success": true})

}
