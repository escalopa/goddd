package main

import (
	"context"
	"log"

	mng "github.com/escalopa/ddd-go/mongo"
	srv "github.com/escalopa/ddd-go/services"
)

func main() {

	mdb, err := mng.NewMongoDatabase(context.Background(), "mongodb://localhost:27017")
	if err != nil {
		log.Fatal("error connecting to mongo", err)
	}
	customerRepository := mng.NewCustomerRepository(mdb)
	productRepository := mng.NewProductRepository(mdb)

	os, err := srv.NewOrderServiceImpl(
		srv.WithCustomerRepository(customerRepository),
		srv.WithProductRepository(productRepository),
	)

	// or use in-memory repositories
	// os, err := srv.NewOrderServiceImpl(
	// 	srv.WithMemoryCustomerRepository(),
	// 	srv.WithMemoryProductRepository(),
	// )

	if err != nil {
		log.Fatal("error creating order service", err)
	}

	_ = os
}
