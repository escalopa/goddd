package services

import (
	"github.com/google/uuid"
	"log"
)

type CafeImpl struct {
	os OrderService
}

func NewCafeImpl(cfgs ...CafeConfig) (CafeService, error) {
	s := &CafeImpl{}
	for _, cfg := range cfgs {
		err := cfg(s)
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

func (s *CafeImpl) CreateOrder(customerID string, itemIDs []uuid.UUID) (float64, error) {
	total, err := s.os.PlaceOrder(customerID, itemIDs)
	if err != nil {
		return 0.0, err
	}
	log.Printf("Order placed for customer %s for a total of %f", customerID, total)

	// TODO: Bill the customer
	return total, nil
}

func (s *CafeImpl) InjectOrderService(os OrderService) {
	s.os = os
}
