package customer

import (
	"errors"

	"github.com/escalopa/ddd-go/aggregate"
)

var (
	ErrorCustomerNotFound = errors.New("customer not found")
	ErrorSaveCustomer     = errors.New("error saving customer")
	ErrorUpdateCustomer   = errors.New("error updating customer")
)

type CustomerRepository interface {
	Find(id string) (*aggregate.Customer, error)
	Save(customer *aggregate.Customer) error
	Update(customer *aggregate.Customer) error
}
