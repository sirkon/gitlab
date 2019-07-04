package gitlab

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/sirkon/gitlab/gitlabdata"
)

// NewAPIAccess creates an access point to gitlab API instance
//   httpClient can be nil, http.DefaultClient will be used if it is
//   url must be a full path to gitlab API, e.g. https://gitlab.com/api/v4, etc
func NewAPIAccess(httpClient *http.Client, url string) APIAccess {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &apiAccess{
		client: httpClient,
		url:    url,
	}
}

type apiAccess struct {
	client *http.Client
	url    string
}

func (a *apiAccess) Client(token string) Client {
	return apiClient{
		token:  token,
		access: a,
	}
}

func (a *apiAccess) makeRequest(ctx context.Context, project, token string, keys map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, a.url+project, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create req to gitlab API: %s", err)
	}

	q := req.URL.Query()
	for key, value := range keys {
		q.Set(key, value)
	}
	req.URL.RawQuery = q.Encode()
	req.Header.Set("PRIVATE-TOKEN", token)
	req = req.WithContext(ctx)

	zerolog.Ctx(ctx).Debug().Str("gitlab-url", req.URL.RawPath).Msg("gitlab remote request")
	resp, err := a.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get a response: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		defer closeBody(ctx, resp)
		res, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			zerolog.Ctx(ctx).Error().Err(err).Msg("failed to read out response content")
			return nil, err
		}
		zerolog.Ctx(ctx).Error().Int("error-code", resp.StatusCode).Str("error-response", string(res)).Msg("gitlab error")
		if resp.StatusCode == http.StatusNotFound {
			return nil, os.ErrNotExist
		}
		if len(res)==0 {
			return nil, errors.New("gitlab error", )
		}
		return nil, errors.Errorf("gitlab error: %s", string(res))
	}

	return resp, nil
}

type apiClient struct {
	token  string
	access *apiAccess
}

func closeBody(ctx context.Context, resp *http.Response) {
	if err := resp.Body.Close(); err != nil {
		zerolog.Ctx(ctx).Error().Err(err).Msg("failed to close response body")
	}
}

func (c apiClient) projectURL(project string, items ...string) string {
	urlItems := make([]string, len(items)+3)
	urlItems[0] = ""
	urlItems[1] = "projects"
	urlItems[2] = url.PathEscape(project)
	copy(urlItems[3:], items)
	return strings.Join(urlItems, "/")
}

func (c apiClient) Tags(ctx context.Context, project, tagPrefix string) ([]*gitlabdata.Tag, error) {
	var urlPath string
	if len(tagPrefix) > 0 {
		urlPath = c.projectURL(project, "repository", "tags", tagPrefix)
	} else {
		urlPath = c.projectURL(project, "repository", "tags")
	}

	logger := zerolog.Ctx(ctx).With().Str("gitlab-request", "tags").Str("project", project).Str("tag-prefix", tagPrefix).Logger()
	ctx = (&logger).WithContext(ctx)

	resp, err := c.access.makeRequest(ctx, urlPath, c.token, nil)
	if err != nil {
		zerolog.Ctx(ctx).Error().Err(err).Msg("failed to get requested tags")
		return nil, err
	}
	defer closeBody(ctx, resp)

	var dest []*gitlabdata.Tag
	unmarshaler := json.NewDecoder(resp.Body)
	if len(tagPrefix) > 0 {
		var tag gitlabdata.Tag
		if err := unmarshaler.Decode(&tag); err != nil {
			zerolog.Ctx(ctx).Error().Err(err).Msg("failed to unmarshal a response")
			return nil, err
		}
		dest = append(dest, &tag)
	} else {
		if err := unmarshaler.Decode(&dest); err != nil {
			zerolog.Ctx(ctx).Error().Err(err).Msg("failed to unmarshal a response")
			return nil, err
		}
	}

	return dest, nil
}

