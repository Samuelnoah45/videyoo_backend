package config

var (
	DEBUG                              bool
	PORT                               string
	HASURA_GRAPHQL_URL                 string
	HASURA_GRAPHQL_ADMIN_SECRET        string
	SERVICE_NAME                       string
	SMTP_HOST                          string
	SMTP_PASSWORD                      string
	SMTP_USERNAME                      string
	SMTP_PORT                          string
	GIN_MODE                           string
	ALG                                string
	SEND_GRID_API_KEY                  string
	SEND_GRID_URL                      string
	SEND_GRID_USER_NAME                string
	SEND_GRID_SENDER_NAME              string
	SEND_GRID_CONTACT_US_TEMPLATE_ID   string
	SEND_GRID_SEND_TO_EMAIL            string
	AWS_ACCESS_KEY                     string
	AWS_SECRET_ACCESS_KEY              string
	AFRO_SMS_TOKEN                     string
	AFRO_SMS_SENDER                    string
	SEND_GRID_SUBSCRIPTION_TEMPLATE_ID string
	SEND_GRID_BOOKING_TEMPLATE_ID      string
	SERVER_PDF_URL                     string
	AFRO_URL                           string
	AFRO_SMS_ACCESS_TOKEN              string
	RESET_PASSWORD_URL                 string

	JWT_SECRET_KEY          string
	VERIFICATION_SECRET_KEY string
	CLOUDINARY_CLOUD_NAME   string
	CLOUDINARY_API_KEY      string
	CLOUDINARY_SECRET       string

	DB_PORT     string
	HOST        string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
)
