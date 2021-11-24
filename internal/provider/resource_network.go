package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/madacluster/netmaker-terraform-provider/helper"
)

func resourceScaffolding() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider scaffolding.",

		CreateContext: resourceNetworkCreate,
		ReadContext:   resourceNetworkRead,
		UpdateContext: resourceScaffoldingUpdate,
		DeleteContext: resourceScaffoldingDelete,

		Schema: AddIdNetworkSchema(),
	}
}

func resourceNetworkCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)
	client := meta.(*helper.Client)
	network := helper.CreateNetworkFromSchemaData(d)
	var diags diag.Diagnostics
	client.CreateNetwork(network)
	return diag.Errorf("not implemented")
}

func resourceNetworkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	client := meta.(*helper.Client)
	networkID := d.Get("id").(string)

	d.SetId(networkID)

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

func resourceScaffoldingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func resourceScaffoldingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}
