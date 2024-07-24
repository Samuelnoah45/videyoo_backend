package transaction_services

// imports
import (
	"context"
	"fmt"
	graphqlClient "server/clients/graphql"
	transactionModel "server/pkgs/transaction/models"
)

func GetTransaction(id string, xHasuraRole string, tokenString string) (transactionModel.Transaction, error) {
	var query struct {
		Transaction_transactions_by_pk struct {
			ID                   string                                `json:"id"`
			Transaction_category string                                `json:"transaction_category"`
			Transaction_products []transactionModel.TransactionProduct `json:"transaction_products"`
		} `graphql:"transaction_project_stock_out_transactions_by_pk"`
	}

	variables := map[string]interface{}{
		"id": id,
	}
	err := graphqlClient.AuthClient(xHasuraRole, tokenString).Query(context.Background(), &query, variables)
	if err != nil {
		fmt.Println(err.Error(), "when querying transaction")
		return transactionModel.Transaction{}, err
	}
	transaction := transactionModel.Transaction{
		// write data from query object
		ID:                   query.Transaction_transactions_by_pk.ID,
		Transaction_products: query.Transaction_transactions_by_pk.Transaction_products,
	}
	return transaction, nil
}
