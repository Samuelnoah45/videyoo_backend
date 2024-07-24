package transaction_controllers

// imports
import (
	"fmt"
	"net/http"

	authService "server/pkgs/auth/services"
	notificationModel "server/pkgs/notification/models"
	notificationService "server/pkgs/notification/services"
	transactionModel "server/pkgs/transaction/models"
	transactionService "server/pkgs/transaction/services"

	"github.com/gin-gonic/gin"
)

// check api controller
func UserStockOutTransactionOnCreate(ctx *gin.Context) {

	//step 1: get request data from body
	var transaction transactionModel.UserStockOutTransaction
	if dataBindError := ctx.ShouldBindJSON(&transaction); dataBindError != nil {
		fmt.Println(dataBindError.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": dataBindError.Error()})
		return
	}
	// step 2: generate code
	verificationCode, verificationCodeGenerationError := transactionService.GenerateTransactionVerificationCode()
	if verificationCodeGenerationError != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": verificationCodeGenerationError.Error()})
		return
	}

	// step 3: get user
	user, getUserError := authService.GetUser(transaction.User_id)

	if getUserError != nil {
		fmt.Println(getUserError.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": getUserError.Error()})
		return
	}

	// step 5: send code by email
	sendResponse, emailSendError := transactionService.SendUserStockOutTransactionVerificationCodeByEmail(user, verificationCode)
	if emailSendError != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": emailSendError.Error()})

	}
	// step 6: create notification
	var notifications []notificationModel.Notification
	subject := "Transaction Verification"

	messageToUser := fmt.Sprintf("Hi %s. \n Transaction has created,  And we have sent verification code to your email please give this number to stock manager", user.FirstName)
	notifications = append(notifications, notificationModel.Notification{
		Subject: subject,
		Message: messageToUser,
		UserId:  transaction.User_id,
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
