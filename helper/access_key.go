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
