# # Campaign
A summary of an individual campaign&#39;s settings and content.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **string** | A string that uniquely identifies this campaign.  | [optional] [readonly]
**WebId**| **int32** | The ID used in the Mailchimp web application. View this campaign in your Mailchimp account at &#x60;https://{dc}.admin.mailchimp.com/campaigns/show/?id&#x3D;{web_id}&#x60;.  | [optional] [readonly]
**ParentCampaignId**| **string** | If this campaign is the child of another campaign, this identifies the parent campaign. For Example, for RSS or Automation children.  | [optional] [readonly]
**Type**| **string** | There are four types of [campaigns](https://mailchimp.com/help/getting-started-with-campaigns/) you can create in Mailchimp. A/B Split campaigns have been deprecated and variate campaigns should be used instead. for more information please, see Model/string.php  | [optional]
**CreateTime**| [**time.Time**](time.Time.md) | The date and time the campaign was created in ISO 8601 format.  | [optional] [readonly]
**ArchiveUrl**| **string** | The link to the campaign&#39;s archive version in ISO 8601 format.  | [optional] [readonly]
**LongArchiveUrl**| **string** | The original link to the campaign&#39;s archive version.  | [optional] [readonly]
**Status**| **string** | The current status of the campaign. for more information please, see Model/string.php  | [optional] [readonly]
**EmailsSent**| **int32** | The total number of emails sent for this campaign.  | [optional] [readonly]
**SendTime**| [**time.Time**](time.Time.md) | The date and time a campaign was sent.  | [optional] [readonly]
**ContentType**| **string** | How the campaign&#39;s content is put together. for more information please, see Model/string.php  | [optional]
**NeedsBlockRefresh**| **bool** | Determines if the campaign needs its blocks refreshed by opening the web-based campaign editor. Deprecated and will always return false.  | [optional] [readonly]
**Resendable**| **bool** | Determines if the campaign qualifies to be resent to non-openers.  | [optional] [readonly]
**Recipients**| [**List3**](List3.md) |   | [optional]
**Settings**| [**CampaignSettings2**](CampaignSettings2.md) |   | [optional]
**VariateSettings**| [**ABTestOptions**](ABTestOptions.md) |   | [optional]
**Tracking**| [**CampaignTrackingOptions1**](CampaignTrackingOptions1.md) |   | [optional]
**RssOpts**| [**RSSOptions**](RSSOptions.md) |   | [optional]
**AbSplitOpts**| [**ABTestingOptions**](ABTestingOptions.md) |   | [optional]
**SocialCard**| [**CampaignSocialCard**](CampaignSocialCard.md) |   | [optional]
**ReportSummary**| [**CampaignReportSummary2**](CampaignReportSummary2.md) |   | [optional]
**DeliveryStatus**| [**CampaignDeliveryStatus**](CampaignDeliveryStatus.md) |   | [optional]
**ResendShortcutEligibility**| [**ResendShortcutEligibility**](ResendShortcutEligibility.md) |   | [optional]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

