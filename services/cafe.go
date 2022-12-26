package services

import "github.com/google/uuid"

type CafeConfig func(cs CafeService) error

type CafeService interface {
	CreateOrder(customerID string, itemIDs []uuid.UUID) (float64, error)
	InjectOrderService(os OrderService)
}

func WithOrderService(os OrderService) CafeConfig {
	return func(cs CafeService) error {
		cs.InjectOrderService(os)
		return nil
	}
}
