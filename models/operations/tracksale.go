// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/dubinc/dub-go/internal/utils"
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
	// The name of the sale event. It can be used to track different types of event for example 'Purchase', 'Upgrade', 'Payment', etc.
	EventName *string `default:"Purchase" json:"eventName"`
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

func (o *TrackSaleRequestBody) GetEventName() *string {
	if o == nil {
		return nil
	}
	return o.EventName
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

type TrackSaleCustomer struct {
	ID     string  `json:"id"`
	Name   *string `json:"name"`
	Email  *string `json:"email"`
	Avatar *string `json:"avatar"`
}

func (o *TrackSaleCustomer) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *TrackSaleCustomer) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TrackSaleCustomer) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *TrackSaleCustomer) GetAvatar() *string {
	if o == nil {
		return nil
	}
	return o.Avatar
}

type Sale struct {
	Amount           float64        `json:"amount"`
	Currency         string         `json:"currency"`
	PaymentProcessor string         `json:"paymentProcessor"`
	InvoiceID        *string        `json:"invoiceId"`
	Metadata         map[string]any `json:"metadata"`
}

func (o *Sale) GetAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.Amount
}

func (o *Sale) GetCurrency() string {
	if o == nil {
		return ""
	}
	return o.Currency
}

func (o *Sale) GetPaymentProcessor() string {
	if o == nil {
		return ""
	}
	return o.PaymentProcessor
}

func (o *Sale) GetInvoiceID() *string {
	if o == nil {
		return nil
	}
	return o.InvoiceID
}

func (o *Sale) GetMetadata() map[string]any {
	if o == nil {
		return nil
	}
	return o.Metadata
}

// TrackSaleResponseBody - A sale was tracked.
type TrackSaleResponseBody struct {
	EventName string            `json:"eventName"`
	Customer  TrackSaleCustomer `json:"customer"`
	Sale      Sale              `json:"sale"`
}

func (o *TrackSaleResponseBody) GetEventName() string {
	if o == nil {
		return ""
	}
	return o.EventName
}

func (o *TrackSaleResponseBody) GetCustomer() TrackSaleCustomer {
	if o == nil {
		return TrackSaleCustomer{}
	}
	return o.Customer
}

func (o *TrackSaleResponseBody) GetSale() Sale {
	if o == nil {
		return Sale{}
	}
	return o.Sale
}
