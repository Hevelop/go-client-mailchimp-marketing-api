# # ListMembers1
Individuals who are currently or have been previously subscribed to this list, including members who have bounced or unsubscribed.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **string** | The MD5 hash of the lowercase version of the list member&#39;s email address.  | [optional] [readonly]
**EmailAddress**| **string** | Email address for a subscriber.  | [optional] [readonly]
**FullName**| **string** | The contact&#39;s full name.  | [optional] [readonly]
**UniqueEmailId**| **string** | An identifier for the address across all of Mailchimp.  | [optional] [readonly]
**EmailType**| **string** | Type of email this member asked to get (&#39;html&#39; or &#39;text&#39;).  | [optional]
**Status**| **string** | Subscriber&#39;s current status. for more information please, see Model/string.php  | [optional]
**MergeFields**| **map[string]map[string]interface{}** | A dictionary of merge fields where the keys are the merge tags. See the [Merge Fields documentation](https://mailchimp.com/developer/marketing/docs/merge-fields/#structure) for more about the structure.  | [optional]
**Interests**| **map[string]bool** | The key of this object&#39;s properties is the ID of the interest in question.  | [optional]
**Stats**| [**SubscriberStats**](SubscriberStats.md) |   | [optional]
**IpSignup**| **string** | IP address the subscriber signed up from.  | [optional] [readonly]
**TimestampSignup**| [**time.Time**](time.Time.md) | The date and time the subscriber signed up for the list in ISO 8601 format.  | [optional] [readonly]
**IpOpt**| **string** | The IP address the subscriber used to confirm their opt-in status.  | [optional] [readonly]
**TimestampOpt**| [**time.Time**](time.Time.md) | The date and time the subscriber confirmed their opt-in status in ISO 8601 format.  | [optional] [readonly]
**MemberRating**| **int32** | Star rating for this member, between 1 and 5.  | [optional] [readonly]
**LastChanged**| [**time.Time**](time.Time.md) | The date and time the member&#39;s info was last changed in ISO 8601 format.  | [optional] [readonly]
**Language**| **string** | If set/detected, the [subscriber&#39;s language](https://mailchimp.com/help/view-and-edit-contact-languages/).  | [optional]
**Vip**| **bool** | [VIP status](https://mailchimp.com/help/designate-and-send-to-vip-contacts/) for subscriber.  | [optional]
**EmailClient**| **string** | The list member&#39;s email client.  | [optional] [readonly]
**Location**| [**Location1**](Location1.md) |   | [optional]
**LastNote**| [**Notes**](Notes.md) |   | [optional]
**ListId**| **string** | The list id.  | [optional] [readonly]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

