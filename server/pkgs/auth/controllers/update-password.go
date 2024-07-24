package auth_controllers

// import (
// 	"context"
// 	graphqlClient "server/clients/graphql"
// 	"server/models"
// 	"server/utilService"
// 	"fmt"
// 	"github.com/gin-gonic/gin"
// )

// // imports

// func UpdatePassword(ctx *gin.Context) {

// 	var updateInput struct {
// 		Password     string `json:"password"`
// 		New_password string `json:"new_password"`
// 		Email        string `json:"email"`
// 	}

// 	var query struct {
// 		Users []struct {
// 			ID                  string `json:"id"`
// 			First_name          string `json:"first_name"`
// 			Last_name           string `json:"last_name"`
// 			Email               string `json:"email"`
// 			Password            string `json:"password"`
// 			Photo_url           string `json:"photo_url"`
// 			Phone_number        string `json:"phone_number"`
// 			Is_account_verified bool   `json:"is_account_verified"`
// 			User_roles          []struct {
// 				Role_name string `json:"role_name"`
// 			} `json:"user_roles"`
// 		} `graphql:"user_users(where: {email: {_eq: $email}})"`
// 	}

// 	variables := map[string]interface{}{
// 		"email": updateInput.Email,
// 	}
// 	err := graphqlClient.AnonymousClient().Query(context.Background(), &query, variables)
// 	if err != nil {
// 		fmt.Println(err.Error(), "when querying user")
// 		ctx.JSON(400, gin.H{"message": err.Error()})
// 		return
// 	}

// 	if len(query.Users) > 0 && !query.Users[0].Is_account_verified {
// 		ctx.JSON(400, gin.H{"message": "Unverified account"})
// 		return
// 	}
// 	user_roles := []string{}
// 	for _, role := range query.Users[0].User_roles {
// 		user_roles = append(user_roles, role.Role_name)
// 	}

// 	if len(query.Users) > 0 && utilService.ComparePasswords(query.Users[0].Password, updateInput.Password) {
// 		var user models.UserData
// 		user.ID = query.Users[0].ID
// 		user.Email = query.Users[0].Email
// 		user.FirstName = query.Users[0].First_name
// 		user.LastName = query.Users[0].Last_name
// 		user.PhoneNumber = query.Users[0].Phone_number
// 		user.PhotoUrl = query.Users[0].Photo_url
// 		user.UserRoles = user_roles
// 		sendTokenAndUserData(ctx, user)
// 		return
// 	} else {
// 		ctx.JSON(400, gin.H{"message": "Invalid credentials"})
// 		return
// 	}
// 	// Get all headers as a map

// 			// 1.Define the GraphQL mutation string
// 			var updateMutation struct {
// 				UpdateUsers struct {
// 					Returning []struct{
// 						ID string `json:"id"`
// 						Role string `json:"role"`
// 						Email string `json:"email"`
// 					} `json:"returning"`
// 				} `graphql:"update_users(where: {email: {_eq: $email}}, _set: {password: $password, firstName: $firstName, lastName: $lastName})"`
// 			}
// 	// 		//2. set variable

// 			muatateVariables := map[string]interface{}{
// 				"password":  user.FirstName,

// 			}
// 	// 		//3. execute the request
// 			err5 :=graphqlClient.SystemClient().Mutate(context.Background(), &updateMutation, variables)
// 			if err5 != nil {
// 				ctx.JSON(400, gin.H{"error": err5.Error()})
// 				return
// 			}

// 		ctx.JSON(200, gin.H{"message": "User data updated"})

// }
