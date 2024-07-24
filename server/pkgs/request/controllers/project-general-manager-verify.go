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
	projectService "server/pkgs/project/services"
	requestService "server/pkgs/request/services"

	"github.com/gin-gonic/gin"
)

// check api controller
func ProjectGeneralManagerVerify(ctx *gin.Context) {

	// step 1: get role and token from context
	xHasuraRole := ctx.GetString("x-hasura-role") // Retrieve a string value
	tokenString := ctx.GetString("tokenString")   // Retrieve a string value
	// step 2: get  data from request body
	var inputData struct {
		Project_stock_out_request_id string `json:"project_stock_out_request_id"`
	}
	if err := ctx.ShouldBindJSON(&inputData); err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// step 3: Get request information
	request, err := requestService.GetProjectStockOutRequest(inputData.Project_stock_out_request_id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err})

	}

	// step 4: Define the GraphQL mutation to verify the request
	var mutation struct {
		Update_notification_notifications struct {
			Affected_rows int `json:"affected_rows"`
		} `graphql:"update_notification_notifications(where: {id: {_eq: $id}}, _set: $set)"`
	}

	// Prepare variables for the mutation
	variables := map[string]interface{}{
		"id": inputData.Project_stock_out_request_id,
		"set": map[string]interface{}{
			"Is_General_manager_approved": true,
		},
	}

	// Execute the GraphQL mutation
	err = graphqlClient.AuthClient(xHasuraRole, tokenString).Mutate(context.Background(), &mutation, variables)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	// Step 5: Send message to general manager
	// 5.1 Get project
	// // step 2: get project
	project, err := projectService.GetProject(request.Project_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// // step 3: create notification

	//5.2. Send message to general manager
	var notifications []notificationModel.Notification
	subject := fmt.Sprintf("New Stock Request from %s project", project.Name)

	// // for general manager
	stockManger, err := authService.GetUser(request.Stock_manager_id)
	if err != nil {
		fmt.Println(err.Error())
	}
	messageToStockManager := fmt.Sprintf("Hi %s. \n  %s project has new stock request created by %s", stockManger.FirstName, project.Name, project.Project_manager.FirstName)
	notifications = append(notifications, notificationModel.Notification{
		Subject: subject,
		Message: messageToStockManager,
		UserId:  request.Stock_manager_id,
	})
	message, err := notificationService.SendNotification(notifications)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(200, gin.H{"message": message})
}
