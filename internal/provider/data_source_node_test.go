package provider

import (
	"regexp"
	"testing"

	"github.com/gravitl/netmaker/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/madacluster/netmaker-terraform-provider/helper"
)

var token string

const node_id = "testnode"
const net_id = "test"
const key_name = "test"
const node_mac = "01:02:03:04:05:06"

var host = "http://localhost:8081"
var pass = "mx4S6JsSg7JWcZ"
var user = "admin"

// const ipRange = "10.101.0.0/24"

func CreateTestData(t *testing.T, createNode bool) {
	c, err := helper.NewClient(&host, &user, &pass)
	if err != nil {
		t.Fatal(err)
	}

	network := &models.Network{
		AddressRange:        "10.102.0.0/24",
		LocalRange:          "",
		IsLocal:             "no",
		IsDualStack:         "",
		AddressRange6:       "",
		DefaultUDPHolePunch: "yes",
		NetID:               net_id,
	}
	got, err := c.CreateNetwork(*network)
	if err != nil {
		t.Fatal(err)
	}
	key := &models.AccessKey{
		Name: key_name,
		Uses: 10,
	}
	accessKey, err := c.CreateKey(got.NetID, *key)
	// token = accessKey.
	if err != nil {
		t.Fatal(err)
	}
	token = accessKey.Value
	if createNode {
		node := models.Node{
			AccessKey: token,
			PublicKey: "DM5qhLAE20PG9BbfBCger+Ac9D2NDOwCtY1rbYDLf34=", Name: node_id, Endpoint: "10.0.0.1", MacAddress: node_mac, Password: "password", Network: net_id,
		}
		_, err := c.CreateNetworkNode(network.NetID, node)
		if err != nil {
			t.Fatal(err)
		}
	}
	t.Cleanup(func() {
		err := CleanTestData()
		if err != nil {
			t.Fatal(err)
		}
	})
}

func CleanTestData() error {
	c, err := helper.NewClient(&host, &user, &pass)
	if err != nil {
		return err
	}
	err = c.DeleteNetworkNode(net_id, node_mac)
	if err != nil {
		return err
	}
	err = c.DeleteKey(net_id, key_name)
	if err != nil {
		return err
	}
	return c.DeleteNetwork(net_id)
}
func TestAccDataSourceNode(t *testing.T) {
	// t.Skip("data source not yet implemented, remove this once you add your own code")
	CreateTestData(t, true)
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNode,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"data.netmaker_node.foo", "name", regexp.MustCompile("^testnode")),
				),
			},
		},
	})
}

const testAccDataSourceNode = `
provider "netmaker" {
	username = "admin"
	password = "mx4S6JsSg7JWcZ"
	host = "http://localhost:8081"
}

data "netmaker_node" "foo" {
	network_id = "test"
	mac = "01:02:03:04:05:06"
}
`
