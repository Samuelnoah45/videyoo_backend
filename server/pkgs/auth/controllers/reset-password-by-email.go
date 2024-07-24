package auth_controllers

// imports
import (
	"context"
	"fmt"
	"net/http"
	graphqlClient "server/clients/graphql"
	authModel "server/pkgs/auth/models"
	"strings"

	"server/utilService"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

// Reset Password by email controller
func ResetPasswordByEmail(ctx *gin.Context) {
	var validate = validator.New()
	//1. get user new password  data from request body
	var inputData struct {
		Password                  string `json:"password" validate:"required,min=6"`
		Email                     string `json:"email" validate:"required,email,min=6"`
		ResetPasswordByEmailToken string `json:"reset_password_by_email_token" validate:"required"`
	}

	if dataBindError := ctx.ShouldBindJSON(&inputData); dataBindError != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": dataBindError.Error()})
		return
	}

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

	//3. Parse the token
	token, validateError := utilService.ValidateVerificationToken(inputData.ResetPasswordByEmailToken)
	if validateError != nil {
		fmt.Println(token, "when parsing token")
		ctx.JSON(400, gin.H{"message": "invalid token"})
		// delete token from database
		return
	}
	//4. If the token is valid then search for the user with the token
	var query struct {
		Users []struct {
			ID                  string `json:"id"`
			First_name          string `json:"first_name"`
			Last_name           string `json:"last_name"`
			Email               string `json:"email"`
			Password            string `json:"password"`
			Photo_url           string `json:"photo_url"`
			Phone_number        string `json:"phone_number"`
			Gender              string `json:"gender"`
			Is_account_verified bool   `json:"is_account_verified"`
			User_roles          []struct {
				Role_name string `json:"role_name"`
			} `json:"user_roles"`
			Reset_password_by_email_token string `json:"reset_password_by_email_token"`
		} `graphql:"user_users(where: {reset_password_by_email_token: {_eq: $resetPasswordByEmailToken}})"`
	}
	//5. construct graphql variables
	variables := map[string]interface{}{
		"resetPasswordByEmailToken": inputData.ResetPasswordByEmailToken,
	}
	//6. execute the request
	fetchError := graphqlClient.SystemClient().Query(context.Background(), &query, variables)
	if fetchError != nil {
		fmt.Println(fetchError.Error())
		ctx.JSON(400, gin.H{"message": "Something went wrong"})
		return
	}
	//7. If the user exists then update the password
	if len(query.Users) > 0 && query.Users[0].Reset_password_by_email_token == inputData.ResetPasswordByEmailToken {
		// if the user exists, send the token with user data
		// change password
		password, hashPasswordError := utilService.HashPassword(inputData.Password)
		if hashPasswordError != nil {
			ctx.JSON(400, gin.H{"message": hashPasswordError.Error()})
			return
		}
		// 1.Define the GraphQL mutation string
		var mutation struct {
			UpdateUsers struct {
				Returning []struct {
					ID string `json:"id"`
				} `json:"returning"`
			} `graphql:"update_user_users(where: {email: {_eq: $email}}, _set: {password: $password, reset_password_by_email_token: $resetToken, is_account_verified:$isAccountVerified})"`
		}

		//2. set variable
		updateVariables := map[string]interface{}{
			"password":          password,
			"resetToken":        "",
			"email":             query.Users[0].Email,
			"isAccountVerified": true,
		}
		//3. execute the request
		updateError := graphqlClient.SystemClient().Mutate(context.Background(), &mutation, updateVariables)
		if updateError != nil {
			ctx.JSON(400, gin.H{"message": updateError.Error()})
			return
		}

		// send data to user
		user_roles := []string{}
		for _, role := range query.Users[0].User_roles {
			editedRole := strings.Replace(role.Role_name, "_", "-", -1)
			user_roles = append(user_roles, editedRole)
		}
		var user authModel.User
		user.ID = query.Users[0].ID
		user.Email = query.Users[0].Email
		user.FirstName = query.Users[0].First_name
		user.LastName = query.Users[0].Last_name
		user.PhoneNumber = query.Users[0].Phone_number
		user.PhotoUrl = query.Users[0].Photo_url
		user.UserRoles = user_roles
		user.Gender = query.Users[0].Gender
		sendTokenAndUserData(ctx, user)

	} else {

		ctx.JSON(400, gin.H{"message": "Invalid credentials"})
	}
}
