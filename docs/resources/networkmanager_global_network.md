---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_networkmanager_global_network Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::NetworkManager::GlobalNetwork type specifies a global network of the user's account
---

# awscc_networkmanager_global_network (Resource)

The AWS::NetworkManager::GlobalNetwork type specifies a global network of the user's account



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **description** (String) The description of the global network.
- **tags** (Attributes List) The tags for the global network. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- **arn** (String) The Amazon Resource Name (ARN) of the global network.
- **id** (String) The ID of the global network.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- **key** (String)
- **value** (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_networkmanager_global_network.example <resource ID>
```
