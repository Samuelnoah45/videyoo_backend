package request_models

type PurchaseRequest struct {
	ID                  string `json:"id"`
	Stock_manager_id    string `json:"stock_manager_id"`
	General_manager_id  string `json:"general_manager_id"`
	Purchase_manager_id string `json:"purchase_manager_id"`
}
