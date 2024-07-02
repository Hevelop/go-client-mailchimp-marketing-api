# # SendingSchedule1
The schedule for sending the RSS Campaign.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hour**| **int32** | The hour to send the campaign in local time. Acceptable hours are 0-23. For example, &#39;4&#39; would be 4am in [your account&#39;s default time zone](https://mailchimp.com/help/set-account-details/).  | [optional]
**DailySend**| [**DailySendingDays**](DailySendingDays.md) |   | [optional]
**WeeklySendDay**| **string** | The day of the week to send a weekly RSS Campaign. for more information please, see Model/string.php  | [optional]
**MonthlySendDate**| **float32** | The day of the month to send a monthly RSS Campaign. Acceptable days are 0-31, where &#39;0&#39; is always the last day of a month. Months with fewer than the selected number of days will not have an RSS campaign sent out that day. For example, RSS Campaigns set to send on the 30th will not go out in February.  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

