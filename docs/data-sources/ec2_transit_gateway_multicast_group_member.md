---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ec2_transit_gateway_multicast_group_member Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::EC2::TransitGatewayMulticastGroupMember
---

# awscc_ec2_transit_gateway_multicast_group_member (Data Source)

Data Source schema for AWS::EC2::TransitGatewayMulticastGroupMember



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **group_ip_address** (String) The IP address assigned to the transit gateway multicast group.
- **group_member** (Boolean) Indicates that the resource is a transit gateway multicast group member.
- **group_source** (Boolean) Indicates that the resource is a transit gateway multicast group member.
- **member_type** (String) The member type (for example, static).
- **network_interface_id** (String) The ID of the transit gateway attachment.
- **resource_id** (String) The ID of the resource.
- **resource_type** (String) The type of resource, for example a VPC attachment.
- **source_type** (String) The source type.
- **subnet_id** (String) The ID of the subnet.
- **transit_gateway_attachment_id** (String) The ID of the transit gateway attachment.
- **transit_gateway_multicast_domain_id** (String) The ID of the transit gateway multicast domain.


