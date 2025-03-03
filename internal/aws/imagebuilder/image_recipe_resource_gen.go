// Code generated by generators/resource/main.go; DO NOT EDIT.

package imagebuilder

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_imagebuilder_image_recipe", imageRecipeResourceType)
}

// imageRecipeResourceType returns the Terraform awscc_imagebuilder_image_recipe resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ImageBuilder::ImageRecipe resource type.
func imageRecipeResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"additional_instance_configuration": {
			// Property: AdditionalInstanceConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Specify additional settings and launch scripts for your build instances.",
			//   "properties": {
			//     "SystemsManagerAgent": {
			//       "additionalProperties": false,
			//       "description": "Contains settings for the SSM agent on your build instance.",
			//       "properties": {
			//         "UninstallAfterBuild": {
			//           "description": "Controls whether the SSM agent is removed from your final build image, prior to creating the new AMI. If this is set to true, then the agent is removed from the final image. If it's set to false, then the agent is left in, so that it is included in the new AMI. The default value is false.",
			//           "type": "boolean"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "UserDataOverride": {
			//       "description": "Use this property to provide commands or a command script to run when you launch your build instance.",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Specify additional settings and launch scripts for your build instances.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"systems_manager_agent": {
						// Property: SystemsManagerAgent
						Description: "Contains settings for the SSM agent on your build instance.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"uninstall_after_build": {
									// Property: UninstallAfterBuild
									Description: "Controls whether the SSM agent is removed from your final build image, prior to creating the new AMI. If this is set to true, then the agent is removed from the final image. If it's set to false, then the agent is left in, so that it is included in the new AMI. The default value is false.",
									Type:        types.BoolType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
					"user_data_override": {
						// Property: UserDataOverride
						Description: "Use this property to provide commands or a command script to run when you launch your build instance.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the image recipe.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the image recipe.",
			Type:        types.StringType,
			Computed:    true,
		},
		"block_device_mappings": {
			// Property: BlockDeviceMappings
			// CloudFormation resource type schema:
			// {
			//   "description": "The block device mappings to apply when creating images from this recipe.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "Defines block device mappings for the instance used to configure your image. ",
			//     "properties": {
			//       "DeviceName": {
			//         "description": "The device to which these mappings apply.",
			//         "type": "string"
			//       },
			//       "Ebs": {
			//         "additionalProperties": false,
			//         "description": "Use to manage Amazon EBS-specific configuration for this mapping.",
			//         "properties": {
			//           "DeleteOnTermination": {
			//             "description": "Use to configure delete on termination of the associated device.",
			//             "type": "boolean"
			//           },
			//           "Encrypted": {
			//             "description": "Use to configure device encryption.",
			//             "type": "boolean"
			//           },
			//           "Iops": {
			//             "description": "Use to configure device IOPS.",
			//             "type": "integer"
			//           },
			//           "KmsKeyId": {
			//             "description": "Use to configure the KMS key to use when encrypting the device.",
			//             "type": "string"
			//           },
			//           "SnapshotId": {
			//             "description": "The snapshot that defines the device contents.",
			//             "type": "string"
			//           },
			//           "VolumeSize": {
			//             "description": "Use to override the device's volume size.",
			//             "type": "integer"
			//           },
			//           "VolumeType": {
			//             "description": "Use to override the device's volume type.",
			//             "enum": [
			//               "standard",
			//               "io1",
			//               "io2",
			//               "gp2",
			//               "gp3",
			//               "sc1",
			//               "st1"
			//             ],
			//             "type": "string"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "NoDevice": {
			//         "description": "Use to remove a mapping from the parent image.",
			//         "type": "string"
			//       },
			//       "VirtualName": {
			//         "description": "Use to manage instance ephemeral devices.",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "The block device mappings to apply when creating images from this recipe.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"device_name": {
						// Property: DeviceName
						Description: "The device to which these mappings apply.",
						Type:        types.StringType,
						Optional:    true,
					},
					"ebs": {
						// Property: Ebs
						Description: "Use to manage Amazon EBS-specific configuration for this mapping.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"delete_on_termination": {
									// Property: DeleteOnTermination
									Description: "Use to configure delete on termination of the associated device.",
									Type:        types.BoolType,
									Optional:    true,
								},
								"encrypted": {
									// Property: Encrypted
									Description: "Use to configure device encryption.",
									Type:        types.BoolType,
									Optional:    true,
								},
								"iops": {
									// Property: Iops
									Description: "Use to configure device IOPS.",
									Type:        types.NumberType,
									Optional:    true,
								},
								"kms_key_id": {
									// Property: KmsKeyId
									Description: "Use to configure the KMS key to use when encrypting the device.",
									Type:        types.StringType,
									Optional:    true,
								},
								"snapshot_id": {
									// Property: SnapshotId
									Description: "The snapshot that defines the device contents.",
									Type:        types.StringType,
									Optional:    true,
								},
								"volume_size": {
									// Property: VolumeSize
									Description: "Use to override the device's volume size.",
									Type:        types.NumberType,
									Optional:    true,
								},
								"volume_type": {
									// Property: VolumeType
									Description: "Use to override the device's volume type.",
									Type:        types.StringType,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringInSlice([]string{
											"standard",
											"io1",
											"io2",
											"gp2",
											"gp3",
											"sc1",
											"st1",
										}),
									},
								},
							},
						),
						Optional: true,
					},
					"no_device": {
						// Property: NoDevice
						Description: "Use to remove a mapping from the parent image.",
						Type:        types.StringType,
						Optional:    true,
					},
					"virtual_name": {
						// Property: VirtualName
						Description: "Use to manage instance ephemeral devices.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"components": {
			// Property: Components
			// CloudFormation resource type schema:
			// {
			//   "description": "The components of the image recipe.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "Configuration details of the component.",
			//     "properties": {
			//       "ComponentArn": {
			//         "description": "The Amazon Resource Name (ARN) of the component.",
			//         "type": "string"
			//       },
			//       "Parameters": {
			//         "description": "A group of parameter settings that are used to configure the component for a specific recipe.",
			//         "items": {
			//           "additionalProperties": false,
			//           "description": "Contains a key/value pair that sets the named component parameter.",
			//           "properties": {
			//             "Name": {
			//               "description": "The name of the component parameter to set.",
			//               "type": "string"
			//             },
			//             "Value": {
			//               "description": "Sets the value for the named component parameter.",
			//               "items": {
			//                 "type": "string"
			//               },
			//               "type": "array"
			//             }
			//           },
			//           "required": [
			//             "Name",
			//             "Value"
			//           ],
			//           "type": "object"
			//         },
			//         "type": "array"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "The components of the image recipe.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"component_arn": {
						// Property: ComponentArn
						Description: "The Amazon Resource Name (ARN) of the component.",
						Type:        types.StringType,
						Optional:    true,
					},
					"parameters": {
						// Property: Parameters
						Description: "A group of parameter settings that are used to configure the component for a specific recipe.",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"name": {
									// Property: Name
									Description: "The name of the component parameter to set.",
									Type:        types.StringType,
									Required:    true,
								},
								"value": {
									// Property: Value
									Description: "Sets the value for the named component parameter.",
									Type:        types.ListType{ElemType: types.StringType},
									Required:    true,
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Optional: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Required: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the image recipe.",
			//   "type": "string"
			// }
			Description: "The description of the image recipe.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the image recipe.",
			//   "type": "string"
			// }
			Description: "The name of the image recipe.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"parent_image": {
			// Property: ParentImage
			// CloudFormation resource type schema:
			// {
			//   "description": "The parent image of the image recipe.",
			//   "type": "string"
			// }
			Description: "The parent image of the image recipe.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The tags of the image recipe.",
			//   "patternProperties": {
			//     "": {
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The tags of the image recipe.",
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"version": {
			// Property: Version
			// CloudFormation resource type schema:
			// {
			//   "description": "The version of the image recipe.",
			//   "type": "string"
			// }
			Description: "The version of the image recipe.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"working_directory": {
			// Property: WorkingDirectory
			// CloudFormation resource type schema:
			// {
			//   "description": "The working directory to be used during build and test workflows.",
			//   "type": "string"
			// }
			Description: "The working directory to be used during build and test workflows.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::ImageBuilder::ImageRecipe",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ImageBuilder::ImageRecipe").WithTerraformTypeName("awscc_imagebuilder_image_recipe")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"additional_instance_configuration": "AdditionalInstanceConfiguration",
		"arn":                               "Arn",
		"block_device_mappings":             "BlockDeviceMappings",
		"component_arn":                     "ComponentArn",
		"components":                        "Components",
		"delete_on_termination":             "DeleteOnTermination",
		"description":                       "Description",
		"device_name":                       "DeviceName",
		"ebs":                               "Ebs",
		"encrypted":                         "Encrypted",
		"iops":                              "Iops",
		"kms_key_id":                        "KmsKeyId",
		"name":                              "Name",
		"no_device":                         "NoDevice",
		"parameters":                        "Parameters",
		"parent_image":                      "ParentImage",
		"snapshot_id":                       "SnapshotId",
		"systems_manager_agent":             "SystemsManagerAgent",
		"tags":                              "Tags",
		"uninstall_after_build":             "UninstallAfterBuild",
		"user_data_override":                "UserDataOverride",
		"value":                             "Value",
		"version":                           "Version",
		"virtual_name":                      "VirtualName",
		"volume_size":                       "VolumeSize",
		"volume_type":                       "VolumeType",
		"working_directory":                 "WorkingDirectory",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
