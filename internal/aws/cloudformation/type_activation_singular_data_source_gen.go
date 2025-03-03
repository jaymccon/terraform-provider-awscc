// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_cloudformation_type_activation", typeActivationDataSourceType)
}

// typeActivationDataSourceType returns the Terraform awscc_cloudformation_type_activation data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::CloudFormation::TypeActivation resource type.
func typeActivationDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the extension.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the extension.",
			Type:        types.StringType,
			Computed:    true,
		},
		"auto_update": {
			// Property: AutoUpdate
			// CloudFormation resource type schema:
			// {
			//   "description": "Whether to automatically update the extension in this account and region when a new minor version is published by the extension publisher. Major versions released by the publisher must be manually updated.",
			//   "type": "boolean"
			// }
			Description: "Whether to automatically update the extension in this account and region when a new minor version is published by the extension publisher. Major versions released by the publisher must be manually updated.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"execution_role_arn": {
			// Property: ExecutionRoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the IAM execution role to use to register the type. If your resource type calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. CloudFormation then assumes that execution role to provide your resource type with the appropriate credentials.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the IAM execution role to use to register the type. If your resource type calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. CloudFormation then assumes that execution role to provide your resource type with the appropriate credentials.",
			Type:        types.StringType,
			Computed:    true,
		},
		"logging_config": {
			// Property: LoggingConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Specifies logging configuration information for a type.",
			//   "properties": {
			//     "LogGroupName": {
			//       "description": "The Amazon CloudWatch log group to which CloudFormation sends error logging information when invoking the type's handlers.",
			//       "maxLength": 512,
			//       "minLength": 1,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "LogRoleArn": {
			//       "description": "The ARN of the role that CloudFormation should assume when sending log entries to CloudWatch logs.",
			//       "maxLength": 256,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Specifies logging configuration information for a type.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"log_group_name": {
						// Property: LogGroupName
						Description: "The Amazon CloudWatch log group to which CloudFormation sends error logging information when invoking the type's handlers.",
						Type:        types.StringType,
						Computed:    true,
					},
					"log_role_arn": {
						// Property: LogRoleArn
						Description: "The ARN of the role that CloudFormation should assume when sending log entries to CloudWatch logs.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"major_version": {
			// Property: MajorVersion
			// CloudFormation resource type schema:
			// {
			//   "description": "The Major Version of the type you want to enable",
			//   "maxLength": 100000,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The Major Version of the type you want to enable",
			Type:        types.StringType,
			Computed:    true,
		},
		"public_type_arn": {
			// Property: PublicTypeArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Number (ARN) assigned to the public extension upon publication",
			//   "maxLength": 1024,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Number (ARN) assigned to the public extension upon publication",
			Type:        types.StringType,
			Computed:    true,
		},
		"publisher_id": {
			// Property: PublisherId
			// CloudFormation resource type schema:
			// {
			//   "description": "The publisher id assigned by CloudFormation for publishing in this region.",
			//   "maxLength": 40,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The publisher id assigned by CloudFormation for publishing in this region.",
			Type:        types.StringType,
			Computed:    true,
		},
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			// {
			//   "description": "The kind of extension",
			//   "enum": [
			//     "RESOURCE",
			//     "MODULE"
			//   ],
			//   "type": "string"
			// }
			Description: "The kind of extension",
			Type:        types.StringType,
			Computed:    true,
		},
		"type_name": {
			// Property: TypeName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the type being registered.\n\nWe recommend that type names adhere to the following pattern: company_or_organization::service::type.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the type being registered.\n\nWe recommend that type names adhere to the following pattern: company_or_organization::service::type.",
			Type:        types.StringType,
			Computed:    true,
		},
		"type_name_alias": {
			// Property: TypeNameAlias
			// CloudFormation resource type schema:
			// {
			//   "description": "An alias to assign to the public extension in this account and region. If you specify an alias for the extension, you must then use the alias to refer to the extension in your templates.",
			//   "maxLength": 204,
			//   "minLength": 10,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "An alias to assign to the public extension in this account and region. If you specify an alias for the extension, you must then use the alias to refer to the extension in your templates.",
			Type:        types.StringType,
			Computed:    true,
		},
		"version_bump": {
			// Property: VersionBump
			// CloudFormation resource type schema:
			// {
			//   "description": "Manually updates a previously-enabled type to a new major or minor version, if available. You can also use this parameter to update the value of AutoUpdateEnabled",
			//   "enum": [
			//     "MAJOR",
			//     "MINOR"
			//   ],
			//   "type": "string"
			// }
			Description: "Manually updates a previously-enabled type to a new major or minor version, if available. You can also use this parameter to update the value of AutoUpdateEnabled",
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
		Description: "Data Source schema for AWS::CloudFormation::TypeActivation",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFormation::TypeActivation").WithTerraformTypeName("awscc_cloudformation_type_activation")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                "Arn",
		"auto_update":        "AutoUpdate",
		"execution_role_arn": "ExecutionRoleArn",
		"log_group_name":     "LogGroupName",
		"log_role_arn":       "LogRoleArn",
		"logging_config":     "LoggingConfig",
		"major_version":      "MajorVersion",
		"public_type_arn":    "PublicTypeArn",
		"publisher_id":       "PublisherId",
		"type":               "Type",
		"type_name":          "TypeName",
		"type_name_alias":    "TypeNameAlias",
		"version_bump":       "VersionBump",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
