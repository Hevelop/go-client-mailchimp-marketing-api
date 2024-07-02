# # ECommerceStore1
An individual store in an account.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **string** | The unique identifier for the store.  |
**ListId**| **string** | The unique identifier for the list associated with the store. The &#x60;list_id&#x60; for a specific store cannot change.  |
**Name**| **string** | The name of the store.  |
**Platform**| **string** | The e-commerce platform of the store.  | [optional]
**Domain**| **string** | The store domain. This parameter is required for Connected Sites and Google Ads.  | [optional]
**IsSyncing**| **bool** | Whether to disable automations because the store is currently [syncing](https://mailchimp.com/developer/marketing/docs/e-commerce/#pausing-store-automations).  | [optional]
**EmailAddress**| **string** | The email address for the store.  | [optional]
**CurrencyCode**| **string** | The three-letter ISO 4217 code for the currency that the store accepts.  |
**MoneyFormat**| **string** | The currency format for the store. For example: &#x60;$&#x60;, &#x60;£&#x60;, etc.  | [optional]
**PrimaryLocale**| **string** | The primary locale for the store. For example: &#x60;en&#x60;, &#x60;de&#x60;, etc.  | [optional]
**Timezone**| **string** | The timezone for the store.  | [optional]
**Phone**| **string** | The store phone number.  | [optional]
**Address**| [**Address1**](Address1.md) |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

