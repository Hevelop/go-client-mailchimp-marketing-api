# Hevelop\MailchimpMarketingApi\CustomerJourneysAPI

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCustomerJourneysJourneysIdStepsIdActionsTrigger**](CustomerJourneysAPI.md#PostCustomerJourneysJourneysIdStepsIdActionsTrigger) | **Post** /customer-journeys/journeys/{journey_id}/steps/{step_id}/actions/trigger | Customer Journeys API trigger for a contact



## PostCustomerJourneysJourneysIdStepsIdActionsTrigger

> map[string]interface{} PostCustomerJourneysJourneysIdStepsIdActionsTrigger(ctx, journeyId, stepId).Body(body).Execute()

Customer Journeys API trigger for a contact



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
	journeyId := int32(56) // int32 | The id for the Journey.
	stepId := int32(56) // int32 | The id for the Step.
	body := *openapiclient.NewSubscriberInCustomerJourneySAudience("EmailAddress_example") // SubscriberInCustomerJourneySAudience | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerJourneysAPI.PostCustomerJourneysJourneysIdStepsIdActionsTrigger(context.Background(), journeyId, stepId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerJourneysAPI.PostCustomerJourneysJourneysIdStepsIdActionsTrigger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCustomerJourneysJourneysIdStepsIdActionsTrigger`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CustomerJourneysAPI.PostCustomerJourneysJourneysIdStepsIdActionsTrigger`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journeyId** | **int32** | The id for the Journey. | 
**stepId** | **int32** | The id for the Step. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCustomerJourneysJourneysIdStepsIdActionsTriggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**SubscriberInCustomerJourneySAudience**](SubscriberInCustomerJourneySAudience.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

