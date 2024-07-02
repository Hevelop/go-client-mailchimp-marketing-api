# Hevelop\MailchimpMarketingApi\AutomationsAPI

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveAutomations**](AutomationsAPI.md#ArchiveAutomations) | **Post** /automations/{workflow_id}/actions/archive | Archive automation
[**DeleteAutomationsIdEmailsId**](AutomationsAPI.md#DeleteAutomationsIdEmailsId) | **Delete** /automations/{workflow_id}/emails/{workflow_email_id} | Delete workflow email
[**GetAutomations**](AutomationsAPI.md#GetAutomations) | **Get** /automations | List automations
[**GetAutomationsId**](AutomationsAPI.md#GetAutomationsId) | **Get** /automations/{workflow_id} | Get automation info
[**GetAutomationsIdEmails**](AutomationsAPI.md#GetAutomationsIdEmails) | **Get** /automations/{workflow_id}/emails | List automated emails
[**GetAutomationsIdEmailsId**](AutomationsAPI.md#GetAutomationsIdEmailsId) | **Get** /automations/{workflow_id}/emails/{workflow_email_id} | Get workflow email info
[**GetAutomationsIdEmailsIdQueue**](AutomationsAPI.md#GetAutomationsIdEmailsIdQueue) | **Get** /automations/{workflow_id}/emails/{workflow_email_id}/queue | List automated email subscribers
[**GetAutomationsIdEmailsIdQueueId**](AutomationsAPI.md#GetAutomationsIdEmailsIdQueueId) | **Get** /automations/{workflow_id}/emails/{workflow_email_id}/queue/{subscriber_hash} | Get automated email subscriber
[**GetAutomationsIdRemovedSubscribers**](AutomationsAPI.md#GetAutomationsIdRemovedSubscribers) | **Get** /automations/{workflow_id}/removed-subscribers | List subscribers removed from workflow
[**GetAutomationsIdRemovedSubscribersId**](AutomationsAPI.md#GetAutomationsIdRemovedSubscribersId) | **Get** /automations/{workflow_id}/removed-subscribers/{subscriber_hash} | Get subscriber removed from workflow
[**PatchAutomationEmailWorkflowId**](AutomationsAPI.md#PatchAutomationEmailWorkflowId) | **Patch** /automations/{workflow_id}/emails/{workflow_email_id} | Update workflow email
[**PostAutomations**](AutomationsAPI.md#PostAutomations) | **Post** /automations | Add automation
[**PostAutomationsIdActionsPauseAllEmails**](AutomationsAPI.md#PostAutomationsIdActionsPauseAllEmails) | **Post** /automations/{workflow_id}/actions/pause-all-emails | Pause automation emails
[**PostAutomationsIdActionsStartAllEmails**](AutomationsAPI.md#PostAutomationsIdActionsStartAllEmails) | **Post** /automations/{workflow_id}/actions/start-all-emails | Start automation emails
[**PostAutomationsIdEmailsIdActionsPause**](AutomationsAPI.md#PostAutomationsIdEmailsIdActionsPause) | **Post** /automations/{workflow_id}/emails/{workflow_email_id}/actions/pause | Pause automated email
[**PostAutomationsIdEmailsIdActionsStart**](AutomationsAPI.md#PostAutomationsIdEmailsIdActionsStart) | **Post** /automations/{workflow_id}/emails/{workflow_email_id}/actions/start | Start automated email
[**PostAutomationsIdEmailsIdQueue**](AutomationsAPI.md#PostAutomationsIdEmailsIdQueue) | **Post** /automations/{workflow_id}/emails/{workflow_email_id}/queue | Add subscriber to workflow email
[**PostAutomationsIdRemovedSubscribers**](AutomationsAPI.md#PostAutomationsIdRemovedSubscribers) | **Post** /automations/{workflow_id}/removed-subscribers | Remove subscriber from workflow



## ArchiveAutomations

> ArchiveAutomations(ctx, workflowId).Execute()

Archive automation



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
	workflowId := "workflowId_example" // string | The unique id for the Automation workflow.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutomationsAPI.ArchiveAutomations(context.Background(), workflowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.ArchiveAutomations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** | The unique id for the Automation workflow. | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveAutomationsRequest struct via the builder pattern


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


## DeleteAutomationsIdEmailsId

> DeleteAutomationsIdEmailsId(ctx, workflowId, workflowEmailId).Execute()

Delete workflow email



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
	workflowId := "workflowId_example" // string | The unique id for the Automation workflow.
	workflowEmailId := "workflowEmailId_example" // string | The unique id for the Automation workflow email.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutomationsAPI.DeleteAutomationsIdEmailsId(context.Background(), workflowId, workflowEmailId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.DeleteAutomationsIdEmailsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** | The unique id for the Automation workflow. | 
**workflowEmailId** | **string** | The unique id for the Automation workflow email. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutomationsIdEmailsIdRequest struct via the builder pattern


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


## GetAutomations

> GetAutomations200Response GetAutomations(ctx).Count(count).Offset(offset).Fields(fields).ExcludeFields(excludeFields).BeforeCreateTime(beforeCreateTime).SinceCreateTime(sinceCreateTime).BeforeStartTime(beforeStartTime).SinceStartTime(sinceStartTime).Status(status).Execute()

List automations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/Hevelop/go-client-mailchimp-marketing-api"
)

func main() {
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	beforeCreateTime := time.Now() // time.Time | Restrict the response to automations created before this time. Uses the ISO 8601 time format: 2015-10-21T15:41:36+00:00. (optional)
	sinceCreateTime := time.Now() // time.Time | Restrict the response to automations created after this time. Uses the ISO 8601 time format: 2015-10-21T15:41:36+00:00. (optional)
	beforeStartTime := time.Now() // time.Time | Restrict the response to automations started before this time. Uses the ISO 8601 time format: 2015-10-21T15:41:36+00:00. (optional)
	sinceStartTime := time.Now() // time.Time | Restrict the response to automations started after this time. Uses the ISO 8601 time format: 2015-10-21T15:41:36+00:00. (optional)
	status := "status_example" // string | Restrict the results to automations with the specified status. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationsAPI.GetAutomations(context.Background()).Count(count).Offset(offset).Fields(fields).ExcludeFields(excludeFields).BeforeCreateTime(beforeCreateTime).SinceCreateTime(sinceCreateTime).BeforeStartTime(beforeStartTime).SinceStartTime(sinceStartTime).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.GetAutomations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomations`: GetAutomations200Response
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.GetAutomations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **beforeCreateTime** | **time.Time** | Restrict the response to automations created before this time. Uses the ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **sinceCreateTime** | **time.Time** | Restrict the response to automations created after this time. Uses the ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **beforeStartTime** | **time.Time** | Restrict the response to automations started before this time. Uses the ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **sinceStartTime** | **time.Time** | Restrict the response to automations started after this time. Uses the ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **status** | **string** | Restrict the results to automations with the specified status. | 

### Return type

[**GetAutomations200Response**](GetAutomations200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutomationsId

> AutomationWorkflow GetAutomationsId(ctx, workflowId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get automation info



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
	workflowId := "workflowId_example" // string | The unique id for the Automation workflow.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationsAPI.GetAutomationsId(context.Background(), workflowId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.GetAutomationsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomationsId`: AutomationWorkflow
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.GetAutomationsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** | The unique id for the Automation workflow. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**AutomationWorkflow**](AutomationWorkflow.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutomationsIdEmails

> AutomationEmails GetAutomationsIdEmails(ctx, workflowId).Execute()

List automated emails



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
	workflowId := "workflowId_example" // string | The unique id for the Automation workflow.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationsAPI.GetAutomationsIdEmails(context.Background(), workflowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.GetAutomationsIdEmails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomationsIdEmails`: AutomationEmails
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.GetAutomationsIdEmails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** | The unique id for the Automation workflow. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationsIdEmailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutomationEmails**](AutomationEmails.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutomationsIdEmailsId

> AutomationWorkflowEmail GetAutomationsIdEmailsId(ctx, workflowId, workflowEmailId).Execute()

Get workflow email info



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
	workflowId := "workflowId_example" // string | The unique id for the Automation workflow.
	workflowEmailId := "workflowEmailId_example" // string | The unique id for the Automation workflow email.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationsAPI.GetAutomationsIdEmailsId(context.Background(), workflowId, workflowEmailId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.GetAutomationsIdEmailsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomationsIdEmailsId`: AutomationWorkflowEmail
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.GetAutomationsIdEmailsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** | The unique id for the Automation workflow. | 
**workflowEmailId** | **string** | The unique id for the Automation workflow email. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationsIdEmailsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AutomationWorkflowEmail**](AutomationWorkflowEmail.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutomationsIdEmailsIdQueue

> GetAutomationsIdEmailsIdQueue200Response GetAutomationsIdEmailsIdQueue(ctx, workflowId, workflowEmailId).Execute()

List automated email subscribers



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
	workflowId := "workflowId_example" // string | The unique id for the Automation workflow.
	workflowEmailId := "workflowEmailId_example" // string | The unique id for the Automation workflow email.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationsAPI.GetAutomationsIdEmailsIdQueue(context.Background(), workflowId, workflowEmailId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.GetAutomationsIdEmailsIdQueue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomationsIdEmailsIdQueue`: GetAutomationsIdEmailsIdQueue200Response
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.GetAutomationsIdEmailsIdQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** | The unique id for the Automation workflow. | 
**workflowEmailId** | **string** | The unique id for the Automation workflow email. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationsIdEmailsIdQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetAutomationsIdEmailsIdQueue200Response**](GetAutomationsIdEmailsIdQueue200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutomationsIdEmailsIdQueueId

> SubscriberInAutomationQueue2 GetAutomationsIdEmailsIdQueueId(ctx, workflowId, workflowEmailId, subscriberHash).Execute()

Get automated email subscriber



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
	workflowId := "workflowId_example" // string | The unique id for the Automation workflow.
	workflowEmailId := "workflowEmailId_example" // string | The unique id for the Automation workflow email.
	subscriberHash := "subscriberHash_example" // string | The MD5 hash of the lowercase version of the list member's email address.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationsAPI.GetAutomationsIdEmailsIdQueueId(context.Background(), workflowId, workflowEmailId, subscriberHash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.GetAutomationsIdEmailsIdQueueId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomationsIdEmailsIdQueueId`: SubscriberInAutomationQueue2
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.GetAutomationsIdEmailsIdQueueId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** | The unique id for the Automation workflow. | 
**workflowEmailId** | **string** | The unique id for the Automation workflow email. | 
**subscriberHash** | **string** | The MD5 hash of the lowercase version of the list member&#39;s email address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationsIdEmailsIdQueueIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SubscriberInAutomationQueue2**](SubscriberInAutomationQueue2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutomationsIdRemovedSubscribers

> RemovedSubscribers GetAutomationsIdRemovedSubscribers(ctx, workflowId).Execute()

List subscribers removed from workflow



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
	workflowId := "workflowId_example" // string | The unique id for the Automation workflow.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationsAPI.GetAutomationsIdRemovedSubscribers(context.Background(), workflowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.GetAutomationsIdRemovedSubscribers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomationsIdRemovedSubscribers`: RemovedSubscribers
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.GetAutomationsIdRemovedSubscribers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** | The unique id for the Automation workflow. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationsIdRemovedSubscribersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RemovedSubscribers**](RemovedSubscribers.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutomationsIdRemovedSubscribersId

> SubscriberRemovedFromAutomationWorkflow GetAutomationsIdRemovedSubscribersId(ctx, workflowId, subscriberHash).Execute()

Get subscriber removed from workflow



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
	workflowId := "workflowId_example" // string | The unique id for the Automation workflow.
	subscriberHash := "subscriberHash_example" // string | The MD5 hash of the lowercase version of the list member's email address.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationsAPI.GetAutomationsIdRemovedSubscribersId(context.Background(), workflowId, subscriberHash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.GetAutomationsIdRemovedSubscribersId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomationsIdRemovedSubscribersId`: SubscriberRemovedFromAutomationWorkflow
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.GetAutomationsIdRemovedSubscribersId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** | The unique id for the Automation workflow. | 
**subscriberHash** | **string** | The MD5 hash of the lowercase version of the list member&#39;s email address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationsIdRemovedSubscribersIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SubscriberRemovedFromAutomationWorkflow**](SubscriberRemovedFromAutomationWorkflow.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAutomationEmailWorkflowId

> AutomationWorkflowEmail PatchAutomationEmailWorkflowId(ctx, workflowId, workflowEmailId).Body(body).Execute()

Update workflow email



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
	workflowId := "workflowId_example" // string | The unique id for the Automation workflow.
	workflowEmailId := "workflowEmailId_example" // string | The unique id for the Automation workflow email.
	body := *openapiclient.NewUpdateInformationAboutASpecificWorkflowEmail() // UpdateInformationAboutASpecificWorkflowEmail | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationsAPI.PatchAutomationEmailWorkflowId(context.Background(), workflowId, workflowEmailId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.PatchAutomationEmailWorkflowId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAutomationEmailWorkflowId`: AutomationWorkflowEmail
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.PatchAutomationEmailWorkflowId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** | The unique id for the Automation workflow. | 
**workflowEmailId** | **string** | The unique id for the Automation workflow email. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAutomationEmailWorkflowIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UpdateInformationAboutASpecificWorkflowEmail**](UpdateInformationAboutASpecificWorkflowEmail.md) |  | 

### Return type

[**AutomationWorkflowEmail**](AutomationWorkflowEmail.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAutomations

> AutomationWorkflow PostAutomations(ctx).Body(body).Execute()

Add automation



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
	body := *openapiclient.NewAutomationWorkflow1(*openapiclient.NewList1(), *openapiclient.NewAutomationTrigger1("WorkflowType_example")) // AutomationWorkflow1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationsAPI.PostAutomations(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.PostAutomations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAutomations`: AutomationWorkflow
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.PostAutomations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAutomationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AutomationWorkflow1**](AutomationWorkflow1.md) |  | 

### Return type

[**AutomationWorkflow**](AutomationWorkflow.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAutomationsIdActionsPauseAllEmails

> PostAutomationsIdActionsPauseAllEmails(ctx, workflowId).Execute()

Pause automation emails



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
	workflowId := "workflowId_example" // string | The unique id for the Automation workflow.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutomationsAPI.PostAutomationsIdActionsPauseAllEmails(context.Background(), workflowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.PostAutomationsIdActionsPauseAllEmails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** | The unique id for the Automation workflow. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAutomationsIdActionsPauseAllEmailsRequest struct via the builder pattern


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


## PostAutomationsIdActionsStartAllEmails

> PostAutomationsIdActionsStartAllEmails(ctx, workflowId).Execute()

Start automation emails



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
	workflowId := "workflowId_example" // string | The unique id for the Automation workflow.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutomationsAPI.PostAutomationsIdActionsStartAllEmails(context.Background(), workflowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.PostAutomationsIdActionsStartAllEmails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** | The unique id for the Automation workflow. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAutomationsIdActionsStartAllEmailsRequest struct via the builder pattern


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


## PostAutomationsIdEmailsIdActionsPause

> PostAutomationsIdEmailsIdActionsPause(ctx, workflowId, workflowEmailId).Execute()

Pause automated email



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
	workflowId := "workflowId_example" // string | The unique id for the Automation workflow.
	workflowEmailId := "workflowEmailId_example" // string | The unique id for the Automation workflow email.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutomationsAPI.PostAutomationsIdEmailsIdActionsPause(context.Background(), workflowId, workflowEmailId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.PostAutomationsIdEmailsIdActionsPause``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** | The unique id for the Automation workflow. | 
**workflowEmailId** | **string** | The unique id for the Automation workflow email. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAutomationsIdEmailsIdActionsPauseRequest struct via the builder pattern


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


## PostAutomationsIdEmailsIdActionsStart

> PostAutomationsIdEmailsIdActionsStart(ctx, workflowId, workflowEmailId).Execute()

Start automated email



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
	workflowId := "workflowId_example" // string | The unique id for the Automation workflow.
	workflowEmailId := "workflowEmailId_example" // string | The unique id for the Automation workflow email.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutomationsAPI.PostAutomationsIdEmailsIdActionsStart(context.Background(), workflowId, workflowEmailId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.PostAutomationsIdEmailsIdActionsStart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** | The unique id for the Automation workflow. | 
**workflowEmailId** | **string** | The unique id for the Automation workflow email. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAutomationsIdEmailsIdActionsStartRequest struct via the builder pattern


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


## PostAutomationsIdEmailsIdQueue

> SubscriberInAutomationQueue2 PostAutomationsIdEmailsIdQueue(ctx, workflowId, workflowEmailId).Body(body).Execute()

Add subscriber to workflow email



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
	workflowId := "workflowId_example" // string | The unique id for the Automation workflow.
	workflowEmailId := "workflowEmailId_example" // string | The unique id for the Automation workflow email.
	body := *openapiclient.NewSubscriberInAutomationQueue1("EmailAddress_example") // SubscriberInAutomationQueue1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationsAPI.PostAutomationsIdEmailsIdQueue(context.Background(), workflowId, workflowEmailId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.PostAutomationsIdEmailsIdQueue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAutomationsIdEmailsIdQueue`: SubscriberInAutomationQueue2
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.PostAutomationsIdEmailsIdQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** | The unique id for the Automation workflow. | 
**workflowEmailId** | **string** | The unique id for the Automation workflow email. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAutomationsIdEmailsIdQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**SubscriberInAutomationQueue1**](SubscriberInAutomationQueue1.md) |  | 

### Return type

[**SubscriberInAutomationQueue2**](SubscriberInAutomationQueue2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAutomationsIdRemovedSubscribers

> SubscriberRemovedFromAutomationWorkflow PostAutomationsIdRemovedSubscribers(ctx, workflowId).Body(body).Execute()

Remove subscriber from workflow



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
	workflowId := "workflowId_example" // string | The unique id for the Automation workflow.
	body := *openapiclient.NewSubscriberInAutomationQueue1("EmailAddress_example") // SubscriberInAutomationQueue1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationsAPI.PostAutomationsIdRemovedSubscribers(context.Background(), workflowId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationsAPI.PostAutomationsIdRemovedSubscribers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAutomationsIdRemovedSubscribers`: SubscriberRemovedFromAutomationWorkflow
	fmt.Fprintf(os.Stdout, "Response from `AutomationsAPI.PostAutomationsIdRemovedSubscribers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** | The unique id for the Automation workflow. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAutomationsIdRemovedSubscribersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SubscriberInAutomationQueue1**](SubscriberInAutomationQueue1.md) |  | 

### Return type

[**SubscriberRemovedFromAutomationWorkflow**](SubscriberRemovedFromAutomationWorkflow.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

