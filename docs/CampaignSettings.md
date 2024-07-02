# # CampaignSettings
Settings for the campaign including the email subject, from name, and from email address.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubjectLine**| **string** | The subject line for the campaign.  | [optional]
**PreviewText**| **string** | The preview text for the campaign.  | [optional]
**Title**| **string** | The title of the campaign.  | [optional]
**FromName**| **string** | The &#39;from&#39; name on the campaign (not an email address).  | [optional]
**ReplyTo**| **string** | The reply-to email address for the campaign.  | [optional]
**Authenticate**| **bool** | Whether Mailchimp [authenticated](https://mailchimp.com/help/about-email-authentication/) the campaign. Defaults to &#x60;true&#x60;.  | [optional]
**AutoFooter**| **bool** | Automatically append Mailchimp&#39;s [default footer](https://mailchimp.com/help/about-campaign-footers/) to the campaign.  | [optional]
**InlineCss**| **bool** | Automatically inline the CSS included with the campaign content.  | [optional]
**AutoTweet**| **bool** | Automatically tweet a link to the [campaign archive](https://mailchimp.com/help/about-email-campaign-archives-and-pages/) page when the campaign is sent.  | [optional]
**AutoFbPost**| **[]string** | An array of [Facebook](https://mailchimp.com/help/connect-or-disconnect-the-facebook-integration/) page ids to auto-post to.  | [optional]
**FbComments**| **bool** | Allows Facebook comments on the campaign (also force-enables the Campaign Archive toolbar). Defaults to &#x60;true&#x60;.  | [optional]
**TemplateId**| **int32** | The id for the template used in this campaign.  | [optional]
**DragAndDrop**| **bool** | Whether the campaign uses the drag-and-drop editor.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

