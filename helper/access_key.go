package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gravitl/netmaker/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (c *Client) CreateAccessKeyFromSchema(d *schema.ResourceData, netID string) (*models.AccessKey, error) {
	key := CreateAccessKeyFromSchema(d)
	return c.CreateKey(netID, *key)
}

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
func CreateAccessKeyFromSchema(d *schema.ResourceData) *models.AccessKey {
	return &models.AccessKey{
		Name:         d.Get("name").(string),
		AccessString: d.Get("key").(string),
		Uses:         d.Get("uses").(int),
	}
}
func (c *Client) GetKeys(networkID string) ([]models.AccessKey, error) {
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

func (c *Client) GetKey(networkID string, keyID string) (*models.AccessKey, error) {
	keys, err := c.GetKeys(networkID)
	if err != nil {
		return nil, err
	}
	for _, key := range keys {
		if key.Name == keyID {
			return &key, nil
		}
	}
	return nil, nil
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

func CreateAccessKeySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the access key",
		},
		"netid": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the access key",
			ForceNew:    true,
		},
		"key": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Key of the access key",
		},
		"uses": {
			Type: schema.TypeInt,
			// Computed:    true,
			Required:    true,
			Description: "Uses of the access key",
		},
	}
}

func SetAccessKeySchemaData(d *schema.ResourceData, key *models.AccessKey, netID string) error {
	ID := fmt.Sprintf("%s-%s", netID, key.Name)
	d.SetId(ID)
	d.Set("name", key.Name)
	d.Set("key", key.AccessString)
	d.Set("uses", key.Uses)
	return nil
}

func (c *Client) UpdateKey(networkID string, key models.AccessKey) error {
	rb, err := json.Marshal(key)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/networks/%s/keyupdate", c.HostURL, networkID), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}
	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) UpdateKeyFromSchema(d *schema.ResourceData, netID string) error {
	key := CreateAccessKeyFromSchema(d)
	return c.UpdateKey(netID, *key)
}
