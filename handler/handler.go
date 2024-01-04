package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Order struct{}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create an order")
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List all orders")
}

func (o *Order) GetById(w http.ResponseWriter, r *http.Request) {
	orderId := chi.URLParam(r, "id")

	fmt.Println("Get an order by ID", orderId)
}

func (o *Order) UpdateById(w http.ResponseWriter, r *http.Request) {
	orderId := chi.URLParam(r, "id")

	fmt.Println("Update an order by ID", orderId)
}

func (o *Order) DeleteById(w http.ResponseWriter, r *http.Request) {
	orderId := chi.URLParam(r, "id")

	fmt.Println("Delete an order by ID", orderId)
}
