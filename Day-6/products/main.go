package main

import (
	"net/http"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

var products = []Product{}

func GetProducts(w http.ResponseWriter, r *http.Request) {
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
}
