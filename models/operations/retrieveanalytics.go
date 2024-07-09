// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dubinc/dub-go/internal/utils"
	"github.com/dubinc/dub-go/models/components"
)

// Event - The type of event to retrieve analytics for. Defaults to 'clicks'.
type Event string

const (
	EventClicks    Event = "clicks"
	EventLeads     Event = "leads"
	EventSales     Event = "sales"
	EventComposite Event = "composite"
)

func (e Event) ToPointer() *Event {
	return &e
}
func (e *Event) UnmarshalJSON(data []byte) error {
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
		*e = Event(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Event: %v", v)
	}
}

// QueryParamGroupBy - The parameter to group the analytics data points by. Defaults to 'count' if undefined.
type QueryParamGroupBy string

const (
	QueryParamGroupByCount      QueryParamGroupBy = "count"
	QueryParamGroupByTimeseries QueryParamGroupBy = "timeseries"
	QueryParamGroupByCountries  QueryParamGroupBy = "countries"
	QueryParamGroupByCities     QueryParamGroupBy = "cities"
	QueryParamGroupByDevices    QueryParamGroupBy = "devices"
	QueryParamGroupByBrowsers   QueryParamGroupBy = "browsers"
	QueryParamGroupByOs         QueryParamGroupBy = "os"
	QueryParamGroupByReferers   QueryParamGroupBy = "referers"
	QueryParamGroupByTopLinks   QueryParamGroupBy = "top_links"
	QueryParamGroupByTopUrls    QueryParamGroupBy = "top_urls"
	QueryParamGroupByTrigger    QueryParamGroupBy = "trigger"
)

func (e QueryParamGroupBy) ToPointer() *QueryParamGroupBy {
	return &e
}
func (e *QueryParamGroupBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "count":
		fallthrough
	case "timeseries":
		fallthrough
	case "countries":
		fallthrough
	case "cities":
		fallthrough
	case "devices":
		fallthrough
	case "browsers":
		fallthrough
	case "os":
		fallthrough
	case "referers":
		fallthrough
	case "top_links":
		fallthrough
	case "top_urls":
		fallthrough
	case "trigger":
		*e = QueryParamGroupBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for QueryParamGroupBy: %v", v)
	}
}

// Interval - The interval to retrieve analytics for. Takes precedence over start and end. If undefined, defaults to 24h.
type Interval string

const (
	IntervalTwentyFourh   Interval = "24h"
	IntervalSevend        Interval = "7d"
	IntervalThirtyd       Interval = "30d"
	IntervalNinetyd       Interval = "90d"
	IntervalYtd           Interval = "ytd"
	IntervalOney          Interval = "1y"
	IntervalAll           Interval = "all"
	IntervalAllUnfiltered Interval = "all_unfiltered"
)

func (e Interval) ToPointer() *Interval {
	return &e
}
func (e *Interval) UnmarshalJSON(data []byte) error {
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
		*e = Interval(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Interval: %v", v)
	}
}

type RetrieveAnalyticsRequest struct {
	// The type of event to retrieve analytics for. Defaults to 'clicks'.
	Event *Event `default:"clicks" queryParam:"style=form,explode=true,name=event"`
	// The parameter to group the analytics data points by. Defaults to 'count' if undefined.
	GroupBy *QueryParamGroupBy `default:"count" queryParam:"style=form,explode=true,name=groupBy"`
	// The domain to filter analytics for.
	Domain *string `queryParam:"style=form,explode=true,name=domain"`
	// The short link slug.
	Key *string `queryParam:"style=form,explode=true,name=key"`
	// The unique ID of the short link on Dub.
	LinkID *string `queryParam:"style=form,explode=true,name=linkId"`
	// This is the ID of the link in the your database. Must be prefixed with 'ext_' when passed as a query parameter.
	ExternalID *string `queryParam:"style=form,explode=true,name=externalId"`
	// The interval to retrieve analytics for. Takes precedence over start and end. If undefined, defaults to 24h.
	Interval *Interval `queryParam:"style=form,explode=true,name=interval"`
	// The start date and time when to retrieve analytics from.
	Start *string `queryParam:"style=form,explode=true,name=start"`
	// The end date and time when to retrieve analytics from. If not provided, defaults to the current date.
	End *string `queryParam:"style=form,explode=true,name=end"`
	// The IANA time zone code for aligning timeseries granularity (e.g. America/New_York). Defaults to UTC.
	Timezone *string `default:"UTC" queryParam:"style=form,explode=true,name=timezone"`
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
	Root *bool `queryParam:"style=form,explode=true,name=root"`
}

