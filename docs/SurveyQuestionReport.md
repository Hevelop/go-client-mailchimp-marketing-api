# # SurveyQuestionReport
The details of a survey question&#39;s report.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **string** | The ID of the survey question.  | [optional] [readonly]
**SurveyId**| **string** | The unique ID of the survey.  | [optional] [readonly]
**Query**| **string** | The query of the survey question.  | [optional] [readonly]
**Type**| **string** | The response type of the survey question. for more information please, see Model/string.php  | [optional] [readonly]
**TotalResponses**| **int32** | The total number of responses to this question.  | [optional] [readonly]
**IsRequired**| **bool** | Whether this survey question is required to answer.  | [optional] [readonly]
**HasOther**| **bool** | Whether this survey question has an &#39;other&#39; option.  | [optional] [readonly]
**OtherLabel**| **string** | Label used for the &#39;other&#39; option of this survey question.  | [optional] [readonly]
**AverageRating**| **float32** | The average rating for this range question.  | [optional] [readonly]
**RangeLowLabel**| **string** | Label for the low end of the range.  | [optional] [readonly]
**RangeHighLabel**| **string** | Label for the high end of the range.  | [optional] [readonly]
**PlaceholderLabel**| **string** | Placeholder text for this survey question&#39;s answer box.  | [optional] [readonly]
**SubscribeCheckboxEnabled**| **bool** | Whether the subscribe checkbox is shown for this email question.  | [optional] [readonly]
**SubscribeCheckboxLabel**| **string** | Label used for the subscribe checkbox for this email question.  | [optional] [readonly]
**MergeField**| [**MergeField3**](MergeField3.md) |   | [optional]
**Options**| [**[]OptionsInner**](OptionsInner.md) | The answer choices for this question.  | [optional] [readonly]
**ContactCounts**| [**ContactCounts**](ContactCounts.md) |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

