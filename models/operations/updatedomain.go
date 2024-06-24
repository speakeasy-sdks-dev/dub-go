// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dubinc/dub-go/internal/utils"
	"github.com/dubinc/dub-go/models/components"
)

type UpdateDomainGlobals struct {
	WorkspaceID *string `queryParam:"style=form,explode=true,name=workspaceId"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	ProjectSlug *string `queryParam:"style=form,explode=true,name=projectSlug"`
}

func (o *UpdateDomainGlobals) GetWorkspaceID() *string {
	if o == nil {
		return nil
	}
	return o.WorkspaceID
}

func (o *UpdateDomainGlobals) GetProjectSlug() *string {
	if o == nil {
		return nil
	}
	return o.ProjectSlug
}

type UpdateDomainRequestBody struct {
	// Name of the domain.
	Slug *string `json:"slug,omitempty"`
	// Redirect users to a specific URL when any link under this domain has expired.
	ExpiredURL *string `json:"expiredUrl,omitempty"`
	// Whether to archive this domain. `false` will unarchive a previously archived domain.
	Archived *bool `default:"false" json:"archived"`
	// Provide context to your teammates in the link creation modal by showing them an example of a link to be shortened.
	Placeholder *string `default:"https://dub.co/help/article/what-is-dub" json:"placeholder"`
}

func (u UpdateDomainRequestBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpdateDomainRequestBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpdateDomainRequestBody) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *UpdateDomainRequestBody) GetExpiredURL() *string {
	if o == nil {
		return nil
	}
	return o.ExpiredURL
}

func (o *UpdateDomainRequestBody) GetArchived() *bool {
	if o == nil {
		return nil
	}
	return o.Archived
}

func (o *UpdateDomainRequestBody) GetPlaceholder() *string {
	if o == nil {
		return nil
	}
	return o.Placeholder
}

type UpdateDomainRequest struct {
	// The domain name.
	Slug        string                   `pathParam:"style=simple,explode=false,name=slug"`
	RequestBody *UpdateDomainRequestBody `request:"mediaType=application/json"`
}

func (o *UpdateDomainRequest) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

func (o *UpdateDomainRequest) GetRequestBody() *UpdateDomainRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type UpdateDomainResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The domain was updated.
	DomainSchema *components.DomainSchema
}

func (o *UpdateDomainResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateDomainResponse) GetDomainSchema() *components.DomainSchema {
	if o == nil {
		return nil
	}
	return o.DomainSchema
}
