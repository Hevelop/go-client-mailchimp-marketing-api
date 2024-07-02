# # SurveyResponse
A single survey response.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId**| **string** | The ID for the survey response.  | [optional] [readonly]
**SubmittedAt**| [**time.Time**](time.Time.md) | The date and time when the survey response was submitted in ISO 8601 format.  | [optional] [readonly]
**Contact**| [**Contact**](Contact.md) |   | [optional]
**IsNewContact**| **bool** | If this contact was added to the Mailchimp audience via this survey.  | [optional]
**Results**| [**[]Response**](Response.md) | The survey questions and the answers to those questions.  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

