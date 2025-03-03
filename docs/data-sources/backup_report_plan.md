---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_backup_report_plan Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Backup::ReportPlan
---

# awscc_backup_report_plan (Data Source)

Data Source schema for AWS::Backup::ReportPlan



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **report_delivery_channel** (Attributes) A structure that contains information about where and how to deliver your reports, specifically your Amazon S3 bucket name, S3 key prefix, and the formats of your reports. (see [below for nested schema](#nestedatt--report_delivery_channel))
- **report_plan_arn** (String) An Amazon Resource Name (ARN) that uniquely identifies a resource. The format of the ARN depends on the resource type.
- **report_plan_description** (String) An optional description of the report plan with a maximum of 1,024 characters.
- **report_plan_name** (String) The unique name of the report plan. The name must be between 1 and 256 characters, starting with a letter, and consisting of letters (a-z, A-Z), numbers (0-9), and underscores (_).
- **report_plan_tags** (Attributes List) Metadata that you can assign to help organize the report plans that you create. Each tag is a key-value pair. (see [below for nested schema](#nestedatt--report_plan_tags))
- **report_setting** (Attributes) Identifies the report template for the report. Reports are built using a report template. (see [below for nested schema](#nestedatt--report_setting))

<a id="nestedatt--report_delivery_channel"></a>
### Nested Schema for `report_delivery_channel`

Read-Only:

- **formats** (Set of String) A list of the format of your reports: CSV, JSON, or both. If not specified, the default format is CSV.
- **s3_bucket_name** (String) The unique name of the S3 bucket that receives your reports.
- **s3_key_prefix** (String) The prefix for where AWS Backup Audit Manager delivers your reports to Amazon S3. The prefix is this part of the following path: s3://your-bucket-name/prefix/Backup/us-west-2/year/month/day/report-name. If not specified, there is no prefix.


<a id="nestedatt--report_plan_tags"></a>
### Nested Schema for `report_plan_tags`

Read-Only:

- **key** (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- **value** (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.


<a id="nestedatt--report_setting"></a>
### Nested Schema for `report_setting`

Read-Only:

- **report_template** (String) Identifies the report template for the report. Reports are built using a report template. The report templates are: `BACKUP_JOB_REPORT | COPY_JOB_REPORT | RESTORE_JOB_REPORT`


