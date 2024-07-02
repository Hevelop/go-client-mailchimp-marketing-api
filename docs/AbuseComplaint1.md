# # AbuseComplaint1
Details of abuse complaints for a specific list. An abuse complaint occurs when your recipient clicks to &#39;report spam&#39; in their email program.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **int32** | The id for the abuse report  | [optional] [readonly]
**CampaignId**| **string** | The campaign id for the abuse report  | [optional] [readonly]
**ListId**| **string** | The unique id of the list for the abuse report.  | [optional] [readonly]
**ListIsActive**| **bool** | The status of the list used, namely if it&#39;s deleted or disabled.  | [optional] [readonly]
**EmailId**| **string** | The MD5 hash of the lowercase version of the list member&#39;s email address.  | [optional] [readonly]
**EmailAddress**| **string** | Email address for a subscriber.  | [optional] [readonly]
**MergeFields**| **map[string]map[string]interface{}** | A dictionary of merge fields where the keys are the merge tags. See the [Merge Fields documentation](https://mailchimp.com/developer/marketing/docs/merge-fields/#structure) for more about the structure.  | [optional]
**Vip**| **bool** | [VIP status](https://mailchimp.com/help/designate-and-send-to-vip-contacts/) for subscriber.  | [optional] [readonly]
**Date**| [**time.Time**](time.Time.md) | Date for the abuse report  | [optional] [readonly]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

