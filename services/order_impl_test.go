package services

import (
	"github.com/escalopa/ddd-go/aggregate"
	"github.com/escalopa/ddd-go/domain/repository"
	"log"
	"strconv"
	"testing"
)

var tos OrderService
var tcust *aggregate.Customer
var tprod []*aggregate.Product

func TestMain(m *testing.M) {
	cr, pr, err := createRandomRepositories()
	if err != nil {
		log.Fatal(err)
	}

	tos, err = NewOrderServiceImpl(
		WithCustomerRepository(cr),
		WithProductRepository(pr),
	)

	if err != nil {
		log.Fatal(err)
	}
	m.Run()
}

func TestOrderServiceImpl_PlaceOrder(t *testing.T) {
	var itemIDs []string
	var expectedTotal float64
	for _, p := range tprod {
		expectedTotal += p.GetPrice()
		itemIDs = append(itemIDs, p.GetID())
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
	cr = repository.NewMemoryCustomerRepository()
	pr = repository.NewMemoryProductRepository()
	tcust, err = createRandomCustomer()
	if err != nil {
		return nil, nil, err
	}
	err = cr.Save(tcust)
	if err != nil {
		return nil, nil, err
	}

	tprod, err = createRandomProducts()
	if err != nil {
		return nil, nil, err
	}
	for _, p := range tprod {
		err = pr.Save(p)
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