func (c apiClient) File(ctx context.Context, project, path, ref string) ([]byte, error) {
	urlPath := c.projectURL(project, "repository", "files", url.PathEscape(path))

	logger := zerolog.Ctx(ctx).With().Str("gitlab-request", "file").Str("project", project).Str("file", path).Logger()
	ctx = (&logger).WithContext(ctx)

	resp, err := c.access.makeRequest(ctx, urlPath, c.token, map[string]string{"ref": ref})
	if err != nil {
		zerolog.Ctx(ctx).Error().Err(err).Msg("failed to get a file")
		return nil, err
	}
	defer closeBody(ctx, resp)

	type responseType struct {
		Encoding      string `json:"encoding"`
		Content       string `json:"content"`
		ContentSHA256 string `json:"content_sha256"`
	}
	dec := json.NewDecoder(resp.Body)
	var dest responseType

	if err := dec.Decode(&dest); err != nil {
		return nil, err
	}

	var content []byte
	switch dest.Encoding {
	case "base64":
		content, err = base64.StdEncoding.DecodeString(dest.Content)
		if err != nil {
			zerolog.Ctx(ctx).Error().Err(err).Msg("failed to decode file content")
			return nil, err
		}
	default:
		zerolog.Ctx(ctx).Error().Msgf("encoding %s is not supported", dest.Encoding)
		return nil, fmt.Errorf("encoding %s is not supported", dest.Encoding)
	}

	return content, nil
}

func (c apiClient) ProjectInfo(ctx context.Context, project string) (*gitlabdata.Project, error) {
	urlPath := c.projectURL(project)

	logger := zerolog.Ctx(ctx).With().Str("gitlab-request", "project-info").Str("project", project).Logger()
	ctx = (&logger).WithContext(ctx)

	resp, err := c.access.makeRequest(ctx, urlPath, c.token, nil)
	if err != nil {
		zerolog.Ctx(ctx).Error().Err(err).Msg("failed to get project info")
		return nil, err
	}
	defer closeBody(ctx, resp)

	var dest gitlabdata.Project
	unmarshaler := json.NewDecoder(resp.Body)
	if err := unmarshaler.Decode(&dest); err != nil {
		zerolog.Ctx(ctx).Error().Err(err).Msg("failed to unmarshal a response")
		return nil, err
	}

	return &dest, nil
}

func (c apiClient) Archive(ctx context.Context, projectID int, tag string) (io.ReadCloser, error) {
	urlPath := c.projectURL(strconv.Itoa(projectID), "repository", "archive.zip")

	logger := zerolog.Ctx(ctx).With().Str("gitlab-request", "archive").Int("project-id", projectID).Str("tag", tag).Logger()
	ctx = (&logger).WithContext(ctx)

	resp, err := c.access.makeRequest(ctx, urlPath, c.token, map[string]string{"sha": tag})
	if err != nil {
		zerolog.Ctx(ctx).Error().Err(err).Msg("failed to get an archive")
		return nil, err
	}

	return resp.Body, nil
}

type referenceItem struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

func (c apiClient) Commits(ctx context.Context, project string, ref string) ([]*gitlabdata.Commit, error) {
	urlPath := c.projectURL(project, "repository", "commits")

	logger := zerolog.Ctx(ctx).With().Str("gitlab-request", "commits").Str("project", project).Logger()
	ctx = (&logger).WithContext(ctx)

	var dest []*gitlabdata.Commit

	resp, err := c.access.makeRequest(ctx, urlPath, c.token, map[string]string{"ref_name": ref})
	if err == nil {
		defer closeBody(ctx, resp)
		unmarshaler := json.NewDecoder(resp.Body)
		if err := unmarshaler.Decode(&dest); err != nil {
			logger.Error().Err(err).Msg("failed to unmarshal commits response")
			return nil, err
		}
		return dest, nil
	}

	logger.Warn().Err(err).Msg("failed to get an archive via branch or tag name, trying to get it via commit SHA")
	referenceURLPath := c.projectURL(project, "repository", "commits", ref, "refs")
	resp, err = c.access.makeRequest(ctx, referenceURLPath, c.token, nil)
	if err != nil {
		logger.Error().Err(err).Msgf("failed to get references for a given commit `%s`", ref)
		return nil, err
	}
	defer closeBody(ctx, resp)
	var references []referenceItem
	unmarshaler := json.NewDecoder(resp.Body)
	if err := unmarshaler.Decode(&references); err != nil {
		logger.Error().Err(err).Msg("failed to unmarshal commit references response")
		return nil, err
	}

	// got commit references, trying them out
	for _, repoRef := range references {
		dest, err = c.Commits(ctx, project, repoRef.Name)
		if err != nil {
			logger.Error().Err(err).Msgf("failed to retrieve commits of %s %s", repoRef.Type, repoRef.Name)
			continue
		}

		// got commits, filter every commit happened after the given one and return the result
		for i, commit := range dest {
			if commit.ShortID == ref || commit.ID == ref {
				return dest[i:], nil
			}
		}
	}

	return nil, fmt.Errorf("failed to get commit history of given branch/tag/commit")
}
