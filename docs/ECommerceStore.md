# # ECommerceStore
An individual store in an account.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **string** | The unique identifier for the store.  | [optional] [readonly]
**ListId**| **string** | The unique identifier for the list that&#39;s associated with the store. The &#x60;list_id&#x60; for a specific store can&#39;t change.  | [optional] [readonly]
**Name**| **string** | The name of the store.  | [optional]
**Platform**| **string** | The e-commerce platform of the store.  | [optional]
**Domain**| **string** | The store domain.  The store domain must be unique within a user account.  | [optional]
**IsSyncing**| **bool** | Whether to disable automations because the store is currently [syncing](https://mailchimp.com/developer/marketing/docs/e-commerce/#pausing-store-automations).  | [optional]
**EmailAddress**| **string** | The email address for the store.  | [optional]
**CurrencyCode**| **string** | The three-letter ISO 4217 code for the currency that the store accepts.  | [optional]
**MoneyFormat**| **string** | The currency format for the store. For example: &#x60;$&#x60;, &#x60;£&#x60;, etc.  | [optional]
**PrimaryLocale**| **string** | The primary locale for the store. For example: &#x60;en&#x60;, &#x60;de&#x60;, etc.  | [optional]
**Timezone**| **string** | The timezone for the store.  | [optional]
**Phone**| **string** | The store phone number.  | [optional]
**Address**| [**Address1**](Address1.md) |   | [optional]
**ConnectedSite**| [**ConnectedSite2**](ConnectedSite2.md) |   | [optional]
**Automations**| [**Automations**](Automations.md) |   | [optional]
**ListIsActive**| **bool** | The status of the list connected to the store, namely if it&#39;s deleted or disabled.  | [optional] [readonly]
**CreatedAt**| [**time.Time**](time.Time.md) | The date and time the store was created in ISO 8601 format.  | [optional] [readonly]
**UpdatedAt**| [**time.Time**](time.Time.md) | The date and time the store was last updated in ISO 8601 format.  | [optional] [readonly]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

