# Hevelop\MailchimpMarketingApi\SurveysAPI

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostListsIdSurveysIdActionsCreateEmail**](SurveysAPI.md#PostListsIdSurveysIdActionsCreateEmail) | **Post** /lists/{list_id}/surveys/{survey_id}/actions/create-email | Create a Survey Campaign
[**PostListsIdSurveysIdActionsPublish**](SurveysAPI.md#PostListsIdSurveysIdActionsPublish) | **Post** /lists/{list_id}/surveys/{survey_id}/actions/publish | Publish a Survey
[**PostListsIdSurveysIdActionsUnpublish**](SurveysAPI.md#PostListsIdSurveysIdActionsUnpublish) | **Post** /lists/{list_id}/surveys/{survey_id}/actions/unpublish | Unpublish a Survey



## PostListsIdSurveysIdActionsCreateEmail

> Campaign3 PostListsIdSurveysIdActionsCreateEmail(ctx, listId, surveyId).Execute()

Create a Survey Campaign



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
	listId := "listId_example" // string | The unique ID for the list.
	surveyId := "surveyId_example" // string | The ID of the survey.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SurveysAPI.PostListsIdSurveysIdActionsCreateEmail(context.Background(), listId, surveyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveysAPI.PostListsIdSurveysIdActionsCreateEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostListsIdSurveysIdActionsCreateEmail`: Campaign3
	fmt.Fprintf(os.Stdout, "Response from `SurveysAPI.PostListsIdSurveysIdActionsCreateEmail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**surveyId** | **string** | The ID of the survey. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostListsIdSurveysIdActionsCreateEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Campaign3**](Campaign3.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostListsIdSurveysIdActionsPublish

> PostListsIdSurveysIdActionsPublish(ctx, listId, surveyId).Execute()

Publish a Survey



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
	listId := "listId_example" // string | The unique ID for the list.
	surveyId := "surveyId_example" // string | The ID of the survey.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SurveysAPI.PostListsIdSurveysIdActionsPublish(context.Background(), listId, surveyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveysAPI.PostListsIdSurveysIdActionsPublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**surveyId** | **string** | The ID of the survey. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostListsIdSurveysIdActionsPublishRequest struct via the builder pattern


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


## PostListsIdSurveysIdActionsUnpublish

> PostListsIdSurveysIdActionsUnpublish(ctx, listId, surveyId).Execute()

Unpublish a Survey



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
	listId := "listId_example" // string | The unique ID for the list.
	surveyId := "surveyId_example" // string | The ID of the survey.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SurveysAPI.PostListsIdSurveysIdActionsUnpublish(context.Background(), listId, surveyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveysAPI.PostListsIdSurveysIdActionsUnpublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**surveyId** | **string** | The ID of the survey. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostListsIdSurveysIdActionsUnpublishRequest struct via the builder pattern


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

