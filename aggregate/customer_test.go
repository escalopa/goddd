package aggregate

import "testing"

func TestCustomer(t *testing.T) {
	tests := []struct {
		name          string
		customerName  string
		expectedError error
	}{
		{
			name:          "empty name",
			customerName:  "",
			expectedError: ErrorEmptyName,
		},
		{
			name:          "valid name",
			customerName:  "John Doe",
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			customer, err := NewCustomer(test.customerName)
			if err != test.expectedError {
				t.Errorf("expected error %v, got %v", test.expectedError, err)
			}
			if customer != nil {
				if customer.person.Name != test.customerName {
					t.Errorf("expected customer name %s, got %s", test.customerName, customer.person.Name)
				}
			}
		})
	}
}
