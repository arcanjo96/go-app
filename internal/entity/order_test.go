package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfItGetsAnErrorIfIDIsBlank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "id is required")
}

func TestIfItGetsAnErrorIfPriceIsBlank(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.Validate(), "price must be greater than zero")
}

func TestIfItGetsAnErrorIfTaxIsBlank(t *testing.T) {
	order := Order{ID: "123", Price: 10.0}
	assert.Error(t, order.Validate(), "invalid tax")
}

func TestFinalPrice(t *testing.T) {
	order := Order{ID: "123", Price: 10.0, Tax: 3.0}
	assert.NoError(t, order.Validate())
	assert.Equal(t, 3.0, order.Tax)
	order.CalculateFinalPrice()
	assert.Equal(t, 13.0, order.FinalPrice)
}
