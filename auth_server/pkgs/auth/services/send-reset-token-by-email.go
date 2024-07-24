package auth_services

// imports
import (
	"bytes"
	"context"
	"fmt"
	graphqlClient "server/clients/graphql"
	"time"

	"html/template"
	"server/utilService"
)

type EmailDataToken struct {
	Link   string
	Header string
}

func SendResetTokenByEmail(email string, resetUrl string) (string, error) {

	//2.  Define the GraphQL query to execute
	var query struct {
		User []struct {
			Email string `json:"email"`
		} `graphql:"user_users(where: {email: {_eq: $email}})"`
	}
	//3. construct graphql variables
	queryVariables := map[string]interface{}{
		"email": email,
	}
	//4 execute the request
	fetchError := graphqlClient.SystemClient().Query(context.Background(), &query, queryVariables)
	if fetchError != nil {
		fmt.Println(fetchError.Error(), "when querying user to forgot password")
		return "Something went  wrong when querying user", nil

	}

	//5.  Check if the user exists
	if len(query.User) == 0 {
		message := fmt.Sprintf("There is no User with email  %s ", email)

		return message, nil
	}
	//6.  Define the GraphQL mutation string that store password reset token in database
	var mutation struct {
		Update_user_users struct {
			Returning []struct {
				ID string `json:"id"`
			} `json:"returning"`
		} `graphql:"update_user_users(where: {email: {_eq: $email}}, _set: {reset_password_by_email_token: $resetPasswordByEmailToken, reset_password_by_email_token_expires_at: $tokenExpiresAt})"`
	}
	//7. create password reset token
	resetPasswordByEmailToken, generateEmailTokenError := utilService.EmailVerificationToken(email)
	if generateEmailTokenError != nil {
		return "Something went  wrong when generating email token", generateEmailTokenError
	}

	//8. construct graphql variables
	// Calculate the expiration time in UTC
	expiresAt := time.Now().UTC().Add(10 * time.Minute)
	type timestamptz string

	// Convert to timestamptz string representation
	timestamptzString := expiresAt.Format(time.RFC3339)
	mutationVariables := map[string]interface{}{
		"resetPasswordByEmailToken": resetPasswordByEmailToken,
		"email":                     email,
		"tokenExpiresAt":            timestamptz(timestamptzString),
	}
	//9. execute the set password Reset token mutation
	mutateError := graphqlClient.SystemClient().Mutate(context.Background(), &mutation, mutationVariables)
	if mutateError != nil {

		return "Something went  wrong when updating user", mutateError
	}
	//10. Send password reset token to user by  email

	// Create the reset URL with the token
	var resetURL string

	if resetUrl == "" {
		resetURL = "http://localhost:3000/resetPasswordByEmail/" + resetPasswordByEmailToken
	} else {
		resetURL = resetUrl + "/" + resetPasswordByEmailToken
	}

	t, _ := template.ParseFiles("./templates/verify-email-template.html")
	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Reset password from vascom \n%s\n\n", mimeHeaders)))
	// Define the email data

	emailData := EmailDataToken{
		Link:   resetURL,
		Header: "Reset your password with above link",
	}
	t.Execute(&body, emailData)
	message, sendEmailError := utilService.SendEmail(email, body)
	if sendEmailError != nil {
		fmt.Println("Something wrong or Not correct email", sendEmailError.Error())
		return "Something wrong or Incorrect email", sendEmailError
	}
	return message, sendEmailError

}
