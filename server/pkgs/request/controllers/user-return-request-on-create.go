package request_controllers

// imports
import (
	"fmt"
	"net/http"
	authService "server/pkgs/auth/services"
	notificationModel "server/pkgs/notification/models"
	notificationService "server/pkgs/notification/services"
	requestModel "server/pkgs/request/models"

	"github.com/gin-gonic/gin"
)

// check api controller
func UserReturnRequestOnCreate(ctx *gin.Context) {

	//step 1: get request data from body
	var request requestModel.UserReturnRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// step 3: create notification
	var notifications []notificationModel.Notification

	user, err := authService.GetUser(request.User_id)
	if err != nil {
		fmt.Println(err.Error())
	}

	stockManger, err := authService.GetUser(request.Stock_manager_id)
	if err != nil {
		fmt.Println(err.Error())
	}

	messageToStockManager := fmt.Sprintf("Hi %s. \n  %s  has new return request", stockManger.FirstName, user.FirstName)
	notifications = append(notifications, notificationModel.Notification{
		Subject: "User return request",
		Message: messageToStockManager,
		UserId:  request.Stock_manager_id,
	})

	// step 4: Send notification
	message, err := notificationService.SendNotification(notifications)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(200, gin.H{"message": message})

}
