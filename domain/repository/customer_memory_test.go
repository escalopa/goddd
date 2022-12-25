package repository

import (
	"testing"

	"github.com/escalopa/ddd-go/aggregate"
)

func TestCustomerMemoryRepository_Find(t *testing.T) {

	// Create a fake customer to add to repository
	customer, err := aggregate.NewCustomer("Percy")
	if err != nil {
		t.Fatal(err)
	}
	id := customer.GetID()
	repo := NewMemoryCustomerRepository()
	err = repo.Save(customer)
	if err != nil {
		return
	}

	testCases := []struct {
		name        string
		id          string
		expectedErr error
	}{
		{
			name:        "No Customer By ID",
			id:          "f47ac10b-58cc-0372-8567-0e02b2c3d479",
			expectedErr: ErrorCustomerNotFound,
		}, {
			name:        "Customer By ID",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			_, err := repo.Find(tc.id)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestCustomerMemoryRepository_Save(t *testing.T) {
	testCases := []struct {
		name        string
		cust        string
		expectedErr error
	}{
		{
			name:        "Add Customer",
			cust:        "Percy",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewMemoryCustomerRepository()

			cust, err := aggregate.NewCustomer(tc.cust)
			if err != nil {
				t.Fatal(err)
			}

			err = repo.Save(cust)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

			found, err := repo.Find(cust.GetID())
			if err != nil {
				t.Fatal(err)
			}

			if found.GetID() != cust.GetID() {
				t.Errorf("Expected %v, got %v", cust.GetID(), found.GetID())
			}
		})
	}
}

func TestCustomerMemoryRepository_Update(t *testing.T) {
	tests := []struct {
		testName    string
		name        string
		newName     string
		expectedErr error
	}{
		{
			testName:    "Update Customer",
			name:        "Percy",
			newName:     "Percy Updated",
			expectedErr: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			repo := NewMemoryCustomerRepository()
			customer, err := aggregate.NewCustomer(tc.name)
			if err != nil {
				t.Fatal(err)
			}
			err = repo.Save(customer)
			if err != nil {
				t.Fatal(err)
			}
			err = customer.UpdateName(tc.newName)
			if err != nil {
				t.Fatal(err)
			}
			err = repo.Update(customer)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			if err == nil {
				found, err := repo.Find(customer.GetID())
				if err != nil {
					t.Fatal(err)
				}
				if found.GetName() != customer.GetName() {
					t.Errorf("Expected %v, got %v", customer.GetName(), found.GetName())
				}
			}
		})
	}
}
