package services

import (
	"context"
	"github.com/escalopa/ddd-go/memory"
	"strconv"
	"testing"

	"github.com/escalopa/ddd-go/aggregate"
	"github.com/escalopa/ddd-go/repository"
	"github.com/google/uuid"
)

var tos OrderService
var tcust *aggregate.Customer
var tprod []*aggregate.Product

func TestOrderServiceImpl_PlaceOrder(t *testing.T) {
	var itemIDs []uuid.UUID
	var expectedTotal float64
	for _, p := range tprod {
		expectedTotal += p.GetPrice()
		itemIDs = append(itemIDs, uuid.MustParse(p.GetID()))
	}

	total, err := tos.PlaceOrder(tcust.GetID(), itemIDs)
	if err != nil {
		t.Fatal(err)
	}
	if total != expectedTotal {
		t.Fatalf("Expected total %f, got %f", expectedTotal, total)
	}
}

func createRandomRepositories() (cr repository.CustomerRepository, pr repository.ProductRepository, err error) {
	cr = memory.NewMemoryCustomerRepository()
	pr = memory.NewMemoryProductRepository()
	tcust, err = createRandomCustomer()
	if err != nil {
		return nil, nil, err
	}
	err = cr.Save(context.Background(), tcust)
	if err != nil {
		return nil, nil, err
	}

	tprod, err = createRandomProducts()
	if err != nil {
		return nil, nil, err
	}
	for _, p := range tprod {
		err = pr.Save(context.Background(), p)
		if err != nil {
			return nil, nil, err
		}
	}

	return cr, pr, nil
}

func createRandomCustomer() (c *aggregate.Customer, err error) {
	c, err = aggregate.NewCustomer("John")
	if err != nil {
		return nil, err
	}
	return c, nil
}

func createRandomProducts() (products []*aggregate.Product, err error) {
	for i := 0; i < 10; i++ {
		p, err := aggregate.NewProduct("Product "+strconv.Itoa(i), "Product Description"+strconv.Itoa(i), 10.0, 1.0)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
