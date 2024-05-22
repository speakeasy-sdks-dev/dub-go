// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/dubinc/dub-go/models/components"
)

type GetTimeseriesByClicksGlobals struct {
	WorkspaceID string `queryParam:"style=form,explode=true,name=workspaceId"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	ProjectSlug *string `queryParam:"style=form,explode=true,name=projectSlug"`
}

func (o *GetTimeseriesByClicksGlobals) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

func (o *GetTimeseriesByClicksGlobals) GetProjectSlug() *string {
	if o == nil {
		return nil
	}
	return o.ProjectSlug
}

// QueryParamInterval - The interval to retrieve analytics for. Takes precedence over start and end. If undefined, defaults to 24h.
type QueryParamInterval string

const (
	QueryParamIntervalOneh        QueryParamInterval = "1h"
	QueryParamIntervalTwentyFourh QueryParamInterval = "24h"
	QueryParamIntervalSevend      QueryParamInterval = "7d"
	QueryParamIntervalThirtyd     QueryParamInterval = "30d"
	QueryParamIntervalNinetyd     QueryParamInterval = "90d"
	QueryParamIntervalYtd         QueryParamInterval = "ytd"
	QueryParamIntervalOney        QueryParamInterval = "1y"
	QueryParamIntervalAll         QueryParamInterval = "all"
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
	case "1h":
		fallthrough
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
		*e = QueryParamInterval(v)
		return nil
	default:
		return fmt.Errorf("invalid value for QueryParamInterval: %v", v)
	}
}

type GetTimeseriesByClicksRequest struct {
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

func (o *GetTimeseriesByClicksRequest) GetDomain() *string {
	if o == nil {
		return nil
	}
	return o.Domain
}

func (o *GetTimeseriesByClicksRequest) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *GetTimeseriesByClicksRequest) GetLinkID() *string {
	if o == nil {
		return nil
	}
	return o.LinkID
}

func (o *GetTimeseriesByClicksRequest) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

func (o *GetTimeseriesByClicksRequest) GetInterval() *QueryParamInterval {
	if o == nil {
		return nil
	}
	return o.Interval
}

func (o *GetTimeseriesByClicksRequest) GetStart() *string {
	if o == nil {
		return nil
	}
	return o.Start
}

func (o *GetTimeseriesByClicksRequest) GetEnd() *string {
	if o == nil {
		return nil
	}
	return o.End
}

func (o *GetTimeseriesByClicksRequest) GetCountry() *components.CountryCode {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *GetTimeseriesByClicksRequest) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *GetTimeseriesByClicksRequest) GetDevice() *string {
	if o == nil {
		return nil
	}
	return o.Device
}

func (o *GetTimeseriesByClicksRequest) GetBrowser() *string {
	if o == nil {
		return nil
	}
	return o.Browser
}

func (o *GetTimeseriesByClicksRequest) GetOs() *string {
	if o == nil {
		return nil
	}
	return o.Os
}

func (o *GetTimeseriesByClicksRequest) GetReferer() *string {
	if o == nil {
		return nil
	}
	return o.Referer
}

func (o *GetTimeseriesByClicksRequest) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *GetTimeseriesByClicksRequest) GetTagID() *string {
	if o == nil {
		return nil
	}
	return o.TagID
}

func (o *GetTimeseriesByClicksRequest) GetQr() *bool {
	if o == nil {
		return nil
	}
	return o.Qr
}

func (o *GetTimeseriesByClicksRequest) GetRoot() *bool {
	if o == nil {
		return nil
	}
	return o.Root
}

type ResponseBody struct {
	// The starting timestamp of the interval
	Start string `json:"start"`
	// The number of clicks in the interval
	Clicks float64 `json:"clicks"`
}

func (o *ResponseBody) GetStart() string {
	if o == nil {
		return ""
	}
	return o.Start
}

func (o *ResponseBody) GetClicks() float64 {
	if o == nil {
		return 0.0
	}
	return o.Clicks
}

type GetTimeseriesByClicksResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The number of clicks over a period of time
	ResponseBodies []ResponseBody
}

func (o *GetTimeseriesByClicksResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetTimeseriesByClicksResponse) GetResponseBodies() []ResponseBody {
	if o == nil {
		return nil
	}
	return o.ResponseBodies
}
