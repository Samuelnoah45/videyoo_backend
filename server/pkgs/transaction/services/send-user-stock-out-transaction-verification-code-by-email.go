package transaction_services

// imports
import (
	"bytes"
	"fmt"
	"html/template"
	authModel "server/pkgs/auth/models"

	"server/utilService"
)

type UserStockOutTransactionVerificationByEmailTemplateData struct {
	UserName         string `json:"userName"`
	VerificationCode string `json:"verificationCode"`
}

func SendUserStockOutTransactionVerificationCodeByEmail(user authModel.User, verificationCode string) (string, error) {

	t, _ := template.ParseFiles("./templates/project-stock-out-transaction-verification-template.html")
	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Reset password from vascom \n%s\n\n", mimeHeaders)))
	// Define the email data

	templateData := UserStockOutTransactionVerificationByEmailTemplateData{
		UserName:         user.FirstName,
		VerificationCode: verificationCode,
	}
	t.Execute(&body, templateData)
	message, err := utilService.SendEmail(user.Email, body)
	if err != nil {
		fmt.Println("There is error when sending email", err.Error())
		return "There is error when sending email", err
	}
	return message, nil

}
