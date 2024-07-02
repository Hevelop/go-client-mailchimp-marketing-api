# # MemberActivityEvents1
The member activity events for a given member.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activity**| **[]map[string]interface{}** | An array of objects, each representing a contact event. There are multiple possible types, see the [activity schema documentation](https://mailchimp.com/developer/marketing/docs/alternative-schemas/#activity-schemas).  | [optional] [readonly]
**EmailId**| **string** | The MD5 hash of the lowercase version of the list member&#39;s email address.  | [optional] [readonly]
**ListId**| **string** | The list id.  | [optional] [readonly]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

