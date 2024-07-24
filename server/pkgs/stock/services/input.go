package stock_services

type Warehouses_warehouse_products_insert_input struct {
	Product_id   string  `json:"product_id"`
	Warehouse_id string  `json:"warehouse_id"`
	Is_new       bool    `json:"is_new"`
	Price        float64 `json:"Price"`
	Quantity     float64 `json:"Quantity"`
}

type Warehouses_warehouse_products_update_input struct {
	ID       string  `json:"id"`
	Price    float64 `json:"Price"`
	Quantity float64 `json:"Quantity"`
}
