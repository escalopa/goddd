package repository

import (
	"errors"

	"github.com/escalopa/ddd-go/aggregate"
)

var (
	ErrorCustomerNotFound     = errors.New("customer not found")
	ErrorCustomerNameOldEqNew = errors.New("customer name old eq new")

	ErrorProductNotFound = errors.New("product not found")
)

type CustomerRepository interface {
	Find(id string) (*aggregate.Customer, error)
	Save(customer *aggregate.Customer) error
	Update(customer *aggregate.Customer) error
}

type ProductRepository interface {
	Find(id string) (*aggregate.Product, error)
	Save(product *aggregate.Product) error
	Update(product *aggregate.Product) error
	FindAll() ([]*aggregate.Product, error)
	Delete(id string) error
}
