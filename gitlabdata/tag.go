/*
All data structures are copied from github.com/xanzy/go-gitlab
*/

package gitlabdata

import (
	"time"
)

// Tag represents a GitLab tag.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/tags.html
type Tag struct {
	Commit  *Commit  `json:"commit"`
	Release *Release `json:"release"`
	Name    string   `json:"name"`
	Message string   `json:"message"`
}

// Commit represents a GitLab commit.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/commits.html
type Commit struct {
	ID             string           `json:"id"`
	ShortID        string           `json:"short_id"`
	Title          string           `json:"title"`
	AuthorName     string           `json:"author_name"`
	AuthorEmail    string           `json:"author_email"`
	AuthoredDate   *time.Time       `json:"authored_date"`
	CommitterName  string           `json:"committer_name"`
	CommitterEmail string           `json:"committer_email"`
	CommittedDate  *time.Time       `json:"committed_date"`
	CreatedAt      *time.Time       `json:"created_at"`
	Message        string           `json:"message"`
	ParentIDs      []string         `json:"parent_ids"`
	Stats          *CommitStats     `json:"stats"`
	Status         *BuildStateValue `json:"status"`
}

// Release represents a GitLab version release.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/tags.html
type Release struct {
	TagName     string `json:"tag_name"`
	Description string `json:"description"`
}

// CommitStats represents the number of added and deleted files in a commit.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/commits.html
type CommitStats struct {
	Additions int `json:"additions"`
	Deletions int `json:"deletions"`
	Total     int `json:"total"`
}
