---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_cloudfront_public_key Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::CloudFront::PublicKey
---

# awscc_cloudfront_public_key (Data Source)

Data Source schema for AWS::CloudFront::PublicKey



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **created_time** (String)
- **public_key_config** (Attributes) (see [below for nested schema](#nestedatt--public_key_config))

<a id="nestedatt--public_key_config"></a>
### Nested Schema for `public_key_config`

Read-Only:

- **caller_reference** (String)
- **comment** (String)
- **encoded_key** (String)
- **name** (String)


