package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type User struct {
	UserName string   `json:"username"`
	Password string   `json:"password"`
	Network  []string `json:"network"`
}

func (c *Client) CreateAdmin(user User) error {

	rb, err := json.Marshal(user)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/users/adm/createadmin", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}
	return nil

}

func (c *Client) CheckAdmin() (bool, error) {
	admin := false
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/users/adm/hasadmin", c.HostURL), nil)
	if err != nil {
		return admin, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return admin, err
	}

	err = json.Unmarshal(body, &admin)
	if err != nil {
		return admin, err
	}

	return admin, nil
}

func (c *Client) CreateUser(user User) error {

	rb, err := json.Marshal(user)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/users/%s", c.HostURL, user.UserName), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteUser(user User) error {

	rb, err := json.Marshal(user)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/users/%s", c.HostURL, user.UserName), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}
	return nil

}
