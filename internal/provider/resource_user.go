package provider

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/madacluster/netmaker-terraform-provider/helper"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider scaffolding.",

		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		UpdateContext: resourceUserUpdate,
		DeleteContext: resourceUserDelete,

		Schema: addLastupdated(helper.CreateUserSchema()),
	}
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)
	var diags diag.Diagnostics

	client := meta.(*helper.Client)
	user, err := client.CreateUserFromSchema(d)

	if err != nil {
		return diag.Errorf("failed to create models.User: %s", err)
	}
	d.SetId(user.UserName)
	return diags

}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	client := meta.(*helper.Client)
	UserID := d.Id()

	// d.SetId(UserID)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	user, err := client.GetUser(UserID)
	if err != nil {
		return diag.FromErr(err)
	}

	if helper.SetUserSchemaData(d, user) != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceUserUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	client := meta.(*helper.Client)
	// UserID := d.Id()
	if d.HasChangesExcept("password") {
		_, err := client.UpdateUserFromSchema(d)
		if err != nil {
			return diag.FromErr(err)
		}
		d.Set("last_updated", time.Now().Format(time.RFC850))

	}

	return resourceUserRead(ctx, d, meta)
}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)
	client := meta.(*helper.Client)
	var diags diag.Diagnostics

	UserID := d.Id()
	err := client.DeleteUser(UserID)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
