package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/madacluster/netmaker-terraform-provider/helper"
)

func resourceIngress() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider scaffolding.",

		CreateContext: resourceIngressCreate,
		ReadContext:   resourceIngressRead,
		UpdateContext: resourceIngressUpdate,
		DeleteContext: resourceIngressDelete,

		Schema: addLastupdated(helper.CreateIngressSchema()),
	}
}

func resourceIngressCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)
	var diags diag.Diagnostics

	client := meta.(*helper.Client)
	netID := d.Get("netid").(string)
	mac := d.Get("mac").(string)
	node, err := client.CreateIngress(netID, mac)

	if err != nil {
		return diag.Errorf("failed to create models.Ingress: %s", err)
	}

	ID := fmt.Sprintf("%s-%s", netID, node.MacAddress)
	d.SetId(ID)
	return diags
}

func resourceIngressRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	helper.SetIngressSchemaData(d, node, netid, mac)

	return diags
}

func resourceIngressUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	resourceIngressDelete(ctx, d, meta)

	return resourceIngressCreate(ctx, d, meta)
}

func resourceIngressDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)
	client := meta.(*helper.Client)
	var diags diag.Diagnostics

	networkID := d.Get("netid").(string)
	mac := d.Get("mac").(string)
	_, err := client.DeleteIngress(networkID, mac)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
