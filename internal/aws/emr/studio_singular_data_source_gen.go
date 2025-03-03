// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package emr

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_emr_studio", studioDataSourceType)
}

// studioDataSourceType returns the Terraform awscc_emr_studio data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::EMR::Studio resource type.
func studioDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the EMR Studio.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the EMR Studio.",
			Type:        types.StringType,
			Computed:    true,
		},
		"auth_mode": {
			// Property: AuthMode
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies whether the Studio authenticates users using single sign-on (SSO) or IAM. Amazon EMR Studio currently only supports SSO authentication.",
			//   "enum": [
			//     "SSO",
			//     "IAM"
			//   ],
			//   "type": "string"
			// }
			Description: "Specifies whether the Studio authenticates users using single sign-on (SSO) or IAM. Amazon EMR Studio currently only supports SSO authentication.",
			Type:        types.StringType,
			Computed:    true,
		},
		"default_s3_location": {
			// Property: DefaultS3Location
			// CloudFormation resource type schema:
			// {
			//   "description": "The default Amazon S3 location to back up EMR Studio Workspaces and notebook files. A Studio user can select an alternative Amazon S3 location when creating a Workspace.",
			//   "maxLength": 10280,
			//   "minLength": 6,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The default Amazon S3 location to back up EMR Studio Workspaces and notebook files. A Studio user can select an alternative Amazon S3 location when creating a Workspace.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "A detailed description of the Studio.",
			//   "maxLength": 256,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "A detailed description of the Studio.",
			Type:        types.StringType,
			Computed:    true,
		},
		"engine_security_group_id": {
			// Property: EngineSecurityGroupId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the Amazon EMR Studio Engine security group. The Engine security group allows inbound network traffic from the Workspace security group, and it must be in the same VPC specified by VpcId.",
			//   "maxLength": 256,
			//   "minLength": 4,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The ID of the Amazon EMR Studio Engine security group. The Engine security group allows inbound network traffic from the Workspace security group, and it must be in the same VPC specified by VpcId.",
			Type:        types.StringType,
			Computed:    true,
		},
		"idp_auth_url": {
			// Property: IdpAuthUrl
			// CloudFormation resource type schema:
			// {
			//   "description": "Your identity provider's authentication endpoint. Amazon EMR Studio redirects federated users to this endpoint for authentication when logging in to a Studio with the Studio URL.",
			//   "maxLength": 4096,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Your identity provider's authentication endpoint. Amazon EMR Studio redirects federated users to this endpoint for authentication when logging in to a Studio with the Studio URL.",
			Type:        types.StringType,
			Computed:    true,
		},
		"idp_relay_state_parameter_name": {
			// Property: IdpRelayStateParameterName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of relay state parameter for external Identity Provider.",
			//   "maxLength": 256,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "The name of relay state parameter for external Identity Provider.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "A descriptive name for the Amazon EMR Studio.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "A descriptive name for the Amazon EMR Studio.",
			Type:        types.StringType,
			Computed:    true,
		},
		"service_role": {
			// Property: ServiceRole
			// CloudFormation resource type schema:
			// {
			//   "description": "The IAM role that will be assumed by the Amazon EMR Studio. The service role provides a way for Amazon EMR Studio to interoperate with other AWS services.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The IAM role that will be assumed by the Amazon EMR Studio. The service role provides a way for Amazon EMR Studio to interoperate with other AWS services.",
			Type:        types.StringType,
			Computed:    true,
		},
		"studio_id": {
			// Property: StudioId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the EMR Studio.",
			//   "maxLength": 256,
			//   "minLength": 4,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The ID of the EMR Studio.",
			Type:        types.StringType,
			Computed:    true,
		},
		"subnet_ids": {
			// Property: SubnetIds
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of up to 5 subnet IDs to associate with the Studio. The subnets must belong to the VPC specified by VpcId. Studio users can create a Workspace in any of the specified subnets.",
			//   "items": {
			//     "description": "Identifier of a subnet",
			//     "pattern": "",
			//     "type": "string"
			//   },
			//   "minItems": 1,
			//   "type": "array"
			// }
			Description: "A list of up to 5 subnet IDs to associate with the Studio. The subnets must belong to the VPC specified by VpcId. Studio users can create a Workspace in any of the specified subnets.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of tags to associate with the Studio. Tags are user-defined key-value pairs that consist of a required key string with a maximum of 128 characters, and an optional value string with a maximum of 256 characters.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "An arbitrary set of tags (key-value pairs) for this EMR Studio.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 255 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 256,
			//         "minLength": 0,
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
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A list of tags to associate with the Studio. Tags are user-defined key-value pairs that consist of a required key string with a maximum of 128 characters, and an optional value string with a maximum of 256 characters.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 255 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				tfsdk.SetNestedAttributesOptions{},
			),
			Computed: true,
		},
		"url": {
			// Property: Url
			// CloudFormation resource type schema:
			// {
			//   "description": "The unique Studio access URL.",
			//   "maxLength": 4096,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The unique Studio access URL.",
			Type:        types.StringType,
			Computed:    true,
		},
		"user_role": {
			// Property: UserRole
			// CloudFormation resource type schema:
			// {
			//   "description": "The IAM user role that will be assumed by users and groups logged in to a Studio. The permissions attached to this IAM role can be scoped down for each user or group using session policies.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The IAM user role that will be assumed by users and groups logged in to a Studio. The permissions attached to this IAM role can be scoped down for each user or group using session policies.",
			Type:        types.StringType,
			Computed:    true,
		},
		"vpc_id": {
			// Property: VpcId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the Amazon Virtual Private Cloud (Amazon VPC) to associate with the Studio.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The ID of the Amazon Virtual Private Cloud (Amazon VPC) to associate with the Studio.",
			Type:        types.StringType,
			Computed:    true,
		},
		"workspace_security_group_id": {
			// Property: WorkspaceSecurityGroupId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the Amazon EMR Studio Workspace security group. The Workspace security group allows outbound network traffic to resources in the Engine security group, and it must be in the same VPC specified by VpcId.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The ID of the Amazon EMR Studio Workspace security group. The Workspace security group allows outbound network traffic to resources in the Engine security group, and it must be in the same VPC specified by VpcId.",
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
		Description: "Data Source schema for AWS::EMR::Studio",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EMR::Studio").WithTerraformTypeName("awscc_emr_studio")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                            "Arn",
		"auth_mode":                      "AuthMode",
		"default_s3_location":            "DefaultS3Location",
		"description":                    "Description",
		"engine_security_group_id":       "EngineSecurityGroupId",
		"idp_auth_url":                   "IdpAuthUrl",
		"idp_relay_state_parameter_name": "IdpRelayStateParameterName",
		"key":                            "Key",
		"name":                           "Name",
		"service_role":                   "ServiceRole",
		"studio_id":                      "StudioId",
		"subnet_ids":                     "SubnetIds",
		"tags":                           "Tags",
		"url":                            "Url",
		"user_role":                      "UserRole",
		"value":                          "Value",
		"vpc_id":                         "VpcId",
		"workspace_security_group_id":    "WorkspaceSecurityGroupId",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
