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
