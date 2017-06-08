package cabify_challenge

// The interface of any object able to calculate prices
type PriceCalculator interface {
  Calculate(units int) float64
}

// A price policy that applies no discount at all.
type DefaultPrice struct {
  PricePerUnit float64
}

func (p DefaultPrice) Calculate(units int) float64 {
  return p.PricePerUnit * float64(units)
}

// A price policy that applies a discount of 2-for-1.
type GetTwoPayOnePrice struct {
  PricePerUnit float64
}

func (p GetTwoPayOnePrice) Calculate(units int) float64 {
  n := units - (units / 2)
  return p.PricePerUnit * float64(n)
}

// A price policy that applies a discount for bulk purchases.
type BulkDiscountPrice struct {
  PricePerUnit float64
  DiscountPricePerUnit float64
  BulkMinUnits int
}

func (p BulkDiscountPrice) Calculate(units int) float64 {
  if units < p.BulkMinUnits {
    return p.PricePerUnit * float64(units)
  } else {
    return p.DiscountPricePerUnit * float64(units)
  }
}
