package cabify_challenge

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyCheckout(t *testing.T) {
	co := NewCheckout(DefaultCatalog())
	assert.Equal(t, 0.00, co.GetTotal())
}

func TestFailedScanningUnknownProducts(t *testing.T) {
	co := NewCheckout(DefaultCatalog())
	err := co.Scan("NAZGUL")
	assert.EqualError(t, err, "cannot scan unknown product with code NAZGUL")
}

func TestNoDiscountAppliedToAny(t *testing.T) {
	co := NewCheckout(DefaultCatalog())
	co.Scan("VOUCHER")
	co.Scan("TSHIRT")
	co.Scan("MUG")
	assert.Equal(t, 32.50, co.GetTotal())
}

func TestGetTwoPayOneDiscountApplied(t *testing.T) {
	co := NewCheckout(DefaultCatalog())
	co.Scan("VOUCHER")
	co.Scan("TSHIRT")
	co.Scan("VOUCHER")
	assert.Equal(t, 25.00, co.GetTotal())
}

func TestBulkDiscountApplied(t *testing.T) {
	co := NewCheckout(DefaultCatalog())
	co.Scan("TSHIRT")
	co.Scan("TSHIRT")
	co.Scan("TSHIRT")
	co.Scan("VOUCHER")
	co.Scan("TSHIRT")
	assert.Equal(t, 81.00, co.GetTotal())
}

func TestMultipleDiscountsApplied(t *testing.T) {
	co := NewCheckout(DefaultCatalog())
	co.Scan("VOUCHER")
	co.Scan("TSHIRT")
	co.Scan("VOUCHER")
	co.Scan("VOUCHER")
	co.Scan("MUG")
	co.Scan("TSHIRT")
	co.Scan("TSHIRT")
	assert.Equal(t, 74.50, co.GetTotal())
}
