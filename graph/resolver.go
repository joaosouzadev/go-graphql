package graph

import "github.com/joaosouzadev/go-graphql/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	customers  []*model.Customer
	items      []*model.Item
	orders     []*model.Order
	orderItems []*model.OrderItem
}
