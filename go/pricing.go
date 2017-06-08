package cabify_challenge

// The interface of any object able to calculate prices
type PriceCalculator interface {
  Calculate(units int) float64
}

// A price policy that applies no discount at all.
type DefaultPrice struct {
  pricePerUnit float64
}

func (p DefaultPrice) Calculate(units int) float64 {
  return p.pricePerUnit * float64(units)
}

// A price policy that applies a discount of 2-for-1.
type GetTwoPayOnePrice struct {
  pricePerUnit float64
}

func (p GetTwoPayOnePrice) Calculate(units int) float64 {
  n := units - (units / 2)
  return p.pricePerUnit * float64(n)
}

// A price policy that applies a discount for bulk purchases.
type BulkDiscountPrice struct {
  pricePerUnit float64
  discountPricePerUnit float64
  bulkMinUnits int
}

func (p BulkDiscountPrice) Calculate(units int) float64 {
  if units < p.bulkMinUnits {
    return p.pricePerUnit * float64(units)
  } else {
    return p.discountPricePerUnit * float64(units)
  }
}
