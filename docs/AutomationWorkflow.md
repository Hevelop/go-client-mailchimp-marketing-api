# # AutomationWorkflow
A summary of an individual Automation workflow&#39;s settings and content.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **string** | A string that identifies the Automation.  | [optional] [readonly]
**CreateTime**| [**time.Time**](time.Time.md) | The date and time the Automation was created in ISO 8601 format.  | [optional] [readonly]
**StartTime**| [**time.Time**](time.Time.md) | The date and time the Automation was started in ISO 8601 format.  | [optional] [readonly]
**Status**| **string** | The current status of the Automation. for more information please, see Model/string.php  | [optional] [readonly]
**EmailsSent**| **int32** | The total number of emails sent for the Automation.  | [optional] [readonly]
**Recipients**| [**List**](List.md) |   | [optional]
**Settings**| [**AutomationCampaignSettings**](AutomationCampaignSettings.md) |   | [optional]
**Tracking**| [**AutomationTrackingOptions**](AutomationTrackingOptions.md) |   | [optional]
**TriggerSettings**| [**AutomationTrigger**](AutomationTrigger.md) |   | [optional]
**ReportSummary**| [**CampaignReportSummary**](CampaignReportSummary.md) |   | [optional]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

