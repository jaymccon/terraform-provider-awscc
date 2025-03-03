---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_memorydb_subnet_group Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::MemoryDB::SubnetGroup resource creates an Amazon MemoryDB Subnet Group.
---

# awscc_memorydb_subnet_group (Resource)

The AWS::MemoryDB::SubnetGroup resource creates an Amazon MemoryDB Subnet Group.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **subnet_group_name** (String) The name of the subnet group. This value must be unique as it also serves as the subnet group identifier.

### Optional

- **description** (String) An optional description of the subnet group.
- **subnet_ids** (Set of String) A list of VPC subnet IDs for the subnet group.
- **tags** (Attributes Set) An array of key-value pairs to apply to this subnet group. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- **arn** (String) The Amazon Resource Name (ARN) of the subnet group.
- **id** (String) Uniquely identifies the resource.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- **key** (String) The key for the tag. May not be null.
- **value** (String) The tag's value. May be null.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_memorydb_subnet_group.example <resource ID>
```
