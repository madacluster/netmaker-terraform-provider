package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceNetwork(t *testing.T) {
	// t.Skip("resource not yet implemented, remove this once you add your own code")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceNetwork,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"netmaker_network.foo", "netid", regexp.MustCompile("^netmakertes")),
				),
			},
		},
	})
}

const testAccResourceNetwork = `
provider "netmaker" {
	username = "admin"
	password = "mx4S6JsSg7JWcZ"
	host = "http://localhost:8081"
}
resource "netmaker_network" "foo" {
  netid = "netmakertes"
  addressrange = "10.100.10.0/24"
}
`
