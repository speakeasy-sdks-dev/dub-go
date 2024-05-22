// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/dubinc/dub-go/internal/utils"
	"github.com/dubinc/dub-go/models/components"
)

// PaymentProcessor - The payment processor via which the sale was made.
type PaymentProcessor string

const (
	PaymentProcessorStripe  PaymentProcessor = "stripe"
	PaymentProcessorShopify PaymentProcessor = "shopify"
	PaymentProcessorPaddle  PaymentProcessor = "paddle"
)

func (e PaymentProcessor) ToPointer() *PaymentProcessor {
	return &e
}
func (e *PaymentProcessor) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "stripe":
		fallthrough
	case "shopify":
		fallthrough
	case "paddle":
		*e = PaymentProcessor(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PaymentProcessor: %v", v)
	}
}

type TrackSaleRequestBody struct {
	// This is the unique identifier for the customer in the client's app. This is used to track the customer's journey.
	CustomerID string `json:"customerId"`
	// The amount of the sale. Should be passed in cents.
	Amount int64 `json:"amount"`
	// The payment processor via which the sale was made.
	PaymentProcessor PaymentProcessor `json:"paymentProcessor"`
	// The invoice ID of the sale.
	InvoiceID *string `default:"null" json:"invoiceId"`
	// The currency of the sale. Accepts ISO 4217 currency codes.
	Currency *string `default:"usd" json:"currency"`
	// Additional metadata to be stored with the sale event.
	Metadata map[string]any `json:"metadata,omitempty"`
}

func (t TrackSaleRequestBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TrackSaleRequestBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TrackSaleRequestBody) GetCustomerID() string {
	if o == nil {
		return ""
	}
	return o.CustomerID
}

func (o *TrackSaleRequestBody) GetAmount() int64 {
	if o == nil {
		return 0
	}
	return o.Amount
}

func (o *TrackSaleRequestBody) GetPaymentProcessor() PaymentProcessor {
	if o == nil {
		return PaymentProcessor("")
	}
	return o.PaymentProcessor
}

func (o *TrackSaleRequestBody) GetInvoiceID() *string {
	if o == nil {
		return nil
	}
	return o.InvoiceID
}

func (o *TrackSaleRequestBody) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *TrackSaleRequestBody) GetMetadata() map[string]any {
	if o == nil {
		return nil
	}
	return o.Metadata
}

// TrackSaleResponseBody - A sale was tracked.
type TrackSaleResponseBody struct {
	CustomerID       string         `json:"customerId"`
	Amount           float64        `json:"amount"`
	PaymentProcessor string         `json:"paymentProcessor"`
	InvoiceID        *string        `json:"invoiceId"`
	Currency         string         `json:"currency"`
	Metadata         map[string]any `json:"metadata"`
}

func (o *TrackSaleResponseBody) GetCustomerID() string {
	if o == nil {
		return ""
	}
	return o.CustomerID
}

func (o *TrackSaleResponseBody) GetAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.Amount
}

func (o *TrackSaleResponseBody) GetPaymentProcessor() string {
	if o == nil {
		return ""
	}
	return o.PaymentProcessor
}

func (o *TrackSaleResponseBody) GetInvoiceID() *string {
	if o == nil {
		return nil
	}
	return o.InvoiceID
}

func (o *TrackSaleResponseBody) GetCurrency() string {
	if o == nil {
		return ""
	}
	return o.Currency
}

func (o *TrackSaleResponseBody) GetMetadata() map[string]any {
	if o == nil {
		return nil
	}
	return o.Metadata
}

type TrackSaleResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// A sale was tracked.
	Object *TrackSaleResponseBody
}

func (o *TrackSaleResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *TrackSaleResponse) GetObject() *TrackSaleResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
