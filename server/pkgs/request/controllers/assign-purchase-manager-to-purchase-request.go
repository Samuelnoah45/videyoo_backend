package request_controllers

// imports
import (
	"context"
	"fmt"
	"net/http"
	graphqlClient "server/clients/graphql"
	authService "server/pkgs/auth/services"
	notificationModel "server/pkgs/notification/models"
	notificationService "server/pkgs/notification/services"

	"github.com/gin-gonic/gin"
)

// check api controller
func AssignPurchaseManagerToPurchaseRequest(ctx *gin.Context) {

	xHasuraRole := ctx.GetString("x-hasura-role") // Retrieve a string value
	tokenString := ctx.GetString("tokenString")   // Retrieve a string value

	//step 1: get request data from body
	var inputData struct {
		Request_id string `json:"request_id"`
		User_id    string `json:"user_id"`
	}
	if err := ctx.ShouldBindJSON(&inputData); err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// step 2: update request
	// Define the GraphQL mutation
	var mutation struct {
		Update_request_purchase_requests_by_pk struct {
			ID string `json:"id"`
		} `graphql:"update_request_purchase_requests_by_pk(where: {id: {_eq: $id}}, _set: $set)"`
	}
	// Prepare variables for the mutation
	variables := map[string]interface{}{
		"id": inputData.Request_id,
		"set": map[string]interface{}{
			"purchase_manager_id": inputData.User_id,
		},
	}

	// Execute the GraphQL mutation
	err := graphqlClient.AuthClient(xHasuraRole, tokenString).Mutate(context.Background(), &mutation, variables)
	if err != nil {
		return
	}

	// step 3: create notification
	var notifications []notificationModel.Notification
	subject := "Purchase Request"

	// for general manager
	purchaseManger, err := authService.GetUser(inputData.User_id)
	if err != nil {
		fmt.Println(err.Error())
	}

	messagePurchaseManager := fmt.Sprintf("Hi %s. \n New purchase request has created and You assigned as purchase manager", purchaseManger.FirstName)
	notifications = append(notifications, notificationModel.Notification{
		Subject: subject,
		Message: messagePurchaseManager,
		UserId:  inputData.User_id,
	})

	// step 4: Send notification
	message, err := notificationService.SendNotification(notifications)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": message})
}
