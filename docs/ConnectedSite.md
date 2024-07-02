# # ConnectedSite
Information about a specific connected site.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForeignId**| **string** | The unique identifier for the site.  | [optional] [readonly]
**StoreId**| **string** | The unique identifier for the ecommerce store that&#39;s associated with the connected site (if any). The store_id for a specific connected site can&#39;t change.  | [optional] [readonly]
**Platform**| **string** | The platform of the connected site.  | [optional] [readonly]
**Domain**| **string** | The connected site domain.  | [optional] [readonly]
**SiteScript**| [**Script**](Script.md) |   | [optional]
**CreatedAt**| [**time.Time**](time.Time.md) | The date and time the connected site was created in ISO 8601 format.  | [optional] [readonly]
**UpdatedAt**| [**time.Time**](time.Time.md) | The date and time the connected site was last updated in ISO 8601 format.  | [optional] [readonly]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

