# # SubscriberList
Information about a specific list.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **string** | A string that uniquely identifies this list.  | [optional] [readonly]
**WebId**| **int32** | The ID used in the Mailchimp web application. View this list in your Mailchimp account at &#x60;https://{dc}.admin.mailchimp.com/lists/members/?id&#x3D;{web_id}&#x60;.  | [optional] [readonly]
**Name**| **string** | The name of the list.  | [optional]
**Contact**| [**ListContact**](ListContact.md) |   | [optional]
**PermissionReminder**| **string** | The [permission reminder](https://mailchimp.com/help/edit-the-permission-reminder/) for the list.  | [optional]
**UseArchiveBar**| **bool** | Whether campaigns for this list use the [Archive Bar](https://mailchimp.com/help/about-email-campaign-archives-and-pages/) in archives by default.  | [optional] [default to false]
**CampaignDefaults**| [**CampaignDefaults**](CampaignDefaults.md) |   | [optional]
**NotifyOnSubscribe**| **string** | The email address to send [subscribe notifications](https://mailchimp.com/help/change-subscribe-and-unsubscribe-notifications/) to.  | [optional] [default to "false"]
**NotifyOnUnsubscribe**| **string** | The email address to send [unsubscribe notifications](https://mailchimp.com/help/change-subscribe-and-unsubscribe-notifications/) to.  | [optional] [default to "false"]
**DateCreated**| [**time.Time**](time.Time.md) | The date and time that this list was created in ISO 8601 format.  | [optional] [readonly]
**ListRating**| **int32** | An auto-generated activity score for the list (0-5).  | [optional] [readonly]
**EmailTypeOption**| **bool** | Whether the list supports [multiple formats for emails](https://mailchimp.com/help/change-audience-name-defaults/). When set to &#x60;true&#x60;, subscribers can choose whether they want to receive HTML or plain-text emails. When set to &#x60;false&#x60;, subscribers will receive HTML emails, with a plain-text alternative backup.  | [optional]
**SubscribeUrlShort**| **string** | Our [url shortened](https://mailchimp.com/help/share-your-signup-form/) version of this list&#39;s subscribe form.  | [optional] [readonly]
**SubscribeUrlLong**| **string** | The full version of this list&#39;s subscribe form (host will vary).  | [optional] [readonly]
**BeamerAddress**| **string** | The list&#39;s [Email Beamer](https://mailchimp.com/help/use-email-beamer-to-create-a-campaign/) address.  | [optional] [readonly]
**Visibility**| **string** | Legacy - visibility settings are no longer used for more information please, see Model/string.php  | [optional]
**DoubleOptin**| **bool** | Whether or not to require the subscriber to confirm subscription via email.  | [optional] [default to false]
**HasWelcome**| **bool** | Whether or not this list has a welcome automation connected. Welcome Automations: welcomeSeries, singleWelcome, emailFollowup.  | [optional] [default to false]
**MarketingPermissions**| **bool** | Whether or not the list has marketing permissions (eg. GDPR) enabled.  | [optional] [default to false]
**Modules**| **[]string** | Any list-specific modules installed for this list.  | [optional] [readonly]
**Stats**| [**Statistics**](Statistics.md) |   | [optional]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

