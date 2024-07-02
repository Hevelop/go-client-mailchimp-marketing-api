# Hevelop\MailchimpMarketingApi\CampaignsAPI

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCampaignsId**](CampaignsAPI.md#DeleteCampaignsId) | **Delete** /campaigns/{campaign_id} | Delete campaign
[**DeleteCampaignsIdFeedbackId**](CampaignsAPI.md#DeleteCampaignsIdFeedbackId) | **Delete** /campaigns/{campaign_id}/feedback/{feedback_id} | Delete campaign feedback message
[**GetCampaigns**](CampaignsAPI.md#GetCampaigns) | **Get** /campaigns | List campaigns
[**GetCampaignsId**](CampaignsAPI.md#GetCampaignsId) | **Get** /campaigns/{campaign_id} | Get campaign info
[**GetCampaignsIdContent**](CampaignsAPI.md#GetCampaignsIdContent) | **Get** /campaigns/{campaign_id}/content | Get campaign content
[**GetCampaignsIdFeedback**](CampaignsAPI.md#GetCampaignsIdFeedback) | **Get** /campaigns/{campaign_id}/feedback | List campaign feedback
[**GetCampaignsIdFeedbackId**](CampaignsAPI.md#GetCampaignsIdFeedbackId) | **Get** /campaigns/{campaign_id}/feedback/{feedback_id} | Get campaign feedback message
[**GetCampaignsIdSendChecklist**](CampaignsAPI.md#GetCampaignsIdSendChecklist) | **Get** /campaigns/{campaign_id}/send-checklist | Get campaign send checklist
[**PatchCampaignsId**](CampaignsAPI.md#PatchCampaignsId) | **Patch** /campaigns/{campaign_id} | Update campaign settings
[**PatchCampaignsIdFeedbackId**](CampaignsAPI.md#PatchCampaignsIdFeedbackId) | **Patch** /campaigns/{campaign_id}/feedback/{feedback_id} | Update campaign feedback message
[**PostCampaigns**](CampaignsAPI.md#PostCampaigns) | **Post** /campaigns | Add campaign
[**PostCampaignsIdActionsCancelSend**](CampaignsAPI.md#PostCampaignsIdActionsCancelSend) | **Post** /campaigns/{campaign_id}/actions/cancel-send | Cancel campaign
[**PostCampaignsIdActionsCreateResend**](CampaignsAPI.md#PostCampaignsIdActionsCreateResend) | **Post** /campaigns/{campaign_id}/actions/create-resend | Resend campaign
[**PostCampaignsIdActionsPause**](CampaignsAPI.md#PostCampaignsIdActionsPause) | **Post** /campaigns/{campaign_id}/actions/pause | Pause rss campaign
[**PostCampaignsIdActionsReplicate**](CampaignsAPI.md#PostCampaignsIdActionsReplicate) | **Post** /campaigns/{campaign_id}/actions/replicate | Replicate campaign
[**PostCampaignsIdActionsResume**](CampaignsAPI.md#PostCampaignsIdActionsResume) | **Post** /campaigns/{campaign_id}/actions/resume | Resume rss campaign
[**PostCampaignsIdActionsSchedule**](CampaignsAPI.md#PostCampaignsIdActionsSchedule) | **Post** /campaigns/{campaign_id}/actions/schedule | Schedule campaign
[**PostCampaignsIdActionsSend**](CampaignsAPI.md#PostCampaignsIdActionsSend) | **Post** /campaigns/{campaign_id}/actions/send | Send campaign
[**PostCampaignsIdActionsTest**](CampaignsAPI.md#PostCampaignsIdActionsTest) | **Post** /campaigns/{campaign_id}/actions/test | Send test email
[**PostCampaignsIdActionsUnschedule**](CampaignsAPI.md#PostCampaignsIdActionsUnschedule) | **Post** /campaigns/{campaign_id}/actions/unschedule | Unschedule campaign
[**PostCampaignsIdFeedback**](CampaignsAPI.md#PostCampaignsIdFeedback) | **Post** /campaigns/{campaign_id}/feedback | Add campaign feedback
[**PutCampaignsIdContent**](CampaignsAPI.md#PutCampaignsIdContent) | **Put** /campaigns/{campaign_id}/content | Set campaign content



## DeleteCampaignsId

> DeleteCampaignsId(ctx, campaignId).Execute()

Delete campaign



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
	campaignId := "campaignId_example" // string | The unique id for the campaign.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CampaignsAPI.DeleteCampaignsId(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.DeleteCampaignsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCampaignsIdRequest struct via the builder pattern


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


## DeleteCampaignsIdFeedbackId

> DeleteCampaignsIdFeedbackId(ctx, campaignId, feedbackId).Execute()

Delete campaign feedback message



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
	campaignId := "campaignId_example" // string | The unique id for the campaign.
	feedbackId := "feedbackId_example" // string | The unique id for the feedback message.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CampaignsAPI.DeleteCampaignsIdFeedbackId(context.Background(), campaignId, feedbackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.DeleteCampaignsIdFeedbackId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 
**feedbackId** | **string** | The unique id for the feedback message. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCampaignsIdFeedbackIdRequest struct via the builder pattern


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


## GetCampaigns

> GetCampaigns200Response GetCampaigns(ctx).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Type_(type_).Status(status).BeforeSendTime(beforeSendTime).SinceSendTime(sinceSendTime).BeforeCreateTime(beforeCreateTime).SinceCreateTime(sinceCreateTime).ListId(listId).FolderId(folderId).MemberId(memberId).SortField(sortField).SortDir(sortDir).IncludeResendShortcutEligibility(includeResendShortcutEligibility).Execute()

List campaigns



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
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)
	type_ := "type__example" // string | The campaign type. (optional)
	status := "status_example" // string | The status of the campaign. (optional)
	beforeSendTime := time.Now() // time.Time | Restrict the response to campaigns sent before the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. (optional)
	sinceSendTime := time.Now() // time.Time | Restrict the response to campaigns sent after the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. (optional)
	beforeCreateTime := time.Now() // time.Time | Restrict the response to campaigns created before the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. (optional)
	sinceCreateTime := time.Now() // time.Time | Restrict the response to campaigns created after the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. (optional)
	listId := "listId_example" // string | The unique id for the list. (optional)
	folderId := "folderId_example" // string | The unique folder id. (optional)
	memberId := "memberId_example" // string | Retrieve campaigns sent to a particular list member. Member ID is The MD5 hash of the lowercase version of the list member’s email address. (optional)
	sortField := "sortField_example" // string | Returns files sorted by the specified field. (optional)
	sortDir := "sortDir_example" // string | Determines the order direction for sorted results. (optional)
	includeResendShortcutEligibility := true // bool | Return the `resend_shortcut_eligibility` field in the response, which tells you if the campaign is eligible for the various Campaign Resend Shortcuts offered. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.GetCampaigns(context.Background()).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Type_(type_).Status(status).BeforeSendTime(beforeSendTime).SinceSendTime(sinceSendTime).BeforeCreateTime(beforeCreateTime).SinceCreateTime(sinceCreateTime).ListId(listId).FolderId(folderId).MemberId(memberId).SortField(sortField).SortDir(sortDir).IncludeResendShortcutEligibility(includeResendShortcutEligibility).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.GetCampaigns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaigns`: GetCampaigns200Response
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.GetCampaigns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **type_** | **string** | The campaign type. | 
 **status** | **string** | The status of the campaign. | 
 **beforeSendTime** | **time.Time** | Restrict the response to campaigns sent before the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **sinceSendTime** | **time.Time** | Restrict the response to campaigns sent after the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **beforeCreateTime** | **time.Time** | Restrict the response to campaigns created before the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **sinceCreateTime** | **time.Time** | Restrict the response to campaigns created after the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **listId** | **string** | The unique id for the list. | 
 **folderId** | **string** | The unique folder id. | 
 **memberId** | **string** | Retrieve campaigns sent to a particular list member. Member ID is The MD5 hash of the lowercase version of the list member’s email address. | 
 **sortField** | **string** | Returns files sorted by the specified field. | 
 **sortDir** | **string** | Determines the order direction for sorted results. | 
 **includeResendShortcutEligibility** | **bool** | Return the &#x60;resend_shortcut_eligibility&#x60; field in the response, which tells you if the campaign is eligible for the various Campaign Resend Shortcuts offered. | 

### Return type

[**GetCampaigns200Response**](GetCampaigns200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignsId

> Campaign GetCampaignsId(ctx, campaignId).Fields(fields).ExcludeFields(excludeFields).IncludeResendShortcutEligibility(includeResendShortcutEligibility).Execute()

Get campaign info



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
	campaignId := "campaignId_example" // string | The unique id for the campaign.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	includeResendShortcutEligibility := true // bool | Return the `resend_shortcut_eligibility` field in the response, which tells you if the campaign is eligible for the various Campaign Resend Shortcuts offered. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.GetCampaignsId(context.Background(), campaignId).Fields(fields).ExcludeFields(excludeFields).IncludeResendShortcutEligibility(includeResendShortcutEligibility).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.GetCampaignsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignsId`: Campaign
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.GetCampaignsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **includeResendShortcutEligibility** | **bool** | Return the &#x60;resend_shortcut_eligibility&#x60; field in the response, which tells you if the campaign is eligible for the various Campaign Resend Shortcuts offered. | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignsIdContent

> CampaignContent GetCampaignsIdContent(ctx, campaignId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get campaign content



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
	campaignId := "campaignId_example" // string | The unique id for the campaign.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.GetCampaignsIdContent(context.Background(), campaignId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.GetCampaignsIdContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignsIdContent`: CampaignContent
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.GetCampaignsIdContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignsIdContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**CampaignContent**](CampaignContent.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignsIdFeedback

> CampaignReports GetCampaignsIdFeedback(ctx, campaignId).Fields(fields).ExcludeFields(excludeFields).Execute()

List campaign feedback



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
	campaignId := "campaignId_example" // string | The unique id for the campaign.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.GetCampaignsIdFeedback(context.Background(), campaignId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.GetCampaignsIdFeedback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignsIdFeedback`: CampaignReports
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.GetCampaignsIdFeedback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignsIdFeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**CampaignReports**](CampaignReports.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignsIdFeedbackId

> CampaignFeedback2 GetCampaignsIdFeedbackId(ctx, campaignId, feedbackId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get campaign feedback message



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
	campaignId := "campaignId_example" // string | The unique id for the campaign.
	feedbackId := "feedbackId_example" // string | The unique id for the feedback message.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.GetCampaignsIdFeedbackId(context.Background(), campaignId, feedbackId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.GetCampaignsIdFeedbackId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignsIdFeedbackId`: CampaignFeedback2
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.GetCampaignsIdFeedbackId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 
**feedbackId** | **string** | The unique id for the feedback message. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignsIdFeedbackIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**CampaignFeedback2**](CampaignFeedback2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignsIdSendChecklist

> SendChecklist GetCampaignsIdSendChecklist(ctx, campaignId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get campaign send checklist



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
	campaignId := "campaignId_example" // string | The unique id for the campaign.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.GetCampaignsIdSendChecklist(context.Background(), campaignId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.GetCampaignsIdSendChecklist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignsIdSendChecklist`: SendChecklist
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.GetCampaignsIdSendChecklist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignsIdSendChecklistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**SendChecklist**](SendChecklist.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCampaignsId

> Campaign PatchCampaignsId(ctx, campaignId).Body(body).Execute()

Update campaign settings



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
	campaignId := "campaignId_example" // string | The unique id for the campaign.
	body := *openapiclient.NewCampaign2(*openapiclient.NewCampaignSettings4("SubjectLine_example", "FromName_example", "ReplyTo_example")) // Campaign2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.PatchCampaignsId(context.Background(), campaignId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.PatchCampaignsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCampaignsId`: Campaign
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.PatchCampaignsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCampaignsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**Campaign2**](Campaign2.md) |  | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCampaignsIdFeedbackId

> CampaignFeedback2 PatchCampaignsIdFeedbackId(ctx, campaignId, feedbackId).Body(body).Execute()

Update campaign feedback message



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
	campaignId := "campaignId_example" // string | The unique id for the campaign.
	feedbackId := "feedbackId_example" // string | The unique id for the feedback message.
	body := *openapiclient.NewCampaignFeedback3() // CampaignFeedback3 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.PatchCampaignsIdFeedbackId(context.Background(), campaignId, feedbackId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.PatchCampaignsIdFeedbackId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCampaignsIdFeedbackId`: CampaignFeedback2
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.PatchCampaignsIdFeedbackId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 
**feedbackId** | **string** | The unique id for the feedback message. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCampaignsIdFeedbackIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**CampaignFeedback3**](CampaignFeedback3.md) |  | 

### Return type

[**CampaignFeedback2**](CampaignFeedback2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCampaigns

> Campaign PostCampaigns(ctx).Body(body).Execute()

Add campaign



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
	body := *openapiclient.NewCampaign1("Type_example") // Campaign1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.PostCampaigns(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.PostCampaigns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCampaigns`: Campaign
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.PostCampaigns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Campaign1**](Campaign1.md) |  | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCampaignsIdActionsCancelSend

> PostCampaignsIdActionsCancelSend(ctx, campaignId).Execute()

Cancel campaign



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
	campaignId := "campaignId_example" // string | The unique id for the campaign.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CampaignsAPI.PostCampaignsIdActionsCancelSend(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.PostCampaignsIdActionsCancelSend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCampaignsIdActionsCancelSendRequest struct via the builder pattern


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


## PostCampaignsIdActionsCreateResend

> Campaign3 PostCampaignsIdActionsCreateResend(ctx, campaignId).Body(body).Execute()

Resend campaign



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
	campaignId := "campaignId_example" // string | The unique id for the campaign.
	body := *openapiclient.NewPostCampaignsIdActionsCreateResendRequest() // PostCampaignsIdActionsCreateResendRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.PostCampaignsIdActionsCreateResend(context.Background(), campaignId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.PostCampaignsIdActionsCreateResend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCampaignsIdActionsCreateResend`: Campaign3
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.PostCampaignsIdActionsCreateResend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCampaignsIdActionsCreateResendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PostCampaignsIdActionsCreateResendRequest**](PostCampaignsIdActionsCreateResendRequest.md) |  | 

### Return type

[**Campaign3**](Campaign3.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCampaignsIdActionsPause

> PostCampaignsIdActionsPause(ctx, campaignId).Execute()

Pause rss campaign



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
	campaignId := "campaignId_example" // string | The unique id for the campaign.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CampaignsAPI.PostCampaignsIdActionsPause(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.PostCampaignsIdActionsPause``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCampaignsIdActionsPauseRequest struct via the builder pattern


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


## PostCampaignsIdActionsReplicate

> Campaign3 PostCampaignsIdActionsReplicate(ctx, campaignId).Execute()

Replicate campaign



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
	campaignId := "campaignId_example" // string | The unique id for the campaign.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.PostCampaignsIdActionsReplicate(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.PostCampaignsIdActionsReplicate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCampaignsIdActionsReplicate`: Campaign3
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.PostCampaignsIdActionsReplicate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCampaignsIdActionsReplicateRequest struct via the builder pattern


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


## PostCampaignsIdActionsResume

> PostCampaignsIdActionsResume(ctx, campaignId).Execute()

Resume rss campaign



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
	campaignId := "campaignId_example" // string | The unique id for the campaign.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CampaignsAPI.PostCampaignsIdActionsResume(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.PostCampaignsIdActionsResume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCampaignsIdActionsResumeRequest struct via the builder pattern


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


## PostCampaignsIdActionsSchedule

> PostCampaignsIdActionsSchedule(ctx, campaignId).Body(body).Execute()

Schedule campaign



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
	campaignId := "campaignId_example" // string | The unique id for the campaign.
	body := *openapiclient.NewPostCampaignsIdActionsScheduleRequest(time.Now()) // PostCampaignsIdActionsScheduleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CampaignsAPI.PostCampaignsIdActionsSchedule(context.Background(), campaignId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.PostCampaignsIdActionsSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCampaignsIdActionsScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PostCampaignsIdActionsScheduleRequest**](PostCampaignsIdActionsScheduleRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCampaignsIdActionsSend

> PostCampaignsIdActionsSend(ctx, campaignId).Execute()

Send campaign



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
	campaignId := "campaignId_example" // string | The unique id for the campaign.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CampaignsAPI.PostCampaignsIdActionsSend(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.PostCampaignsIdActionsSend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCampaignsIdActionsSendRequest struct via the builder pattern


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


## PostCampaignsIdActionsTest

> PostCampaignsIdActionsTest(ctx, campaignId).Body(body).Execute()

Send test email



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
	campaignId := "campaignId_example" // string | The unique id for the campaign.
	body := *openapiclient.NewPostCampaignsIdActionsTestRequest([]string{"TestEmails_example"}, "SendType_example") // PostCampaignsIdActionsTestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CampaignsAPI.PostCampaignsIdActionsTest(context.Background(), campaignId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.PostCampaignsIdActionsTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCampaignsIdActionsTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PostCampaignsIdActionsTestRequest**](PostCampaignsIdActionsTestRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCampaignsIdActionsUnschedule

> PostCampaignsIdActionsUnschedule(ctx, campaignId).Execute()

Unschedule campaign



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
	campaignId := "campaignId_example" // string | The unique id for the campaign.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CampaignsAPI.PostCampaignsIdActionsUnschedule(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.PostCampaignsIdActionsUnschedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCampaignsIdActionsUnscheduleRequest struct via the builder pattern


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


## PostCampaignsIdFeedback

> CampaignFeedback2 PostCampaignsIdFeedback(ctx, campaignId).Body(body).Execute()

Add campaign feedback



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
	campaignId := "campaignId_example" // string | The unique id for the campaign.
	body := *openapiclient.NewCampaignFeedback1("Message_example") // CampaignFeedback1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.PostCampaignsIdFeedback(context.Background(), campaignId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.PostCampaignsIdFeedback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCampaignsIdFeedback`: CampaignFeedback2
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.PostCampaignsIdFeedback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCampaignsIdFeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CampaignFeedback1**](CampaignFeedback1.md) |  | 

### Return type

[**CampaignFeedback2**](CampaignFeedback2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCampaignsIdContent

> CampaignContent PutCampaignsIdContent(ctx, campaignId).Body(body).Execute()

Set campaign content



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
	campaignId := "campaignId_example" // string | The unique id for the campaign.
	body := *openapiclient.NewCampaignContent1() // CampaignContent1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.PutCampaignsIdContent(context.Background(), campaignId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.PutCampaignsIdContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCampaignsIdContent`: CampaignContent
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.PutCampaignsIdContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCampaignsIdContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CampaignContent1**](CampaignContent1.md) |  | 

### Return type

[**CampaignContent**](CampaignContent.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

