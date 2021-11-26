package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceUser(t *testing.T) {
	// t.Skip("resource not yet implemented, remove this once you add your own code")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAdminUser,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"netmaker_user.foo", "username", regexp.MustCompile("^netmakertes")),
				),
			},
		},
	})
}

const testAccResourceAdminUser = `
provider "netmaker" {
	username = "admin"
	password = "mx4S6JsSg7JWcZ"
	host 	 = "http://localhost:8081"
}
resource "netmaker_network" "foo" {
	netid = "netmakertes"
	addressrange = "10.100.10.0/24"
  }
resource "netmaker_user" "foo" {
  username = "netmakertes"
  password = "netmaker"
  networks = ["netmakertes"]
}
`
