package services

import (
	"github.com/google/uuid"
	"testing"
)

func TestCafeImpl_CreateOrder(t *testing.T) {
	cafe, err := NewCafeImpl(WithOrderService(tos))
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name          string
		customerID    string
		itemIDs       []uuid.UUID
		expectedTotal float64
		expectedErr   error
	}{
		{
			name:          "valid order",
			customerID:    tcust.GetID(),
			itemIDs:       []uuid.UUID{uuid.MustParse(tprod[0].GetID()), uuid.MustParse(tprod[1].GetID())},
			expectedTotal: tprod[0].GetPrice() + tprod[1].GetPrice(),
			expectedErr:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cafe.CreateOrder(tt.customerID, tt.itemIDs)
			if err != tt.expectedErr {
				t.Errorf("CafeImpl.CreateOrder() error = %v, expectedErr %v", err, tt.expectedErr)
			}
			if got != tt.expectedTotal {
				t.Errorf("CafeImpl.CreateOrder() total = %v, want %v", got, tt.expectedTotal)
			}
		})
	}
}
