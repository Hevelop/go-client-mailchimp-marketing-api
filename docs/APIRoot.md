# # APIRoot
The API root resource links to all other resources available in the API.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId**| **string** | The Mailchimp account ID.  | [optional] [readonly]
**LoginId**| **string** | The ID associated with the user who owns this API key. If you can login to multiple accounts, this ID will be the same for each account.  | [optional] [readonly]
**AccountName**| **string** | The name of the account.  | [optional] [readonly]
**Email**| **string** | The account email address.  | [optional] [readonly]
**FirstName**| **string** | The first name tied to the account.  | [optional] [readonly]
**LastName**| **string** | The last name tied to the account.  | [optional] [readonly]
**Username**| **string** | The username tied to the account.  | [optional] [readonly]
**AvatarUrl**| **string** | URL of the avatar for the user.  | [optional] [readonly]
**Role**| **string** | The [user role](https://mailchimp.com/help/manage-user-levels-in-your-account/) for the account.  | [optional] [readonly]
**MemberSince**| [**time.Time**](time.Time.md) | The date and time that the account was created in ISO 8601 format.  | [optional] [readonly]
**PricingPlanType**| **string** | The type of pricing plan the account is on. for more information please, see Model/string.php  | [optional] [readonly]
**FirstPayment**| [**time.Time**](time.Time.md) | Date of first payment for monthly plans.  | [optional] [readonly]
**AccountTimezone**| **string** | The timezone currently set for the account.  | [optional] [readonly]
**AccountIndustry**| **string** | The user-specified industry associated with the account.  | [optional] [readonly]
**Contact**| [**AccountContact**](AccountContact.md) |   | [optional]
**ProEnabled**| **bool** | Legacy - whether the account includes [Mailchimp Pro](https://mailchimp.com/help/about-mailchimp-pro/).  | [optional] [readonly]
**LastLogin**| [**time.Time**](time.Time.md) | The date and time of the last login for this account in ISO 8601 format.  | [optional] [readonly]
**TotalSubscribers**| **int32** | The total number of subscribers across all lists in the account.  | [optional] [readonly]
**IndustryStats**| [**IndustryStats**](IndustryStats.md) |   | [optional]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

