package gitlabdata

import (
	"time"
)

// User represents a GitLab user.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/users.html
type User struct {
	ID                        int                `json:"id"`
	Username                  string             `json:"username"`
	Email                     string             `json:"email"`
	Name                      string             `json:"name"`
	State                     string             `json:"state"`
	CreatedAt                 *time.Time         `json:"created_at"`
	Bio                       string             `json:"bio"`
	Location                  string             `json:"location"`
	PublicEmail               string             `json:"public_email"`
	Skype                     string             `json:"skype"`
	Linkedin                  string             `json:"linkedin"`
	Twitter                   string             `json:"twitter"`
	WebsiteURL                string             `json:"website_url"`
	Organization              string             `json:"organization"`
	ExternUID                 string             `json:"extern_uid"`
	Provider                  string             `json:"provider"`
	ThemeID                   int                `json:"theme_id"`
	LastActivityOn            *ISOTime           `json:"last_activity_on"`
	ColorSchemeID             int                `json:"color_scheme_id"`
	IsAdmin                   bool               `json:"is_admin"`
	AvatarURL                 string             `json:"avatar_url"`
	CanCreateGroup            bool               `json:"can_create_group"`
	CanCreateProject          bool               `json:"can_create_project"`
	ProjectsLimit             int                `json:"projects_limit"`
	CurrentSignInAt           *time.Time         `json:"current_sign_in_at"`
	LastSignInAt              *time.Time         `json:"last_sign_in_at"`
	ConfirmedAt               *time.Time         `json:"confirmed_at"`
	TwoFactorEnabled          bool               `json:"two_factor_enabled"`
	Identities                []*UserIdentity    `json:"identities"`
	External                  bool               `json:"external"`
	PrivateProfile            bool               `json:"private_profile"`
	SharedRunnersMinutesLimit int                `json:"shared_runners_minutes_limit"`
	CustomAttributes          []*CustomAttribute `json:"custom_attributes"`
}

// UserIdentity represents a user identity.
type UserIdentity struct {
	Provider  string `json:"provider"`
	ExternUID string `json:"extern_uid"`
}

// ListUsersOptions represents the available ListUsers() options.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/users.html#list-users
type ListUsersOptions struct {
	ListOptions
	Active  *bool `url:"active,omitempty" json:"active,omitempty"`
	Blocked *bool `url:"blocked,omitempty" json:"blocked,omitempty"`

	// The options below are only available for admins.
	Search               *string    `url:"search,omitempty" json:"search,omitempty"`
	Username             *string    `url:"username,omitempty" json:"username,omitempty"`
	ExternalUID          *string    `url:"extern_uid,omitempty" json:"extern_uid,omitempty"`
	Provider             *string    `url:"provider,omitempty" json:"provider,omitempty"`
	CreatedBefore        *time.Time `url:"created_before,omitempty" json:"created_before,omitempty"`
	CreatedAfter         *time.Time `url:"created_after,omitempty" json:"created_after,omitempty"`
	OrderBy              *string    `url:"order_by,omitempty" json:"order_by,omitempty"`
	Sort                 *string    `url:"sort,omitempty" json:"sort,omitempty"`
	WithCustomAttributes *bool      `url:"with_custom_attributes,omitempty" json:"with_custom_attributes,omitempty"`
}
