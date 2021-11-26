package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gravitl/netmaker/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// GetNetworks - Returns list of coffees (no auth required)
func (c *Client) GetNetworks() ([]models.Network, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/networks", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	networks := []models.Network{}
	err = json.Unmarshal(body, &networks)
	if err != nil {
		return nil, err
	}

	return networks, nil
}

// GetNetworks - Returns a network by ID
func (c *Client) GetNetwork(networkID string) (*models.Network, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/networks/%s", c.HostURL, networkID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	network := models.Network{}
	err = json.Unmarshal(body, &network)
	if err != nil {
		return nil, err
	}

	return &network, nil
}

func (c *Client) CreateNetworkFromSchema(d *schema.ResourceData) (*models.Network, error) {
	network := CreateNetworkFromSchemaData(d)
	return c.CreateNetwork(*network)
}

// GetNetworks - Create a new network
func (c *Client) CreateNetwork(network models.Network) (*models.Network, error) {

	rb, err := json.Marshal(network)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/networks", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return nil, err
	}

	return &network, nil
}

// UpdateNetwork - Updates a network
func (c *Client) UpdateNetwork(network models.Network) (*models.Network, error) {

	rb, err := json.Marshal(network)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/networks/%s", c.HostURL, network.NetID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	network = models.Network{}
	err = json.Unmarshal(body, &network)
	if err != nil {
		return nil, err
	}

	return &network, nil
}

func (c *Client) UpdateNetworkFromSchema(d *schema.ResourceData) (*models.Network, error) {
	network := CreateNetworkFromSchemaData(d)
	return c.UpdateNetwork(*network)
}

func (c *Client) UpdateNetworkMap(data map[string]string) (*models.Network, error) {
	network := models.Network{}
	mapFiels(data, &network)
	return c.UpdateNetwork(network)
}

func (c *Client) DeleteNetwork(networkID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/networks/%s", c.HostURL, networkID), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func mapFiels(data map[string]string, network *models.Network) {
	for k, v := range data {
		switch k {
		case "displayname":
			network.DisplayName = v
		case "addressrange":
			network.AddressRange = v
		case "netid":
			network.NetID = v
		case "islocal":
			network.IsLocal = v
		case "isDualStack":
			network.IsDualStack = v
		case "isIPv4":
			network.IsIPv4 = v
		case "isIPv6":
			network.IsIPv6 = v
		case "isGRPCHub":
			network.IsGRPCHub = v
		case "localrange":
			network.LocalRange = v
		case "checkininterval":
			i, _ := strconv.Atoi(v)
			network.DefaultCheckInInterval = int32(i)
		case "defaultudpholepunch":
			network.DefaultUDPHolePunch = v
		case "defaultextclientdns":
			network.DefaultExtClientDNS = v
		case "defaultmtu":
			i, _ := strconv.Atoi(v)
			network.DefaultMTU = int32(i)
		case "defaultkeepalive":
			i, _ := strconv.Atoi(v)
			network.DefaultKeepalive = int32(i)
		case "allowmanualsignup":
			network.AllowManualSignUp = v
		case "nodeslastmodified":
			network.NodesLastModified, _ = strconv.ParseInt(v, 10, 64)
		case "defaultinterface":
			network.DefaultInterface = v
		case "defaultlistenport":
			i, _ := strconv.ParseInt(v, 10, 32)
			network.DefaultListenPort = int32(i)
		case "defaultsaveconfig":
			network.DefaultSaveConfig = v
		case "nodelimit":
			i, _ := strconv.ParseInt(v, 10, 32)
			network.NodeLimit = int32(i)
		case "defaultpostup":
			network.DefaultPostUp = v
		case "defaultpostdown":
			network.DefaultPostDown = v
		}
	}
}

func mapFielsRevert(network *models.Network) map[string]string {
	data := make(map[string]string)
	data["displayname"] = network.DisplayName
	data["addressrange"] = network.AddressRange
	data["netid"] = network.NetID
	data["defaultinterface"] = network.DefaultInterface
	data["defaultlistenport"] = strconv.FormatInt(int64(network.DefaultListenPort), 10)
	data["nodelimit"] = strconv.FormatInt(int64(network.NodeLimit), 10)
	data["defaultpostup"] = network.DefaultPostUp
	data["defaultpostdown"] = network.DefaultPostDown
	data["defaultsaveconfig"] = network.DefaultSaveConfig
	data["defaultmtu"] = strconv.FormatInt(int64(network.DefaultMTU), 10)
	data["defaultkeepalive"] = strconv.FormatInt(int64(network.DefaultKeepalive), 10)
	data["allowmanualsignup"] = network.AllowManualSignUp
	data["defaultudpholepunch"] = network.DefaultUDPHolePunch
	data["defaultextclientdns"] = network.DefaultExtClientDNS
	data["islocal"] = network.IsLocal
	data["isDualStack"] = network.IsDualStack
	data["isIPv4"] = network.IsIPv4
	data["isIPv6"] = network.IsIPv6
	data["isGRPCHub"] = network.IsGRPCHub
	data["localrange"] = network.LocalRange
	data["checkininterval"] = strconv.FormatInt(int64(network.DefaultCheckInInterval), 10)
	return data
}

func CreateNetworkSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"netid": {
			Type:     schema.TypeString,
			Required: true,
		},
		"addressrange": {
			Type:     schema.TypeString,
			Required: true,
		},
		"addressrange6": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"displayname": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"islocal": {
			Type:     schema.TypeString,
			Computed: true,
			Default:  nil,
		},
		"isdualstack": {
			Type:     schema.TypeString,
			Computed: true,
			Default:  nil,
		},
		"isipv4": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"isipv6": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"isgrpchub": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"localrange": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"checkininterval": {
			Type:     schema.TypeInt,
			Computed: true,
		},

		"defaultudpholepunch": {
			Type:     schema.TypeString,
			Computed: true,
			Default:  nil,
		},
		"defaultextclientdns": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"defaultmtu": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"defaultkeepalive": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"allowmanualsignup": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"defaultinterface": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"defaultlistenport": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"defaultsaveconfig": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"nodelimit": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"defaultpostup": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"defaultpostdown": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}

}

