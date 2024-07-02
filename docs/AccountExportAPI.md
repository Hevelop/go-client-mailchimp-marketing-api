# Hevelop\MailchimpMarketingApi\AccountExportAPI

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountExportId**](AccountExportAPI.md#GetAccountExportId) | **Get** /account-exports/{export_id} | Get account export info



## GetAccountExportId

> AccountExportsInner GetAccountExportId(ctx, exportId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get account export info



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
	exportId := "exportId_example" // string | The unique id for the account export.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountExportAPI.GetAccountExportId(context.Background(), exportId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountExportAPI.GetAccountExportId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountExportId`: AccountExportsInner
	fmt.Fprintf(os.Stdout, "Response from `AccountExportAPI.GetAccountExportId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exportId** | **string** | The unique id for the account export. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountExportIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**AccountExportsInner**](AccountExportsInner.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

