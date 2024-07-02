# # SubscriberInAutomationQueue2
Information about subscribers in an Automation email queue.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **string** | The MD5 hash of the lowercase version of the list member&#39;s email address.  | [optional] [readonly]
**WorkflowId**| **string** | A string that uniquely identifies an Automation workflow.  | [optional] [readonly]
**EmailId**| **string** | A string that uniquely identifies an email in an Automation workflow.  | [optional] [readonly]
**ListId**| **string** | A string that uniquely identifies a list.  | [optional] [readonly]
**ListIsActive**| **bool** | The status of the list used, namely if it&#39;s deleted or disabled.  | [optional] [readonly]
**EmailAddress**| **string** | The list member&#39;s email address.  | [optional]
**NextSend**| [**time.Time**](time.Time.md) | The date and time of the next send for the workflow email in ISO 8601 format.  | [optional] [readonly]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

