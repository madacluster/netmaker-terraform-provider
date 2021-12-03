package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gravitl/netmaker/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (c *Client) GetNodes() ([]models.Node, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/nodes", c.HostURL), nil)
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	nodes := []models.Node{}
	err = json.Unmarshal(body, &nodes)
	if err != nil {
		return nil, err
	}
	return nodes, nil
}

func (c *Client) GetNetworkNodes(networkID string) ([]models.Node, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/nodes/%s", c.HostURL, networkID), nil)
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	nodes := []models.Node{}
	err = json.Unmarshal(body, &nodes)
	if err != nil {
		return nil, err
	}
	return nodes, nil
}

func (c *Client) CreateNetworkNode(networkID string, node models.Node) (*models.Node, error) {
	rb, err := json.Marshal(node)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/nodes/%s", c.HostURL, networkID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	node = models.Node{}
	err = json.Unmarshal(body, &node)
	if err != nil {
		return nil, err
	}
	return &node, nil
}

func (c *Client) DeleteNetworkNode(networkID, mac string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/nodes/%s/%s", c.HostURL, networkID, mac), nil)
	if err != nil {
		return err
	}
	_, err = c.doRequest(req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GetNode(networkID, mac string) (*models.Node, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/nodes/%s/%s", c.HostURL, networkID, mac), nil)
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	node := &models.Node{}
	err = json.Unmarshal(body, node)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (c *Client) GetNetworkIngress(networkID string) ([]models.Node, error) {
	nodes, err := c.GetNetworkNodes(networkID)
	if err != nil {
		return nil, err
	}
	filter := []models.Node{}
	for _, node := range nodes {
		if node.IsIngressGateway == "yes" {
			filter = append(filter, node)
		}
	}
	return filter, nil
}

func (c *Client) CreateIngress(networkID, mac string) (*models.Node, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/nodes/%s/%s/createingress", c.HostURL, networkID, mac), nil)
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	node := &models.Node{}

	err = json.Unmarshal(body, node)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (c *Client) DeleteIngress(networkID, mac string) (*models.Node, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/nodes/%s/%s/deleteingress", c.HostURL, networkID, mac), nil)
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	node := &models.Node{}

	err = json.Unmarshal(body, node)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (c *Client) GetNetworkEgress(networkID string) ([]models.Node, error) {
	nodes, err := c.GetNetworkNodes(networkID)
	if err != nil {
		return nil, err
	}
	filter := []models.Node{}
	for _, node := range nodes {
		if node.IsEgressGateway == "yes" {
			filter = append(filter, node)
		}
	}
	return filter, nil
}

func (c *Client) CreateEgress(networkID, mac string, gateway *models.EgressGatewayRequest) (*models.Node, error) {
	rb, err := json.Marshal(gateway)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/nodes/%s/%s/creategateway", c.HostURL, networkID, mac), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	node := &models.Node{}

	err = json.Unmarshal(body, node)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (c *Client) DeleteEgress(networkID, mac string) (*models.Node, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/nodes/%s/%s/deletegateway", c.HostURL, networkID, mac), nil)
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	node := &models.Node{}

	err = json.Unmarshal(body, node)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func CreateNodeDataSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The name of the node",
		},
		"mac": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The MAC address of the node",
		},
		"network_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The ID of the network the node belongs to",
		},
		"is_ingress_gateway": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Is the node an ingress gateway",
		},
		"is_egress_gateway": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Is the node an egress gateway",
		},
	}
}

func SetNodeSchemaData(d *schema.ResourceData, node *models.Node, networkID string) {
	d.SetId(node.ID)
	d.Set("name", node.Name)
	d.Set("mac", node.MacAddress)
	d.Set("network_id", node.Network)
	d.Set("is_ingress_gateway", node.IsIngressGateway)
	d.Set("is_egress_gateway", node.IsEgressGateway)
}
