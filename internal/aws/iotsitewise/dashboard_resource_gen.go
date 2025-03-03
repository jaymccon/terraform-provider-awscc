// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotsitewise

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("awscc_iotsitewise_dashboard", dashboardResourceType)
}

// dashboardResourceType returns the Terraform awscc_iotsitewise_dashboard resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IoTSiteWise::Dashboard resource type.
func dashboardResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"dashboard_arn": {
			// Property: DashboardArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the dashboard.",
			//   "type": "string"
			// }
			Description: "The ARN of the dashboard.",
			Type:        types.StringType,
			Computed:    true,
		},
		"dashboard_definition": {
			// Property: DashboardDefinition
			// CloudFormation resource type schema:
			// {
			//   "description": "The dashboard definition specified in a JSON literal.",
			//   "type": "string"
			// }
			Description: "The dashboard definition specified in a JSON literal.",
			Type:        types.StringType,
			Required:    true,
		},
		"dashboard_description": {
			// Property: DashboardDescription
			// CloudFormation resource type schema:
			// {
			//   "description": "A description for the dashboard.",
			//   "type": "string"
			// }
			Description: "A description for the dashboard.",
			Type:        types.StringType,
			Required:    true,
		},
		"dashboard_id": {
			// Property: DashboardId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the dashboard.",
			//   "type": "string"
			// }
			Description: "The ID of the dashboard.",
			Type:        types.StringType,
			Computed:    true,
		},
		"dashboard_name": {
			// Property: DashboardName
			// CloudFormation resource type schema:
			// {
			//   "description": "A friendly name for the dashboard.",
			//   "type": "string"
			// }
			Description: "A friendly name for the dashboard.",
			Type:        types.StringType,
			Required:    true,
		},
		"project_id": {
			// Property: ProjectId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the project in which to create the dashboard.",
			//   "type": "string"
			// }
			Description: "The ID of the project in which to create the dashboard.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of key-value pairs that contain metadata for the dashboard.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "To add or update tag, provide both key and value. To delete tag, provide only tag key to be deleted",
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "A list of key-value pairs that contain metadata for the dashboard.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::IoTSiteWise::Dashboard",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTSiteWise::Dashboard").WithTerraformTypeName("awscc_iotsitewise_dashboard")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"dashboard_arn":         "DashboardArn",
		"dashboard_definition":  "DashboardDefinition",
		"dashboard_description": "DashboardDescription",
		"dashboard_id":          "DashboardId",
		"dashboard_name":        "DashboardName",
		"key":                   "Key",
		"project_id":            "ProjectId",
		"tags":                  "Tags",
		"value":                 "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
