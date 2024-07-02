# # AutomationWorkflowEmail
A summary of an individual Automation workflow email.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **string** | A string that uniquely identifies the Automation email.  | [optional] [readonly]
**WebId**| **int32** | The ID used in the Mailchimp web application. View this automation in your Mailchimp account at &#x60;https://{dc}.admin.mailchimp.com/campaigns/show/?id&#x3D;{web_id}&#x60;.  | [optional] [readonly]
**WorkflowId**| **string** | A string that uniquely identifies an Automation workflow.  | [optional] [readonly]
**Position**| **int32** | The position of an Automation email in a workflow.  | [optional] [readonly]
**Delay**| [**AutomationDelay**](AutomationDelay.md) |   | [optional]
**CreateTime**| [**time.Time**](time.Time.md) | The date and time the campaign was created in ISO 8601 format.  | [optional] [readonly]
**StartTime**| [**time.Time**](time.Time.md) | The date and time the campaign was started in ISO 8601 format.  | [optional] [readonly]
**ArchiveUrl**| **string** | The link to the campaign&#39;s archive version in ISO 8601 format.  | [optional] [readonly]
**Status**| **string** | The current status of the campaign. for more information please, see Model/string.php  | [optional] [readonly]
**EmailsSent**| **int32** | The total number of emails sent for this campaign.  | [optional] [readonly]
**SendTime**| [**time.Time**](time.Time.md) |  The date and time a campaign was sent in ISO 8601 format  | [optional] [readonly]
**ContentType**| **string** | How the campaign&#39;s content is put together (&#39;template&#39;, &#39;drag_and_drop&#39;, &#39;html&#39;, &#39;url&#39;).  | [optional] [readonly]
**NeedsBlockRefresh**| **bool** | Determines if the automation email needs its blocks refreshed by opening the web-based campaign editor.  | [optional] [readonly]
**HasLogoMergeTag**| **bool** | Determines if the campaign contains the *|BRAND:LOGO|* merge tag.  | [optional] [readonly]
**Recipients**| [**List2**](List2.md) |   | [optional]
**Settings**| [**CampaignSettings**](CampaignSettings.md) |   | [optional]
**Tracking**| [**CampaignTrackingOptions**](CampaignTrackingOptions.md) |   | [optional]
**SocialCard**| [**CampaignSocialCard**](CampaignSocialCard.md) |   | [optional]
**TriggerSettings**| [**AutomationTrigger**](AutomationTrigger.md) |   | [optional]
**ReportSummary**| [**CampaignReportSummary1**](CampaignReportSummary1.md) |   | [optional]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

