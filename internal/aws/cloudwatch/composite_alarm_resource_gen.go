// Code generated by generators/resource/main.go; DO NOT EDIT.

package cloudwatch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_cloudwatch_composite_alarm", compositeAlarmResourceType)
}

// compositeAlarmResourceType returns the Terraform awscc_cloudwatch_composite_alarm resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::CloudWatch::CompositeAlarm resource type.
func compositeAlarmResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"actions_enabled": {
			// Property: ActionsEnabled
			// CloudFormation resource type schema:
			// {
			//   "description": "Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.",
			//   "type": "boolean"
			// }
			Description: "Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"alarm_actions": {
			// Property: AlarmActions
			// CloudFormation resource type schema:
			// {
			//   "description": "The list of actions to execute when this alarm transitions into an ALARM state from any other state. Specify each action as an Amazon Resource Name (ARN).",
			//   "items": {
			//     "description": "Amazon Resource Name (ARN) of the action",
			//     "maxLength": 1024,
			//     "minLength": 1,
			//     "type": "string"
			//   },
			//   "maxItems": 5,
			//   "type": "array"
			// }
			Description: "The list of actions to execute when this alarm transitions into an ALARM state from any other state. Specify each action as an Amazon Resource Name (ARN).",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(5),
				validate.ArrayForEach(validate.StringLenBetween(1, 1024)),
			},
		},
		"alarm_description": {
			// Property: AlarmDescription
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the alarm",
			//   "maxLength": 1024,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "The description of the alarm",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 1024),
			},
		},
		"alarm_name": {
			// Property: AlarmName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the Composite Alarm",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The name of the Composite Alarm",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 255),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"alarm_rule": {
			// Property: AlarmRule
			// CloudFormation resource type schema:
			// {
			//   "description": "Expression which aggregates the state of other Alarms (Metric or Composite Alarms)",
			//   "maxLength": 10240,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Expression which aggregates the state of other Alarms (Metric or Composite Alarms)",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 10240),
			},
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Amazon Resource Name (ARN) of the alarm",
			//   "maxLength": 1600,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Amazon Resource Name (ARN) of the alarm",
			Type:        types.StringType,
			Computed:    true,
		},
		"insufficient_data_actions": {
			// Property: InsufficientDataActions
			// CloudFormation resource type schema:
			// {
			//   "description": "The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).",
			//   "items": {
			//     "description": "Amazon Resource Name (ARN) of the action",
			//     "maxLength": 1024,
			//     "minLength": 1,
			//     "type": "string"
			//   },
			//   "maxItems": 5,
			//   "type": "array"
			// }
			Description: "The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(5),
				validate.ArrayForEach(validate.StringLenBetween(1, 1024)),
			},
		},
		"ok_actions": {
			// Property: OKActions
			// CloudFormation resource type schema:
			// {
			//   "description": "The actions to execute when this alarm transitions to the OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).",
			//   "items": {
			//     "description": "Amazon Resource Name (ARN) of the action",
			//     "maxLength": 1024,
			//     "minLength": 1,
			//     "type": "string"
			//   },
			//   "maxItems": 5,
			//   "type": "array"
			// }
			Description: "The actions to execute when this alarm transitions to the OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(5),
				validate.ArrayForEach(validate.StringLenBetween(1, 1024)),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "The AWS::CloudWatch::CompositeAlarm type specifies an alarm which aggregates the states of other Alarms (Metric or Composite Alarms) as defined by the AlarmRule expression",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudWatch::CompositeAlarm").WithTerraformTypeName("awscc_cloudwatch_composite_alarm")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"actions_enabled":           "ActionsEnabled",
		"alarm_actions":             "AlarmActions",
		"alarm_description":         "AlarmDescription",
		"alarm_name":                "AlarmName",
		"alarm_rule":                "AlarmRule",
		"arn":                       "Arn",
		"insufficient_data_actions": "InsufficientDataActions",
		"ok_actions":                "OKActions",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
