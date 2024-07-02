# # ECommerceProduct1
Information about a specific product.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **string** | A unique identifier for the product.  |
**Title**| **string** | The title of a product.  |
**Handle**| **string** | The handle of a product.  | [optional]
**Url**| **string** | The URL for a product.  | [optional]
**Description**| **string** | The description of a product.  | [optional]
**Type**| **string** | The type of product.  | [optional]
**Vendor**| **string** | The vendor for a product.  | [optional]
**ImageUrl**| **string** | The image URL for a product.  | [optional]
**Variants**| [**[]ECommerceProductVariant1**](ECommerceProductVariant1.md) | An array of the product&#39;s variants. At least one variant is required for each product. A variant can use the same &#x60;id&#x60; and &#x60;title&#x60; as the parent product.  |
**Images**| [**[]ECommerceProductImage1**](ECommerceProductImage1.md) | An array of the product&#39;s images.  | [optional]
**PublishedAtForeign**| [**time.Time**](time.Time.md) | The date and time the product was published.  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

