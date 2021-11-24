package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/madacluster/netmaker-terraform-provider/helper"
)

func addId() map[string]*schema.Schema {
	networkSchema := helper.CreateNetworkSchema()
	networkSchema["id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return networkSchema
}

func dataSourceNetwork() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Network Data source in the Terraform provider Netmaker.",

		ReadContext: dataSourceNetworkRead,

		Schema: addId(),
	}
}

func dataSourceNetworkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	client := meta.(*helper.Client)
	networkID := d.Get("id").(string)
	// idFromAPI := "my-id"
	// d.SetId(idFromAPI)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	network, err := client.GetNetwork(networkID)
	if err != nil {
		return diag.FromErr(err)
	}
	err = helper.SetNetworkSchemaData(d, network)

	if err != nil {
		return diag.FromErr(err)
	}
	// always run
	d.SetId(networkID)

	return diags
}
