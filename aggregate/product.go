package aggregate

import (
	"errors"

	"github.com/escalopa/ddd-go/entity"
	"github.com/google/uuid"
)

var (
	ErrorEmptyDescription = errors.New("description is required")
	ErrorSmallPrice       = errors.New("price must be greater than zero")
	ErrorSmallQuantity    = errors.New("quantity must be greater than zero")
	ErrorEmptyItem        = errors.New("item is cannot be null")
)

type Product struct {
	item     *entity.Item
	price    float64
	quantity int
}

func NewProduct(name string, description string, price float64, quantity int) (*Product, error) {
	p := &Product{
		item: &entity.Item{
			ID:          uuid.New(),
			Name:        name,
			Description: description,
		},
		price:    price,
		quantity: quantity,
	}
	if err := p.Validate(); err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Product) GetID() string {
	return p.item.ID.String()
}

func (p *Product) GetItem() *entity.Item {
	return p.item
}

func (p *Product) GetPrice() float64 {
	return p.price
}

func (p *Product) GetQuantity() int {
	return p.quantity
}

func (p *Product) SetID(id string) error {
	if err := p.item.SetID(id); err != nil {
		return err
	}
	if err := p.Validate(); err != nil {
		return err
	}
	return nil
}

func (p *Product) SetName(name string) error {
	if err := p.item.SetName(name); err != nil {
		return err
	}
	if err := p.Validate(); err != nil {
		return err
	}
	return nil
}

func (p *Product) SetDescription(description string) error {
	if err := p.item.SetDescription(description); err != nil {
		return err
	}
	if err := p.Validate(); err != nil {
		return err
	}
	return nil
}

func (p *Product) SetPrice(price float64) error {
	p.price = price
	if err := p.Validate(); err != nil {
		return err
	}
	return nil
}

func (p *Product) SetQuantity(quantity int) error {
	p.quantity = quantity
	if err := p.Validate(); err != nil {
		return err
	}
	return nil
}

func (p *Product) Validate() error {
	if p.item == nil {
		return ErrorEmptyItem
	}
	if p.item.Name == "" {
		return ErrorEmptyName
	}
	if p.item.Description == "" {
		return ErrorEmptyDescription
	}
	if p.price <= 0 {
		return ErrorSmallPrice
	}
	if p.quantity <= 0 {
		return ErrorSmallQuantity
	}
	return nil
}
