package transaction_services

// package transaction_controllers

// import (
// 	"context"
// 	graphqlClient "server/clients/graphql"
// 	warehouseHandler "server/handlers/warehouses"
// 	"server/models/transactions"
// 	warehouseModel "server/models/warehouses"
// )

// func ProcessStockInTransaction(transaction transactions.Transaction) error {

// 	// step 3: update existing transaction_products in warehouse
// 	// step 3.1 define update mutation
// 	var updateWarehouseProductsMutation struct {
// 		Update_many_werehouses_werehouse_products struct {
// 			Affected_rows int `json:"affected_rows"`
// 		} `graphql:"update_many_werehouses_werehouse_products(where: {}, _set: $set)"`
// 	}
// 	// step 3.2. Define update variables
// 	// /3.2.1. Fetch warehouse product information
// 	selectedWarehouseProducts, fetchError := warehouseHandler.GetWarehouseProducts(transaction.Transaction_products)
// 	if fetchError != nil {
// 		return fetchError
// 	}

// 	// step 3.4. Define update variables
// 	var updateInputs []warehouseModel.Werehouses_werehouse_products_update_input
// 	for _, existingProduct := range transaction.Transaction_products {
// 		// get warehouse product
// 		warehouseProduct, findError := warehouseHandler.FindWarehouseProduct(selectedWarehouseProducts, existingProduct)
// 		if findError != nil {
// 			return findError
// 		}
// 		updateInput := warehouseModel.Werehouses_werehouse_products_update_input{
// 			ID:       warehouseProduct.ID,
// 			Quantity: warehouseProduct.Quantity + existingProduct.Quantity, // Assuming you want to update the Quantity with the value from existingInWarehouseProducts
// 		}
// 		updateInputs = append(updateInputs, updateInput)
// 	}
// 	updateVariable := map[string]interface{}{
// 		"_set": updateInputs, //
// 	}
// 	// step 3.5 Execute mutation
// 	updateError := graphqlClient.SystemClient().Mutate(context.Background(), &updateWarehouseProductsMutation, updateVariable)
// 	if updateError != nil {
// 		return updateError
// 	}
// 	return nil

// }
