package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceAccessKey(t *testing.T) {
	// t.Skip("data source not yet implemented, remove this once you add your own code")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAccessKey,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"data.netmaker_access_key.foo", "netid", regexp.MustCompile("^netmakertest")),
				),
			},
		},
	})
}

const testAccDataSourceAccessKey = `
provider "netmaker" {
	username = "admin"
	password = "mx4S6JsSg7JWcZ"
	host = "http://localhost:8081"
}
data "netmaker_access_key" "foo" {
	name = "netmakertest"
	netid = "netmakertest"
}
`
