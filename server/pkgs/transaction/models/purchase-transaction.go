package transaction_models

type PurchaseTransaction struct {
	ID                  string `json:"id"`
	Stock_manager_id    string `json:"stock_manager_id"`
	Purchase_manager_id string `json:"purchase_manager_id"`
	Purchase_request_id string `json:"user_stock_out_request_id"`
	Transaction_id      string `json:"transaction_id"`
	Is_verified         bool   `json:"is_verified"`
}
