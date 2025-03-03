---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_memorydb_user Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::MemoryDB::User
---

# awscc_memorydb_user (Data Source)

Data Source schema for AWS::MemoryDB::User



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **access_string** (String) Access permissions string used for this user account.
- **arn** (String) The Amazon Resource Name (ARN) of the user account.
- **authentication_mode** (Attributes) (see [below for nested schema](#nestedatt--authentication_mode))
- **status** (String) Indicates the user status. Can be "active", "modifying" or "deleting".
- **tags** (Attributes Set) An array of key-value pairs to apply to this user. (see [below for nested schema](#nestedatt--tags))
- **user_name** (String) The name of the user.

<a id="nestedatt--authentication_mode"></a>
### Nested Schema for `authentication_mode`

Read-Only:

- **passwords** (List of String) Passwords used for this user account. You can create up to two passwords for each user.
- **type** (String) Type of authentication strategy for this user.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- **key** (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws: or memorydb:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- **value** (String) The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length and cannot be prefixed with aws: or memorydb:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.


