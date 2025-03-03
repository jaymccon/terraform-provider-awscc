---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_sagemaker_app_image_config Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::SageMaker::AppImageConfig
---

# awscc_sagemaker_app_image_config (Data Source)

Data Source schema for AWS::SageMaker::AppImageConfig



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **app_image_config_arn** (String) The Amazon Resource Name (ARN) of the AppImageConfig.
- **app_image_config_name** (String) The Name of the AppImageConfig.
- **kernel_gateway_image_config** (Attributes) The KernelGatewayImageConfig. (see [below for nested schema](#nestedatt--kernel_gateway_image_config))
- **tags** (Attributes List) A list of tags to apply to the AppImageConfig. (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--kernel_gateway_image_config"></a>
### Nested Schema for `kernel_gateway_image_config`

Read-Only:

- **file_system_config** (Attributes) The Amazon Elastic File System (EFS) storage configuration for a SageMaker image. (see [below for nested schema](#nestedatt--kernel_gateway_image_config--file_system_config))
- **kernel_specs** (Attributes List) The specification of the Jupyter kernels in the image. (see [below for nested schema](#nestedatt--kernel_gateway_image_config--kernel_specs))

<a id="nestedatt--kernel_gateway_image_config--file_system_config"></a>
### Nested Schema for `kernel_gateway_image_config.file_system_config`

Read-Only:

- **default_gid** (Number) The default POSIX group ID (GID). If not specified, defaults to 100.
- **default_uid** (Number) The default POSIX user ID (UID). If not specified, defaults to 1000.
- **mount_path** (String) The path within the image to mount the user's EFS home directory. The directory should be empty. If not specified, defaults to /home/sagemaker-user.


<a id="nestedatt--kernel_gateway_image_config--kernel_specs"></a>
### Nested Schema for `kernel_gateway_image_config.kernel_specs`

Read-Only:

- **display_name** (String) The display name of the kernel.
- **name** (String) The name of the kernel.



<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- **key** (String)
- **value** (String)


