package data

import (
	"time"

	"ecommerce-graphql-go/graph/model"
)

type Data struct {
	products []*model.Product
	orders   []*model.Order
	users    []*model.User
}

func NewData() *Data {
	desc1 := "High-performance laptop"
	desc2 := "Latest model smartphone"
	desc3 := "Noise-cancelling headphones"

	admin := &model.User{
		ID:    "1",
		Name:  "Admin",
		Email: "admin@example.com",
		Role:  model.RoleAdmin,
	}

	customer := &model.User{
		ID:    "2",
		Name:  "Customer",
		Email: "customer@example.com",
		Role:  model.RoleCustomer,
	}

	products := []*model.Product{
		{
			ID:          "1",
			Name:        "Laptop",
			Description: &desc1,
			Price:       999.99,
			Category:    "Electronics",
			InStock:     10,
		},
		{
			ID:          "2",
			Name:        "Smartphone",
			Description: &desc2,
			Price:       699.99,
			Category:    "Electronics",
			InStock:     15,
		},
		{
			ID:          "3",
			Name:        "Headphones",
			Description: &desc3,
			Price:       199.99,
			Category:    "Audio",
			InStock:     20,
		},
	}

	orders := []*model.Order{
		{
			ID:        "1",
			Products:  []*model.Product{products[0], products[2]},
			Total:     1199.98,
			CreatedAt: time.Now().Add(-24 * time.Hour).Format(time.RFC3339),
			Status:    "DELIVERED",
			User:      customer,
		},
		{
			ID:        "2",
			Products:  []*model.Product{products[1]},
			Total:     699.99,
			CreatedAt: time.Now().Add(-12 * time.Hour).Format(time.RFC3339),
			Status:    "PROCESSING",
			User:      customer,
		},
	}

	return &Data{
		products: products,
		orders:   orders,
		users:    []*model.User{admin, customer},
	}
}

func (s *Data) GetProduct(id string) *model.Product {
	for _, p := range s.products {
		if p.ID == id {
			return p
		}
	}
	return nil
}

func (s *Data) GetProducts(limit, offset int32, category string) []*model.Product {
	var products []*model.Product

	// Filter category
	for _, p := range s.products {
		if category == "" || p.Category == category {
			products = append(products, p)
		}
	}

	// Pagination
	start := offset
	end := offset + limit
	length := int32(len(products))
	if start >= length {
		return []*model.Product{}
	}
	if end > length {
		end = length
	}

	return products[start:end]
}

func (s *Data) CreateProduct(p *model.Product) {
	s.products = append(s.products, p)
}

func (s *Data) UpdateProduct(p *model.Product) {
	for i, product := range s.products {
		if product.ID == p.ID {
			s.products[i] = p
			return
		}
	}
}

func (s *Data) GetOrder(id string, userID string) *model.Order {
	for _, o := range s.orders {
		if o.ID == id && o.User.ID == userID {
			return o
		}
	}
	return nil
}

func (s *Data) GetOrders(userID string) []*model.Order {
	var userOrders []*model.Order
	for _, o := range s.orders {
		if o.User.ID == userID {
			userOrders = append(userOrders, o)
		}
	}
	return userOrders
}

func (s *Data) CreateOrder(o *model.Order) {
	o.CreatedAt = time.Now().Format(time.RFC3339)
	s.orders = append(s.orders, o)
}

func (s *Data) GetUser(id string) *model.User {
	for _, u := range s.users {
		if u.ID == id {
			return u
		}
	}
	return nil
}