func (r RetrieveAnalyticsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RetrieveAnalyticsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RetrieveAnalyticsRequest) GetEvent() *Event {
	if o == nil {
		return nil
	}
	return o.Event
}

func (o *RetrieveAnalyticsRequest) GetGroupBy() *QueryParamGroupBy {
	if o == nil {
		return nil
	}
	return o.GroupBy
}

func (o *RetrieveAnalyticsRequest) GetDomain() *string {
	if o == nil {
		return nil
	}
	return o.Domain
}

func (o *RetrieveAnalyticsRequest) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *RetrieveAnalyticsRequest) GetLinkID() *string {
	if o == nil {
		return nil
	}
	return o.LinkID
}

func (o *RetrieveAnalyticsRequest) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

func (o *RetrieveAnalyticsRequest) GetInterval() *Interval {
	if o == nil {
		return nil
	}
	return o.Interval
}

func (o *RetrieveAnalyticsRequest) GetStart() *string {
	if o == nil {
		return nil
	}
	return o.Start
}

func (o *RetrieveAnalyticsRequest) GetEnd() *string {
	if o == nil {
		return nil
	}
	return o.End
}

func (o *RetrieveAnalyticsRequest) GetTimezone() *string {
	if o == nil {
		return nil
	}
	return o.Timezone
}

func (o *RetrieveAnalyticsRequest) GetCountry() *components.CountryCode {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *RetrieveAnalyticsRequest) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *RetrieveAnalyticsRequest) GetDevice() *string {
	if o == nil {
		return nil
	}
	return o.Device
}

func (o *RetrieveAnalyticsRequest) GetBrowser() *string {
	if o == nil {
		return nil
	}
	return o.Browser
}

func (o *RetrieveAnalyticsRequest) GetOs() *string {
	if o == nil {
		return nil
	}
	return o.Os
}

func (o *RetrieveAnalyticsRequest) GetReferer() *string {
	if o == nil {
		return nil
	}
	return o.Referer
}

func (o *RetrieveAnalyticsRequest) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *RetrieveAnalyticsRequest) GetTagID() *string {
	if o == nil {
		return nil
	}
	return o.TagID
}

func (o *RetrieveAnalyticsRequest) GetQr() *bool {
	if o == nil {
		return nil
	}
	return o.Qr
}

func (o *RetrieveAnalyticsRequest) GetRoot() *bool {
	if o == nil {
		return nil
	}
	return o.Root
}

type RetrieveAnalyticsResponseBodyType string

