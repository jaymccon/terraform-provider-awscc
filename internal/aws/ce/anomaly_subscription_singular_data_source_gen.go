// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ce

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_ce_anomaly_subscription", anomalySubscriptionDataSourceType)
}

// anomalySubscriptionDataSourceType returns the Terraform awscc_ce_anomaly_subscription data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::CE::AnomalySubscription resource type.
func anomalySubscriptionDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"account_id": {
			// Property: AccountId
			// CloudFormation resource type schema:
			// {
			//   "description": "The accountId",
			//   "maxLength": 1024,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "The accountId",
			Type:        types.StringType,
			Computed:    true,
		},
		"frequency": {
			// Property: Frequency
			// CloudFormation resource type schema:
			// {
			//   "description": "The frequency at which anomaly reports are sent over email. ",
			//   "enum": [
			//     "DAILY",
			//     "IMMEDIATE",
			//     "WEEKLY"
			//   ],
			//   "type": "string"
			// }
			Description: "The frequency at which anomaly reports are sent over email. ",
			Type:        types.StringType,
			Computed:    true,
		},
		"monitor_arn_list": {
			// Property: MonitorArnList
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of cost anomaly monitors.",
			//   "insertionOrder": false,
			//   "items": {
			//     "description": "Subscription ARN",
			//     "pattern": "",
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "A list of cost anomaly monitors.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"subscribers": {
			// Property: Subscribers
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of subscriber",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Address": {
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Status": {
			//         "enum": [
			//           "CONFIRMED",
			//           "DECLINED"
			//         ],
			//         "type": "string"
			//       },
			//       "Type": {
			//         "enum": [
			//           "EMAIL",
			//           "SNS"
			//         ],
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Address",
			//       "Type"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "A list of subscriber",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"address": {
						// Property: Address
						Type:     types.StringType,
						Computed: true,
					},
					"status": {
						// Property: Status
						Type:     types.StringType,
						Computed: true,
					},
					"type": {
						// Property: Type
						Type:     types.StringType,
						Computed: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"subscription_arn": {
			// Property: SubscriptionArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Subscription ARN",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Subscription ARN",
			Type:        types.StringType,
			Computed:    true,
		},
		"subscription_name": {
			// Property: SubscriptionName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the subscription.",
			//   "maxLength": 1024,
			//   "minLength": 0,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the subscription.",
			Type:        types.StringType,
			Computed:    true,
		},
		"threshold": {
			// Property: Threshold
			// CloudFormation resource type schema:
			// {
			//   "description": "The dollar value that triggers a notification if the threshold is exceeded. ",
			//   "minimum": 0,
			//   "type": "number"
			// }
			Description: "The dollar value that triggers a notification if the threshold is exceeded. ",
			Type:        types.NumberType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::CE::AnomalySubscription",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CE::AnomalySubscription").WithTerraformTypeName("awscc_ce_anomaly_subscription")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_id":        "AccountId",
		"address":           "Address",
		"frequency":         "Frequency",
		"monitor_arn_list":  "MonitorArnList",
		"status":            "Status",
		"subscribers":       "Subscribers",
		"subscription_arn":  "SubscriptionArn",
		"subscription_name": "SubscriptionName",
		"threshold":         "Threshold",
		"type":              "Type",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
