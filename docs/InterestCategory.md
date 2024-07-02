# # InterestCategory
Interest categories organize interests, which are used to group subscribers based on their preferences. These correspond to Group Titles the application.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListId**| **string** | The unique list id for the category.  | [optional] [readonly]
**Id**| **string** | The id for the interest category.  | [optional] [readonly]
**Title**| **string** | The text description of this category. This field appears on signup forms and is often phrased as a question.  | [optional]
**DisplayOrder**| **int32** | The order that the categories are displayed in the list. Lower numbers display first.  | [optional]
**Type**| **string** | Determines how this category’s interests appear on signup forms. for more information please, see Model/string.php  | [optional]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

