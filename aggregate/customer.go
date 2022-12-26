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

func NewCustomer(name string) (*Customer, error) {
	customer := &Customer{
		person: &entity.Person{
			Name: name,
			ID:   uuid.New(),
		},
		items:        []*entity.Item{},
		transactions: []*valueobject.Transaction{},
	}

	if err := customer.Validate(); err != nil {
		return nil, err
	}

	return customer, nil
}

func (c *Customer) GetID() string {
	return c.person.ID.String()
}

func (c *Customer) GetName() string {
	return c.person.Name
}

func (c *Customer) UpdateName(name string) error {
	if name == "" {
		return ErrorEmptyName
	}

	c.person.Name = name
	return nil
}

func (c *Customer) Validate() error {
	if c.person.Name == "" {
		return ErrorEmptyName
	}

	return nil
}

func (c *Customer) AddItem(product *entity.Item) {
	c.items = append(c.items, product)
}
