---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ssmincidents_replication_set Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource type definition for AWS::SSMIncidents::ReplicationSet
---

# awscc_ssmincidents_replication_set (Resource)

Resource type definition for AWS::SSMIncidents::ReplicationSet



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **regions** (Attributes Set) The ReplicationSet configuration. (see [below for nested schema](#nestedatt--regions))

### Optional

- **deletion_protected** (Boolean) Configures the ReplicationSet deletion protection.

### Read-Only

- **arn** (String) The ARN of the ReplicationSet.
- **id** (String) Uniquely identifies the resource.

<a id="nestedatt--regions"></a>
### Nested Schema for `regions`

Required:

- **region_configuration** (Attributes) The ReplicationSet regional configuration. (see [below for nested schema](#nestedatt--regions--region_configuration))
- **region_name** (String) The AWS region name.

<a id="nestedatt--regions--region_configuration"></a>
### Nested Schema for `regions.region_configuration`

Required:

- **sse_kms_key_id** (String) The ARN of the ReplicationSet.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_ssmincidents_replication_set.example <resource ID>
```
