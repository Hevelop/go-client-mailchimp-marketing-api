# # MergeField
A [merge field](https://mailchimp.com/developer/marketing/docs/merge-fields/) for an audience.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MergeId**| **int32** | An unchanging id for the merge field.  | [optional] [readonly]
**Tag**| **string** | The merge tag used for Mailchimp campaigns and [adding contact information](https://mailchimp.com/developer/marketing/docs/merge-fields/#add-merge-data-to-contacts).  | [optional]
**Name**| **string** | The name of the merge field (audience field).  | [optional]
**Type**| **string** | The [type](https://mailchimp.com/developer/marketing/docs/merge-fields/#structure) for the merge field. for more information please, see Model/string.php  | [optional] [readonly]
**Required**| **bool** | The boolean value if the merge field is required.  | [optional]
**DefaultValue**| **string** | The default value for the merge field if &#x60;null&#x60;.  | [optional]
**Public**| **bool** | Whether the merge field is displayed on the signup form.  | [optional]
**DisplayOrder**| **int32** | The order that the merge field displays on the list signup form.  | [optional]
**Options**| [**MergeFieldOptions**](MergeFieldOptions.md) |   | [optional]
**HelpText**| **string** | Extra text to help the subscriber fill out the form.  | [optional]
**ListId**| **string** | The ID that identifies this merge field&#39;s audience&#39;.  | [optional] [readonly]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

