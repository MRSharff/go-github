// Copyright 2013 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import "fmt"

// RepositoriesService handles communication with the repository related
// methods of the GitHub API.
//
// GitHub API docs: http://developer.github.com/v3/repos/
type RepositoriesService struct {
	client *Client
}

// Repository represents a GitHub repository.
type Repository struct {
	ID              *int       `json:"id,omitempty"`
	Owner           *User      `json:"owner,omitempty"`
	Name            *string    `json:"name,omitempty"`
	Description     *string    `json:"description,omitempty"`
	Homepage        *string    `json:"homepage,omitempty"`
	DefaultBranch   *string    `json:"default_branch,omitempty"`
	MasterBranch    *string    `json:"master_branch,omitempty"`
	CreatedAt       *Timestamp `json:"created_at,omitempty"`
	PushedAt        *Timestamp `json:"pushed_at,omitempty"`
	UpdatedAt       *Timestamp `json:"updated_at,omitempty"`
	URL             *string    `json:"url,omitempty"`
	HTMLURL         *string    `json:"html_url,omitempty"`
	CloneURL        *string    `json:"clone_url,omitempty"`
	GitURL          *string    `json:"git_url,omitempty"`
	MirrorURL       *string    `json:"mirror_url,omitempty"`
	SSHURL          *string    `json:"ssh_url,omitempty"`
	SVNURL          *string    `json:"svn_url,omitempty"`
	Language        *string    `json:"language,omitempty"`
	Fork            *bool      `json:"fork"`
	ForksCount      *int       `json:"forks_count,omitempty"`
	WatchersCount   *int       `json:"watchers_count,omitempty"`
	OpenIssuesCount *int       `json:"open_issues_count,omitempty"`
	Size            *int       `json:"size,omitempty"`

	// Additional mutable fields when creating and editing a repository
	Private   *bool `json:"private"`
	HasIssues *bool `json:"has_issues"`
	HasWiki   *bool `json:"has_wiki"`
}

func (r Repository) String() string {
	return Stringify(r)
}

// RepositoryListOptions specifies the optional parameters to the
// RepositoriesService.List method.
type RepositoryListOptions struct {
	// Type of repositories to list.  Possible values are: all, owner, public,
	// private, member.  Default is "all".
	Type string `url:"type,omitempty"`

	// How to sort the repository list.  Possible values are: created, updated,
	// pushed, full_name.  Default is "full_name".
	Sort string `url:"sort,omitempty"`

	// Direction in which to sort repositories.  Possible values are: asc, desc.
	// Default is "asc" when sort is "full_name", otherwise default is "desc".
	Direction string `url:"direction,omitempty"`

	ListOptions
}

// List the repositories for a user.  Passing the empty string will list
// repositories for the authenticated user.
//
// GitHub API docs: http://developer.github.com/v3/repos/#list-user-repositories
func (s *RepositoriesService) List(user string, opt *RepositoryListOptions) ([]Repository, *Response, error) {
	var u string
	if user != "" {
		u = fmt.Sprintf("users/%v/repos", user)
	} else {
		u = "user/repos"
	}
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	repos := new([]Repository)
	resp, err := s.client.Do(req, repos)
	if err != nil {
		return nil, resp, err
	}

	return *repos, resp, err
}

// RepositoryListByOrgOptions specifies the optional parameters to the
// RepositoriesService.ListByOrg method.
type RepositoryListByOrgOptions struct {
	// Type of repositories to list.  Possible values are: all, public, private,
	// forks, sources, member.  Default is "all".
	Type string `url:"type,omitempty"`

	ListOptions
}

// ListByOrg lists the repositories for an organization.
//
// GitHub API docs: http://developer.github.com/v3/repos/#list-organization-repositories
func (s *RepositoriesService) ListByOrg(org string, opt *RepositoryListByOrgOptions) ([]Repository, *Response, error) {
	u := fmt.Sprintf("orgs/%v/repos", org)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	repos := new([]Repository)
	resp, err := s.client.Do(req, repos)
	if err != nil {
		return nil, resp, err
	}

	return *repos, resp, err
}

// RepositoryListAllOptions specifies the optional parameters to the
// RepositoriesService.ListAll method.
type RepositoryListAllOptions struct {
	// ID of the last repository seen
	Since int `url:"since,omitempty"`

	ListOptions
}

