package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceEgress(t *testing.T) {
	// t.Skip("resource not yet implemented, remove this once you add your own code")
	CreateTestData(t, true)
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceEgress,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"netmaker_egress.foo", "netid", regexp.MustCompile("^test")),
				),
			},
		},
	})
}

const testAccResourceEgress = `
provider "netmaker" {
	username = "admin"
	password = "mx4S6JsSg7JWcZ"
	host = "http://localhost:8081"
}

resource "netmaker_egress" "foo" {
	netid = "test"
	mac = "01:02:03:04:05:06"
	interface = "nm-test"
	ranges = ["0.0.0.0/0"]
}
`
