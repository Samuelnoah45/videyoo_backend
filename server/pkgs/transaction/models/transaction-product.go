package transaction_models

import productModel "server/pkgs/product/models"

type TransactionProduct struct {
	ID                  string               `json:"id"`
	Product_id          string               `json:"product_id"`
	Product             productModel.Product `json:"product"`
	Quantity            float64              `json:"quantity"`
	Unit_price          float64              `json:"unit_price"`
	Is_new              bool                 `json:"is_new"`
	Is_returnable       bool                 `json:"is_returnable"`
	Is_new_to_warehouse bool                 `json:"is_new_to_warehouse"`
	Warehouse_id        string               `json:"warehouse_id"`
	Price               float64              `json:"Price"`
}
