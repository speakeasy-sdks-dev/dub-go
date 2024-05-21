// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/dubinc/dub-go/internal/utils"
	"github.com/dubinc/dub-go/models/components"
)

type GetOSByClicksGlobals struct {
	WorkspaceID string `queryParam:"style=form,explode=true,name=workspaceId"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	ProjectSlug *string `queryParam:"style=form,explode=true,name=projectSlug"`
}

func (o *GetOSByClicksGlobals) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

func (o *GetOSByClicksGlobals) GetProjectSlug() *string {
	if o == nil {
		return nil
	}
	return o.ProjectSlug
}

// GetOSByClicksQueryParamInterval - The interval to retrieve analytics for. Takes precedence over start and end. If undefined, defaults to 24h.
type GetOSByClicksQueryParamInterval string

const (
	GetOSByClicksQueryParamIntervalOneh        GetOSByClicksQueryParamInterval = "1h"
	GetOSByClicksQueryParamIntervalTwentyFourh GetOSByClicksQueryParamInterval = "24h"
	GetOSByClicksQueryParamIntervalSevend      GetOSByClicksQueryParamInterval = "7d"
	GetOSByClicksQueryParamIntervalThirtyd     GetOSByClicksQueryParamInterval = "30d"
	GetOSByClicksQueryParamIntervalNinetyd     GetOSByClicksQueryParamInterval = "90d"
	GetOSByClicksQueryParamIntervalYtd         GetOSByClicksQueryParamInterval = "ytd"
	GetOSByClicksQueryParamIntervalOney        GetOSByClicksQueryParamInterval = "1y"
	GetOSByClicksQueryParamIntervalAll         GetOSByClicksQueryParamInterval = "all"
)

func (e GetOSByClicksQueryParamInterval) ToPointer() *GetOSByClicksQueryParamInterval {
	return &e
}

func (e *GetOSByClicksQueryParamInterval) UnmarshalJSON(data []byte) error {
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
		*e = GetOSByClicksQueryParamInterval(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetOSByClicksQueryParamInterval: %v", v)
	}
}

type GetOSByClicksRequest struct {
	// The domain to filter analytics for.
	Domain *string `queryParam:"style=form,explode=true,name=domain"`
	// The short link slug.
	Key *string `queryParam:"style=form,explode=true,name=key"`
	// The unique ID of the short link on Dub.
	LinkID *string `queryParam:"style=form,explode=true,name=linkId"`
	// This is the ID of the link in the your database. Must be prefixed with 'ext_' when passed as a query parameter.
	ExternalID *string `queryParam:"style=form,explode=true,name=externalId"`
	// The interval to retrieve analytics for. Takes precedence over start and end. If undefined, defaults to 24h.
	Interval *GetOSByClicksQueryParamInterval `default:"24h" queryParam:"style=form,explode=true,name=interval"`
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

func (g GetOSByClicksRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetOSByClicksRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetOSByClicksRequest) GetDomain() *string {
	if o == nil {
		return nil
	}
	return o.Domain
}

func (o *GetOSByClicksRequest) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *GetOSByClicksRequest) GetLinkID() *string {
	if o == nil {
		return nil
	}
	return o.LinkID
}

func (o *GetOSByClicksRequest) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

func (o *GetOSByClicksRequest) GetInterval() *GetOSByClicksQueryParamInterval {
	if o == nil {
		return nil
	}
	return o.Interval
}

func (o *GetOSByClicksRequest) GetStart() *string {
	if o == nil {
		return nil
	}
	return o.Start
}

func (o *GetOSByClicksRequest) GetEnd() *string {
	if o == nil {
		return nil
	}
	return o.End
}

func (o *GetOSByClicksRequest) GetCountry() *components.CountryCode {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *GetOSByClicksRequest) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *GetOSByClicksRequest) GetDevice() *string {
	if o == nil {
		return nil
	}
	return o.Device
}

func (o *GetOSByClicksRequest) GetBrowser() *string {
	if o == nil {
		return nil
	}
	return o.Browser
}

func (o *GetOSByClicksRequest) GetOs() *string {
	if o == nil {
		return nil
	}
	return o.Os
}

func (o *GetOSByClicksRequest) GetReferer() *string {
	if o == nil {
		return nil
	}
	return o.Referer
}

func (o *GetOSByClicksRequest) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *GetOSByClicksRequest) GetTagID() *string {
	if o == nil {
		return nil
	}
	return o.TagID
}

func (o *GetOSByClicksRequest) GetQr() *bool {
	if o == nil {
		return nil
	}
	return o.Qr
}

func (o *GetOSByClicksRequest) GetRoot() *bool {
	if o == nil {
		return nil
	}
	return o.Root
}

type GetOSByClicksResponseBody struct {
	// The name of the OS
	Os string `json:"os"`
	// The number of clicks from this OS
	Clicks float64 `json:"clicks"`
}

func (o *GetOSByClicksResponseBody) GetOs() string {
	if o == nil {
		return ""
	}
	return o.Os
}

func (o *GetOSByClicksResponseBody) GetClicks() float64 {
	if o == nil {
		return 0.0
	}
	return o.Clicks
}

type GetOSByClicksResponse struct {
	HTTPMeta components.HTTPMetadata
	// The top OS by number of clicks
	ResponseBodies []GetOSByClicksResponseBody
}

func (o *GetOSByClicksResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetOSByClicksResponse) GetResponseBodies() []GetOSByClicksResponseBody {
	if o == nil {
		return nil
	}
	return o.ResponseBodies
}
