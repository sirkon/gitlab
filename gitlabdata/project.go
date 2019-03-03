package gitlabdata

import (
	"time"
)

// Project represents a GitLab project.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/projects.html
type Project struct {
	ID                                        int               `json:"id"`
	Description                               string            `json:"description"`
	DefaultBranch                             string            `json:"default_branch"`
	Public                                    bool              `json:"public"`
	Visibility                                VisibilityValue   `json:"visibility"`
	SSHURLToRepo                              string            `json:"ssh_url_to_repo"`
	HTTPURLToRepo                             string            `json:"http_url_to_repo"`
	WebURL                                    string            `json:"web_url"`
	ReadmeURL                                 string            `json:"readme_url"`
	TagList                                   []string          `json:"tag_list"`
	Owner                                     *User             `json:"owner"`
	Name                                      string            `json:"name"`
	NameWithNamespace                         string            `json:"name_with_namespace"`
	Path                                      string            `json:"path"`
	PathWithNamespace                         string            `json:"path_with_namespace"`
	IssuesEnabled                             bool              `json:"issues_enabled"`
	OpenIssuesCount                           int               `json:"open_issues_count"`
	MergeRequestsEnabled                      bool              `json:"merge_requests_enabled"`
	ApprovalsBeforeMerge                      int               `json:"approvals_before_merge"`
	JobsEnabled                               bool              `json:"jobs_enabled"`
	WikiEnabled                               bool              `json:"wiki_enabled"`
	SnippetsEnabled                           bool              `json:"snippets_enabled"`
	ContainerRegistryEnabled                  bool              `json:"container_registry_enabled"`
	CreatedAt                                 *time.Time        `json:"created_at,omitempty"`
	LastActivityAt                            *time.Time        `json:"last_activity_at,omitempty"`
	CreatorID                                 int               `json:"creator_id"`
	Namespace                                 *ProjectNamespace `json:"namespace"`
	ImportStatus                              string            `json:"import_status"`
	ImportError                               string            `json:"import_error"`
	Permissions                               *Permissions      `json:"permissions"`
	Archived                                  bool              `json:"archived"`
	AvatarURL                                 string            `json:"avatar_url"`
	SharedRunnersEnabled                      bool              `json:"shared_runners_enabled"`
	ForksCount                                int               `json:"forks_count"`
	StarCount                                 int               `json:"star_count"`
	RunnersToken                              string            `json:"runners_token"`
	PublicBuilds                              bool              `json:"public_builds"`
	OnlyAllowMergeIfPipelineSucceeds          bool              `json:"only_allow_merge_if_pipeline_succeeds"`
	OnlyAllowMergeIfAllDiscussionsAreResolved bool              `json:"only_allow_merge_if_all_discussions_are_resolved"`
	LFSEnabled                                bool              `json:"lfs_enabled"`
	RequestAccessEnabled                      bool              `json:"request_access_enabled"`
	MergeMethod                               MergeMethodValue  `json:"merge_method"`
	ForkedFromProject                         *ForkParent       `json:"forked_from_project"`
	Mirror                                    bool              `json:"mirror"`
	MirrorUserID                              int               `json:"mirror_user_id"`
	MirrorTriggerBuilds                       bool              `json:"mirror_trigger_builds"`
	OnlyMirrorProtectedBranches               bool              `json:"only_mirror_protected_branches"`
	MirrorOverwritesDivergedBranches          bool              `json:"mirror_overwrites_diverged_branches"`
	SharedWithGroups                          []struct {
		GroupID          int    `json:"group_id"`
		GroupName        string `json:"group_name"`
		GroupAccessLevel int    `json:"group_access_level"`
	} `json:"shared_with_groups"`
	Statistics       *ProjectStatistics `json:"statistics"`
	Links            *Links             `json:"_links,omitempty"`
	CIConfigPath     *string            `json:"ci_config_path"`
	CustomAttributes []*CustomAttribute `json:"custom_attributes"`
}

