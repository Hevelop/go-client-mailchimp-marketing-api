# Hevelop\MailchimpMarketingApi\SearchMembersAPI

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSearchMembers**](SearchMembersAPI.md#GetSearchMembers) | **Get** /search-members | Search members



## GetSearchMembers

> Members GetSearchMembers(ctx).Query(query).Fields(fields).ExcludeFields(excludeFields).ListId(listId).Execute()

Search members



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
	query := "query_example" // string | The search query used to filter results. Query should be a valid email, or a string representing a contact's first or last name.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	listId := "listId_example" // string | The unique id for the list. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchMembersAPI.GetSearchMembers(context.Background()).Query(query).Fields(fields).ExcludeFields(excludeFields).ListId(listId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchMembersAPI.GetSearchMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSearchMembers`: Members
	fmt.Fprintf(os.Stdout, "Response from `SearchMembersAPI.GetSearchMembers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | The search query used to filter results. Query should be a valid email, or a string representing a contact&#39;s first or last name. | 
 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **listId** | **string** | The unique id for the list. | 

### Return type

[**Members**](Members.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

