# # ECommerceCart
Information about a specific cart.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **string** | A unique identifier for the cart.  | [optional] [readonly]
**Customer**| [**ECommerceCustomer**](ECommerceCustomer.md) |   | [optional]
**CampaignId**| **string** | A string that uniquely identifies the campaign associated with a cart.  | [optional]
**CheckoutUrl**| **string** | The URL for the cart. This parameter is required for [Abandoned Cart](https://mailchimp.com/help/create-an-abandoned-cart-email/) automations.  | [optional]
**CurrencyCode**| **string** | The three-letter ISO 4217 code for the currency that the cart uses.  | [optional]
**OrderTotal**| **float32** | The order total for the cart.  | [optional]
**TaxTotal**| **float32** | The total tax for the cart.  | [optional]
**Lines**| [**[]ECommerceCartLineItem**](ECommerceCartLineItem.md) | An array of the cart&#39;s line items.  | [optional]
**CreatedAt**| [**time.Time**](time.Time.md) | The date and time the cart was created in ISO 8601 format.  | [optional] [readonly]
**UpdatedAt**| [**time.Time**](time.Time.md) | The date and time the cart was last updated in ISO 8601 format.  | [optional] [readonly]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

