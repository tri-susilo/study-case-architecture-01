package main

import (
	"go_crud_1/config"
	category_controller "go_crud_1/controllers/category"
	home_controller "go_crud_1/controllers/home"
	product_controller "go_crud_1/controllers/product"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	//1. Homepage
	http.HandleFunc("/", home_controller.Welcome)

	//2. Categories
	http.HandleFunc("/categories", category_controller.Index)
	http.HandleFunc("/categories/add", category_controller.Add)
	http.HandleFunc("/categories/edit", category_controller.Edit)
	http.HandleFunc("/categories/delete", category_controller.Delete)

	//3. Products
	http.HandleFunc("/products", product_controller.Index)
	http.HandleFunc("/products/add", product_controller.Add)
	http.HandleFunc("/products/detail", product_controller.Detail)
	http.HandleFunc("/products/edit", product_controller.Edit)
	http.HandleFunc("/products/delete", product_controller.Delete)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
