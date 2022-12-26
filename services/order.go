package services

import (
	repo "github.com/escalopa/ddd-go/domain/repository"
)

type OrderConfig func(os OrderService) error

type OrderService interface {
	PlaceOrder(customerID string, itemIDs []string) (float64, error)
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
	cr := repo.NewMemoryCustomerRepository()
	return WithCustomerRepository(cr)
}

func WithProductRepository(pr repo.ProductRepository) OrderConfig {
	return func(os OrderService) error {
		os.InjectProductRepository(pr)
		return nil
	}
}

func WithMemoryProductRepository() OrderConfig {
	pr := repo.NewMemoryProductRepository()
	return WithProductRepository(pr)
}
