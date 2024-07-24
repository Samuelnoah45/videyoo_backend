package auth_services

// imports
import (
	"context"

	graphqlClient "server/clients/graphql"
	authModel "server/pkgs/auth/models"
)

func GetUser(id string) (authModel.User, error) {
	var query struct {
		User_users_by_pk struct {
			ID                  string `json:"id"`
			First_name          string `json:"first_name"`
			Last_name           string `json:"last_name"`
			Email               string `json:"email"`
			Password            string `json:"password"`
			Phone_number        string `json:"phone_number"`
			Is_account_verified bool   `json:"is_account_verified"`
			User_roles          []struct {
				Role_name string `json:"role_name"`
			} `json:"user_roles"`
		} `graphql:"user_users_by_pk(id: $id)"`
	}

	type uuid string
	variables := map[string]interface{}{
		"id": uuid(id),
	}

	err := graphqlClient.SystemClient().Query(context.Background(), &query, variables)
	if err != nil {
		return authModel.User{}, err

	}
	user := authModel.User{
		ID:        query.User_users_by_pk.ID,
		FirstName: query.User_users_by_pk.First_name,
		LastName:  query.User_users_by_pk.Last_name,
		Email:     query.User_users_by_pk.Email,
		UserRoles: []string{},
	}
	return user, nil
}
