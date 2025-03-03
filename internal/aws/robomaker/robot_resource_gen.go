// Code generated by generators/resource/main.go; DO NOT EDIT.

package robomaker

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_robomaker_robot", robotResourceType)
}

// robotResourceType returns the Terraform awscc_robomaker_robot resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::RoboMaker::Robot resource type.
func robotResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"architecture": {
			// Property: Architecture
			// CloudFormation resource type schema:
			// {
			//   "description": "The target architecture of the robot.",
			//   "enum": [
			//     "X86_64",
			//     "ARM64",
			//     "ARMHF"
			//   ],
			//   "type": "string"
			// }
			Description: "The target architecture of the robot.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"X86_64",
					"ARM64",
					"ARMHF",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"fleet": {
			// Property: Fleet
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the fleet.",
			//   "maxLength": 1224,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the fleet.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 1224),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"greengrass_group_id": {
			// Property: GreengrassGroupId
			// CloudFormation resource type schema:
			// {
			//   "description": "The Greengrass group id.",
			//   "maxLength": 1224,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The Greengrass group id.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 1224),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name for the robot.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The name for the robot.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 255),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A key-value pair to associate with a resource.",
			//   "patternProperties": {
			//     "": {
			//       "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//       "maxLength": 256,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "A key-value pair to associate with a resource.",
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Optional: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "AWS::RoboMaker::Robot resource creates an AWS RoboMaker fleet.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::RoboMaker::Robot").WithTerraformTypeName("awscc_robomaker_robot")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"architecture":        "Architecture",
		"arn":                 "Arn",
		"fleet":               "Fleet",
		"greengrass_group_id": "GreengrassGroupId",
		"name":                "Name",
		"tags":                "Tags",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
