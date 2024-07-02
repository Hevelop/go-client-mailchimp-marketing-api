# # CampaignFeedback2
A specific feedback message from a specific campaign.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeedbackId**| **int32** | The individual id for the feedback item.  | [optional] [readonly]
**ParentId**| **int32** | If a reply, the id of the parent feedback item.  | [optional] [readonly]
**BlockId**| **int32** | The block id for the editable block that the feedback addresses.  | [optional]
**Message**| **string** | The content of the feedback.  | [optional]
**IsComplete**| **bool** | The status of feedback.  | [optional]
**CreatedBy**| **string** | The login name of the user who created the feedback.  | [optional] [readonly]
**CreatedAt**| [**time.Time**](time.Time.md) | The date and time the feedback item was created in ISO 8601 format.  | [optional] [readonly]
**UpdatedAt**| [**time.Time**](time.Time.md) | The date and time the feedback was last updated in ISO 8601 format.  | [optional] [readonly]
**Source**| **string** | The source of the feedback. for more information please, see Model/string.php  | [optional] [readonly]
**CampaignId**| **string** | The unique id for the campaign.  | [optional] [readonly]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

