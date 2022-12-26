package services

import (
	"log"
	"testing"
)

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
