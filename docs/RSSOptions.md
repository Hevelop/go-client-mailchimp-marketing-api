# # RSSOptions
[RSS](https://mailchimp.com/help/share-your-blog-posts-with-mailchimp/) options for a campaign.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeedUrl**| **string** | The URL for the RSS feed.  | [optional]
**Frequency**| **string** | The frequency of the RSS Campaign. for more information please, see Model/string.php  | [optional]
**Schedule**| [**SendingSchedule**](SendingSchedule.md) |   | [optional]
**LastSent**| [**time.Time**](time.Time.md) | The date the campaign was last sent.  | [optional] [readonly]
**ConstrainRssImg**| **bool** | Whether to add CSS to images in the RSS feed to constrain their width in campaigns.  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

