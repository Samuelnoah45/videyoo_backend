package request_router

import (
	requestController "server/pkgs/request/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Request controller

	router.POST("/project_stock_out_request_on_create", requestController.ProjectStockOutRequestOnCreate)
	router.POST("/project_return_request_on_create", requestController.ProjectReturnRequestOnCreate)
	router.POST("/user_stock_out_request_on_create", requestController.UserStockOutRequestOnCreate)
	router.POST("/user_return_request_on_create", requestController.UserReturnRequestOnCreate)
	router.POST("/purchase_request_on_create", requestController.PurchaseRequestOnCreate)
	router.POST("/on_gm_verify_project_stock_out_request", requestController.OnGeneralManagerVerifyProjectStockOutRequest)
	router.POST("/on_tpm_verify_project_stock_out_request", requestController.OnTechnicalProjectManagerVerifyProjectStockOutRequest)
}
