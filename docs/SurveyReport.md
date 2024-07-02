# # SurveyReport
The report for a survey.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **string** | A string that uniquely identifies this survey.  | [optional] [readonly]
**WebId**| **int32** | The ID used in the Mailchimp web application. View this survey report in your Mailchimp account at &#x60;https://{dc}.admin.mailchimp.com/lists/surveys/results?survey_id&#x3D;{web_id}&#x60;.  | [optional] [readonly]
**ListId**| **string** | The ID of the list connected to this survey.  | [optional] [readonly]
**ListName**| **string** | The name of the list connected to this survey.  | [optional] [readonly]
**Title**| **string** | The title of the survey.  | [optional] [readonly]
**Url**| **string** | The URL for the survey.  | [optional] [readonly]
**Status**| **string** | The survey&#39;s status. for more information please, see Model/string.php  | [optional] [readonly]
**PublishedAt**| [**time.Time**](time.Time.md) | The date and time the survey was published in ISO 8601 format.  | [optional] [readonly]
**CreatedAt**| [**time.Time**](time.Time.md) | The date and time the survey was created in ISO 8601 format.  | [optional] [readonly]
**UpdatedAt**| [**time.Time**](time.Time.md) | The date and time the survey was last updated in ISO 8601 format.  | [optional] [readonly]
**TotalResponses**| **int32** | The total number of responses to this survey.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

