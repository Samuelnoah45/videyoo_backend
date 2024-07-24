package transaction_models

type ProjectStockOutTransaction struct {
	ID                            string `json:"id"`
	Project_manager_id            string `json:"project_manager_id"`
	Stock_manager_id              string `json:"stock_manager_id"`
	Project_id                    string `json:"project_id"`
	Is_verified                   bool   `json:"is_verified"`
	Project_stock_out_request_id  string `json:"project_stock_out_request_id"`
	Transaction_verification_code string `json:"transaction_verification_code"`
	Transaction_id                string `json:"transaction_id"`
}