// Repository represents a repository.
type Repository struct {
	Name              string          `json:"name"`
	Description       string          `json:"description"`
	WebURL            string          `json:"web_url"`
	AvatarURL         string          `json:"avatar_url"`
	GitSSHURL         string          `json:"git_ssh_url"`
	GitHTTPURL        string          `json:"git_http_url"`
	Namespace         string          `json:"namespace"`
	Visibility        VisibilityValue `json:"visibility"`
	PathWithNamespace string          `json:"path_with_namespace"`
	DefaultBranch     string          `json:"default_branch"`
	Homepage          string          `json:"homepage"`
	URL               string          `json:"url"`
	SSHURL            string          `json:"ssh_url"`
	HTTPURL           string          `json:"http_url"`
}

// ProjectNamespace represents a project namespace.
type ProjectNamespace struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	Kind     string `json:"kind"`
	FullPath string `json:"full_path"`
}

// StorageStatistics represents a statistics record for a group or project.
type StorageStatistics struct {
	StorageSize      int64 `json:"storage_size"`
	RepositorySize   int64 `json:"repository_size"`
	LfsObjectsSize   int64 `json:"lfs_objects_size"`
	JobArtifactsSize int64 `json:"job_artifacts_size"`
}

// ProjectStatistics represents a statistics record for a project.
type ProjectStatistics struct {
	StorageStatistics
	CommitCount int `json:"commit_count"`
}

// Permissions represents permissions.
type Permissions struct {
	ProjectAccess *ProjectAccess `json:"project_access"`
	GroupAccess   *GroupAccess   `json:"group_access"`
}

// ProjectAccess represents project access.
type ProjectAccess struct {
	AccessLevel       AccessLevelValue       `json:"access_level"`
	NotificationLevel NotificationLevelValue `json:"notification_level"`
}

// GroupAccess represents group access.
type GroupAccess struct {
	AccessLevel       AccessLevelValue       `json:"access_level"`
	NotificationLevel NotificationLevelValue `json:"notification_level"`
}

// ForkParent represents the parent project when this is a fork.
type ForkParent struct {
	HTTPURLToRepo     string `json:"http_url_to_repo"`
	ID                int    `json:"id"`
	Name              string `json:"name"`
	NameWithNamespace string `json:"name_with_namespace"`
	Path              string `json:"path"`
	PathWithNamespace string `json:"path_with_namespace"`
	WebURL            string `json:"web_url"`
}

// Links represents a project web links for self, issues, merge_requests,
// repo_branches, labels, events, members.
type Links struct {
	Self          string `json:"self"`
	Issues        string `json:"issues"`
	MergeRequests string `json:"merge_requests"`
	RepoBranches  string `json:"repo_branches"`
	Labels        string `json:"labels"`
	Events        string `json:"events"`
	Members       string `json:"members"`
}

// ListProjectsOptions represents the available ListProjects() options.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/projects.html#list-projects
type ListProjectsOptions struct {
	ListOptions
	Archived                 *bool             `url:"archived,omitempty" json:"archived,omitempty"`
	OrderBy                  *string           `url:"order_by,omitempty" json:"order_by,omitempty"`
	Sort                     *string           `url:"sort,omitempty" json:"sort,omitempty"`
	Search                   *string           `url:"search,omitempty" json:"search,omitempty"`
	Simple                   *bool             `url:"simple,omitempty" json:"simple,omitempty"`
	Owned                    *bool             `url:"owned,omitempty" json:"owned,omitempty"`
	Membership               *bool             `url:"membership,omitempty" json:"membership,omitempty"`
	Starred                  *bool             `url:"starred,omitempty" json:"starred,omitempty"`
	Statistics               *bool             `url:"statistics,omitempty" json:"statistics,omitempty"`
	Visibility               *VisibilityValue  `url:"visibility,omitempty" json:"visibility,omitempty"`
	WithIssuesEnabled        *bool             `url:"with_issues_enabled,omitempty" json:"with_issues_enabled,omitempty"`
	WithMergeRequestsEnabled *bool             `url:"with_merge_requests_enabled,omitempty" json:"with_merge_requests_enabled,omitempty"`
	MinAccessLevel           *AccessLevelValue `url:"min_access_level,omitempty" json:"min_access_level,omitempty"`
	WithCustomAttributes     *bool             `url:"with_custom_attributes,omitempty" json:"with_custom_attributes,omitempty"`
}
