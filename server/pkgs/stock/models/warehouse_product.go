package stock_models

type WarehouseProduct struct {
	ID           string  `json:"id"`
	Product_id   string  `json:"product_id"`
	Quantity     float64 `json:"quantity"`
	Unit_price   float64 `json:"unit_price"`
	Is_new       bool    `json:"is_new"`
	Warehouse_id string  `json:"warehouse_id"`
	Price        float64 `json:"Price"`
}
