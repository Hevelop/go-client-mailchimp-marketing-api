# Hevelop\MailchimpMarketingApi\BatchesAPI

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBatchesId**](BatchesAPI.md#DeleteBatchesId) | **Delete** /batches/{batch_id} | Delete batch request
[**GetBatches**](BatchesAPI.md#GetBatches) | **Get** /batches | List batch requests
[**GetBatchesId**](BatchesAPI.md#GetBatchesId) | **Get** /batches/{batch_id} | Get batch operation status
[**PostBatches**](BatchesAPI.md#PostBatches) | **Post** /batches | Start batch operation



## DeleteBatchesId

> DeleteBatchesId(ctx, batchId).Execute()

Delete batch request



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
	batchId := "batchId_example" // string | The unique id for the batch operation.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BatchesAPI.DeleteBatchesId(context.Background(), batchId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchesAPI.DeleteBatchesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchId** | **string** | The unique id for the batch operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBatchesIdRequest struct via the builder pattern


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


## GetBatches

> BatchOperations GetBatches(ctx).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()

List batch requests



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
	resp, r, err := apiClient.BatchesAPI.GetBatches(context.Background()).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchesAPI.GetBatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBatches`: BatchOperations
	fmt.Fprintf(os.Stdout, "Response from `BatchesAPI.GetBatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**BatchOperations**](BatchOperations.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBatchesId

> Batch GetBatchesId(ctx, batchId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get batch operation status



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
	batchId := "batchId_example" // string | The unique id for the batch operation.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchesAPI.GetBatchesId(context.Background(), batchId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchesAPI.GetBatchesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBatchesId`: Batch
	fmt.Fprintf(os.Stdout, "Response from `BatchesAPI.GetBatchesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchId** | **string** | The unique id for the batch operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBatchesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**Batch**](Batch.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBatches

> Batch PostBatches(ctx).Body(body).Execute()

Start batch operation



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
	body := *openapiclient.NewPostBatchesRequest([]openapiclient.Operations{*openapiclient.NewOperations("POST", "/lists")}) // PostBatchesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchesAPI.PostBatches(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchesAPI.PostBatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBatches`: Batch
	fmt.Fprintf(os.Stdout, "Response from `BatchesAPI.PostBatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PostBatchesRequest**](PostBatchesRequest.md) |  | 

### Return type

[**Batch**](Batch.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

