package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gravitl/netmaker/models"
)

func (c *Client) CreateKey(networkID string, key models.AccessKey) (*models.AccessKey, error) {
	rb, err := json.Marshal(key)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/networks/%s/keys", c.HostURL, networkID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &key)
	if err != nil {
		return nil, err
	}

	return &key, nil
}

func (c *Client) GetKey(networkID string) ([]models.AccessKey, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/networks/%s/keys", c.HostURL, networkID), nil)
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var key []models.AccessKey
	err = json.Unmarshal(body, &key)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func (c *Client) DeleteKey(networkID string, keyID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/networks/%s/keys/%s", c.HostURL, networkID, keyID), nil)
	if err != nil {
		return err
	}
	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
