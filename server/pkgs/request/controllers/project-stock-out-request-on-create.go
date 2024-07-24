package request_controllers

// imports
import (
	"fmt"
	"net/http"

	notificationModel "server/pkgs/notification/models"
	notificationService "server/pkgs/notification/services"
	projectService "server/pkgs/project/services"
	requestModel "server/pkgs/request/models"

	"github.com/gin-gonic/gin"
)

// check api controller
func ProjectStockOutRequestOnCreate(ctx *gin.Context) {

	//step 1: get request data from body
	var request requestModel.ProjectStockOutRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// step 2: get project
	project, err := projectService.GetProject(request.Project_id)
	if err != nil {
		fmt.Println(err.Error())

		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	// step 3: create notification

	var notifications []notificationModel.Notification
	subject := fmt.Sprintf("New Stock Request from %s project", project.Name)
	// for technical project manager
	messageToTechnicalProjectManager := fmt.Sprintf("Hi %s. \n  %s project has new stock request created by %s", project.Technical_project_manager.FirstName, project.Name, project.Project_manager.FirstName)
	notifications = append(notifications, notificationModel.Notification{
		Subject: subject,
		Message: messageToTechnicalProjectManager,
		UserId:  request.Technical_project_manager_id,
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
