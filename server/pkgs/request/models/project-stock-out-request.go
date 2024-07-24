package request_models

type ProjectStockOutRequest struct {
	ID                           string `json:"id"`
	Project_manager_id           string `json:"project_manager_id"`
	Stock_manager_id             string `json:"stock_manager_id"`
	Technical_project_manager_id string `json:"technical_project_manager_id"`
	General_manager_id           string `json:"general_manager_id"`
	Project_id                   string `json:"project_id"`
}
