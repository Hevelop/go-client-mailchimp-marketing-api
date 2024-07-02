# Hevelop\MailchimpMarketingApi\VerifiedDomainsAPI

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVerifiedDomain**](VerifiedDomainsAPI.md#CreateVerifiedDomain) | **Post** /verified-domains | Add domain to account
[**DeleteVerifiedDomain**](VerifiedDomainsAPI.md#DeleteVerifiedDomain) | **Delete** /verified-domains/{domain_name} | Delete domain
[**GetVerifiedDomain**](VerifiedDomainsAPI.md#GetVerifiedDomain) | **Get** /verified-domains/{domain_name} | Get domain info
[**GetVerifiedDomains**](VerifiedDomainsAPI.md#GetVerifiedDomains) | **Get** /verified-domains | List sending domains
[**VerifyDomain**](VerifiedDomainsAPI.md#VerifyDomain) | **Post** /verified-domains/{domain_name}/actions/verify | Verify domain



## CreateVerifiedDomain

> VerifiedDomains CreateVerifiedDomain(ctx).Body(body).Execute()

Add domain to account



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
	body := *openapiclient.NewVerifiedDomains2("VerificationEmail_example") // VerifiedDomains2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerifiedDomainsAPI.CreateVerifiedDomain(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerifiedDomainsAPI.CreateVerifiedDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVerifiedDomain`: VerifiedDomains
	fmt.Fprintf(os.Stdout, "Response from `VerifiedDomainsAPI.CreateVerifiedDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVerifiedDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**VerifiedDomains2**](VerifiedDomains2.md) |  | 

### Return type

[**VerifiedDomains**](VerifiedDomains.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVerifiedDomain

> DeleteVerifiedDomain(ctx, domainName).Execute()

Delete domain



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
	domainName := "domainName_example" // string | The domain name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VerifiedDomainsAPI.DeleteVerifiedDomain(context.Background(), domainName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerifiedDomainsAPI.DeleteVerifiedDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainName** | **string** | The domain name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVerifiedDomainRequest struct via the builder pattern


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


## GetVerifiedDomain

> VerifiedDomains GetVerifiedDomain(ctx, domainName).Execute()

Get domain info



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
	domainName := "domainName_example" // string | The domain name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerifiedDomainsAPI.GetVerifiedDomain(context.Background(), domainName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerifiedDomainsAPI.GetVerifiedDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVerifiedDomain`: VerifiedDomains
	fmt.Fprintf(os.Stdout, "Response from `VerifiedDomainsAPI.GetVerifiedDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainName** | **string** | The domain name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVerifiedDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VerifiedDomains**](VerifiedDomains.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVerifiedDomains

> VerifiedDomains1 GetVerifiedDomains(ctx).Execute()

List sending domains



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerifiedDomainsAPI.GetVerifiedDomains(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerifiedDomainsAPI.GetVerifiedDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVerifiedDomains`: VerifiedDomains1
	fmt.Fprintf(os.Stdout, "Response from `VerifiedDomainsAPI.GetVerifiedDomains`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVerifiedDomainsRequest struct via the builder pattern


### Return type

[**VerifiedDomains1**](VerifiedDomains1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyDomain

> VerifiedDomains VerifyDomain(ctx, domainName).Body(body).Execute()

Verify domain



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
	domainName := "domainName_example" // string | The domain name.
	body := *openapiclient.NewVerifyADomainForSending("Code_example") // VerifyADomainForSending | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerifiedDomainsAPI.VerifyDomain(context.Background(), domainName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerifiedDomainsAPI.VerifyDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyDomain`: VerifiedDomains
	fmt.Fprintf(os.Stdout, "Response from `VerifiedDomainsAPI.VerifyDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainName** | **string** | The domain name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**VerifyADomainForSending**](VerifyADomainForSending.md) |  | 

### Return type

[**VerifiedDomains**](VerifiedDomains.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

