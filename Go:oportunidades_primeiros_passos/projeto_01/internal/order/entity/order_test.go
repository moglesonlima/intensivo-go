package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Testes de unidade (testing é um pacote do go & o T é uma struct que executa os testes)

func TestGivenAnEmptyID_WhenCreateAnOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{}
	assert.Error(t, order.isValid(), "Invalid ID")
}

func TestGivenAnEmptyPrice_WhenCreateAnOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.isValid(), "Invalid Price")
}

func TestGivenAnEmptyTax_WhenCreateAnOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{ID: "123", Price: 10}
	assert.Error(t, order.isValid(), "Invalid Tax")
}

func TestGivenAvalidParams_WhenOrder_ThenIShouldReceiveCreateOrderWithAllParams(t *testing.T) {
	order := Order{ID: "123", Price: 10, Tax: 2}
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
	assert.Nil(t, order.isValid())
	// assert.Equal(t, order.FinalPrice, 12.0)
}
func TestGivenAvalidParams_WhenOrderFunc_ThenIShouldReceiveCreateOrderWithAllParams(t *testing.T) {
	order, err := NewOrder("123", 10, 2)
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
	assert.Nil(t, err)
}

func TestGivenAPriceAndTax_WhenICallCalculatePrice_thenIShouldSetFinalPrice(t *testing.T) {
	order, err := NewOrder("123", 10, 2)
	assert.Nil(t, err)
	assert.Nil(t, order.CalculateFinalPrice())
	assert.Equal(t, 12.0, order.FinalPrice)
}
