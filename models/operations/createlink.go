// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"errors"
	"fmt"
	"github.com/dubinc/dub-go/internal/utils"
	"github.com/dubinc/dub-go/models/components"
)

type TagIdsType string

const (
	TagIdsTypeStr        TagIdsType = "str"
	TagIdsTypeArrayOfStr TagIdsType = "arrayOfStr"
)

// TagIds - The unique IDs of the tags assigned to the short link.
type TagIds struct {
	Str        *string
	ArrayOfStr []string

	Type TagIdsType
}

func CreateTagIdsStr(str string) TagIds {
	typ := TagIdsTypeStr

	return TagIds{
		Str:  &str,
		Type: typ,
	}
}

func CreateTagIdsArrayOfStr(arrayOfStr []string) TagIds {
	typ := TagIdsTypeArrayOfStr

	return TagIds{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func (u *TagIds) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = TagIdsTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = TagIdsTypeArrayOfStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for TagIds", string(data))
}

func (u TagIds) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	return nil, errors.New("could not marshal union type TagIds: all fields are null")
}

type TagNamesType string

const (
	TagNamesTypeStr        TagNamesType = "str"
	TagNamesTypeArrayOfStr TagNamesType = "arrayOfStr"
)

// TagNames - The unique name of the tags assigned to the short link (case insensitive).
type TagNames struct {
	Str        *string
	ArrayOfStr []string

	Type TagNamesType
}

func CreateTagNamesStr(str string) TagNames {
	typ := TagNamesTypeStr

	return TagNames{
		Str:  &str,
		Type: typ,
	}
}

func CreateTagNamesArrayOfStr(arrayOfStr []string) TagNames {
	typ := TagNamesTypeArrayOfStr

	return TagNames{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func (u *TagNames) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = TagNamesTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = TagNamesTypeArrayOfStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for TagNames", string(data))
}

func (u TagNames) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	return nil, errors.New("could not marshal union type TagNames: all fields are null")
}

type CreateLinkRequestBody struct {
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
	TagIds *TagIds `json:"tagIds,omitempty"`
	// The unique name of the tags assigned to the short link (case insensitive).
	TagNames *TagNames `json:"tagNames,omitempty"`
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
	// The UTM source of the short link. If set, this will populate or override the UTM source in the destination URL.
	UtmSource *string `json:"utm_source,omitempty"`
	// The UTM medium of the short link. If set, this will populate or override the UTM medium in the destination URL.
	UtmMedium *string `json:"utm_medium,omitempty"`
	// The UTM campaign of the short link. If set, this will populate or override the UTM campaign in the destination URL.
	UtmCampaign *string `json:"utm_campaign,omitempty"`
	// The UTM term of the short link. If set, this will populate or override the UTM term in the destination URL.
	UtmTerm *string `json:"utm_term,omitempty"`
	// The UTM content of the short link. If set, this will populate or override the UTM content in the destination URL.
	UtmContent *string `json:"utm_content,omitempty"`
	// The referral tag of the short link. If set, this will populate or override the `ref` query parameter in the destination URL.
	Ref *string `json:"ref,omitempty"`
}

func (c CreateLinkRequestBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateLinkRequestBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateLinkRequestBody) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

func (o *CreateLinkRequestBody) GetDomain() *string {
	if o == nil {
		return nil
	}
	return o.Domain
}

func (o *CreateLinkRequestBody) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *CreateLinkRequestBody) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

func (o *CreateLinkRequestBody) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *CreateLinkRequestBody) GetTrackConversion() *bool {
	if o == nil {
		return nil
	}
	return o.TrackConversion
}

func (o *CreateLinkRequestBody) GetArchived() *bool {
	if o == nil {
		return nil
	}
	return o.Archived
}

func (o *CreateLinkRequestBody) GetPublicStats() *bool {
	if o == nil {
		return nil
	}
	return o.PublicStats
}

func (o *CreateLinkRequestBody) GetTagID() *string {
	if o == nil {
		return nil
	}
	return o.TagID
}

func (o *CreateLinkRequestBody) GetTagIds() *TagIds {
	if o == nil {
		return nil
	}
	return o.TagIds
}

func (o *CreateLinkRequestBody) GetTagNames() *TagNames {
	if o == nil {
		return nil
	}
	return o.TagNames
}

func (o *CreateLinkRequestBody) GetComments() *string {
	if o == nil {
		return nil
	}
	return o.Comments
}

func (o *CreateLinkRequestBody) GetExpiresAt() *string {
	if o == nil {
		return nil
	}
	return o.ExpiresAt
}

func (o *CreateLinkRequestBody) GetExpiredURL() *string {
	if o == nil {
		return nil
	}
	return o.ExpiredURL
}

func (o *CreateLinkRequestBody) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *CreateLinkRequestBody) GetProxy() *bool {
	if o == nil {
		return nil
	}
	return o.Proxy
}

func (o *CreateLinkRequestBody) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *CreateLinkRequestBody) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateLinkRequestBody) GetImage() *string {
	if o == nil {
		return nil
	}
	return o.Image
}

func (o *CreateLinkRequestBody) GetVideo() *string {
	if o == nil {
		return nil
	}
	return o.Video
}

func (o *CreateLinkRequestBody) GetRewrite() *bool {
	if o == nil {
		return nil
	}
	return o.Rewrite
}

func (o *CreateLinkRequestBody) GetIos() *string {
	if o == nil {
		return nil
	}
	return o.Ios
}

func (o *CreateLinkRequestBody) GetAndroid() *string {
	if o == nil {
		return nil
	}
	return o.Android
}

func (o *CreateLinkRequestBody) GetGeo() *components.LinkGeoTargeting {
	if o == nil {
		return nil
	}
	return o.Geo
}

func (o *CreateLinkRequestBody) GetDoIndex() *bool {
	if o == nil {
		return nil
	}
	return o.DoIndex
}

func (o *CreateLinkRequestBody) GetUtmSource() *string {
	if o == nil {
		return nil
	}
	return o.UtmSource
}

func (o *CreateLinkRequestBody) GetUtmMedium() *string {
	if o == nil {
		return nil
	}
	return o.UtmMedium
}

func (o *CreateLinkRequestBody) GetUtmCampaign() *string {
	if o == nil {
		return nil
	}
	return o.UtmCampaign
}

func (o *CreateLinkRequestBody) GetUtmTerm() *string {
	if o == nil {
		return nil
	}
	return o.UtmTerm
}

func (o *CreateLinkRequestBody) GetUtmContent() *string {
	if o == nil {
		return nil
	}
	return o.UtmContent
}

func (o *CreateLinkRequestBody) GetRef() *string {
	if o == nil {
		return nil
	}
	return o.Ref
}
