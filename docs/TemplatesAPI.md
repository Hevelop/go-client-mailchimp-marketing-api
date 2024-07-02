# Hevelop\MailchimpMarketingApi\TemplatesAPI

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTemplatesId**](TemplatesAPI.md#DeleteTemplatesId) | **Delete** /templates/{template_id} | Delete template
[**GetTemplates**](TemplatesAPI.md#GetTemplates) | **Get** /templates | List templates
[**GetTemplatesId**](TemplatesAPI.md#GetTemplatesId) | **Get** /templates/{template_id} | Get template info
[**GetTemplatesIdDefaultContent**](TemplatesAPI.md#GetTemplatesIdDefaultContent) | **Get** /templates/{template_id}/default-content | View default content
[**PatchTemplatesId**](TemplatesAPI.md#PatchTemplatesId) | **Patch** /templates/{template_id} | Update template
[**PostTemplates**](TemplatesAPI.md#PostTemplates) | **Post** /templates | Add template



## DeleteTemplatesId

> DeleteTemplatesId(ctx, templateId).Execute()

Delete template



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
	templateId := "templateId_example" // string | The unique id for the template.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TemplatesAPI.DeleteTemplatesId(context.Background(), templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.DeleteTemplatesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The unique id for the template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTemplatesIdRequest struct via the builder pattern


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


## GetTemplates

> Templates GetTemplates(ctx).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).CreatedBy(createdBy).SinceDateCreated(sinceDateCreated).BeforeDateCreated(beforeDateCreated).Type_(type_).Category(category).FolderId(folderId).SortField(sortField).ContentType(contentType).SortDir(sortDir).Execute()

List templates



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
	createdBy := "createdBy_example" // string | The Mailchimp account user who created the template. (optional)
	sinceDateCreated := "sinceDateCreated_example" // string | Restrict the response to templates created after the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. (optional)
	beforeDateCreated := "beforeDateCreated_example" // string | Restrict the response to templates created before the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. (optional)
	type_ := "type__example" // string | Limit results based on template type. (optional)
	category := "category_example" // string | Limit results based on category. (optional)
	folderId := "folderId_example" // string | The unique folder id. (optional)
	sortField := "sortField_example" // string | Returns user templates sorted by the specified field. (optional)
	contentType := "contentType_example" // string | Limit results based on how the template's content is put together. Only templates of type `user` can be filtered by `content_type`. If you want to retrieve saved templates created with the legacy email editor, then filter `content_type` to `template`. If you'd rather pull your saved templates for the new editor, filter to `multichannel`. For code your own templates, filter to `html`. (optional)
	sortDir := "sortDir_example" // string | Determines the order direction for sorted results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.GetTemplates(context.Background()).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).CreatedBy(createdBy).SinceDateCreated(sinceDateCreated).BeforeDateCreated(beforeDateCreated).Type_(type_).Category(category).FolderId(folderId).SortField(sortField).ContentType(contentType).SortDir(sortDir).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.GetTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplates`: Templates
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.GetTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **createdBy** | **string** | The Mailchimp account user who created the template. | 
 **sinceDateCreated** | **string** | Restrict the response to templates created after the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **beforeDateCreated** | **string** | Restrict the response to templates created before the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **type_** | **string** | Limit results based on template type. | 
 **category** | **string** | Limit results based on category. | 
 **folderId** | **string** | The unique folder id. | 
 **sortField** | **string** | Returns user templates sorted by the specified field. | 
 **contentType** | **string** | Limit results based on how the template&#39;s content is put together. Only templates of type &#x60;user&#x60; can be filtered by &#x60;content_type&#x60;. If you want to retrieve saved templates created with the legacy email editor, then filter &#x60;content_type&#x60; to &#x60;template&#x60;. If you&#39;d rather pull your saved templates for the new editor, filter to &#x60;multichannel&#x60;. For code your own templates, filter to &#x60;html&#x60;. | 
 **sortDir** | **string** | Determines the order direction for sorted results. | 

### Return type

[**Templates**](Templates.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplatesId

> TemplateInstance GetTemplatesId(ctx, templateId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get template info



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
	templateId := "templateId_example" // string | The unique id for the template.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.GetTemplatesId(context.Background(), templateId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.GetTemplatesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplatesId`: TemplateInstance
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.GetTemplatesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The unique id for the template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplatesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**TemplateInstance**](TemplateInstance.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplatesIdDefaultContent

> TemplateDefaultContent GetTemplatesIdDefaultContent(ctx, templateId).Fields(fields).ExcludeFields(excludeFields).Execute()

View default content



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
	templateId := "templateId_example" // string | The unique id for the template.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.GetTemplatesIdDefaultContent(context.Background(), templateId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.GetTemplatesIdDefaultContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplatesIdDefaultContent`: TemplateDefaultContent
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.GetTemplatesIdDefaultContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The unique id for the template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplatesIdDefaultContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**TemplateDefaultContent**](TemplateDefaultContent.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchTemplatesId

> TemplateInstance PatchTemplatesId(ctx, templateId).Body(body).Execute()

Update template



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
	templateId := "templateId_example" // string | The unique id for the template.
	body := *openapiclient.NewTemplateInstance1("Freddie's Jokes", "Html_example") // TemplateInstance1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.PatchTemplatesId(context.Background(), templateId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.PatchTemplatesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchTemplatesId`: TemplateInstance
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.PatchTemplatesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The unique id for the template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchTemplatesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**TemplateInstance1**](TemplateInstance1.md) |  | 

### Return type

[**TemplateInstance**](TemplateInstance.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTemplates

> TemplateInstance PostTemplates(ctx).Body(body).Execute()

Add template



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
	body := *openapiclient.NewTemplateInstance1("Freddie's Jokes", "Html_example") // TemplateInstance1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.PostTemplates(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.PostTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTemplates`: TemplateInstance
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.PostTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TemplateInstance1**](TemplateInstance1.md) |  | 

### Return type

[**TemplateInstance**](TemplateInstance.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

