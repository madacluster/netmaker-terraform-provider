package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	client "github.com/madacluster/netmaker-terraform-provider/helper"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			DataSourcesMap: map[string]*schema.Resource{
				"netmaker_networks":   dataSourceNetworks(),
				"netmaker_network":    dataSourceNetwork(),
				"netmaker_access_key": dataSourceAccessKey(),
				"netmaker_node":       dataSourceNode(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"netmaker_network":    resourceNetwork(),
				"netmaker_user":       resourceUser(),
				"netmaker_access_key": resourceAccessKey(),
				"netmaker_egress":     resourceEgress(),
				"netmaker_ingress":    resourceIngress(),
			},
			Schema: map[string]*schema.Schema{
				"username": {
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("NETMAKER_USERNAME", nil),
				},
				"password": {
					Type:        schema.TypeString,
					Required:    true,
					Sensitive:   true,
					DefaultFunc: schema.EnvDefaultFunc("NETMAKER_PASSWORD", nil),
				},
				"host": {
					Type:        schema.TypeString,
					Required:    true,
					Sensitive:   true,
					DefaultFunc: schema.EnvDefaultFunc("NETMAKER_HOST", nil),
				},
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		// Setup a models.User-Agent for your API client (replace the provider name for yours):
		// userAgent := p.UserAgent("terraform-provider-scaffolding", version)
		// TODO: myClient.UserAgent = userAgent
		username := d.Get("username").(string)
		password := d.Get("password").(string)
		host := d.Get("host").(string)

		// Warning or errors can be collected in a slice type
		var diags diag.Diagnostics
		if (username != "") && (password != "") && (host != "") {
			c, err := client.NewClient(&host, &username, &password)
			if err != nil {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Unable to create HashiCups client",
					Detail:   "Unable to auth user for authenticated HashiCups client",
				})
			}

			return c, diags
		}

		c, err := client.NewClient(nil, nil, nil)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create HashiCups client",
				Detail:   "Unable to auth user for authenticated HashiCups client",
			})
		}

		return c, diags
	}
}
