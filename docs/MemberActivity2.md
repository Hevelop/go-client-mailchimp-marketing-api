# # MemberActivity2
A summary of the interaction with the campaign.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action**| **string** | One of the following actions: &#39;open&#39;, &#39;click&#39;, or &#39;bounce&#39;  | [optional]
**Type**| **string** | If the action is a &#39;bounce&#39;, the type of bounce received: &#39;hard&#39;, &#39;soft&#39;.  | [optional]
**Timestamp**| [**time.Time**](time.Time.md) | The date and time recorded for the action in ISO 8601 format.  | [optional]
**Url**| **string** | If the action is a &#39;click&#39;, the URL on which the member clicked.  | [optional]
**Ip**| **string** | The IP address recorded for the action.  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

