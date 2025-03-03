// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package datasync

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_datasync_location_s3", locationS3DataSourceType)
}

// locationS3DataSourceType returns the Terraform awscc_datasync_location_s3 data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::DataSync::LocationS3 resource type.
func locationS3DataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"location_arn": {
			// Property: LocationArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the Amazon S3 bucket location.",
			//   "maxLength": 128,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the Amazon S3 bucket location.",
			Type:        types.StringType,
			Computed:    true,
		},
		"location_uri": {
			// Property: LocationUri
			// CloudFormation resource type schema:
			// {
			//   "description": "The URL of the S3 location that was described.",
			//   "maxLength": 4356,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The URL of the S3 location that was described.",
			Type:        types.StringType,
			Computed:    true,
		},
		"s3_bucket_arn": {
			// Property: S3BucketArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the Amazon S3 bucket.",
			//   "maxLength": 156,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the Amazon S3 bucket.",
			Type:        types.StringType,
			Computed:    true,
		},
		"s3_config": {
			// Property: S3Config
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The Amazon Resource Name (ARN) of the AWS IAM role that is used to access an Amazon S3 bucket.",
			//   "properties": {
			//     "BucketAccessRoleArn": {
			//       "description": "The ARN of the IAM role of the Amazon S3 bucket.",
			//       "maxLength": 2048,
			//       "pattern": "",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "BucketAccessRoleArn"
			//   ],
			//   "type": "object"
			// }
			Description: "The Amazon Resource Name (ARN) of the AWS IAM role that is used to access an Amazon S3 bucket.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"bucket_access_role_arn": {
						// Property: BucketAccessRoleArn
						Description: "The ARN of the IAM role of the Amazon S3 bucket.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"s3_storage_class": {
			// Property: S3StorageClass
			// CloudFormation resource type schema:
			// {
			//   "default": "STANDARD",
			//   "description": "The Amazon S3 storage class you want to store your files in when this location is used as a task destination.",
			//   "enum": [
			//     "STANDARD",
			//     "STANDARD_IA",
			//     "ONEZONE_IA",
			//     "INTELLIGENT_TIERING",
			//     "GLACIER",
			//     "DEEP_ARCHIVE"
			//   ],
			//   "type": "string"
			// }
			Description: "The Amazon S3 storage class you want to store your files in when this location is used as a task destination.",
			Type:        types.StringType,
			Computed:    true,
		},
		"subdirectory": {
			// Property: Subdirectory
			// CloudFormation resource type schema:
			// {
			//   "description": "A subdirectory in the Amazon S3 bucket. This subdirectory in Amazon S3 is used to read data from the S3 source location or write data to the S3 destination.",
			//   "maxLength": 4096,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "A subdirectory in the Amazon S3 bucket. This subdirectory in Amazon S3 is used to read data from the S3 source location or write data to the S3 destination.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key for an AWS resource tag.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for an AWS resource tag.",
			//         "maxLength": 256,
			//         "minLength": 1,
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
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key for an AWS resource tag.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for an AWS resource tag.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				tfsdk.SetNestedAttributesOptions{},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::DataSync::LocationS3",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataSync::LocationS3").WithTerraformTypeName("awscc_datasync_location_s3")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"bucket_access_role_arn": "BucketAccessRoleArn",
		"key":                    "Key",
		"location_arn":           "LocationArn",
		"location_uri":           "LocationUri",
		"s3_bucket_arn":          "S3BucketArn",
		"s3_config":              "S3Config",
		"s3_storage_class":       "S3StorageClass",
		"subdirectory":           "Subdirectory",
		"tags":                   "Tags",
		"value":                  "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
