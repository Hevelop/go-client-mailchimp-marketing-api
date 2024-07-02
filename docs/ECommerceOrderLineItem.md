# # ECommerceOrderLineItem
Information about a specific order line.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **string** | A unique identifier for an order line item.  | [optional] [readonly]
**ProductId**| **string** | A unique identifier for the product associated with an order line item.  | [optional]
**ProductTitle**| **string** | The name of the product for an order line item.  | [optional] [readonly]
**ProductVariantId**| **string** | A unique identifier for the product variant associated with an order line item.  | [optional]
**ProductVariantTitle**| **string** | The name of the product variant for an order line item.  | [optional] [readonly]
**ImageUrl**| **string** | The image URL for a product.  | [optional] [readonly]
**Quantity**| **int32** | The order line item quantity.  | [optional]
**Price**| **float32** | The order line item price.  | [optional]
**Discount**| **float32** | The total discount amount applied to a line item.  | [optional]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

