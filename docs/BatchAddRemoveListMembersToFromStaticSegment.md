# # BatchAddRemoveListMembersToFromStaticSegment
Batch add/remove List members to/from static segment

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MembersAdded**| [**[]ListMembers**](ListMembers.md) | An array of objects, each representing a new member that was added to the static segment.  | [optional]
**MembersRemoved**| [**[]ListMembers**](ListMembers.md) | An array of objects, each representing an existing list member that got deleted from the static segment.  | [optional]
**Errors**| [**[]ErrorsInner1**](ErrorsInner1.md) | An array of objects, each representing an array of email addresses that could not be added to the segment or removed and an error message providing more details.  | [optional]
**TotalAdded**| **int32** | The total number of items matching the query, irrespective of pagination.  | [optional]
**TotalRemoved**| **int32** | The total number of items matching the query, irrespective of pagination.  | [optional]
**ErrorCount**| **int32** | The total number of items matching the query, irrespective of pagination.  | [optional]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

