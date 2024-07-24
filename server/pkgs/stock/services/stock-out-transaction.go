package stock_services

import (
	"context"
	graphqlClient "server/clients/graphql"

	stockModel "server/pkgs/stock/models"
	transactionModel "server/pkgs/transaction/models"
)

func ProcessStockOutTransaction(transaction transactionModel.Transaction) error {
	// step 1: split transaction_products in two based on is_new_to_warehouse

	newToWarehouseProducts := make([]transactionModel.TransactionProduct, 0)
	existingInWarehouseProducts := make([]transactionModel.TransactionProduct, 0)
	for _, transactionProduct := range transaction.Transaction_products {
		if transactionProduct.Is_new_to_warehouse {
			newToWarehouseProducts = append(newToWarehouseProducts, transactionProduct)
		} else {
			existingInWarehouseProducts = append(existingInWarehouseProducts, transactionProduct)
		}
	}

	// step 2: insert new transaction_products into  warehouse
	// step 2.1 define insert mutation
	var insertWarehouseProductsMutation struct {
		Insert_warehouses_warehouse_products struct {
			Affected_rows int `json:"affected_rows"`
		} `graphql:"insert_warehouses_warehouse_products(objects: $objects)"`
	}

	// step 2.2. Define variables
	var insertObjects []stockModel.Warehouses_warehouse_products_insert_input
	for _, newToWarehouseProduct := range newToWarehouseProducts {
		insertObjects = append(insertObjects, stockModel.Warehouses_warehouse_products_insert_input{
			Product_id:   newToWarehouseProduct.Product_id,
			Warehouse_id: newToWarehouseProduct.Warehouse_id,
			Price:        newToWarehouseProduct.Price,
			Is_new:       newToWarehouseProduct.Is_new,
			Quantity:     newToWarehouseProduct.Quantity,
		})
	}

	insertVariable := map[string]interface{}{
		"objects": insertObjects,
	}

	// step 2.3. Execute the insert mutation
	err := graphqlClient.SystemClient().Mutate(context.Background(), &insertWarehouseProductsMutation, insertVariable)
	if err != nil {
		return err
	}

	// step 3: update existing transaction_products in warehouse
	// step 3.1 define update mutation
	var updateWarehouseProductsMutation struct {
		Update_many_warehouses_warehouse_products struct {
			Affected_rows int `json:"affected_rows"`
		} `graphql:"update_many_warehouses_warehouse_products(where: {}, _set: $set)"`
	}

	// step 3.2. Define update variables
	selectedWarehouseProducts, fetchError := GetWarehouseProducts(existingInWarehouseProducts)
	if fetchError != nil {
		return fetchError
	}
	var updateInputs []stockModel.Warehouses_warehouse_products_update_input
	for _, existingProduct := range existingInWarehouseProducts {
		// get warehouse product
		warehouseProduct, findError := FindWarehouseProduct(selectedWarehouseProducts, existingProduct)
		if findError != nil {
			return findError

		}
		updateInput := stockModel.Warehouses_warehouse_products_update_input{
			ID:       warehouseProduct.ID,
			Quantity: warehouseProduct.Quantity - existingProduct.Quantity, // Assuming you want to update the Quantity with the value from existingInWarehouseProducts
		}
		updateInputs = append(updateInputs, updateInput)
	}
	updateVariable := map[string]interface{}{
		"_set": updateInputs, //
	}
	// step 3.5 Execute mutation
	updateError := graphqlClient.SystemClient().Mutate(context.Background(), &updateWarehouseProductsMutation, updateVariable)
	if updateError != nil {
		return updateError
	}
	return nil

}
