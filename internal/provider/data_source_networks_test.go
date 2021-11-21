package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceScaffolding(t *testing.T) {
	t.Skip("data source not yet implemented, remove this once you add your own code")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceScaffolding,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"data.netmaker_networks.foo", "networks", regexp.MustCompile("^ba")),
				),
			},
		},
	})
}

const testAccDataSourceScaffolding = `
data "netmaker_networks" "foo" {
	networks = "bar"
}
`