const (
	RetrieveAnalyticsResponseBodyTypeClicksCount             RetrieveAnalyticsResponseBodyType = "ClicksCount"
	RetrieveAnalyticsResponseBodyTypeArrayOfClicksTimeseries RetrieveAnalyticsResponseBodyType = "arrayOfClicksTimeseries"
	RetrieveAnalyticsResponseBodyTypeArrayOfClicksCountries  RetrieveAnalyticsResponseBodyType = "arrayOfClicksCountries"
	RetrieveAnalyticsResponseBodyTypeArrayOfClicksCities     RetrieveAnalyticsResponseBodyType = "arrayOfClicksCities"
	RetrieveAnalyticsResponseBodyTypeArrayOfClicksDevices    RetrieveAnalyticsResponseBodyType = "arrayOfClicksDevices"
	RetrieveAnalyticsResponseBodyTypeArrayOfClicksBrowsers   RetrieveAnalyticsResponseBodyType = "arrayOfClicksBrowsers"
	RetrieveAnalyticsResponseBodyTypeArrayOfClicksOS         RetrieveAnalyticsResponseBodyType = "arrayOfClicksOS"
	RetrieveAnalyticsResponseBodyTypeArrayOfClicksReferers   RetrieveAnalyticsResponseBodyType = "arrayOfClicksReferers"
	RetrieveAnalyticsResponseBodyTypeArrayOfClicksTopLinks   RetrieveAnalyticsResponseBodyType = "arrayOfClicksTopLinks"
	RetrieveAnalyticsResponseBodyTypeArrayOfClicksTopUrls    RetrieveAnalyticsResponseBodyType = "arrayOfClicksTopUrls"
	RetrieveAnalyticsResponseBodyTypeLeadsCount              RetrieveAnalyticsResponseBodyType = "LeadsCount"
	RetrieveAnalyticsResponseBodyTypeArrayOfLeadsTimeseries  RetrieveAnalyticsResponseBodyType = "arrayOfLeadsTimeseries"
	RetrieveAnalyticsResponseBodyTypeArrayOfLeadsCountries   RetrieveAnalyticsResponseBodyType = "arrayOfLeadsCountries"
	RetrieveAnalyticsResponseBodyTypeArrayOfLeadsCities      RetrieveAnalyticsResponseBodyType = "arrayOfLeadsCities"
	RetrieveAnalyticsResponseBodyTypeArrayOfLeadsDevices     RetrieveAnalyticsResponseBodyType = "arrayOfLeadsDevices"
	RetrieveAnalyticsResponseBodyTypeArrayOfLeadsBrowsers    RetrieveAnalyticsResponseBodyType = "arrayOfLeadsBrowsers"
	RetrieveAnalyticsResponseBodyTypeArrayOfLeadsOS          RetrieveAnalyticsResponseBodyType = "arrayOfLeadsOS"
	RetrieveAnalyticsResponseBodyTypeArrayOfLeadsReferers    RetrieveAnalyticsResponseBodyType = "arrayOfLeadsReferers"
	RetrieveAnalyticsResponseBodyTypeArrayOfLeadsTopLinks    RetrieveAnalyticsResponseBodyType = "arrayOfLeadsTopLinks"
	RetrieveAnalyticsResponseBodyTypeArrayOfLeadsTopUrls     RetrieveAnalyticsResponseBodyType = "arrayOfLeadsTopUrls"
	RetrieveAnalyticsResponseBodyTypeSalesCount              RetrieveAnalyticsResponseBodyType = "SalesCount"
	RetrieveAnalyticsResponseBodyTypeArrayOfSalesTimeseries  RetrieveAnalyticsResponseBodyType = "arrayOfSalesTimeseries"
	RetrieveAnalyticsResponseBodyTypeArrayOfSalesCountries   RetrieveAnalyticsResponseBodyType = "arrayOfSalesCountries"
	RetrieveAnalyticsResponseBodyTypeArrayOfSalesCities      RetrieveAnalyticsResponseBodyType = "arrayOfSalesCities"
	RetrieveAnalyticsResponseBodyTypeArrayOfSalesDevices     RetrieveAnalyticsResponseBodyType = "arrayOfSalesDevices"
	RetrieveAnalyticsResponseBodyTypeArrayOfSalesBrowsers    RetrieveAnalyticsResponseBodyType = "arrayOfSalesBrowsers"
	RetrieveAnalyticsResponseBodyTypeArrayOfSalesOS          RetrieveAnalyticsResponseBodyType = "arrayOfSalesOS"
	RetrieveAnalyticsResponseBodyTypeArrayOfSalesReferers    RetrieveAnalyticsResponseBodyType = "arrayOfSalesReferers"
	RetrieveAnalyticsResponseBodyTypeArrayOfSalesTopLinks    RetrieveAnalyticsResponseBodyType = "arrayOfSalesTopLinks"
	RetrieveAnalyticsResponseBodyTypeArrayOfSalesTopUrls     RetrieveAnalyticsResponseBodyType = "arrayOfSalesTopUrls"
)

// RetrieveAnalyticsResponseBody - Analytics data
type RetrieveAnalyticsResponseBody struct {
	ClicksCount             *components.ClicksCount
	ArrayOfClicksTimeseries []components.ClicksTimeseries
	ArrayOfClicksCountries  []components.ClicksCountries
	ArrayOfClicksCities     []components.ClicksCities
	ArrayOfClicksDevices    []components.ClicksDevices
	ArrayOfClicksBrowsers   []components.ClicksBrowsers
	ArrayOfClicksOS         []components.ClicksOS
	ArrayOfClicksReferers   []components.ClicksReferers
	ArrayOfClicksTopLinks   []components.ClicksTopLinks
	ArrayOfClicksTopUrls    []components.ClicksTopUrls
	LeadsCount              *components.LeadsCount
	ArrayOfLeadsTimeseries  []components.LeadsTimeseries
	ArrayOfLeadsCountries   []components.LeadsCountries
	ArrayOfLeadsCities      []components.LeadsCities
	ArrayOfLeadsDevices     []components.LeadsDevices
	ArrayOfLeadsBrowsers    []components.LeadsBrowsers
	ArrayOfLeadsOS          []components.LeadsOS
	ArrayOfLeadsReferers    []components.LeadsReferers
	ArrayOfLeadsTopLinks    []components.LeadsTopLinks
	ArrayOfLeadsTopUrls     []components.LeadsTopUrls
	SalesCount              *components.SalesCount
	ArrayOfSalesTimeseries  []components.SalesTimeseries
	ArrayOfSalesCountries   []components.SalesCountries
	ArrayOfSalesCities      []components.SalesCities
	ArrayOfSalesDevices     []components.SalesDevices
	ArrayOfSalesBrowsers    []components.SalesBrowsers
	ArrayOfSalesOS          []components.SalesOS
	ArrayOfSalesReferers    []components.SalesReferers
	ArrayOfSalesTopLinks    []components.SalesTopLinks
	ArrayOfSalesTopUrls     []components.SalesTopUrls

	Type RetrieveAnalyticsResponseBodyType
}

