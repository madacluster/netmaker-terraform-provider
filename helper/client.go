package helper

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// AuthStruct -
type AuthStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthResponse -
type AuthResponse struct {
	Code     int    `json:"Code"`
	Message  string `json:"Message"`
	Response struct {
		UserName  string `json:"UserName"`
		AuthToken string `json:"AuthToken"`
	} `json:"Response"`
}

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
	Auth       AuthStruct
}

const HostURL string = "http://localhost:19090"

func NewClient(host, username, password *string) (*Client, error) {

	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default Hashicups URL
		HostURL: HostURL,
		Auth: AuthStruct{
			Username: *username,
			Password: *password,
		},
	}

	if host != nil {
		c.HostURL = *host
	}
	admin, err := c.CheckAdmin()
	if err != nil {
		return nil, err
	}
	if !admin {
		user := User{
			UserName: *username,
			Password: *password,
		}
		c.CreateAdmin(user)
	}
	ar, err := c.SignIn()
	if err != nil {
		return nil, err
	}

	c.Token = ar.Response.AuthToken

	return &c, nil

}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", "Bearer "+c.Token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
