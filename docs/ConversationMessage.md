# # ConversationMessage
An individual message in a conversation. Conversation tracking is a feature available to paid accounts that lets you view replies to your campaigns in your Mailchimp account.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **string** | A string that uniquely identifies this message  | [optional] [readonly]
**ConversationId**| **string** | A string that identifies this message&#39;s conversation  | [optional] [readonly]
**ListId**| **int32** | The list&#39;s web ID  | [optional] [readonly]
**FromLabel**| **string** | A label representing the sender of this message  | [optional] [readonly]
**FromEmail**| **string** | A label representing the email of the sender of this message  | [optional]
**Subject**| **string** | The subject of this message  | [optional]
**Message**| **string** | The plain-text content of the message  | [optional]
**Read**| **bool** | Whether this message has been marked as read  | [optional]
**Timestamp**| [**time.Time**](time.Time.md) | The date and time the message was either sent or received in ISO 8601 format.  | [optional] [readonly]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

