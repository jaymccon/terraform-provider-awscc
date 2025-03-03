---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_networkmanager_site Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::NetworkManager::Site type describes a site.
---

# awscc_networkmanager_site (Resource)

The AWS::NetworkManager::Site type describes a site.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **global_network_id** (String) The ID of the global network.

### Optional

- **description** (String) The description of the site.
- **location** (Attributes) The location of the site. (see [below for nested schema](#nestedatt--location))
- **tags** (Attributes List) The tags for the site. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- **id** (String) Uniquely identifies the resource.
- **site_arn** (String) The Amazon Resource Name (ARN) of the site.
- **site_id** (String) The ID of the site.

<a id="nestedatt--location"></a>
### Nested Schema for `location`

Optional:

- **address** (String) The physical address.
- **latitude** (String) The latitude.
- **longitude** (String) The longitude.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- **key** (String)
- **value** (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_networkmanager_site.example <resource ID>
```
