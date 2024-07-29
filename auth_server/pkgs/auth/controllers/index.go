package auth_controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"

	authModel "server/pkgs/auth/models"
	"server/utilService"
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

	ctx.JSON(200, gin.H{"token": token, "success": true})
}
