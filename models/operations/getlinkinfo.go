// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

type GetLinkInfoRequest struct {
	Domain *string `queryParam:"style=form,explode=true,name=domain"`
	// The key of the link to retrieve. E.g. for `d.to/github`, the key is `github`.
	Key *string `queryParam:"style=form,explode=true,name=key"`
	// The unique ID of the short link.
	LinkID *string `queryParam:"style=form,explode=true,name=linkId"`
	// This is the ID of the link in the your database. Must be prefixed with `ext_` when passed as a query parameter.
	ExternalID *string `queryParam:"style=form,explode=true,name=externalId"`
}

func (o *GetLinkInfoRequest) GetDomain() *string {
	if o == nil {
		return nil
	}
	return o.Domain
}

func (o *GetLinkInfoRequest) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *GetLinkInfoRequest) GetLinkID() *string {
	if o == nil {
		return nil
	}
	return o.LinkID
}

func (o *GetLinkInfoRequest) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}
