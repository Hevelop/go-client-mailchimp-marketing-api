# # List9
Information about a specific list segment.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name**| **string** | The name of the segment.  |
**StaticSegment**| **[]string** | An array of emails to be used for a static segment. Any emails provided that are not present on the list will be ignored. Passing an empty array for an existing static segment will reset that segment and remove all members. This field cannot be provided with the &#x60;options&#x60; field.  | [optional]
**Options**| [**Conditions2**](Conditions2.md) |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

