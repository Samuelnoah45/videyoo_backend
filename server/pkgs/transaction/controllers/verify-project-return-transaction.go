package transaction_controllers

// imports
import (
	"fmt"
	"net/http"

	notificationModel "server/pkgs/notification/models"
	notificationService "server/pkgs/notification/services"
	projectService "server/pkgs/project/services"
	transactionService "server/pkgs/transaction/services"

	"github.com/gin-gonic/gin"
)

// check api controller
func VerifyProjectReturnTransaction(ctx *gin.Context) {
	// step 1: get role and token from context
	xHasuraRole := ctx.GetString("x-hasura-role") // Retrieve a string value
	tokenString := ctx.GetString("tokenString")   // Retrieve a string value
	//step 1: get request data from body
	var inputData struct {
		Transaction_id string `json:"transaction_id"`
	}
	if dataBindError := ctx.ShouldBindJSON(&inputData); dataBindError != nil {
		fmt.Println(dataBindError.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": dataBindError.Error()})
		return
	}
	projectReturnTransaction, dataFetchError := transactionService.GetProjectReturnTransaction(inputData.Transaction_id)
	if dataFetchError != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": dataFetchError.Error()})
	}

	// check if transaction is already verified
	if !projectReturnTransaction.Is_verified {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Transaction is already verified"})
		return
	}

	// update is verified field of project stock out transaction

	err := transactionService.UpdateProjectReturnTransaction(inputData.Transaction_id, true, xHasuraRole, tokenString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// modify the warehouse
	transaction, err := transactionService.GetTransaction(projectReturnTransaction.Transaction_id, xHasuraRole, tokenString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		transactionService.UpdateProjectReturnTransaction(inputData.Transaction_id, false, xHasuraRole, tokenString)
	}
	fmt.Println(transaction)
	// err = transactionService.ProcessStockInTransaction(transaction)
	// if err != nil {
	// 	transactionService.UpdateProjectReturnTransaction(inputData.Transaction_id, false, xHasuraRole, tokenString)
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// }

	// step 6: create notification
	// step 2: get project
	project, getProjectError := projectService.GetProject(projectReturnTransaction.Project_id)
	if getProjectError != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": getProjectError.Error()})
		return
	}

	// step 6: create notification
	var notifications []notificationModel.Notification
	subject := "Transaction Verification"
	messageToProjectManager := fmt.Sprintf("Hi %s. \n Return request has verified  for  %s ", project.Project_manager.FirstName, project.Name)
	notifications = append(notifications, notificationModel.Notification{
		Subject: subject,
		Message: messageToProjectManager,
		UserId:  projectReturnTransaction.Project_manager_id,
	})
	// step 7: Send notification
	message, sendNotificationError := notificationService.SendNotification(notifications)
	if sendNotificationError != nil {
		fmt.Println("error", sendNotificationError.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": sendNotificationError.Error()})
		return
	}
	ctx.JSON(200, gin.H{"notificationMessage": message})
}
