package main

import (
	"log"
	"net/http"
	"test/controller"
)

func main() {
	http.HandleFunc("/transactions", controller.TransactionHandler)
	log.Println("HTTP Listen port :8080")
	http.ListenAndServe(":8080", nil)
}
