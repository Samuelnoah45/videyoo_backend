package request_services

// imports
import (
	"context"
	graphqlClient "server/clients/graphql"
	requestModel "server/pkgs/request/models"
)

func GetProjectStockOutRequest(id string) (requestModel.ProjectStockOutRequest, error) {
	var query struct {
		Requests_project_stock_out_requests_by_pk struct {
			ID                 string `json:"id"`
			Project_manager_id string `json:"project_manager_id"`
			Stock_manager_id   string `json:"stock_manager_id"`
			Project_id         string `json:"project_id"`
			Request_id         string `json:"request_id"`
		} `graphql:"requests_project_stock_out_requests_by_pk"`
	}
	type uuid string
	variables := map[string]interface{}{
		"id": uuid(id),
	}
	err := graphqlClient.AnonymousClient().Query(context.Background(), &query, variables)
	if err != nil {
		return requestModel.ProjectStockOutRequest{}, err
	}
	request := requestModel.ProjectStockOutRequest{
		// write data from query object
		ID:                 query.Requests_project_stock_out_requests_by_pk.ID,
		Project_manager_id: query.Requests_project_stock_out_requests_by_pk.Project_manager_id,
		Stock_manager_id:   query.Requests_project_stock_out_requests_by_pk.Stock_manager_id,
		Project_id:         query.Requests_project_stock_out_requests_by_pk.Project_id,
	}
	return request, nil
}
