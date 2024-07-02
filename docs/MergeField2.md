# # MergeField2
A [merge field](https://mailchimp.com/developer/marketing/docs/merge-fields/) for an audience.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag**| **string** | The merge tag used for Mailchimp campaigns and [adding contact information](https://mailchimp.com/developer/marketing/docs/merge-fields/#add-merge-data-to-contacts).  | [optional]
**Name**| **string** | The name of the merge field (audience field).  |
**Required**| **bool** | Whether the merge field is required to import a contact.  | [optional]
**DefaultValue**| **string** | The default value for the merge field if &#x60;null&#x60;.  | [optional]
**Public**| **bool** | Whether the merge field is displayed on the signup form.  | [optional]
**DisplayOrder**| **int32** | The order that the merge field displays on the list signup form.  | [optional]
**Options**| [**MergeFieldOptions2**](MergeFieldOptions2.md) |   | [optional]
**HelpText**| **string** | Extra text to help the subscriber fill out the form.  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

