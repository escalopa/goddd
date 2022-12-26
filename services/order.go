package services

import (
	"github.com/escalopa/ddd-go/memory"
	repo "github.com/escalopa/ddd-go/repository"
	"github.com/google/uuid"
)

type OrderConfig func(os OrderService) error

type OrderService interface {
	PlaceOrder(customerID string, itemIDs []uuid.UUID) (float64, error)
	InjectCustomerRepository(cr repo.CustomerRepository)
	InjectProductRepository(pr repo.ProductRepository)
}

func WithCustomerRepository(cr repo.CustomerRepository) OrderConfig {
	return func(os OrderService) error {
		os.InjectCustomerRepository(cr)
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfig {
	cr := memory.NewMemoryCustomerRepository()
	return WithCustomerRepository(cr)
}

func WithProductRepository(pr repo.ProductRepository) OrderConfig {
	return func(os OrderService) error {
		os.InjectProductRepository(pr)
		return nil
	}
}

func WithMemoryProductRepository() OrderConfig {
	pr := memory.NewMemoryProductRepository()
	return WithProductRepository(pr)
}
