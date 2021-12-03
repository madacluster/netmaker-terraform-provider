package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceIngress(t *testing.T) {
	// t.Skip("resource not yet implemented, remove this once you add your own code")
	CreateTestData(t, true)
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIngress,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"netmaker_ingress.foo", "netid", regexp.MustCompile("^test")),
				),
			},
		},
	})
}

const testAccResourceIngress = `
provider "netmaker" {
	username = "admin"
	password = "mx4S6JsSg7JWcZ"
	host = "http://localhost:8081"
}

resource "netmaker_ingress" "foo" {
	netid = "test"
	mac = "01:02:03:04:05:06"
}
`
