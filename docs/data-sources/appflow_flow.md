---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_appflow_flow Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::AppFlow::Flow
---

# awscc_appflow_flow (Data Source)

Data Source schema for AWS::AppFlow::Flow



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **description** (String) Description of the flow.
- **destination_flow_config_list** (Attributes List) List of Destination connectors of the flow. (see [below for nested schema](#nestedatt--destination_flow_config_list))
- **flow_arn** (String) ARN identifier of the flow.
- **flow_name** (String) Name of the flow.
- **kms_arn** (String) The ARN of the AWS Key Management Service (AWS KMS) key that's used to encrypt your function's environment variables. If it's not provided, AWS Lambda uses a default service key.
- **source_flow_config** (Attributes) Configurations of Source connector of the flow. (see [below for nested schema](#nestedatt--source_flow_config))
- **tags** (Attributes List) List of Tags. (see [below for nested schema](#nestedatt--tags))
- **tasks** (Attributes List) List of tasks for the flow. (see [below for nested schema](#nestedatt--tasks))
- **trigger_config** (Attributes) Trigger settings of the flow. (see [below for nested schema](#nestedatt--trigger_config))

<a id="nestedatt--destination_flow_config_list"></a>
### Nested Schema for `destination_flow_config_list`

Read-Only:

- **connector_profile_name** (String) Name of destination connector profile
- **connector_type** (String) Destination connector type
- **destination_connector_properties** (Attributes) Destination connector details (see [below for nested schema](#nestedatt--destination_flow_config_list--destination_connector_properties))

<a id="nestedatt--destination_flow_config_list--destination_connector_properties"></a>
### Nested Schema for `destination_flow_config_list.destination_connector_properties`

Read-Only:

- **event_bridge** (Attributes) (see [below for nested schema](#nestedatt--destination_flow_config_list--destination_connector_properties--event_bridge))
- **lookout_metrics** (Attributes) (see [below for nested schema](#nestedatt--destination_flow_config_list--destination_connector_properties--lookout_metrics))
- **redshift** (Attributes) (see [below for nested schema](#nestedatt--destination_flow_config_list--destination_connector_properties--redshift))
- **s3** (Attributes) (see [below for nested schema](#nestedatt--destination_flow_config_list--destination_connector_properties--s3))
- **salesforce** (Attributes) (see [below for nested schema](#nestedatt--destination_flow_config_list--destination_connector_properties--salesforce))
- **snowflake** (Attributes) (see [below for nested schema](#nestedatt--destination_flow_config_list--destination_connector_properties--snowflake))
- **upsolver** (Attributes) (see [below for nested schema](#nestedatt--destination_flow_config_list--destination_connector_properties--upsolver))
- **zendesk** (Attributes) (see [below for nested schema](#nestedatt--destination_flow_config_list--destination_connector_properties--zendesk))

<a id="nestedatt--destination_flow_config_list--destination_connector_properties--event_bridge"></a>
### Nested Schema for `destination_flow_config_list.destination_connector_properties.event_bridge`

Read-Only:

- **error_handling_config** (Attributes) (see [below for nested schema](#nestedatt--destination_flow_config_list--destination_connector_properties--event_bridge--error_handling_config))
- **object** (String)

<a id="nestedatt--destination_flow_config_list--destination_connector_properties--event_bridge--error_handling_config"></a>
### Nested Schema for `destination_flow_config_list.destination_connector_properties.event_bridge.object`

Read-Only:

- **bucket_name** (String)
- **bucket_prefix** (String)
- **fail_on_first_error** (Boolean)



<a id="nestedatt--destination_flow_config_list--destination_connector_properties--lookout_metrics"></a>
### Nested Schema for `destination_flow_config_list.destination_connector_properties.lookout_metrics`

Read-Only:

- **object** (String)


<a id="nestedatt--destination_flow_config_list--destination_connector_properties--redshift"></a>
### Nested Schema for `destination_flow_config_list.destination_connector_properties.redshift`

Read-Only:

- **bucket_prefix** (String)
- **error_handling_config** (Attributes) (see [below for nested schema](#nestedatt--destination_flow_config_list--destination_connector_properties--redshift--error_handling_config))
- **intermediate_bucket_name** (String)
- **object** (String)

<a id="nestedatt--destination_flow_config_list--destination_connector_properties--redshift--error_handling_config"></a>
### Nested Schema for `destination_flow_config_list.destination_connector_properties.redshift.object`

Read-Only:

- **bucket_name** (String)
- **bucket_prefix** (String)
- **fail_on_first_error** (Boolean)



<a id="nestedatt--destination_flow_config_list--destination_connector_properties--s3"></a>
### Nested Schema for `destination_flow_config_list.destination_connector_properties.s3`

Read-Only:

- **bucket_name** (String)
- **bucket_prefix** (String)
- **s3_output_format_config** (Attributes) (see [below for nested schema](#nestedatt--destination_flow_config_list--destination_connector_properties--s3--s3_output_format_config))

<a id="nestedatt--destination_flow_config_list--destination_connector_properties--s3--s3_output_format_config"></a>
### Nested Schema for `destination_flow_config_list.destination_connector_properties.s3.s3_output_format_config`

Read-Only:

- **aggregation_config** (Attributes) (see [below for nested schema](#nestedatt--destination_flow_config_list--destination_connector_properties--s3--s3_output_format_config--aggregation_config))
- **file_type** (String)
- **prefix_config** (Attributes) (see [below for nested schema](#nestedatt--destination_flow_config_list--destination_connector_properties--s3--s3_output_format_config--prefix_config))

<a id="nestedatt--destination_flow_config_list--destination_connector_properties--s3--s3_output_format_config--aggregation_config"></a>
### Nested Schema for `destination_flow_config_list.destination_connector_properties.s3.s3_output_format_config.aggregation_config`

Read-Only:

- **aggregation_type** (String)


<a id="nestedatt--destination_flow_config_list--destination_connector_properties--s3--s3_output_format_config--prefix_config"></a>
### Nested Schema for `destination_flow_config_list.destination_connector_properties.s3.s3_output_format_config.prefix_config`

Read-Only:

- **prefix_format** (String)
- **prefix_type** (String)




<a id="nestedatt--destination_flow_config_list--destination_connector_properties--salesforce"></a>
### Nested Schema for `destination_flow_config_list.destination_connector_properties.salesforce`

Read-Only:

- **error_handling_config** (Attributes) (see [below for nested schema](#nestedatt--destination_flow_config_list--destination_connector_properties--salesforce--error_handling_config))
- **id_field_names** (List of String) List of fields used as ID when performing a write operation.
- **object** (String)
- **write_operation_type** (String)

<a id="nestedatt--destination_flow_config_list--destination_connector_properties--salesforce--error_handling_config"></a>
### Nested Schema for `destination_flow_config_list.destination_connector_properties.salesforce.write_operation_type`

Read-Only:

- **bucket_name** (String)
- **bucket_prefix** (String)
- **fail_on_first_error** (Boolean)



<a id="nestedatt--destination_flow_config_list--destination_connector_properties--snowflake"></a>
### Nested Schema for `destination_flow_config_list.destination_connector_properties.snowflake`

Read-Only:

- **bucket_prefix** (String)
- **error_handling_config** (Attributes) (see [below for nested schema](#nestedatt--destination_flow_config_list--destination_connector_properties--snowflake--error_handling_config))
- **intermediate_bucket_name** (String)
- **object** (String)

<a id="nestedatt--destination_flow_config_list--destination_connector_properties--snowflake--error_handling_config"></a>
### Nested Schema for `destination_flow_config_list.destination_connector_properties.snowflake.object`

Read-Only:

- **bucket_name** (String)
- **bucket_prefix** (String)
- **fail_on_first_error** (Boolean)



<a id="nestedatt--destination_flow_config_list--destination_connector_properties--upsolver"></a>
### Nested Schema for `destination_flow_config_list.destination_connector_properties.upsolver`

Read-Only:

- **bucket_name** (String)
- **bucket_prefix** (String)
- **s3_output_format_config** (Attributes) (see [below for nested schema](#nestedatt--destination_flow_config_list--destination_connector_properties--upsolver--s3_output_format_config))

<a id="nestedatt--destination_flow_config_list--destination_connector_properties--upsolver--s3_output_format_config"></a>
### Nested Schema for `destination_flow_config_list.destination_connector_properties.upsolver.s3_output_format_config`

Read-Only:

- **aggregation_config** (Attributes) (see [below for nested schema](#nestedatt--destination_flow_config_list--destination_connector_properties--upsolver--s3_output_format_config--aggregation_config))
- **file_type** (String)
- **prefix_config** (Attributes) (see [below for nested schema](#nestedatt--destination_flow_config_list--destination_connector_properties--upsolver--s3_output_format_config--prefix_config))

<a id="nestedatt--destination_flow_config_list--destination_connector_properties--upsolver--s3_output_format_config--aggregation_config"></a>
### Nested Schema for `destination_flow_config_list.destination_connector_properties.upsolver.s3_output_format_config.aggregation_config`

Read-Only:

- **aggregation_type** (String)


<a id="nestedatt--destination_flow_config_list--destination_connector_properties--upsolver--s3_output_format_config--prefix_config"></a>
### Nested Schema for `destination_flow_config_list.destination_connector_properties.upsolver.s3_output_format_config.prefix_config`

Read-Only:

- **prefix_format** (String)
- **prefix_type** (String)




<a id="nestedatt--destination_flow_config_list--destination_connector_properties--zendesk"></a>
### Nested Schema for `destination_flow_config_list.destination_connector_properties.zendesk`

Read-Only:

- **error_handling_config** (Attributes) (see [below for nested schema](#nestedatt--destination_flow_config_list--destination_connector_properties--zendesk--error_handling_config))
- **id_field_names** (List of String) List of fields used as ID when performing a write operation.
- **object** (String)
- **write_operation_type** (String)

<a id="nestedatt--destination_flow_config_list--destination_connector_properties--zendesk--error_handling_config"></a>
### Nested Schema for `destination_flow_config_list.destination_connector_properties.zendesk.write_operation_type`

Read-Only:

- **bucket_name** (String)
- **bucket_prefix** (String)
- **fail_on_first_error** (Boolean)





<a id="nestedatt--source_flow_config"></a>
### Nested Schema for `source_flow_config`

Read-Only:

- **connector_profile_name** (String) Name of source connector profile
- **connector_type** (String) Type of source connector
- **incremental_pull_config** (Attributes) Configuration for scheduled incremental data pull (see [below for nested schema](#nestedatt--source_flow_config--incremental_pull_config))
- **source_connector_properties** (Attributes) Source connector details required to query a connector (see [below for nested schema](#nestedatt--source_flow_config--source_connector_properties))

<a id="nestedatt--source_flow_config--incremental_pull_config"></a>
### Nested Schema for `source_flow_config.incremental_pull_config`

Read-Only:

- **datetime_type_field_name** (String) Name of the datetime/timestamp data type field to be used for importing incremental records from the source


<a id="nestedatt--source_flow_config--source_connector_properties"></a>
### Nested Schema for `source_flow_config.source_connector_properties`

Read-Only:

- **amplitude** (Attributes) (see [below for nested schema](#nestedatt--source_flow_config--source_connector_properties--amplitude))
- **datadog** (Attributes) (see [below for nested schema](#nestedatt--source_flow_config--source_connector_properties--datadog))
- **dynatrace** (Attributes) (see [below for nested schema](#nestedatt--source_flow_config--source_connector_properties--dynatrace))
- **google_analytics** (Attributes) (see [below for nested schema](#nestedatt--source_flow_config--source_connector_properties--google_analytics))
- **infor_nexus** (Attributes) (see [below for nested schema](#nestedatt--source_flow_config--source_connector_properties--infor_nexus))
- **marketo** (Attributes) (see [below for nested schema](#nestedatt--source_flow_config--source_connector_properties--marketo))
- **s3** (Attributes) (see [below for nested schema](#nestedatt--source_flow_config--source_connector_properties--s3))
- **salesforce** (Attributes) (see [below for nested schema](#nestedatt--source_flow_config--source_connector_properties--salesforce))
- **service_now** (Attributes) (see [below for nested schema](#nestedatt--source_flow_config--source_connector_properties--service_now))
- **singular** (Attributes) (see [below for nested schema](#nestedatt--source_flow_config--source_connector_properties--singular))
- **slack** (Attributes) (see [below for nested schema](#nestedatt--source_flow_config--source_connector_properties--slack))
- **trendmicro** (Attributes) (see [below for nested schema](#nestedatt--source_flow_config--source_connector_properties--trendmicro))
- **veeva** (Attributes) (see [below for nested schema](#nestedatt--source_flow_config--source_connector_properties--veeva))
- **zendesk** (Attributes) (see [below for nested schema](#nestedatt--source_flow_config--source_connector_properties--zendesk))

<a id="nestedatt--source_flow_config--source_connector_properties--amplitude"></a>
### Nested Schema for `source_flow_config.source_connector_properties.amplitude`

Read-Only:

- **object** (String)


<a id="nestedatt--source_flow_config--source_connector_properties--datadog"></a>
### Nested Schema for `source_flow_config.source_connector_properties.datadog`

Read-Only:

- **object** (String)


<a id="nestedatt--source_flow_config--source_connector_properties--dynatrace"></a>
### Nested Schema for `source_flow_config.source_connector_properties.dynatrace`

Read-Only:

- **object** (String)


<a id="nestedatt--source_flow_config--source_connector_properties--google_analytics"></a>
### Nested Schema for `source_flow_config.source_connector_properties.google_analytics`

Read-Only:

- **object** (String)


<a id="nestedatt--source_flow_config--source_connector_properties--infor_nexus"></a>
### Nested Schema for `source_flow_config.source_connector_properties.infor_nexus`

Read-Only:

- **object** (String)


<a id="nestedatt--source_flow_config--source_connector_properties--marketo"></a>
### Nested Schema for `source_flow_config.source_connector_properties.marketo`

Read-Only:

- **object** (String)


<a id="nestedatt--source_flow_config--source_connector_properties--s3"></a>
### Nested Schema for `source_flow_config.source_connector_properties.s3`

Read-Only:

- **bucket_name** (String)
- **bucket_prefix** (String)


<a id="nestedatt--source_flow_config--source_connector_properties--salesforce"></a>
### Nested Schema for `source_flow_config.source_connector_properties.salesforce`

Read-Only:

- **enable_dynamic_field_update** (Boolean)
- **include_deleted_records** (Boolean)
- **object** (String)


<a id="nestedatt--source_flow_config--source_connector_properties--service_now"></a>
### Nested Schema for `source_flow_config.source_connector_properties.service_now`

Read-Only:

- **object** (String)


<a id="nestedatt--source_flow_config--source_connector_properties--singular"></a>
### Nested Schema for `source_flow_config.source_connector_properties.singular`

Read-Only:

- **object** (String)


<a id="nestedatt--source_flow_config--source_connector_properties--slack"></a>
### Nested Schema for `source_flow_config.source_connector_properties.slack`

Read-Only:

- **object** (String)


<a id="nestedatt--source_flow_config--source_connector_properties--trendmicro"></a>
### Nested Schema for `source_flow_config.source_connector_properties.trendmicro`

Read-Only:

- **object** (String)


<a id="nestedatt--source_flow_config--source_connector_properties--veeva"></a>
### Nested Schema for `source_flow_config.source_connector_properties.veeva`

Read-Only:

- **document_type** (String)
- **include_all_versions** (Boolean)
- **include_renditions** (Boolean)
- **include_source_files** (Boolean)
- **object** (String)


<a id="nestedatt--source_flow_config--source_connector_properties--zendesk"></a>
### Nested Schema for `source_flow_config.source_connector_properties.zendesk`

Read-Only:

- **object** (String)




<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- **key** (String) A string used to identify this tag
- **value** (String) A string containing the value for the tag


<a id="nestedatt--tasks"></a>
### Nested Schema for `tasks`

Read-Only:

- **connector_operator** (Attributes) Operation to be performed on provided source fields (see [below for nested schema](#nestedatt--tasks--connector_operator))
- **destination_field** (String) A field value on which source field should be validated
- **source_fields** (List of String) Source fields on which particular task will be applied
- **task_properties** (Attributes List) A Map used to store task related info (see [below for nested schema](#nestedatt--tasks--task_properties))
- **task_type** (String) Type of task

<a id="nestedatt--tasks--connector_operator"></a>
### Nested Schema for `tasks.connector_operator`

Read-Only:

- **amplitude** (String)
- **datadog** (String)
- **dynatrace** (String)
- **google_analytics** (String)
- **infor_nexus** (String)
- **marketo** (String)
- **s3** (String)
- **salesforce** (String)
- **service_now** (String)
- **singular** (String)
- **slack** (String)
- **trendmicro** (String)
- **veeva** (String)
- **zendesk** (String)


<a id="nestedatt--tasks--task_properties"></a>
### Nested Schema for `tasks.task_properties`

Read-Only:

- **key** (String)
- **value** (String)



<a id="nestedatt--trigger_config"></a>
### Nested Schema for `trigger_config`

Read-Only:

- **trigger_properties** (Attributes) Details required based on the type of trigger (see [below for nested schema](#nestedatt--trigger_config--trigger_properties))
- **trigger_type** (String) Trigger type of the flow

<a id="nestedatt--trigger_config--trigger_properties"></a>
### Nested Schema for `trigger_config.trigger_properties`

Read-Only:

- **data_pull_mode** (String)
- **schedule_end_time** (Number)
- **schedule_expression** (String)
- **schedule_offset** (Number)
- **schedule_start_time** (Number)
- **time_zone** (String)


