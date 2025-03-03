// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package apprunner

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_apprunner_service", serviceDataSourceType)
}

// serviceDataSourceType returns the Terraform awscc_apprunner_service data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::AppRunner::Service resource type.
func serviceDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"auto_scaling_configuration_arn": {
			// Property: AutoScalingConfigurationArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Autoscaling configuration ARN",
			//   "maxLength": 1011,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Autoscaling configuration ARN",
			Type:        types.StringType,
			Computed:    true,
		},
		"encryption_configuration": {
			// Property: EncryptionConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Encryption configuration (KMS key)",
			//   "properties": {
			//     "KmsKey": {
			//       "description": "The KMS Key",
			//       "maxLength": 256,
			//       "minLength": 0,
			//       "pattern": "",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "KmsKey"
			//   ],
			//   "type": "object"
			// }
			Description: "Encryption configuration (KMS key)",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"kms_key": {
						// Property: KmsKey
						Description: "The KMS Key",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"health_check_configuration": {
			// Property: HealthCheckConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Health check configuration",
			//   "properties": {
			//     "HealthyThreshold": {
			//       "description": "Health check Healthy Threshold",
			//       "maximum": 20,
			//       "minimum": 1,
			//       "type": "integer"
			//     },
			//     "Interval": {
			//       "description": "Health check Interval",
			//       "type": "integer"
			//     },
			//     "Path": {
			//       "description": "Health check Path",
			//       "type": "string"
			//     },
			//     "Protocol": {
			//       "description": "Health Check Protocol",
			//       "enum": [
			//         "TCP",
			//         "HTTP"
			//       ],
			//       "type": "string"
			//     },
			//     "Timeout": {
			//       "description": "Health check Timeout",
			//       "maximum": 20,
			//       "minimum": 1,
			//       "type": "integer"
			//     },
			//     "UnhealthyThreshold": {
			//       "description": "Health check Unhealthy Threshold",
			//       "maximum": 20,
			//       "minimum": 1,
			//       "type": "integer"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Health check configuration",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"healthy_threshold": {
						// Property: HealthyThreshold
						Description: "Health check Healthy Threshold",
						Type:        types.NumberType,
						Computed:    true,
					},
					"interval": {
						// Property: Interval
						Description: "Health check Interval",
						Type:        types.NumberType,
						Computed:    true,
					},
					"path": {
						// Property: Path
						Description: "Health check Path",
						Type:        types.StringType,
						Computed:    true,
					},
					"protocol": {
						// Property: Protocol
						Description: "Health Check Protocol",
						Type:        types.StringType,
						Computed:    true,
					},
					"timeout": {
						// Property: Timeout
						Description: "Health check Timeout",
						Type:        types.NumberType,
						Computed:    true,
					},
					"unhealthy_threshold": {
						// Property: UnhealthyThreshold
						Description: "Health check Unhealthy Threshold",
						Type:        types.NumberType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"instance_configuration": {
			// Property: InstanceConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Instance Configuration",
			//   "properties": {
			//     "Cpu": {
			//       "description": "CPU",
			//       "maxLength": 6,
			//       "minLength": 4,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "InstanceRoleArn": {
			//       "description": "Instance Role Arn",
			//       "maxLength": 102,
			//       "minLength": 29,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "Memory": {
			//       "description": "Memory",
			//       "maxLength": 4,
			//       "minLength": 4,
			//       "pattern": "",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Instance Configuration",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"cpu": {
						// Property: Cpu
						Description: "CPU",
						Type:        types.StringType,
						Computed:    true,
					},
					"instance_role_arn": {
						// Property: InstanceRoleArn
						Description: "Instance Role Arn",
						Type:        types.StringType,
						Computed:    true,
					},
					"memory": {
						// Property: Memory
						Description: "Memory",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"service_arn": {
			// Property: ServiceArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the AppRunner Service.",
			//   "maxLength": 1011,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the AppRunner Service.",
			Type:        types.StringType,
			Computed:    true,
		},
		"service_id": {
			// Property: ServiceId
			// CloudFormation resource type schema:
			// {
			//   "description": "The AppRunner Service Id",
			//   "maxLength": 32,
			//   "minLength": 32,
			//   "type": "string"
			// }
			Description: "The AppRunner Service Id",
			Type:        types.StringType,
			Computed:    true,
		},
		"service_name": {
			// Property: ServiceName
			// CloudFormation resource type schema:
			// {
			//   "description": "The AppRunner Service Name.",
			//   "maxLength": 40,
			//   "minLength": 4,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The AppRunner Service Name.",
			Type:        types.StringType,
			Computed:    true,
		},
		"service_url": {
			// Property: ServiceUrl
			// CloudFormation resource type schema:
			// {
			//   "description": "The Service Url of the AppRunner Service.",
			//   "type": "string"
			// }
			Description: "The Service Url of the AppRunner Service.",
			Type:        types.StringType,
			Computed:    true,
		},
		"source_configuration": {
			// Property: SourceConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Source Code configuration",
			//   "properties": {
			//     "AuthenticationConfiguration": {
			//       "additionalProperties": false,
			//       "description": "Authentication Configuration",
			//       "properties": {
			//         "AccessRoleArn": {
			//           "description": "Access Role Arn",
			//           "maxLength": 102,
			//           "minLength": 29,
			//           "pattern": "",
			//           "type": "string"
			//         },
			//         "ConnectionArn": {
			//           "description": "Connection Arn",
			//           "maxLength": 1011,
			//           "minLength": 1,
			//           "pattern": "",
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "AutoDeploymentsEnabled": {
			//       "description": "Auto Deployment enabled",
			//       "type": "boolean"
			//     },
			//     "CodeRepository": {
			//       "additionalProperties": false,
			//       "description": "Source Code Repository",
			//       "properties": {
			//         "CodeConfiguration": {
			//           "additionalProperties": false,
			//           "description": "Code Configuration",
			//           "properties": {
			//             "CodeConfigurationValues": {
			//               "additionalProperties": false,
			//               "description": "Code Configuration Values",
			//               "properties": {
			//                 "BuildCommand": {
			//                   "description": "Build Command",
			//                   "type": "string"
			//                 },
			//                 "Port": {
			//                   "description": "Port",
			//                   "type": "string"
			//                 },
			//                 "Runtime": {
			//                   "description": "Runtime",
			//                   "enum": [
			//                     "PYTHON_3",
			//                     "NODEJS_12"
			//                   ],
			//                   "type": "string"
			//                 },
			//                 "RuntimeEnvironmentVariables": {
			//                   "items": {
			//                     "additionalProperties": false,
			//                     "properties": {
			//                       "Name": {
			//                         "type": "string"
			//                       },
			//                       "Value": {
			//                         "type": "string"
			//                       }
			//                     },
			//                     "type": "object"
			//                   },
			//                   "type": "array"
			//                 },
			//                 "StartCommand": {
			//                   "description": "Start Command",
			//                   "type": "string"
			//                 }
			//               },
			//               "required": [
			//                 "Runtime"
			//               ],
			//               "type": "object"
			//             },
			//             "ConfigurationSource": {
			//               "description": "Configuration Source",
			//               "enum": [
			//                 "REPOSITORY",
			//                 "API"
			//               ],
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "ConfigurationSource"
			//           ],
			//           "type": "object"
			//         },
			//         "RepositoryUrl": {
			//           "description": "Repository Url",
			//           "type": "string"
			//         },
			//         "SourceCodeVersion": {
			//           "additionalProperties": false,
			//           "description": "Source Code Version",
			//           "properties": {
			//             "Type": {
			//               "description": "Source Code Version Type",
			//               "enum": [
			//                 "BRANCH"
			//               ],
			//               "type": "string"
			//             },
			//             "Value": {
			//               "description": "Source Code Version Value",
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "Type",
			//             "Value"
			//           ],
			//           "type": "object"
			//         }
			//       },
			//       "required": [
			//         "RepositoryUrl",
			//         "SourceCodeVersion"
			//       ],
			//       "type": "object"
			//     },
			//     "ImageRepository": {
			//       "additionalProperties": false,
			//       "description": "Image Repository",
			//       "properties": {
			//         "ImageConfiguration": {
			//           "additionalProperties": false,
			//           "description": "Image Configuration",
			//           "properties": {
			//             "Port": {
			//               "description": "Port",
			//               "type": "string"
			//             },
			//             "RuntimeEnvironmentVariables": {
			//               "items": {
			//                 "additionalProperties": false,
			//                 "properties": {
			//                   "Name": {
			//                     "type": "string"
			//                   },
			//                   "Value": {
			//                     "type": "string"
			//                   }
			//                 },
			//                 "type": "object"
			//               },
			//               "type": "array"
			//             },
			//             "StartCommand": {
			//               "description": "Start Command",
			//               "type": "string"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "ImageIdentifier": {
			//           "description": "Image Identifier",
			//           "maxLength": 1024,
			//           "minLength": 1,
			//           "pattern": "",
			//           "type": "string"
			//         },
			//         "ImageRepositoryType": {
			//           "description": "Image Repository Type",
			//           "enum": [
			//             "ECR",
			//             "ECR_PUBLIC"
			//           ],
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "ImageIdentifier",
			//         "ImageRepositoryType"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Source Code configuration",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"authentication_configuration": {
						// Property: AuthenticationConfiguration
						Description: "Authentication Configuration",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"access_role_arn": {
									// Property: AccessRoleArn
									Description: "Access Role Arn",
									Type:        types.StringType,
									Computed:    true,
								},
								"connection_arn": {
									// Property: ConnectionArn
									Description: "Connection Arn",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"auto_deployments_enabled": {
						// Property: AutoDeploymentsEnabled
						Description: "Auto Deployment enabled",
						Type:        types.BoolType,
						Computed:    true,
					},
					"code_repository": {
						// Property: CodeRepository
						Description: "Source Code Repository",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"code_configuration": {
									// Property: CodeConfiguration
									Description: "Code Configuration",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"code_configuration_values": {
												// Property: CodeConfigurationValues
												Description: "Code Configuration Values",
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"build_command": {
															// Property: BuildCommand
															Description: "Build Command",
															Type:        types.StringType,
															Computed:    true,
														},
														"port": {
															// Property: Port
															Description: "Port",
															Type:        types.StringType,
															Computed:    true,
														},
														"runtime": {
															// Property: Runtime
															Description: "Runtime",
															Type:        types.StringType,
															Computed:    true,
														},
														"runtime_environment_variables": {
															// Property: RuntimeEnvironmentVariables
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
														"start_command": {
															// Property: StartCommand
															Description: "Start Command",
															Type:        types.StringType,
															Computed:    true,
														},
													},
												),
												Computed: true,
											},
											"configuration_source": {
												// Property: ConfigurationSource
												Description: "Configuration Source",
												Type:        types.StringType,
												Computed:    true,
											},
										},
									),
									Computed: true,
								},
								"repository_url": {
									// Property: RepositoryUrl
									Description: "Repository Url",
									Type:        types.StringType,
									Computed:    true,
								},
								"source_code_version": {
									// Property: SourceCodeVersion
									Description: "Source Code Version",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"type": {
												// Property: Type
												Description: "Source Code Version Type",
												Type:        types.StringType,
												Computed:    true,
											},
											"value": {
												// Property: Value
												Description: "Source Code Version Value",
												Type:        types.StringType,
												Computed:    true,
											},
										},
									),
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"image_repository": {
						// Property: ImageRepository
						Description: "Image Repository",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"image_configuration": {
									// Property: ImageConfiguration
									Description: "Image Configuration",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"port": {
												// Property: Port
												Description: "Port",
												Type:        types.StringType,
												Computed:    true,
											},
											"runtime_environment_variables": {
												// Property: RuntimeEnvironmentVariables
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
											"start_command": {
												// Property: StartCommand
												Description: "Start Command",
												Type:        types.StringType,
												Computed:    true,
											},
										},
									),
									Computed: true,
								},
								"image_identifier": {
									// Property: ImageIdentifier
									Description: "Image Identifier",
									Type:        types.StringType,
									Computed:    true,
								},
								"image_repository_type": {
									// Property: ImageRepositoryType
									Description: "Image Repository Type",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "description": "AppRunner Service status.",
			//   "type": "string"
			// }
			Description: "AppRunner Service status.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
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
		Description: "Data Source schema for AWS::AppRunner::Service",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::AppRunner::Service").WithTerraformTypeName("awscc_apprunner_service")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_role_arn":                "AccessRoleArn",
		"authentication_configuration":   "AuthenticationConfiguration",
		"auto_deployments_enabled":       "AutoDeploymentsEnabled",
		"auto_scaling_configuration_arn": "AutoScalingConfigurationArn",
		"build_command":                  "BuildCommand",
		"code_configuration":             "CodeConfiguration",
		"code_configuration_values":      "CodeConfigurationValues",
		"code_repository":                "CodeRepository",
		"configuration_source":           "ConfigurationSource",
		"connection_arn":                 "ConnectionArn",
		"cpu":                            "Cpu",
		"encryption_configuration":       "EncryptionConfiguration",
		"health_check_configuration":     "HealthCheckConfiguration",
		"healthy_threshold":              "HealthyThreshold",
		"image_configuration":            "ImageConfiguration",
		"image_identifier":               "ImageIdentifier",
		"image_repository":               "ImageRepository",
		"image_repository_type":          "ImageRepositoryType",
		"instance_configuration":         "InstanceConfiguration",
		"instance_role_arn":              "InstanceRoleArn",
		"interval":                       "Interval",
		"key":                            "Key",
		"kms_key":                        "KmsKey",
		"memory":                         "Memory",
		"name":                           "Name",
		"path":                           "Path",
		"port":                           "Port",
		"protocol":                       "Protocol",
		"repository_url":                 "RepositoryUrl",
		"runtime":                        "Runtime",
		"runtime_environment_variables":  "RuntimeEnvironmentVariables",
		"service_arn":                    "ServiceArn",
		"service_id":                     "ServiceId",
		"service_name":                   "ServiceName",
		"service_url":                    "ServiceUrl",
		"source_code_version":            "SourceCodeVersion",
		"source_configuration":           "SourceConfiguration",
		"start_command":                  "StartCommand",
		"status":                         "Status",
		"tags":                           "Tags",
		"timeout":                        "Timeout",
		"type":                           "Type",
		"unhealthy_threshold":            "UnhealthyThreshold",
		"value":                          "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
