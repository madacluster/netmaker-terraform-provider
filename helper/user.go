package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gravitl/netmaker/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (c *Client) CreateAdmin(user models.User) error {

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

func (c *Client) CreateUser(user models.User) error {

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

func (c *Client) DeleteUser(username string) error {

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/users/%s", c.HostURL, username), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}
	return nil

}

func (c *Client) UpdateUser(user models.User) error {
	rb, err := json.Marshal(user)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/users/%s", c.HostURL, user.UserName), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}
	return nil
}

func CreateUserSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"username": {
			Type:     schema.TypeString,
			Required: true,
		},
		"password": {
			Type:      schema.TypeString,
			Sensitive: true,
			Optional:  true,
		},
		"networks": {
			Type:     schema.TypeList,
			Computed: true,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}

func (c *Client) GetUser(username string) (*models.User, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/users/%s", c.HostURL, username), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUserFromSchemaData(d *schema.ResourceData) *models.User {
	user := &models.User{}
	user.UserName = d.Get("username").(string)
	user.Password = d.Get("password").(string)
	networks := d.Get("network")
	if networks != nil {
		user.Networks = d.Get("network").([]string)
	} else {
		user.Networks = []string{}
	}

	return user
}

func (c *Client) CreateUserFromSchema(d *schema.ResourceData) (*models.User, error) {
	user := CreateUserFromSchemaData(d)
	return user, c.CreateUser(*user)
}

func (c *Client) CreateAdminUserFromSchema(d *schema.ResourceData) (*models.User, error) {
	user := CreateUserFromSchemaData(d)
	return user, c.CreateAdmin(*user)
}

func SetUserSchemaData(d *schema.ResourceData, user *models.User) error {
	d.Set("username", user.UserName)
	d.Set("network", user.Networks)
	return nil
}

func (c *Client) UpdateUserFromSchema(d *schema.ResourceData) (*models.User, error) {
	user := CreateUserFromSchemaData(d)
	return user, c.UpdateUser(*user)
}
