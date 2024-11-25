package service

import (
	"test/model"
	"time"
)

func ProcessTransaction(req model.ProductModelRequest) model.ProductModelResponse {
	transactionID := "TR001"
	transactionTotal := 0
	categoryMap := make(map[string][]model.ItemModel)

	for _, product := range req.Product {
		transactionTotal += int(product.Price)
		items := model.ItemModel{
			Id:    product.Id,
			Name:  product.Name,
			Price: product.Price,
		}
		categoryMap[product.Category] = append(categoryMap[product.Category], items)
	}

	var transactionDetails []model.TransactionDetailModel
	for category, items := range categoryMap {
		transactionDetails = append(transactionDetails, model.TransactionDetailModel{
			Category: category,
			Items:    items,
		})
	}

	currentTime := time.Now().Format("02 Jan 2006 15:04:05")

	// build response
	response := model.ProductModelResponse{
		Id:                transactionID,
		TransactionTotal:  int32(transactionTotal),
		TransactionDate:   currentTime,
		TransactionDetail: transactionDetails,
	}

	return response
}
