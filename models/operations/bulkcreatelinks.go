// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"errors"
	"fmt"
	"github.com/dubinc/dub-go/internal/utils"
	"github.com/dubinc/dub-go/models/components"
)

type BulkCreateLinksTagIdsType string

const (
	BulkCreateLinksTagIdsTypeStr        BulkCreateLinksTagIdsType = "str"
	BulkCreateLinksTagIdsTypeArrayOfStr BulkCreateLinksTagIdsType = "arrayOfStr"
)

// BulkCreateLinksTagIds - The unique IDs of the tags assigned to the short link.
type BulkCreateLinksTagIds struct {
	Str        *string
	ArrayOfStr []string

	Type BulkCreateLinksTagIdsType
}

func CreateBulkCreateLinksTagIdsStr(str string) BulkCreateLinksTagIds {
	typ := BulkCreateLinksTagIdsTypeStr

	return BulkCreateLinksTagIds{
		Str:  &str,
		Type: typ,
	}
}

func CreateBulkCreateLinksTagIdsArrayOfStr(arrayOfStr []string) BulkCreateLinksTagIds {
	typ := BulkCreateLinksTagIdsTypeArrayOfStr

	return BulkCreateLinksTagIds{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func (u *BulkCreateLinksTagIds) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = BulkCreateLinksTagIdsTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = BulkCreateLinksTagIdsTypeArrayOfStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for BulkCreateLinksTagIds", string(data))
}

func (u BulkCreateLinksTagIds) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	return nil, errors.New("could not marshal union type BulkCreateLinksTagIds: all fields are null")
}

type BulkCreateLinksTagNamesType string

const (
	BulkCreateLinksTagNamesTypeStr        BulkCreateLinksTagNamesType = "str"
	BulkCreateLinksTagNamesTypeArrayOfStr BulkCreateLinksTagNamesType = "arrayOfStr"
)

// BulkCreateLinksTagNames - The unique name of the tags assigned to the short link (case insensitive).
type BulkCreateLinksTagNames struct {
	Str        *string
	ArrayOfStr []string

	Type BulkCreateLinksTagNamesType
}

func CreateBulkCreateLinksTagNamesStr(str string) BulkCreateLinksTagNames {
	typ := BulkCreateLinksTagNamesTypeStr

	return BulkCreateLinksTagNames{
		Str:  &str,
		Type: typ,
	}
}

func CreateBulkCreateLinksTagNamesArrayOfStr(arrayOfStr []string) BulkCreateLinksTagNames {
	typ := BulkCreateLinksTagNamesTypeArrayOfStr

	return BulkCreateLinksTagNames{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func (u *BulkCreateLinksTagNames) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = BulkCreateLinksTagNamesTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = BulkCreateLinksTagNamesTypeArrayOfStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for BulkCreateLinksTagNames", string(data))
}

func (u BulkCreateLinksTagNames) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	return nil, errors.New("could not marshal union type BulkCreateLinksTagNames: all fields are null")
}

