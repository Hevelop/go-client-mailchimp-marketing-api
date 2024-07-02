# Hevelop\MailchimpMarketingApi\ReportsAPI

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetReports**](ReportsAPI.md#GetReports) | **Get** /reports | List campaign reports
[**GetReportsId**](ReportsAPI.md#GetReportsId) | **Get** /reports/{campaign_id} | Get campaign report
[**GetReportsIdAbuseReportsId**](ReportsAPI.md#GetReportsIdAbuseReportsId) | **Get** /reports/{campaign_id}/abuse-reports | List abuse reports
[**GetReportsIdAbuseReportsIdId**](ReportsAPI.md#GetReportsIdAbuseReportsIdId) | **Get** /reports/{campaign_id}/abuse-reports/{report_id} | Get abuse report
[**GetReportsIdAdvice**](ReportsAPI.md#GetReportsIdAdvice) | **Get** /reports/{campaign_id}/advice | List campaign feedback
[**GetReportsIdClickDetails**](ReportsAPI.md#GetReportsIdClickDetails) | **Get** /reports/{campaign_id}/click-details | List campaign details
[**GetReportsIdClickDetailsId**](ReportsAPI.md#GetReportsIdClickDetailsId) | **Get** /reports/{campaign_id}/click-details/{link_id} | Get campaign link details
[**GetReportsIdClickDetailsIdMembers**](ReportsAPI.md#GetReportsIdClickDetailsIdMembers) | **Get** /reports/{campaign_id}/click-details/{link_id}/members | List clicked link subscribers
[**GetReportsIdClickDetailsIdMembersId**](ReportsAPI.md#GetReportsIdClickDetailsIdMembersId) | **Get** /reports/{campaign_id}/click-details/{link_id}/members/{subscriber_hash} | Get clicked link subscriber
[**GetReportsIdDomainPerformance**](ReportsAPI.md#GetReportsIdDomainPerformance) | **Get** /reports/{campaign_id}/domain-performance | List domain performance stats
[**GetReportsIdEcommerceProductActivity**](ReportsAPI.md#GetReportsIdEcommerceProductActivity) | **Get** /reports/{campaign_id}/ecommerce-product-activity | List campaign product activity
[**GetReportsIdEepurl**](ReportsAPI.md#GetReportsIdEepurl) | **Get** /reports/{campaign_id}/eepurl | List EepURL activity
[**GetReportsIdEmailActivity**](ReportsAPI.md#GetReportsIdEmailActivity) | **Get** /reports/{campaign_id}/email-activity | List email activity
[**GetReportsIdEmailActivityId**](ReportsAPI.md#GetReportsIdEmailActivityId) | **Get** /reports/{campaign_id}/email-activity/{subscriber_hash} | Get subscriber email activity
[**GetReportsIdLocations**](ReportsAPI.md#GetReportsIdLocations) | **Get** /reports/{campaign_id}/locations | List top open activities
[**GetReportsIdOpenDetails**](ReportsAPI.md#GetReportsIdOpenDetails) | **Get** /reports/{campaign_id}/open-details | List campaign open details
[**GetReportsIdOpenDetailsIdMembersId**](ReportsAPI.md#GetReportsIdOpenDetailsIdMembersId) | **Get** /reports/{campaign_id}/open-details/{subscriber_hash} | Get opened campaign subscriber
[**GetReportsIdSentTo**](ReportsAPI.md#GetReportsIdSentTo) | **Get** /reports/{campaign_id}/sent-to | List campaign recipients
[**GetReportsIdSentToId**](ReportsAPI.md#GetReportsIdSentToId) | **Get** /reports/{campaign_id}/sent-to/{subscriber_hash} | Get campaign recipient info
[**GetReportsIdSubReportsId**](ReportsAPI.md#GetReportsIdSubReportsId) | **Get** /reports/{campaign_id}/sub-reports | List child campaign reports
[**GetReportsIdUnsubscribed**](ReportsAPI.md#GetReportsIdUnsubscribed) | **Get** /reports/{campaign_id}/unsubscribed | List unsubscribed members
[**GetReportsIdUnsubscribedId**](ReportsAPI.md#GetReportsIdUnsubscribedId) | **Get** /reports/{campaign_id}/unsubscribed/{subscriber_hash} | Get unsubscribed member



## GetReports

> CampaignReports1 GetReports(ctx).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Type_(type_).BeforeSendTime(beforeSendTime).SinceSendTime(sinceSendTime).Execute()

List campaign reports



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
	beforeSendTime := time.Now() // time.Time | Restrict the response to campaigns sent before the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. (optional)
	sinceSendTime := time.Now() // time.Time | Restrict the response to campaigns sent after the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetReports(context.Background()).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Type_(type_).BeforeSendTime(beforeSendTime).SinceSendTime(sinceSendTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReports`: CampaignReports1
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReports`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **type_** | **string** | The campaign type. | 
 **beforeSendTime** | **time.Time** | Restrict the response to campaigns sent before the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **sinceSendTime** | **time.Time** | Restrict the response to campaigns sent after the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 

### Return type

[**CampaignReports1**](CampaignReports1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportsId

> CampaignReport GetReportsId(ctx, campaignId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get campaign report



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
	resp, r, err := apiClient.ReportsAPI.GetReportsId(context.Background(), campaignId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportsId`: CampaignReport
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReportsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**CampaignReport**](CampaignReport.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportsIdAbuseReportsId

> AbuseComplaints1 GetReportsIdAbuseReportsId(ctx, campaignId).Fields(fields).ExcludeFields(excludeFields).Execute()

List abuse reports



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
	resp, r, err := apiClient.ReportsAPI.GetReportsIdAbuseReportsId(context.Background(), campaignId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportsIdAbuseReportsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportsIdAbuseReportsId`: AbuseComplaints1
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReportsIdAbuseReportsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsIdAbuseReportsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**AbuseComplaints1**](AbuseComplaints1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportsIdAbuseReportsIdId

> AbuseComplaint1 GetReportsIdAbuseReportsIdId(ctx, campaignId, reportId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get abuse report



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
	reportId := "reportId_example" // string | The id for the abuse report.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetReportsIdAbuseReportsIdId(context.Background(), campaignId, reportId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportsIdAbuseReportsIdId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportsIdAbuseReportsIdId`: AbuseComplaint1
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReportsIdAbuseReportsIdId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 
**reportId** | **string** | The id for the abuse report. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsIdAbuseReportsIdIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**AbuseComplaint1**](AbuseComplaint1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportsIdAdvice

> CampaignAdviceReport GetReportsIdAdvice(ctx, campaignId).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	resp, r, err := apiClient.ReportsAPI.GetReportsIdAdvice(context.Background(), campaignId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportsIdAdvice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportsIdAdvice`: CampaignAdviceReport
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReportsIdAdvice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsIdAdviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**CampaignAdviceReport**](CampaignAdviceReport.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportsIdClickDetails

> ClickDetailReport GetReportsIdClickDetails(ctx, campaignId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).SortField(sortField).SortDir(sortDir).Execute()

List campaign details



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
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)
	sortField := "sortField_example" // string | Returns click reports sorted by the specified field. (optional)
	sortDir := "sortDir_example" // string | Determines the order direction for sorted results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetReportsIdClickDetails(context.Background(), campaignId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).SortField(sortField).SortDir(sortDir).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportsIdClickDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportsIdClickDetails`: ClickDetailReport
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReportsIdClickDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsIdClickDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **sortField** | **string** | Returns click reports sorted by the specified field. | 
 **sortDir** | **string** | Determines the order direction for sorted results. | 

### Return type

[**ClickDetailReport**](ClickDetailReport.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportsIdClickDetailsId

> ClickDetailReport GetReportsIdClickDetailsId(ctx, campaignId, linkId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get campaign link details



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
	linkId := "linkId_example" // string | The id for the link.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetReportsIdClickDetailsId(context.Background(), campaignId, linkId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportsIdClickDetailsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportsIdClickDetailsId`: ClickDetailReport
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReportsIdClickDetailsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 
**linkId** | **string** | The id for the link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsIdClickDetailsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**ClickDetailReport**](ClickDetailReport.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportsIdClickDetailsIdMembers

> ClickDetailMembers GetReportsIdClickDetailsIdMembers(ctx, campaignId, linkId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()

List clicked link subscribers



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
	linkId := "linkId_example" // string | The id for the link.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetReportsIdClickDetailsIdMembers(context.Background(), campaignId, linkId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportsIdClickDetailsIdMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportsIdClickDetailsIdMembers`: ClickDetailMembers
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReportsIdClickDetailsIdMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 
**linkId** | **string** | The id for the link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsIdClickDetailsIdMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**ClickDetailMembers**](ClickDetailMembers.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportsIdClickDetailsIdMembersId

> ClickDetailMember GetReportsIdClickDetailsIdMembersId(ctx, campaignId, linkId, subscriberHash).Fields(fields).ExcludeFields(excludeFields).Execute()

Get clicked link subscriber



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
	linkId := "linkId_example" // string | The id for the link.
	subscriberHash := "subscriberHash_example" // string | The MD5 hash of the lowercase version of the list member's email address.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetReportsIdClickDetailsIdMembersId(context.Background(), campaignId, linkId, subscriberHash).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportsIdClickDetailsIdMembersId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportsIdClickDetailsIdMembersId`: ClickDetailMember
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReportsIdClickDetailsIdMembersId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 
**linkId** | **string** | The id for the link. | 
**subscriberHash** | **string** | The MD5 hash of the lowercase version of the list member&#39;s email address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsIdClickDetailsIdMembersIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**ClickDetailMember**](ClickDetailMember.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportsIdDomainPerformance

> DomainPerformance GetReportsIdDomainPerformance(ctx, campaignId).Fields(fields).ExcludeFields(excludeFields).Execute()

List domain performance stats



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
	resp, r, err := apiClient.ReportsAPI.GetReportsIdDomainPerformance(context.Background(), campaignId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportsIdDomainPerformance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportsIdDomainPerformance`: DomainPerformance
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReportsIdDomainPerformance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsIdDomainPerformanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**DomainPerformance**](DomainPerformance.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportsIdEcommerceProductActivity

> GetReportsIdEcommerceProductActivity200Response GetReportsIdEcommerceProductActivity(ctx, campaignId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).SortField(sortField).Execute()

List campaign product activity



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
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)
	sortField := "sortField_example" // string | Returns files sorted by the specified field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetReportsIdEcommerceProductActivity(context.Background(), campaignId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).SortField(sortField).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportsIdEcommerceProductActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportsIdEcommerceProductActivity`: GetReportsIdEcommerceProductActivity200Response
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReportsIdEcommerceProductActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsIdEcommerceProductActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **sortField** | **string** | Returns files sorted by the specified field. | 

### Return type

[**GetReportsIdEcommerceProductActivity200Response**](GetReportsIdEcommerceProductActivity200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportsIdEepurl

> EepurlActivity GetReportsIdEepurl(ctx, campaignId).Fields(fields).ExcludeFields(excludeFields).Execute()

List EepURL activity



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
	resp, r, err := apiClient.ReportsAPI.GetReportsIdEepurl(context.Background(), campaignId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportsIdEepurl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportsIdEepurl`: EepurlActivity
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReportsIdEepurl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsIdEepurlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**EepurlActivity**](EepurlActivity.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportsIdEmailActivity

> EmailActivity GetReportsIdEmailActivity(ctx, campaignId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Since(since).Execute()

List email activity



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
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)
	since := "since_example" // string | Restrict results to email activity events that occur after a specific time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetReportsIdEmailActivity(context.Background(), campaignId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Since(since).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportsIdEmailActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportsIdEmailActivity`: EmailActivity
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReportsIdEmailActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsIdEmailActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **since** | **string** | Restrict results to email activity events that occur after a specific time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 

### Return type

[**EmailActivity**](EmailActivity.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportsIdEmailActivityId

> EmailActivity GetReportsIdEmailActivityId(ctx, campaignId, subscriberHash).Fields(fields).ExcludeFields(excludeFields).Since(since).Execute()

Get subscriber email activity



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
	subscriberHash := "subscriberHash_example" // string | The MD5 hash of the lowercase version of the list member's email address.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	since := "since_example" // string | Restrict results to email activity events that occur after a specific time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetReportsIdEmailActivityId(context.Background(), campaignId, subscriberHash).Fields(fields).ExcludeFields(excludeFields).Since(since).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportsIdEmailActivityId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportsIdEmailActivityId`: EmailActivity
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReportsIdEmailActivityId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 
**subscriberHash** | **string** | The MD5 hash of the lowercase version of the list member&#39;s email address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsIdEmailActivityIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **since** | **string** | Restrict results to email activity events that occur after a specific time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 

### Return type

[**EmailActivity**](EmailActivity.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportsIdLocations

> OpenLocations GetReportsIdLocations(ctx, campaignId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()

List top open activities



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
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetReportsIdLocations(context.Background(), campaignId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportsIdLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportsIdLocations`: OpenLocations
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReportsIdLocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsIdLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**OpenLocations**](OpenLocations.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportsIdOpenDetails

> OpenDetailReport GetReportsIdOpenDetails(ctx, campaignId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Since(since).SortField(sortField).SortDir(sortDir).Execute()

List campaign open details



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
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)
	since := "2016-04-12 12:00:00" // string | Restrict results to campaign open events that occur after a specific time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. (optional)
	sortField := "sortField_example" // string | Returns open reports sorted by the specified field. (optional)
	sortDir := "sortDir_example" // string | Determines the order direction for sorted results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetReportsIdOpenDetails(context.Background(), campaignId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Since(since).SortField(sortField).SortDir(sortDir).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportsIdOpenDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportsIdOpenDetails`: OpenDetailReport
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReportsIdOpenDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsIdOpenDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **since** | **string** | Restrict results to campaign open events that occur after a specific time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **sortField** | **string** | Returns open reports sorted by the specified field. | 
 **sortDir** | **string** | Determines the order direction for sorted results. | 

### Return type

[**OpenDetailReport**](OpenDetailReport.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportsIdOpenDetailsIdMembersId

> OpenActivity GetReportsIdOpenDetailsIdMembersId(ctx, campaignId, subscriberHash).Fields(fields).ExcludeFields(excludeFields).Execute()

Get opened campaign subscriber



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
	subscriberHash := "subscriberHash_example" // string | The MD5 hash of the lowercase version of the list member's email address.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetReportsIdOpenDetailsIdMembersId(context.Background(), campaignId, subscriberHash).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportsIdOpenDetailsIdMembersId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportsIdOpenDetailsIdMembersId`: OpenActivity
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReportsIdOpenDetailsIdMembersId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 
**subscriberHash** | **string** | The MD5 hash of the lowercase version of the list member&#39;s email address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsIdOpenDetailsIdMembersIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**OpenActivity**](OpenActivity.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportsIdSentTo

> SentTo GetReportsIdSentTo(ctx, campaignId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()

List campaign recipients



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
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetReportsIdSentTo(context.Background(), campaignId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportsIdSentTo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportsIdSentTo`: SentTo
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReportsIdSentTo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsIdSentToRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**SentTo**](SentTo.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportsIdSentToId

> SentTo GetReportsIdSentToId(ctx, campaignId, subscriberHash).Fields(fields).ExcludeFields(excludeFields).Execute()

Get campaign recipient info



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
	subscriberHash := "subscriberHash_example" // string | The MD5 hash of the lowercase version of the list member's email address.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetReportsIdSentToId(context.Background(), campaignId, subscriberHash).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportsIdSentToId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportsIdSentToId`: SentTo
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReportsIdSentToId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 
**subscriberHash** | **string** | The MD5 hash of the lowercase version of the list member&#39;s email address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsIdSentToIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**SentTo**](SentTo.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportsIdSubReportsId

> CampaignSubReports GetReportsIdSubReportsId(ctx, campaignId).Fields(fields).ExcludeFields(excludeFields).Execute()

List child campaign reports



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
	resp, r, err := apiClient.ReportsAPI.GetReportsIdSubReportsId(context.Background(), campaignId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportsIdSubReportsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportsIdSubReportsId`: CampaignSubReports
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReportsIdSubReportsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsIdSubReportsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**CampaignSubReports**](CampaignSubReports.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportsIdUnsubscribed

> Unsubscribes GetReportsIdUnsubscribed(ctx, campaignId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()

List unsubscribed members



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
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetReportsIdUnsubscribed(context.Background(), campaignId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportsIdUnsubscribed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportsIdUnsubscribed`: Unsubscribes
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReportsIdUnsubscribed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsIdUnsubscribedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**Unsubscribes**](Unsubscribes.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportsIdUnsubscribedId

> Unsubscribes GetReportsIdUnsubscribedId(ctx, campaignId, subscriberHash).Fields(fields).ExcludeFields(excludeFields).Execute()

Get unsubscribed member



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
	subscriberHash := "subscriberHash_example" // string | The MD5 hash of the lowercase version of the list member's email address.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetReportsIdUnsubscribedId(context.Background(), campaignId, subscriberHash).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportsIdUnsubscribedId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportsIdUnsubscribedId`: Unsubscribes
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReportsIdUnsubscribedId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The unique id for the campaign. | 
**subscriberHash** | **string** | The MD5 hash of the lowercase version of the list member&#39;s email address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsIdUnsubscribedIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**Unsubscribes**](Unsubscribes.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

