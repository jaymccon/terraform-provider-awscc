// Code generated by generators/resource/main.go; DO NOT EDIT.

package ce

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_ce_anomaly_monitor", anomalyMonitorResourceType)
}

// anomalyMonitorResourceType returns the Terraform awscc_ce_anomaly_monitor resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::CE::AnomalyMonitor resource type.
func anomalyMonitorResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"creation_date": {
			// Property: CreationDate
			// CloudFormation resource type schema:
			// {
			//   "description": "The date when the monitor was created. ",
			//   "maxLength": 40,
			//   "minLength": 0,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The date when the monitor was created. ",
			Type:        types.StringType,
			Computed:    true,
		},
		"dimensional_value_count": {
			// Property: DimensionalValueCount
			// CloudFormation resource type schema:
			// {
			//   "description": "The value for evaluated dimensions.",
			//   "minimum": 0,
			//   "type": "integer"
			// }
			Description: "The value for evaluated dimensions.",
			Type:        types.NumberType,
			Computed:    true,
		},
		"last_evaluated_date": {
			// Property: LastEvaluatedDate
			// CloudFormation resource type schema:
			// {
			//   "description": "The date when the monitor last evaluated for anomalies.",
			//   "maxLength": 40,
			//   "minLength": 0,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The date when the monitor last evaluated for anomalies.",
			Type:        types.StringType,
			Computed:    true,
		},
		"last_updated_date": {
			// Property: LastUpdatedDate
			// CloudFormation resource type schema:
			// {
			//   "description": "The date when the monitor was last updated.",
			//   "maxLength": 40,
			//   "minLength": 0,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The date when the monitor was last updated.",
			Type:        types.StringType,
			Computed:    true,
		},
		"monitor_arn": {
			// Property: MonitorArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Monitor ARN",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Monitor ARN",
			Type:        types.StringType,
			Computed:    true,
		},
		"monitor_dimension": {
			// Property: MonitorDimension
			// CloudFormation resource type schema:
			// {
			//   "description": "The dimensions to evaluate",
			//   "enum": [
			//     "SERVICE"
			//   ],
			//   "type": "string"
			// }
			Description: "The dimensions to evaluate",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"SERVICE",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"monitor_name": {
			// Property: MonitorName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the monitor.",
			//   "maxLength": 1024,
			//   "minLength": 0,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the monitor.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 1024),
			},
		},
		"monitor_specification": {
			// Property: MonitorSpecification
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"monitor_type": {
			// Property: MonitorType
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "DIMENSIONAL",
			//     "CUSTOM"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"DIMENSIONAL",
					"CUSTOM",
				}),
			},
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
		Description: "AWS Cost Anomaly Detection leverages advanced Machine Learning technologies to identify anomalous spend and root causes, so you can quickly take action. You can use Cost Anomaly Detection by creating monitor.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CE::AnomalyMonitor").WithTerraformTypeName("awscc_ce_anomaly_monitor")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"creation_date":           "CreationDate",
		"dimensional_value_count": "DimensionalValueCount",
		"last_evaluated_date":     "LastEvaluatedDate",
		"last_updated_date":       "LastUpdatedDate",
		"monitor_arn":             "MonitorArn",
		"monitor_dimension":       "MonitorDimension",
		"monitor_name":            "MonitorName",
		"monitor_specification":   "MonitorSpecification",
		"monitor_type":            "MonitorType",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
