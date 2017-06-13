package cabify_challenge

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyCheckout(t *testing.T) {
	co := NewCheckout(DefaultPrices())
	assert.Equal(t, 0.00, co.Total())
}

func TestNoDiscountAppliedToAny(t *testing.T) {
	co := NewCheckout(DefaultPrices())
	co.Scan("VOUCHER")
	co.Scan("TSHIRT")
	co.Scan("MUG")
	assert.Equal(t, 32.50, co.Total())
}

func TestGetTwoPayOneDiscountApplied(t *testing.T) {
	co := NewCheckout(DefaultPrices())
	co.Scan("VOUCHER")
	co.Scan("TSHIRT")
	co.Scan("VOUCHER")
	assert.Equal(t, 25.00, co.Total())
}

func TestBulkDiscountApplied(t *testing.T) {
	co := NewCheckout(DefaultPrices())
	co.Scan("TSHIRT")
	co.Scan("TSHIRT")
	co.Scan("TSHIRT")
	co.Scan("VOUCHER")
	co.Scan("TSHIRT")
	assert.Equal(t, 81.00, co.Total())
}

func TestMultipleDiscountsApplied(t *testing.T) {
	co := NewCheckout(DefaultPrices())
	co.Scan("VOUCHER")
	co.Scan("TSHIRT")
	co.Scan("VOUCHER")
	co.Scan("VOUCHER")
	co.Scan("MUG")
	co.Scan("TSHIRT")
	co.Scan("TSHIRT")
	assert.Equal(t, 74.50, co.Total())
}
