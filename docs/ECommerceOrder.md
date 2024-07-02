# # ECommerceOrder
Information about a specific order.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **string** | A unique identifier for the order.  | [optional] [readonly]
**Customer**| [**ECommerceCustomer**](ECommerceCustomer.md) |   | [optional]
**StoreId**| **string** | The unique identifier for the store.  | [optional] [readonly]
**CampaignId**| **string** | A string that uniquely identifies the campaign associated with an order.  | [optional]
**LandingSite**| **string** | The URL for the page where the buyer landed when entering the shop.  | [optional]
**FinancialStatus**| **string** | The order status. Use this parameter to trigger [Order Notifications](https://mailchimp.com/developer/marketing/docs/e-commerce/#order-notifications).  | [optional]
**FulfillmentStatus**| **string** | The fulfillment status for the order. Use this parameter to trigger [Order Notifications](https://mailchimp.com/developer/marketing/docs/e-commerce/#order-notifications).  | [optional]
**CurrencyCode**| **string** | The three-letter ISO 4217 code for the currency that the store accepts.  | [optional]
**OrderTotal**| **float32** | The order total associated with an order.  | [optional]
**OrderUrl**| **string** | The URL for the order.  | [optional]
**DiscountTotal**| **float32** | The total amount of the discounts to be applied to the price of the order.  | [optional]
**TaxTotal**| **float32** | The tax total associated with an order.  | [optional]
**ShippingTotal**| **float32** | The shipping total for the order.  | [optional]
**TrackingCode**| **string** | The Mailchimp tracking code for the order. Uses the &#39;mc_tc&#39; parameter in E-Commerce tracking URLs. for more information please, see Model/string.php  | [optional]
**ProcessedAtForeign**| [**time.Time**](time.Time.md) | The date and time the order was processed in ISO 8601 format.  | [optional]
**CancelledAtForeign**| [**time.Time**](time.Time.md) | The date and time the order was cancelled in ISO 8601 format.  | [optional]
**UpdatedAtForeign**| [**time.Time**](time.Time.md) | The date and time the order was updated in ISO 8601 format.  | [optional]
**ShippingAddress**| [**ShippingAddress**](ShippingAddress.md) |   | [optional]
**BillingAddress**| [**BillingAddress**](BillingAddress.md) |   | [optional]
**Promos**| [**[]PromosInner**](PromosInner.md) | The promo codes applied on the order  | [optional]
**Lines**| [**[]ECommerceOrderLineItem**](ECommerceOrderLineItem.md) | An array of the order&#39;s line items.  | [optional]
**Outreach**| [**Outreach**](Outreach.md) |   | [optional]
**TrackingNumber**| **string** | The tracking number associated with the order.  | [optional]
**TrackingCarrier**| **string** | The tracking carrier associated with the order.  | [optional]
**TrackingUrl**| **string** | The tracking URL associated with the order.  | [optional]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