func CreateRetrieveAnalyticsResponseBodyClicksCount(clicksCount components.ClicksCount) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeClicksCount

	return RetrieveAnalyticsResponseBody{
		ClicksCount: &clicksCount,
		Type:        typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfClicksTimeseries(arrayOfClicksTimeseries []components.ClicksTimeseries) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfClicksTimeseries

	return RetrieveAnalyticsResponseBody{
		ArrayOfClicksTimeseries: arrayOfClicksTimeseries,
		Type:                    typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfClicksCountries(arrayOfClicksCountries []components.ClicksCountries) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfClicksCountries

	return RetrieveAnalyticsResponseBody{
		ArrayOfClicksCountries: arrayOfClicksCountries,
		Type:                   typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfClicksCities(arrayOfClicksCities []components.ClicksCities) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfClicksCities

	return RetrieveAnalyticsResponseBody{
		ArrayOfClicksCities: arrayOfClicksCities,
		Type:                typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfClicksDevices(arrayOfClicksDevices []components.ClicksDevices) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfClicksDevices

	return RetrieveAnalyticsResponseBody{
		ArrayOfClicksDevices: arrayOfClicksDevices,
		Type:                 typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfClicksBrowsers(arrayOfClicksBrowsers []components.ClicksBrowsers) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfClicksBrowsers

	return RetrieveAnalyticsResponseBody{
		ArrayOfClicksBrowsers: arrayOfClicksBrowsers,
		Type:                  typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfClicksOS(arrayOfClicksOS []components.ClicksOS) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfClicksOS

	return RetrieveAnalyticsResponseBody{
		ArrayOfClicksOS: arrayOfClicksOS,
		Type:            typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfClicksReferers(arrayOfClicksReferers []components.ClicksReferers) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfClicksReferers

	return RetrieveAnalyticsResponseBody{
		ArrayOfClicksReferers: arrayOfClicksReferers,
		Type:                  typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfClicksTopLinks(arrayOfClicksTopLinks []components.ClicksTopLinks) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfClicksTopLinks

	return RetrieveAnalyticsResponseBody{
		ArrayOfClicksTopLinks: arrayOfClicksTopLinks,
		Type:                  typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfClicksTopUrls(arrayOfClicksTopUrls []components.ClicksTopUrls) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfClicksTopUrls

	return RetrieveAnalyticsResponseBody{
		ArrayOfClicksTopUrls: arrayOfClicksTopUrls,
		Type:                 typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyLeadsCount(leadsCount components.LeadsCount) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeLeadsCount

	return RetrieveAnalyticsResponseBody{
		LeadsCount: &leadsCount,
		Type:       typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfLeadsTimeseries(arrayOfLeadsTimeseries []components.LeadsTimeseries) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfLeadsTimeseries

	return RetrieveAnalyticsResponseBody{
		ArrayOfLeadsTimeseries: arrayOfLeadsTimeseries,
		Type:                   typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfLeadsCountries(arrayOfLeadsCountries []components.LeadsCountries) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfLeadsCountries

	return RetrieveAnalyticsResponseBody{
		ArrayOfLeadsCountries: arrayOfLeadsCountries,
		Type:                  typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfLeadsCities(arrayOfLeadsCities []components.LeadsCities) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfLeadsCities

	return RetrieveAnalyticsResponseBody{
		ArrayOfLeadsCities: arrayOfLeadsCities,
		Type:               typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfLeadsDevices(arrayOfLeadsDevices []components.LeadsDevices) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfLeadsDevices

	return RetrieveAnalyticsResponseBody{
		ArrayOfLeadsDevices: arrayOfLeadsDevices,
		Type:                typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfLeadsBrowsers(arrayOfLeadsBrowsers []components.LeadsBrowsers) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfLeadsBrowsers

	return RetrieveAnalyticsResponseBody{
		ArrayOfLeadsBrowsers: arrayOfLeadsBrowsers,
		Type:                 typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfLeadsOS(arrayOfLeadsOS []components.LeadsOS) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfLeadsOS

	return RetrieveAnalyticsResponseBody{
		ArrayOfLeadsOS: arrayOfLeadsOS,
		Type:           typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfLeadsReferers(arrayOfLeadsReferers []components.LeadsReferers) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfLeadsReferers

	return RetrieveAnalyticsResponseBody{
		ArrayOfLeadsReferers: arrayOfLeadsReferers,
		Type:                 typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfLeadsTopLinks(arrayOfLeadsTopLinks []components.LeadsTopLinks) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfLeadsTopLinks

	return RetrieveAnalyticsResponseBody{
		ArrayOfLeadsTopLinks: arrayOfLeadsTopLinks,
		Type:                 typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfLeadsTopUrls(arrayOfLeadsTopUrls []components.LeadsTopUrls) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfLeadsTopUrls

	return RetrieveAnalyticsResponseBody{
		ArrayOfLeadsTopUrls: arrayOfLeadsTopUrls,
		Type:                typ,
	}
}

func CreateRetrieveAnalyticsResponseBodySalesCount(salesCount components.SalesCount) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeSalesCount

	return RetrieveAnalyticsResponseBody{
		SalesCount: &salesCount,
		Type:       typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfSalesTimeseries(arrayOfSalesTimeseries []components.SalesTimeseries) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfSalesTimeseries

	return RetrieveAnalyticsResponseBody{
		ArrayOfSalesTimeseries: arrayOfSalesTimeseries,
		Type:                   typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfSalesCountries(arrayOfSalesCountries []components.SalesCountries) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfSalesCountries

	return RetrieveAnalyticsResponseBody{
		ArrayOfSalesCountries: arrayOfSalesCountries,
		Type:                  typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfSalesCities(arrayOfSalesCities []components.SalesCities) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfSalesCities

	return RetrieveAnalyticsResponseBody{
		ArrayOfSalesCities: arrayOfSalesCities,
		Type:               typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfSalesDevices(arrayOfSalesDevices []components.SalesDevices) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfSalesDevices

	return RetrieveAnalyticsResponseBody{
		ArrayOfSalesDevices: arrayOfSalesDevices,
		Type:                typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfSalesBrowsers(arrayOfSalesBrowsers []components.SalesBrowsers) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfSalesBrowsers

	return RetrieveAnalyticsResponseBody{
		ArrayOfSalesBrowsers: arrayOfSalesBrowsers,
		Type:                 typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfSalesOS(arrayOfSalesOS []components.SalesOS) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfSalesOS

	return RetrieveAnalyticsResponseBody{
		ArrayOfSalesOS: arrayOfSalesOS,
		Type:           typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfSalesReferers(arrayOfSalesReferers []components.SalesReferers) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfSalesReferers

	return RetrieveAnalyticsResponseBody{
		ArrayOfSalesReferers: arrayOfSalesReferers,
		Type:                 typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfSalesTopLinks(arrayOfSalesTopLinks []components.SalesTopLinks) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfSalesTopLinks

	return RetrieveAnalyticsResponseBody{
		ArrayOfSalesTopLinks: arrayOfSalesTopLinks,
		Type:                 typ,
	}
}

func CreateRetrieveAnalyticsResponseBodyArrayOfSalesTopUrls(arrayOfSalesTopUrls []components.SalesTopUrls) RetrieveAnalyticsResponseBody {
	typ := RetrieveAnalyticsResponseBodyTypeArrayOfSalesTopUrls

	return RetrieveAnalyticsResponseBody{
		ArrayOfSalesTopUrls: arrayOfSalesTopUrls,
		Type:                typ,
	}
}

func (u *RetrieveAnalyticsResponseBody) UnmarshalJSON(data []byte) error {

	var leadsCount components.LeadsCount = components.LeadsCount{}
	if err := utils.UnmarshalJSON(data, &leadsCount, "", true, true); err == nil {
		u.LeadsCount = &leadsCount
		u.Type = RetrieveAnalyticsResponseBodyTypeLeadsCount
		return nil
	}

	var clicksCount components.ClicksCount = components.ClicksCount{}
	if err := utils.UnmarshalJSON(data, &clicksCount, "", true, true); err == nil {
		u.ClicksCount = &clicksCount
		u.Type = RetrieveAnalyticsResponseBodyTypeClicksCount
		return nil
	}

	var salesCount components.SalesCount = components.SalesCount{}
	if err := utils.UnmarshalJSON(data, &salesCount, "", true, true); err == nil {
		u.SalesCount = &salesCount
		u.Type = RetrieveAnalyticsResponseBodyTypeSalesCount
		return nil
	}

	var arrayOfLeadsDevices []components.LeadsDevices = []components.LeadsDevices{}
	if err := utils.UnmarshalJSON(data, &arrayOfLeadsDevices, "", true, true); err == nil {
		u.ArrayOfLeadsDevices = arrayOfLeadsDevices
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfLeadsDevices
		return nil
	}

	var arrayOfLeadsOS []components.LeadsOS = []components.LeadsOS{}
	if err := utils.UnmarshalJSON(data, &arrayOfLeadsOS, "", true, true); err == nil {
		u.ArrayOfLeadsOS = arrayOfLeadsOS
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfLeadsOS
		return nil
	}

	var arrayOfClicksBrowsers []components.ClicksBrowsers = []components.ClicksBrowsers{}
	if err := utils.UnmarshalJSON(data, &arrayOfClicksBrowsers, "", true, true); err == nil {
		u.ArrayOfClicksBrowsers = arrayOfClicksBrowsers
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfClicksBrowsers
		return nil
	}

	var arrayOfClicksOS []components.ClicksOS = []components.ClicksOS{}
	if err := utils.UnmarshalJSON(data, &arrayOfClicksOS, "", true, true); err == nil {
		u.ArrayOfClicksOS = arrayOfClicksOS
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfClicksOS
		return nil
	}

	var arrayOfClicksReferers []components.ClicksReferers = []components.ClicksReferers{}
	if err := utils.UnmarshalJSON(data, &arrayOfClicksReferers, "", true, true); err == nil {
		u.ArrayOfClicksReferers = arrayOfClicksReferers
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfClicksReferers
		return nil
	}

	var arrayOfClicksTopLinks []components.ClicksTopLinks = []components.ClicksTopLinks{}
	if err := utils.UnmarshalJSON(data, &arrayOfClicksTopLinks, "", true, true); err == nil {
		u.ArrayOfClicksTopLinks = arrayOfClicksTopLinks
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfClicksTopLinks
		return nil
	}

	var arrayOfClicksTopUrls []components.ClicksTopUrls = []components.ClicksTopUrls{}
	if err := utils.UnmarshalJSON(data, &arrayOfClicksTopUrls, "", true, true); err == nil {
		u.ArrayOfClicksTopUrls = arrayOfClicksTopUrls
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfClicksTopUrls
		return nil
	}

	var arrayOfClicksCities []components.ClicksCities = []components.ClicksCities{}
	if err := utils.UnmarshalJSON(data, &arrayOfClicksCities, "", true, true); err == nil {
		u.ArrayOfClicksCities = arrayOfClicksCities
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfClicksCities
		return nil
	}

	var arrayOfLeadsTimeseries []components.LeadsTimeseries = []components.LeadsTimeseries{}
	if err := utils.UnmarshalJSON(data, &arrayOfLeadsTimeseries, "", true, true); err == nil {
		u.ArrayOfLeadsTimeseries = arrayOfLeadsTimeseries
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfLeadsTimeseries
		return nil
	}

	var arrayOfLeadsCountries []components.LeadsCountries = []components.LeadsCountries{}
	if err := utils.UnmarshalJSON(data, &arrayOfLeadsCountries, "", true, true); err == nil {
		u.ArrayOfLeadsCountries = arrayOfLeadsCountries
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfLeadsCountries
		return nil
	}

	var arrayOfLeadsCities []components.LeadsCities = []components.LeadsCities{}
	if err := utils.UnmarshalJSON(data, &arrayOfLeadsCities, "", true, true); err == nil {
		u.ArrayOfLeadsCities = arrayOfLeadsCities
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfLeadsCities
		return nil
	}

	var arrayOfClicksCountries []components.ClicksCountries = []components.ClicksCountries{}
	if err := utils.UnmarshalJSON(data, &arrayOfClicksCountries, "", true, true); err == nil {
		u.ArrayOfClicksCountries = arrayOfClicksCountries
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfClicksCountries
		return nil
	}

	var arrayOfLeadsBrowsers []components.LeadsBrowsers = []components.LeadsBrowsers{}
	if err := utils.UnmarshalJSON(data, &arrayOfLeadsBrowsers, "", true, true); err == nil {
		u.ArrayOfLeadsBrowsers = arrayOfLeadsBrowsers
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfLeadsBrowsers
		return nil
	}

	var arrayOfClicksDevices []components.ClicksDevices = []components.ClicksDevices{}
	if err := utils.UnmarshalJSON(data, &arrayOfClicksDevices, "", true, true); err == nil {
		u.ArrayOfClicksDevices = arrayOfClicksDevices
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfClicksDevices
		return nil
	}

	var arrayOfLeadsReferers []components.LeadsReferers = []components.LeadsReferers{}
	if err := utils.UnmarshalJSON(data, &arrayOfLeadsReferers, "", true, true); err == nil {
		u.ArrayOfLeadsReferers = arrayOfLeadsReferers
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfLeadsReferers
		return nil
	}

	var arrayOfLeadsTopLinks []components.LeadsTopLinks = []components.LeadsTopLinks{}
	if err := utils.UnmarshalJSON(data, &arrayOfLeadsTopLinks, "", true, true); err == nil {
		u.ArrayOfLeadsTopLinks = arrayOfLeadsTopLinks
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfLeadsTopLinks
		return nil
	}

	var arrayOfLeadsTopUrls []components.LeadsTopUrls = []components.LeadsTopUrls{}
	if err := utils.UnmarshalJSON(data, &arrayOfLeadsTopUrls, "", true, true); err == nil {
		u.ArrayOfLeadsTopUrls = arrayOfLeadsTopUrls
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfLeadsTopUrls
		return nil
	}

	var arrayOfClicksTimeseries []components.ClicksTimeseries = []components.ClicksTimeseries{}
	if err := utils.UnmarshalJSON(data, &arrayOfClicksTimeseries, "", true, true); err == nil {
		u.ArrayOfClicksTimeseries = arrayOfClicksTimeseries
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfClicksTimeseries
		return nil
	}

	var arrayOfSalesTimeseries []components.SalesTimeseries = []components.SalesTimeseries{}
	if err := utils.UnmarshalJSON(data, &arrayOfSalesTimeseries, "", true, true); err == nil {
		u.ArrayOfSalesTimeseries = arrayOfSalesTimeseries
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfSalesTimeseries
		return nil
	}

	var arrayOfSalesCountries []components.SalesCountries = []components.SalesCountries{}
	if err := utils.UnmarshalJSON(data, &arrayOfSalesCountries, "", true, true); err == nil {
		u.ArrayOfSalesCountries = arrayOfSalesCountries
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfSalesCountries
		return nil
	}

	var arrayOfSalesCities []components.SalesCities = []components.SalesCities{}
	if err := utils.UnmarshalJSON(data, &arrayOfSalesCities, "", true, true); err == nil {
		u.ArrayOfSalesCities = arrayOfSalesCities
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfSalesCities
		return nil
	}

	var arrayOfSalesDevices []components.SalesDevices = []components.SalesDevices{}
	if err := utils.UnmarshalJSON(data, &arrayOfSalesDevices, "", true, true); err == nil {
		u.ArrayOfSalesDevices = arrayOfSalesDevices
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfSalesDevices
		return nil
	}

	var arrayOfSalesBrowsers []components.SalesBrowsers = []components.SalesBrowsers{}
	if err := utils.UnmarshalJSON(data, &arrayOfSalesBrowsers, "", true, true); err == nil {
		u.ArrayOfSalesBrowsers = arrayOfSalesBrowsers
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfSalesBrowsers
		return nil
	}

	var arrayOfSalesOS []components.SalesOS = []components.SalesOS{}
	if err := utils.UnmarshalJSON(data, &arrayOfSalesOS, "", true, true); err == nil {
		u.ArrayOfSalesOS = arrayOfSalesOS
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfSalesOS
		return nil
	}

	var arrayOfSalesReferers []components.SalesReferers = []components.SalesReferers{}
	if err := utils.UnmarshalJSON(data, &arrayOfSalesReferers, "", true, true); err == nil {
		u.ArrayOfSalesReferers = arrayOfSalesReferers
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfSalesReferers
		return nil
	}

	var arrayOfSalesTopLinks []components.SalesTopLinks = []components.SalesTopLinks{}
	if err := utils.UnmarshalJSON(data, &arrayOfSalesTopLinks, "", true, true); err == nil {
		u.ArrayOfSalesTopLinks = arrayOfSalesTopLinks
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfSalesTopLinks
		return nil
	}

	var arrayOfSalesTopUrls []components.SalesTopUrls = []components.SalesTopUrls{}
	if err := utils.UnmarshalJSON(data, &arrayOfSalesTopUrls, "", true, true); err == nil {
		u.ArrayOfSalesTopUrls = arrayOfSalesTopUrls
		u.Type = RetrieveAnalyticsResponseBodyTypeArrayOfSalesTopUrls
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for RetrieveAnalyticsResponseBody", string(data))
}

func (u RetrieveAnalyticsResponseBody) MarshalJSON() ([]byte, error) {
	if u.ClicksCount != nil {
		return utils.MarshalJSON(u.ClicksCount, "", true)
	}

	if u.ArrayOfClicksTimeseries != nil {
		return utils.MarshalJSON(u.ArrayOfClicksTimeseries, "", true)
	}

	if u.ArrayOfClicksCountries != nil {
		return utils.MarshalJSON(u.ArrayOfClicksCountries, "", true)
	}

	if u.ArrayOfClicksCities != nil {
		return utils.MarshalJSON(u.ArrayOfClicksCities, "", true)
	}

	if u.ArrayOfClicksDevices != nil {
		return utils.MarshalJSON(u.ArrayOfClicksDevices, "", true)
	}

	if u.ArrayOfClicksBrowsers != nil {
		return utils.MarshalJSON(u.ArrayOfClicksBrowsers, "", true)
	}

	if u.ArrayOfClicksOS != nil {
		return utils.MarshalJSON(u.ArrayOfClicksOS, "", true)
	}

	if u.ArrayOfClicksReferers != nil {
		return utils.MarshalJSON(u.ArrayOfClicksReferers, "", true)
	}

	if u.ArrayOfClicksTopLinks != nil {
		return utils.MarshalJSON(u.ArrayOfClicksTopLinks, "", true)
	}

	if u.ArrayOfClicksTopUrls != nil {
		return utils.MarshalJSON(u.ArrayOfClicksTopUrls, "", true)
	}

	if u.LeadsCount != nil {
		return utils.MarshalJSON(u.LeadsCount, "", true)
	}

	if u.ArrayOfLeadsTimeseries != nil {
		return utils.MarshalJSON(u.ArrayOfLeadsTimeseries, "", true)
	}

	if u.ArrayOfLeadsCountries != nil {
		return utils.MarshalJSON(u.ArrayOfLeadsCountries, "", true)
	}

	if u.ArrayOfLeadsCities != nil {
		return utils.MarshalJSON(u.ArrayOfLeadsCities, "", true)
	}

	if u.ArrayOfLeadsDevices != nil {
		return utils.MarshalJSON(u.ArrayOfLeadsDevices, "", true)
	}

	if u.ArrayOfLeadsBrowsers != nil {
		return utils.MarshalJSON(u.ArrayOfLeadsBrowsers, "", true)
	}

	if u.ArrayOfLeadsOS != nil {
		return utils.MarshalJSON(u.ArrayOfLeadsOS, "", true)
	}

	if u.ArrayOfLeadsReferers != nil {
		return utils.MarshalJSON(u.ArrayOfLeadsReferers, "", true)
	}

	if u.ArrayOfLeadsTopLinks != nil {
		return utils.MarshalJSON(u.ArrayOfLeadsTopLinks, "", true)
	}

	if u.ArrayOfLeadsTopUrls != nil {
		return utils.MarshalJSON(u.ArrayOfLeadsTopUrls, "", true)
	}

	if u.SalesCount != nil {
		return utils.MarshalJSON(u.SalesCount, "", true)
	}

	if u.ArrayOfSalesTimeseries != nil {
		return utils.MarshalJSON(u.ArrayOfSalesTimeseries, "", true)
	}

	if u.ArrayOfSalesCountries != nil {
		return utils.MarshalJSON(u.ArrayOfSalesCountries, "", true)
	}

	if u.ArrayOfSalesCities != nil {
		return utils.MarshalJSON(u.ArrayOfSalesCities, "", true)
	}

	if u.ArrayOfSalesDevices != nil {
		return utils.MarshalJSON(u.ArrayOfSalesDevices, "", true)
	}

	if u.ArrayOfSalesBrowsers != nil {
		return utils.MarshalJSON(u.ArrayOfSalesBrowsers, "", true)
	}

	if u.ArrayOfSalesOS != nil {
		return utils.MarshalJSON(u.ArrayOfSalesOS, "", true)
	}

	if u.ArrayOfSalesReferers != nil {
		return utils.MarshalJSON(u.ArrayOfSalesReferers, "", true)
	}

	if u.ArrayOfSalesTopLinks != nil {
		return utils.MarshalJSON(u.ArrayOfSalesTopLinks, "", true)
	}

	if u.ArrayOfSalesTopUrls != nil {
		return utils.MarshalJSON(u.ArrayOfSalesTopUrls, "", true)
	}

	return nil, errors.New("could not marshal union type RetrieveAnalyticsResponseBody: all fields are null")
}
