// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dubinc/dub-go/models/components"
)

type GetTagsGlobals struct {
	WorkspaceID string `queryParam:"style=form,explode=true,name=workspaceId"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	ProjectSlug *string `queryParam:"style=form,explode=true,name=projectSlug"`
}

func (o *GetTagsGlobals) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

func (o *GetTagsGlobals) GetProjectSlug() *string {
	if o == nil {
		return nil
	}
	return o.ProjectSlug
}

type GetTagsRequest struct {
}

type GetTagsResponse struct {
	HTTPMeta components.HTTPMetadata
	// A list of tags
	TagSchemas []components.TagSchema
}

func (o *GetTagsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetTagsResponse) GetTagSchemas() []components.TagSchema {
	if o == nil {
		return nil
	}
	return o.TagSchemas
}
