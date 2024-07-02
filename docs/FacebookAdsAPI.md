# Hevelop\MailchimpMarketingApi\FacebookAdsAPI

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllFacebookAds**](FacebookAdsAPI.md#GetAllFacebookAds) | **Get** /facebook-ads | List facebook ads
[**GetFacebookAdsId**](FacebookAdsAPI.md#GetFacebookAdsId) | **Get** /facebook-ads/{outreach_id} | Get facebook ad info



## GetAllFacebookAds

> GetAllFacebookAds200Response GetAllFacebookAds(ctx).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).SortField(sortField).SortDir(sortDir).Execute()

List facebook ads



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
	sortField := "sortField_example" // string | Returns files sorted by the specified field. (optional)
	sortDir := "sortDir_example" // string | Determines the order direction for sorted results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FacebookAdsAPI.GetAllFacebookAds(context.Background()).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).SortField(sortField).SortDir(sortDir).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FacebookAdsAPI.GetAllFacebookAds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllFacebookAds`: GetAllFacebookAds200Response
	fmt.Fprintf(os.Stdout, "Response from `FacebookAdsAPI.GetAllFacebookAds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllFacebookAdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **sortField** | **string** | Returns files sorted by the specified field. | 
 **sortDir** | **string** | Determines the order direction for sorted results. | 

### Return type

[**GetAllFacebookAds200Response**](GetAllFacebookAds200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFacebookAdsId

> GetFacebookAdsId200Response GetFacebookAdsId(ctx, outreachId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get facebook ad info



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
	outreachId := "outreachId_example" // string | The outreach id.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FacebookAdsAPI.GetFacebookAdsId(context.Background(), outreachId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FacebookAdsAPI.GetFacebookAdsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFacebookAdsId`: GetFacebookAdsId200Response
	fmt.Fprintf(os.Stdout, "Response from `FacebookAdsAPI.GetFacebookAdsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**outreachId** | **string** | The outreach id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFacebookAdsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**GetFacebookAdsId200Response**](GetFacebookAdsId200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

