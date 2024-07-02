# # ECommercePromoCode2
Information about an Ecommerce Store&#39;s specific Promo Code.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code**| **string** | The discount code. Restricted to UTF-8 characters with max length 50.  | [optional]
**RedemptionUrl**| **string** | The url that should be used in the promotion campaign restricted to UTF-8 characters with max length 2000.  | [optional]
**UsageCount**| **int32** | Number of times promo code has been used.  | [optional]
**Enabled**| **bool** | Whether the promo code is currently enabled.  | [optional]
**CreatedAtForeign**| [**time.Time**](time.Time.md) | The date and time the promotion was created in ISO 8601 format.  | [optional]
**UpdatedAtForeign**| [**time.Time**](time.Time.md) | The date and time the promotion was updated in ISO 8601 format.  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

