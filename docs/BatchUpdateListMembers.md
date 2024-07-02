# # BatchUpdateListMembers
Batch update list members.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewMembers**| [**[]ListMembers**](ListMembers.md) | An array of objects, each representing a new member that was added to the list.  | [optional]
**UpdatedMembers**| [**[]ListMembers**](ListMembers.md) | An array of objects, each representing an existing list member whose subscription status was updated.  | [optional]
**Errors**| [**[]ErrorsInner**](ErrorsInner.md) | An array of objects, each representing an email address that could not be added to the list or updated and an error message providing more details.  | [optional]
**TotalCreated**| **int32** | The total number of items matching the query, irrespective of pagination.  | [optional]
**TotalUpdated**| **int32** | The total number of items matching the query, irrespective of pagination.  | [optional]
**ErrorCount**| **int32** | The total number of items matching the query, irrespective of pagination.  | [optional]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

