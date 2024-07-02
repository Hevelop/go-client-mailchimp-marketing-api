# # ABTestOptions
The settings specific to A/B test campaigns.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WinningCombinationId**| **string** | ID for the winning combination.  | [optional] [readonly]
**WinningCampaignId**| **string** | ID of the campaign that was sent to the remaining recipients based on the winning combination.  | [optional] [readonly]
**WinnerCriteria**| **string** | The combination that performs the best. This may be determined automatically by click rate, open rate, or total revenue -- or you may choose manually based on the reporting data you find the most valuable. For Multivariate Campaigns testing send_time, winner_criteria is ignored. For Multivariate Campaigns with &#39;manual&#39; as the winner_criteria, the winner must be chosen in the Mailchimp web application. for more information please, see Model/string.php  | [optional]
**WaitTime**| **int32** | The number of minutes to wait before choosing the winning campaign. The value of wait_time must be greater than 0 and in whole hours, specified in minutes.  | [optional]
**TestSize**| **int32** | The percentage of recipients to send the test combinations to, must be a value between 10 and 100.  | [optional]
**SubjectLines**| **[]string** | The possible subject lines to test. If no subject lines are provided, settings.subject_line will be used.  | [optional]
**SendTimes**| [**[]time.Time**](time.Time.md) | The possible send times to test. The times provided should be in the format YYYY-MM-DD HH:MM:SS. If send_times are provided to test, the test_size will be set to 100% and winner_criteria will be ignored.  | [optional]
**FromNames**| **[]string** | The possible from names. The number of from_names provided must match the number of reply_to_addresses. If no from_names are provided, settings.from_name will be used.  | [optional]
**ReplyToAddresses**| **[]string** | The possible reply-to addresses. The number of reply_to_addresses provided must match the number of from_names. If no reply_to_addresses are provided, settings.reply_to will be used.  | [optional]
**Contents**| **[]string** | Descriptions of possible email contents. To set campaign contents, make a PUT request to /campaigns/{campaign_id}/content with the field &#39;variate_contents&#39;.  | [optional] [readonly]
**Combinations**| [**[]CombinationsInner**](CombinationsInner.md) | Combinations of possible variables used to build emails.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

