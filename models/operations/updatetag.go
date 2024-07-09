// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
)

// UpdateTagColor - The color of the tag. If not provided, a random color will be used from the list: red, yellow, green, blue, purple, pink, brown.
type UpdateTagColor string

const (
	UpdateTagColorRed    UpdateTagColor = "red"
	UpdateTagColorYellow UpdateTagColor = "yellow"
	UpdateTagColorGreen  UpdateTagColor = "green"
	UpdateTagColorBlue   UpdateTagColor = "blue"
	UpdateTagColorPurple UpdateTagColor = "purple"
	UpdateTagColorPink   UpdateTagColor = "pink"
	UpdateTagColorBrown  UpdateTagColor = "brown"
)

func (e UpdateTagColor) ToPointer() *UpdateTagColor {
	return &e
}
func (e *UpdateTagColor) UnmarshalJSON(data []byte) error {
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
		*e = UpdateTagColor(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateTagColor: %v", v)
	}
}

type UpdateTagRequestBody struct {
	// The name of the tag to create.
	Name *string `json:"name,omitempty"`
	// The color of the tag. If not provided, a random color will be used from the list: red, yellow, green, blue, purple, pink, brown.
	Color *UpdateTagColor `json:"color,omitempty"`
	// The name of the tag to create.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Tag *string `json:"tag,omitempty"`
}

func (o *UpdateTagRequestBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateTagRequestBody) GetColor() *UpdateTagColor {
	if o == nil {
		return nil
	}
	return o.Color
}

func (o *UpdateTagRequestBody) GetTag() *string {
	if o == nil {
		return nil
	}
	return o.Tag
}

type UpdateTagRequest struct {
	// The ID of the tag to update.
	ID          string                `pathParam:"style=simple,explode=false,name=id"`
	RequestBody *UpdateTagRequestBody `request:"mediaType=application/json"`
}

func (o *UpdateTagRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateTagRequest) GetRequestBody() *UpdateTagRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}
