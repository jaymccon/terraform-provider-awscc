---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_s3outposts_access_point Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type Definition for AWS::S3Outposts::AccessPoint
---

# awscc_s3outposts_access_point (Resource)

Resource Type Definition for AWS::S3Outposts::AccessPoint



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **bucket** (String) The Amazon Resource Name (ARN) of the bucket you want to associate this AccessPoint with.
- **name** (String) A name for the AccessPoint.
- **vpc_configuration** (Attributes) Virtual Private Cloud (VPC) from which requests can be made to the AccessPoint. (see [below for nested schema](#nestedatt--vpc_configuration))

### Optional

- **policy** (Map of String) The access point policy associated with this access point.

### Read-Only

- **arn** (String) The Amazon Resource Name (ARN) of the specified AccessPoint.
- **id** (String) Uniquely identifies the resource.

<a id="nestedatt--vpc_configuration"></a>
### Nested Schema for `vpc_configuration`

Required:

- **vpc_id** (String) Virtual Private Cloud (VPC) Id from which AccessPoint will allow requests.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_s3outposts_access_point.example <resource ID>
```
