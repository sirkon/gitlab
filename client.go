package gitlab

import (
	"context"
	"github.com/sirkon/gitlab/gitlabdata"
	"io"
)

// APIAccess spawns API clients for a given user
type APIAccess interface {
	Client(token string) Client
}

// Client an implementation of gitlab API access for a given user
type Client interface {
	// Tags get all tags for a given project
	Tags(ctx context.Context, project, tagPrefix string) ([]*gitlabdata.Tag, error)

	// File gets a file with given path from a given project
	File(ctx context.Context, project, path, tag string) ([]byte, error)

	// ProjectInfo gets an info for a given project
	ProjectInfo(ctx context.Context, project string) (*gitlabdata.Project, error)

	// Archive gets an archive for a given project. Needs explicit numeric project ID unlike other methods
	Archive(ctx context.Context, projectID int, tag string) (io.ReadCloser, error)
}
