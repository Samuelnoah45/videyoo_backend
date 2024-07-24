package request_controllers

// imports
import (
	"fmt"
	"net/http"
	authService "server/pkgs/auth/services"
	notificationModel "server/pkgs/notification/models"
	notificationService "server/pkgs/notification/services"
	projectService "server/pkgs/project/services"

	"github.com/gin-gonic/gin"
)

// check api controller
func OnGeneralManagerVerifyProjectStockOutRequest(ctx *gin.Context) {

	//step 1: get request data from body
	var request struct {
		Project_id       string `json:"project_id"`
		Stock_manager_id string `json:"stock_manager_id"`
		Old_verify_value bool   `json:"old_verify_value"`
		New_verify_value bool   `json:"new_verify_value"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// step 2: get project
	project, err := projectService.GetProject(request.Project_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	// step 3: create notification

	var notifications []notificationModel.Notification
	subject := fmt.Sprintf("New Stock Request from %s project", project.Name)

	// for general manager
	stockManager, err := authService.GetUser(request.Stock_manager_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

	}
	messageToStockManager := fmt.Sprintf("Hi %s. \n  %s project has new stock request created by %s", stockManager.FirstName, project.Name, project.Project_manager.FirstName)
	notifications = append(notifications, notificationModel.Notification{
		Subject: subject,
		Message: messageToStockManager,
		UserId:  request.Stock_manager_id,
	})

	// step 4: Send notification
	message, err := notificationService.SendNotification(notifications)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return

	}
	ctx.JSON(200, gin.H{"message": message})

}
