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
func PurchaseRequestOnCreate(ctx *gin.Context) {
	//step 1: get request data from body
	var request requestModel.PurchaseRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	// step 3: create notification
	var notifications []notificationModel.Notification
	subject := "Purchase Request"

	// for general manager
	generalManger, err := authService.GetUser(request.General_manager_id)
	if err != nil {
		fmt.Println(err.Error())
	}

	messageToGeneralManager := fmt.Sprintf("Hi %s. \n New purchase request has created", generalManger.FirstName)
	notifications = append(notifications, notificationModel.Notification{
		Subject: subject,
		Message: messageToGeneralManager,
		UserId:  request.General_manager_id,
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
