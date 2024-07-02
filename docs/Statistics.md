# # Statistics
Stats for the list. Many of these are cached for at least five minutes.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberCount**| **int32** | The number of active members in the list.  | [optional] [readonly]
**TotalContacts**| **int32** | The number of contacts in the list, including subscribed, unsubscribed, pending, cleaned, deleted, transactional, and those that need to be reconfirmed. Requires include_total_contacts query parameter to be included.  | [optional] [readonly]
**UnsubscribeCount**| **int32** | The number of members who have unsubscribed from the list.  | [optional] [readonly]
**CleanedCount**| **int32** | The number of members cleaned from the list.  | [optional] [readonly]
**MemberCountSinceSend**| **int32** | The number of active members in the list since the last campaign was sent.  | [optional] [readonly]
**UnsubscribeCountSinceSend**| **int32** | The number of members who have unsubscribed since the last campaign was sent.  | [optional] [readonly]
**CleanedCountSinceSend**| **int32** | The number of members cleaned from the list since the last campaign was sent.  | [optional] [readonly]
**CampaignCount**| **int32** | The number of campaigns in any status that use this list.  | [optional] [readonly]
**CampaignLastSent**| [**time.Time**](time.Time.md) | The date and time the last campaign was sent to this list in ISO 8601 format. This is updated when a campaign is sent to 10 or more recipients.  | [optional] [readonly]
**MergeFieldCount**| **int32** | The number of merge fields ([audience field](https://mailchimp.com/help/getting-started-with-merge-tags/)) for this list (doesn&#39;t include EMAIL).  | [optional] [readonly]
**AvgSubRate**| **float32** | The average number of subscriptions per month for the list (not returned if we haven&#39;t calculated it yet).  | [optional] [readonly]
**AvgUnsubRate**| **float32** | The average number of unsubscriptions per month for the list (not returned if we haven&#39;t calculated it yet).  | [optional] [readonly]
**TargetSubRate**| **float32** | The target number of subscriptions per month for the list to keep it growing (not returned if we haven&#39;t calculated it yet).  | [optional] [readonly]
**OpenRate**| **float32** | The average open rate (a percentage represented as a number between 0 and 100) per campaign for the list (not returned if we haven&#39;t calculated it yet).  | [optional] [readonly]
**ClickRate**| **float32** | The average click rate (a percentage represented as a number between 0 and 100) per campaign for the list (not returned if we haven&#39;t calculated it yet).  | [optional] [readonly]
**LastSubDate**| [**time.Time**](time.Time.md) | The date and time of the last time someone subscribed to this list in ISO 8601 format.  | [optional] [readonly]
**LastUnsubDate**| [**time.Time**](time.Time.md) | The date and time of the last time someone unsubscribed from this list in ISO 8601 format.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

