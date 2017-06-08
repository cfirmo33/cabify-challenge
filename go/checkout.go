package cabify_challenge

type PricingRules map[string]PriceCalculator

func DefaultPrices() PricingRules {
  return PricingRules {
    "VOUCHER": GetTwoPayOnePrice { PricePerUnit: 5.00 },
    "TSHIRT": BulkDiscountPrice { PricePerUnit: 20.00, DiscountPricePerUnit: 19.00, BulkMinUnits: 3 },
    "MUG": DefaultPrice { PricePerUnit: 7.50 },
  }
}

type Checkout struct {
  prices PricingRules
  scanChan chan string
  totalChan chan chan float64
}

func CheckoutProcess(prices PricingRules,
                     scanChan chan string,
                     totalChan chan chan float64) {
  cart := make(map[string]int)
  for {
    select {
    case s := <- scanChan:
      count, present := cart[s]
      if present {
        cart[s] = count + 1
      } else {
        cart[s] = 1
      }
    case t := <- totalChan:
      var total = 0.0
      for k, v := range cart {
        total += prices[k].Calculate(v)
      }
      t <- total
    }
  }
}

func NewCheckout(prices PricingRules) *Checkout {
  co := new(Checkout)
  co.prices = prices
  co.scanChan = make(chan string)
  co.totalChan = make(chan chan float64)
  go CheckoutProcess(prices, co.scanChan, co.totalChan)
  return co
}

func (co Checkout) Scan(productName string) {
  co.scanChan <- productName
}

func (co Checkout) Total() float64 {
  rep := make (chan float64)
  co.totalChan <- rep
  return <- rep
}
