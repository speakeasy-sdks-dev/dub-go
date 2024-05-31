// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dubinc/dub-go/models/components"
)

type TrackLeadGlobals struct {
	WorkspaceID string `queryParam:"style=form,explode=true,name=workspaceId"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	ProjectSlug *string `queryParam:"style=form,explode=true,name=projectSlug"`
}

func (o *TrackLeadGlobals) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

func (o *TrackLeadGlobals) GetProjectSlug() *string {
	if o == nil {
		return nil
	}
	return o.ProjectSlug
}

type TrackLeadRequestBody struct {
	// The ID of the click in th Dub. You can read this value from `dclid` cookie.
	ClickID string `json:"clickId"`
	// The name of the event to track.
	EventName string `json:"eventName"`
	// This is the unique identifier for the customer in the client's app. This is used to track the customer's journey.
	CustomerID string `json:"customerId"`
	// Name of the customer in the client's app.
	CustomerName *string `json:"customerName,omitempty"`
	// Email of the customer in the client's app.
	CustomerEmail *string `json:"customerEmail,omitempty"`
	// Avatar of the customer in the client's app.
	CustomerAvatar *string `json:"customerAvatar,omitempty"`
	// Additional metadata to be stored with the lead event
	Metadata map[string]any `json:"metadata,omitempty"`
}

func (o *TrackLeadRequestBody) GetClickID() string {
	if o == nil {
		return ""
	}
	return o.ClickID
}

func (o *TrackLeadRequestBody) GetEventName() string {
	if o == nil {
		return ""
	}
	return o.EventName
}

func (o *TrackLeadRequestBody) GetCustomerID() string {
	if o == nil {
		return ""
	}
	return o.CustomerID
}

func (o *TrackLeadRequestBody) GetCustomerName() *string {
	if o == nil {
		return nil
	}
	return o.CustomerName
}

func (o *TrackLeadRequestBody) GetCustomerEmail() *string {
	if o == nil {
		return nil
	}
	return o.CustomerEmail
}

func (o *TrackLeadRequestBody) GetCustomerAvatar() *string {
	if o == nil {
		return nil
	}
	return o.CustomerAvatar
}

func (o *TrackLeadRequestBody) GetMetadata() map[string]any {
	if o == nil {
		return nil
	}
	return o.Metadata
}

// TrackLeadResponseBody - A lead was tracked.
type TrackLeadResponseBody struct {
	ClickID        string         `json:"clickId"`
	EventName      string         `json:"eventName"`
	CustomerID     string         `json:"customerId"`
	CustomerName   *string        `json:"customerName"`
	CustomerEmail  *string        `json:"customerEmail"`
	CustomerAvatar *string        `json:"customerAvatar"`
	Metadata       map[string]any `json:"metadata,omitempty"`
}

func (o *TrackLeadResponseBody) GetClickID() string {
	if o == nil {
		return ""
	}
	return o.ClickID
}

func (o *TrackLeadResponseBody) GetEventName() string {
	if o == nil {
		return ""
	}
	return o.EventName
}

func (o *TrackLeadResponseBody) GetCustomerID() string {
	if o == nil {
		return ""
	}
	return o.CustomerID
}

func (o *TrackLeadResponseBody) GetCustomerName() *string {
	if o == nil {
		return nil
	}
	return o.CustomerName
}

func (o *TrackLeadResponseBody) GetCustomerEmail() *string {
	if o == nil {
		return nil
	}
	return o.CustomerEmail
}

func (o *TrackLeadResponseBody) GetCustomerAvatar() *string {
	if o == nil {
		return nil
	}
	return o.CustomerAvatar
}

func (o *TrackLeadResponseBody) GetMetadata() map[string]any {
	if o == nil {
		return nil
	}
	return o.Metadata
}

type TrackLeadResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// A lead was tracked.
	Object *TrackLeadResponseBody
}

func (o *TrackLeadResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *TrackLeadResponse) GetObject() *TrackLeadResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
