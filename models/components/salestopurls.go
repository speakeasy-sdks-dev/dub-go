// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type SalesTopUrls struct {
	// The destination URL
	URL string `json:"url"`
	// The number of sales from this URL
	Sales float64 `json:"sales"`
	// The total amount of sales from this URL
	Amount float64 `json:"amount"`
}

func (o *SalesTopUrls) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

func (o *SalesTopUrls) GetSales() float64 {
	if o == nil {
		return 0.0
	}
	return o.Sales
}

func (o *SalesTopUrls) GetAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.Amount
}
