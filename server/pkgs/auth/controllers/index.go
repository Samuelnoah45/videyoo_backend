package auth_controllers

import (
	"fmt"

	authModel "server/pkgs/auth/models"

	"server/utilService"

	"github.com/gin-gonic/gin"
)

// this function accept payload to create token and then call util service with payload
// finally send token to client

// to send email to user
type EmailDataToken struct {
	Link   string
	Header string
}

func sendTokenAndUserData(ctx *gin.Context, user authModel.User) {
	token, err := utilService.HasuraAccessToken(user)

	if err != nil {
		fmt.Println(err.Error(), "when sending token ")

		ctx.JSON(400, gin.H{"message": "Something went wrong when creating token"})
		return
	}
	var response authModel.AuthResponse
	response.ID = user.ID
	response.User = user
	response.HasuraAccessToken = token
	ctx.JSON(200, response)
}
