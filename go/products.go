package cabify_challenge

import (
	"fmt"
)

// A product that can be sell in our online store
type Product struct {
	Code  string
	Name  string
	Price PriceCalculator
}

// A convenience function to calculate price for the given units
func (p Product) CalculatePrice(units int) float64 {
	return p.Price.Calculate(units)
}

// A catalog of the products available in the online store
type Catalog struct {
	Content map[string]Product
}

// Create a new fresh and empty catalog
func NewCatalog() *Catalog {
	cat := new(Catalog)
	cat.Content = make(map[string]Product)
	return cat
}

// Check whether the catalog is empty
func (c Catalog) IsEmpty() bool {
	return len(c.Content) == 0
}

// Check if there is a product with the given code
func (c Catalog) ContainsProduct(code string) bool {
	_, found := c.Content[code]
	return found
}

// Add a new element to the catalog
func (c Catalog) AddProduct(p Product) {
	c.Content[p.Code] = p
}

// Get a product by its code, or return error if not found
func (c Catalog) GetProduct(code string) (*Product, error) {
	prod, ok := c.Content[code]
	if ok {
		return &prod, nil
	} else {
		return nil, fmt.Errorf("no product found for such code %s", code)
	}
}

// The default catalog of our online store
func DefaultCatalog() *Catalog {
	cat := NewCatalog()
	cat.AddProduct(Product{
		"VOUCHER",
		"Cabify Voucher",
		GetTwoPayOnePrice{PricePerUnit: 5.00},
	})
	cat.AddProduct(Product{
		"TSHIRT",
		"Cabify T-Shirt",
		BulkDiscountPrice{
			PricePerUnit:         20.00,
			DiscountPricePerUnit: 19.00,
			BulkMinUnits:         3},
	})
	cat.AddProduct(Product{
		"MUG",
		"Cafify Coffee Mug",
		DefaultPrice{PricePerUnit: 7.50},
	})
	return cat
}
