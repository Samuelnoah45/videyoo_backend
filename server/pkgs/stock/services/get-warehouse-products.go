package stock_services

// // imports
import (
	"context"
	graphqlClient "server/clients/graphql"

	stockModel "server/pkgs/stock/models"
	transactionModel "server/pkgs/transaction/models"
)

func GetWarehouseProducts(transactionsProducts []transactionModel.TransactionProduct) ([]stockModel.WarehouseProduct, error) {
	var getWareHouseProducts struct {
		Warehouse_products []stockModel.WarehouseProduct `json:"warehouses_warehouse_products"`
	}
	// Create a map to store unique warehouse IDs efficiently
	distinctExistingTransactionProducts := make(map[string]struct{})
	// Iterate through existingInWarehouseProducts and add unique warehouse IDs to the map
	for _, wHProduct := range transactionsProducts {
		distinctExistingTransactionProducts[wHProduct.Product_id] = struct{}{} // Use an empty struct to create unique map keys
	}
	// Convert the map keys back to an array to maintain order
	productIDs := make([]string, 0, len(distinctExistingTransactionProducts))
	for productID := range distinctExistingTransactionProducts {
		productIDs = append(productIDs, productID)
	}
	fetchVariable := map[string]interface{}{
		"product_id": map[string]interface{}{
			"_in": productIDs,
		},
	}
	fetchError := graphqlClient.SystemClient().Query(context.Background(), &getWareHouseProducts, fetchVariable)
	if fetchError != nil {
		return nil, fetchError
	}
	return getWareHouseProducts.Warehouse_products, nil

}
