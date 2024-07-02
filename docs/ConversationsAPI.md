# Hevelop\MailchimpMarketingApi\ConversationsAPI

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConversations**](ConversationsAPI.md#GetConversations) | **Get** /conversations | List conversations
[**GetConversationsId**](ConversationsAPI.md#GetConversationsId) | **Get** /conversations/{conversation_id} | Get conversation
[**GetConversationsIdMessages**](ConversationsAPI.md#GetConversationsIdMessages) | **Get** /conversations/{conversation_id}/messages | List messages
[**GetConversationsIdMessagesId**](ConversationsAPI.md#GetConversationsIdMessagesId) | **Get** /conversations/{conversation_id}/messages/{message_id} | Get message



## GetConversations

> TrackedConversations GetConversations(ctx).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).HasUnreadMessages(hasUnreadMessages).ListId(listId).CampaignId(campaignId).Execute()

List conversations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)
	hasUnreadMessages := "hasUnreadMessages_example" // string | Whether the conversation has any unread messages. (optional)
	listId := "listId_example" // string | The unique id for the list. (optional)
	campaignId := "campaignId_example" // string | The unique id for the campaign. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.GetConversations(context.Background()).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).HasUnreadMessages(hasUnreadMessages).ListId(listId).CampaignId(campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.GetConversations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConversations`: TrackedConversations
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.GetConversations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConversationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **hasUnreadMessages** | **string** | Whether the conversation has any unread messages. | 
 **listId** | **string** | The unique id for the list. | 
 **campaignId** | **string** | The unique id for the campaign. | 

### Return type

[**TrackedConversations**](TrackedConversations.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConversationsId

> Conversation GetConversationsId(ctx, conversationId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get conversation



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	conversationId := "conversationId_example" // string | The unique id for the conversation.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.GetConversationsId(context.Background(), conversationId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.GetConversationsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConversationsId`: Conversation
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.GetConversationsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** | The unique id for the conversation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConversationsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**Conversation**](Conversation.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConversationsIdMessages

> CollectionOfConversationMessages GetConversationsIdMessages(ctx, conversationId).Fields(fields).ExcludeFields(excludeFields).IsRead(isRead).BeforeTimestamp(beforeTimestamp).SinceTimestamp(sinceTimestamp).Execute()

List messages



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	conversationId := "conversationId_example" // string | The unique id for the conversation.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	isRead := "isRead_example" // string | Whether a conversation message has been marked as read. (optional)
	beforeTimestamp := time.Now() // time.Time | Restrict the response to messages created before the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. (optional)
	sinceTimestamp := time.Now() // time.Time | Restrict the response to messages created after the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.GetConversationsIdMessages(context.Background(), conversationId).Fields(fields).ExcludeFields(excludeFields).IsRead(isRead).BeforeTimestamp(beforeTimestamp).SinceTimestamp(sinceTimestamp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.GetConversationsIdMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConversationsIdMessages`: CollectionOfConversationMessages
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.GetConversationsIdMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** | The unique id for the conversation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConversationsIdMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **isRead** | **string** | Whether a conversation message has been marked as read. | 
 **beforeTimestamp** | **time.Time** | Restrict the response to messages created before the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **sinceTimestamp** | **time.Time** | Restrict the response to messages created after the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 

### Return type

[**CollectionOfConversationMessages**](CollectionOfConversationMessages.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConversationsIdMessagesId

> ConversationMessage GetConversationsIdMessagesId(ctx, conversationId, messageId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get message



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	conversationId := "conversationId_example" // string | The unique id for the conversation.
	messageId := "messageId_example" // string | The unique id for the conversation message.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.GetConversationsIdMessagesId(context.Background(), conversationId, messageId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.GetConversationsIdMessagesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConversationsIdMessagesId`: ConversationMessage
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.GetConversationsIdMessagesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** | The unique id for the conversation. | 
**messageId** | **string** | The unique id for the conversation message. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConversationsIdMessagesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**ConversationMessage**](ConversationMessage.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

