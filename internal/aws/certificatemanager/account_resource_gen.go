// Code generated by generators/resource/main.go; DO NOT EDIT.

package certificatemanager

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_certificatemanager_account", accountResourceType)
}

// accountResourceType returns the Terraform awscc_certificatemanager_account resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::CertificateManager::Account resource type.
func accountResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"account_id": {
			// Property: AccountId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"expiry_events_configuration": {
			// Property: ExpiryEventsConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "DaysBeforeExpiry": {
			//       "maximum": 45,
			//       "minimum": 1,
			//       "type": "integer"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"days_before_expiry": {
						// Property: DaysBeforeExpiry
						Type:     types.NumberType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntBetween(1, 45),
						},
					},
				},
			),
			Required: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::CertificateManager::Account.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CertificateManager::Account").WithTerraformTypeName("awscc_certificatemanager_account")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_id":                  "AccountId",
		"days_before_expiry":          "DaysBeforeExpiry",
		"expiry_events_configuration": "ExpiryEventsConfiguration",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
