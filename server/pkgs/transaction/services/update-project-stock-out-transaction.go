package transaction_services

import (
	"context"
	graphqlClient "server/clients/graphql"
)

func UpdateProjectStockOutTransaction(transactionId string, isVerified bool) error {

	// Define the GraphQL mutation
	var mutation struct {
		Update_transaction_project_stock_out_transactions_by_pk struct {
			ID string `json:"id"`
		} `graphql:"update_transaction_user_stock_out_transactions_by_pk(where: {id: {_eq: $id}}, _set: $set)"`
	}

	// Prepare variables for the mutation
	variables := map[string]interface{}{
		"id": transactionId,
		"set": map[string]interface{}{
			"Is_verified":                   isVerified,
			"Transaction_verification_code": "",
		},
	}

	// Execute the GraphQL mutation
	err := graphqlClient.SystemClient().Mutate(context.Background(), &mutation, variables)
	if err != nil {
		return err
	}
	return nil
}
