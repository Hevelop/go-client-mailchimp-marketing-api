# # List8
Information about a specific list segment.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name**| **string** | The name of the segment.  |
**StaticSegment**| **[]string** | An array of emails to be used for a static segment. Any emails provided that are not present on the list will be ignored. Passing an empty array will create a static segment without any subscribers. This field cannot be provided with the options field.  | [optional]
**Options**| [**Conditions1**](Conditions1.md) |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

