package services

import (
	"fmt"
	"gostock/repositories"
	"time"
)

type OrderService struct {
	repo *repositories.OrderRepository
}

func NewOrderService(repo *repositories.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) ProcessOrder(order repositories.Order) {
	order.Status = "pending"
	s.repo.Save(order)

	go func(o repositories.Order) {
		time.Sleep(2 * time.Second)
		o.Status = "processed"
		s.repo.Save(o)
		fmt.Printf("Order %s has been processed\n", o.ID)
	}(order)
}
