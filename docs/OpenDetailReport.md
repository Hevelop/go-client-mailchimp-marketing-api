# # OpenDetailReport
A detailed report of any campaign emails that were opened by a list member.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members**| [**[]OpenActivity**](OpenActivity.md) | An array of objects, each representing a list member who opened a campaign email. Each members object will contain information about the number of total opens by a single member, as well as timestamps for each open event.  | [optional]
**CampaignId**| **string** | The campaign id.  | [optional] [readonly]
**TotalOpens**| **int32** | The total number of opens matching the query regardless of pagination.  | [optional]
**TotalItems**| **int32** | The total number of items matching the query regardless of pagination.  | [optional] [readonly]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

