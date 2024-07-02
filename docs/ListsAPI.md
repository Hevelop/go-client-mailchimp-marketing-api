# Hevelop\MailchimpMarketingApi\ListsAPI

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteListsId**](ListsAPI.md#DeleteListsId) | **Delete** /lists/{list_id} | Delete list
[**DeleteListsIdInterestCategoriesId**](ListsAPI.md#DeleteListsIdInterestCategoriesId) | **Delete** /lists/{list_id}/interest-categories/{interest_category_id} | Delete interest category
[**DeleteListsIdInterestCategoriesIdInterestsId**](ListsAPI.md#DeleteListsIdInterestCategoriesIdInterestsId) | **Delete** /lists/{list_id}/interest-categories/{interest_category_id}/interests/{interest_id} | Delete interest in category
[**DeleteListsIdMembersId**](ListsAPI.md#DeleteListsIdMembersId) | **Delete** /lists/{list_id}/members/{subscriber_hash} | Archive list member
[**DeleteListsIdMembersIdNotesId**](ListsAPI.md#DeleteListsIdMembersIdNotesId) | **Delete** /lists/{list_id}/members/{subscriber_hash}/notes/{note_id} | Delete note
[**DeleteListsIdMergeFieldsId**](ListsAPI.md#DeleteListsIdMergeFieldsId) | **Delete** /lists/{list_id}/merge-fields/{merge_id} | Delete merge field
[**DeleteListsIdSegmentsId**](ListsAPI.md#DeleteListsIdSegmentsId) | **Delete** /lists/{list_id}/segments/{segment_id} | Delete segment
[**DeleteListsIdSegmentsIdMembersId**](ListsAPI.md#DeleteListsIdSegmentsIdMembersId) | **Delete** /lists/{list_id}/segments/{segment_id}/members/{subscriber_hash} | Remove list member from segment
[**DeleteListsIdWebhooksId**](ListsAPI.md#DeleteListsIdWebhooksId) | **Delete** /lists/{list_id}/webhooks/{webhook_id} | Delete webhook
[**GetListMemberTags**](ListsAPI.md#GetListMemberTags) | **Get** /lists/{list_id}/members/{subscriber_hash}/tags | List member tags
[**GetLists**](ListsAPI.md#GetLists) | **Get** /lists | Get lists info
[**GetListsId**](ListsAPI.md#GetListsId) | **Get** /lists/{list_id} | Get list info
[**GetListsIdAbuseReports**](ListsAPI.md#GetListsIdAbuseReports) | **Get** /lists/{list_id}/abuse-reports | List abuse reports
[**GetListsIdAbuseReportsId**](ListsAPI.md#GetListsIdAbuseReportsId) | **Get** /lists/{list_id}/abuse-reports/{report_id} | Get abuse report
[**GetListsIdActivity**](ListsAPI.md#GetListsIdActivity) | **Get** /lists/{list_id}/activity | List recent activity
[**GetListsIdClients**](ListsAPI.md#GetListsIdClients) | **Get** /lists/{list_id}/clients | List top email clients
[**GetListsIdGrowthHistory**](ListsAPI.md#GetListsIdGrowthHistory) | **Get** /lists/{list_id}/growth-history | List growth history data
[**GetListsIdGrowthHistoryId**](ListsAPI.md#GetListsIdGrowthHistoryId) | **Get** /lists/{list_id}/growth-history/{month} | Get growth history by month
[**GetListsIdInterestCategories**](ListsAPI.md#GetListsIdInterestCategories) | **Get** /lists/{list_id}/interest-categories | List interest categories
[**GetListsIdInterestCategoriesId**](ListsAPI.md#GetListsIdInterestCategoriesId) | **Get** /lists/{list_id}/interest-categories/{interest_category_id} | Get interest category info
[**GetListsIdInterestCategoriesIdInterests**](ListsAPI.md#GetListsIdInterestCategoriesIdInterests) | **Get** /lists/{list_id}/interest-categories/{interest_category_id}/interests | List interests in category
[**GetListsIdInterestCategoriesIdInterestsId**](ListsAPI.md#GetListsIdInterestCategoriesIdInterestsId) | **Get** /lists/{list_id}/interest-categories/{interest_category_id}/interests/{interest_id} | Get interest in category
[**GetListsIdLocations**](ListsAPI.md#GetListsIdLocations) | **Get** /lists/{list_id}/locations | List locations
[**GetListsIdMembers**](ListsAPI.md#GetListsIdMembers) | **Get** /lists/{list_id}/members | List members info
[**GetListsIdMembersId**](ListsAPI.md#GetListsIdMembersId) | **Get** /lists/{list_id}/members/{subscriber_hash} | Get member info
[**GetListsIdMembersIdActivity**](ListsAPI.md#GetListsIdMembersIdActivity) | **Get** /lists/{list_id}/members/{subscriber_hash}/activity | View recent activity 50
[**GetListsIdMembersIdActivityFeed**](ListsAPI.md#GetListsIdMembersIdActivityFeed) | **Get** /lists/{list_id}/members/{subscriber_hash}/activity-feed | View recent activity
[**GetListsIdMembersIdEvents**](ListsAPI.md#GetListsIdMembersIdEvents) | **Get** /lists/{list_id}/members/{subscriber_hash}/events | List member events
[**GetListsIdMembersIdGoals**](ListsAPI.md#GetListsIdMembersIdGoals) | **Get** /lists/{list_id}/members/{subscriber_hash}/goals | List member goal events
[**GetListsIdMembersIdNotes**](ListsAPI.md#GetListsIdMembersIdNotes) | **Get** /lists/{list_id}/members/{subscriber_hash}/notes | List recent member notes
[**GetListsIdMembersIdNotesId**](ListsAPI.md#GetListsIdMembersIdNotesId) | **Get** /lists/{list_id}/members/{subscriber_hash}/notes/{note_id} | Get member note
[**GetListsIdMergeFields**](ListsAPI.md#GetListsIdMergeFields) | **Get** /lists/{list_id}/merge-fields | List merge fields
[**GetListsIdMergeFieldsId**](ListsAPI.md#GetListsIdMergeFieldsId) | **Get** /lists/{list_id}/merge-fields/{merge_id} | Get merge field
[**GetListsIdSegmentsId**](ListsAPI.md#GetListsIdSegmentsId) | **Get** /lists/{list_id}/segments/{segment_id} | Get segment info
[**GetListsIdSegmentsIdMembers**](ListsAPI.md#GetListsIdSegmentsIdMembers) | **Get** /lists/{list_id}/segments/{segment_id}/members | List members in segment
[**GetListsIdSignupForms**](ListsAPI.md#GetListsIdSignupForms) | **Get** /lists/{list_id}/signup-forms | List signup forms
[**GetListsIdSurveys**](ListsAPI.md#GetListsIdSurveys) | **Get** /lists/{list_id}/surveys | Get information about all surveys for a list
[**GetListsIdSurveysId**](ListsAPI.md#GetListsIdSurveysId) | **Get** /lists/{list_id}/surveys/{survey_id} | Get survey
[**GetListsIdWebhooks**](ListsAPI.md#GetListsIdWebhooks) | **Get** /lists/{list_id}/webhooks | List webhooks
[**GetListsIdWebhooksId**](ListsAPI.md#GetListsIdWebhooksId) | **Get** /lists/{list_id}/webhooks/{webhook_id} | Get webhook info
[**PatchListsId**](ListsAPI.md#PatchListsId) | **Patch** /lists/{list_id} | Update lists
[**PatchListsIdInterestCategoriesId**](ListsAPI.md#PatchListsIdInterestCategoriesId) | **Patch** /lists/{list_id}/interest-categories/{interest_category_id} | Update interest category
[**PatchListsIdInterestCategoriesIdInterestsId**](ListsAPI.md#PatchListsIdInterestCategoriesIdInterestsId) | **Patch** /lists/{list_id}/interest-categories/{interest_category_id}/interests/{interest_id} | Update interest in category
[**PatchListsIdMembersId**](ListsAPI.md#PatchListsIdMembersId) | **Patch** /lists/{list_id}/members/{subscriber_hash} | Update list member
[**PatchListsIdMembersIdNotesId**](ListsAPI.md#PatchListsIdMembersIdNotesId) | **Patch** /lists/{list_id}/members/{subscriber_hash}/notes/{note_id} | Update note
[**PatchListsIdMergeFieldsId**](ListsAPI.md#PatchListsIdMergeFieldsId) | **Patch** /lists/{list_id}/merge-fields/{merge_id} | Update merge field
[**PatchListsIdSegmentsId**](ListsAPI.md#PatchListsIdSegmentsId) | **Patch** /lists/{list_id}/segments/{segment_id} | Update segment
[**PatchListsIdWebhooksId**](ListsAPI.md#PatchListsIdWebhooksId) | **Patch** /lists/{list_id}/webhooks/{webhook_id} | Update webhook
[**PostListMemberEvents**](ListsAPI.md#PostListMemberEvents) | **Post** /lists/{list_id}/members/{subscriber_hash}/events | Add event
[**PostListMemberTags**](ListsAPI.md#PostListMemberTags) | **Post** /lists/{list_id}/members/{subscriber_hash}/tags | Add or remove member tags
[**PostLists**](ListsAPI.md#PostLists) | **Post** /lists | Add list
[**PostListsId**](ListsAPI.md#PostListsId) | **Post** /lists/{list_id} | Batch subscribe or unsubscribe
[**PostListsIdInterestCategories**](ListsAPI.md#PostListsIdInterestCategories) | **Post** /lists/{list_id}/interest-categories | Add interest category
[**PostListsIdInterestCategoriesIdInterests**](ListsAPI.md#PostListsIdInterestCategoriesIdInterests) | **Post** /lists/{list_id}/interest-categories/{interest_category_id}/interests | Add interest in category
[**PostListsIdMembers**](ListsAPI.md#PostListsIdMembers) | **Post** /lists/{list_id}/members | Add member to list
[**PostListsIdMembersHashActionsDeletePermanent**](ListsAPI.md#PostListsIdMembersHashActionsDeletePermanent) | **Post** /lists/{list_id}/members/{subscriber_hash}/actions/delete-permanent | Delete list member
[**PostListsIdMembersIdNotes**](ListsAPI.md#PostListsIdMembersIdNotes) | **Post** /lists/{list_id}/members/{subscriber_hash}/notes | Add member note
[**PostListsIdMergeFields**](ListsAPI.md#PostListsIdMergeFields) | **Post** /lists/{list_id}/merge-fields | Add merge field
[**PostListsIdSegments**](ListsAPI.md#PostListsIdSegments) | **Post** /lists/{list_id}/segments | Add segment
[**PostListsIdSegmentsId**](ListsAPI.md#PostListsIdSegmentsId) | **Post** /lists/{list_id}/segments/{segment_id} | Batch add or remove members
[**PostListsIdSegmentsIdMembers**](ListsAPI.md#PostListsIdSegmentsIdMembers) | **Post** /lists/{list_id}/segments/{segment_id}/members | Add member to segment
[**PostListsIdSignupForms**](ListsAPI.md#PostListsIdSignupForms) | **Post** /lists/{list_id}/signup-forms | Customize signup form
[**PostListsIdWebhooks**](ListsAPI.md#PostListsIdWebhooks) | **Post** /lists/{list_id}/webhooks | Add webhook
[**PreviewASegment**](ListsAPI.md#PreviewASegment) | **Get** /lists/{list_id}/segments | List segments
[**PutListsIdMembersId**](ListsAPI.md#PutListsIdMembersId) | **Put** /lists/{list_id}/members/{subscriber_hash} | Add or update list member
[**SearchTagsByName**](ListsAPI.md#SearchTagsByName) | **Get** /lists/{list_id}/tag-search | Search for tags on a list by name.



## DeleteListsId

> DeleteListsId(ctx, listId).Execute()

Delete list



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListsAPI.DeleteListsId(context.Background(), listId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.DeleteListsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteListsIdRequest struct via the builder pattern


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


## DeleteListsIdInterestCategoriesId

> DeleteListsIdInterestCategoriesId(ctx, listId, interestCategoryId).Execute()

Delete interest category



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
	interestCategoryId := "interestCategoryId_example" // string | The unique ID for the interest category.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListsAPI.DeleteListsIdInterestCategoriesId(context.Background(), listId, interestCategoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.DeleteListsIdInterestCategoriesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**interestCategoryId** | **string** | The unique ID for the interest category. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteListsIdInterestCategoriesIdRequest struct via the builder pattern


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


## DeleteListsIdInterestCategoriesIdInterestsId

> DeleteListsIdInterestCategoriesIdInterestsId(ctx, listId, interestCategoryId, interestId).Execute()

Delete interest in category



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
	interestCategoryId := "interestCategoryId_example" // string | The unique ID for the interest category.
	interestId := "interestId_example" // string | The specific interest or 'group name'.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListsAPI.DeleteListsIdInterestCategoriesIdInterestsId(context.Background(), listId, interestCategoryId, interestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.DeleteListsIdInterestCategoriesIdInterestsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**interestCategoryId** | **string** | The unique ID for the interest category. | 
**interestId** | **string** | The specific interest or &#39;group name&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteListsIdInterestCategoriesIdInterestsIdRequest struct via the builder pattern


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


## DeleteListsIdMembersId

> DeleteListsIdMembersId(ctx, listId, subscriberHash).Execute()

Archive list member



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
	subscriberHash := "subscriberHash_example" // string | The MD5 hash of the lowercase version of the list member's email address. This endpoint also accepts a list member's email address or contact_id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListsAPI.DeleteListsIdMembersId(context.Background(), listId, subscriberHash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.DeleteListsIdMembersId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**subscriberHash** | **string** | The MD5 hash of the lowercase version of the list member&#39;s email address. This endpoint also accepts a list member&#39;s email address or contact_id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteListsIdMembersIdRequest struct via the builder pattern


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


## DeleteListsIdMembersIdNotesId

> DeleteListsIdMembersIdNotesId(ctx, listId, subscriberHash, noteId).Execute()

Delete note



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
	subscriberHash := "subscriberHash_example" // string | The MD5 hash of the lowercase version of the list member's email address. This endpoint also accepts a list member's email address or contact_id.
	noteId := "noteId_example" // string | The id for the note.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListsAPI.DeleteListsIdMembersIdNotesId(context.Background(), listId, subscriberHash, noteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.DeleteListsIdMembersIdNotesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**subscriberHash** | **string** | The MD5 hash of the lowercase version of the list member&#39;s email address. This endpoint also accepts a list member&#39;s email address or contact_id. | 
**noteId** | **string** | The id for the note. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteListsIdMembersIdNotesIdRequest struct via the builder pattern


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


## DeleteListsIdMergeFieldsId

> DeleteListsIdMergeFieldsId(ctx, listId, mergeId).Execute()

Delete merge field



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
	mergeId := "mergeId_example" // string | The id for the merge field.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListsAPI.DeleteListsIdMergeFieldsId(context.Background(), listId, mergeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.DeleteListsIdMergeFieldsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**mergeId** | **string** | The id for the merge field. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteListsIdMergeFieldsIdRequest struct via the builder pattern


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


## DeleteListsIdSegmentsId

> DeleteListsIdSegmentsId(ctx, listId, segmentId).Execute()

Delete segment



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
	segmentId := "segmentId_example" // string | The unique id for the segment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListsAPI.DeleteListsIdSegmentsId(context.Background(), listId, segmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.DeleteListsIdSegmentsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**segmentId** | **string** | The unique id for the segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteListsIdSegmentsIdRequest struct via the builder pattern


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


## DeleteListsIdSegmentsIdMembersId

> DeleteListsIdSegmentsIdMembersId(ctx, listId, segmentId, subscriberHash).Execute()

Remove list member from segment



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
	segmentId := "segmentId_example" // string | The unique id for the segment.
	subscriberHash := "subscriberHash_example" // string | The MD5 hash of the lowercase version of the list member's email address.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListsAPI.DeleteListsIdSegmentsIdMembersId(context.Background(), listId, segmentId, subscriberHash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.DeleteListsIdSegmentsIdMembersId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**segmentId** | **string** | The unique id for the segment. | 
**subscriberHash** | **string** | The MD5 hash of the lowercase version of the list member&#39;s email address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteListsIdSegmentsIdMembersIdRequest struct via the builder pattern


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


## DeleteListsIdWebhooksId

> DeleteListsIdWebhooksId(ctx, listId, webhookId).Execute()

Delete webhook



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
	webhookId := "webhookId_example" // string | The webhook's id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListsAPI.DeleteListsIdWebhooksId(context.Background(), listId, webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.DeleteListsIdWebhooksId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**webhookId** | **string** | The webhook&#39;s id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteListsIdWebhooksIdRequest struct via the builder pattern


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


## GetListMemberTags

> CollectionOfTags GetListMemberTags(ctx, listId, subscriberHash).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()

List member tags



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
	subscriberHash := "subscriberHash_example" // string | The MD5 hash of the lowercase version of the list member's email address. This endpoint also accepts a list member's email address or contact_id.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.GetListMemberTags(context.Background(), listId, subscriberHash).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.GetListMemberTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListMemberTags`: CollectionOfTags
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.GetListMemberTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**subscriberHash** | **string** | The MD5 hash of the lowercase version of the list member&#39;s email address. This endpoint also accepts a list member&#39;s email address or contact_id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListMemberTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**CollectionOfTags**](CollectionOfTags.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLists

> SubscriberLists GetLists(ctx).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).BeforeDateCreated(beforeDateCreated).SinceDateCreated(sinceDateCreated).BeforeCampaignLastSent(beforeCampaignLastSent).SinceCampaignLastSent(sinceCampaignLastSent).Email(email).SortField(sortField).SortDir(sortDir).HasEcommerceStore(hasEcommerceStore).IncludeTotalContacts(includeTotalContacts).Execute()

Get lists info



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
	beforeDateCreated := "beforeDateCreated_example" // string | Restrict response to lists created before the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. (optional)
	sinceDateCreated := "sinceDateCreated_example" // string | Restrict results to lists created after the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. (optional)
	beforeCampaignLastSent := "beforeCampaignLastSent_example" // string | Restrict results to lists created before the last campaign send date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. (optional)
	sinceCampaignLastSent := "sinceCampaignLastSent_example" // string | Restrict results to lists created after the last campaign send date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. (optional)
	email := "email_example" // string | Restrict results to lists that include a specific subscriber's email address. (optional)
	sortField := "sortField_example" // string | Returns files sorted by the specified field. (optional)
	sortDir := "sortDir_example" // string | Determines the order direction for sorted results. (optional)
	hasEcommerceStore := true // bool | Restrict results to lists that contain an active, connected, undeleted ecommerce store. (optional)
	includeTotalContacts := true // bool | Return the total_contacts field in the stats response, which contains an approximate count of all contacts in any state. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.GetLists(context.Background()).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).BeforeDateCreated(beforeDateCreated).SinceDateCreated(sinceDateCreated).BeforeCampaignLastSent(beforeCampaignLastSent).SinceCampaignLastSent(sinceCampaignLastSent).Email(email).SortField(sortField).SortDir(sortDir).HasEcommerceStore(hasEcommerceStore).IncludeTotalContacts(includeTotalContacts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.GetLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLists`: SubscriberLists
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.GetLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **beforeDateCreated** | **string** | Restrict response to lists created before the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **sinceDateCreated** | **string** | Restrict results to lists created after the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **beforeCampaignLastSent** | **string** | Restrict results to lists created before the last campaign send date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **sinceCampaignLastSent** | **string** | Restrict results to lists created after the last campaign send date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **email** | **string** | Restrict results to lists that include a specific subscriber&#39;s email address. | 
 **sortField** | **string** | Returns files sorted by the specified field. | 
 **sortDir** | **string** | Determines the order direction for sorted results. | 
 **hasEcommerceStore** | **bool** | Restrict results to lists that contain an active, connected, undeleted ecommerce store. | 
 **includeTotalContacts** | **bool** | Return the total_contacts field in the stats response, which contains an approximate count of all contacts in any state. | 

### Return type

[**SubscriberLists**](SubscriberLists.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListsId

> SubscriberList GetListsId(ctx, listId).Fields(fields).ExcludeFields(excludeFields).IncludeTotalContacts(includeTotalContacts).Execute()

Get list info



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
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	includeTotalContacts := true // bool | Return the total_contacts field in the stats response, which contains an approximate count of all contacts in any state. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.GetListsId(context.Background(), listId).Fields(fields).ExcludeFields(excludeFields).IncludeTotalContacts(includeTotalContacts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.GetListsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListsId`: SubscriberList
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.GetListsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **includeTotalContacts** | **bool** | Return the total_contacts field in the stats response, which contains an approximate count of all contacts in any state. | 

### Return type

[**SubscriberList**](SubscriberList.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListsIdAbuseReports

> AbuseComplaints GetListsIdAbuseReports(ctx, listId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()

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
	listId := "listId_example" // string | The unique ID for the list.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.GetListsIdAbuseReports(context.Background(), listId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.GetListsIdAbuseReports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListsIdAbuseReports`: AbuseComplaints
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.GetListsIdAbuseReports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListsIdAbuseReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**AbuseComplaints**](AbuseComplaints.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListsIdAbuseReportsId

> AbuseComplaint GetListsIdAbuseReportsId(ctx, listId, reportId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()

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
	listId := "listId_example" // string | The unique ID for the list.
	reportId := "reportId_example" // string | The id for the abuse report.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.GetListsIdAbuseReportsId(context.Background(), listId, reportId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.GetListsIdAbuseReportsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListsIdAbuseReportsId`: AbuseComplaint
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.GetListsIdAbuseReportsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**reportId** | **string** | The id for the abuse report. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListsIdAbuseReportsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**AbuseComplaint**](AbuseComplaint.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListsIdActivity

> ListActivity GetListsIdActivity(ctx, listId).Count(count).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()

List recent activity



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
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.GetListsIdActivity(context.Background(), listId).Count(count).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.GetListsIdActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListsIdActivity`: ListActivity
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.GetListsIdActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListsIdActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**ListActivity**](ListActivity.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListsIdClients

> EmailClients GetListsIdClients(ctx, listId).Fields(fields).ExcludeFields(excludeFields).Execute()

List top email clients



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
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.GetListsIdClients(context.Background(), listId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.GetListsIdClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListsIdClients`: EmailClients
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.GetListsIdClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListsIdClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**EmailClients**](EmailClients.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListsIdGrowthHistory

> GrowthHistory GetListsIdGrowthHistory(ctx, listId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).SortField(sortField).SortDir(sortDir).Execute()

List growth history data



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
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)
	sortField := "sortField_example" // string | Returns files sorted by the specified field. (optional)
	sortDir := "sortDir_example" // string | Determines the order direction for sorted results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.GetListsIdGrowthHistory(context.Background(), listId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).SortField(sortField).SortDir(sortDir).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.GetListsIdGrowthHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListsIdGrowthHistory`: GrowthHistory
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.GetListsIdGrowthHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListsIdGrowthHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **sortField** | **string** | Returns files sorted by the specified field. | 
 **sortDir** | **string** | Determines the order direction for sorted results. | 

### Return type

[**GrowthHistory**](GrowthHistory.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListsIdGrowthHistoryId

> GrowthHistory GetListsIdGrowthHistoryId(ctx, listId, month).Fields(fields).ExcludeFields(excludeFields).Execute()

Get growth history by month



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
	month := "month_example" // string | A specific month of list growth history.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.GetListsIdGrowthHistoryId(context.Background(), listId, month).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.GetListsIdGrowthHistoryId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListsIdGrowthHistoryId`: GrowthHistory
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.GetListsIdGrowthHistoryId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**month** | **string** | A specific month of list growth history. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListsIdGrowthHistoryIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**GrowthHistory**](GrowthHistory.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListsIdInterestCategories

> InterestGroupings GetListsIdInterestCategories(ctx, listId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Type_(type_).Execute()

List interest categories



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
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)
	type_ := "type__example" // string | Restrict results a type of interest group (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.GetListsIdInterestCategories(context.Background(), listId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.GetListsIdInterestCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListsIdInterestCategories`: InterestGroupings
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.GetListsIdInterestCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListsIdInterestCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **type_** | **string** | Restrict results a type of interest group | 

### Return type

[**InterestGroupings**](InterestGroupings.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListsIdInterestCategoriesId

> InterestCategory GetListsIdInterestCategoriesId(ctx, listId, interestCategoryId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get interest category info



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
	interestCategoryId := "interestCategoryId_example" // string | The unique ID for the interest category.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.GetListsIdInterestCategoriesId(context.Background(), listId, interestCategoryId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.GetListsIdInterestCategoriesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListsIdInterestCategoriesId`: InterestCategory
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.GetListsIdInterestCategoriesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**interestCategoryId** | **string** | The unique ID for the interest category. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListsIdInterestCategoriesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**InterestCategory**](InterestCategory.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListsIdInterestCategoriesIdInterests

> Interests GetListsIdInterestCategoriesIdInterests(ctx, listId, interestCategoryId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()

List interests in category



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
	interestCategoryId := "interestCategoryId_example" // string | The unique ID for the interest category.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.GetListsIdInterestCategoriesIdInterests(context.Background(), listId, interestCategoryId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.GetListsIdInterestCategoriesIdInterests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListsIdInterestCategoriesIdInterests`: Interests
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.GetListsIdInterestCategoriesIdInterests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**interestCategoryId** | **string** | The unique ID for the interest category. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListsIdInterestCategoriesIdInterestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**Interests**](Interests.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListsIdInterestCategoriesIdInterestsId

> Interest GetListsIdInterestCategoriesIdInterestsId(ctx, listId, interestCategoryId, interestId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get interest in category



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
	interestCategoryId := "interestCategoryId_example" // string | The unique ID for the interest category.
	interestId := "interestId_example" // string | The specific interest or 'group name'.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.GetListsIdInterestCategoriesIdInterestsId(context.Background(), listId, interestCategoryId, interestId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.GetListsIdInterestCategoriesIdInterestsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListsIdInterestCategoriesIdInterestsId`: Interest
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.GetListsIdInterestCategoriesIdInterestsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**interestCategoryId** | **string** | The unique ID for the interest category. | 
**interestId** | **string** | The specific interest or &#39;group name&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListsIdInterestCategoriesIdInterestsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**Interest**](Interest.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListsIdLocations

> ListLocations GetListsIdLocations(ctx, listId).Fields(fields).ExcludeFields(excludeFields).Execute()

List locations



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
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.GetListsIdLocations(context.Background(), listId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.GetListsIdLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListsIdLocations`: ListLocations
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.GetListsIdLocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListsIdLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**ListLocations**](ListLocations.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListsIdMembers

> ListMembers2 GetListsIdMembers(ctx, listId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).EmailType(emailType).Status(status).SinceTimestampOpt(sinceTimestampOpt).BeforeTimestampOpt(beforeTimestampOpt).SinceLastChanged(sinceLastChanged).BeforeLastChanged(beforeLastChanged).UniqueEmailId(uniqueEmailId).VipOnly(vipOnly).InterestCategoryId(interestCategoryId).InterestIds(interestIds).InterestMatch(interestMatch).SortField(sortField).SortDir(sortDir).SinceLastCampaign(sinceLastCampaign).UnsubscribedSince(unsubscribedSince).Execute()

List members info



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
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)
	emailType := "emailType_example" // string | The email type. (optional)
	status := "status_example" // string | The subscriber's status. (optional)
	sinceTimestampOpt := "sinceTimestampOpt_example" // string | Restrict results to subscribers who opted-in after the set timeframe. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. (optional)
	beforeTimestampOpt := "beforeTimestampOpt_example" // string | Restrict results to subscribers who opted-in before the set timeframe. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. (optional)
	sinceLastChanged := "sinceLastChanged_example" // string | Restrict results to subscribers whose information changed after the set timeframe. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. (optional)
	beforeLastChanged := "beforeLastChanged_example" // string | Restrict results to subscribers whose information changed before the set timeframe. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. (optional)
	uniqueEmailId := "uniqueEmailId_example" // string | A unique identifier for the email address across all Mailchimp lists. (optional)
	vipOnly := true // bool | A filter to return only the list's VIP members. Passing `true` will restrict results to VIP list members, passing `false` will return all list members. (optional)
	interestCategoryId := "interestCategoryId_example" // string | The unique id for the interest category. (optional)
	interestIds := "interestIds_example" // string | Used to filter list members by interests. Must be accompanied by interest_category_id and interest_match. The value must be a comma separated list of interest ids present for any supplied interest categories. (optional)
	interestMatch := "interestMatch_example" // string | Used to filter list members by interests. Must be accompanied by interest_category_id and interest_ids. \"any\" will match a member with any of the interest supplied, \"all\" will only match members with every interest supplied, and \"none\" will match members without any of the interest supplied. (optional)
	sortField := "sortField_example" // string | Returns files sorted by the specified field. (optional)
	sortDir := "sortDir_example" // string | Determines the order direction for sorted results. (optional)
	sinceLastCampaign := true // bool | Filter subscribers by those subscribed/unsubscribed/pending/cleaned since last email campaign send. Member status is required to use this filter. (optional)
	unsubscribedSince := "unsubscribedSince_example" // string | Filter subscribers by those unsubscribed since a specific date. Using any status other than unsubscribed with this filter will result in an error. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.GetListsIdMembers(context.Background(), listId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).EmailType(emailType).Status(status).SinceTimestampOpt(sinceTimestampOpt).BeforeTimestampOpt(beforeTimestampOpt).SinceLastChanged(sinceLastChanged).BeforeLastChanged(beforeLastChanged).UniqueEmailId(uniqueEmailId).VipOnly(vipOnly).InterestCategoryId(interestCategoryId).InterestIds(interestIds).InterestMatch(interestMatch).SortField(sortField).SortDir(sortDir).SinceLastCampaign(sinceLastCampaign).UnsubscribedSince(unsubscribedSince).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.GetListsIdMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListsIdMembers`: ListMembers2
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.GetListsIdMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListsIdMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **emailType** | **string** | The email type. | 
 **status** | **string** | The subscriber&#39;s status. | 
 **sinceTimestampOpt** | **string** | Restrict results to subscribers who opted-in after the set timeframe. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **beforeTimestampOpt** | **string** | Restrict results to subscribers who opted-in before the set timeframe. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **sinceLastChanged** | **string** | Restrict results to subscribers whose information changed after the set timeframe. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **beforeLastChanged** | **string** | Restrict results to subscribers whose information changed before the set timeframe. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **uniqueEmailId** | **string** | A unique identifier for the email address across all Mailchimp lists. | 
 **vipOnly** | **bool** | A filter to return only the list&#39;s VIP members. Passing &#x60;true&#x60; will restrict results to VIP list members, passing &#x60;false&#x60; will return all list members. | 
 **interestCategoryId** | **string** | The unique id for the interest category. | 
 **interestIds** | **string** | Used to filter list members by interests. Must be accompanied by interest_category_id and interest_match. The value must be a comma separated list of interest ids present for any supplied interest categories. | 
 **interestMatch** | **string** | Used to filter list members by interests. Must be accompanied by interest_category_id and interest_ids. \&quot;any\&quot; will match a member with any of the interest supplied, \&quot;all\&quot; will only match members with every interest supplied, and \&quot;none\&quot; will match members without any of the interest supplied. | 
 **sortField** | **string** | Returns files sorted by the specified field. | 
 **sortDir** | **string** | Determines the order direction for sorted results. | 
 **sinceLastCampaign** | **bool** | Filter subscribers by those subscribed/unsubscribed/pending/cleaned since last email campaign send. Member status is required to use this filter. | 
 **unsubscribedSince** | **string** | Filter subscribers by those unsubscribed since a specific date. Using any status other than unsubscribed with this filter will result in an error. | 

### Return type

[**ListMembers2**](ListMembers2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListsIdMembersId

> ListMembers2 GetListsIdMembersId(ctx, listId, subscriberHash).Fields(fields).ExcludeFields(excludeFields).Execute()

Get member info



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
	subscriberHash := "subscriberHash_example" // string | The MD5 hash of the lowercase version of the list member's email address. This endpoint also accepts a list member's email address or contact_id.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.GetListsIdMembersId(context.Background(), listId, subscriberHash).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.GetListsIdMembersId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListsIdMembersId`: ListMembers2
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.GetListsIdMembersId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**subscriberHash** | **string** | The MD5 hash of the lowercase version of the list member&#39;s email address. This endpoint also accepts a list member&#39;s email address or contact_id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListsIdMembersIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**ListMembers2**](ListMembers2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListsIdMembersIdActivity

> MemberActivityEvents GetListsIdMembersIdActivity(ctx, listId, subscriberHash).Fields(fields).ExcludeFields(excludeFields).Action(action).Execute()

View recent activity 50



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
	subscriberHash := "subscriberHash_example" // string | The MD5 hash of the lowercase version of the list member's email address. This endpoint also accepts a list member's email address or contact_id.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	action := []string{"Inner_example"} // []string | A comma seperated list of actions to return. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.GetListsIdMembersIdActivity(context.Background(), listId, subscriberHash).Fields(fields).ExcludeFields(excludeFields).Action(action).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.GetListsIdMembersIdActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListsIdMembersIdActivity`: MemberActivityEvents
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.GetListsIdMembersIdActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**subscriberHash** | **string** | The MD5 hash of the lowercase version of the list member&#39;s email address. This endpoint also accepts a list member&#39;s email address or contact_id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListsIdMembersIdActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **action** | **[]string** | A comma seperated list of actions to return. | 

### Return type

[**MemberActivityEvents**](MemberActivityEvents.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListsIdMembersIdActivityFeed

> MemberActivityEvents1 GetListsIdMembersIdActivityFeed(ctx, listId, subscriberHash).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).ActivityFilters(activityFilters).Execute()

View recent activity



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
	subscriberHash := "subscriberHash_example" // string | The MD5 hash of the lowercase version of the list member's email address. This endpoint also accepts a list member's email address or contact_id.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)
	activityFilters := []string{"Inner_example"} // []string | A comma-separated list of activity filters that correspond to a set of activity types, e.g \"?activity_filters=open,bounce,click\". (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.GetListsIdMembersIdActivityFeed(context.Background(), listId, subscriberHash).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).ActivityFilters(activityFilters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.GetListsIdMembersIdActivityFeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListsIdMembersIdActivityFeed`: MemberActivityEvents1
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.GetListsIdMembersIdActivityFeed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**subscriberHash** | **string** | The MD5 hash of the lowercase version of the list member&#39;s email address. This endpoint also accepts a list member&#39;s email address or contact_id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListsIdMembersIdActivityFeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **activityFilters** | **[]string** | A comma-separated list of activity filters that correspond to a set of activity types, e.g \&quot;?activity_filters&#x3D;open,bounce,click\&quot;. | 

### Return type

[**MemberActivityEvents1**](MemberActivityEvents1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListsIdMembersIdEvents

> CollectionOfEvents GetListsIdMembersIdEvents(ctx, listId, subscriberHash).Count(count).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()

List member events



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
	subscriberHash := "subscriberHash_example" // string | The MD5 hash of the lowercase version of the list member's email address. This endpoint also accepts a list member's email address or contact_id.
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.GetListsIdMembersIdEvents(context.Background(), listId, subscriberHash).Count(count).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.GetListsIdMembersIdEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListsIdMembersIdEvents`: CollectionOfEvents
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.GetListsIdMembersIdEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**subscriberHash** | **string** | The MD5 hash of the lowercase version of the list member&#39;s email address. This endpoint also accepts a list member&#39;s email address or contact_id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListsIdMembersIdEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**CollectionOfEvents**](CollectionOfEvents.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListsIdMembersIdGoals

> CollectionOfMemberActivityEvents GetListsIdMembersIdGoals(ctx, listId, subscriberHash).Fields(fields).ExcludeFields(excludeFields).Execute()

List member goal events



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
	subscriberHash := "subscriberHash_example" // string | The MD5 hash of the lowercase version of the list member's email address. This endpoint also accepts a list member's email address or contact_id.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.GetListsIdMembersIdGoals(context.Background(), listId, subscriberHash).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.GetListsIdMembersIdGoals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListsIdMembersIdGoals`: CollectionOfMemberActivityEvents
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.GetListsIdMembersIdGoals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**subscriberHash** | **string** | The MD5 hash of the lowercase version of the list member&#39;s email address. This endpoint also accepts a list member&#39;s email address or contact_id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListsIdMembersIdGoalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**CollectionOfMemberActivityEvents**](CollectionOfMemberActivityEvents.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListsIdMembersIdNotes

> CollectionOfNotes GetListsIdMembersIdNotes(ctx, listId, subscriberHash).SortField(sortField).SortDir(sortDir).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()

List recent member notes



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
	subscriberHash := "subscriberHash_example" // string | The MD5 hash of the lowercase version of the list member's email address.
	sortField := "sortField_example" // string | Returns notes sorted by the specified field. (optional)
	sortDir := "sortDir_example" // string | Determines the order direction for sorted results. (optional)
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.GetListsIdMembersIdNotes(context.Background(), listId, subscriberHash).SortField(sortField).SortDir(sortDir).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.GetListsIdMembersIdNotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListsIdMembersIdNotes`: CollectionOfNotes
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.GetListsIdMembersIdNotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**subscriberHash** | **string** | The MD5 hash of the lowercase version of the list member&#39;s email address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListsIdMembersIdNotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sortField** | **string** | Returns notes sorted by the specified field. | 
 **sortDir** | **string** | Determines the order direction for sorted results. | 
 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**CollectionOfNotes**](CollectionOfNotes.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListsIdMembersIdNotesId

> MemberNotes GetListsIdMembersIdNotesId(ctx, listId, subscriberHash, noteId).Fields(fields).ExcludeFields(excludeFields).Execute()

Get member note



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
	subscriberHash := "subscriberHash_example" // string | The MD5 hash of the lowercase version of the list member's email address. This endpoint also accepts a list member's email address or contact_id.
	noteId := "noteId_example" // string | The id for the note.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.GetListsIdMembersIdNotesId(context.Background(), listId, subscriberHash, noteId).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.GetListsIdMembersIdNotesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListsIdMembersIdNotesId`: MemberNotes
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.GetListsIdMembersIdNotesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**subscriberHash** | **string** | The MD5 hash of the lowercase version of the list member&#39;s email address. This endpoint also accepts a list member&#39;s email address or contact_id. | 
**noteId** | **string** | The id for the note. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListsIdMembersIdNotesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**MemberNotes**](MemberNotes.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListsIdMergeFields

> CollectionOfMergeFields GetListsIdMergeFields(ctx, listId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Type_(type_).Required(required).Execute()

List merge fields



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
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)
	type_ := "type__example" // string | The merge field type. (optional)
	required := true // bool | Whether it's a required merge field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.GetListsIdMergeFields(context.Background(), listId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Type_(type_).Required(required).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.GetListsIdMergeFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListsIdMergeFields`: CollectionOfMergeFields
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.GetListsIdMergeFields`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListsIdMergeFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **type_** | **string** | The merge field type. | 
 **required** | **bool** | Whether it&#39;s a required merge field. | 

### Return type

[**CollectionOfMergeFields**](CollectionOfMergeFields.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListsIdMergeFieldsId

> MergeField GetListsIdMergeFieldsId(ctx, listId, mergeId).ExcludeFields(excludeFields).Fields(fields).Execute()

Get merge field



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
	mergeId := "mergeId_example" // string | The id for the merge field.
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.GetListsIdMergeFieldsId(context.Background(), listId, mergeId).ExcludeFields(excludeFields).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.GetListsIdMergeFieldsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListsIdMergeFieldsId`: MergeField
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.GetListsIdMergeFieldsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**mergeId** | **string** | The id for the merge field. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListsIdMergeFieldsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 

### Return type

[**MergeField**](MergeField.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListsIdSegmentsId

> List7 GetListsIdSegmentsId(ctx, listId, segmentId).Fields(fields).ExcludeFields(excludeFields).IncludeCleaned(includeCleaned).IncludeTransactional(includeTransactional).IncludeUnsubscribed(includeUnsubscribed).Execute()

Get segment info



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
	segmentId := "segmentId_example" // string | The unique id for the segment.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	includeCleaned := false // bool | Include cleaned members in response (optional)
	includeTransactional := false // bool | Include transactional members in response (optional)
	includeUnsubscribed := false // bool | Include unsubscribed members in response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.GetListsIdSegmentsId(context.Background(), listId, segmentId).Fields(fields).ExcludeFields(excludeFields).IncludeCleaned(includeCleaned).IncludeTransactional(includeTransactional).IncludeUnsubscribed(includeUnsubscribed).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.GetListsIdSegmentsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListsIdSegmentsId`: List7
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.GetListsIdSegmentsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**segmentId** | **string** | The unique id for the segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListsIdSegmentsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **includeCleaned** | **bool** | Include cleaned members in response | 
 **includeTransactional** | **bool** | Include transactional members in response | 
 **includeUnsubscribed** | **bool** | Include unsubscribed members in response | 

### Return type

[**List7**](List7.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListsIdSegmentsIdMembers

> SegmentMembers GetListsIdSegmentsIdMembers(ctx, listId, segmentId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).IncludeCleaned(includeCleaned).IncludeTransactional(includeTransactional).IncludeUnsubscribed(includeUnsubscribed).Execute()

List members in segment



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
	segmentId := "segmentId_example" // string | The unique id for the segment.
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)
	includeCleaned := false // bool | Include cleaned members in response (optional)
	includeTransactional := false // bool | Include transactional members in response (optional)
	includeUnsubscribed := false // bool | Include unsubscribed members in response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.GetListsIdSegmentsIdMembers(context.Background(), listId, segmentId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).IncludeCleaned(includeCleaned).IncludeTransactional(includeTransactional).IncludeUnsubscribed(includeUnsubscribed).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.GetListsIdSegmentsIdMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListsIdSegmentsIdMembers`: SegmentMembers
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.GetListsIdSegmentsIdMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**segmentId** | **string** | The unique id for the segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListsIdSegmentsIdMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **includeCleaned** | **bool** | Include cleaned members in response | 
 **includeTransactional** | **bool** | Include transactional members in response | 
 **includeUnsubscribed** | **bool** | Include unsubscribed members in response | 

### Return type

[**SegmentMembers**](SegmentMembers.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListsIdSignupForms

> ListSignupForms GetListsIdSignupForms(ctx, listId).Execute()

List signup forms



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.GetListsIdSignupForms(context.Background(), listId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.GetListsIdSignupForms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListsIdSignupForms`: ListSignupForms
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.GetListsIdSignupForms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListsIdSignupFormsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListSignupForms**](ListSignupForms.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListsIdSurveys

> GetListsIdSurveys(ctx, listId).Execute()

Get information about all surveys for a list



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListsAPI.GetListsIdSurveys(context.Background(), listId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.GetListsIdSurveys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListsIdSurveysRequest struct via the builder pattern


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


## GetListsIdSurveysId

> GetListsIdSurveysId(ctx, listId, surveyId).Execute()

Get survey



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
	r, err := apiClient.ListsAPI.GetListsIdSurveysId(context.Background(), listId, surveyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.GetListsIdSurveysId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetListsIdSurveysIdRequest struct via the builder pattern


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


## GetListsIdWebhooks

> ListWebhooks GetListsIdWebhooks(ctx, listId).Execute()

List webhooks



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.GetListsIdWebhooks(context.Background(), listId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.GetListsIdWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListsIdWebhooks`: ListWebhooks
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.GetListsIdWebhooks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListsIdWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListWebhooks**](ListWebhooks.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListsIdWebhooksId

> ListWebhooks GetListsIdWebhooksId(ctx, listId, webhookId).Execute()

Get webhook info



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
	webhookId := "webhookId_example" // string | The webhook's id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.GetListsIdWebhooksId(context.Background(), listId, webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.GetListsIdWebhooksId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListsIdWebhooksId`: ListWebhooks
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.GetListsIdWebhooksId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**webhookId** | **string** | The webhook&#39;s id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListsIdWebhooksIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListWebhooks**](ListWebhooks.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchListsId

> SubscriberList PatchListsId(ctx, listId).Body(body).Execute()

Update lists



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
	body := *openapiclient.NewSubscriberList2("Name_example", *openapiclient.NewListContact2("Company_example", "Address1_example", "City_example", "State_example", "Zip_example", "Country_example"), "PermissionReminder_example", *openapiclient.NewCampaignDefaults1("FromName_example", "FromEmail_example", "Subject_example", "Language_example"), false) // SubscriberList2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.PatchListsId(context.Background(), listId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.PatchListsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchListsId`: SubscriberList
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.PatchListsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchListsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SubscriberList2**](SubscriberList2.md) |  | 

### Return type

[**SubscriberList**](SubscriberList.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchListsIdInterestCategoriesId

> InterestCategory PatchListsIdInterestCategoriesId(ctx, listId, interestCategoryId).Body(body).Execute()

Update interest category



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
	interestCategoryId := "interestCategoryId_example" // string | The unique ID for the interest category.
	body := *openapiclient.NewInterestCategory1("Title_example", "Type_example") // InterestCategory1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.PatchListsIdInterestCategoriesId(context.Background(), listId, interestCategoryId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.PatchListsIdInterestCategoriesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchListsIdInterestCategoriesId`: InterestCategory
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.PatchListsIdInterestCategoriesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**interestCategoryId** | **string** | The unique ID for the interest category. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchListsIdInterestCategoriesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**InterestCategory1**](InterestCategory1.md) |  | 

### Return type

[**InterestCategory**](InterestCategory.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchListsIdInterestCategoriesIdInterestsId

> Interest PatchListsIdInterestCategoriesIdInterestsId(ctx, listId, interestCategoryId, interestId).Body(body).Execute()

Update interest in category



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
	interestCategoryId := "interestCategoryId_example" // string | The unique ID for the interest category.
	interestId := "interestId_example" // string | The specific interest or 'group name'.
	body := *openapiclient.NewInterest1("Name_example") // Interest1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.PatchListsIdInterestCategoriesIdInterestsId(context.Background(), listId, interestCategoryId, interestId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.PatchListsIdInterestCategoriesIdInterestsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchListsIdInterestCategoriesIdInterestsId`: Interest
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.PatchListsIdInterestCategoriesIdInterestsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**interestCategoryId** | **string** | The unique ID for the interest category. | 
**interestId** | **string** | The specific interest or &#39;group name&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchListsIdInterestCategoriesIdInterestsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**Interest1**](Interest1.md) |  | 

### Return type

[**Interest**](Interest.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchListsIdMembersId

> ListMembers2 PatchListsIdMembersId(ctx, listId, subscriberHash).Body(body).SkipMergeValidation(skipMergeValidation).Execute()

Update list member



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
	subscriberHash := "subscriberHash_example" // string | The MD5 hash of the lowercase version of the list member's email address. This endpoint also accepts a list member's email address or contact_id.
	body := *openapiclient.NewAddListMembers3() // AddListMembers3 | 
	skipMergeValidation := true // bool | If skip_merge_validation is true, member data will be accepted without merge field values, even if the merge field is usually required. This defaults to false. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.PatchListsIdMembersId(context.Background(), listId, subscriberHash).Body(body).SkipMergeValidation(skipMergeValidation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.PatchListsIdMembersId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchListsIdMembersId`: ListMembers2
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.PatchListsIdMembersId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**subscriberHash** | **string** | The MD5 hash of the lowercase version of the list member&#39;s email address. This endpoint also accepts a list member&#39;s email address or contact_id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchListsIdMembersIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**AddListMembers3**](AddListMembers3.md) |  | 
 **skipMergeValidation** | **bool** | If skip_merge_validation is true, member data will be accepted without merge field values, even if the merge field is usually required. This defaults to false. | 

### Return type

[**ListMembers2**](ListMembers2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchListsIdMembersIdNotesId

> MemberNotes PatchListsIdMembersIdNotesId(ctx, listId, subscriberHash, noteId).Body(body).Execute()

Update note



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
	subscriberHash := "subscriberHash_example" // string | The MD5 hash of the lowercase version of the list member's email address. This endpoint also accepts a list member's email address or contact_id.
	noteId := "noteId_example" // string | The id for the note.
	body := *openapiclient.NewMemberNotes1() // MemberNotes1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.PatchListsIdMembersIdNotesId(context.Background(), listId, subscriberHash, noteId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.PatchListsIdMembersIdNotesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchListsIdMembersIdNotesId`: MemberNotes
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.PatchListsIdMembersIdNotesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**subscriberHash** | **string** | The MD5 hash of the lowercase version of the list member&#39;s email address. This endpoint also accepts a list member&#39;s email address or contact_id. | 
**noteId** | **string** | The id for the note. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchListsIdMembersIdNotesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**MemberNotes1**](MemberNotes1.md) |  | 

### Return type

[**MemberNotes**](MemberNotes.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchListsIdMergeFieldsId

> MergeField PatchListsIdMergeFieldsId(ctx, listId, mergeId).Body(body).Execute()

Update merge field



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
	mergeId := "mergeId_example" // string | The id for the merge field.
	body := *openapiclient.NewMergeField2("Name_example") // MergeField2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.PatchListsIdMergeFieldsId(context.Background(), listId, mergeId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.PatchListsIdMergeFieldsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchListsIdMergeFieldsId`: MergeField
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.PatchListsIdMergeFieldsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**mergeId** | **string** | The id for the merge field. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchListsIdMergeFieldsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MergeField2**](MergeField2.md) |  | 

### Return type

[**MergeField**](MergeField.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchListsIdSegmentsId

> List7 PatchListsIdSegmentsId(ctx, listId, segmentId).Body(body).Execute()

Update segment



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
	segmentId := "segmentId_example" // string | The unique id for the segment.
	body := *openapiclient.NewList9("Name_example") // List9 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.PatchListsIdSegmentsId(context.Background(), listId, segmentId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.PatchListsIdSegmentsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchListsIdSegmentsId`: List7
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.PatchListsIdSegmentsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**segmentId** | **string** | The unique id for the segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchListsIdSegmentsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**List9**](List9.md) |  | 

### Return type

[**List7**](List7.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchListsIdWebhooksId

> ListWebhooks PatchListsIdWebhooksId(ctx, listId, webhookId).Body(body).Execute()

Update webhook



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
	webhookId := "webhookId_example" // string | The webhook's id.
	body := *openapiclient.NewAddWebhook() // AddWebhook | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.PatchListsIdWebhooksId(context.Background(), listId, webhookId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.PatchListsIdWebhooksId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchListsIdWebhooksId`: ListWebhooks
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.PatchListsIdWebhooksId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**webhookId** | **string** | The webhook&#39;s id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchListsIdWebhooksIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**AddWebhook**](AddWebhook.md) |  | 

### Return type

[**ListWebhooks**](ListWebhooks.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostListMemberEvents

> PostListMemberEvents(ctx, listId, subscriberHash).Body(body).Execute()

Add event



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
	subscriberHash := "subscriberHash_example" // string | The MD5 hash of the lowercase version of the list member's email address. This endpoint also accepts a list member's email address or contact_id.
	body := *openapiclient.NewEvents("Name_example") // Events | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListsAPI.PostListMemberEvents(context.Background(), listId, subscriberHash).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.PostListMemberEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**subscriberHash** | **string** | The MD5 hash of the lowercase version of the list member&#39;s email address. This endpoint also accepts a list member&#39;s email address or contact_id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostListMemberEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**Events**](Events.md) |  | 

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


## PostListMemberTags

> PostListMemberTags(ctx, listId, subscriberHash).Body(body).Execute()

Add or remove member tags



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
	subscriberHash := "subscriberHash_example" // string | The MD5 hash of the lowercase version of the list member's email address.
	body := *openapiclient.NewMemberTags([]openapiclient.MemberTag{*openapiclient.NewMemberTag("Name_example", "Status_example")}) // MemberTags | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListsAPI.PostListMemberTags(context.Background(), listId, subscriberHash).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.PostListMemberTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**subscriberHash** | **string** | The MD5 hash of the lowercase version of the list member&#39;s email address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostListMemberTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MemberTags**](MemberTags.md) |  | 

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


## PostLists

> SubscriberList PostLists(ctx).Body(body).Execute()

Add list



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
	body := *openapiclient.NewSubscriberList1("Name_example", *openapiclient.NewListContact1("Company_example", "Address1_example", "City_example", "Country_example"), "PermissionReminder_example", *openapiclient.NewCampaignDefaults1("FromName_example", "FromEmail_example", "Subject_example", "Language_example"), false) // SubscriberList1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.PostLists(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.PostLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLists`: SubscriberList
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.PostLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SubscriberList1**](SubscriberList1.md) |  | 

### Return type

[**SubscriberList**](SubscriberList.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostListsId

> BatchUpdateListMembers PostListsId(ctx, listId).Body(body).SkipMergeValidation(skipMergeValidation).SkipDuplicateCheck(skipDuplicateCheck).Execute()

Batch subscribe or unsubscribe



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
	body := *openapiclient.NewMembersToSubscribeUnsubscribeToFromAListInBatch([]openapiclient.AddListMembers{*openapiclient.NewAddListMembers()}) // MembersToSubscribeUnsubscribeToFromAListInBatch | 
	skipMergeValidation := true // bool | If skip_merge_validation is true, member data will be accepted without merge field values, even if the merge field is usually required. This defaults to false. (optional)
	skipDuplicateCheck := true // bool | If skip_duplicate_check is true, we will ignore duplicates sent in the request when using the batch sub/unsub on the lists endpoint. The status of the first appearance in the request will be saved. This defaults to false. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.PostListsId(context.Background(), listId).Body(body).SkipMergeValidation(skipMergeValidation).SkipDuplicateCheck(skipDuplicateCheck).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.PostListsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostListsId`: BatchUpdateListMembers
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.PostListsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostListsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MembersToSubscribeUnsubscribeToFromAListInBatch**](MembersToSubscribeUnsubscribeToFromAListInBatch.md) |  | 
 **skipMergeValidation** | **bool** | If skip_merge_validation is true, member data will be accepted without merge field values, even if the merge field is usually required. This defaults to false. | 
 **skipDuplicateCheck** | **bool** | If skip_duplicate_check is true, we will ignore duplicates sent in the request when using the batch sub/unsub on the lists endpoint. The status of the first appearance in the request will be saved. This defaults to false. | 

### Return type

[**BatchUpdateListMembers**](BatchUpdateListMembers.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostListsIdInterestCategories

> InterestCategory PostListsIdInterestCategories(ctx, listId).Body(body).Execute()

Add interest category



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
	body := *openapiclient.NewInterestCategory1("Title_example", "Type_example") // InterestCategory1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.PostListsIdInterestCategories(context.Background(), listId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.PostListsIdInterestCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostListsIdInterestCategories`: InterestCategory
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.PostListsIdInterestCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostListsIdInterestCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**InterestCategory1**](InterestCategory1.md) |  | 

### Return type

[**InterestCategory**](InterestCategory.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostListsIdInterestCategoriesIdInterests

> Interest PostListsIdInterestCategoriesIdInterests(ctx, listId, interestCategoryId).Body(body).Execute()

Add interest in category



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
	interestCategoryId := "interestCategoryId_example" // string | The unique ID for the interest category.
	body := *openapiclient.NewInterest1("Name_example") // Interest1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.PostListsIdInterestCategoriesIdInterests(context.Background(), listId, interestCategoryId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.PostListsIdInterestCategoriesIdInterests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostListsIdInterestCategoriesIdInterests`: Interest
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.PostListsIdInterestCategoriesIdInterests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**interestCategoryId** | **string** | The unique ID for the interest category. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostListsIdInterestCategoriesIdInterestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**Interest1**](Interest1.md) |  | 

### Return type

[**Interest**](Interest.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostListsIdMembers

> ListMembers2 PostListsIdMembers(ctx, listId).Body(body).SkipMergeValidation(skipMergeValidation).Execute()

Add member to list



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
	body := *openapiclient.NewAddListMembers1("EmailAddress_example", "Status_example") // AddListMembers1 | 
	skipMergeValidation := true // bool | If skip_merge_validation is true, member data will be accepted without merge field values, even if the merge field is usually required. This defaults to false. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.PostListsIdMembers(context.Background(), listId).Body(body).SkipMergeValidation(skipMergeValidation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.PostListsIdMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostListsIdMembers`: ListMembers2
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.PostListsIdMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostListsIdMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AddListMembers1**](AddListMembers1.md) |  | 
 **skipMergeValidation** | **bool** | If skip_merge_validation is true, member data will be accepted without merge field values, even if the merge field is usually required. This defaults to false. | 

### Return type

[**ListMembers2**](ListMembers2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostListsIdMembersHashActionsDeletePermanent

> PostListsIdMembersHashActionsDeletePermanent(ctx, listId, subscriberHash).Execute()

Delete list member



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
	subscriberHash := "subscriberHash_example" // string | The MD5 hash of the lowercase version of the list member's email address.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListsAPI.PostListsIdMembersHashActionsDeletePermanent(context.Background(), listId, subscriberHash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.PostListsIdMembersHashActionsDeletePermanent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**subscriberHash** | **string** | The MD5 hash of the lowercase version of the list member&#39;s email address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostListsIdMembersHashActionsDeletePermanentRequest struct via the builder pattern


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


## PostListsIdMembersIdNotes

> MemberNotes PostListsIdMembersIdNotes(ctx, listId, subscriberHash).Body(body).Execute()

Add member note



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
	subscriberHash := "subscriberHash_example" // string | The MD5 hash of the lowercase version of the list member's email address.
	body := *openapiclient.NewMemberNotes1() // MemberNotes1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.PostListsIdMembersIdNotes(context.Background(), listId, subscriberHash).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.PostListsIdMembersIdNotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostListsIdMembersIdNotes`: MemberNotes
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.PostListsIdMembersIdNotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**subscriberHash** | **string** | The MD5 hash of the lowercase version of the list member&#39;s email address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostListsIdMembersIdNotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MemberNotes1**](MemberNotes1.md) |  | 

### Return type

[**MemberNotes**](MemberNotes.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostListsIdMergeFields

> MergeField PostListsIdMergeFields(ctx, listId).Body(body).Execute()

Add merge field



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
	body := *openapiclient.NewMergeField1("Name_example", "Type_example") // MergeField1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.PostListsIdMergeFields(context.Background(), listId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.PostListsIdMergeFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostListsIdMergeFields`: MergeField
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.PostListsIdMergeFields`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostListsIdMergeFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MergeField1**](MergeField1.md) |  | 

### Return type

[**MergeField**](MergeField.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostListsIdSegments

> List7 PostListsIdSegments(ctx, listId).Body(body).Execute()

Add segment



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
	body := *openapiclient.NewList8("Name_example") // List8 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.PostListsIdSegments(context.Background(), listId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.PostListsIdSegments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostListsIdSegments`: List7
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.PostListsIdSegments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostListsIdSegmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**List8**](List8.md) |  | 

### Return type

[**List7**](List7.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostListsIdSegmentsId

> BatchAddRemoveListMembersToFromStaticSegment PostListsIdSegmentsId(ctx, listId, segmentId).Body(body).Execute()

Batch add or remove members



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
	segmentId := "segmentId_example" // string | The unique id for the segment.
	body := *openapiclient.NewMembersToAddRemoveToFromAStaticSegment() // MembersToAddRemoveToFromAStaticSegment | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.PostListsIdSegmentsId(context.Background(), listId, segmentId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.PostListsIdSegmentsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostListsIdSegmentsId`: BatchAddRemoveListMembersToFromStaticSegment
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.PostListsIdSegmentsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**segmentId** | **string** | The unique id for the segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostListsIdSegmentsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MembersToAddRemoveToFromAStaticSegment**](MembersToAddRemoveToFromAStaticSegment.md) |  | 

### Return type

[**BatchAddRemoveListMembersToFromStaticSegment**](BatchAddRemoveListMembersToFromStaticSegment.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostListsIdSegmentsIdMembers

> ListMembers1 PostListsIdSegmentsIdMembers(ctx, listId, segmentId).Body(body).Execute()

Add member to segment



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
	segmentId := "segmentId_example" // string | The unique id for the segment.
	body := *openapiclient.NewPostListsIdSegmentsIdMembersRequest("EmailAddress_example") // PostListsIdSegmentsIdMembersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.PostListsIdSegmentsIdMembers(context.Background(), listId, segmentId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.PostListsIdSegmentsIdMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostListsIdSegmentsIdMembers`: ListMembers1
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.PostListsIdSegmentsIdMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**segmentId** | **string** | The unique id for the segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostListsIdSegmentsIdMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**PostListsIdSegmentsIdMembersRequest**](PostListsIdSegmentsIdMembersRequest.md) |  | 

### Return type

[**ListMembers1**](ListMembers1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostListsIdSignupForms

> SignupForm PostListsIdSignupForms(ctx, listId).Body(body).Execute()

Customize signup form



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
	body := *openapiclient.NewSignupForm1() // SignupForm1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.PostListsIdSignupForms(context.Background(), listId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.PostListsIdSignupForms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostListsIdSignupForms`: SignupForm
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.PostListsIdSignupForms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostListsIdSignupFormsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SignupForm1**](SignupForm1.md) |  | 

### Return type

[**SignupForm**](SignupForm.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostListsIdWebhooks

> ListWebhooks PostListsIdWebhooks(ctx, listId).Body(body).Execute()

Add webhook



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
	body := *openapiclient.NewAddWebhook() // AddWebhook | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.PostListsIdWebhooks(context.Background(), listId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.PostListsIdWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostListsIdWebhooks`: ListWebhooks
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.PostListsIdWebhooks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostListsIdWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AddWebhook**](AddWebhook.md) |  | 

### Return type

[**ListWebhooks**](ListWebhooks.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewASegment

> CollectionOfSegments PreviewASegment(ctx, listId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Type_(type_).SinceCreatedAt(sinceCreatedAt).BeforeCreatedAt(beforeCreatedAt).IncludeCleaned(includeCleaned).IncludeTransactional(includeTransactional).IncludeUnsubscribed(includeUnsubscribed).SinceUpdatedAt(sinceUpdatedAt).BeforeUpdatedAt(beforeUpdatedAt).Execute()

List segments



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
	fields := []string{"Inner_example"} // []string | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. (optional)
	excludeFields := []string{"Inner_example"} // []string | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. (optional)
	count := int32(56) // int32 | The number of records to return. Default value is 10. Maximum value is 1000 (optional) (default to 10)
	offset := int32(56) // int32 | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. (optional) (default to 0)
	type_ := "type__example" // string | Limit results based on segment type. (optional)
	sinceCreatedAt := "sinceCreatedAt_example" // string | Restrict results to segments created after the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. (optional)
	beforeCreatedAt := "beforeCreatedAt_example" // string | Restrict results to segments created before the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. (optional)
	includeCleaned := false // bool | Include cleaned members in response (optional)
	includeTransactional := false // bool | Include transactional members in response (optional)
	includeUnsubscribed := false // bool | Include unsubscribed members in response (optional)
	sinceUpdatedAt := "sinceUpdatedAt_example" // string | Restrict results to segments update after the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. (optional)
	beforeUpdatedAt := "beforeUpdatedAt_example" // string | Restrict results to segments update before the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.PreviewASegment(context.Background(), listId).Fields(fields).ExcludeFields(excludeFields).Count(count).Offset(offset).Type_(type_).SinceCreatedAt(sinceCreatedAt).BeforeCreatedAt(beforeCreatedAt).IncludeCleaned(includeCleaned).IncludeTransactional(includeTransactional).IncludeUnsubscribed(includeUnsubscribed).SinceUpdatedAt(sinceUpdatedAt).BeforeUpdatedAt(beforeUpdatedAt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.PreviewASegment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PreviewASegment`: CollectionOfSegments
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.PreviewASegment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewASegmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | **[]string** | A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **int32** | The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **int32** | Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **type_** | **string** | Limit results based on segment type. | 
 **sinceCreatedAt** | **string** | Restrict results to segments created after the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **beforeCreatedAt** | **string** | Restrict results to segments created before the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **includeCleaned** | **bool** | Include cleaned members in response | 
 **includeTransactional** | **bool** | Include transactional members in response | 
 **includeUnsubscribed** | **bool** | Include unsubscribed members in response | 
 **sinceUpdatedAt** | **string** | Restrict results to segments update after the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **beforeUpdatedAt** | **string** | Restrict results to segments update before the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 

### Return type

[**CollectionOfSegments**](CollectionOfSegments.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutListsIdMembersId

> ListMembers2 PutListsIdMembersId(ctx, listId, subscriberHash).Body(body).SkipMergeValidation(skipMergeValidation).Execute()

Add or update list member



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
	subscriberHash := "subscriberHash_example" // string | The MD5 hash of the lowercase version of the list member's email address. This endpoint also accepts a list member's email address or contact_id.
	body := *openapiclient.NewAddListMembers2("EmailAddress_example", "StatusIfNew_example") // AddListMembers2 | 
	skipMergeValidation := true // bool | If skip_merge_validation is true, member data will be accepted without merge field values, even if the merge field is usually required. This defaults to false. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.PutListsIdMembersId(context.Background(), listId, subscriberHash).Body(body).SkipMergeValidation(skipMergeValidation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.PutListsIdMembersId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutListsIdMembersId`: ListMembers2
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.PutListsIdMembersId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 
**subscriberHash** | **string** | The MD5 hash of the lowercase version of the list member&#39;s email address. This endpoint also accepts a list member&#39;s email address or contact_id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutListsIdMembersIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**AddListMembers2**](AddListMembers2.md) |  | 
 **skipMergeValidation** | **bool** | If skip_merge_validation is true, member data will be accepted without merge field values, even if the merge field is usually required. This defaults to false. | 

### Return type

[**ListMembers2**](ListMembers2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchTagsByName

> TagSearchResults SearchTagsByName(ctx, listId).Name(name).Execute()

Search for tags on a list by name.



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
	name := "name_example" // string | The search query used to filter tags.  The search query will be compared to each tag as a prefix, so all tags that have a name starting with this field will be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.SearchTagsByName(context.Background(), listId).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.SearchTagsByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchTagsByName`: TagSearchResults
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.SearchTagsByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | The unique ID for the list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchTagsByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | The search query used to filter tags.  The search query will be compared to each tag as a prefix, so all tags that have a name starting with this field will be returned. | 

### Return type

[**TagSearchResults**](TagSearchResults.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

