# Hevelop\MailchimpMarketingApi\TemplateFoldersAPI

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTemplateFoldersId**](TemplateFoldersAPI.md#DeleteTemplateFoldersId) | **Delete** /template-folders/{folder_id} | Delete template folder
[**GetTemplateFolders**](TemplateFoldersAPI.md#GetTemplateFolders) | **Get** /template-folders | List template folders
[**GetTemplateFoldersId**](TemplateFoldersAPI.md#GetTemplateFoldersId) | **Get** /template-folders/{folder_id} | Get template folder
[**PatchTemplateFoldersId**](TemplateFoldersAPI.md#PatchTemplateFoldersId) | **Patch** /template-folders/{folder_id} | Update template folder
[**PostTemplateFolders**](TemplateFoldersAPI.md#PostTemplateFolders) | **Post** /template-folders | Add template folder



## DeleteTemplateFoldersId

> DeleteTemplateFoldersId(ctx, folderId).Execute()

Delete template folder



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
	folderId := "folderId_example" // string | The unique id for the template folder.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TemplateFoldersAPI.DeleteTemplateFoldersId(context.Background(), folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplateFoldersAPI.DeleteTemplateFoldersId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** | The unique id for the template folder. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTemplateFoldersIdRequest struct via the builder pattern


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


## GetTemplateFolders

> TemplateFolders GetTemplateFolders(ctx).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()

List template folders



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
	resp, r, err := apiClient.TemplateFoldersAPI.GetTemplateFolders(context.Background()).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplateFoldersAPI.GetTemplateFolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplateFolders`: TemplateFolders
	fmt.Fprintf(os.Stdout, "Response from `TemplateFoldersAPI.GetTemplateFolders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**TemplateFolders**](TemplateFolders.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplateFoldersId

> TemplateFolder GetTemplateFoldersId(ctx, folderId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get template folder



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
	folderId := "folderId_example" // string | The unique id for the template folder.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplateFoldersAPI.GetTemplateFoldersId(context.Background(), folderId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplateFoldersAPI.GetTemplateFoldersId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplateFoldersId`: TemplateFolder
	fmt.Fprintf(os.Stdout, "Response from `TemplateFoldersAPI.GetTemplateFoldersId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** | The unique id for the template folder. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateFoldersIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**TemplateFolder**](TemplateFolder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchTemplateFoldersId

> TemplateFolder PatchTemplateFoldersId(ctx, folderId).Body(body).Execute()

Update template folder



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
	folderId := "folderId_example" // string | The unique id for the template folder.
	body := *openapiclient.NewTemplateFolder1("Name_example") // TemplateFolder1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplateFoldersAPI.PatchTemplateFoldersId(context.Background(), folderId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplateFoldersAPI.PatchTemplateFoldersId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchTemplateFoldersId`: TemplateFolder
	fmt.Fprintf(os.Stdout, "Response from `TemplateFoldersAPI.PatchTemplateFoldersId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** | The unique id for the template folder. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchTemplateFoldersIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**TemplateFolder1**](TemplateFolder1.md) |  | 

### Return type

[**TemplateFolder**](TemplateFolder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTemplateFolders

> TemplateFolder PostTemplateFolders(ctx).Body(body).Execute()

Add template folder



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
	body := *openapiclient.NewTemplateFolder1("Name_example") // TemplateFolder1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplateFoldersAPI.PostTemplateFolders(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplateFoldersAPI.PostTemplateFolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTemplateFolders`: TemplateFolder
	fmt.Fprintf(os.Stdout, "Response from `TemplateFoldersAPI.PostTemplateFolders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTemplateFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TemplateFolder1**](TemplateFolder1.md) |  | 

### Return type

[**TemplateFolder**](TemplateFolder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

