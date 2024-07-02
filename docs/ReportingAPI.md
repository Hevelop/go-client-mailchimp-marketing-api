# Hevelop\MailchimpMarketingApi\ReportingAPI

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetReportingLandingPages**](ReportingAPI.md#GetReportingLandingPages) | **Get** /reporting/landing-pages | List landing pages reports
[**GetReportingLandingPagesId**](ReportingAPI.md#GetReportingLandingPagesId) | **Get** /reporting/landing-pages/{outreach_id} | Get landing page report
[**GetReportingSurveys**](ReportingAPI.md#GetReportingSurveys) | **Get** /reporting/surveys | List survey reports
[**GetReportingSurveysId**](ReportingAPI.md#GetReportingSurveysId) | **Get** /reporting/surveys/{survey_id} | Get survey report
[**GetReportingSurveysIdQuestions**](ReportingAPI.md#GetReportingSurveysIdQuestions) | **Get** /reporting/surveys/{survey_id}/questions | List survey question reports
[**GetReportingSurveysIdQuestionsId**](ReportingAPI.md#GetReportingSurveysIdQuestionsId) | **Get** /reporting/surveys/{survey_id}/questions/{question_id} | Get survey question report
[**GetReportingSurveysIdQuestionsIdAnswers**](ReportingAPI.md#GetReportingSurveysIdQuestionsIdAnswers) | **Get** /reporting/surveys/{survey_id}/questions/{question_id}/answers | List answers for question
[**GetReportingSurveysIdResponses**](ReportingAPI.md#GetReportingSurveysIdResponses) | **Get** /reporting/surveys/{survey_id}/responses | List survey responses
[**GetReportingSurveysIdResponsesId**](ReportingAPI.md#GetReportingSurveysIdResponsesId) | **Get** /reporting/surveys/{survey_id}/responses/{response_id} | Get survey response



## GetReportingLandingPages

> GetReportingLandingPages200Response GetReportingLandingPages(ctx).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()

List landing pages reports



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
	resp, r, err := apiClient.ReportingAPI.GetReportingLandingPages(context.Background()).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportingAPI.GetReportingLandingPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportingLandingPages`: GetReportingLandingPages200Response
	fmt.Fprintf(os.Stdout, "Response from `ReportingAPI.GetReportingLandingPages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReportingLandingPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**GetReportingLandingPages200Response**](GetReportingLandingPages200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportingLandingPagesId

> LandingPageReport GetReportingLandingPagesId(ctx, outreachId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get landing page report



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
	outreachId := "outreachId_example" // string | The outreach id.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportingAPI.GetReportingLandingPagesId(context.Background(), outreachId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportingAPI.GetReportingLandingPagesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportingLandingPagesId`: LandingPageReport
	fmt.Fprintf(os.Stdout, "Response from `ReportingAPI.GetReportingLandingPagesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**outreachId** | **string** | The outreach id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportingLandingPagesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**LandingPageReport**](LandingPageReport.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportingSurveys

> GetReportingSurveys200Response GetReportingSurveys(ctx).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()

List survey reports



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
	resp, r, err := apiClient.ReportingAPI.GetReportingSurveys(context.Background()).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportingAPI.GetReportingSurveys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportingSurveys`: GetReportingSurveys200Response
	fmt.Fprintf(os.Stdout, "Response from `ReportingAPI.GetReportingSurveys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReportingSurveysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**GetReportingSurveys200Response**](GetReportingSurveys200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportingSurveysId

> SurveyReport GetReportingSurveysId(ctx, surveyId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get survey report



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
	surveyId := "surveyId_example" // string | The ID of the survey.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportingAPI.GetReportingSurveysId(context.Background(), surveyId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportingAPI.GetReportingSurveysId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportingSurveysId`: SurveyReport
	fmt.Fprintf(os.Stdout, "Response from `ReportingAPI.GetReportingSurveysId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**surveyId** | **string** | The ID of the survey. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportingSurveysIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**SurveyReport**](SurveyReport.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportingSurveysIdQuestions

> GetReportingSurveysIdQuestions200Response GetReportingSurveysIdQuestions(ctx, surveyId).Fields(fields).ExcludeFields(excludeFields).Execute()

List survey question reports



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
	surveyId := "surveyId_example" // string | The ID of the survey.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportingAPI.GetReportingSurveysIdQuestions(context.Background(), surveyId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportingAPI.GetReportingSurveysIdQuestions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportingSurveysIdQuestions`: GetReportingSurveysIdQuestions200Response
	fmt.Fprintf(os.Stdout, "Response from `ReportingAPI.GetReportingSurveysIdQuestions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**surveyId** | **string** | The ID of the survey. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportingSurveysIdQuestionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**GetReportingSurveysIdQuestions200Response**](GetReportingSurveysIdQuestions200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportingSurveysIdQuestionsId

> SurveyQuestionReport GetReportingSurveysIdQuestionsId(ctx, surveyId, questionId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get survey question report



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
	surveyId := "surveyId_example" // string | The ID of the survey.
	questionId := "questionId_example" // string | The ID of the survey question.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportingAPI.GetReportingSurveysIdQuestionsId(context.Background(), surveyId, questionId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportingAPI.GetReportingSurveysIdQuestionsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportingSurveysIdQuestionsId`: SurveyQuestionReport
	fmt.Fprintf(os.Stdout, "Response from `ReportingAPI.GetReportingSurveysIdQuestionsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**surveyId** | **string** | The ID of the survey. | 
**questionId** | **string** | The ID of the survey question. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportingSurveysIdQuestionsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**SurveyQuestionReport**](SurveyQuestionReport.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportingSurveysIdQuestionsIdAnswers

> GetReportingSurveysIdQuestionsIdAnswers200Response GetReportingSurveysIdQuestionsIdAnswers(ctx, surveyId, questionId).Fields(fields).ExcludeFields(excludeFields).RespondentFamiliarityIs(respondentFamiliarityIs).Execute()

List answers for question



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
	surveyId := "surveyId_example" // string | The ID of the survey.
	questionId := "questionId_example" // string | The ID of the survey question.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	respondentFamiliarityIs := "respondentFamiliarityIs_example" // string | Filter survey responses by familiarity of the respondents. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportingAPI.GetReportingSurveysIdQuestionsIdAnswers(context.Background(), surveyId, questionId).Fields(fields).ExcludeFields(excludeFields).RespondentFamiliarityIs(respondentFamiliarityIs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportingAPI.GetReportingSurveysIdQuestionsIdAnswers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportingSurveysIdQuestionsIdAnswers`: GetReportingSurveysIdQuestionsIdAnswers200Response
	fmt.Fprintf(os.Stdout, "Response from `ReportingAPI.GetReportingSurveysIdQuestionsIdAnswers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**surveyId** | **string** | The ID of the survey. | 
**questionId** | **string** | The ID of the survey question. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportingSurveysIdQuestionsIdAnswersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **respondentFamiliarityIs** | **string** | Filter survey responses by familiarity of the respondents. | 

### Return type

[**GetReportingSurveysIdQuestionsIdAnswers200Response**](GetReportingSurveysIdQuestionsIdAnswers200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportingSurveysIdResponses

> GetReportingSurveysIdResponses200Response GetReportingSurveysIdResponses(ctx, surveyId).Fields(fields).ExcludeFields(excludeFields).AnsweredQuestion(answeredQuestion).ChoseAnswer(choseAnswer).RespondentFamiliarityIs(respondentFamiliarityIs).Execute()

List survey responses



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
	surveyId := "surveyId_example" // string | The ID of the survey.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	answeredQuestion := int32(56) // int32 | The ID of the question that was answered. (optional)
	choseAnswer := "choseAnswer_example" // string | The ID of the option chosen to filter responses on. (optional)
	respondentFamiliarityIs := "respondentFamiliarityIs_example" // string | Filter survey responses by familiarity of the respondents. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportingAPI.GetReportingSurveysIdResponses(context.Background(), surveyId).Fields(fields).ExcludeFields(excludeFields).AnsweredQuestion(answeredQuestion).ChoseAnswer(choseAnswer).RespondentFamiliarityIs(respondentFamiliarityIs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportingAPI.GetReportingSurveysIdResponses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportingSurveysIdResponses`: GetReportingSurveysIdResponses200Response
	fmt.Fprintf(os.Stdout, "Response from `ReportingAPI.GetReportingSurveysIdResponses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**surveyId** | **string** | The ID of the survey. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportingSurveysIdResponsesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **answeredQuestion** | **int32** | The ID of the question that was answered. | 
 **choseAnswer** | **string** | The ID of the option chosen to filter responses on. | 
 **respondentFamiliarityIs** | **string** | Filter survey responses by familiarity of the respondents. | 

### Return type

[**GetReportingSurveysIdResponses200Response**](GetReportingSurveysIdResponses200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportingSurveysIdResponsesId

> SurveyResponse GetReportingSurveysIdResponsesId(ctx, surveyId, responseId).Execute()

Get survey response



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
	surveyId := "surveyId_example" // string | The ID of the survey.
	responseId := "responseId_example" // string | The ID of the survey response.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportingAPI.GetReportingSurveysIdResponsesId(context.Background(), surveyId, responseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportingAPI.GetReportingSurveysIdResponsesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportingSurveysIdResponsesId`: SurveyResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportingAPI.GetReportingSurveysIdResponsesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**surveyId** | **string** | The ID of the survey. | 
**responseId** | **string** | The ID of the survey response. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportingSurveysIdResponsesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SurveyResponse**](SurveyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

