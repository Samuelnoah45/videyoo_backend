package transaction_services

// imports
import (
	"context"
	"fmt"
	graphqlClient "server/clients/graphql"

	transactionModel "server/pkgs/transaction/models"
)

func GetProjectStockOutTransaction(id string) (transactionModel.ProjectStockOutTransaction, error) {
	var query struct {
		Transaction_project_stock_out_transactions_by_pk struct {
			ID                            string `json:"id"`
			Is_verified                   bool   `json:"is_verified"`
			Project_manager_id            string `json:"project_manager_id"`
			Stock_manager_id              string `json:"stock_manager_id"`
			Project_id                    string `json:"project_id"`
			Project_stock_out_request_id  string `json:"project_stock_out_request_id"`
			Transaction_id                string `json:"transaction_id"`
			Transaction_verification_code string `json:"transaction_verification_code"`
		} `graphql:"transaction_project_stock_out_transactions_by_pk"`
	}
	type uuid string
	variables := map[string]interface{}{
		"id": uuid(id),
	}
	err := graphqlClient.AnonymousClient().Query(context.Background(), &query, variables)
	if err != nil {
		fmt.Println(err.Error(), "when querying project")
		return transactionModel.ProjectStockOutTransaction{}, err
	}
	transaction := transactionModel.ProjectStockOutTransaction{
		// write data from query object
		ID:                            query.Transaction_project_stock_out_transactions_by_pk.ID,
		Is_verified:                   query.Transaction_project_stock_out_transactions_by_pk.Is_verified,
		Project_manager_id:            query.Transaction_project_stock_out_transactions_by_pk.Project_manager_id,
		Stock_manager_id:              query.Transaction_project_stock_out_transactions_by_pk.Stock_manager_id,
		Project_id:                    query.Transaction_project_stock_out_transactions_by_pk.Project_id,
		Project_stock_out_request_id:  query.Transaction_project_stock_out_transactions_by_pk.Project_stock_out_request_id,
		Transaction_id:                query.Transaction_project_stock_out_transactions_by_pk.Transaction_id,
		Transaction_verification_code: query.Transaction_project_stock_out_transactions_by_pk.Transaction_verification_code,
	}
	return transaction, nil
}
