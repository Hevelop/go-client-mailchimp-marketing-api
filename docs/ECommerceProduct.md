# # ECommerceProduct
Information about a specific product.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **string** | A unique identifier for the product.  | [optional] [readonly]
**CurrencyCode**| **string** | The currency code  | [optional] [readonly]
**Title**| **string** | The title of a product.  | [optional]
**Handle**| **string** | The handle of a product.  | [optional]
**Url**| **string** | The URL for a product.  | [optional]
**Description**| **string** | The description of a product.  | [optional]
**Type**| **string** | The type of product.  | [optional]
**Vendor**| **string** | The vendor for a product.  | [optional]
**ImageUrl**| **string** | The image URL for a product.  | [optional]
**Variants**| [**[]ECommerceProductVariant**](ECommerceProductVariant.md) | Returns up to 50 of the product&#39;s variants. To retrieve all variants use [Product Variants](https://mailchimp.com/developer/marketing/api/ecommerce-product-variants/).  | [optional]
**Images**| [**[]ECommerceProductImage**](ECommerceProductImage.md) | An array of the product&#39;s images.  | [optional]
**PublishedAtForeign**| [**time.Time**](time.Time.md) | The date and time the product was published in ISO 8601 format.  | [optional]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

