package repository

import (
	"context"
	"errors"

	"github.com/escalopa/ddd-go/aggregate"
)

var (
	ErrorCustomerNotFound     = errors.New("customer not found")
	ErrorCustomerNameOldEqNew = errors.New("customer name old eq new")

	ErrorProductNotFound = errors.New("product not found")
)

type CustomerRepository interface {
	Find(ctx context.Context, uuid string) (*aggregate.Customer, error)
	Save(ctx context.Context, customer *aggregate.Customer) error
	Update(ctx context.Context, customer *aggregate.Customer) error
}

type ProductRepository interface {
	Find(ctx context.Context, uuid string) (*aggregate.Product, error)
	Save(ctx context.Context, product *aggregate.Product) error
	Update(ctx context.Context, product *aggregate.Product) error
	FindAll(ctx context.Context) ([]*aggregate.Product, error)
	Delete(ctx context.Context, uuid string) error
}
