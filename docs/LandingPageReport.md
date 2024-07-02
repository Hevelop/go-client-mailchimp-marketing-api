# # LandingPageReport
A summary of an individual landing page&#39;s settings and content.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **string** | A string that uniquely identifies this landing page.  | [optional] [readonly]
**Name**| **string** | The name of this landing page the user will see.  | [optional] [readonly]
**Title**| **string** | The name of the landing page the user&#39;s customers will see.  | [optional] [readonly]
**Url**| **string** | The landing page url.  | [optional] [readonly]
**PublishedAt**| [**time.Time**](time.Time.md) | The time this landing page was published.  | [optional] [readonly]
**UnpublishedAt**| [**time.Time**](time.Time.md) | The time this landing page was unpublished.  | [optional] [readonly]
**Status**| **string** | The status of the landing page.  | [optional] [readonly]
**ListId**| **string** | The list id connected to this landing page.  | [optional] [readonly]
**Visits**| **int32** | The number of visits to this landing pages.  | [optional] [readonly]
**UniqueVisits**| **int32** | The number of unique visits to this landing pages.  | [optional] [readonly]
**Subscribes**| **int32** | The number of subscribes to this landing pages.  | [optional] [readonly]
**Clicks**| **int32** | The number of clicks to this landing pages.  | [optional] [readonly]
**ConversionRate**| **float32** | The percentage of people who visited your landing page and were added to your list.  | [optional] [readonly]
**Timeseries**| [**LandingPageReportTimeseries**](LandingPageReportTimeseries.md) |   | [optional]
**Ecommerce**| [**LandingPageReportEcommerce**](LandingPageReportEcommerce.md) |   | [optional]
**WebId**| **int32** | The ID used in the Mailchimp web application.  | [optional] [readonly]
**ListName**| **string** | List Name  | [optional] [readonly]
**SignupTags**| [**[]Tag**](Tag.md) | A list of tags associated to the landing page.  | [optional] [readonly]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

