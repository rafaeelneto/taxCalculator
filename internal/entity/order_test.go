package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IfIGetAnErrorIfIDIsBlank (t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "Invalid ID")
}
func Test_IfIGetAnErrorIfPriceIsBlank (t *testing.T) {
	order := Order{ID: "fdasfa"}
	assert.Error(t, order.Validate(), "Invalid Price")
}
func Test_IfIGetAnErrorIfTaxIsBlank (t *testing.T) {
	order := Order{ID: "fdasfa", Price: 1.2}
	assert.Error(t, order.Validate(), "Invalid Tax")
}
func Test_IfIGetAnErrorIfQuantityIsBlank (t *testing.T) {
	order := Order{ID: "fdasfa", Price: 1.2, Tax: 0.1}
	assert.Error(t, order.Validate(), "Invalid Quantity")
}
func Test_AllFiedsValid (t *testing.T) {
	order, _ := NewOrder("123", 10.3, 0.1, 10)
	assert.NoError(t, order.Validate())
	assert.Equal(t, 10.3, order.Price)
	assert.Equal(t, 0.1, order.Tax)
	assert.Equal(t, 10, order.Quantity)
	assert.Equal(t, 104.0, order.FinalPrice)
}