func CreateNetworkFromSchemaData(d *schema.ResourceData) *models.Network {
	network := &models.Network{}
	network.NetID = d.Get("netid").(string)
	network.AddressRange = d.Get("addressrange").(string)
	network.AddressRange6 = d.Get("addressrange6").(string)
	network.DisplayName = d.Get("displayname").(string)
	network.IsLocal = d.Get("islocal").(string)
	network.IsDualStack = d.Get("isdualstack").(string)
	network.IsIPv4 = d.Get("isipv4").(string)
	network.IsIPv6 = d.Get("isipv6").(string)
	network.IsGRPCHub = d.Get("isgrpchub").(string)
	network.LocalRange = d.Get("localrange").(string)
	network.DefaultCheckInInterval = int32(d.Get("checkininterval").(int))
	network.DefaultUDPHolePunch = d.Get("defaultudpholepunch").(string)
	network.DefaultExtClientDNS = d.Get("defaultextclientdns").(string)
	network.DefaultMTU = int32(d.Get("defaultmtu").(int))
	network.AllowManualSignUp = d.Get("allowmanualsignup").(string)
	network.DefaultInterface = d.Get("defaultinterface").(string)
	network.DefaultListenPort = int32(d.Get("defaultlistenport").(int))
	network.DefaultSaveConfig = d.Get("defaultsaveconfig").(string)
	network.NodeLimit = int32(d.Get("nodelimit").(int))
	network.DefaultPostUp = d.Get("defaultpostup").(string)
	network.DefaultPostDown = d.Get("defaultpostdown").(string)
	return network
}

func SetNetworkSchemaData(d *schema.ResourceData, network *models.Network) error {
	if err := d.Set("netid", network.NetID); err != nil {
		return err
	}
	if err := d.Set("addressrange", network.AddressRange); err != nil {
		return err
	}
	if err := d.Set("addressrange6", network.AddressRange6); err != nil {
		return err
	}
	if err := d.Set("displayname", network.DisplayName); err != nil {
		return err
	}
	if err := d.Set("islocal", network.IsLocal); err != nil {
		return err
	}
	if err := d.Set("isdualstack", network.IsDualStack); err != nil {
		return err
	}
	if err := d.Set("isipv4", network.IsIPv4); err != nil {
		return err
	}
	if err := d.Set("isipv6", network.IsIPv6); err != nil {
		return err
	}
	if err := d.Set("isgrpchub", network.IsGRPCHub); err != nil {
		return err
	}
	if err := d.Set("localrange", network.LocalRange); err != nil {
		return err
	}
	if err := d.Set("checkininterval", network.DefaultCheckInInterval); err != nil {
		return err
	}

	if err := d.Set("defaultudpholepunch", network.DefaultUDPHolePunch); err != nil {
		return err
	}
	if err := d.Set("defaultextclientdns", network.DefaultExtClientDNS); err != nil {
		return err
	}
	if err := d.Set("defaultmtu", network.DefaultMTU); err != nil {
		return err
	}
	if err := d.Set("defaultkeepalive", network.DefaultKeepalive); err != nil {
		return err
	}
	if err := d.Set("allowmanualsignup", network.AllowManualSignUp); err != nil {
		return err
	}
	return nil
}

func FlattenNetworkData(network *models.Network) map[string]interface{} {
	oi := make(map[string]interface{})
	oi["netid"] = network.NetID
	oi["addressrange"] = network.AddressRange
	oi["addressrange6"] = network.AddressRange6
	oi["displayname"] = network.DisplayName
	oi["islocal"] = network.IsLocal
	oi["isdualstack"] = network.IsDualStack
	oi["isipv4"] = network.IsIPv4
	oi["isipv6"] = network.IsIPv6
	oi["isgrpchub"] = network.IsGRPCHub
	oi["localrange"] = network.LocalRange
	oi["checkininterval"] = network.DefaultCheckInInterval
	oi["defaultinterface"] = network.DefaultInterface
	oi["defaultlistenport"] = network.DefaultListenPort
	oi["nodelimit"] = network.NodeLimit
	oi["defaultpostup"] = network.DefaultPostUp
	oi["defaultpostdown"] = network.DefaultPostDown
	oi["defaultsaveconfig"] = network.DefaultSaveConfig
	oi["defaultmtu"] = network.DefaultMTU
	oi["defaultkeepalive"] = network.DefaultKeepalive
	oi["allowmanualsignup"] = network.AllowManualSignUp
	oi["defaultudpholepunch"] = network.DefaultUDPHolePunch
	oi["defaultextclientdns"] = network.DefaultExtClientDNS
	oi["islocal"] = network.IsLocal
	oi["isdualstack"] = network.IsDualStack
	oi["isipv4"] = network.IsIPv4
	oi["isipv6"] = network.IsIPv6
	oi["isgrpchub"] = network.IsGRPCHub
	oi["localrange"] = network.LocalRange
	oi["checkininterval"] = network.DefaultCheckInInterval

	return oi
}

func FlattenNetworksData(networks *[]models.Network) []interface{} {
	if networks != nil {
		ois := make([]interface{}, len(*networks))
		for i, network := range *networks {
			ois[i] = FlattenNetworkData(&network)
		}

		return ois
	}

	return make([]interface{}, 0)
}
