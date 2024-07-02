# # ECommercePromoRule2
Information about an Ecommerce Store&#39;s specific Promo Rule.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title**| **string** | The title that will show up in promotion campaign. Restricted to UTF-8 characters with max length of 100 bytes.  | [optional]
**Description**| **string** | The description of a promotion restricted to UTF-8 characters with max length 255.  | [optional]
**StartsAt**| [**time.Time**](time.Time.md) | The date and time when the promotion is in effect in ISO 8601 format.  | [optional]
**EndsAt**| **string** | The date and time when the promotion ends. Must be after starts_at and in ISO 8601 format.  | [optional]
**Amount**| **float32** | The amount of the promo code discount. If &#39;type&#39; is &#39;fixed&#39;, the amount is treated as a monetary value. If &#39;type&#39; is &#39;percentage&#39;, amount must be a decimal value between 0.0 and 1.0, inclusive.  | [optional]
**Type**| **string** | Type of discount. For free shipping set type to fixed. for more information please, see Model/string.php  | [optional]
**Target**| **string** | The target that the discount applies to. for more information please, see Model/string.php  | [optional]
**Enabled**| **bool** | Whether the promo rule is currently enabled.  | [optional]
**CreatedAtForeign**| [**time.Time**](time.Time.md) | The date and time the promotion was created in ISO 8601 format.  | [optional]
**UpdatedAtForeign**| [**time.Time**](time.Time.md) | The date and time the promotion was updated in ISO 8601 format.  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

