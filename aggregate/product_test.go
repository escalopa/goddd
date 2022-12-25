package aggregate

import "testing"

func TestProduct_NewProduct(t *testing.T) {
	tests := []struct {
		testName    string
		name        string
		description string
		price       float64
		quantity    int
		expectedErr error
	}{
		{
			testName:    "valid product",
			name:        "name",
			description: "description",
			price:       1.0,
			quantity:    1,
			expectedErr: nil,
		},
		{
			testName:    "empty name",
			name:        "",
			description: "description",
			price:       1.0,
			quantity:    1,
			expectedErr: ErrorEmptyName,
		},
		{
			testName:    "empty description",
			name:        "name",
			description: "",
			price:       1.0,
			quantity:    1,
			expectedErr: ErrorEmptyDescription,
		}, {
			testName:    "small price",
			name:        "name",
			description: "description",
			price:       0.0,
			quantity:    1,
			expectedErr: ErrorSmallPrice,
		}, {
			testName:    "small quantity",
			name:        "name",
			description: "description",
			price:       1.0,
			quantity:    0,
			expectedErr: ErrorSmallQuantity,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			_, err := NewProduct(test.name, test.description, test.price, test.quantity)
			if err != test.expectedErr {
				t.Errorf("expected error %v, got %v", test.expectedErr, err)
			}
		})
	}
}
