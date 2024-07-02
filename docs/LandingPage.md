# # LandingPage
A summary of an individual landing page&#39;s settings and content.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **string** | A string that uniquely identifies this landing page.  | [optional] [readonly]
**Name**| **string** | The name of this landing page.  | [optional]
**Title**| **string** | The title of this landing page seen in the browser&#39;s title bar.  | [optional]
**Description**| **string** | The description of this landing page.  | [optional]
**TemplateId**| **int32** | The template_id of this landing page.  | [optional] [readonly]
**Status**| **string** | The status of this landing page. for more information please, see Model/string.php  | [optional] [readonly]
**ListId**| **string** | The list&#39;s ID associated with this landing page.  | [optional]
**StoreId**| **string** | The ID of the store associated with this landing page.  | [optional]
**WebId**| **int32** | The ID used in the Mailchimp web application.  | [optional]
**CreatedBySource**| **string** | Created by mobile or web  | [optional] [readonly]
**Url**| **string** | The url of the published landing page.  | [optional] [readonly]
**CreatedAt**| [**time.Time**](time.Time.md) | The time this landing page was created.  | [optional] [readonly]
**PublishedAt**| [**time.Time**](time.Time.md) | The time this landing page was published.  | [optional] [readonly]
**UnpublishedAt**| [**time.Time**](time.Time.md) | The time this landing page was unpublished.  | [optional] [readonly]
**UpdatedAt**| [**time.Time**](time.Time.md) | The time this landing page was updated at.  | [optional] [readonly]
**Tracking**| [**TrackingSettings**](TrackingSettings.md) |   | [optional]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

