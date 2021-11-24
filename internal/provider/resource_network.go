package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/madacluster/netmaker-terraform-provider/helper"
)

func addLastupdated() map[string]*schema.Schema {
	networkSchema := helper.CreateNetworkSchema()
	networkSchema["last_updated"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
	}
	return networkSchema
}

func resourceNetwork() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider scaffolding.",

		CreateContext: resourceNetworkCreate,
		ReadContext:   resourceNetworkRead,
		UpdateContext: resourceNetworkUpdate,
		DeleteContext: resourceScaffoldingDelete,

		Schema: addLastupdated(),
	}
}

func resourceNetworkCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)
	var diags diag.Diagnostics

	client := meta.(*helper.Client)
	network, err := client.CreateNetworkFromSchema(d)

	if err != nil {
		return diag.Errorf("failed to create network: %s", err)
	}
	d.SetId(network.Netid)
	return diags

}

func resourceNetworkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	client := meta.(*helper.Client)
	networkID := d.Id()

	// d.SetId(networkID)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	network, err := client.GetNetwork(networkID)
	if err != nil {
		return diag.FromErr(err)
	}

	if helper.SetNetworkSchemaData(d, network) != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceNetworkUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	client := meta.(*helper.Client)
	networkID := d.Id()
	if d.HasChange("addressrange") {
		name := d.Get("addressrange").(string)
		err := client.UpdateNetworkName(networkID, name)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return diag.Errorf("not implemented")
}

func resourceNetworkDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}