type RequestBody struct {
	// The destination URL of the short link.
	URL string `json:"url"`
	// The domain of the short link. If not provided, the primary domain for the workspace will be used (or `dub.sh` if the workspace has no domains).
	Domain *string `json:"domain,omitempty"`
	// The short link slug. If not provided, a random 7-character slug will be generated.
	Key *string `json:"key,omitempty"`
	// This is the ID of the link in your database. If set, it can be used to identify the link in the future. Must be prefixed with `ext_` when passed as a query parameter.
	ExternalID *string `json:"externalId,omitempty"`
	// The prefix of the short link slug for randomly-generated keys (e.g. if prefix is `/c/`, generated keys will be in the `/c/:key` format). Will be ignored if `key` is provided.
	Prefix *string `json:"prefix,omitempty"`
	// Whether to track conversions for the short link.
	TrackConversion *bool `default:"false" json:"trackConversion"`
	// Whether the short link is archived.
	Archived *bool `default:"false" json:"archived"`
	// Whether the short link's stats are publicly accessible.
	PublicStats *bool `default:"false" json:"publicStats"`
	// The unique ID of the tag assigned to the short link. This field is deprecated – use `tagIds` instead.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	TagID *string `json:"tagId,omitempty"`
	// The unique IDs of the tags assigned to the short link.
	TagIds *BulkCreateLinksTagIds `json:"tagIds,omitempty"`
	// The unique name of the tags assigned to the short link (case insensitive).
	TagNames *BulkCreateLinksTagNames `json:"tagNames,omitempty"`
	// The comments for the short link.
	Comments *string `json:"comments,omitempty"`
	// The date and time when the short link will expire at.
	ExpiresAt *string `json:"expiresAt,omitempty"`
	// The URL to redirect to when the short link has expired.
	ExpiredURL *string `json:"expiredUrl,omitempty"`
	// The password required to access the destination URL of the short link.
	Password *string `json:"password,omitempty"`
	// Whether the short link uses Custom Social Media Cards feature.
	Proxy *bool `default:"false" json:"proxy"`
	// The custom link preview title (og:title). Will be used for Custom Social Media Cards if `proxy` is true. Learn more: https://d.to/og
	Title *string `json:"title,omitempty"`
	// The custom link preview description (og:description). Will be used for Custom Social Media Cards if `proxy` is true. Learn more: https://d.to/og
	Description *string `json:"description,omitempty"`
	// The custom link preview image (og:image). Will be used for Custom Social Media Cards if `proxy` is true. Learn more: https://d.to/og
	Image *string `json:"image,omitempty"`
	// The custom link preview video (og:video). Will be used for Custom Social Media Cards if `proxy` is true. Learn more: https://d.to/og
	Video *string `json:"video,omitempty"`
	// Whether the short link uses link cloaking.
	Rewrite *bool `default:"false" json:"rewrite"`
	// The iOS destination URL for the short link for iOS device targeting.
	Ios *string `json:"ios,omitempty"`
	// The Android destination URL for the short link for Android device targeting.
	Android *string `json:"android,omitempty"`
	// Geo targeting information for the short link in JSON format `{[COUNTRY]: https://example.com }`.
	Geo *components.LinkGeoTargeting `json:"geo,omitempty"`
	// Allow search engines to index your short link. Defaults to `false` if not provided. Learn more: https://d.to/noindex
	DoIndex *bool `default:"false" json:"doIndex"`
}

func (r RequestBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RequestBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RequestBody) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

func (o *RequestBody) GetDomain() *string {
	if o == nil {
		return nil
	}
	return o.Domain
}

func (o *RequestBody) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *RequestBody) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

func (o *RequestBody) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *RequestBody) GetTrackConversion() *bool {
	if o == nil {
		return nil
	}
	return o.TrackConversion
}

func (o *RequestBody) GetArchived() *bool {
	if o == nil {
		return nil
	}
	return o.Archived
}

func (o *RequestBody) GetPublicStats() *bool {
	if o == nil {
		return nil
	}
	return o.PublicStats
}

func (o *RequestBody) GetTagID() *string {
	if o == nil {
		return nil
	}
	return o.TagID
}

func (o *RequestBody) GetTagIds() *BulkCreateLinksTagIds {
	if o == nil {
		return nil
	}
	return o.TagIds
}

func (o *RequestBody) GetTagNames() *BulkCreateLinksTagNames {
	if o == nil {
		return nil
	}
	return o.TagNames
}

func (o *RequestBody) GetComments() *string {
	if o == nil {
		return nil
	}
	return o.Comments
}

func (o *RequestBody) GetExpiresAt() *string {
	if o == nil {
		return nil
	}
	return o.ExpiresAt
}

func (o *RequestBody) GetExpiredURL() *string {
	if o == nil {
		return nil
	}
	return o.ExpiredURL
}

func (o *RequestBody) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *RequestBody) GetProxy() *bool {
	if o == nil {
		return nil
	}
	return o.Proxy
}

func (o *RequestBody) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *RequestBody) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *RequestBody) GetImage() *string {
	if o == nil {
		return nil
	}
	return o.Image
}

func (o *RequestBody) GetVideo() *string {
	if o == nil {
		return nil
	}
	return o.Video
}

func (o *RequestBody) GetRewrite() *bool {
	if o == nil {
		return nil
	}
	return o.Rewrite
}

func (o *RequestBody) GetIos() *string {
	if o == nil {
		return nil
	}
	return o.Ios
}

func (o *RequestBody) GetAndroid() *string {
	if o == nil {
		return nil
	}
	return o.Android
}

func (o *RequestBody) GetGeo() *components.LinkGeoTargeting {
	if o == nil {
		return nil
	}
	return o.Geo
}

func (o *RequestBody) GetDoIndex() *bool {
	if o == nil {
		return nil
	}
	return o.DoIndex
}
