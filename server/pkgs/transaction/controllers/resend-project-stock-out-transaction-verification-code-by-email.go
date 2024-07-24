package transaction_controllers

// imports
import (
	"fmt"
	"net/http"

	projectService "server/pkgs/project/services"
	transactionService "server/pkgs/transaction/services"

	"github.com/gin-gonic/gin"
)

// check api controller
func ResendProjectStockOutTransactionVerificationCodeByEmail(ctx *gin.Context) {

	//step 1: get request data from body
	var inputData struct {
		Transaction_id string `json:"transaction_id"`
	}
	if dataBindError := ctx.ShouldBindJSON(&inputData); dataBindError != nil {
		fmt.Println(dataBindError.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": dataBindError.Error()})
		return
	}

	transaction, dataFetchError := transactionService.GetProjectStockOutTransaction(inputData.Transaction_id)
	if dataFetchError != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": dataFetchError.Error()})
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

	ctx.JSON(200, gin.H{"emailMessage": sendResponse})
}
