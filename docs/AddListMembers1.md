# # AddListMembers1
Individuals who are currently or have been previously subscribed to this list, including members who have bounced or unsubscribed.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress**| **string** | Email address for a subscriber.  |
**EmailType**| **string** | Type of email this member asked to get (&#39;html&#39; or &#39;text&#39;).  | [optional]
**Status**| **string** | Subscriber&#39;s current status. for more information please, see Model/string.php  |
**MergeFields**| **map[string]map[string]interface{}** | A dictionary of merge fields where the keys are the merge tags. See the [Merge Fields documentation](https://mailchimp.com/developer/marketing/docs/merge-fields/#structure) for more about the structure.  | [optional]
**Interests**| **map[string]bool** | The key of this object&#39;s properties is the ID of the interest in question.  | [optional]
**Language**| **string** | If set/detected, the [subscriber&#39;s language](https://mailchimp.com/help/view-and-edit-contact-languages/).  | [optional]
**Vip**| **bool** | [VIP status](https://mailchimp.com/help/designate-and-send-to-vip-contacts/) for subscriber.  | [optional]
**Location**| [**Location**](Location.md) |   | [optional]
**MarketingPermissions**| [**[]MarketingPermission1**](MarketingPermission1.md) | The marketing permissions for the subscriber.  | [optional]
**IpSignup**| **string** | IP address the subscriber signed up from.  | [optional]
**TimestampSignup**| [**time.Time**](time.Time.md) | The date and time the subscriber signed up for the list in ISO 8601 format.  | [optional]
**IpOpt**| **string** | The IP address the subscriber used to confirm their opt-in status.  | [optional]
**TimestampOpt**| [**time.Time**](time.Time.md) | The date and time the subscriber confirmed their opt-in status in ISO 8601 format.  | [optional]
**Tags**| **[]string** | The tags that are associated with a member.  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

