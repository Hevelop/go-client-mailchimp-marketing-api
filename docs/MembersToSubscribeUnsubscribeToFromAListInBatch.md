# # MembersToSubscribeUnsubscribeToFromAListInBatch
Members to subscribe to or unsubscribe from a list.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members**| [**[]AddListMembers**](AddListMembers.md) | An array of objects, each representing an email address and the subscription status for a specific list. Up to 500 members may be added or updated with each API call.  |
**SyncTags**| **bool** | Whether this batch operation will replace all existing tags with tags in request.  | [optional]
**UpdateExisting**| **bool** | Whether this batch operation will change existing members&#39; subscription status.  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

