// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/dubinc/dub-go/models/components"
)

type GetTopURLsByClicksGlobals struct {
	WorkspaceID string `queryParam:"style=form,explode=true,name=workspaceId"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	ProjectSlug *string `queryParam:"style=form,explode=true,name=projectSlug"`
}

func (o *GetTopURLsByClicksGlobals) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

func (o *GetTopURLsByClicksGlobals) GetProjectSlug() *string {
	if o == nil {
		return nil
	}
	return o.ProjectSlug
}

// GetTopURLsByClicksQueryParamInterval - The interval to retrieve analytics for. Takes precedence over start and end. If undefined, defaults to 24h.
type GetTopURLsByClicksQueryParamInterval string

const (
	GetTopURLsByClicksQueryParamIntervalOneh        GetTopURLsByClicksQueryParamInterval = "1h"
	GetTopURLsByClicksQueryParamIntervalTwentyFourh GetTopURLsByClicksQueryParamInterval = "24h"
	GetTopURLsByClicksQueryParamIntervalSevend      GetTopURLsByClicksQueryParamInterval = "7d"
	GetTopURLsByClicksQueryParamIntervalThirtyd     GetTopURLsByClicksQueryParamInterval = "30d"
	GetTopURLsByClicksQueryParamIntervalNinetyd     GetTopURLsByClicksQueryParamInterval = "90d"
	GetTopURLsByClicksQueryParamIntervalYtd         GetTopURLsByClicksQueryParamInterval = "ytd"
	GetTopURLsByClicksQueryParamIntervalOney        GetTopURLsByClicksQueryParamInterval = "1y"
	GetTopURLsByClicksQueryParamIntervalAll         GetTopURLsByClicksQueryParamInterval = "all"
)

func (e GetTopURLsByClicksQueryParamInterval) ToPointer() *GetTopURLsByClicksQueryParamInterval {
	return &e
}
func (e *GetTopURLsByClicksQueryParamInterval) UnmarshalJSON(data []byte) error {
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
		*e = GetTopURLsByClicksQueryParamInterval(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetTopURLsByClicksQueryParamInterval: %v", v)
	}
}

type GetTopURLsByClicksRequest struct {
	// The domain to filter analytics for.
	Domain *string `queryParam:"style=form,explode=true,name=domain"`
	// The short link slug.
	Key *string `queryParam:"style=form,explode=true,name=key"`
	// The unique ID of the short link on Dub.
	LinkID *string `queryParam:"style=form,explode=true,name=linkId"`
	// This is the ID of the link in the your database. Must be prefixed with 'ext_' when passed as a query parameter.
	ExternalID *string `queryParam:"style=form,explode=true,name=externalId"`
	// The interval to retrieve analytics for. Takes precedence over start and end. If undefined, defaults to 24h.
	Interval *GetTopURLsByClicksQueryParamInterval `queryParam:"style=form,explode=true,name=interval"`
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

func (o *GetTopURLsByClicksRequest) GetDomain() *string {
	if o == nil {
		return nil
	}
	return o.Domain
}

func (o *GetTopURLsByClicksRequest) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *GetTopURLsByClicksRequest) GetLinkID() *string {
	if o == nil {
		return nil
	}
	return o.LinkID
}

func (o *GetTopURLsByClicksRequest) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

func (o *GetTopURLsByClicksRequest) GetInterval() *GetTopURLsByClicksQueryParamInterval {
	if o == nil {
		return nil
	}
	return o.Interval
}

func (o *GetTopURLsByClicksRequest) GetStart() *string {
	if o == nil {
		return nil
	}
	return o.Start
}

func (o *GetTopURLsByClicksRequest) GetEnd() *string {
	if o == nil {
		return nil
	}
	return o.End
}

func (o *GetTopURLsByClicksRequest) GetCountry() *components.CountryCode {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *GetTopURLsByClicksRequest) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *GetTopURLsByClicksRequest) GetDevice() *string {
	if o == nil {
		return nil
	}
	return o.Device
}

func (o *GetTopURLsByClicksRequest) GetBrowser() *string {
	if o == nil {
		return nil
	}
	return o.Browser
}

func (o *GetTopURLsByClicksRequest) GetOs() *string {
	if o == nil {
		return nil
	}
	return o.Os
}

func (o *GetTopURLsByClicksRequest) GetReferer() *string {
	if o == nil {
		return nil
	}
	return o.Referer
}

func (o *GetTopURLsByClicksRequest) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *GetTopURLsByClicksRequest) GetTagID() *string {
	if o == nil {
		return nil
	}
	return o.TagID
}

func (o *GetTopURLsByClicksRequest) GetQr() *bool {
	if o == nil {
		return nil
	}
	return o.Qr
}

func (o *GetTopURLsByClicksRequest) GetRoot() *bool {
	if o == nil {
		return nil
	}
	return o.Root
}

type GetTopURLsByClicksResponseBody struct {
	// The destination URL
	URL string `json:"url"`
	// The number of clicks from this URL
	Clicks float64 `json:"clicks"`
}

func (o *GetTopURLsByClicksResponseBody) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

func (o *GetTopURLsByClicksResponseBody) GetClicks() float64 {
	if o == nil {
		return 0.0
	}
	return o.Clicks
}

type GetTopURLsByClicksResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The top URLs by number of clicks
	ResponseBodies []GetTopURLsByClicksResponseBody
}

func (o *GetTopURLsByClicksResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetTopURLsByClicksResponse) GetResponseBodies() []GetTopURLsByClicksResponseBody {
	if o == nil {
		return nil
	}
	return o.ResponseBodies
}
