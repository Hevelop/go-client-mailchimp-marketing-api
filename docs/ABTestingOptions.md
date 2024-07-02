# # ABTestingOptions
[A/B Testing](https://mailchimp.com/help/about-ab-testing-campaigns/) options for a campaign.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SplitTest**| **string** | The type of AB split to run. for more information please, see Model/string.php  | [optional]
**PickWinner**| **string** | How we should evaluate a winner. Based on &#39;opens&#39;, &#39;clicks&#39;, or &#39;manual&#39;. for more information please, see Model/string.php  | [optional]
**WaitUnits**| **string** | How unit of time for measuring the winner (&#39;hours&#39; or &#39;days&#39;). This cannot be changed after a campaign is sent. for more information please, see Model/string.php  | [optional]
**WaitTime**| **int32** | The amount of time to wait before picking a winner. This cannot be changed after a campaign is sent.  | [optional]
**SplitSize**| **int32** | The size of the split groups. Campaigns split based on &#39;schedule&#39; are forced to have a 50/50 split. Valid split integers are between 1-50.  | [optional]
**FromNameA**| **string** | For campaigns split on &#39;From Name&#39;, the name for Group A.  | [optional]
**FromNameB**| **string** | For campaigns split on &#39;From Name&#39;, the name for Group B.  | [optional]
**ReplyEmailA**| **string** | For campaigns split on &#39;From Name&#39;, the reply-to address for Group A.  | [optional]
**ReplyEmailB**| **string** | For campaigns split on &#39;From Name&#39;, the reply-to address for Group B.  | [optional]
**SubjectA**| **string** | For campaigns split on &#39;Subject Line&#39;, the subject line for Group A.  | [optional]
**SubjectB**| **string** | For campaigns split on &#39;Subject Line&#39;, the subject line for Group B.  | [optional]
**SendTimeA**| [**time.Time**](time.Time.md) | The send time for Group A.  | [optional]
**SendTimeB**| [**time.Time**](time.Time.md) | The send time for Group B.  | [optional]
**SendTimeWinner**| **string** | The send time for the winning version.  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

