# # CampaignContent1
The HTML and plain-text content for a campaign

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlainText**| **string** | The plain-text portion of the campaign. If left unspecified, we&#39;ll generate this automatically.  | [optional]
**Html**| **string** | The raw HTML for the campaign.  | [optional]
**Url**| **string** | When importing a campaign, the URL where the HTML lives.  | [optional]
**Template**| [**TemplateContent**](TemplateContent.md) |   | [optional]
**Archive**| [**UploadArchive**](UploadArchive.md) |   | [optional]
**VariateContents**| [**[]VariateContentsInner1**](VariateContentsInner1.md) | Content options for [Multivariate Campaigns](https://mailchimp.com/help/about-multivariate-campaigns/). Each content option must provide HTML content and may optionally provide plain text. For campaigns not testing content, only one object should be provided.  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

