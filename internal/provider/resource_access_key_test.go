package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceAccessKey(t *testing.T) {
	// t.Skip("resource not yet implemented, remove this once you add your own code")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAccessKey,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"netmaker_access_key.foo", "netid", regexp.MustCompile("^test")),
				),
			},
		},
	})
}

const testAccResourceAccessKey = `
provider "netmaker" {
	username = "admin"
	password = "mx4S6JsSg7JWcZ"
	host = "http://localhost:8081"
}
resource "netmaker_network" "foo" {
	netid = "test"
	addressrange = "10.100.10.0/24"
}
resource "netmaker_access_key" "foo" {
	depends_on = ["netmaker_network.foo"]
	netid = "test"
	uses = 10
	name = "test"
}
`
