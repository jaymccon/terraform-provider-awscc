// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package imagebuilder

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_imagebuilder_image_pipeline", imagePipelineDataSourceType)
}

// imagePipelineDataSourceType returns the Terraform awscc_imagebuilder_image_pipeline data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::ImageBuilder::ImagePipeline resource type.
func imagePipelineDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the image pipeline.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the image pipeline.",
			Type:        types.StringType,
			Computed:    true,
		},
		"container_recipe_arn": {
			// Property: ContainerRecipeArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the container recipe that defines how images are configured and tested.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the container recipe that defines how images are configured and tested.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the image pipeline.",
			//   "type": "string"
			// }
			Description: "The description of the image pipeline.",
			Type:        types.StringType,
			Computed:    true,
		},
		"distribution_configuration_arn": {
			// Property: DistributionConfigurationArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the distribution configuration associated with this image pipeline.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the distribution configuration associated with this image pipeline.",
			Type:        types.StringType,
			Computed:    true,
		},
		"enhanced_image_metadata_enabled": {
			// Property: EnhancedImageMetadataEnabled
			// CloudFormation resource type schema:
			// {
			//   "description": "Collects additional information about the image being created, including the operating system (OS) version and package list.",
			//   "type": "boolean"
			// }
			Description: "Collects additional information about the image being created, including the operating system (OS) version and package list.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"image_recipe_arn": {
			// Property: ImageRecipeArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the image recipe that defines how images are configured, tested, and assessed.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the image recipe that defines how images are configured, tested, and assessed.",
			Type:        types.StringType,
			Computed:    true,
		},
		"image_tests_configuration": {
			// Property: ImageTestsConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The image tests configuration of the image pipeline.",
			//   "properties": {
			//     "ImageTestsEnabled": {
			//       "description": "Defines if tests should be executed when building this image.",
			//       "type": "boolean"
			//     },
			//     "TimeoutMinutes": {
			//       "description": "The maximum time in minutes that tests are permitted to run.",
			//       "maximum": 1440,
			//       "minimum": 60,
			//       "type": "integer"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The image tests configuration of the image pipeline.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"image_tests_enabled": {
						// Property: ImageTestsEnabled
						Description: "Defines if tests should be executed when building this image.",
						Type:        types.BoolType,
						Computed:    true,
					},
					"timeout_minutes": {
						// Property: TimeoutMinutes
						Description: "The maximum time in minutes that tests are permitted to run.",
						Type:        types.NumberType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"infrastructure_configuration_arn": {
			// Property: InfrastructureConfigurationArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the infrastructure configuration associated with this image pipeline.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the infrastructure configuration associated with this image pipeline.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the image pipeline.",
			//   "type": "string"
			// }
			Description: "The name of the image pipeline.",
			Type:        types.StringType,
			Computed:    true,
		},
		"schedule": {
			// Property: Schedule
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The schedule of the image pipeline.",
			//   "properties": {
			//     "PipelineExecutionStartCondition": {
			//       "description": "The condition configures when the pipeline should trigger a new image build.",
			//       "enum": [
			//         "EXPRESSION_MATCH_ONLY",
			//         "EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE"
			//       ],
			//       "type": "string"
			//     },
			//     "ScheduleExpression": {
			//       "description": "The expression determines how often EC2 Image Builder evaluates your pipelineExecutionStartCondition.",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The schedule of the image pipeline.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"pipeline_execution_start_condition": {
						// Property: PipelineExecutionStartCondition
						Description: "The condition configures when the pipeline should trigger a new image build.",
						Type:        types.StringType,
						Computed:    true,
					},
					"schedule_expression": {
						// Property: ScheduleExpression
						Description: "The expression determines how often EC2 Image Builder evaluates your pipelineExecutionStartCondition.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "description": "The status of the image pipeline.",
			//   "enum": [
			//     "DISABLED",
			//     "ENABLED"
			//   ],
			//   "type": "string"
			// }
			Description: "The status of the image pipeline.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The tags of this image pipeline.",
			//   "patternProperties": {
			//     "": {
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The tags of this image pipeline.",
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::ImageBuilder::ImagePipeline",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ImageBuilder::ImagePipeline").WithTerraformTypeName("awscc_imagebuilder_image_pipeline")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                                "Arn",
		"container_recipe_arn":               "ContainerRecipeArn",
		"description":                        "Description",
		"distribution_configuration_arn":     "DistributionConfigurationArn",
		"enhanced_image_metadata_enabled":    "EnhancedImageMetadataEnabled",
		"image_recipe_arn":                   "ImageRecipeArn",
		"image_tests_configuration":          "ImageTestsConfiguration",
		"image_tests_enabled":                "ImageTestsEnabled",
		"infrastructure_configuration_arn":   "InfrastructureConfigurationArn",
		"name":                               "Name",
		"pipeline_execution_start_condition": "PipelineExecutionStartCondition",
		"schedule":                           "Schedule",
		"schedule_expression":                "ScheduleExpression",
		"status":                             "Status",
		"tags":                               "Tags",
		"timeout_minutes":                    "TimeoutMinutes",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
