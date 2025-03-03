// Code generated by generators/resource/main.go; DO NOT EDIT.

package s3outposts

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_s3outposts_access_point", accessPointResourceType)
}

// accessPointResourceType returns the Terraform awscc_s3outposts_access_point resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::S3Outposts::AccessPoint resource type.
func accessPointResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the specified AccessPoint.",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the specified AccessPoint.",
			Type:        types.StringType,
			Computed:    true,
		},
		"bucket": {
			// Property: Bucket
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the bucket you want to associate this AccessPoint with.",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the bucket you want to associate this AccessPoint with.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(20, 2048),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "A name for the AccessPoint.",
			//   "maxLength": 50,
			//   "minLength": 3,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "A name for the AccessPoint.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(3, 50),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"policy": {
			// Property: Policy
			// CloudFormation resource type schema:
			// {
			//   "description": "The access point policy associated with this access point.",
			//   "type": "object"
			// }
			Description: "The access point policy associated with this access point.",
			Type:        types.MapType{ElemType: types.StringType},
			Optional:    true,
		},
		"vpc_configuration": {
			// Property: VpcConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Virtual Private Cloud (VPC) from which requests can be made to the AccessPoint.",
			//   "properties": {
			//     "VpcId": {
			//       "description": "Virtual Private Cloud (VPC) Id from which AccessPoint will allow requests.",
			//       "maxLength": 1024,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Virtual Private Cloud (VPC) from which requests can be made to the AccessPoint.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"vpc_id": {
						// Property: VpcId
						Description: "Virtual Private Cloud (VPC) Id from which AccessPoint will allow requests.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 1024),
						},
					},
				},
			),
			Required: true,
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
		Description: "Resource Type Definition for AWS::S3Outposts::AccessPoint",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::S3Outposts::AccessPoint").WithTerraformTypeName("awscc_s3outposts_access_point")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":               "Arn",
		"bucket":            "Bucket",
		"name":              "Name",
		"policy":            "Policy",
		"vpc_configuration": "VpcConfiguration",
		"vpc_id":            "VpcId",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
