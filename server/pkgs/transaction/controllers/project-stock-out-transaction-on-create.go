package transaction_controllers

// imports
import (
	"fmt"
	"net/http"

	projectService "server/pkgs/project/services"

	notificationModel "server/pkgs/notification/models"
	notificationService "server/pkgs/notification/services"
	transactionModel "server/pkgs/transaction/models"
	transactionService "server/pkgs/transaction/services"

	"github.com/gin-gonic/gin"
)

// check api controller
func ProjectStockOutTransactionOnCreate(ctx *gin.Context) {

	//step 1: get request data from body
	var transaction transactionModel.ProjectStockOutTransaction
	if dataBindError := ctx.ShouldBindJSON(&transaction); dataBindError != nil {
		fmt.Println(dataBindError.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": dataBindError.Error()})
		return
	}

	// step 2: get project
	project, getProjectError := projectService.GetProject(transaction.Project_id)
	if getProjectError != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": getProjectError.Error()})
		return
	}

	// step 4: generate code
	verificationCode, verificationCodeGenerationError := transactionService.GenerateTransactionVerificationCode()
	if verificationCodeGenerationError != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": verificationCodeGenerationError.Error()})
		return
	}

	// step 5: send code by email
	sendResponse, emailSendError := transactionService.SendProjectStockOutTransactionVerificationCodeByEmail(project.Project_manager, project.Name, verificationCode)
	if emailSendError != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": emailSendError.Error()})

	}
	// if code sended successfully store code on database

	// step 6: create notification
	var notifications []notificationModel.Notification
	subject := "Transaction Verification"

	messageToProjectManager := fmt.Sprintf("Hi %s. \n Transaction has created  for  %s and we have sent verification code to your email please give this number to stock manager", project.Project_manager.FirstName, project.Name)
	notifications = append(notifications, notificationModel.Notification{
		Subject: subject,
		Message: messageToProjectManager,
		UserId:  transaction.Project_manager_id,
	})

	// step 7: Send notification
	message, sendNotificationError := notificationService.SendNotification(notifications)
	if sendNotificationError != nil {
		fmt.Println("error", sendNotificationError.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": sendNotificationError.Error()})
		return
	}
	ctx.JSON(200, gin.H{"notificationMessage": message, "emailMessage": sendResponse})

}
