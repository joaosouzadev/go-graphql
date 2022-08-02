package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/google/uuid"
	"github.com/joaosouzadev/go-graphql/graph/generated"
	"github.com/joaosouzadev/go-graphql/graph/model"
)

// CreateCustomer is the resolver for the createCustomer field.
func (r *mutationResolver) CreateCustomer(ctx context.Context, input model.NewCustomer) (*model.Customer, error) {
	customer := &model.Customer{
		ID:    uuid.New().String(),
		Name:  input.Name,
		Email: input.Email,
	}
	r.customers = append(r.customers, customer)
	return customer, nil
}

// CreateItem is the resolver for the createItem field.
func (r *mutationResolver) CreateItem(ctx context.Context, input model.NewItem) (*model.Item, error) {
	item := &model.Item{
		ID:    uuid.New().String(),
		Name:  input.Name,
		Price: input.Price,
	}
	r.items = append(r.items, item)
	return item, nil
}

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input model.NewOrder) (*model.Order, error) {
	var customer *model.Customer
	for _, v := range r.customers {
		if v.ID == input.CustomerID {
			customer = v
			break
		}
	}

	var total float64
	for _, v := range r.items {
		for _, k := range input.Items {
			if v.ID == k.ItemID {
				total = total + (v.Price * k.Quantity)
			}
		}
	}

	orderId := uuid.New().String()
	order := &model.Order{
		ID:       orderId,
		Customer: customer,
		Total:    total,
	}
	r.orders = append(r.orders, order)

	var orderItems []*model.OrderItem
	for _, v := range input.Items {
		v.OrderID = &orderId
		orderIt, err := r.CreateOrderItem(ctx, *v)
		if err != nil {
			panic(err)
		}
		orderItems = append(orderItems, orderIt)
	}

	order.Items = orderItems
	return order, nil
}

// CreateOrderItem is the resolver for the createOrderItem field.
func (r *mutationResolver) CreateOrderItem(ctx context.Context, input model.NewOrderItem) (*model.OrderItem, error) {
	var product *model.Item
	for _, v := range r.items {
		if v.ID == input.ItemID {
			product = v
		}
	}

	var order *model.Order
	for _, v := range r.orders {
		if v.ID == *input.OrderID {
			order = v
		}
	}

	orderItem := &model.OrderItem{
		ID:       uuid.New().String(),
		Order:    order,
		Item:     product,
		Quantity: input.Quantity,
	}
	r.orderItems = append(r.orderItems, orderItem)
	return orderItem, nil
}

// Customers is the resolver for the customers field.
func (r *queryResolver) Customers(ctx context.Context) ([]*model.Customer, error) {
	return r.Resolver.customers, nil
}

// Items is the resolver for the items field.
func (r *queryResolver) Items(ctx context.Context) ([]*model.Item, error) {
	return r.Resolver.items, nil
}

// Orders is the resolver for the orders field.
func (r *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {
	return r.Resolver.orders, nil
}

// OrderItems is the resolver for the orderItems field.
func (r *queryResolver) OrderItems(ctx context.Context) ([]*model.OrderItem, error) {
	return r.Resolver.orderItems, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
