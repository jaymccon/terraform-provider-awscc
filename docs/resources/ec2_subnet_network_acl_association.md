---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ec2_subnet_network_acl_association Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::EC2::SubnetNetworkAclAssociation
---

# awscc_ec2_subnet_network_acl_association (Resource)

Resource Type definition for AWS::EC2::SubnetNetworkAclAssociation



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **network_acl_id** (String) The ID of the network ACL
- **subnet_id** (String) The ID of the subnet

### Read-Only

- **association_id** (String)
- **id** (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_ec2_subnet_network_acl_association.example <resource ID>
```
