package customer

import (
	"sync"

	"github.com/escalopa/ddd-go/aggregate"
)

type MemoryRepository struct {
	customers map[string]*aggregate.Customer
	sync.Mutex
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		customers: map[string]*aggregate.Customer{},
	}
}

func (r *MemoryRepository) Find(id string) (*aggregate.Customer, error) {
	customer, ok := r.customers[id]
	if !ok {
		return nil, ErrorCustomerNotFound
	}
	return customer, nil
}

func (r *MemoryRepository) Save(customer *aggregate.Customer) error {
	r.Lock()
	r.customers[customer.GetID()] = customer
	r.Unlock()
	return nil
}

func (r *MemoryRepository) Update(customer *aggregate.Customer) error {
	if _, ok := r.customers[customer.GetID()]; !ok {
		return ErrorCustomerNotFound
	}

	r.Lock()
	r.customers[customer.GetID()] = customer
	r.Unlock()
	return nil
}
