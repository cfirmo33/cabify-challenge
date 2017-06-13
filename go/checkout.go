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
			HandleScanRequest(cart, cat, scan)
		case total := <-totalChan:
			HandleGetTotalRequest(cart, cat, total)
		}
	}
}

func HandleScanRequest(cart map[string]int, cat *Catalog, req ScanReq) {
	code := req.ProductCode
	if cat.ContainsProduct(code) {
		count, present := cart[code]
		if present {
			cart[code] = count + 1
		} else {
			cart[code] = 1
		}
		req.Rep <- nil
	} else {
		req.Rep <- fmt.Errorf(
			"cannot scan unknown product with code %s", code)
	}
}

func HandleGetTotalRequest(cart map[string]int, cat *Catalog, req GetTotalReq) {
	var result = 0.0
	for k, v := range cart {
		// Here we can safely ignore the error since `scan()` operation
		// already ensures the product exists in the catalog
		prod, _ := cat.GetProduct(k)
		result += prod.CalculatePrice(v)
	}
	req.Rep <- result
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

func (co Checkout) GetTotal() float64 {
	req := GetTotalReq{make(chan float64)}
	co.totalChan <- req
	return <-req.Rep
}
