package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/madacluster/netmaker-terraform-provider/helper"
)

func resourceEgress() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider scaffolding.",

		CreateContext: resourceEgressCreate,
		ReadContext:   resourceEgressRead,
		UpdateContext: resourceEgressUpdate,
		DeleteContext: resourceEgressDelete,

		Schema: addLastupdated(helper.CreateEgressSchema()),
	}
}

func resourceEgressCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)
	var diags diag.Diagnostics

	client := meta.(*helper.Client)
	netID := d.Get("netid").(string)
	mac := d.Get("mac").(string)
	node, err := client.CreateEgressFromSchema(d, netID, mac)

	if err != nil {
		return diag.Errorf("failed to create models.Egress: %s", err)
	}

	ID := fmt.Sprintf("%s-%s", netID, node.MacAddress)
	d.SetId(ID)
	return diags
}

func resourceEgressRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	client := meta.(*helper.Client)
	mac := d.Get("mac").(string)
	netid := d.Get("netid").(string)
	// d.SetId(UserID)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	node, err := client.GetNode(netid, mac)
	if err != nil {
		return diag.FromErr(err)
	}
	helper.SetEgressSchemaData(d, node, netid, mac)

	return diags
}

func resourceEgressUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	resourceEgressDelete(ctx, d, meta)

	return resourceEgressCreate(ctx, d, meta)
}

func resourceEgressDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)
	client := meta.(*helper.Client)
	var diags diag.Diagnostics

	networkID := d.Get("netid").(string)
	mac := d.Get("mac").(string)
	_, err := client.DeleteEgress(networkID, mac)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
