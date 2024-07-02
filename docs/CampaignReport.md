# # CampaignReport
Report details about a sent campaign.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **string** | A string that uniquely identifies this campaign.  | [optional]
**CampaignTitle**| **string** | The title of the campaign.  | [optional] [readonly]
**Type**| **string** | The type of campaign (regular, plain-text, ab_split, rss, automation, variate, or auto).  | [optional]
**ListId**| **string** | The unique list id.  | [optional] [readonly]
**ListIsActive**| **bool** | The status of the list used, namely if it&#39;s deleted or disabled.  | [optional] [readonly]
**ListName**| **string** | The name of the list.  | [optional] [readonly]
**SubjectLine**| **string** | The subject line for the campaign.  | [optional] [readonly]
**PreviewText**| **string** | The preview text for the campaign.  | [optional]
**EmailsSent**| **int32** | The total number of emails sent for this campaign.  | [optional]
**AbuseReports**| **int32** | The number of abuse reports generated for this campaign.  | [optional]
**Unsubscribed**| **int32** | The total number of unsubscribed members for this campaign.  | [optional] [readonly]
**SendTime**| [**time.Time**](time.Time.md) | The date and time a campaign was sent in ISO 8601 format.  | [optional] [readonly]
**RssLastSend**| [**time.Time**](time.Time.md) | For RSS campaigns, the date and time of the last send in ISO 8601 format.  | [optional] [readonly]
**Bounces**| [**Bounces**](Bounces.md) |   | [optional]
**Forwards**| [**Forwards**](Forwards.md) |   | [optional]
**Opens**| [**Opens**](Opens.md) |   | [optional]
**Clicks**| [**Clicks**](Clicks.md) |   | [optional]
**FacebookLikes**| [**FacebookLikes**](FacebookLikes.md) |   | [optional]
**IndustryStats**| [**IndustryStats1**](IndustryStats1.md) |   | [optional]
**ListStats**| [**ListStats**](ListStats.md) |   | [optional]
**AbSplit**| [**ABSplitStats**](ABSplitStats.md) |   | [optional]
**Timewarp**| [**[]TimewarpStatsInner**](TimewarpStatsInner.md) | An hourly breakdown of sends, opens, and clicks if a campaign is sent using timewarp.  | [optional]
**Timeseries**| [**[]TimeseriesInner**](TimeseriesInner.md) | An hourly breakdown of the performance of the campaign over the first 24 hours.  | [optional]
**ShareReport**| [**ShareReport**](ShareReport.md) |   | [optional]
**Ecommerce**| [**ECommerceReport1**](ECommerceReport1.md) |   | [optional]
**DeliveryStatus**| [**CampaignDeliveryStatus**](CampaignDeliveryStatus.md) |   | [optional]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

