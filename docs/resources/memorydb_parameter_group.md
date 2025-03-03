---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_memorydb_parameter_group Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::MemoryDB::ParameterGroup resource creates an Amazon MemoryDB ParameterGroup.
---

# awscc_memorydb_parameter_group (Resource)

The AWS::MemoryDB::ParameterGroup resource creates an Amazon MemoryDB ParameterGroup.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **parameter_group_name** (String) The name of the parameter group.

### Optional

- **description** (String) A description of the parameter group.
- **family** (String) The name of the parameter group family that this parameter group is compatible with.
- **parameters** (Map of String) An map of parameter names and values for the parameter update. You must supply at least one parameter name and value; subsequent arguments are optional.
- **tags** (Attributes Set) An array of key-value pairs to apply to this parameter group. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- **arn** (String) The Amazon Resource Name (ARN) of the parameter group.
- **id** (String) Uniquely identifies the resource.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- **key** (String) The key for the tag. May not be null.
- **value** (String) The tag's value. May be null.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_memorydb_parameter_group.example <resource ID>
```
