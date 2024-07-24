package transaction_services

import (
	"context"
	graphqlClient "server/clients/graphql"
)

func UpdateUserReturnTransaction(transactionId string, isVerified bool, xHasuraRole string, stringToken string) error {

	// Define the GraphQL mutation
	var mutation struct {
		Update_transaction_user_return_transactions_by_pk struct {
			ID string `json:"id"`
		} `graphql:"update_transaction_user_return_transactions_by_pk(where: {id: {_eq: $id}}, _set: $set)"`
	}
	// Prepare variables for the mutation
	variables := map[string]interface{}{
		"id": transactionId,
		"set": map[string]interface{}{
			"Is_verified": isVerified,
		},
	}

	// Execute the GraphQL mutation
	err := graphqlClient.AuthClient(xHasuraRole, stringToken).Mutate(context.Background(), &mutation, variables)
	if err != nil {
		return err
	}
	return nil
}
