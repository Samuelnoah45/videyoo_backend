package utilService

import (
	"fmt"
	"server/config"

	authModel "server/pkgs/auth/models"
	"time"

	"github.com/golang-jwt/jwt"
)

func createJWTToken(payload map[string]interface{}, secretKey string, tokenExpiration int64) (string, error) {
	// Create a new JWT token with the given payload and secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(payload))
	// Set the token expiration time to 1 hour from now
	token.Claims.(jwt.MapClaims)["exp"] = tokenExpiration
	// Sign the token with the secret key and return the signed token as a string
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func HasuraAccessToken(user authModel.User) (string, error) {
	payload := map[string]interface{}{
		"sub": "12345",                               // The user ID
		"iat": time.Now().Unix(),                     // The token issue time (UNIX timestamp)
		"exp": time.Now().Add(time.Hour * 48).Unix(), // The token expiration time (UNIX timestamp)
		"https://hasura.io/jwt/claims": map[string]interface{}{ // Hasura claims
			"x-hasura-allowed-roles": user.UserRoles, // The allowed roles for the user
			"x-hasura-default-role":  "org-member",   // The default role for the user
			"x-hasura-user-id":       user.ID,        // The user ID
		},
	}
	tokenExpiration := time.Now().Add(time.Hour * 48).Unix()
	token, err := createJWTToken(payload, config.JWT_SECRET_KEY, tokenExpiration)
	if err != nil {
		return "", err
	}
	return token, nil
}

func EmailVerificationToken(email string) (string, error) {
	payload := map[string]interface{}{
		"sub":   "12345",                                 // The user ID
		"iat":   time.Now().Unix(),                       // The token issue time (UNIX timestamp)
		"exp":   time.Now().Add(time.Minute * 10).Unix(), // The token expiration time (UNIX timestamp)
		"email": email,
	}
	tokenExpiration := time.Now().Add(time.Minute * 10).Unix()
	token, err := createJWTToken(payload, config.VERIFICATION_SECRET_KEY, tokenExpiration)
	if err != nil {
		return "", err
	}
	return token, nil
}

func ValidateVerificationToken(signedToken string) (map[string]interface{}, error) {
	// Parse the token
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.VERIFICATION_SECRET_KEY), nil
	})
	if err != nil {
		return nil, err
	}

	// Verify the token
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Extract and return the payload
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims format")
	}
	return claims, nil
}

func ValidateJWTToken(signedToken string) (map[string]interface{}, error) {
	// Parse the token
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_SECRET_KEY), nil
	})
	if err != nil {
		return nil, err
	}

	// Verify the token
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Extract and return the payload
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims format")
	}

	return claims, nil
}
