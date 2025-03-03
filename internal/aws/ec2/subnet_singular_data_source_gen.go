// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_ec2_subnet", subnetDataSourceType)
}

// subnetDataSourceType returns the Terraform awscc_ec2_subnet data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::EC2::Subnet resource type.
func subnetDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"assign_ipv_6_address_on_creation": {
			// Property: AssignIpv6AddressOnCreation
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Computed: true,
		},
		"availability_zone": {
			// Property: AvailabilityZone
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"cidr_block": {
			// Property: CidrBlock
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"ipv_6_cidr_block": {
			// Property: Ipv6CidrBlock
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"ipv_6_cidr_blocks": {
			// Property: Ipv6CidrBlocks
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Computed: true,
		},
		"map_public_ip_on_launch": {
			// Property: MapPublicIpOnLaunch
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Computed: true,
		},
		"network_acl_association_id": {
			// Property: NetworkAclAssociationId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"outpost_arn": {
			// Property: OutpostArn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"subnet_id": {
			// Property: SubnetId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
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
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Computed: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"vpc_id": {
			// Property: VpcId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::EC2::Subnet",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::Subnet").WithTerraformTypeName("awscc_ec2_subnet")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"assign_ipv_6_address_on_creation": "AssignIpv6AddressOnCreation",
		"availability_zone":                "AvailabilityZone",
		"cidr_block":                       "CidrBlock",
		"ipv_6_cidr_block":                 "Ipv6CidrBlock",
		"ipv_6_cidr_blocks":                "Ipv6CidrBlocks",
		"key":                              "Key",
		"map_public_ip_on_launch":          "MapPublicIpOnLaunch",
		"network_acl_association_id":       "NetworkAclAssociationId",
		"outpost_arn":                      "OutpostArn",
		"subnet_id":                        "SubnetId",
		"tags":                             "Tags",
		"value":                            "Value",
		"vpc_id":                           "VpcId",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
