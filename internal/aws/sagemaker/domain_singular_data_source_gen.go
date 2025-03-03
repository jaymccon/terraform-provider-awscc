// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_sagemaker_domain", domainDataSourceType)
}

// domainDataSourceType returns the Terraform awscc_sagemaker_domain data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::SageMaker::Domain resource type.
func domainDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"app_network_access_type": {
			// Property: AppNetworkAccessType
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the VPC used for non-EFS traffic. The default value is PublicInternetOnly.",
			//   "enum": [
			//     "PublicInternetOnly",
			//     "VpcOnly"
			//   ],
			//   "type": "string"
			// }
			Description: "Specifies the VPC used for non-EFS traffic. The default value is PublicInternetOnly.",
			Type:        types.StringType,
			Computed:    true,
		},
		"auth_mode": {
			// Property: AuthMode
			// CloudFormation resource type schema:
			// {
			//   "description": "The mode of authentication that members use to access the domain.",
			//   "enum": [
			//     "SSO",
			//     "IAM"
			//   ],
			//   "type": "string"
			// }
			Description: "The mode of authentication that members use to access the domain.",
			Type:        types.StringType,
			Computed:    true,
		},
		"default_user_settings": {
			// Property: DefaultUserSettings
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The default user settings.",
			//   "properties": {
			//     "ExecutionRole": {
			//       "description": "The user profile Amazon Resource Name (ARN).",
			//       "maxLength": 2048,
			//       "minLength": 20,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "JupyterServerAppSettings": {
			//       "additionalProperties": false,
			//       "description": "The Jupyter server's app settings.",
			//       "properties": {
			//         "DefaultResourceSpec": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "InstanceType": {
			//               "description": "The instance type that the image version runs on.",
			//               "enum": [
			//                 "system",
			//                 "ml.t3.micro",
			//                 "ml.t3.small",
			//                 "ml.t3.medium",
			//                 "ml.t3.large",
			//                 "ml.t3.xlarge",
			//                 "ml.t3.2xlarge",
			//                 "ml.m5.large",
			//                 "ml.m5.xlarge",
			//                 "ml.m5.2xlarge",
			//                 "ml.m5.4xlarge",
			//                 "ml.m5.8xlarge",
			//                 "ml.m5.12xlarge",
			//                 "ml.m5.16xlarge",
			//                 "ml.m5.24xlarge",
			//                 "ml.c5.large",
			//                 "ml.c5.xlarge",
			//                 "ml.c5.2xlarge",
			//                 "ml.c5.4xlarge",
			//                 "ml.c5.9xlarge",
			//                 "ml.c5.12xlarge",
			//                 "ml.c5.18xlarge",
			//                 "ml.c5.24xlarge",
			//                 "ml.p3.2xlarge",
			//                 "ml.p3.8xlarge",
			//                 "ml.p3.16xlarge",
			//                 "ml.g4dn.xlarge",
			//                 "ml.g4dn.2xlarge",
			//                 "ml.g4dn.4xlarge",
			//                 "ml.g4dn.8xlarge",
			//                 "ml.g4dn.12xlarge",
			//                 "ml.g4dn.16xlarge"
			//               ],
			//               "type": "string"
			//             },
			//             "SageMakerImageArn": {
			//               "description": "The ARN of the SageMaker image that the image version belongs to.",
			//               "maxLength": 256,
			//               "pattern": "",
			//               "type": "string"
			//             },
			//             "SageMakerImageVersionArn": {
			//               "description": "The ARN of the image version created on the instance.",
			//               "maxLength": 256,
			//               "pattern": "",
			//               "type": "string"
			//             }
			//           },
			//           "type": "object"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "KernelGatewayAppSettings": {
			//       "additionalProperties": false,
			//       "description": "The kernel gateway app settings.",
			//       "properties": {
			//         "CustomImages": {
			//           "description": "A list of custom SageMaker images that are configured to run as a KernelGateway app.",
			//           "items": {
			//             "additionalProperties": false,
			//             "description": "A custom SageMaker image.",
			//             "properties": {
			//               "AppImageConfigName": {
			//                 "description": "The Name of the AppImageConfig.",
			//                 "maxLength": 63,
			//                 "pattern": "",
			//                 "type": "string"
			//               },
			//               "ImageName": {
			//                 "description": "The name of the CustomImage. Must be unique to your account.",
			//                 "maxLength": 63,
			//                 "pattern": "",
			//                 "type": "string"
			//               },
			//               "ImageVersionNumber": {
			//                 "description": "The version number of the CustomImage.",
			//                 "minimum": 0,
			//                 "type": "integer"
			//               }
			//             },
			//             "required": [
			//               "AppImageConfigName",
			//               "ImageName"
			//             ],
			//             "type": "object"
			//           },
			//           "maxItems": 30,
			//           "minItems": 0,
			//           "type": "array",
			//           "uniqueItems": false
			//         },
			//         "DefaultResourceSpec": {
			//           "additionalProperties": false,
			//           "description": "The default instance type and the Amazon Resource Name (ARN) of the default SageMaker image used by the KernelGateway app.",
			//           "properties": {
			//             "InstanceType": {
			//               "description": "The instance type that the image version runs on.",
			//               "enum": [
			//                 "system",
			//                 "ml.t3.micro",
			//                 "ml.t3.small",
			//                 "ml.t3.medium",
			//                 "ml.t3.large",
			//                 "ml.t3.xlarge",
			//                 "ml.t3.2xlarge",
			//                 "ml.m5.large",
			//                 "ml.m5.xlarge",
			//                 "ml.m5.2xlarge",
			//                 "ml.m5.4xlarge",
			//                 "ml.m5.8xlarge",
			//                 "ml.m5.12xlarge",
			//                 "ml.m5.16xlarge",
			//                 "ml.m5.24xlarge",
			//                 "ml.c5.large",
			//                 "ml.c5.xlarge",
			//                 "ml.c5.2xlarge",
			//                 "ml.c5.4xlarge",
			//                 "ml.c5.9xlarge",
			//                 "ml.c5.12xlarge",
			//                 "ml.c5.18xlarge",
			//                 "ml.c5.24xlarge",
			//                 "ml.p3.2xlarge",
			//                 "ml.p3.8xlarge",
			//                 "ml.p3.16xlarge",
			//                 "ml.g4dn.xlarge",
			//                 "ml.g4dn.2xlarge",
			//                 "ml.g4dn.4xlarge",
			//                 "ml.g4dn.8xlarge",
			//                 "ml.g4dn.12xlarge",
			//                 "ml.g4dn.16xlarge"
			//               ],
			//               "type": "string"
			//             },
			//             "SageMakerImageArn": {
			//               "description": "The ARN of the SageMaker image that the image version belongs to.",
			//               "maxLength": 256,
			//               "pattern": "",
			//               "type": "string"
			//             },
			//             "SageMakerImageVersionArn": {
			//               "description": "The ARN of the image version created on the instance.",
			//               "maxLength": 256,
			//               "pattern": "",
			//               "type": "string"
			//             }
			//           },
			//           "type": "object"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "SecurityGroups": {
			//       "description": "The security groups for the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.",
			//       "items": {
			//         "maxLength": 32,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "maxItems": 5,
			//       "minItems": 0,
			//       "type": "array",
			//       "uniqueItems": false
			//     },
			//     "SharingSettings": {
			//       "additionalProperties": false,
			//       "description": "The sharing settings.",
			//       "properties": {
			//         "NotebookOutputOption": {
			//           "description": "Whether to include the notebook cell output when sharing the notebook. The default is Disabled.",
			//           "enum": [
			//             "Allowed",
			//             "Disabled"
			//           ],
			//           "type": "string"
			//         },
			//         "S3KmsKeyId": {
			//           "description": "When NotebookOutputOption is Allowed, the AWS Key Management Service (KMS) encryption key ID used to encrypt the notebook cell output in the Amazon S3 bucket.",
			//           "maxLength": 2048,
			//           "pattern": "",
			//           "type": "string"
			//         },
			//         "S3OutputPath": {
			//           "description": "When NotebookOutputOption is Allowed, the Amazon S3 bucket used to store the shared notebook snapshots.",
			//           "maxLength": 1024,
			//           "pattern": "",
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The default user settings.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"execution_role": {
						// Property: ExecutionRole
						Description: "The user profile Amazon Resource Name (ARN).",
						Type:        types.StringType,
						Computed:    true,
					},
					"jupyter_server_app_settings": {
						// Property: JupyterServerAppSettings
						Description: "The Jupyter server's app settings.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"default_resource_spec": {
									// Property: DefaultResourceSpec
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"instance_type": {
												// Property: InstanceType
												Description: "The instance type that the image version runs on.",
												Type:        types.StringType,
												Computed:    true,
											},
											"sage_maker_image_arn": {
												// Property: SageMakerImageArn
												Description: "The ARN of the SageMaker image that the image version belongs to.",
												Type:        types.StringType,
												Computed:    true,
											},
											"sage_maker_image_version_arn": {
												// Property: SageMakerImageVersionArn
												Description: "The ARN of the image version created on the instance.",
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
					"kernel_gateway_app_settings": {
						// Property: KernelGatewayAppSettings
						Description: "The kernel gateway app settings.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"custom_images": {
									// Property: CustomImages
									Description: "A list of custom SageMaker images that are configured to run as a KernelGateway app.",
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"app_image_config_name": {
												// Property: AppImageConfigName
												Description: "The Name of the AppImageConfig.",
												Type:        types.StringType,
												Computed:    true,
											},
											"image_name": {
												// Property: ImageName
												Description: "The name of the CustomImage. Must be unique to your account.",
												Type:        types.StringType,
												Computed:    true,
											},
											"image_version_number": {
												// Property: ImageVersionNumber
												Description: "The version number of the CustomImage.",
												Type:        types.NumberType,
												Computed:    true,
											},
										},
										tfsdk.ListNestedAttributesOptions{},
									),
									Computed: true,
								},
								"default_resource_spec": {
									// Property: DefaultResourceSpec
									Description: "The default instance type and the Amazon Resource Name (ARN) of the default SageMaker image used by the KernelGateway app.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"instance_type": {
												// Property: InstanceType
												Description: "The instance type that the image version runs on.",
												Type:        types.StringType,
												Computed:    true,
											},
											"sage_maker_image_arn": {
												// Property: SageMakerImageArn
												Description: "The ARN of the SageMaker image that the image version belongs to.",
												Type:        types.StringType,
												Computed:    true,
											},
											"sage_maker_image_version_arn": {
												// Property: SageMakerImageVersionArn
												Description: "The ARN of the image version created on the instance.",
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
					"security_groups": {
						// Property: SecurityGroups
						Description: "The security groups for the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.",
						Type:        types.ListType{ElemType: types.StringType},
						Computed:    true,
					},
					"sharing_settings": {
						// Property: SharingSettings
						Description: "The sharing settings.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"notebook_output_option": {
									// Property: NotebookOutputOption
									Description: "Whether to include the notebook cell output when sharing the notebook. The default is Disabled.",
									Type:        types.StringType,
									Computed:    true,
								},
								"s3_kms_key_id": {
									// Property: S3KmsKeyId
									Description: "When NotebookOutputOption is Allowed, the AWS Key Management Service (KMS) encryption key ID used to encrypt the notebook cell output in the Amazon S3 bucket.",
									Type:        types.StringType,
									Computed:    true,
								},
								"s3_output_path": {
									// Property: S3OutputPath
									Description: "When NotebookOutputOption is Allowed, the Amazon S3 bucket used to store the shared notebook snapshots.",
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
		"domain_arn": {
			// Property: DomainArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the created domain.",
			//   "maxLength": 256,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the created domain.",
			Type:        types.StringType,
			Computed:    true,
		},
		"domain_id": {
			// Property: DomainId
			// CloudFormation resource type schema:
			// {
			//   "description": "The domain name.",
			//   "maxLength": 63,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The domain name.",
			Type:        types.StringType,
			Computed:    true,
		},
		"domain_name": {
			// Property: DomainName
			// CloudFormation resource type schema:
			// {
			//   "description": "A name for the domain.",
			//   "maxLength": 63,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "A name for the domain.",
			Type:        types.StringType,
			Computed:    true,
		},
		"home_efs_file_system_id": {
			// Property: HomeEfsFileSystemId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the Amazon Elastic File System (EFS) managed by this Domain.",
			//   "maxLength": 32,
			//   "type": "string"
			// }
			Description: "The ID of the Amazon Elastic File System (EFS) managed by this Domain.",
			Type:        types.StringType,
			Computed:    true,
		},
		"kms_key_id": {
			// Property: KmsKeyId
			// CloudFormation resource type schema:
			// {
			//   "description": "SageMaker uses AWS KMS to encrypt the EFS volume attached to the domain with an AWS managed customer master key (CMK) by default.",
			//   "maxLength": 2048,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "SageMaker uses AWS KMS to encrypt the EFS volume attached to the domain with an AWS managed customer master key (CMK) by default.",
			Type:        types.StringType,
			Computed:    true,
		},
		"single_sign_on_managed_application_instance_id": {
			// Property: SingleSignOnManagedApplicationInstanceId
			// CloudFormation resource type schema:
			// {
			//   "description": "The SSO managed application instance ID.",
			//   "maxLength": 256,
			//   "type": "string"
			// }
			Description: "The SSO managed application instance ID.",
			Type:        types.StringType,
			Computed:    true,
		},
		"subnet_ids": {
			// Property: SubnetIds
			// CloudFormation resource type schema:
			// {
			//   "description": "The VPC subnets that Studio uses for communication.",
			//   "items": {
			//     "maxLength": 32,
			//     "pattern": "",
			//     "type": "string"
			//   },
			//   "maxItems": 16,
			//   "minItems": 1,
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "The VPC subnets that Studio uses for communication.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of tags to apply to the user profile.",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "minItems": 0,
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "A list of tags to apply to the user profile.",
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
		"url": {
			// Property: Url
			// CloudFormation resource type schema:
			// {
			//   "description": "The URL to the created domain.",
			//   "maxLength": 1024,
			//   "type": "string"
			// }
			Description: "The URL to the created domain.",
			Type:        types.StringType,
			Computed:    true,
		},
		"vpc_id": {
			// Property: VpcId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.",
			//   "maxLength": 32,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The ID of the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.",
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
		Description: "Data Source schema for AWS::SageMaker::Domain",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SageMaker::Domain").WithTerraformTypeName("awscc_sagemaker_domain")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"app_image_config_name":        "AppImageConfigName",
		"app_network_access_type":      "AppNetworkAccessType",
		"auth_mode":                    "AuthMode",
		"custom_images":                "CustomImages",
		"default_resource_spec":        "DefaultResourceSpec",
		"default_user_settings":        "DefaultUserSettings",
		"domain_arn":                   "DomainArn",
		"domain_id":                    "DomainId",
		"domain_name":                  "DomainName",
		"execution_role":               "ExecutionRole",
		"home_efs_file_system_id":      "HomeEfsFileSystemId",
		"image_name":                   "ImageName",
		"image_version_number":         "ImageVersionNumber",
		"instance_type":                "InstanceType",
		"jupyter_server_app_settings":  "JupyterServerAppSettings",
		"kernel_gateway_app_settings":  "KernelGatewayAppSettings",
		"key":                          "Key",
		"kms_key_id":                   "KmsKeyId",
		"notebook_output_option":       "NotebookOutputOption",
		"s3_kms_key_id":                "S3KmsKeyId",
		"s3_output_path":               "S3OutputPath",
		"sage_maker_image_arn":         "SageMakerImageArn",
		"sage_maker_image_version_arn": "SageMakerImageVersionArn",
		"security_groups":              "SecurityGroups",
		"sharing_settings":             "SharingSettings",
		"single_sign_on_managed_application_instance_id": "SingleSignOnManagedApplicationInstanceId",
		"subnet_ids": "SubnetIds",
		"tags":       "Tags",
		"url":        "Url",
		"value":      "Value",
		"vpc_id":     "VpcId",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
