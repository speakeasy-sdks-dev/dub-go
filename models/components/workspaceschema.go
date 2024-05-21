// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/dubinc/dub-go/internal/utils"
)

// Plan - The plan of the workspace.
type Plan string

const (
	PlanFree          Plan = "free"
	PlanPro           Plan = "pro"
	PlanBusiness      Plan = "business"
	PlanBusinessPlus  Plan = "business plus"
	PlanBusinessExtra Plan = "business extra"
	PlanBusinessMax   Plan = "business max"
	PlanEnterprise    Plan = "enterprise"
)

func (e Plan) ToPointer() *Plan {
	return &e
}

func (e *Plan) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "free":
		fallthrough
	case "pro":
		fallthrough
	case "business":
		fallthrough
	case "business plus":
		fallthrough
	case "business extra":
		fallthrough
	case "business max":
		fallthrough
	case "enterprise":
		*e = Plan(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Plan: %v", v)
	}
}

// Role - The role of the authenticated user in the workspace.
type Role string

const (
	RoleOwner  Role = "owner"
	RoleMember Role = "member"
)

func (e Role) ToPointer() *Role {
	return &e
}

func (e *Role) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "owner":
		fallthrough
	case "member":
		*e = Role(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Role: %v", v)
	}
}

type Users struct {
	// The role of the authenticated user in the workspace.
	Role Role `json:"role"`
}

func (o *Users) GetRole() Role {
	if o == nil {
		return Role("")
	}
	return o.Role
}

type Domains struct {
	// The domain name.
	Slug string `json:"slug"`
	// Whether the domain is the primary domain for the workspace.
	Primary *bool `default:"false" json:"primary"`
}

func (d Domains) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *Domains) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Domains) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

func (o *Domains) GetPrimary() *bool {
	if o == nil {
		return nil
	}
	return o.Primary
}

type WorkspaceSchema struct {
	// The unique ID of the workspace.
	ID string `json:"id"`
	// The name of the workspace.
	Name string `json:"name"`
	// The slug of the workspace.
	Slug string `json:"slug"`
	// The logo of the workspace.
	Logo *string `default:"null" json:"logo"`
	// The usage of the workspace.
	Usage float64 `json:"usage"`
	// The usage limit of the workspace.
	UsageLimit float64 `json:"usageLimit"`
	// The links usage of the workspace.
	LinksUsage float64 `json:"linksUsage"`
	// The links limit of the workspace.
	LinksLimit float64 `json:"linksLimit"`
	// The domains limit of the workspace.
	DomainsLimit float64 `json:"domainsLimit"`
	// The tags limit of the workspace.
	TagsLimit float64 `json:"tagsLimit"`
	// The users limit of the workspace.
	UsersLimit float64 `json:"usersLimit"`
	// The plan of the workspace.
	Plan Plan `json:"plan"`
	// The Stripe ID of the workspace.
	StripeID *string `json:"stripeId"`
	// The date and time when the billing cycle starts for the workspace.
	BillingCycleStart float64 `json:"billingCycleStart"`
	// The date and time when the workspace was created.
	CreatedAt string `json:"createdAt"`
	// The role of the authenticated user in the workspace.
	Users []Users `json:"users"`
	// The domains of the workspace.
	Domains []Domains `json:"domains"`
	// The invite code of the workspace.
	InviteCode *string `json:"inviteCode"`
}

func (w WorkspaceSchema) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WorkspaceSchema) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WorkspaceSchema) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *WorkspaceSchema) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *WorkspaceSchema) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

func (o *WorkspaceSchema) GetLogo() *string {
	if o == nil {
		return nil
	}
	return o.Logo
}

func (o *WorkspaceSchema) GetUsage() float64 {
	if o == nil {
		return 0.0
	}
	return o.Usage
}

func (o *WorkspaceSchema) GetUsageLimit() float64 {
	if o == nil {
		return 0.0
	}
	return o.UsageLimit
}

func (o *WorkspaceSchema) GetLinksUsage() float64 {
	if o == nil {
		return 0.0
	}
	return o.LinksUsage
}

func (o *WorkspaceSchema) GetLinksLimit() float64 {
	if o == nil {
		return 0.0
	}
	return o.LinksLimit
}

func (o *WorkspaceSchema) GetDomainsLimit() float64 {
	if o == nil {
		return 0.0
	}
	return o.DomainsLimit
}

func (o *WorkspaceSchema) GetTagsLimit() float64 {
	if o == nil {
		return 0.0
	}
	return o.TagsLimit
}

func (o *WorkspaceSchema) GetUsersLimit() float64 {
	if o == nil {
		return 0.0
	}
	return o.UsersLimit
}

func (o *WorkspaceSchema) GetPlan() Plan {
	if o == nil {
		return Plan("")
	}
	return o.Plan
}

func (o *WorkspaceSchema) GetStripeID() *string {
	if o == nil {
		return nil
	}
	return o.StripeID
}

func (o *WorkspaceSchema) GetBillingCycleStart() float64 {
	if o == nil {
		return 0.0
	}
	return o.BillingCycleStart
}

func (o *WorkspaceSchema) GetCreatedAt() string {
	if o == nil {
		return ""
	}
	return o.CreatedAt
}

func (o *WorkspaceSchema) GetUsers() []Users {
	if o == nil {
		return []Users{}
	}
	return o.Users
}

func (o *WorkspaceSchema) GetDomains() []Domains {
	if o == nil {
		return []Domains{}
	}
	return o.Domains
}

func (o *WorkspaceSchema) GetInviteCode() *string {
	if o == nil {
		return nil
	}
	return o.InviteCode
}
