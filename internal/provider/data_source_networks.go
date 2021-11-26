package provider

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/madacluster/netmaker-terraform-provider/helper"
)

func dataSourceNetworks() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "models.Network Data source in the Terraform provider Netmaker.",

		ReadContext: dataSourceNetworksRead,

		Schema: map[string]*schema.Schema{
			"networks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: helper.CreateNetworkSchema(),
				},
			},
		},
	}
}

func dataSourceNetworksRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	client := meta.(*helper.Client)

	// idFromAPI := "my-id"
	// d.SetId(idFromAPI)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	networks, err := client.GetNetworks()
	networksFlatten := helper.FlattenNetworksData(&networks)

	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("networks", networksFlatten); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
