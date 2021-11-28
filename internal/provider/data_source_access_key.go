package provider

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/madacluster/netmaker-terraform-provider/helper"
)

func AddIdAccessKeySchema() map[string]*schema.Schema {
	result := helper.CreateAccessKeyDataSchema()
	result["id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return result
}

func dataSourceAccessKey() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "AccessKey Data source in the Terraform provider Netmaker.",

		ReadContext: dataSourceAccessKeyRead,

		Schema: helper.CreateAccessKeyDataSchema(),
	}
}

func dataSourceAccessKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	client := meta.(*helper.Client)
	accessKeyID := d.Get("name").(string)
	keyName := d.Get("netid").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	key, err := client.GetKey(accessKeyID, keyName)
	if err != nil {
		return diag.FromErr(err)
	}
	err = helper.SetAccessKeySchemaData(d, key)

	if err != nil {
		return diag.FromErr(err)
	}
	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
