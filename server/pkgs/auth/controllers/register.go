package auth_controllers

// imports
import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	graphqlClient "server/clients/graphql"
	authModel "server/pkgs/auth/models"
	authService "server/pkgs/auth/services"
)

type user_user_roles_arr_rel_insert_input map[string]interface{}
type user_user_roles_insert_input map[string]interface{}

// signup controller
func Register(ctx *gin.Context) {
	var validate = validator.New()

	userId := ctx.GetString("user-id") // Retrieve a string value

	//1. Get the user input from the request body
	var newUser authModel.User
	// var newUser inputUser
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	fmt.Println(newUser)

	if err := validate.Struct(newUser); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var errorMsg string
		for _, e := range validationErrors {
			errorMsg = fmt.Sprintf("Field %s validation failed on the %s tag", e.Field(), e.Tag())
			break
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"message": errorMsg})
		return
	}
	// //3. Define the GraphQL mutation string
	var mutation struct {
		InsertUsers struct {
			Returning []struct {
				ID string `json:"id"`
			} `json:"returning"`
		} `graphql:"insert_user_users(objects: {first_name: $first_name, last_name: $last_name, email: $email, user_roles: $user_roles})"`
	}

	//4.  construct graphql variable

	var userRoles []user_user_roles_insert_input
	for _, role := range newUser.UserRoles {
		userRoles = append(userRoles, user_user_roles_insert_input{
			"role_name":  role,
			"created_by": userId,
		})
	}
	userRoles = append(userRoles, user_user_roles_insert_input{
		"role_name":  "org-member",
		"created_by": userId,
	})

	variables := map[string]interface{}{
		"first_name": newUser.FirstName,
		"last_name":  newUser.LastName,
		"email":      newUser.Email,
		"user_roles": user_user_roles_arr_rel_insert_input{
			"data": userRoles,
		},
	}
	// //5. execute the request
	err := graphqlClient.SystemClient().Mutate(context.Background(), &mutation, variables)
	if err != nil {
		fmt.Println(err.Error(), "when executing register  mutation")
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}
	message, sendError := authService.SendResetTokenByEmail(newUser.Email, "http://localhost:3000")
	if sendError != nil {
		ctx.JSON(400, gin.H{"message": sendError.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": message, "success": true})
}
