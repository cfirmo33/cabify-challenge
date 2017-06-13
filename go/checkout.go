package cabify_challenge

type Checkout struct {
	scanChan  chan string
	totalChan chan chan float64
}

func CheckoutProcess(cat *Catalog,
	scanChan chan string,
	totalChan chan chan float64) {
	cart := make(map[string]int)
	for {
		select {
		case s := <-scanChan:
			count, present := cart[s]
			if present {
				cart[s] = count + 1
			} else {
				cart[s] = 1
			}
		case t := <-totalChan:
			var total = 0.0
			for k, v := range cart {
				prod, _ := cat.GetProduct(k)
				total += prod.CalculatePrice(v)
			}
			t <- total
		}
	}
}

func NewCheckout(cat *Catalog) *Checkout {
	co := new(Checkout)
	co.scanChan = make(chan string)
	co.totalChan = make(chan chan float64)
	go CheckoutProcess(cat, co.scanChan, co.totalChan)
	return co
}

func (co Checkout) Scan(productName string) {
	co.scanChan <- productName
}

func (co Checkout) Total() float64 {
	rep := make(chan float64)
	co.totalChan <- rep
	return <-rep
}
