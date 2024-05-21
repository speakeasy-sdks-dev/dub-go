// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dubinc/dub-go/internal/utils"
	"github.com/dubinc/dub-go/models/components"
)

type GetLinksCountGlobals struct {
	WorkspaceID string `queryParam:"style=form,explode=true,name=workspaceId"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	ProjectSlug *string `queryParam:"style=form,explode=true,name=projectSlug"`
}

func (o *GetLinksCountGlobals) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

func (o *GetLinksCountGlobals) GetProjectSlug() *string {
	if o == nil {
		return nil
	}
	return o.ProjectSlug
}

type GetLinksCountQueryParamTagIdsType string

const (
	GetLinksCountQueryParamTagIdsTypeStr        GetLinksCountQueryParamTagIdsType = "str"
	GetLinksCountQueryParamTagIdsTypeArrayOfStr GetLinksCountQueryParamTagIdsType = "arrayOfStr"
)

// GetLinksCountQueryParamTagIds - The tag IDs to filter the links by.
type GetLinksCountQueryParamTagIds struct {
	Str        *string
	ArrayOfStr []string

	Type GetLinksCountQueryParamTagIdsType
}

func CreateGetLinksCountQueryParamTagIdsStr(str string) GetLinksCountQueryParamTagIds {
	typ := GetLinksCountQueryParamTagIdsTypeStr

	return GetLinksCountQueryParamTagIds{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetLinksCountQueryParamTagIdsArrayOfStr(arrayOfStr []string) GetLinksCountQueryParamTagIds {
	typ := GetLinksCountQueryParamTagIdsTypeArrayOfStr

	return GetLinksCountQueryParamTagIds{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func (u *GetLinksCountQueryParamTagIds) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = GetLinksCountQueryParamTagIdsTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = GetLinksCountQueryParamTagIdsTypeArrayOfStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for GetLinksCountQueryParamTagIds", string(data))
}

func (u GetLinksCountQueryParamTagIds) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	return nil, errors.New("could not marshal union type GetLinksCountQueryParamTagIds: all fields are null")
}

type GetLinksCountQueryParamTagNamesType string

const (
	GetLinksCountQueryParamTagNamesTypeStr        GetLinksCountQueryParamTagNamesType = "str"
	GetLinksCountQueryParamTagNamesTypeArrayOfStr GetLinksCountQueryParamTagNamesType = "arrayOfStr"
)

// GetLinksCountQueryParamTagNames - The unique name of the tags assigned to the short link (case insensitive).
type GetLinksCountQueryParamTagNames struct {
	Str        *string
	ArrayOfStr []string

	Type GetLinksCountQueryParamTagNamesType
}

func CreateGetLinksCountQueryParamTagNamesStr(str string) GetLinksCountQueryParamTagNames {
	typ := GetLinksCountQueryParamTagNamesTypeStr

	return GetLinksCountQueryParamTagNames{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetLinksCountQueryParamTagNamesArrayOfStr(arrayOfStr []string) GetLinksCountQueryParamTagNames {
	typ := GetLinksCountQueryParamTagNamesTypeArrayOfStr

	return GetLinksCountQueryParamTagNames{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func (u *GetLinksCountQueryParamTagNames) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = GetLinksCountQueryParamTagNamesTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = GetLinksCountQueryParamTagNamesTypeArrayOfStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for GetLinksCountQueryParamTagNames", string(data))
}

func (u GetLinksCountQueryParamTagNames) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	return nil, errors.New("could not marshal union type GetLinksCountQueryParamTagNames: all fields are null")
}

type Two string

const (
	TwoTagID Two = "tagId"
)

func (e Two) ToPointer() *Two {
	return &e
}
func (e *Two) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "tagId":
		*e = Two(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Two: %v", v)
	}
}

type One string

const (
	OneDomain One = "domain"
)

func (e One) ToPointer() *One {
	return &e
}
func (e *One) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "domain":
		*e = One(v)
		return nil
	default:
		return fmt.Errorf("invalid value for One: %v", v)
	}
}

type GroupByType string

const (
	GroupByTypeOne GroupByType = "1"
	GroupByTypeTwo GroupByType = "2"
)

// GroupBy - The field to group the links by.
type GroupBy struct {
	One *One
	Two *Two

	Type GroupByType
}

func CreateGroupByOne(one One) GroupBy {
	typ := GroupByTypeOne

	return GroupBy{
		One:  &one,
		Type: typ,
	}
}

func CreateGroupByTwo(two Two) GroupBy {
	typ := GroupByTypeTwo

	return GroupBy{
		Two:  &two,
		Type: typ,
	}
}

func (u *GroupBy) UnmarshalJSON(data []byte) error {

	var one One = One("")
	if err := utils.UnmarshalJSON(data, &one, "", true, true); err == nil {
		u.One = &one
		u.Type = GroupByTypeOne
		return nil
	}

	var two Two = Two("")
	if err := utils.UnmarshalJSON(data, &two, "", true, true); err == nil {
		u.Two = &two
		u.Type = GroupByTypeTwo
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for GroupBy", string(data))
}

func (u GroupBy) MarshalJSON() ([]byte, error) {
	if u.One != nil {
		return utils.MarshalJSON(u.One, "", true)
	}

	if u.Two != nil {
		return utils.MarshalJSON(u.Two, "", true)
	}

	return nil, errors.New("could not marshal union type GroupBy: all fields are null")
}

type GetLinksCountRequest struct {
	// The domain to filter the links by. E.g. `ac.me`. If not provided, all links for the workspace will be returned.
	Domain *string `queryParam:"style=form,explode=true,name=domain"`
	// The tag ID to filter the links by. This field is deprecated – use `tagIds` instead.
	TagID *string `queryParam:"style=form,explode=true,name=tagId"`
	// The tag IDs to filter the links by.
	TagIds *GetLinksCountQueryParamTagIds `queryParam:"style=form,explode=true,name=tagIds"`
	// The unique name of the tags assigned to the short link (case insensitive).
	TagNames *GetLinksCountQueryParamTagNames `queryParam:"style=form,explode=true,name=tagNames"`
	// The search term to filter the links by. The search term will be matched against the short link slug and the destination url.
	Search *string `queryParam:"style=form,explode=true,name=search"`
	// The user ID to filter the links by.
	UserID *string `queryParam:"style=form,explode=true,name=userId"`
	// Whether to include archived links in the response. Defaults to `false` if not provided.
	ShowArchived *bool `default:"false" queryParam:"style=form,explode=true,name=showArchived"`
	// Whether to include tags in the response. Defaults to `false` if not provided.
	WithTags *bool `default:"false" queryParam:"style=form,explode=true,name=withTags"`
	// The field to group the links by.
	GroupBy *GroupBy `queryParam:"style=form,explode=true,name=groupBy"`
}

func (g GetLinksCountRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetLinksCountRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetLinksCountRequest) GetDomain() *string {
	if o == nil {
		return nil
	}
	return o.Domain
}

func (o *GetLinksCountRequest) GetTagID() *string {
	if o == nil {
		return nil
	}
	return o.TagID
}

func (o *GetLinksCountRequest) GetTagIds() *GetLinksCountQueryParamTagIds {
	if o == nil {
		return nil
	}
	return o.TagIds
}

func (o *GetLinksCountRequest) GetTagNames() *GetLinksCountQueryParamTagNames {
	if o == nil {
		return nil
	}
	return o.TagNames
}

func (o *GetLinksCountRequest) GetSearch() *string {
	if o == nil {
		return nil
	}
	return o.Search
}

func (o *GetLinksCountRequest) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

func (o *GetLinksCountRequest) GetShowArchived() *bool {
	if o == nil {
		return nil
	}
	return o.ShowArchived
}

func (o *GetLinksCountRequest) GetWithTags() *bool {
	if o == nil {
		return nil
	}
	return o.WithTags
}

func (o *GetLinksCountRequest) GetGroupBy() *GroupBy {
	if o == nil {
		return nil
	}
	return o.GroupBy
}

type GetLinksCountResponse struct {
	HTTPMeta components.HTTPMetadata
	// A list of links
	Number *float64
}

func (o *GetLinksCountResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetLinksCountResponse) GetNumber() *float64 {
	if o == nil {
		return nil
	}
	return o.Number
}
