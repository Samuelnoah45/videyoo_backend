package notification_services

import (
	"context"
	graphqlClient "server/clients/graphql"
	notificationModel "server/pkgs/notification/models"
)

type notification_notifications_insert_input struct {
	Subject string `json:"subject"`
	Message string `json:"message"`
	UserId  string `json:"user_id"`
}

func SendNotification(notifications []notificationModel.Notification) (string, error) {

	// step 1: define mutations for notification
	var mutation struct {
		Insert_notification_notifications struct {
			Affected_rows int `json:"affected_rows"`
		} `graphql:"insert_notification_notifications(objects: $objects)"`
	}

	// step 2. Define variables
	var insertObjects []notification_notifications_insert_input

	for _, notification := range notifications {
		insertObjects = append(insertObjects, notification_notifications_insert_input{
			Subject: notification.Subject,
			Message: notification.Message,
			UserId:  notification.UserId,
		})
	}

	variables := map[string]interface{}{
		"objects": insertObjects,
	}

	// 4. Execute the request
	err := graphqlClient.SystemClient().Mutate(context.Background(), &mutation, variables)
	if err != nil {
		return "fail", err
	}
	return "message sent successfully", nil
}
