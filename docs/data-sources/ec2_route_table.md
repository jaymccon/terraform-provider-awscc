---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ec2_route_table Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::EC2::RouteTable
---

# awscc_ec2_route_table (Data Source)

Data Source schema for AWS::EC2::RouteTable



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **route_table_id** (String) The route table ID.
- **tags** (Attributes List) Any tags assigned to the route table. (see [below for nested schema](#nestedatt--tags))
- **vpc_id** (String) The ID of the VPC.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- **key** (String)
- **value** (String)


