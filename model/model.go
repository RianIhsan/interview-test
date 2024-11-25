package model

type ProductModelRequest struct {
	Product []ProductItem `json:"product"`
}

type ProductItem struct {
	Id       int32  `json:"id"`
	Name     string `json:"name"`
	Price    int32  `json:"price"`
	Category string `json:"category"`
}

type ProductModelResponse struct {
	Id                string                   `json:"id"`
	TransactionTotal  int32                    `json:"transaction_total"`
	TransactionDate   string                   `json:"transaction_date"`
	TransactionDetail []TransactionDetailModel `json:"transaction_detail"`
}

type TransactionDetailModel struct {
	Category string      `json:"category"`
	Items    []ItemModel `json:"items"`
}

type ItemModel struct {
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Price int32  `json:"price"`
}
