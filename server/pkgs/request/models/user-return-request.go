package request_models

type UserReturnRequest struct {
	ID               string `json:"id"`
	User_id          string `json:"user_id"`
	Stock_manager_id string `json:"stock_manager_id"`
}
