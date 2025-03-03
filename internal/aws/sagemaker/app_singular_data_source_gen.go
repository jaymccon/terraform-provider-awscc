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
	registry.AddDataSourceTypeFactory("awscc_sagemaker_app", appDataSourceType)
}

// appDataSourceType returns the Terraform awscc_sagemaker_app data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::SageMaker::App resource type.
func appDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"app_arn": {
			// Property: AppArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the app.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the app.",
			Type:        types.StringType,
			Computed:    true,
		},
		"app_name": {
			// Property: AppName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the app.",
			//   "maxLength": 63,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the app.",
			Type:        types.StringType,
			Computed:    true,
		},
		"app_type": {
			// Property: AppType
			// CloudFormation resource type schema:
			// {
			//   "description": "The type of app.",
			//   "enum": [
			//     "JupyterServer",
			//     "KernelGateway"
			//   ],
			//   "type": "string"
			// }
			Description: "The type of app.",
			Type:        types.StringType,
			Computed:    true,
		},
		"domain_id": {
			// Property: DomainId
			// CloudFormation resource type schema:
			// {
			//   "description": "The domain ID.",
			//   "maxLength": 63,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The domain ID.",
			Type:        types.StringType,
			Computed:    true,
		},
		"resource_spec": {
			// Property: ResourceSpec
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance.",
			//   "properties": {
			//     "InstanceType": {
			//       "description": "The instance type that the image version runs on.",
			//       "enum": [
			//         "system",
			//         "ml.t3.micro",
			//         "ml.t3.small",
			//         "ml.t3.medium",
			//         "ml.t3.large",
			//         "ml.t3.xlarge",
			//         "ml.t3.2xlarge",
			//         "ml.m5.large",
			//         "ml.m5.xlarge",
			//         "ml.m5.2xlarge",
			//         "ml.m5.4xlarge",
			//         "ml.m5.8xlarge",
			//         "ml.m5.12xlarge",
			//         "ml.m5.16xlarge",
			//         "ml.m5.24xlarge",
			//         "ml.c5.large",
			//         "ml.c5.xlarge",
			//         "ml.c5.2xlarge",
			//         "ml.c5.4xlarge",
			//         "ml.c5.9xlarge",
			//         "ml.c5.12xlarge",
			//         "ml.c5.18xlarge",
			//         "ml.c5.24xlarge",
			//         "ml.p3.2xlarge",
			//         "ml.p3.8xlarge",
			//         "ml.p3.16xlarge",
			//         "ml.g4dn.xlarge",
			//         "ml.g4dn.2xlarge",
			//         "ml.g4dn.4xlarge",
			//         "ml.g4dn.8xlarge",
			//         "ml.g4dn.12xlarge",
			//         "ml.g4dn.16xlarge"
			//       ],
			//       "type": "string"
			//     },
			//     "SageMakerImageArn": {
			//       "description": "The ARN of the SageMaker image that the image version belongs to.",
			//       "maxLength": 256,
			//       "minLength": 1,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "SageMakerImageVersionArn": {
			//       "description": "The ARN of the image version created on the instance.",
			//       "maxLength": 256,
			//       "minLength": 1,
			//       "pattern": "",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance.",
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
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of tags to apply to the app.",
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
			Description: "A list of tags to apply to the app.",
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
		"user_profile_name": {
			// Property: UserProfileName
			// CloudFormation resource type schema:
			// {
			//   "description": "The user profile name.",
			//   "maxLength": 63,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The user profile name.",
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
		Description: "Data Source schema for AWS::SageMaker::App",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SageMaker::App").WithTerraformTypeName("awscc_sagemaker_app")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"app_arn":                      "AppArn",
		"app_name":                     "AppName",
		"app_type":                     "AppType",
		"domain_id":                    "DomainId",
		"instance_type":                "InstanceType",
		"key":                          "Key",
		"resource_spec":                "ResourceSpec",
		"sage_maker_image_arn":         "SageMakerImageArn",
		"sage_maker_image_version_arn": "SageMakerImageVersionArn",
		"tags":                         "Tags",
		"user_profile_name":            "UserProfileName",
		"value":                        "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
