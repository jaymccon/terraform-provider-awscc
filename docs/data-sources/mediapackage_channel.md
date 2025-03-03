---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_mediapackage_channel Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::MediaPackage::Channel
---

# awscc_mediapackage_channel (Data Source)

Data Source schema for AWS::MediaPackage::Channel



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **arn** (String) The Amazon Resource Name (ARN) assigned to the Channel.
- **description** (String) A short text description of the Channel.
- **egress_access_logs** (Attributes) The configuration parameters for egress access logging. (see [below for nested schema](#nestedatt--egress_access_logs))
- **hls_ingest** (Attributes) A short text description of the Channel. (see [below for nested schema](#nestedatt--hls_ingest))
- **ingress_access_logs** (Attributes) The configuration parameters for egress access logging. (see [below for nested schema](#nestedatt--ingress_access_logs))
- **tags** (Attributes List) A collection of tags associated with a resource (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--egress_access_logs"></a>
### Nested Schema for `egress_access_logs`

Read-Only:

- **log_group_name** (String) Sets a custom AWS CloudWatch log group name for access logs. If a log group name isn't specified, the defaults are used: /aws/MediaPackage/EgressAccessLogs for egress access logs and /aws/MediaPackage/IngressAccessLogs for ingress access logs.


<a id="nestedatt--hls_ingest"></a>
### Nested Schema for `hls_ingest`

Read-Only:

- **ingest_endpoints** (Attributes List) A list of endpoints to which the source stream should be sent. (see [below for nested schema](#nestedatt--hls_ingest--ingest_endpoints))

<a id="nestedatt--hls_ingest--ingest_endpoints"></a>
### Nested Schema for `hls_ingest.ingest_endpoints`

Read-Only:

- **id** (String) The system generated unique identifier for the IngestEndpoint
- **password** (String) The system generated password for ingest authentication.
- **url** (String) The ingest URL to which the source stream should be sent.
- **username** (String) The system generated username for ingest authentication.



<a id="nestedatt--ingress_access_logs"></a>
### Nested Schema for `ingress_access_logs`

Read-Only:

- **log_group_name** (String) Sets a custom AWS CloudWatch log group name for access logs. If a log group name isn't specified, the defaults are used: /aws/MediaPackage/EgressAccessLogs for egress access logs and /aws/MediaPackage/IngressAccessLogs for ingress access logs.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- **key** (String)
- **value** (String)


