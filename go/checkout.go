package cabify_challenge

import "fmt"

type ScanReq struct {
	ProductCode string
	Rep         chan error
}

type GetTotalReq struct {
	Rep chan float64
}

type Checkout struct {
	scanChan  chan ScanReq
	totalChan chan GetTotalReq
}

func CheckoutProcess(cat *Catalog,
	scanChan chan ScanReq,
	totalChan chan GetTotalReq) {
	cart := make(map[string]int)
	for {
		select {
		case scan := <-scanChan:
			code := scan.ProductCode
			if cat.ContainsProduct(code) {
				count, present := cart[code]
				if present {
					cart[code] = count + 1
				} else {
					cart[code] = 1
				}
				scan.Rep <- nil
			} else {
				scan.Rep <- fmt.Errorf(
					"cannot scan unknown product with code %s", code)
			}
		case total := <-totalChan:
			var result = 0.0
			for k, v := range cart {
				prod, _ := cat.GetProduct(k)
				result += prod.CalculatePrice(v)
			}
			total.Rep <- result
		}
	}
}

func NewCheckout(cat *Catalog) *Checkout {
	co := new(Checkout)
	co.scanChan = make(chan ScanReq)
	co.totalChan = make(chan GetTotalReq)
	go CheckoutProcess(cat, co.scanChan, co.totalChan)
	return co
}

func (co Checkout) Scan(productCode string) error {
	req := ScanReq{productCode, make(chan error)}
	co.scanChan <- req
	return <-req.Rep
}

func (co Checkout) Total() float64 {
	req := GetTotalReq{make(chan float64)}
	co.totalChan <- req
	return <-req.Rep
}
