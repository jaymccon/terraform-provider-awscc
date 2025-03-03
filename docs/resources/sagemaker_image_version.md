---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_sagemaker_image_version Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::SageMaker::ImageVersion
---

# awscc_sagemaker_image_version (Resource)

Resource Type definition for AWS::SageMaker::ImageVersion



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **base_image** (String) The registry path of the container image on which this image version is based.
- **image_name** (String) The name of the image this version belongs to.

### Read-Only

- **container_image** (String) The registry path of the container image that contains this image version.
- **id** (String) Uniquely identifies the resource.
- **image_arn** (String) The Amazon Resource Name (ARN) of the parent image.
- **image_version_arn** (String) The Amazon Resource Name (ARN) of the image version.
- **version** (Number) The version number of the image version.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_sagemaker_image_version.example <resource ID>
```
