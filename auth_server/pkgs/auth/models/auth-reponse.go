package auth_models

type AuthResponse struct {
	ID                string `json:"id"`
	HasuraAccessToken string `json:"hasura_access_token"`
	User              User   `json:"user"`
}
