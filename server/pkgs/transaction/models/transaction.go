package transaction_models

type Transaction struct {
	ID                   string               `json:"id"`
	Transaction_products []TransactionProduct `json:"transaction_products"`
	Transaction_category string               `json:"transaction_category"`
}
