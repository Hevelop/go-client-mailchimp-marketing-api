# # ECommerceCustomer
Information about a specific customer.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **string** | A unique identifier for the customer.  | [optional] [readonly]
**EmailAddress**| **string** | The customer&#39;s email address.  | [optional] [readonly]
**OptInStatus**| **bool** | The customer&#39;s opt-in status. This value will never overwrite the opt-in status of a pre-existing Mailchimp list member, but will apply to list members that are added through the e-commerce API endpoints. Customers who don&#39;t opt in to your Mailchimp list [will be added as &#x60;Transactional&#x60; members](https://mailchimp.com/developer/marketing/docs/e-commerce/#customers).  | [optional]
**Company**| **string** | The customer&#39;s company.  | [optional]
**FirstName**| **string** | The customer&#39;s first name.  | [optional]
**LastName**| **string** | The customer&#39;s last name.  | [optional]
**OrdersCount**| **int32** | The customer&#39;s total order count.  | [optional] [readonly]
**TotalSpent**| **float32** | The total amount the customer has spent.  | [optional] [readonly]
**Address**| [**Address**](Address.md) |   | [optional]
**CreatedAt**| [**time.Time**](time.Time.md) | The date and time the customer was created in ISO 8601 format.  | [optional] [readonly]
**UpdatedAt**| [**time.Time**](time.Time.md) | The date and time the customer was last updated in ISO 8601 format.  | [optional] [readonly]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

