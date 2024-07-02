# Hevelop\MailchimpMarketingApi\BatchWebhooksAPI

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBatchWebhookId**](BatchWebhooksAPI.md#DeleteBatchWebhookId) | **Delete** /batch-webhooks/{batch_webhook_id} | Delete batch webhook
[**GetBatchWebhook**](BatchWebhooksAPI.md#GetBatchWebhook) | **Get** /batch-webhooks/{batch_webhook_id} | Get batch webhook info
[**GetBatchWebhooks**](BatchWebhooksAPI.md#GetBatchWebhooks) | **Get** /batch-webhooks | List batch webhooks
[**PatchBatchWebhooks**](BatchWebhooksAPI.md#PatchBatchWebhooks) | **Patch** /batch-webhooks/{batch_webhook_id} | Update batch webhook
[**PostBatchWebhooks**](BatchWebhooksAPI.md#PostBatchWebhooks) | **Post** /batch-webhooks | Add batch webhook



## DeleteBatchWebhookId

> DeleteBatchWebhookId(ctx, batchWebhookId).Execute()

Delete batch webhook



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
	batchWebhookId := "batchWebhookId_example" // string | The unique id for the batch webhook.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BatchWebhooksAPI.DeleteBatchWebhookId(context.Background(), batchWebhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchWebhooksAPI.DeleteBatchWebhookId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchWebhookId** | **string** | The unique id for the batch webhook. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBatchWebhookIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBatchWebhook

> BatchWebhook GetBatchWebhook(ctx, batchWebhookId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get batch webhook info



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
	batchWebhookId := "batchWebhookId_example" // string | The unique id for the batch webhook.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchWebhooksAPI.GetBatchWebhook(context.Background(), batchWebhookId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchWebhooksAPI.GetBatchWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBatchWebhook`: BatchWebhook
	fmt.Fprintf(os.Stdout, "Response from `BatchWebhooksAPI.GetBatchWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchWebhookId** | **string** | The unique id for the batch webhook. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBatchWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**BatchWebhook**](BatchWebhook.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBatchWebhooks

> BatchWebhooks GetBatchWebhooks(ctx).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()

List batch webhooks



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchWebhooksAPI.GetBatchWebhooks(context.Background()).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchWebhooksAPI.GetBatchWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBatchWebhooks`: BatchWebhooks
	fmt.Fprintf(os.Stdout, "Response from `BatchWebhooksAPI.GetBatchWebhooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBatchWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**BatchWebhooks**](BatchWebhooks.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchBatchWebhooks

> BatchWebhook PatchBatchWebhooks(ctx, batchWebhookId).Body(body).Execute()

Update batch webhook



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
	batchWebhookId := "batchWebhookId_example" // string | The unique id for the batch webhook.
	body := *openapiclient.NewBatchWebhook2() // BatchWebhook2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchWebhooksAPI.PatchBatchWebhooks(context.Background(), batchWebhookId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchWebhooksAPI.PatchBatchWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchBatchWebhooks`: BatchWebhook
	fmt.Fprintf(os.Stdout, "Response from `BatchWebhooksAPI.PatchBatchWebhooks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchWebhookId** | **string** | The unique id for the batch webhook. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchBatchWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**BatchWebhook2**](BatchWebhook2.md) |  | 

### Return type

[**BatchWebhook**](BatchWebhook.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBatchWebhooks

> BatchWebhook PostBatchWebhooks(ctx).Body(body).Execute()

Add batch webhook



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
	body := *openapiclient.NewBatchWebhook1("http://yourdomain.com/webhook") // BatchWebhook1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchWebhooksAPI.PostBatchWebhooks(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchWebhooksAPI.PostBatchWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBatchWebhooks`: BatchWebhook
	fmt.Fprintf(os.Stdout, "Response from `BatchWebhooksAPI.PostBatchWebhooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBatchWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**BatchWebhook1**](BatchWebhook1.md) |  | 

### Return type

[**BatchWebhook**](BatchWebhook.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

