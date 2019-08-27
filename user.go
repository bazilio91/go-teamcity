package teamcity

import "fmt"


type getUserGroupsJson struct {
	Items []UserGroup `json:"group"`
}

func (c client) GetUserGroups() ([]UserGroup, error) {
	response := getUserGroupsJson{}
	err := c.httpGet("/userGroups", nil, &response)
	if err != nil {
		errorf("GetUserGroups failed with %s", err)
		return nil, err
	}

	return response.Items, nil
}



func (c client) GetUserGroup(key string) (*UserGroup, error) {
	var response *UserGroup
	err := c.httpGet(fmt.Sprintf("/userGroups/key:%s", key), nil, &response)
	if err != nil {
		errorf("GetUserGroup failed with %s", err)
		return nil, err
	}

	return response, nil
}

func (c client) CreateUserGroup(group UserGroup) (*UserGroup, error) {
	var response *UserGroup
	err := c.httpPost("/userGroups", nil, group, &response)
	if err != nil {
		errorf("CreateUserGroup failed with %s", err)
		return nil, err
	}

	return response, nil
}

func (c client) GetUser(userLocator string) (*User, error) {
	var response *User
	err := c.httpGet(fmt.Sprintf("/users/%s", userLocator), nil, &response)
	if err != nil {
		errorf("GetUser failed with %s", err)
		return nil, err
	}

	return response, nil
}

func (c client) UpdateUserGroups(userLocator string, groups []UserGroup) ([]UserGroup, error) {
	var response []UserGroup
	err := c.httpPost(fmt.Sprintf("/users/%s/groups", userLocator), nil, groups, &response)
	if err != nil {
		errorf("UpdateUserGroups failed with %s", err)
		return nil, err
	}

	return response, nil
}
