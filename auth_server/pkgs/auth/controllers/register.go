package auth_controllers

// imports
import (
	"context"
	"fmt"
	"net/http"
	graphqlClient "server/clients/graphql"
	authModel "server/pkgs/auth/models"
	"server/utilService"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type user_roles_arr_rel_insert_input map[string]interface{}
type roles_insert_input map[string]interface{}

// signup controller
func Register(ctx *gin.Context) {

	// // Read the raw request body
	// bodyBytes, bodyErr := ioutil.ReadAll(ctx.Request.Body)
	// if bodyErr != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to read request body"})
	// 	return
	// }

	// // Print the raw request body
	// fmt.Println("Raw request body:", string(bodyBytes))

	var validate = validator.New()

	//1. Get the user input from the request body

	type RequestBody struct {
		Input struct {
			FirstName string `json:"first_name" validate:"required"`
			LastName  string `json:"last_name" validate:"required"`
			Password  string `json:"password" validate:"required,min=6"`
			Email     string `json:"email" validate:"required,email"`
		}
	}
	var requestBody RequestBody
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	fmt.Println("###############", requestBody)
	var newUser = authModel.User{
		Email:     requestBody.Input.Email,
		FirstName: requestBody.Input.FirstName,
		LastName:  requestBody.Input.LastName,
		Password:  requestBody.Input.Password,
		UserRoles: []string{"user"},
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
		} `graphql:"insert_users(objects: {first_name: $first_name, last_name: $last_name,  password: $password, email: $email, user_roles: $user_roles})"`
	}

	//4.  construct graphql variable

	var userRoles []roles_insert_input
	for _, role := range newUser.UserRoles {
		userRoles = append(userRoles, roles_insert_input{
			"role_name": role,
		})
	}
	userRoles = append(userRoles, roles_insert_input{
		"role_name": "user",
	})
	password, hashErr := utilService.HashPassword(newUser.Password)

	if hashErr != nil {
		ctx.JSON(400, gin.H{"message": hashErr.Error()})
		return
	}
	variables := map[string]interface{}{
		"first_name": newUser.FirstName,
		"last_name":  newUser.LastName,
		"email":      newUser.Email,
		"password":   password,
		"user_roles": user_roles_arr_rel_insert_input{
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

	token, tokenError := utilService.HasuraAccessToken(newUser)
	if tokenError != nil {
		ctx.JSON(400, gin.H{"message": tokenError.Error()})
		return
	}

	ctx.JSON(200, gin.H{"token": token, "success": true})
}
