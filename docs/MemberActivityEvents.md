# # MemberActivityEvents
The last 50 member events for a list.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activity**| [**[]MemberActivity**](MemberActivity.md) | An array of objects, each representing a member event.  | [optional]
**EmailId**| **string** | The MD5 hash of the lowercase version of the list member&#39;s email address.  | [optional] [readonly]
**ContactId**| **string** | As Mailchimp evolves beyond email, you may eventually have contacts without email addresses. While the &#x60;email_id&#x60; is the MD5 hash of their email address, this &#x60;contact_id&#x60; is agnostic of contact’s inclusion of an email address.  | [optional] [readonly]
**ListId**| **string** | The list id.  | [optional] [readonly]
**TotalItems**| **int32** | The total number of items matching the query regardless of pagination.  | [optional] [readonly]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

