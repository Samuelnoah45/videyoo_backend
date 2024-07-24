package transaction_controllers

// imports
import (
	"fmt"
	"net/http"

	authService "server/pkgs/auth/services"
	notificationModel "server/pkgs/notification/models"
	notificationService "server/pkgs/notification/services"
	transactionService "server/pkgs/transaction/services"

	"github.com/gin-gonic/gin"
)

// check api controller
func VerifyPurchaseTransaction(ctx *gin.Context) {
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
	purchaseTransaction, dataFetchError := transactionService.GetPurchaseTransaction(inputData.Transaction_id)
	if dataFetchError != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": dataFetchError.Error()})
	}

	// check if transaction is already verified
	if !purchaseTransaction.Is_verified {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Transaction is already verified"})
		return
	}

	// update is verified field of project stock out transaction
	err := transactionService.UpdatePurchaseTransaction(inputData.Transaction_id, true, xHasuraRole, tokenString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// modify the warehouse
	transaction, err := transactionService.GetTransaction(purchaseTransaction.Transaction_id, xHasuraRole, tokenString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		transactionService.UpdatePurchaseTransaction(inputData.Transaction_id, false, xHasuraRole, tokenString)
	}

	fmt.Println(transaction)
	// err = transactionService.ProcessStockInTransaction(transaction)
	// if err != nil {
	// 	transactionService.UpdatePurchaseTransaction(inputData.Transaction_id, false, xHasuraRole, tokenString)
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// }

	// step 6: create notification
	purchaseManger, getUserError := authService.GetUser(purchaseTransaction.Purchase_manager_id)
	if getUserError != nil {
		fmt.Println(getUserError.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": getUserError.Error()})
		return
	}
	var notifications []notificationModel.Notification
	subject := "Transaction Verification"

	messageToUser := fmt.Sprintf("Hi %s. \nTransaction has been verified", purchaseManger.FirstName)
	notifications = append(notifications, notificationModel.Notification{
		Subject: subject,
		Message: messageToUser,
		UserId:  purchaseManger.ID,
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
