# # ECommerceCart2
Information about a specific cart.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customer**| [**ECommerceCustomer2**](ECommerceCustomer2.md) |   | [optional]
**CampaignId**| **string** | A string that uniquely identifies the campaign associated with a cart.  | [optional]
**CheckoutUrl**| **string** | The URL for the cart. This parameter is required for [Abandoned Cart](https://mailchimp.com/help/create-an-abandoned-cart-email/) automations.  | [optional]
**CurrencyCode**| **string** | The three-letter ISO 4217 code for the currency that the cart uses.  | [optional]
**OrderTotal**| **float32** | The order total for the cart.  | [optional]
**TaxTotal**| **float32** | The total tax for the cart.  | [optional]
**Lines**| [**[]ECommerceCartLineItem2**](ECommerceCartLineItem2.md) | An array of the cart&#39;s line items.  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

