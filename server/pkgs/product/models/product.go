package product_models

type Product struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Model            string `json:"model"`
	Measurement_unit string `json:"measurement_unit"`
	Sku              string `json:"sku"`
}
