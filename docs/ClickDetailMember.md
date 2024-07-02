# # ClickDetailMember
A subscriber who clicked a specific URL in a specific campaign.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailId**| **string** | The MD5 hash of the lowercase version of the list member&#39;s email address.  | [optional] [readonly]
**EmailAddress**| **string** | Email address for a subscriber.  | [optional] [readonly]
**MergeFields**| **map[string]map[string]interface{}** | A dictionary of merge fields where the keys are the merge tags. See the [Merge Fields documentation](https://mailchimp.com/developer/marketing/docs/merge-fields/#structure) for more about the structure.  | [optional]
**Vip**| **bool** | [VIP status](https://mailchimp.com/help/designate-and-send-to-vip-contacts/) for subscriber.  | [optional] [readonly]
**Clicks**| **int32** | The total number of times the subscriber clicked on the link.  | [optional] [readonly]
**CampaignId**| **string** | The campaign id.  | [optional] [readonly]
**UrlId**| **string** | The id for the tracked URL in the campaign.  | [optional] [readonly]
**ListId**| **string** | The list id.  | [optional] [readonly]
**ListIsActive**| **bool** | The status of the list used, namely if it&#39;s deleted or disabled.  | [optional] [readonly]
**ContactStatus**| **string** | The status of the member, namely if they are subscribed, unsubscribed, deleted, non-subscribed, transactional, pending, or need reconfirmation.  | [optional] [readonly]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

