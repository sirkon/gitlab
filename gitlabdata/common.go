package gitlabdata

import (
	"encoding/json"
	"fmt"
	"time"
)

// ListOptions specifies the optional parameters to various List methods that
// support pagination.
type ListOptions struct {
	// For paginated result sets, page of results to retrieve.
	Page int `url:"page,omitempty" json:"page,omitempty"`

	// For paginated result sets, the number of results to include per page.
	PerPage int `url:"per_page,omitempty" json:"per_page,omitempty"`
}

// VisibilityValue represents a visibility level within GitLab.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/
type VisibilityValue string

// List of available visibility levels
//
// GitLab API docs: https://docs.gitlab.com/ce/api/
const (
	PrivateVisibility  VisibilityValue = "private"
	InternalVisibility VisibilityValue = "internal"
	PublicVisibility   VisibilityValue = "public"
)

// MergeMethodValue represents a project merge type within GitLab.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/projects.html#project-merge-method
type MergeMethodValue string

// List of available merge type
//
// GitLab API docs: https://docs.gitlab.com/ce/api/projects.html#project-merge-method
const (
	NoFastForwardMerge MergeMethodValue = "merge"
	FastForwardMerge   MergeMethodValue = "ff"
	RebaseMerge        MergeMethodValue = "rebase_merge"
)

// EventTypeValue represents actions type for contribution events
type EventTypeValue string

// List of available action type
//
// GitLab API docs: https://docs.gitlab.com/ce/api/events.html#action-types
const (
	CreatedEventType   EventTypeValue = "created"
	UpdatedEventType   EventTypeValue = "updated"
	ClosedEventType    EventTypeValue = "closed"
	ReopenedEventType  EventTypeValue = "reopened"
	PushedEventType    EventTypeValue = "pushed"
	CommentedEventType EventTypeValue = "commented"
	MergedEventType    EventTypeValue = "merged"
	JoinedEventType    EventTypeValue = "joined"
	LeftEventType      EventTypeValue = "left"
	DestroyedEventType EventTypeValue = "destroyed"
	ExpiredEventType   EventTypeValue = "expired"
)

// EventTargetTypeValue represents actions type value for contribution events
type EventTargetTypeValue string

// List of available action type
//
// GitLab API docs: https://docs.gitlab.com/ce/api/events.html#target-types
const (
	IssueEventTargetType        EventTargetTypeValue = "issue"
	MilestoneEventTargetType    EventTargetTypeValue = "milestone"
	MergeRequestEventTargetType EventTargetTypeValue = "merge_request"
	NoteEventTargetType         EventTargetTypeValue = "note"
	ProjectEventTargetType      EventTargetTypeValue = "project"
	SnippetEventTargetType      EventTargetTypeValue = "snippet"
	UserEventTargetType         EventTargetTypeValue = "user"
)

// CustomAttribute struct is used to unmarshal response to api calls.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/custom_attributes.html
type CustomAttribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// AccessLevelValue represents a permission level within GitLab.
//
// GitLab API docs: https://docs.gitlab.com/ce/permissions/permissions.html
type AccessLevelValue int

// List of available access levels
//
// GitLab API docs: https://docs.gitlab.com/ce/permissions/permissions.html
const (
	NoPermissions         AccessLevelValue = 0
	GuestPermissions      AccessLevelValue = 10
	ReporterPermissions   AccessLevelValue = 20
	DeveloperPermissions  AccessLevelValue = 30
	MaintainerPermissions AccessLevelValue = 40
	OwnerPermissions      AccessLevelValue = 50

	// These are deprecated and should be removed in a future version
	MasterPermissions AccessLevelValue = 40
	OwnerPermission   AccessLevelValue = 50
)

// BuildStateValue represents a GitLab build state.
type BuildStateValue string

// These constants represent all valid build states.
const (
	Pending  BuildStateValue = "pending"
	Running  BuildStateValue = "running"
	Success  BuildStateValue = "success"
	Failed   BuildStateValue = "failed"
	Canceled BuildStateValue = "canceled"
	Skipped  BuildStateValue = "skipped"
)

// ISOTime represents an ISO 8601 formatted date
type ISOTime time.Time

// NotificationLevelValue represents a notification level.
type NotificationLevelValue int

// String implements the fmt.Stringer interface.
func (l NotificationLevelValue) String() string {
	return notificationLevelNames[l]
}

// MarshalJSON implements the json.Marshaler interface.
func (l NotificationLevelValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (l *NotificationLevelValue) UnmarshalJSON(data []byte) error {
	var raw interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch raw := raw.(type) {
	case float64:
		*l = NotificationLevelValue(raw)
	case string:
		*l = notificationLevelTypes[raw]
	case nil:
		// No action needed.
	default:
		return fmt.Errorf("json: cannot unmarshal %T into Go value of type %T", raw, *l)
	}

	return nil
}

// List of valid notification levels.
const (
	DisabledNotificationLevel NotificationLevelValue = iota
	ParticipatingNotificationLevel
	WatchNotificationLevel
	GlobalNotificationLevel
	MentionNotificationLevel
	CustomNotificationLevel
)

var notificationLevelNames = [...]string{
	"disabled",
	"participating",
	"watch",
	"global",
	"mention",
	"custom",
}

var notificationLevelTypes = map[string]NotificationLevelValue{
	"disabled":      DisabledNotificationLevel,
	"participating": ParticipatingNotificationLevel,
	"watch":         WatchNotificationLevel,
	"global":        GlobalNotificationLevel,
	"mention":       MentionNotificationLevel,
	"custom":        CustomNotificationLevel,
}
