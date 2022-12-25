package aggregate

import (
	"errors"

	"github.com/escalopa/ddd-go/entity"
	"github.com/escalopa/ddd-go/valueobject"
	"github.com/google/uuid"
)

var (
	ErrorEmptyName = errors.New("name is required")
)

type Customer struct {
	person       *entity.Person
	items        []*entity.Item
	transactions []*valueobject.Transaction
}

func (c *Customer) GetID() string {
	return c.person.ID.String()
}

func (c *Customer) GetName() string {
	return c.person.Name
}

func (c *Customer) SetName(name string) error {
	if name == "" {
		return ErrorEmptyName
	}

	c.person.Name = name
	return nil
}

func NewCustomer(name string) (*Customer, error) {
	if name == "" {
		return nil, ErrorEmptyName
	}

	customer := &Customer{
		person: &entity.Person{
			Name: name,
			ID:   uuid.New(),
		},
		items:        []*entity.Item{},
		transactions: []*valueobject.Transaction{},
	}

	return customer, nil
}
