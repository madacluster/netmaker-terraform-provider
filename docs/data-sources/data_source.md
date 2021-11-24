---
page_title: "Network data soutce - terraform-provider-netmaker"
subcategory: ""
description: |-
  Sample data source in the Terraform provider Netmaker.
---

# Data Source `netmaker_networks`

Sample data source in the netmaker Terraform provider.

## Example Usage

```terraform
data "netmaker_networks" "example" {
  sample_attribute = "foo"
}
```

## Schema

### Required

- **sample_attribute** (String, Required) Sample attribute.

### Optional

- **id** (String, Optional) The ID of this resource.


