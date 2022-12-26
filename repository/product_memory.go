package repository

import (
	"context"
	"sync"

	"github.com/escalopa/ddd-go/aggregate"
)

type ProductMemoryRepository struct {
	products map[string]*aggregate.Product
	sync.Mutex
}

func NewMemoryProductRepository() ProductRepository {
	return &ProductMemoryRepository{
		products: map[string]*aggregate.Product{},
	}
}

func (r *ProductMemoryRepository) Find(_ context.Context, id string) (*aggregate.Product, error) {
	product, ok := r.products[id]
	if !ok {
		return nil, ErrorProductNotFound
	}
	return product, nil
}

func (r *ProductMemoryRepository) FindAll(_ context.Context) ([]*aggregate.Product, error) {
	var products []*aggregate.Product
	for _, product := range r.products {
		products = append(products, product)
	}
	return products, nil
}

func (r *ProductMemoryRepository) Save(_ context.Context, product *aggregate.Product) error {
	r.Lock()
	defer r.Unlock()

	r.products[product.GetID()] = product
	return nil
}

func (r *ProductMemoryRepository) Update(_ context.Context, product *aggregate.Product) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.products[product.GetID()]; !ok {
		return ErrorProductNotFound
	}

	r.products[product.GetID()] = product
	return nil
}

func (r *ProductMemoryRepository) Delete(_ context.Context, id string) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.products[id]; !ok {
		return ErrorProductNotFound
	}

	delete(r.products, id)
	return nil
}
