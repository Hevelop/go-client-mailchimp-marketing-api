# # MemberNotes
A specific note for a specific member.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **int32** | The note id.  | [optional] [readonly]
**CreatedAt**| [**time.Time**](time.Time.md) | The date and time the note was created in ISO 8601 format.  | [optional] [readonly]
**CreatedBy**| **string** | The author of the note.  | [optional] [readonly]
**UpdatedAt**| [**time.Time**](time.Time.md) | The date and time the note was last updated in ISO 8601 format.  | [optional] [readonly]
**Note**| **string** | The content of the note.  | [optional]
**ListId**| **string** | The unique id for the list.  | [optional] [readonly]
**EmailId**| **string** | The MD5 hash of the lowercase version of the list member&#39;s email address.  | [optional] [readonly]
**ContactId**| **string** | As Mailchimp evolves beyond email, you may eventually have contacts without email addresses. While the &#x60;email_id&#x60; is the MD5 hash of their email address, this &#x60;contact_id&#x60; is agnostic of contact’s inclusion of an email address.  | [optional] [readonly]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

