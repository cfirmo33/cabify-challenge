package cabify_challenge

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCatalogIsEmpty(t *testing.T) {
	cat := NewCatalog()

	assert.True(t, cat.IsEmpty())
}

func TestAddAndGetCatalogProducts(t *testing.T) {
	prod1 := Product{"XX", "description", DefaultPrice{PricePerUnit: 1.00}}
	prod2 := Product{"YY", "description", DefaultPrice{PricePerUnit: 2.00}}
	prod3 := Product{"ZZ", "description", DefaultPrice{PricePerUnit: 3.00}}
	cat := NewCatalog()
	cat.AddProduct(prod1)
	cat.AddProduct(prod2)
	cat.AddProduct(prod3)

	rep1, _ := cat.GetProduct("XX")
	assert.Equal(t, prod1, *rep1)
	rep2, _ := cat.GetProduct("YY")
	assert.Equal(t, prod2, *rep2)
	rep3, _ := cat.GetProduct("ZZ")
	assert.Equal(t, prod3, *rep3)
}

func TestAddProductsToCatalogOverwriteExisting(t *testing.T) {
	prod1 := Product{"XX", "description", DefaultPrice{PricePerUnit: 1.00}}
	prod2 := Product{"XX", "other description", DefaultPrice{PricePerUnit: 2.00}}
	cat := NewCatalog()
	cat.AddProduct(prod1)
	cat.AddProduct(prod2)

	rep, _ := cat.GetProduct("XX")
	assert.Equal(t, prod2, *rep)
}

func TestContainsProduct(t *testing.T) {
	prod := Product{"XX", "description", DefaultPrice{PricePerUnit: 1.00}}
	cat := NewCatalog()
	cat.AddProduct(prod)

	assert.True(t, cat.ContainsProduct("XX"))
	assert.False(t, cat.ContainsProduct("YY"))
}

func TestGetProductFailsIfNotPresent(t *testing.T) {
	cat := NewCatalog()
	_, err := cat.GetProduct("XX")
	assert.EqualError(t, err, "no product found for such code XX")
}

func TestDefaultCatalogIsNotEmpty(t *testing.T) {
	cat := DefaultCatalog()

	assert.False(t, cat.IsEmpty())
}
