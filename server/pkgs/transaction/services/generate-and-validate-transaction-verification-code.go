package transaction_services

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func hashVerificationCode(verificationCode string) (string, error) {
	// Generate a salt for the password hash
	// The cost parameter is used to determine the complexity of the hash
	// Increasing the cost parameter makes the hash more computationally expensive to generate
	// and therefore harder to crack
	code, err := bcrypt.GenerateFromPassword([]byte(verificationCode), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(code), nil
}
func GenerateTransactionVerificationCode() (string, error) {
	rand.Seed(time.Now().UnixNano())
	code := fmt.Sprintf("%06d", rand.Intn(999999)) // Generate a 6-digit code
	verificationCode, err := hashVerificationCode(code)
	if err != nil {
		return "", err
	}
	return verificationCode, nil
}
func ValidatedVerificationCode(hashedVerificationCode string, verificationCode string) bool {
	// Convert the hashed password string to a byte slice
	hashedVerificationCodeBytes := []byte(hashedVerificationCode)
	// Compare the hashed password byte slice with the plaintext password byte slice
	err := bcrypt.CompareHashAndPassword(hashedVerificationCodeBytes, []byte(verificationCode))
	// If the error is nil, the passwords match
	return err == nil
}
