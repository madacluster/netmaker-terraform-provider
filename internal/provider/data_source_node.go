package provider

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/madacluster/netmaker-terraform-provider/helper"
)

func dataSourceNode() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Node Data source in the Terraform provider Netmaker.",

		ReadContext: dataSourceNodeRead,

		Schema: helper.CreateNodeDataSchema(),
	}
}

func dataSourceNodeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	client := meta.(*helper.Client)
	mac := d.Get("mac").(string)
	netID := d.Get("network_id").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	node, err := client.GetNode(netID, mac)
	if err != nil {
		return diag.FromErr(err)
	}
	helper.SetNodeSchemaData(d, node, netID)

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
