package auth_models

type AuthResponse struct {
	Token string `json:"hasura_access_token"`
}
