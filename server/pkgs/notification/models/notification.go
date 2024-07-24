package notification_models

type Notification struct {
	Id      string `json:"id"`
	Subject string `json:"subject"`
	Message string `json:"message"`
	UserId  string `json:"user_id"`
}
