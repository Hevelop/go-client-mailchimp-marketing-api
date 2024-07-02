# # GetFacebookAdsId200Response


## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **string** | Unique ID of an Outreach.  | [optional]
**WebId**| **int32** | The ID used in the Mailchimp web application. For example, for a &#x60;regular&#x60; outreach, you can view this campaign in your Mailchimp account at &#x60;https://{dc}.admin.mailchimp.com/campaigns/show/?id&#x3D;{web_id}&#x60;.  | [optional]
**Name**| **string** | Title or name of an Outreach.  | [optional]
**Type**| **string** | The type of outreach this object is. for more information please, see Model/string.php  | [optional]
**Status**| **string** | The status of this outreach. for more information please, see Model/string.php  | [optional]
**ShowReport**| **bool** | Outreach report availability. Note: This property is hotly debated in what it _should_ convey. See [MCP-1371](https://jira.mailchimp.com/browse/MCP-1371) for more context.  | [optional]
**CreateTime**| [**time.Time**](time.Time.md) | The date and time the outreach was created in ISO 8601 format.  | [optional]
**StartTime**| [**time.Time**](time.Time.md) | The date and time the outreach was started in ISO 8601 format.  | [optional]
**UpdatedAt**| [**time.Time**](time.Time.md) | The date and time the outreach was last updated in ISO 8601 format.  | [optional]
**CanceledAt**| [**time.Time**](time.Time.md) | The date and time the outreach was canceled in ISO 8601 format.  | [optional]
**PublishedTime**| [**time.Time**](time.Time.md) | The date and time the outreach was (or will be) published in ISO 8601 format.  | [optional]
**HasSegment**| **bool** | If this outreach targets a segment of your audience.  | [optional]
**ReportSummary**| [**ReportSummary**](ReportSummary.md) |   | [optional]
**Recipients**| [**Recipients**](Recipients.md) |   | [optional]
**Thumbnail**| **string** | The URL of the thumbnail for this outreach.  | [optional]
**EmailSourceName**| **string** |   | [optional]
**PausedAt**| [**time.Time**](time.Time.md) | The date and time the ad was paused in ISO 8601 format.  | [optional]
**EndTime**| [**time.Time**](time.Time.md) | The date and time the ad was ended in ISO 8601 format.  | [optional]
**NeedsAttention**| **bool** | If the ad has a problem and needs attention.  | [optional]
**WasCanceledByFacebook**| **bool** |   | [optional]
**IsConnected**| **bool** | Check if this ad is connected to a facebook page  | [optional]
**HasAudience**| **bool** | Check if this ad has audience setup  | [optional]
**HasContent**| **bool** | Check if this ad has content  | [optional]
**Channel**| [**GetAllFacebookAds200ResponseFacebookAdsInnerAllOfChannel**](GetAllFacebookAds200ResponseFacebookAdsInnerAllOfChannel.md) |   | [optional]
**Feedback**| [**GetAllFacebookAds200ResponseFacebookAdsInnerAllOfFeedback**](GetAllFacebookAds200ResponseFacebookAdsInnerAllOfFeedback.md) |   | [optional]
**Site**| [**GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite**](GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite.md) |   | [optional]
**Audience**| [**GetAllFacebookAds200ResponseFacebookAdsInnerAllOfAudience**](GetAllFacebookAds200ResponseFacebookAdsInnerAllOfAudience.md) |   | [optional]
**Budget**| [**GetAllFacebookAds200ResponseFacebookAdsInnerAllOfBudget**](GetAllFacebookAds200ResponseFacebookAdsInnerAllOfBudget.md) |   | [optional]
**Content**| [**GetAllFacebookAds200ResponseFacebookAdsInnerAllOfContent**](GetAllFacebookAds200ResponseFacebookAdsInnerAllOfContent.md) |   | [optional]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

