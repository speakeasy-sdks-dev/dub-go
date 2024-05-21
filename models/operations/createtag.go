// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/dubinc/dub-go/models/components"
)

type CreateTagGlobals struct {
	WorkspaceID string `queryParam:"style=form,explode=true,name=workspaceId"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	ProjectSlug *string `queryParam:"style=form,explode=true,name=projectSlug"`
}

func (o *CreateTagGlobals) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

func (o *CreateTagGlobals) GetProjectSlug() *string {
	if o == nil {
		return nil
	}
	return o.ProjectSlug
}

// Color - The color of the tag. If not provided, a random color will be used from the list: red, yellow, green, blue, purple, pink, brown.
type Color string

const (
	ColorRed    Color = "red"
	ColorYellow Color = "yellow"
	ColorGreen  Color = "green"
	ColorBlue   Color = "blue"
	ColorPurple Color = "purple"
	ColorPink   Color = "pink"
	ColorBrown  Color = "brown"
)

func (e Color) ToPointer() *Color {
	return &e
}
func (e *Color) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "red":
		fallthrough
	case "yellow":
		fallthrough
	case "green":
		fallthrough
	case "blue":
		fallthrough
	case "purple":
		fallthrough
	case "pink":
		fallthrough
	case "brown":
		*e = Color(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Color: %v", v)
	}
}

type CreateTagRequestBody struct {
	// The name of the tag to create.
	Tag string `json:"tag"`
	// The color of the tag. If not provided, a random color will be used from the list: red, yellow, green, blue, purple, pink, brown.
	Color *Color `json:"color,omitempty"`
}

func (o *CreateTagRequestBody) GetTag() string {
	if o == nil {
		return ""
	}
	return o.Tag
}

func (o *CreateTagRequestBody) GetColor() *Color {
	if o == nil {
		return nil
	}
	return o.Color
}

type CreateTagResponse struct {
	HTTPMeta components.HTTPMetadata
	// The created tag
	TagSchema *components.TagSchema
}

func (o *CreateTagResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateTagResponse) GetTagSchema() *components.TagSchema {
	if o == nil {
		return nil
	}
	return o.TagSchema
}
