package repository

import (
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

func (r *CustomerMemoryRepository) Find(id string) (*aggregate.Customer, error) {
	customer, ok := r.customers[id]
	if !ok {
		return nil, ErrorCustomerNotFound
	}
	return customer, nil
}

func (r *CustomerMemoryRepository) Save(customer *aggregate.Customer) error {
	r.Lock()
	defer r.Unlock()

	if err := customer.Validate(); err != nil {
		return err
	}

	r.customers[customer.GetID()] = customer
	return nil
}

func (r *CustomerMemoryRepository) Update(customer *aggregate.Customer) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.customers[customer.GetID()]; !ok {
		return ErrorCustomerNotFound
	}

	r.customers[customer.GetID()] = customer
	return nil
}
