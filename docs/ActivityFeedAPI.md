# Hevelop\MailchimpMarketingApi\ActivityFeedAPI

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetActivityFeedChimpChatter**](ActivityFeedAPI.md#GetActivityFeedChimpChatter) | **Get** /activity-feed/chimp-chatter | Get latest chimp chatter



## GetActivityFeedChimpChatter

> GetActivityFeedChimpChatter200Response GetActivityFeedChimpChatter(ctx).Count(count).Offset(offset).Execute()

Get latest chimp chatter



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
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivityFeedAPI.GetActivityFeedChimpChatter(context.Background()).Count(count).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityFeedAPI.GetActivityFeedChimpChatter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActivityFeedChimpChatter`: GetActivityFeedChimpChatter200Response
	fmt.Fprintf(os.Stdout, "Response from `ActivityFeedAPI.GetActivityFeedChimpChatter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActivityFeedChimpChatterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**GetActivityFeedChimpChatter200Response**](GetActivityFeedChimpChatter200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

