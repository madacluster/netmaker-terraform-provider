package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNetwork(t *testing.T) {
	// t.Skip("data source not yet implemented, remove this once you add your own code")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNetwork,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"data.netmaker_network.foo", "netid", regexp.MustCompile("^private")),
				),
			},
		},
	})
}

const testAccDataSourceNetwork = `
provider "netmaker" {
	username = "admin"
	password = "mx4S6JsSg7JWcZ"
	host = "http://localhost:8081"
}
data "netmaker_network" "foo" {
	id = "private"
	netid = "private"
	addressrange = "10.100.10.0/24"
}
`
