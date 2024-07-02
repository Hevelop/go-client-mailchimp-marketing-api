# Hevelop\MailchimpMarketingApi\ConnectedSitesAPI

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteConnectedSitesId**](ConnectedSitesAPI.md#DeleteConnectedSitesId) | **Delete** /connected-sites/{connected_site_id} | Delete connected site
[**GetConnectedSites**](ConnectedSitesAPI.md#GetConnectedSites) | **Get** /connected-sites | List connected sites
[**GetConnectedSitesId**](ConnectedSitesAPI.md#GetConnectedSitesId) | **Get** /connected-sites/{connected_site_id} | Get connected site
[**PostConnectedSites**](ConnectedSitesAPI.md#PostConnectedSites) | **Post** /connected-sites | Add connected site
[**PostConnectedSitesIdActionsVerifyScriptInstallation**](ConnectedSitesAPI.md#PostConnectedSitesIdActionsVerifyScriptInstallation) | **Post** /connected-sites/{connected_site_id}/actions/verify-script-installation | Verify connected site script



## DeleteConnectedSitesId

> DeleteConnectedSitesId(ctx, connectedSiteId).Execute()

Delete connected site



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
	connectedSiteId := "connectedSiteId_example" // string | The unique identifier for the site.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConnectedSitesAPI.DeleteConnectedSitesId(context.Background(), connectedSiteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectedSitesAPI.DeleteConnectedSitesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectedSiteId** | **string** | The unique identifier for the site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectedSitesIdRequest struct via the builder pattern


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


## GetConnectedSites

> ConnectedSites GetConnectedSites(ctx).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()

List connected sites



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
	resp, r, err := apiClient.ConnectedSitesAPI.GetConnectedSites(context.Background()).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectedSitesAPI.GetConnectedSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectedSites`: ConnectedSites
	fmt.Fprintf(os.Stdout, "Response from `ConnectedSitesAPI.GetConnectedSites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectedSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**ConnectedSites**](ConnectedSites.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectedSitesId

> ConnectedSite GetConnectedSitesId(ctx, connectedSiteId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get connected site



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
	connectedSiteId := "connectedSiteId_example" // string | The unique identifier for the site.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectedSitesAPI.GetConnectedSitesId(context.Background(), connectedSiteId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectedSitesAPI.GetConnectedSitesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectedSitesId`: ConnectedSite
	fmt.Fprintf(os.Stdout, "Response from `ConnectedSitesAPI.GetConnectedSitesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectedSiteId** | **string** | The unique identifier for the site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectedSitesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**ConnectedSite**](ConnectedSite.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostConnectedSites

> ConnectedSite PostConnectedSites(ctx).Body(body).Execute()

Add connected site



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
	body := *openapiclient.NewConnectedSite1("MC001", "example.com") // ConnectedSite1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectedSitesAPI.PostConnectedSites(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectedSitesAPI.PostConnectedSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostConnectedSites`: ConnectedSite
	fmt.Fprintf(os.Stdout, "Response from `ConnectedSitesAPI.PostConnectedSites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostConnectedSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ConnectedSite1**](ConnectedSite1.md) |  | 

### Return type

[**ConnectedSite**](ConnectedSite.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostConnectedSitesIdActionsVerifyScriptInstallation

> PostConnectedSitesIdActionsVerifyScriptInstallation(ctx, connectedSiteId).Execute()

Verify connected site script



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
	connectedSiteId := "connectedSiteId_example" // string | The unique identifier for the site.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConnectedSitesAPI.PostConnectedSitesIdActionsVerifyScriptInstallation(context.Background(), connectedSiteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectedSitesAPI.PostConnectedSitesIdActionsVerifyScriptInstallation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectedSiteId** | **string** | The unique identifier for the site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostConnectedSitesIdActionsVerifyScriptInstallationRequest struct via the builder pattern


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

