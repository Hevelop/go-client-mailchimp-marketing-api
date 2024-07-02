# # DailyListActivity
One day&#39;s worth of list activity. Doesn&#39;t include Automation activity.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Day**| **string** | The date for the activity summary.  | [optional] [readonly]
**EmailsSent**| **int32** | The total number of emails sent on the date for the activity summary.  | [optional] [readonly]
**UniqueOpens**| **int32** | The number of unique opens.  | [optional] [readonly]
**RecipientClicks**| **int32** | The number of clicks.  | [optional] [readonly]
**HardBounce**| **int32** | The number of hard bounces.  | [optional] [readonly]
**SoftBounce**| **int32** | The number of soft bounces  | [optional] [readonly]
**Subs**| **int32** | The number of subscribes.  | [optional] [readonly]
**Unsubs**| **int32** | The number of unsubscribes.  | [optional] [readonly]
**OtherAdds**| **int32** | The number of subscribers who may have been added outside of the [double opt-in process](https://mailchimp.com/help/about-double-opt-in/), such as imports or API activity.  | [optional] [readonly]
**OtherRemoves**| **int32** | The number of subscribers who may have been removed outside of unsubscribing or reporting an email as spam (for example, deleted subscribers).  | [optional] [readonly]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

