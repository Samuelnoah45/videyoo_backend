package auth_models

type JWTPayload struct {
	Subject      string `json:"sub"`
	IssuedAt     int64  `json:"iat"`
	ExpiresAt    int64  `json:"exp"`
	HasuraClaims struct {
		AllowedRoles []string `json:"x-hasura-allowed-roles"`
		DefaultRole  string   `json:"x-hasura-default-role"`
		UserID       int      `json:"x-hasura-user-id"`
	} `json:"https://hasura.io/jwt/claims"`
}
