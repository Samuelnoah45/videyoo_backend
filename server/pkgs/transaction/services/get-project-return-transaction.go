package transaction_services

// imports
import (
	"context"
	"fmt"
	graphqlClient "server/clients/graphql"

	transactionModel "server/pkgs/transaction/models"
)

func GetProjectReturnTransaction(id string) (transactionModel.ProjectReturnTransaction, error) {
	var query struct {
		Transaction_project_return_transactions_by_pk struct {
			ID                        string `json:"id"`
			Is_verified               bool   `json:"is_verified"`
			Stock_manager_id          string `json:"stock_manager_id"`
			Project_manager_id        string `json:"project_manager_id"`
			Project_id                string `json:"project_id"`
			Project_return_request_id string `json:"user_return_request_id"`
			Transaction_id            string `json:"transaction_id"`
		} `graphql:"transaction_user_return_transactions_by_pk"`
	}

	type uuid string
	variables := map[string]interface{}{
		"id": uuid(id),
	}

	err := graphqlClient.SystemClient().Query(context.Background(), &query, variables)
	if err != nil {
		fmt.Println(err.Error(), "when querying user stock out transaction")
		return transactionModel.ProjectReturnTransaction{}, err
	}
	transaction := transactionModel.ProjectReturnTransaction{
		// write data from query object
		ID:                           query.Transaction_project_return_transactions_by_pk.ID,
		Stock_manager_id:             query.Transaction_project_return_transactions_by_pk.Stock_manager_id,
		Project_id:                   query.Transaction_project_return_transactions_by_pk.Project_id,
		Project_stock_out_request_id: query.Transaction_project_return_transactions_by_pk.Project_return_request_id,
		Project_manager_id:           query.Transaction_project_return_transactions_by_pk.Project_manager_id,
	}
	return transaction, nil
}
