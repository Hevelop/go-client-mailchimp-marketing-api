# Hevelop\MailchimpMarketingApi\LandingPagesAPI

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLandingPageId**](LandingPagesAPI.md#DeleteLandingPageId) | **Delete** /landing-pages/{page_id} | Delete landing page
[**GetAllLandingPages**](LandingPagesAPI.md#GetAllLandingPages) | **Get** /landing-pages | List landing pages
[**GetLandingPageId**](LandingPagesAPI.md#GetLandingPageId) | **Get** /landing-pages/{page_id} | Get landing page info
[**GetLandingPageIdContent**](LandingPagesAPI.md#GetLandingPageIdContent) | **Get** /landing-pages/{page_id}/content | Get landing page content
[**PatchLandingPageId**](LandingPagesAPI.md#PatchLandingPageId) | **Patch** /landing-pages/{page_id} | Update landing page
[**PostAllLandingPages**](LandingPagesAPI.md#PostAllLandingPages) | **Post** /landing-pages | Add landing page
[**PostLandingPageIdActionsPublish**](LandingPagesAPI.md#PostLandingPageIdActionsPublish) | **Post** /landing-pages/{page_id}/actions/publish | Publish landing page
[**PostLandingPageIdActionsUnpublish**](LandingPagesAPI.md#PostLandingPageIdActionsUnpublish) | **Post** /landing-pages/{page_id}/actions/unpublish | Unpublish landing page



## DeleteLandingPageId

> DeleteLandingPageId(ctx, pageId).Execute()

Delete landing page



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
	pageId := "pageId_example" // string | The unique id for the page.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LandingPagesAPI.DeleteLandingPageId(context.Background(), pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LandingPagesAPI.DeleteLandingPageId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | The unique id for the page. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLandingPageIdRequest struct via the builder pattern


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


## GetAllLandingPages

> GetAllLandingPages200Response GetAllLandingPages(ctx).SortDir(sortDir).SortField(sortField).Fields(fields).ExcludeFields(excludeFields).Count(count).Execute()

List landing pages



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
	sortDir := "sortDir_example" // string | Determines the order direction for sorted results. (optional)
	sortField := "sortField_example" // string | Returns files sorted by the specified field. (optional)
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LandingPagesAPI.GetAllLandingPages(context.Background()).SortDir(sortDir).SortField(sortField).Fields(fields).ExcludeFields(excludeFields).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LandingPagesAPI.GetAllLandingPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllLandingPages`: GetAllLandingPages200Response
	fmt.Fprintf(os.Stdout, "Response from `LandingPagesAPI.GetAllLandingPages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllLandingPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortDir** | **string** | Determines the order direction for sorted results. | 
 **sortField** | **string** | Returns files sorted by the specified field. | 
 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]

### Return type

[**GetAllLandingPages200Response**](GetAllLandingPages200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLandingPageId

> LandingPage GetLandingPageId(ctx, pageId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get landing page info



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
	pageId := "pageId_example" // string | The unique id for the page.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LandingPagesAPI.GetLandingPageId(context.Background(), pageId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LandingPagesAPI.GetLandingPageId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLandingPageId`: LandingPage
	fmt.Fprintf(os.Stdout, "Response from `LandingPagesAPI.GetLandingPageId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | The unique id for the page. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLandingPageIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**LandingPage**](LandingPage.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLandingPageIdContent

> LandingPageContent GetLandingPageIdContent(ctx, pageId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get landing page content



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
	pageId := "pageId_example" // string | The unique id for the page.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LandingPagesAPI.GetLandingPageIdContent(context.Background(), pageId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LandingPagesAPI.GetLandingPageIdContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLandingPageIdContent`: LandingPageContent
	fmt.Fprintf(os.Stdout, "Response from `LandingPagesAPI.GetLandingPageIdContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | The unique id for the page. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLandingPageIdContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**LandingPageContent**](LandingPageContent.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchLandingPageId

> LandingPage PatchLandingPageId(ctx, pageId).Body(body).Execute()

Update landing page



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
	pageId := "pageId_example" // string | The unique id for the page.
	body := *openapiclient.NewLandingPage2() // LandingPage2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LandingPagesAPI.PatchLandingPageId(context.Background(), pageId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LandingPagesAPI.PatchLandingPageId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchLandingPageId`: LandingPage
	fmt.Fprintf(os.Stdout, "Response from `LandingPagesAPI.PatchLandingPageId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | The unique id for the page. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchLandingPageIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**LandingPage2**](LandingPage2.md) |  | 

### Return type

[**LandingPage**](LandingPage.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAllLandingPages

> LandingPage PostAllLandingPages(ctx).Body(body).UseDefaultList(useDefaultList).Execute()

Add landing page



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
	body := *openapiclient.NewLandingPage1() // LandingPage1 | 
	useDefaultList := true // bool | Will create the Landing Page using the account's Default List instead of requiring a list_id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LandingPagesAPI.PostAllLandingPages(context.Background()).Body(body).UseDefaultList(useDefaultList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LandingPagesAPI.PostAllLandingPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAllLandingPages`: LandingPage
	fmt.Fprintf(os.Stdout, "Response from `LandingPagesAPI.PostAllLandingPages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAllLandingPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LandingPage1**](LandingPage1.md) |  | 
 **useDefaultList** | **bool** | Will create the Landing Page using the account&#39;s Default List instead of requiring a list_id. | 

### Return type

[**LandingPage**](LandingPage.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLandingPageIdActionsPublish

> PostLandingPageIdActionsPublish(ctx, pageId).Execute()

Publish landing page



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
	pageId := "pageId_example" // string | The unique id for the page.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LandingPagesAPI.PostLandingPageIdActionsPublish(context.Background(), pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LandingPagesAPI.PostLandingPageIdActionsPublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | The unique id for the page. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostLandingPageIdActionsPublishRequest struct via the builder pattern


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


## PostLandingPageIdActionsUnpublish

> PostLandingPageIdActionsUnpublish(ctx, pageId).Execute()

Unpublish landing page



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
	pageId := "pageId_example" // string | The unique id for the page.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LandingPagesAPI.PostLandingPageIdActionsUnpublish(context.Background(), pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LandingPagesAPI.PostLandingPageIdActionsUnpublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | The unique id for the page. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostLandingPageIdActionsUnpublishRequest struct via the builder pattern


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

