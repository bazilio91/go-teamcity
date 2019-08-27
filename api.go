package teamcity

import (
	"net/url"
	"time"
)

// Authorizer is a TeamCity client authorizer
type Authorizer interface {
	// ResolveUrl provides a full absolute root URL.
	// It should use the following format: baseURL + PREFIX.
	// PREFIX might be either "/guestAuth/app/rest" or "/httpAuth/app/rest" depending on authorization mode.
	ResolveBaseURL(baseURL string) string

	// GetUserInfo provides credentials for HTTP basic auth.
	// It returns nil for guest access mode.
	GetUserInfo() *url.Userinfo
}

// Client is a TeamCity client
type Client interface {
	// Get a project by its ID
	GetProjectByID(id string) (Project, error)
	// Get a project by its name
	GetProjectByName(name string) (Project, error)
	// Get list of projects
	GetProjects() ([]Project, error)

	// Get build type by its ID
	GetBuildTypeByID(id string) (BuildType, error)
	// Get list of all build types
	GetBuildTypes() ([]BuildType, error)
	// Get list of build types for a project
	GetBuildTypesForProject(id string) ([]BuildType, error)
	// Get statistics for last build
	GetBuildTypeStatistics(id int) (BuildStatistics, error)

	// Get build by its ID
	GetBuildByID(id int) (Build, error)
	// Get N latest builds
	GetBuilds(count int) ([]Build, error)
	// Get N latest builds for a build type
	GetBuildsForBuildType(id string, count int) ([]Build, error)

	// Get change by its ID
	GetChangeByID(id int) (Change, error)
	// Get N latest changes
	GetChanges(count int) ([]Change, error)
	// Get N latest changes for a project
	GetChangesForProject(id string, count int) ([]Change, error)
	// Get changes for a build
	GetChangesForBuild(id int) ([]Change, error)
	// Get changes for build type since a particular change
	GetChangesForBuildTypeSinceChange(btId string, cId int) ([]Change, error)
	// Get pending changes for build type
	GetChangesForBuildTypePending(id string) ([]Change, error)

	GetUserGroup(key string) (*UserGroup, error)
	GetUserGroups() ([]UserGroup, error)
	CreateUserGroup(UserGroup) (*UserGroup, error)

	GetUser(userLocator string) (*User, error)
	UpdateUserGroups(userLocator string, groups []UserGroup) ([]UserGroup, error)

	GetServerLicensingData() (*ServerLicensingData, error)
}

// Project is a TeamCity project
type Project struct {
	// Project ID
	ID string `json:"id"`
	// Project name
	Name string `json:"name"`
	// Project description
	Description string `json:"description"`
	// Parent project ID
	ParentProjectID string `json:"parentProjectId"`
}

// BuildType is a TeamCity project build configuration
type BuildType struct {
	// Project ID
	ID string `json:"id"`
	// Project name
	Name string `json:"name"`
	// Project description
	Description string `json:"description"`
	// Project ID
	ProjectID string `json:"projectId"`
}

// BuildStatus is a build status enum
type BuildStatus int

const (
	// StatusUnknown is a zero value of BuildStatus
	StatusUnknown BuildStatus = iota

	// StatusSuccess is a status of successful build
	StatusSuccess

	// StatusRunning is a status of build that is currently running
	StatusRunning

	// StatusFailure is a status of failed build
	StatusFailure
)

const DATE_LAYOUT = "20060102T150405-0700"

// Build is a TeamCity project build
type Build struct {
	// Build ID
	ID int `json:"id"`
	// Build Number
	Number string `json:"number"`
	// Build Status
	Status BuildStatus `json:"status"`
	// Build Status Text
	StatusText string `json:"statusText"`
	// Build Progress Percentage
	Progress int `json:"progress"`
	// Build type ID
	BuildTypeID string `json:"buildTypeId"`
	// Build start time
	QueuedDateRaw string `json:"queuedDate"`
	StartDateRaw  string `json:"startDate"`
	FinishDateRaw string `json:"finishDate"`
}

func (b Build) QueuedDate() (time.Time, error) {
	return time.Parse(DATE_LAYOUT, b.QueuedDateRaw)
}

func (b Build) StartDate() (time.Time, error) {
	return time.Parse(DATE_LAYOUT, b.StartDateRaw)
}

func (b Build) FinishDate() (time.Time, error) {
	return time.Parse(DATE_LAYOUT, b.FinishDateRaw)
}

// Change is a TeamCity project change
type Change struct {
	// Change ID
	ID int `json:"id"`
	// VCS revision id
	Version string `json:"version"`
	// Change author username
	Username string `json:"username"`
	// Change date
	Date string `json:"date"`
}

type User struct {
	LastLogin   string      `json:"lastLogin"`
	Roles       []Role      `json:"roles"`
	Groups      []UserGroup `json:"groups"`
	HasPassword bool        `json:"hasPassword"`
	Password    string      `json:"password"`
	Name        string      `json:"name"`
	Realm       string      `json:"realm"`
	ID          int         `json:"id"`
	Href        string      `json:"href"`
	Locator     string      `json:"locator"`
	Email       string      `json:"email"`
	Properties  struct {
		Count    int `json:"count"`
		Property []struct {
			Inherited bool   `json:"inherited"`
			Name      string `json:"name"`
			Type      struct {
				RawValue string `json:"rawValue"`
			} `json:"type"`
			Value string `json:"value"`
		} `json:"property"`
		Href string `json:"href"`
	} `json:"properties"`
	Username string `json:"username"`
}

type Role struct {
	RoleID string `json:"roleId"`
	Scope  string `json:"scope"`
	Href   string `json:"href"`
}

type UserGroup struct {
	// Group id
	Key string `json:"key"`
	// Group name
	Name string `json:"name"`

	Description string `json:"description"`

	ParentGroups []UserGroup `json:"parent-groups"`

	Users []User `json:"users"`

	Roles []Role `json:"roles"`
}
