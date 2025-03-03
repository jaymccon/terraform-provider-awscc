// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package wafv2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_wafv2_logging_configuration", loggingConfigurationDataSourceType)
}

// loggingConfigurationDataSourceType returns the Terraform awscc_wafv2_logging_configuration data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::WAFv2::LoggingConfiguration resource type.
func loggingConfigurationDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"log_destination_configs": {
			// Property: LogDestinationConfigs
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Kinesis Data Firehose Amazon Resource Name (ARNs) that you want to associate with the web ACL.",
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "The Amazon Kinesis Data Firehose Amazon Resource Name (ARNs) that you want to associate with the web ACL.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"logging_filter": {
			// Property: LoggingFilter
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Filtering that specifies which web requests are kept in the logs and which are dropped. You can filter on the rule action and on the web request labels that were applied by matching rules during web ACL evaluation.",
			//   "properties": {
			//     "DefaultBehavior": {
			//       "description": "Default handling for logs that don't match any of the specified filtering conditions.",
			//       "enum": [
			//         "KEEP",
			//         "DROP"
			//       ],
			//       "type": "string"
			//     },
			//     "Filters": {
			//       "description": "The filters that you want to apply to the logs.",
			//       "items": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "Behavior": {
			//             "description": "How to handle logs that satisfy the filter's conditions and requirement. ",
			//             "enum": [
			//               "KEEP",
			//               "DROP"
			//             ],
			//             "type": "string"
			//           },
			//           "Conditions": {
			//             "description": "Match conditions for the filter.",
			//             "items": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "ActionCondition": {
			//                   "additionalProperties": false,
			//                   "description": "A single action condition.",
			//                   "properties": {
			//                     "Action": {
			//                       "description": "Logic to apply to the filtering conditions. You can specify that, in order to satisfy the filter, a log must match all conditions or must match at least one condition.",
			//                       "enum": [
			//                         "ALLOW",
			//                         "BLOCK",
			//                         "COUNT"
			//                       ],
			//                       "type": "string"
			//                     }
			//                   },
			//                   "required": [
			//                     "Action"
			//                   ],
			//                   "type": "object"
			//                 },
			//                 "LabelNameCondition": {
			//                   "additionalProperties": false,
			//                   "description": "A single label name condition.",
			//                   "properties": {
			//                     "LabelName": {
			//                       "description": "The label name that a log record must contain in order to meet the condition. This must be a fully qualified label name. Fully qualified labels have a prefix, optional namespaces, and label name. The prefix identifies the rule group or web ACL context of the rule that added the label. ",
			//                       "type": "string"
			//                     }
			//                   },
			//                   "required": [
			//                     "LabelName"
			//                   ],
			//                   "type": "object"
			//                 }
			//               },
			//               "type": "object"
			//             },
			//             "minItems": 1,
			//             "type": "array"
			//           },
			//           "Requirement": {
			//             "description": "Logic to apply to the filtering conditions. You can specify that, in order to satisfy the filter, a log must match all conditions or must match at least one condition.",
			//             "enum": [
			//               "MEETS_ALL",
			//               "MEETS_ANY"
			//             ],
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "Behavior",
			//           "Conditions",
			//           "Requirement"
			//         ],
			//         "type": "object"
			//       },
			//       "minItems": 1,
			//       "type": "array"
			//     }
			//   },
			//   "required": [
			//     "DefaultBehavior",
			//     "Filters"
			//   ],
			//   "type": "object"
			// }
			Description: "Filtering that specifies which web requests are kept in the logs and which are dropped. You can filter on the rule action and on the web request labels that were applied by matching rules during web ACL evaluation.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"default_behavior": {
						// Property: DefaultBehavior
						Description: "Default handling for logs that don't match any of the specified filtering conditions.",
						Type:        types.StringType,
						Computed:    true,
					},
					"filters": {
						// Property: Filters
						Description: "The filters that you want to apply to the logs.",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"behavior": {
									// Property: Behavior
									Description: "How to handle logs that satisfy the filter's conditions and requirement. ",
									Type:        types.StringType,
									Computed:    true,
								},
								"conditions": {
									// Property: Conditions
									Description: "Match conditions for the filter.",
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"action_condition": {
												// Property: ActionCondition
												Description: "A single action condition.",
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"action": {
															// Property: Action
															Description: "Logic to apply to the filtering conditions. You can specify that, in order to satisfy the filter, a log must match all conditions or must match at least one condition.",
															Type:        types.StringType,
															Computed:    true,
														},
													},
												),
												Computed: true,
											},
											"label_name_condition": {
												// Property: LabelNameCondition
												Description: "A single label name condition.",
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"label_name": {
															// Property: LabelName
															Description: "The label name that a log record must contain in order to meet the condition. This must be a fully qualified label name. Fully qualified labels have a prefix, optional namespaces, and label name. The prefix identifies the rule group or web ACL context of the rule that added the label. ",
															Type:        types.StringType,
															Computed:    true,
														},
													},
												),
												Computed: true,
											},
										},
										tfsdk.ListNestedAttributesOptions{},
									),
									Computed: true,
								},
								"requirement": {
									// Property: Requirement
									Description: "Logic to apply to the filtering conditions. You can specify that, in order to satisfy the filter, a log must match all conditions or must match at least one condition.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"managed_by_firewall_manager": {
			// Property: ManagedByFirewallManager
			// CloudFormation resource type schema:
			// {
			//   "description": "Indicates whether the logging configuration was created by AWS Firewall Manager, as part of an AWS WAF policy configuration. If true, only Firewall Manager can modify or delete the configuration.",
			//   "type": "boolean"
			// }
			Description: "Indicates whether the logging configuration was created by AWS Firewall Manager, as part of an AWS WAF policy configuration. If true, only Firewall Manager can modify or delete the configuration.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"redacted_fields": {
			// Property: RedactedFields
			// CloudFormation resource type schema:
			// {
			//   "description": "The parts of the request that you want to keep out of the logs. For example, if you redact the HEADER field, the HEADER field in the firehose will be xxx.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "JsonBody": {
			//         "additionalProperties": false,
			//         "description": "Inspect the request body as JSON. The request body immediately follows the request headers. This is the part of a request that contains any additional data that you want to send to your web server as the HTTP request body, such as data from a form. ",
			//         "properties": {
			//           "InvalidFallbackBehavior": {
			//             "description": "What AWS WAF should do if it fails to completely parse the JSON body.",
			//             "enum": [
			//               "MATCH",
			//               "NO_MATCH",
			//               "EVALUATE_AS_STRING"
			//             ],
			//             "type": "string"
			//           },
			//           "MatchPattern": {
			//             "additionalProperties": false,
			//             "description": "The patterns to look for in the JSON body. AWS WAF inspects the results of these pattern matches against the rule inspection criteria. ",
			//             "properties": {
			//               "All": {
			//                 "description": "Match all of the elements. See also MatchScope in JsonBody. You must specify either this setting or the IncludedPaths setting, but not both.",
			//                 "type": "object"
			//               },
			//               "IncludedPaths": {
			//                 "description": "Match only the specified include paths. See also MatchScope in JsonBody.",
			//                 "items": {
			//                   "pattern": "",
			//                   "type": "string"
			//                 },
			//                 "minItems": 1,
			//                 "type": "array"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "MatchScope": {
			//             "description": "The parts of the JSON to match against using the MatchPattern. If you specify All, AWS WAF matches against keys and values. ",
			//             "enum": [
			//               "ALL",
			//               "KEY",
			//               "VALUE"
			//             ],
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "MatchPattern",
			//           "MatchScope"
			//         ],
			//         "type": "object"
			//       },
			//       "Method": {
			//         "description": "Inspect the HTTP method. The method indicates the type of operation that the request is asking the origin to perform. ",
			//         "type": "object"
			//       },
			//       "QueryString": {
			//         "description": "Inspect the query string. This is the part of a URL that appears after a ? character, if any. ",
			//         "type": "object"
			//       },
			//       "SingleHeader": {
			//         "additionalProperties": false,
			//         "description": "Inspect a single header. Provide the name of the header to inspect, for example, User-Agent or Referer. This setting isn't case sensitive.",
			//         "properties": {
			//           "Name": {
			//             "description": "The name of the query header to inspect.",
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "Name"
			//         ],
			//         "type": "object"
			//       },
			//       "UriPath": {
			//         "description": "Inspect the request URI path. This is the part of a web request that identifies a resource, for example, /images/daily-ad.jpg. ",
			//         "type": "object"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "The parts of the request that you want to keep out of the logs. For example, if you redact the HEADER field, the HEADER field in the firehose will be xxx.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"json_body": {
						// Property: JsonBody
						Description: "Inspect the request body as JSON. The request body immediately follows the request headers. This is the part of a request that contains any additional data that you want to send to your web server as the HTTP request body, such as data from a form. ",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"invalid_fallback_behavior": {
									// Property: InvalidFallbackBehavior
									Description: "What AWS WAF should do if it fails to completely parse the JSON body.",
									Type:        types.StringType,
									Computed:    true,
								},
								"match_pattern": {
									// Property: MatchPattern
									Description: "The patterns to look for in the JSON body. AWS WAF inspects the results of these pattern matches against the rule inspection criteria. ",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"all": {
												// Property: All
												Description: "Match all of the elements. See also MatchScope in JsonBody. You must specify either this setting or the IncludedPaths setting, but not both.",
												Type:        types.MapType{ElemType: types.StringType},
												Computed:    true,
											},
											"included_paths": {
												// Property: IncludedPaths
												Description: "Match only the specified include paths. See also MatchScope in JsonBody.",
												Type:        types.ListType{ElemType: types.StringType},
												Computed:    true,
											},
										},
									),
									Computed: true,
								},
								"match_scope": {
									// Property: MatchScope
									Description: "The parts of the JSON to match against using the MatchPattern. If you specify All, AWS WAF matches against keys and values. ",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"method": {
						// Property: Method
						Description: "Inspect the HTTP method. The method indicates the type of operation that the request is asking the origin to perform. ",
						Type:        types.MapType{ElemType: types.StringType},
						Computed:    true,
					},
					"query_string": {
						// Property: QueryString
						Description: "Inspect the query string. This is the part of a URL that appears after a ? character, if any. ",
						Type:        types.MapType{ElemType: types.StringType},
						Computed:    true,
					},
					"single_header": {
						// Property: SingleHeader
						Description: "Inspect a single header. Provide the name of the header to inspect, for example, User-Agent or Referer. This setting isn't case sensitive.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"name": {
									// Property: Name
									Description: "The name of the query header to inspect.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"uri_path": {
						// Property: UriPath
						Description: "Inspect the request URI path. This is the part of a web request that identifies a resource, for example, /images/daily-ad.jpg. ",
						Type:        types.MapType{ElemType: types.StringType},
						Computed:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"resource_arn": {
			// Property: ResourceArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the web ACL that you want to associate with LogDestinationConfigs.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the web ACL that you want to associate with LogDestinationConfigs.",
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
		Description: "Data Source schema for AWS::WAFv2::LoggingConfiguration",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::WAFv2::LoggingConfiguration").WithTerraformTypeName("awscc_wafv2_logging_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"action":                      "Action",
		"action_condition":            "ActionCondition",
		"all":                         "All",
		"behavior":                    "Behavior",
		"conditions":                  "Conditions",
		"default_behavior":            "DefaultBehavior",
		"filters":                     "Filters",
		"included_paths":              "IncludedPaths",
		"invalid_fallback_behavior":   "InvalidFallbackBehavior",
		"json_body":                   "JsonBody",
		"label_name":                  "LabelName",
		"label_name_condition":        "LabelNameCondition",
		"log_destination_configs":     "LogDestinationConfigs",
		"logging_filter":              "LoggingFilter",
		"managed_by_firewall_manager": "ManagedByFirewallManager",
		"match_pattern":               "MatchPattern",
		"match_scope":                 "MatchScope",
		"method":                      "Method",
		"name":                        "Name",
		"query_string":                "QueryString",
		"redacted_fields":             "RedactedFields",
		"requirement":                 "Requirement",
		"resource_arn":                "ResourceArn",
		"single_header":               "SingleHeader",
		"uri_path":                    "UriPath",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
