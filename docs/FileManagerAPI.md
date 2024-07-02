# Hevelop\MailchimpMarketingApi\FileManagerAPI

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFileManagerFilesId**](FileManagerAPI.md#DeleteFileManagerFilesId) | **Delete** /file-manager/files/{file_id} | Delete file
[**DeleteFileManagerFoldersId**](FileManagerAPI.md#DeleteFileManagerFoldersId) | **Delete** /file-manager/folders/{folder_id} | Delete folder
[**GetFileManagerFiles**](FileManagerAPI.md#GetFileManagerFiles) | **Get** /file-manager/files | List stored files
[**GetFileManagerFilesId**](FileManagerAPI.md#GetFileManagerFilesId) | **Get** /file-manager/files/{file_id} | Get file
[**GetFileManagerFolders**](FileManagerAPI.md#GetFileManagerFolders) | **Get** /file-manager/folders | List folders
[**GetFileManagerFoldersFiles**](FileManagerAPI.md#GetFileManagerFoldersFiles) | **Get** /file-manager/folders/{folder_id}/files | List stored files
[**GetFileManagerFoldersId**](FileManagerAPI.md#GetFileManagerFoldersId) | **Get** /file-manager/folders/{folder_id} | Get folder
[**PatchFileManagerFilesId**](FileManagerAPI.md#PatchFileManagerFilesId) | **Patch** /file-manager/files/{file_id} | Update file
[**PatchFileManagerFoldersId**](FileManagerAPI.md#PatchFileManagerFoldersId) | **Patch** /file-manager/folders/{folder_id} | Update folder
[**PostFileManagerFiles**](FileManagerAPI.md#PostFileManagerFiles) | **Post** /file-manager/files | Add file
[**PostFileManagerFolders**](FileManagerAPI.md#PostFileManagerFolders) | **Post** /file-manager/folders | Add folder



## DeleteFileManagerFilesId

> DeleteFileManagerFilesId(ctx, fileId).Execute()

Delete file



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
	fileId := "fileId_example" // string | The unique id for the File Manager file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FileManagerAPI.DeleteFileManagerFilesId(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileManagerAPI.DeleteFileManagerFilesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | The unique id for the File Manager file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFileManagerFilesIdRequest struct via the builder pattern


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


## DeleteFileManagerFoldersId

> DeleteFileManagerFoldersId(ctx, folderId).Execute()

Delete folder



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
	folderId := "folderId_example" // string | The unique id for the File Manager folder.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FileManagerAPI.DeleteFileManagerFoldersId(context.Background(), folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileManagerAPI.DeleteFileManagerFoldersId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** | The unique id for the File Manager folder. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFileManagerFoldersIdRequest struct via the builder pattern


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


## GetFileManagerFiles

> FileManager GetFileManagerFiles(ctx).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Type_(type_).CreatedBy(createdBy).BeforeCreatedAt(beforeCreatedAt).SinceCreatedAt(sinceCreatedAt).SortField(sortField).SortDir(sortDir).Execute()

List stored files



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
	type_ := "type__example" // string | The file type for the File Manager file. (optional)
	createdBy := "createdBy_example" // string | The Mailchimp account user who created the File Manager file. (optional)
	beforeCreatedAt := "beforeCreatedAt_example" // string | Restrict the response to files created before the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. (optional)
	sinceCreatedAt := "sinceCreatedAt_example" // string | Restrict the response to files created after the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. (optional)
	sortField := "sortField_example" // string | Returns files sorted by the specified field. (optional)
	sortDir := "sortDir_example" // string | Determines the order direction for sorted results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileManagerAPI.GetFileManagerFiles(context.Background()).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Type_(type_).CreatedBy(createdBy).BeforeCreatedAt(beforeCreatedAt).SinceCreatedAt(sinceCreatedAt).SortField(sortField).SortDir(sortDir).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileManagerAPI.GetFileManagerFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileManagerFiles`: FileManager
	fmt.Fprintf(os.Stdout, "Response from `FileManagerAPI.GetFileManagerFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFileManagerFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **type_** | **string** | The file type for the File Manager file. | 
 **createdBy** | **string** | The Mailchimp account user who created the File Manager file. | 
 **beforeCreatedAt** | **string** | Restrict the response to files created before the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **sinceCreatedAt** | **string** | Restrict the response to files created after the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **sortField** | **string** | Returns files sorted by the specified field. | 
 **sortDir** | **string** | Determines the order direction for sorted results. | 

### Return type

[**FileManager**](FileManager.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileManagerFilesId

> GalleryFile GetFileManagerFilesId(ctx, fileId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get file



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
	fileId := "fileId_example" // string | The unique id for the File Manager file.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileManagerAPI.GetFileManagerFilesId(context.Background(), fileId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileManagerAPI.GetFileManagerFilesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileManagerFilesId`: GalleryFile
	fmt.Fprintf(os.Stdout, "Response from `FileManagerAPI.GetFileManagerFilesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | The unique id for the File Manager file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileManagerFilesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**GalleryFile**](GalleryFile.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileManagerFolders

> FileManagerFolders GetFileManagerFolders(ctx).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).CreatedBy(createdBy).BeforeCreatedAt(beforeCreatedAt).SinceCreatedAt(sinceCreatedAt).Execute()

List folders



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
	createdBy := "createdBy_example" // string | The Mailchimp account user who created the File Manager file. (optional)
	beforeCreatedAt := "beforeCreatedAt_example" // string | Restrict the response to files created before the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. (optional)
	sinceCreatedAt := "sinceCreatedAt_example" // string | Restrict the response to files created after the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileManagerAPI.GetFileManagerFolders(context.Background()).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).CreatedBy(createdBy).BeforeCreatedAt(beforeCreatedAt).SinceCreatedAt(sinceCreatedAt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileManagerAPI.GetFileManagerFolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileManagerFolders`: FileManagerFolders
	fmt.Fprintf(os.Stdout, "Response from `FileManagerAPI.GetFileManagerFolders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFileManagerFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **createdBy** | **string** | The Mailchimp account user who created the File Manager file. | 
 **beforeCreatedAt** | **string** | Restrict the response to files created before the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **sinceCreatedAt** | **string** | Restrict the response to files created after the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 

### Return type

[**FileManagerFolders**](FileManagerFolders.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileManagerFoldersFiles

> FileManager GetFileManagerFoldersFiles(ctx, folderId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Type_(type_).CreatedBy(createdBy).BeforeCreatedAt(beforeCreatedAt).SinceCreatedAt(sinceCreatedAt).SortField(sortField).SortDir(sortDir).Execute()

List stored files



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
	folderId := "folderId_example" // string | The unique id for the File Manager folder.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)
	type_ := "type__example" // string | The file type for the File Manager file. (optional)
	createdBy := "createdBy_example" // string | The Mailchimp account user who created the File Manager file. (optional)
	beforeCreatedAt := "beforeCreatedAt_example" // string | Restrict the response to files created before the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. (optional)
	sinceCreatedAt := "sinceCreatedAt_example" // string | Restrict the response to files created after the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. (optional)
	sortField := "sortField_example" // string | Returns files sorted by the specified field. (optional)
	sortDir := "sortDir_example" // string | Determines the order direction for sorted results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileManagerAPI.GetFileManagerFoldersFiles(context.Background(), folderId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Type_(type_).CreatedBy(createdBy).BeforeCreatedAt(beforeCreatedAt).SinceCreatedAt(sinceCreatedAt).SortField(sortField).SortDir(sortDir).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileManagerAPI.GetFileManagerFoldersFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileManagerFoldersFiles`: FileManager
	fmt.Fprintf(os.Stdout, "Response from `FileManagerAPI.GetFileManagerFoldersFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** | The unique id for the File Manager folder. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileManagerFoldersFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **type_** | **string** | The file type for the File Manager file. | 
 **createdBy** | **string** | The Mailchimp account user who created the File Manager file. | 
 **beforeCreatedAt** | **string** | Restrict the response to files created before the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **sinceCreatedAt** | **string** | Restrict the response to files created after the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **sortField** | **string** | Returns files sorted by the specified field. | 
 **sortDir** | **string** | Determines the order direction for sorted results. | 

### Return type

[**FileManager**](FileManager.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileManagerFoldersId

> GalleryFolder GetFileManagerFoldersId(ctx, folderId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get folder



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
	folderId := "folderId_example" // string | The unique id for the File Manager folder.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileManagerAPI.GetFileManagerFoldersId(context.Background(), folderId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileManagerAPI.GetFileManagerFoldersId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileManagerFoldersId`: GalleryFolder
	fmt.Fprintf(os.Stdout, "Response from `FileManagerAPI.GetFileManagerFoldersId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** | The unique id for the File Manager folder. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileManagerFoldersIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**GalleryFolder**](GalleryFolder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchFileManagerFilesId

> GalleryFile PatchFileManagerFilesId(ctx, fileId).Body(body).Execute()

Update file



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
	fileId := "fileId_example" // string | The unique id for the File Manager file.
	body := *openapiclient.NewGalleryFile2() // GalleryFile2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileManagerAPI.PatchFileManagerFilesId(context.Background(), fileId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileManagerAPI.PatchFileManagerFilesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFileManagerFilesId`: GalleryFile
	fmt.Fprintf(os.Stdout, "Response from `FileManagerAPI.PatchFileManagerFilesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | The unique id for the File Manager file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFileManagerFilesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**GalleryFile2**](GalleryFile2.md) |  | 

### Return type

[**GalleryFile**](GalleryFile.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchFileManagerFoldersId

> GalleryFolder PatchFileManagerFoldersId(ctx, folderId).Body(body).Execute()

Update folder



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
	folderId := "folderId_example" // string | The unique id for the File Manager folder.
	body := *openapiclient.NewGalleryFolder1("Name_example") // GalleryFolder1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileManagerAPI.PatchFileManagerFoldersId(context.Background(), folderId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileManagerAPI.PatchFileManagerFoldersId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFileManagerFoldersId`: GalleryFolder
	fmt.Fprintf(os.Stdout, "Response from `FileManagerAPI.PatchFileManagerFoldersId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** | The unique id for the File Manager folder. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFileManagerFoldersIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**GalleryFolder1**](GalleryFolder1.md) |  | 

### Return type

[**GalleryFolder**](GalleryFolder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFileManagerFiles

> GalleryFile PostFileManagerFiles(ctx).Body(body).Execute()

Add file



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
	body := *openapiclient.NewGalleryFile1("Name_example", "FileData_example") // GalleryFile1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileManagerAPI.PostFileManagerFiles(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileManagerAPI.PostFileManagerFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFileManagerFiles`: GalleryFile
	fmt.Fprintf(os.Stdout, "Response from `FileManagerAPI.PostFileManagerFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFileManagerFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GalleryFile1**](GalleryFile1.md) |  | 

### Return type

[**GalleryFile**](GalleryFile.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFileManagerFolders

> GalleryFolder PostFileManagerFolders(ctx).Body(body).Execute()

Add folder



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
	body := *openapiclient.NewGalleryFolder1("Name_example") // GalleryFolder1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileManagerAPI.PostFileManagerFolders(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileManagerAPI.PostFileManagerFolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFileManagerFolders`: GalleryFolder
	fmt.Fprintf(os.Stdout, "Response from `FileManagerAPI.PostFileManagerFolders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFileManagerFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GalleryFolder1**](GalleryFolder1.md) |  | 

### Return type

[**GalleryFolder**](GalleryFolder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

