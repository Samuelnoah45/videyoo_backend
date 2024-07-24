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
func UserStockOutRequestOnCreate(ctx *gin.Context) {

	//step 1: get request data from body
	var request requestModel.UserStockOutRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// step 3: create notification
	var notifications []notificationModel.Notification
	user, err := authService.GetUser(request.User_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	stockManger, err := authService.GetUser(request.Stock_manager_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	messageToStockManager := fmt.Sprintf("Hi %s. \n  %s  has stock out  request", stockManger.FirstName, user.FirstName)
	notifications = append(notifications, notificationModel.Notification{
		Subject: "User stock out request",
		Message: messageToStockManager,
		UserId:  request.Stock_manager_id,
	})

	// step 4: Send notification
	message, err := notificationService.SendNotification(notifications)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": message})

}
