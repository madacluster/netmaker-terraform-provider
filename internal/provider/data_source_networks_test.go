package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNetworks(t *testing.T) {
	// t.Skip("data source not yet implemented, remove this once you add your own code")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNetworks,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"data.netmaker_networks.foo", "networks.0.netid", regexp.MustCompile("^private")),
				),
			},
		},
	})
}

const testAccDataSourceNetworks = `
provider "netmaker" {
	username = "admin"
	password = "mx4S6JsSg7JWcZ"
	host = "https://api.netmaker.madacluster.tech"
  }
data "netmaker_networks" "foo" {
}
`
