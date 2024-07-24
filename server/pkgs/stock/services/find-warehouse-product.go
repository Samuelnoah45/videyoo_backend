package stock_services

import (
	"errors"

	stockModel "server/pkgs/stock/models"
	transactionModel "server/pkgs/transaction/models"
)

func FindWarehouseProduct(products []stockModel.WarehouseProduct, transactionProduct transactionModel.TransactionProduct) (*stockModel.WarehouseProduct, error) {
	for _, product := range products {
		if product.Product_id == transactionProduct.Product_id && product.Warehouse_id == transactionProduct.Warehouse_id && product.Is_new == transactionProduct.Is_new {
			return &product, nil // Found a match, return a pointer to the product
		}
	}
	return nil, errors.New("warehouse product not found") // No match found
}
