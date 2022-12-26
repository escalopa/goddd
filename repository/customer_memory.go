package repository

import (
	"context"
	"sync"

	"github.com/escalopa/ddd-go/aggregate"
)

type CustomerMemoryRepository struct {
	customers map[string]*aggregate.Customer
	sync.Mutex
}

func NewMemoryCustomerRepository() *CustomerMemoryRepository {
	return &CustomerMemoryRepository{
		customers: map[string]*aggregate.Customer{},
	}
}

func (r *CustomerMemoryRepository) Find(_ context.Context, id string) (*aggregate.Customer, error) {
	customer, ok := r.customers[id]
	if !ok {
		return nil, ErrorCustomerNotFound
	}
	return customer, nil
}

func (r *CustomerMemoryRepository) Save(_ context.Context, customer *aggregate.Customer) error {
	r.Lock()
	defer r.Unlock()

	if err := customer.Validate(); err != nil {
		return err
	}

	r.customers[customer.GetID()] = customer
	return nil
}

func (r *CustomerMemoryRepository) Update(_ context.Context, customer *aggregate.Customer) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.customers[customer.GetID()]; !ok {
		return ErrorCustomerNotFound
	}

	r.customers[customer.GetID()] = customer
	return nil
}
