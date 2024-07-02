# # Conversation
Details about an individual conversation. Conversation tracking is a feature available to paid accounts that lets you view replies to your campaigns in your Mailchimp account.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **string** | A string that uniquely identifies this conversation.  | [optional] [readonly]
**MessageCount**| **int32** | The total number of messages in this conversation.  | [optional] [readonly]
**CampaignId**| **string** | The unique identifier of the campaign for this conversation.  | [optional] [readonly]
**ListId**| **string** | The unique identifier of the list for this conversation.  | [optional] [readonly]
**UnreadMessages**| **int32** | The number of unread messages in this conversation.  | [optional] [readonly]
**FromLabel**| **string** | A label representing the sender of this message.  | [optional] [readonly]
**FromEmail**| **string** | A label representing the email of the sender of this message.  | [optional] [readonly]
**Subject**| **string** | The subject of the message.  | [optional] [readonly]
**LastMessage**| [**LastMessage**](LastMessage.md) |   | [optional]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