// ListAll lists all GitHub repositories in the order that they were created.
//
// GitHub API docs: http://developer.github.com/v3/repos/#list-all-public-repositories
func (s *RepositoriesService) ListAll(opt *RepositoryListAllOptions) ([]Repository, *Response, error) {
	u, err := addOptions("repositories", opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	repos := new([]Repository)
	resp, err := s.client.Do(req, repos)
	if err != nil {
		return nil, resp, err
	}

	return *repos, resp, err
}

// Create a new repository.  If an organization is specified, the new
// repository will be created under that org.  If the empty string is
// specified, it will be created for the authenticated user.
//
// GitHub API docs: http://developer.github.com/v3/repos/#create
func (s *RepositoriesService) Create(org string, repo *Repository) (*Repository, *Response, error) {
	var u string
	if org != "" {
		u = fmt.Sprintf("orgs/%v/repos", org)
	} else {
		u = "user/repos"
	}

	req, err := s.client.NewRequest("POST", u, repo)
	if err != nil {
		return nil, nil, err
	}

	r := new(Repository)
	resp, err := s.client.Do(req, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, err
}

// Get fetches a repository.
//
// GitHub API docs: http://developer.github.com/v3/repos/#get
func (s *RepositoriesService) Get(owner, repo string) (*Repository, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v", owner, repo)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	repository := new(Repository)
	resp, err := s.client.Do(req, repository)
	if err != nil {
		return nil, resp, err
	}

	return repository, resp, err
}

// Edit updates a repository.
//
// GitHub API docs: http://developer.github.com/v3/repos/#edit
func (s *RepositoriesService) Edit(owner, repo string, repository *Repository) (*Repository, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v", owner, repo)
	req, err := s.client.NewRequest("PATCH", u, repository)
	if err != nil {
		return nil, nil, err
	}

	r := new(Repository)
	resp, err := s.client.Do(req, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, err
}

// Delete a repository.
//
// GitHub API docs: https://developer.github.com/v3/repos/#delete-a-repository
func (s *RepositoriesService) Delete(owner, repo string) (*Response, error) {
	u := fmt.Sprintf("repos/%v/%v", owner, repo)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// Contributor represents a repository contributor
type Contributor struct {
	Login             *string `json:"login,omitempty"`
	ID                *int    `json:"id,omitempty"`
	AvatarURL         *string `json:"avatar_url,omitempty"`
	GravatarID        *string `json:"gravatar_id,omitempty"`
	URL               *string `json:"url,omitempty"`
	HTMLURL           *string `json:"html_url,omitempty"`
	FollowersURL      *string `json:"followers_url,omitempty"`
	FollowingURL      *string `json:"following_url,omitempty"`
	GistsURL          *string `json:"gists_url,omitempty"`
	StarredURL        *string `json:"starred_url,omitempty"`
	SubscriptionsURL  *string `json:"subscriptions_url,omitempty"`
	OrganizationsURL  *string `json:"organizations_url,omitempty"`
	ReposURL          *string `json:"repos_url,omitempty"`
	EventsURL         *string `json:"events_url,omitempty"`
	ReceivedEventsURL *string `json:"received_events_url,omitempty"`
	Type              *string `json:"type,omitempty"`
	SiteAdmin         *bool   `json:"site_admin"`
	Contributions     *int    `json:"contributions,omitempty"`
}

// ListContributorsOptions specifies the optional parameters to the
// RepositoriesService.ListContributors method.
type ListContributorsOptions struct {
	// Include anonymous contributors in results or not
	Anon string `url:"anon,omitempty"`
}

// ListContributors lists contributors for a repository.
//
// GitHub API docs: http://developer.github.com/v3/repos/#list-contributors
func (s *RepositoriesService) ListContributors(owner string, repository string, opt *ListContributorsOptions) ([]Contributor, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/contributors", owner, repository)

	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	contributor := new([]Contributor)
	resp, err := s.client.Do(req, contributor)
	if err != nil {
		return nil, nil, err
	}

	return *contributor, resp, err
}

// ListLanguages lists languages for the specified repository. The returned map
// specifies the languages and the number of bytes of code written in that
// language. For example:
//
//     {
//       "C": 78769,
//       "Python": 7769
//     }
//
// GitHub API Docs: http://developer.github.com/v3/repos/#list-languages
func (s *RepositoriesService) ListLanguages(owner string, repo string) (map[string]int, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/languages", owner, repo)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	languages := make(map[string]int)
	resp, err := s.client.Do(req, &languages)
	if err != nil {
		return nil, resp, err
	}

	return languages, resp, err
}

// ListTeams lists the teams for the specified repository.
//
// GitHub API docs: https://developer.github.com/v3/repos/#list-teams
func (s *RepositoriesService) ListTeams(owner string, repo string) ([]Team, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/teams", owner, repo)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	teams := new([]Team)
	resp, err := s.client.Do(req, teams)
	if err != nil {
		return nil, resp, err
	}

	return *teams, resp, err
}

// RepositoryTag represents a repository tag.
type RepositoryTag struct {
	Name       *string `json:"name,omitempty"`
	Commit     *Commit `json:"commit,omitempty"`
	ZipballURL *string `json:"zipball_url,omitempty"`
	TarballURL *string `json:"tarball_url,omitempty"`
}

// ListTags lists tags for the specified repository.
//
// GitHub API docs: https://developer.github.com/v3/repos/#list-tags
func (s *RepositoriesService) ListTags(owner string, repo string) ([]RepositoryTag, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/tags", owner, repo)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	tags := new([]RepositoryTag)
	resp, err := s.client.Do(req, tags)
	if err != nil {
		return nil, resp, err
	}

	return *tags, resp, err
}

// Branch represents a repository branch
type Branch struct {
	Name   *string `json:"name,omitempty"`
	Commit *Commit `json:"commit,omitempty"`
}

// ListBranches lists branches for the specified repository.
//
// GitHub API docs: http://developer.github.com/v3/repos/#list-branches
func (s *RepositoriesService) ListBranches(owner string, repo string) ([]Branch, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/branches", owner, repo)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	branches := new([]Branch)
	resp, err := s.client.Do(req, branches)
	if err != nil {
		return nil, resp, err
	}

	return *branches, resp, err
}

// GetBranch gets the specified branch for a repository.
//
// GitHub API docs: https://developer.github.com/v3/repos/#get-branch
func (s *RepositoriesService) GetBranch(owner, repo, branch string) (*Branch, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/branches/%v", owner, repo, branch)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	b := new(Branch)
	resp, err := s.client.Do(req, b)
	if err != nil {
		return nil, resp, err
	}

	return b, resp, err
}
