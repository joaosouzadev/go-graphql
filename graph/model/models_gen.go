// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Customer struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Item struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type NewCustomer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type NewItem struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type NewOrder struct {
	CustomerID string          `json:"customerId"`
	Items      []*NewOrderItem `json:"items"`
}

type NewOrderItem struct {
	OrderID  *string `json:"orderId"`
	ItemID   string  `json:"itemId"`
	Quantity float64 `json:"quantity"`
}

type Order struct {
	ID       string       `json:"id"`
	Customer *Customer    `json:"customer"`
	Items    []*OrderItem `json:"items"`
	Total    float64      `json:"total"`
}

type OrderItem struct {
	ID       string  `json:"id"`
	Order    *Order  `json:"order"`
	Item     *Item   `json:"item"`
	Quantity float64 `json:"quantity"`
}