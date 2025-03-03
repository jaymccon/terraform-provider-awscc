// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package memorydb

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_memorydb_user", userDataSourceType)
}

// userDataSourceType returns the Terraform awscc_memorydb_user data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::MemoryDB::User resource type.
func userDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"access_string": {
			// Property: AccessString
			// CloudFormation resource type schema:
			// {
			//   "description": "Access permissions string used for this user account.",
			//   "type": "string"
			// }
			Description: "Access permissions string used for this user account.",
			Type:        types.StringType,
			Computed:    true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the user account.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the user account.",
			Type:        types.StringType,
			Computed:    true,
		},
		"authentication_mode": {
			// Property: AuthenticationMode
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Passwords": {
			//       "$comment": "List of passwords.",
			//       "description": "Passwords used for this user account. You can create up to two passwords for each user.",
			//       "insertionOrder": true,
			//       "items": {
			//         "type": "string"
			//       },
			//       "maxItems": 2,
			//       "minItems": 1,
			//       "type": "array",
			//       "uniqueItems": true
			//     },
			//     "Type": {
			//       "description": "Type of authentication strategy for this user.",
			//       "enum": [
			//         "password"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"passwords": {
						// Property: Passwords
						Description: "Passwords used for this user account. You can create up to two passwords for each user.",
						Type:        types.ListType{ElemType: types.StringType},
						Computed:    true,
					},
					"type": {
						// Property: Type
						Description: "Type of authentication strategy for this user.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "description": "Indicates the user status. Can be \"active\", \"modifying\" or \"deleting\".",
			//   "type": "string"
			// }
			Description: "Indicates the user status. Can be \"active\", \"modifying\" or \"deleting\".",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this user.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws: or memorydb:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length and cannot be prefixed with aws: or memorydb:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this user.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws: or memorydb:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length and cannot be prefixed with aws: or memorydb:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				tfsdk.SetNestedAttributesOptions{},
			),
			Computed: true,
		},
		"user_name": {
			// Property: UserName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the user.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the user.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::MemoryDB::User",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::MemoryDB::User").WithTerraformTypeName("awscc_memorydb_user")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_string":       "AccessString",
		"arn":                 "Arn",
		"authentication_mode": "AuthenticationMode",
		"key":                 "Key",
		"passwords":           "Passwords",
		"status":              "Status",
		"tags":                "Tags",
		"type":                "Type",
		"user_name":           "UserName",
		"value":               "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
