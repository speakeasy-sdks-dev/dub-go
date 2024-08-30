// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/dubinc/dub-go/internal/utils"
	"github.com/dubinc/dub-go/models/components"
)

// QueryParamEvent - The type of event to retrieve analytics for. Defaults to 'clicks'.
type QueryParamEvent string

const (
	QueryParamEventClicks    QueryParamEvent = "clicks"
	QueryParamEventLeads     QueryParamEvent = "leads"
	QueryParamEventSales     QueryParamEvent = "sales"
	QueryParamEventComposite QueryParamEvent = "composite"
)

func (e QueryParamEvent) ToPointer() *QueryParamEvent {
	return &e
}
func (e *QueryParamEvent) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "clicks":
		fallthrough
	case "leads":
		fallthrough
	case "sales":
		fallthrough
	case "composite":
		*e = QueryParamEvent(v)
		return nil
	default:
		return fmt.Errorf("invalid value for QueryParamEvent: %v", v)
	}
}

// QueryParamInterval - The interval to retrieve analytics for. Takes precedence over start and end. If undefined, defaults to 24h.
type QueryParamInterval string

const (
	QueryParamIntervalTwentyFourh   QueryParamInterval = "24h"
	QueryParamIntervalSevend        QueryParamInterval = "7d"
	QueryParamIntervalThirtyd       QueryParamInterval = "30d"
	QueryParamIntervalNinetyd       QueryParamInterval = "90d"
	QueryParamIntervalYtd           QueryParamInterval = "ytd"
	QueryParamIntervalOney          QueryParamInterval = "1y"
	QueryParamIntervalAll           QueryParamInterval = "all"
	QueryParamIntervalAllUnfiltered QueryParamInterval = "all_unfiltered"
)

func (e QueryParamInterval) ToPointer() *QueryParamInterval {
	return &e
}
func (e *QueryParamInterval) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "24h":
		fallthrough
	case "7d":
		fallthrough
	case "30d":
		fallthrough
	case "90d":
		fallthrough
	case "ytd":
		fallthrough
	case "1y":
		fallthrough
	case "all":
		fallthrough
	case "all_unfiltered":
		*e = QueryParamInterval(v)
		return nil
	default:
		return fmt.Errorf("invalid value for QueryParamInterval: %v", v)
	}
}

type Order string

const (
	OrderAsc  Order = "asc"
	OrderDesc Order = "desc"
)

func (e Order) ToPointer() *Order {
	return &e
}
func (e *Order) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = Order(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Order: %v", v)
	}
}

type SortBy string

const (
	SortByTimestamp SortBy = "timestamp"
)

func (e SortBy) ToPointer() *SortBy {
	return &e
}
func (e *SortBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "timestamp":
		*e = SortBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SortBy: %v", v)
	}
}

