package transaction_services

// imports
import (
	"context"
	"fmt"
	graphqlClient "server/clients/graphql"
	transactionModel "server/pkgs/transaction/models"
)

func GetPurchaseTransaction(id string) (transactionModel.PurchaseTransaction, error) {
	var query struct {
		Transaction_purchase_transactions_by_pk struct {
			ID                  string `json:"id"`
			Is_verified         bool   `json:"is_verified"`
			Purchase_request_id string `json:"user_return_request_id"`
			Transaction_id      string `json:"transaction_id"`
			Stock_manager_id    string `json:"stock_manager_id"`
			Purchase_manager_id string `json:"purchase_manager_id"`
		} `graphql:"transaction_user_return_transactions_by_pk"`
	}

	type uuid string
	variables := map[string]interface{}{
		"id": uuid(id),
	}
	err := graphqlClient.SystemClient().Query(context.Background(), &query, variables)
	if err != nil {
		fmt.Println(err.Error(), "when querying user stock out transaction")
		return transactionModel.PurchaseTransaction{}, err
	}
	transaction := transactionModel.PurchaseTransaction{
		// write data from query object
		ID:                  query.Transaction_purchase_transactions_by_pk.ID,
		Stock_manager_id:    query.Transaction_purchase_transactions_by_pk.Stock_manager_id,
		Purchase_manager_id: query.Transaction_purchase_transactions_by_pk.Purchase_manager_id,
		Purchase_request_id: query.Transaction_purchase_transactions_by_pk.Purchase_request_id,
		Transaction_id:      query.Transaction_purchase_transactions_by_pk.Transaction_id,
		Is_verified:         query.Transaction_purchase_transactions_by_pk.Is_verified,
	}
	return transaction, nil
}
