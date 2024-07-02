# # CollectionOfNotes
The last 10 notes for a specific list member, based on date created.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notes**| [**[]MemberNotes**](MemberNotes.md) | An array of objects, each representing a note resource.  | [optional]
**EmailId**| **string** | The MD5 hash of the lowercase version of the list member&#39;s email address.  | [optional] [readonly]
**ListId**| **string** | The list id.  | [optional] [readonly]
**TotalItems**| **int32** | The total number of items matching the query regardless of pagination.  | [optional] [readonly]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

