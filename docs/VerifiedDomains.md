# # VerifiedDomains
The verified domains currently on the account.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain**| **string** | The name of this domain.  | [optional] [readonly]
**Verified**| **bool** | Whether the domain has been verified for sending.  | [optional] [readonly]
**Authenticated**| **bool** | Whether domain authentication is enabled for this domain.  | [optional] [readonly]
**VerificationEmail**| **string** | The e-mail address receiving the two-factor challenge for this domain.  | [optional] [readonly]
**VerificationSent**| [**time.Time**](time.Time.md) | The date/time that the two-factor challenge was sent to the verification email.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

