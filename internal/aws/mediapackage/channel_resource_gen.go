// Code generated by generators/resource/main.go; DO NOT EDIT.

package mediapackage

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_mediapackage_channel", channelResourceType)
}

// channelResourceType returns the Terraform awscc_mediapackage_channel resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::MediaPackage::Channel resource type.
func channelResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) assigned to the Channel.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) assigned to the Channel.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "A short text description of the Channel.",
			//   "type": "string"
			// }
			Description: "A short text description of the Channel.",
			Type:        types.StringType,
			Optional:    true,
		},
		"egress_access_logs": {
			// Property: EgressAccessLogs
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The configuration parameters for egress access logging.",
			//   "properties": {
			//     "LogGroupName": {
			//       "description": "Sets a custom AWS CloudWatch log group name for access logs. If a log group name isn't specified, the defaults are used: /aws/MediaPackage/EgressAccessLogs for egress access logs and /aws/MediaPackage/IngressAccessLogs for ingress access logs.",
			//       "maxLength": 256,
			//       "minLength": 1,
			//       "pattern": "",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The configuration parameters for egress access logging.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"log_group_name": {
						// Property: LogGroupName
						Description: "Sets a custom AWS CloudWatch log group name for access logs. If a log group name isn't specified, the defaults are used: /aws/MediaPackage/EgressAccessLogs for egress access logs and /aws/MediaPackage/IngressAccessLogs for ingress access logs.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 256),
						},
					},
				},
			),
			Optional: true,
		},
		"hls_ingest": {
			// Property: HlsIngest
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A short text description of the Channel.",
			//   "properties": {
			//     "ingestEndpoints": {
			//       "description": "A list of endpoints to which the source stream should be sent.",
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "An endpoint for ingesting source content for a Channel.",
			//         "properties": {
			//           "Id": {
			//             "description": "The system generated unique identifier for the IngestEndpoint",
			//             "type": "string"
			//           },
			//           "Password": {
			//             "description": "The system generated password for ingest authentication.",
			//             "type": "string"
			//           },
			//           "Url": {
			//             "description": "The ingest URL to which the source stream should be sent.",
			//             "type": "string"
			//           },
			//           "Username": {
			//             "description": "The system generated username for ingest authentication.",
			//             "type": "string"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "type": "array"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "A short text description of the Channel.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"ingest_endpoints": {
						// Property: ingestEndpoints
						Description: "A list of endpoints to which the source stream should be sent.",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"id": {
									// Property: Id
									Description: "The system generated unique identifier for the IngestEndpoint",
									Type:        types.StringType,
									Optional:    true,
								},
								"password": {
									// Property: Password
									Description: "The system generated password for ingest authentication.",
									Type:        types.StringType,
									Optional:    true,
								},
								"url": {
									// Property: Url
									Description: "The ingest URL to which the source stream should be sent.",
									Type:        types.StringType,
									Optional:    true,
								},
								"username": {
									// Property: Username
									Description: "The system generated username for ingest authentication.",
									Type:        types.StringType,
									Optional:    true,
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Optional: true,
					},
				},
			),
			Computed: true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the Channel.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The ID of the Channel.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 256),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"ingress_access_logs": {
			// Property: IngressAccessLogs
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The configuration parameters for egress access logging.",
			//   "properties": {
			//     "LogGroupName": {
			//       "description": "Sets a custom AWS CloudWatch log group name for access logs. If a log group name isn't specified, the defaults are used: /aws/MediaPackage/EgressAccessLogs for egress access logs and /aws/MediaPackage/IngressAccessLogs for ingress access logs.",
			//       "maxLength": 256,
			//       "minLength": 1,
			//       "pattern": "",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The configuration parameters for egress access logging.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"log_group_name": {
						// Property: LogGroupName
						Description: "Sets a custom AWS CloudWatch log group name for access logs. If a log group name isn't specified, the defaults are used: /aws/MediaPackage/EgressAccessLogs for egress access logs and /aws/MediaPackage/IngressAccessLogs for ingress access logs.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 256),
						},
					},
				},
			),
			Optional: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A collection of tags associated with a resource",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A collection of tags associated with a resource",
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
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.UniqueItems(),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::MediaPackage::Channel",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::MediaPackage::Channel").WithTerraformTypeName("awscc_mediapackage_channel")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                 "Arn",
		"description":         "Description",
		"egress_access_logs":  "EgressAccessLogs",
		"hls_ingest":          "HlsIngest",
		"id":                  "Id",
		"ingest_endpoints":    "ingestEndpoints",
		"ingress_access_logs": "IngressAccessLogs",
		"key":                 "Key",
		"log_group_name":      "LogGroupName",
		"password":            "Password",
		"tags":                "Tags",
		"url":                 "Url",
		"username":            "Username",
		"value":               "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
