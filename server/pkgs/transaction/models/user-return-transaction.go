package transaction_models

type UserReturnTransaction struct {
	ID                     string `json:"id"`
	User_id                string `json:"project_manager_id"`
	Stock_manager_id       string `json:"stock_manager_id"`
	User_return_request_id string `json:"user_stock_out_request_id"`
	Transaction_id         string `json:"transaction_id"`
	Is_verified            bool   `json:"is_verified"`
}
