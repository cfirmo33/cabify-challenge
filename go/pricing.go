package cabify_challenge

// The interface of any object able to calculate prices
type PriceCalculator interface {
  Calculate(units int) float64
}

// A price policy that applies no discount at all.
type DefaultPrice struct {
  pricePerUnit float64
}

func NewDefaultPrice(p float64) *DefaultPrice {
  dp := new(DefaultPrice)
  dp.pricePerUnit = p
  return dp
}

func (p DefaultPrice) Calculate(units int) float64 {
  return p.pricePerUnit * float64(units)
}

// A price policy that applies a discount of 2-for-1.
type GetTwoPayOnePrice struct {
  pricePerUnit float64
}

func NewGetTwoPayOnePrice(p float64) *GetTwoPayOnePrice {
  dp := new(GetTwoPayOnePrice)
  dp.pricePerUnit = p
  return dp
}

func (p GetTwoPayOnePrice) Calculate(units int) float64 {
  n := units - (units / 2)
  return p.pricePerUnit * float64(n)
}
