package transaction_router

import (
	"server/middlewares"
	transactionController "server/pkgs/transaction/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	// 2. Transaction controller
	// 2.1 Send verification
	router.POST("/project_stock_out_transaction_on_create", transactionController.ProjectStockOutTransactionOnCreate)
	router.POST("/user_stock_out_transaction_on_create", middlewares.AuthMiddleware(), transactionController.UserStockOutTransactionOnCreate)
	router.POST("/resend_project_stock_out_transaction_verification_code_by_email", transactionController.ResendProjectStockOutTransactionVerificationCodeByEmail)
	router.POST("/resend_user_stock_out_transaction_verification_code_by_email", transactionController.ResendUserStockOutTransactionVerificationCodeByEmail)

	//2.2 Verify transaction
	//2.2.2.Stock out transaction
	router.POST("/verify_project_stock_out_transaction", transactionController.VerifyProjectStockOutTransaction)
	router.POST("/verify_user_stock_out_transaction", transactionController.VerifyUserStockOutTransaction)
	//2.2.2.Stock in transaction
	router.POST("/verify_purchase_transaction", transactionController.VerifyPurchaseTransaction)
	router.POST("/verify_user_return_transaction", transactionController.VerifyUserReturnTransaction)
	router.POST("/verify_project_return_transaction", transactionController.VerifyProjectReturnTransaction)

}
