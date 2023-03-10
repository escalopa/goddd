package services

import (
	"context"
	"log"

	"github.com/escalopa/ddd-go/aggregate"
	repo "github.com/escalopa/ddd-go/repository"
	"github.com/google/uuid"
)

type OrderServiceImpl struct {
	cp repo.CustomerRepository
	rp repo.ProductRepository
}

func NewOrderServiceImpl(cfgs ...OrderConfig) (OrderService, error) {
	s := &OrderServiceImpl{}
	for _, cfg := range cfgs {
		err := cfg(s)
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

func (s *OrderServiceImpl) PlaceOrder(customerID string, itemIDs []uuid.UUID) (float64, error) {
	customer, err := s.cp.Find(context.Background(), customerID)
	if err != nil {
		return 0.0, err
	}

	var total float64
	var products []*aggregate.Product

	for _, itemID := range itemIDs {
		product, err := s.rp.Find(context.Background(), itemID.String())
		if err != nil {
			return 0.0, err
		}
		total += product.GetPrice()
		products = append(products, product)
	}

	log.Printf("Customer: %s has ordered %d products for a total of %f", customer.GetName(), len(products), total)
	return total, nil
}

func (s *OrderServiceImpl) InjectCustomerRepository(cr repo.CustomerRepository) {
	s.cp = cr
}

func (s *OrderServiceImpl) InjectProductRepository(pr repo.ProductRepository) {
	s.rp = pr
}
