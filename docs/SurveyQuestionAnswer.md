# # SurveyQuestionAnswer
The details of a survey question&#39;s answer.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **string** | The ID of the answer.  | [optional] [readonly]
**Value**| **string** | The raw text answer.  | [optional] [readonly]
**ResponseId**| **string** | The ID of the survey response.  | [optional] [readonly]
**SubmittedAt**| [**time.Time**](time.Time.md) | The date and time when the survey response was submitted in ISO 8601 format.  | [optional] [readonly]
**Contact**| [**Contact**](Contact.md) |   | [optional]
**IsNewContact**| **bool** | If this contact was added to the Mailchimp audience via this survey.  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

