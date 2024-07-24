package transaction_controllers

// imports
import (
	"fmt"
	"net/http"

	transactionService "server/pkgs/transaction/services"

	"github.com/gin-gonic/gin"
)

// check api controller
func VerifyUserStockOutTransaction(ctx *gin.Context) {

	// step 1: get role and token from context
	xHasuraRole := ctx.GetString("x-hasura-role") // Retrieve a string value
	tokenString := ctx.GetString("tokenString")   // Retrieve a string value
	//step 1: get request data from body
	var inputData struct {
		Transaction_id   string `json:"transaction_id"`
		VerificationCode string `json:"verification_code"`
	}
	if dataBindError := ctx.ShouldBindJSON(&inputData); dataBindError != nil {
		fmt.Println(dataBindError.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": dataBindError.Error()})
		return
	}
	userStockOutTransaction, dataFetchError := transactionService.GetUserStockOutTransaction(inputData.Transaction_id)
	if dataFetchError != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": dataFetchError.Error()})
	}

	// check if transaction is already verified
	if !userStockOutTransaction.Is_verified {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Transaction is already verified"})
		return
	}

	// check if verification code is correct
	if !transactionService.ValidatedVerificationCode(inputData.VerificationCode, userStockOutTransaction.Transaction_verification_code) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Transaction code is Invalid"})
		return
	}

	// update is verified field of project stock out transaction

	err := transactionService.UpdateUserStockOutTransaction(inputData.Transaction_id, true)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// modify the warehouse
	transaction, err := transactionService.GetTransaction(userStockOutTransaction.Transaction_id, xHasuraRole, tokenString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		transactionService.UpdateUserStockOutTransaction(inputData.Transaction_id, false)
	}

	fmt.Println(transaction)

	// err = transactionService.ProcessStockOutTransaction(transaction)
	// if err != nil {
	// 	transactionService.UpdateUserStockOutTransaction(inputData.Transaction_id, false)
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// }

}
