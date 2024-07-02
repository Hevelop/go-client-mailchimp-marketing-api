# # LastMessage
The most recent message in the conversation.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromLabel**| **string** | A label representing the sender of this message.  | [optional] [readonly]
**FromEmail**| **string** | A label representing the email of the sender of this message.  | [optional] [readonly]
**Subject**| **string** | The subject of this message.  | [optional] [readonly]
**Message**| **string** | The plain-text content of the message.  | [optional] [readonly]
**Read**| **bool** | Whether this message has been marked as read.  | [optional]
**Timestamp**| [**time.Time**](time.Time.md) | The date and time the message was either sent or received in ISO 8601 format.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

