// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ecs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_ecs_cluster", clusterDataSourceType)
}

// clusterDataSourceType returns the Terraform awscc_ecs_cluster data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::ECS::Cluster resource type.
func clusterDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the Amazon ECS cluster, such as arn:aws:ecs:us-east-2:123456789012:cluster/MyECSCluster.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the Amazon ECS cluster, such as arn:aws:ecs:us-east-2:123456789012:cluster/MyECSCluster.",
			Type:        types.StringType,
			Computed:    true,
		},
		"capacity_providers": {
			// Property: CapacityProviders
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Computed: true,
		},
		"cluster_name": {
			// Property: ClusterName
			// CloudFormation resource type schema:
			// {
			//   "description": "A user-generated string that you use to identify your cluster. If you don't specify a name, AWS CloudFormation generates a unique physical ID for the name.",
			//   "type": "string"
			// }
			Description: "A user-generated string that you use to identify your cluster. If you don't specify a name, AWS CloudFormation generates a unique physical ID for the name.",
			Type:        types.StringType,
			Computed:    true,
		},
		"cluster_settings": {
			// Property: ClusterSettings
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "The setting to use when creating a cluster. This parameter is used to enable CloudWatch Container Insights for a cluster. If this value is specified, it will override the containerInsights value set with PutAccountSetting or PutAccountSettingDefault.",
			//     "properties": {
			//       "Name": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"name": {
						// Property: Name
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Computed: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"configuration": {
			// Property: Configuration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The configurations to be set at cluster level.",
			//   "properties": {
			//     "ExecuteCommandConfiguration": {
			//       "additionalProperties": false,
			//       "description": "The configuration for ExecuteCommand.",
			//       "properties": {
			//         "KmsKeyId": {
			//           "type": "string"
			//         },
			//         "LogConfiguration": {
			//           "additionalProperties": false,
			//           "description": "The session logging configuration for ExecuteCommand.",
			//           "properties": {
			//             "CloudWatchEncryptionEnabled": {
			//               "type": "boolean"
			//             },
			//             "CloudWatchLogGroupName": {
			//               "type": "string"
			//             },
			//             "S3BucketName": {
			//               "type": "string"
			//             },
			//             "S3EncryptionEnabled": {
			//               "type": "boolean"
			//             },
			//             "S3KeyPrefix": {
			//               "type": "string"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "Logging": {
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The configurations to be set at cluster level.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"execute_command_configuration": {
						// Property: ExecuteCommandConfiguration
						Description: "The configuration for ExecuteCommand.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"kms_key_id": {
									// Property: KmsKeyId
									Type:     types.StringType,
									Computed: true,
								},
								"log_configuration": {
									// Property: LogConfiguration
									Description: "The session logging configuration for ExecuteCommand.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"cloudwatch_encryption_enabled": {
												// Property: CloudWatchEncryptionEnabled
												Type:     types.BoolType,
												Computed: true,
											},
											"cloudwatch_log_group_name": {
												// Property: CloudWatchLogGroupName
												Type:     types.StringType,
												Computed: true,
											},
											"s3_bucket_name": {
												// Property: S3BucketName
												Type:     types.StringType,
												Computed: true,
											},
											"s3_encryption_enabled": {
												// Property: S3EncryptionEnabled
												Type:     types.BoolType,
												Computed: true,
											},
											"s3_key_prefix": {
												// Property: S3KeyPrefix
												Type:     types.StringType,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"logging": {
									// Property: Logging
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"default_capacity_provider_strategy": {
			// Property: DefaultCapacityProviderStrategy
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A capacity provider strategy consists of one or more capacity providers along with the `base` and `weight` to assign to them. A capacity provider must be associated with the cluster to be used in a capacity provider strategy. The PutClusterCapacityProviders API is used to associate a capacity provider with a cluster. Only capacity providers with an `ACTIVE` or `UPDATING` status can be used.",
			//     "properties": {
			//       "Base": {
			//         "type": "integer"
			//       },
			//       "CapacityProvider": {
			//         "type": "string"
			//       },
			//       "Weight": {
			//         "type": "integer"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"base": {
						// Property: Base
						Type:     types.NumberType,
						Computed: true,
					},
					"capacity_provider": {
						// Property: CapacityProvider
						Type:     types.StringType,
						Computed: true,
					},
					"weight": {
						// Property: Weight
						Type:     types.NumberType,
						Computed: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "The metadata that you apply to the cluster to help you categorize and organize them. Each tag consists of a key and an optional value, both of which you define.",
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Computed: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::ECS::Cluster",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ECS::Cluster").WithTerraformTypeName("awscc_ecs_cluster")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                                "Arn",
		"base":                               "Base",
		"capacity_provider":                  "CapacityProvider",
		"capacity_providers":                 "CapacityProviders",
		"cloudwatch_encryption_enabled":      "CloudWatchEncryptionEnabled",
		"cloudwatch_log_group_name":          "CloudWatchLogGroupName",
		"cluster_name":                       "ClusterName",
		"cluster_settings":                   "ClusterSettings",
		"configuration":                      "Configuration",
		"default_capacity_provider_strategy": "DefaultCapacityProviderStrategy",
		"execute_command_configuration":      "ExecuteCommandConfiguration",
		"key":                                "Key",
		"kms_key_id":                         "KmsKeyId",
		"log_configuration":                  "LogConfiguration",
		"logging":                            "Logging",
		"name":                               "Name",
		"s3_bucket_name":                     "S3BucketName",
		"s3_encryption_enabled":              "S3EncryptionEnabled",
		"s3_key_prefix":                      "S3KeyPrefix",
		"tags":                               "Tags",
		"value":                              "Value",
		"weight":                             "Weight",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
