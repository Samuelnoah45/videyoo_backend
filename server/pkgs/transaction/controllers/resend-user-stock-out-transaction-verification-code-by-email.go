package transaction_controllers

// imports
import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	authService "server/pkgs/auth/services"
	transactionModel "server/pkgs/transaction/models"
	transactionService "server/pkgs/transaction/services"
)

// check api controller
func ResendUserStockOutTransactionVerificationCodeByEmail(ctx *gin.Context) {

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
	ctx.JSON(200, gin.H{"emailMessage": sendResponse})
}
