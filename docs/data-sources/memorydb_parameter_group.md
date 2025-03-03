---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_memorydb_parameter_group Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::MemoryDB::ParameterGroup
---

# awscc_memorydb_parameter_group (Data Source)

Data Source schema for AWS::MemoryDB::ParameterGroup



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **arn** (String) The Amazon Resource Name (ARN) of the parameter group.
- **description** (String) A description of the parameter group.
- **family** (String) The name of the parameter group family that this parameter group is compatible with.
- **parameter_group_name** (String) The name of the parameter group.
- **parameters** (Map of String) An map of parameter names and values for the parameter update. You must supply at least one parameter name and value; subsequent arguments are optional.
- **tags** (Attributes Set) An array of key-value pairs to apply to this parameter group. (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- **key** (String) The key for the tag. May not be null.
- **value** (String) The tag's value. May be null.


