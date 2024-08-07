# # List7
Information about a specific segment.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **int32** | The unique id for the segment.  | [optional] [readonly]
**Name**| **string** | The name of the segment.  | [optional]
**MemberCount**| **int32** | The number of active subscribers currently included in the segment.  | [optional] [readonly]
**Type**| **string** | The type of segment. Static segments are now known as tags. Learn more about [tags](https://mailchimp.com/help/getting-started-tags?utm_source&#x3D;mc-api&amp;utm_medium&#x3D;docs&amp;utm_campaign&#x3D;apidocs). for more information please, see Model/string.php  | [optional] [readonly]
**CreatedAt**| [**time.Time**](time.Time.md) | The date and time the segment was created in ISO 8601 format.  | [optional] [readonly]
**UpdatedAt**| [**time.Time**](time.Time.md) | The date and time the segment was last updated in ISO 8601 format.  | [optional] [readonly]
**Options**| [**Conditions**](Conditions.md) |   | [optional]
**ListId**| **string** | The list id.  | [optional] [readonly]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

