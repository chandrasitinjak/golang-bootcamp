package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	_ "sesi_8/swaggo/docs"
	"strconv"
	"time"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /

//Order represents the model for an order
type Order struct {
	OrderID      string    `json:"orderId" example:"1"`
	CustomerName string    `json:"customerName" example:"Leo Messi"`
	OrderedAt    time.Time `json:"ordereAt" example:"2019-11-09T21:21:46+00:00"`
	Items        []Item    `json:"items"`
}

var (
	orders []Order
)

//item represents the model for an item in the order
type Item struct {
	ItemID      string `json:"itemId", example:"A1B2C3"`
	Description string `json:"description", example:"A random description"`
	Quantity    int    `json:"quantity", example:"1"`
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/orders", createOrder).Methods("POST")
	//router.HandleFunc("/order", getOrders).Methods("GET")

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	log.Println(http.ListenAndServe(":8080", router))

}

// ShowAccount godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  Order
// @Router       /orders [post]
func createOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order Order

	json.NewDecoder(r.Body).Decode(&order)

	prevOrderID := 1

	order.OrderID = strconv.Itoa(prevOrderID)

	orders = append(orders, order)
	json.NewEncoder(w).Encode(order)

}
