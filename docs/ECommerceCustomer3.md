# # ECommerceCustomer3
Information about a specific customer.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **string** | A unique identifier for the customer. Limited to 50 characters.  |
**EmailAddress**| **string** | The customer&#39;s email address.  |
**OptInStatus**| **bool** | The customer&#39;s opt-in status. This value will never overwrite the opt-in status of a pre-existing Mailchimp list member, but will apply to list members that are added through the e-commerce API endpoints. Customers who don&#39;t opt in to your Mailchimp list [will be added as &#x60;Transactional&#x60; members](https://mailchimp.com/developer/marketing/docs/e-commerce/#customers).  |
**Company**| **string** | The customer&#39;s company.  | [optional]
**FirstName**| **string** | The customer&#39;s first name.  | [optional]
**LastName**| **string** | The customer&#39;s last name.  | [optional]
**Address**| [**Address**](Address.md) |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