type ListEventsRequest struct {
	// The type of event to retrieve analytics for. Defaults to 'clicks'.
	Event *QueryParamEvent `default:"clicks" queryParam:"style=form,explode=true,name=event"`
	// The domain to filter analytics for.
	Domain *string `queryParam:"style=form,explode=true,name=domain"`
	// The short link slug.
	Key *string `queryParam:"style=form,explode=true,name=key"`
	// The unique ID of the short link on Dub.
	LinkID *string `queryParam:"style=form,explode=true,name=linkId"`
	// This is the ID of the link in the your database. Must be prefixed with 'ext_' when passed as a query parameter.
	ExternalID *string `queryParam:"style=form,explode=true,name=externalId"`
	// The interval to retrieve analytics for. Takes precedence over start and end. If undefined, defaults to 24h.
	Interval *QueryParamInterval `queryParam:"style=form,explode=true,name=interval"`
	// The start date and time when to retrieve analytics from.
	Start *string `queryParam:"style=form,explode=true,name=start"`
	// The end date and time when to retrieve analytics from. If not provided, defaults to the current date.
	End *string `queryParam:"style=form,explode=true,name=end"`
	// The IANA time zone code for aligning timeseries granularity (e.g. America/New_York). Defaults to UTC.
	Timezone *string `default:"UTC" queryParam:"style=form,explode=true,name=timezone"`
	// The continent to retrieve analytics for.
	Continent *components.ContinentCode `queryParam:"style=form,explode=true,name=continent"`
	// The country to retrieve analytics for.
	Country *components.CountryCode `queryParam:"style=form,explode=true,name=country"`
	// The city to retrieve analytics for.
	City *string `queryParam:"style=form,explode=true,name=city"`
	// The device to retrieve analytics for.
	Device *string `queryParam:"style=form,explode=true,name=device"`
	// The browser to retrieve analytics for.
	Browser *string `queryParam:"style=form,explode=true,name=browser"`
	// The OS to retrieve analytics for.
	Os *string `queryParam:"style=form,explode=true,name=os"`
	// The referer to retrieve analytics for.
	Referer *string `queryParam:"style=form,explode=true,name=referer"`
	// The URL to retrieve analytics for.
	URL *string `queryParam:"style=form,explode=true,name=url"`
	// The tag ID to retrieve analytics for.
	TagID *string `queryParam:"style=form,explode=true,name=tagId"`
	// Filter for QR code scans. If true, filter for QR codes only. If false, filter for links only. If undefined, return both.
	Qr *bool `queryParam:"style=form,explode=true,name=qr"`
	// Filter for root domains. If true, filter for domains only. If false, filter for links only. If undefined, return both.
	Root   *bool    `queryParam:"style=form,explode=true,name=root"`
	Page   *float64 `default:"0" queryParam:"style=form,explode=true,name=page"`
	Limit  *float64 `default:"100" queryParam:"style=form,explode=true,name=limit"`
	Order  *Order   `default:"desc" queryParam:"style=form,explode=true,name=order"`
	SortBy *SortBy  `default:"timestamp" queryParam:"style=form,explode=true,name=sortBy"`
}

func (l ListEventsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListEventsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListEventsRequest) GetEvent() *QueryParamEvent {
	if o == nil {
		return nil
	}
	return o.Event
}

func (o *ListEventsRequest) GetDomain() *string {
	if o == nil {
		return nil
	}
	return o.Domain
}

func (o *ListEventsRequest) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *ListEventsRequest) GetLinkID() *string {
	if o == nil {
		return nil
	}
	return o.LinkID
}

func (o *ListEventsRequest) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

func (o *ListEventsRequest) GetInterval() *QueryParamInterval {
	if o == nil {
		return nil
	}
	return o.Interval
}

func (o *ListEventsRequest) GetStart() *string {
	if o == nil {
		return nil
	}
	return o.Start
}

func (o *ListEventsRequest) GetEnd() *string {
	if o == nil {
		return nil
	}
	return o.End
}

func (o *ListEventsRequest) GetTimezone() *string {
	if o == nil {
		return nil
	}
	return o.Timezone
}

func (o *ListEventsRequest) GetContinent() *components.ContinentCode {
	if o == nil {
		return nil
	}
	return o.Continent
}

func (o *ListEventsRequest) GetCountry() *components.CountryCode {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *ListEventsRequest) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *ListEventsRequest) GetDevice() *string {
	if o == nil {
		return nil
	}
	return o.Device
}

func (o *ListEventsRequest) GetBrowser() *string {
	if o == nil {
		return nil
	}
	return o.Browser
}

func (o *ListEventsRequest) GetOs() *string {
	if o == nil {
		return nil
	}
	return o.Os
}

func (o *ListEventsRequest) GetReferer() *string {
	if o == nil {
		return nil
	}
	return o.Referer
}

func (o *ListEventsRequest) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *ListEventsRequest) GetTagID() *string {
	if o == nil {
		return nil
	}
	return o.TagID
}

func (o *ListEventsRequest) GetQr() *bool {
	if o == nil {
		return nil
	}
	return o.Qr
}

func (o *ListEventsRequest) GetRoot() *bool {
	if o == nil {
		return nil
	}
	return o.Root
}

func (o *ListEventsRequest) GetPage() *float64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListEventsRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListEventsRequest) GetOrder() *Order {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListEventsRequest) GetSortBy() *SortBy {
	if o == nil {
		return nil
	}
	return o.SortBy
}
