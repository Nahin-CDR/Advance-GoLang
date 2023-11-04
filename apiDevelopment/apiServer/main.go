package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	Id       string
	Name     string
	Quantity int
	Price    float64
}

var Products []Product

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welome to home page")
	fmt.Println("Endpoint hit : homepage")
}

func returnAllProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit : return All Products")

	json.NewEncoder(w).Encode(Products)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)

	key := r.URL.Path[len("/products/"):]

	for _, product := range Products {
		if string(product.Id) == key {
			json.NewEncoder(w).Encode(product)
		}
	}
}

func handleRequest() {
	http.HandleFunc("/products/", getProduct)
	http.HandleFunc("/products", returnAllProducts)
	http.HandleFunc("/", homePage)
	http.ListenAndServe("localhost:10000", nil)
}

func main() {

	Products = []Product{
		{
			Id:       "1",
			Name:     "Chair",
			Quantity: 100,
			Price:    200.00,
		},
		{
			Id:       "2",
			Name:     "Desk",
			Quantity: 300,
			Price:    450.0,
		},
	}
	handleRequest()

}
