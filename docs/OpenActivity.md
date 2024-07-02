# # OpenActivity
A list of a member&#39;s opens activity in a specific campaign.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId**| **string** | The unique id for the campaign.  | [optional] [readonly]
**ListId**| **string** | The unique id for the list.  | [optional] [readonly]
**ListIsActive**| **bool** | The status of the list used, namely if it&#39;s deleted or disabled.  | [optional] [readonly]
**ContactStatus**| **string** | The status of the member, namely if they are subscribed, unsubscribed, deleted, non-subscribed, transactional, pending, or need reconfirmation.  | [optional] [readonly]
**EmailId**| **string** | The MD5 hash of the lowercase version of the list member&#39;s email address.  | [optional] [readonly]
**EmailAddress**| **string** | Email address for a subscriber.  | [optional] [readonly]
**MergeFields**| **map[string]map[string]interface{}** | A dictionary of merge fields where the keys are the merge tags. See the [Merge Fields documentation](https://mailchimp.com/developer/marketing/docs/merge-fields/#structure) for more about the structure.  | [optional]
**Vip**| **bool** | [VIP status](https://mailchimp.com/help/designate-and-send-to-vip-contacts/) for subscriber.  | [optional] [readonly]
**OpensCount**| **int32** | The total number of times the this campaign was opened by the list member.  | [optional] [readonly]
**Opens**| [**[]MemberActivity1**](MemberActivity1.md) | An array of timestamps for each time a list member opened the campaign. If a list member opens an email multiple times, this will return a separate timestamp for each open event.  | [optional]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

