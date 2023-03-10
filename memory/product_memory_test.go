package memory

import (
	"context"
	"github.com/escalopa/ddd-go/repository"
	"testing"

	"github.com/escalopa/ddd-go/aggregate"
)

func TestProductMemoryRepository_Find(t *testing.T) {
	testCases := []struct {
		name        string
		id          string
		expectedErr error
	}{
		{
			name:        "No Product By ID",
			id:          "f47ac10b-58cc-0372-8567-0e02b2c3d479",
			expectedErr: repository.ErrorProductNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewMemoryProductRepository()
			_, err := repo.Find(context.Background(), tc.id)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestProductMemoryRepository_FindAll(t *testing.T) {
	testCases := []struct {
		name        string
		expectedErr error
	}{
		{
			name:        "All Products",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewMemoryProductRepository()
			_, err := repo.FindAll(context.Background())
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestProductMemoryRepository_Save(t *testing.T) {
	testCases := []struct {
		name        string
		prod        string
		description string
		expectedErr error
	}{
		{
			name:        "Add Product",
			prod:        "Product 1",
			description: "Product 1 Description",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewMemoryProductRepository()

			prod, err := aggregate.NewProduct(tc.prod, tc.description, 10.00, 1)
			if err != nil {
				t.Fatal(err)
			}

			err = repo.Save(context.Background(), prod)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

		})
	}
}

func TestProductMemoryRepository_Update(t *testing.T) {
	testCases := []struct {
		name        string
		prod        string
		description string
		expectedErr error
	}{
		{
			name:        "Update Product",
			prod:        "Product 1",
			description: "Product 1 Description",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewMemoryProductRepository()

			prod, err := aggregate.NewProduct(tc.prod, tc.description, 10.00, 1)
			if err != nil {
				t.Fatal(err)
			}

			err = repo.Save(context.Background(), prod)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

			// TODO - Update Product
			err = repo.Update(context.Background(), prod)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

		})
	}
}

func TestProductMemoryRepository_Delete(t *testing.T) {
	testCases := []struct {
		name        string
		prod        string
		description string
		expectedErr error
	}{
		{
			name:        "Delete Product",
			prod:        "Product 1",
			description: "Product 1 Description",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewMemoryProductRepository()

			prod, err := aggregate.NewProduct(tc.prod, tc.description, 10.00, 1)
			if err != nil {
				t.Fatal(err)
			}

			err = repo.Save(context.Background(), prod)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

			err = repo.Delete(context.Background(), prod.GetID())
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

			if err == nil {
				_, err = repo.Find(context.Background(), prod.GetID())
				if err != repository.ErrorProductNotFound {
					t.Errorf("Expected error %v, got %v", repository.ErrorProductNotFound, err)
				}
			}
		})
	}
}
