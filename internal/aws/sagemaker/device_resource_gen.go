// Code generated by generators/resource/main.go; DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_sagemaker_device", deviceResourceType)
}

// deviceResourceType returns the Terraform awscc_sagemaker_device resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::SageMaker::Device resource type.
func deviceResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"device": {
			// Property: Device
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The Edge Device you want to register against a device fleet",
			//   "properties": {
			//     "Description": {
			//       "description": "Description of the device",
			//       "maxLength": 40,
			//       "minLength": 1,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "DeviceName": {
			//       "description": "The name of the device",
			//       "maxLength": 63,
			//       "minLength": 1,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "IotThingName": {
			//       "description": "AWS Internet of Things (IoT) object name.",
			//       "maxLength": 128,
			//       "pattern": "",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "DeviceName"
			//   ],
			//   "type": "object"
			// }
			Description: "The Edge Device you want to register against a device fleet",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"description": {
						// Property: Description
						Description: "Description of the device",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 40),
						},
					},
					"device_name": {
						// Property: DeviceName
						Description: "The name of the device",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 63),
						},
					},
					"iot_thing_name": {
						// Property: IotThingName
						Description: "AWS Internet of Things (IoT) object name.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(128),
						},
					},
				},
			),
			Optional: true,
		},
		"device_fleet_name": {
			// Property: DeviceFleetName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the edge device fleet",
			//   "maxLength": 63,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the edge device fleet",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 63),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "Associate tags with the resource",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The key value of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "Associate tags with the resource",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The key value of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::SageMaker::Device",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SageMaker::Device").WithTerraformTypeName("awscc_sagemaker_device")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"description":       "Description",
		"device":            "Device",
		"device_fleet_name": "DeviceFleetName",
		"device_name":       "DeviceName",
		"iot_thing_name":    "IotThingName",
		"key":               "Key",
		"tags":              "Tags",
		"value":             "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
