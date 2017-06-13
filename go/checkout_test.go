package cabify_challenge

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScenario1(t *testing.T) {
	co := NewCheckout(DefaultPrices())
	co.Scan("VOUCHER")
	co.Scan("TSHIRT")
	co.Scan("MUG")
	assert.Equal(t, 32.50, co.Total())
}

func TestScenario2(t *testing.T) {
	co := NewCheckout(DefaultPrices())
	co.Scan("VOUCHER")
	co.Scan("TSHIRT")
	co.Scan("VOUCHER")
	assert.Equal(t, 25.00, co.Total())
}

func TestScenario3(t *testing.T) {
	co := NewCheckout(DefaultPrices())
	co.Scan("TSHIRT")
	co.Scan("TSHIRT")
	co.Scan("TSHIRT")
	co.Scan("VOUCHER")
	co.Scan("TSHIRT")
	assert.Equal(t, 81.00, co.Total())
}

func TestScenario4(t *testing.T) {
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
