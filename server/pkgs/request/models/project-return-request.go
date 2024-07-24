package request_models

type ProjectReturnRequest struct {
	ID                 string `json:"id"`
	Project_manager_id string `json:"project_manager_id"`
	Stock_manager_id   string `json:"stock_manager_id"`
	Project_id         string `json:"project_id"`
}
