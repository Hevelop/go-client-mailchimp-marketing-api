# # SegmentOptions2
An object representing all segmentation options. This object should contain a &#x60;saved_segment_id&#x60; to use an existing segment, or you can create a new segment by including both &#x60;match&#x60; and &#x60;conditions&#x60; options.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SavedSegmentId**| **int32** | The id for an existing saved segment.  | [optional]
**Match**| **string** | Segment match type. for more information please, see Model/string.php  | [optional]
**Conditions**| **[]map[string]interface{}** | Segment match conditions. There are multiple possible types, see the [condition types documentation](https://mailchimp.com/developer/marketing/docs/alternative-schemas/#segment-condition-schemas).  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

