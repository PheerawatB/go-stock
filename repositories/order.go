package repositories

import (
	"sync"
)

type Order struct {
	ID           string   `json:"id"`
	CustomerName string   `json:"customerName"`
	Items        []string `json:"items"`
	Total        float64  `json:"total"`
	Status       string   `json:"status"`
}

type OrderRepository struct {
	data map[string]Order
	mu   sync.Mutex
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{data: make(map[string]Order)}
}

func (r *OrderRepository) Save(order Order) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[order.ID] = order
}

func (r *OrderRepository) FindByID(id string) (Order, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	order, exists := r.data[id]
	return order, exists
}

func (r *OrderRepository) FindAll() []Order {
	r.mu.Lock()
	defer r.mu.Unlock()
	orders := make([]Order, 0, len(r.data))
	for _, order := range r.data {
		orders = append(orders, order)
	}
	return orders
}

func (r *OrderRepository) DeleteByID(id string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.data, id)
}
