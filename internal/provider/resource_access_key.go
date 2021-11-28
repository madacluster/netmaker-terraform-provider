package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/madacluster/netmaker-terraform-provider/helper"
)

func resourceAccessKey() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider scaffolding.",

		CreateContext: resourceAccessKeyCreate,
		ReadContext:   resourceAccessKeyRead,
		UpdateContext: resourceAccessKeyUpdate,
		DeleteContext: resourceAccessKeyDelete,

		Schema: addLastupdated(helper.CreateAccessKeySchema()),
	}
}

func resourceAccessKeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)
	var diags diag.Diagnostics

	client := meta.(*helper.Client)
	netID := d.Get("netid").(string)
	key, err := client.CreateAccessKeyFromSchema(d, netID)

	if err != nil {
		return diag.Errorf("failed to create models.AccessKey: %s", err)
	}

	ID := fmt.Sprintf("%s-%s", netID, key.Name)
	d.SetId(ID)
	return diags
}

func resourceAccessKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	client := meta.(*helper.Client)
	accessKeyID := d.Get("name").(string)
	netid := d.Get("netid").(string)
	// d.SetId(UserID)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	accessKey, err := client.GetKey(netid, accessKeyID)
	if err != nil {
		return diag.FromErr(err)
	}

	if helper.SetAccessKeySchemaData(d, accessKey, netid) != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceAccessKeyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	client := meta.(*helper.Client)
	if d.HasChanges("name") {
		networkID := d.Get("netid").(string)
		err := client.UpdateKeyFromSchema(d, networkID)
		if err != nil {
			return diag.FromErr(err)
		}
		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	return resourceAccessKeyRead(ctx, d, meta)
}

func resourceAccessKeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)
	client := meta.(*helper.Client)
	var diags diag.Diagnostics

	networkID := d.Get("netid").(string)
	name := d.Get("name").(string)
	err := client.DeleteKey(networkID, name)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
