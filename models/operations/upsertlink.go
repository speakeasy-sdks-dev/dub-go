// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"errors"
	"fmt"
	"github.com/dubinc/dub-go/internal/utils"
	"github.com/dubinc/dub-go/models/components"
)

type UpsertLinkTagIdsType string

const (
	UpsertLinkTagIdsTypeStr        UpsertLinkTagIdsType = "str"
	UpsertLinkTagIdsTypeArrayOfStr UpsertLinkTagIdsType = "arrayOfStr"
)

// UpsertLinkTagIds - The unique IDs of the tags assigned to the short link.
type UpsertLinkTagIds struct {
	Str        *string
	ArrayOfStr []string

	Type UpsertLinkTagIdsType
}

func CreateUpsertLinkTagIdsStr(str string) UpsertLinkTagIds {
	typ := UpsertLinkTagIdsTypeStr

	return UpsertLinkTagIds{
		Str:  &str,
		Type: typ,
	}
}

func CreateUpsertLinkTagIdsArrayOfStr(arrayOfStr []string) UpsertLinkTagIds {
	typ := UpsertLinkTagIdsTypeArrayOfStr

	return UpsertLinkTagIds{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func (u *UpsertLinkTagIds) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = UpsertLinkTagIdsTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = UpsertLinkTagIdsTypeArrayOfStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for UpsertLinkTagIds", string(data))
}

func (u UpsertLinkTagIds) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	return nil, errors.New("could not marshal union type UpsertLinkTagIds: all fields are null")
}

type UpsertLinkTagNamesType string

const (
	UpsertLinkTagNamesTypeStr        UpsertLinkTagNamesType = "str"
	UpsertLinkTagNamesTypeArrayOfStr UpsertLinkTagNamesType = "arrayOfStr"
)

// UpsertLinkTagNames - The unique name of the tags assigned to the short link (case insensitive).
type UpsertLinkTagNames struct {
	Str        *string
	ArrayOfStr []string

	Type UpsertLinkTagNamesType
}

func CreateUpsertLinkTagNamesStr(str string) UpsertLinkTagNames {
	typ := UpsertLinkTagNamesTypeStr

	return UpsertLinkTagNames{
		Str:  &str,
		Type: typ,
	}
}

func CreateUpsertLinkTagNamesArrayOfStr(arrayOfStr []string) UpsertLinkTagNames {
	typ := UpsertLinkTagNamesTypeArrayOfStr

	return UpsertLinkTagNames{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func (u *UpsertLinkTagNames) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = UpsertLinkTagNamesTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = UpsertLinkTagNamesTypeArrayOfStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for UpsertLinkTagNames", string(data))
}

func (u UpsertLinkTagNames) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	return nil, errors.New("could not marshal union type UpsertLinkTagNames: all fields are null")
}

type UpsertLinkRequestBody struct {
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
	TagIds *UpsertLinkTagIds `json:"tagIds,omitempty"`
	// The unique name of the tags assigned to the short link (case insensitive).
	TagNames *UpsertLinkTagNames `json:"tagNames,omitempty"`
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

func (u UpsertLinkRequestBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpsertLinkRequestBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpsertLinkRequestBody) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

func (o *UpsertLinkRequestBody) GetDomain() *string {
	if o == nil {
		return nil
	}
	return o.Domain
}

func (o *UpsertLinkRequestBody) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *UpsertLinkRequestBody) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

func (o *UpsertLinkRequestBody) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *UpsertLinkRequestBody) GetTrackConversion() *bool {
	if o == nil {
		return nil
	}
	return o.TrackConversion
}

func (o *UpsertLinkRequestBody) GetArchived() *bool {
	if o == nil {
		return nil
	}
	return o.Archived
}

func (o *UpsertLinkRequestBody) GetPublicStats() *bool {
	if o == nil {
		return nil
	}
	return o.PublicStats
}

func (o *UpsertLinkRequestBody) GetTagID() *string {
	if o == nil {
		return nil
	}
	return o.TagID
}

func (o *UpsertLinkRequestBody) GetTagIds() *UpsertLinkTagIds {
	if o == nil {
		return nil
	}
	return o.TagIds
}

func (o *UpsertLinkRequestBody) GetTagNames() *UpsertLinkTagNames {
	if o == nil {
		return nil
	}
	return o.TagNames
}

func (o *UpsertLinkRequestBody) GetComments() *string {
	if o == nil {
		return nil
	}
	return o.Comments
}

func (o *UpsertLinkRequestBody) GetExpiresAt() *string {
	if o == nil {
		return nil
	}
	return o.ExpiresAt
}

func (o *UpsertLinkRequestBody) GetExpiredURL() *string {
	if o == nil {
		return nil
	}
	return o.ExpiredURL
}

func (o *UpsertLinkRequestBody) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *UpsertLinkRequestBody) GetProxy() *bool {
	if o == nil {
		return nil
	}
	return o.Proxy
}

func (o *UpsertLinkRequestBody) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *UpsertLinkRequestBody) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UpsertLinkRequestBody) GetImage() *string {
	if o == nil {
		return nil
	}
	return o.Image
}

func (o *UpsertLinkRequestBody) GetVideo() *string {
	if o == nil {
		return nil
	}
	return o.Video
}

func (o *UpsertLinkRequestBody) GetRewrite() *bool {
	if o == nil {
		return nil
	}
	return o.Rewrite
}

func (o *UpsertLinkRequestBody) GetIos() *string {
	if o == nil {
		return nil
	}
	return o.Ios
}

func (o *UpsertLinkRequestBody) GetAndroid() *string {
	if o == nil {
		return nil
	}
	return o.Android
}

func (o *UpsertLinkRequestBody) GetGeo() *components.LinkGeoTargeting {
	if o == nil {
		return nil
	}
	return o.Geo
}

func (o *UpsertLinkRequestBody) GetDoIndex() *bool {
	if o == nil {
		return nil
	}
	return o.DoIndex
}
