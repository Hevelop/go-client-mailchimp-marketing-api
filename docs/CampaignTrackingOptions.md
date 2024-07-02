# # CampaignTrackingOptions
The tracking options for a campaign.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Opens**| **bool** | Whether to [track opens](https://mailchimp.com/help/about-open-tracking/). Defaults to &#x60;true&#x60;.  | [optional]
**HtmlClicks**| **bool** | Whether to [track clicks](https://mailchimp.com/help/enable-and-view-click-tracking/) in the HTML version of the campaign. Defaults to &#x60;true&#x60;.  | [optional]
**TextClicks**| **bool** | Whether to [track clicks](https://mailchimp.com/help/enable-and-view-click-tracking/) in the plain-text version of the campaign. Defaults to &#x60;true&#x60;.  | [optional]
**GoalTracking**| **bool** | Deprecated  | [optional]
**Ecomm360**| **bool** | Whether to enable e-commerce tracking.  | [optional]
**GoogleAnalytics**| **string** | The custom slug for [Google Analytics](https://mailchimp.com/help/integrate-google-analytics-with-mailchimp/) tracking (max of 50 bytes).  | [optional]
**Clicktale**| **string** | The custom slug for [Click Tale](https://mailchimp.com/help/additional-tracking-options-for-campaigns/) tracking (max of 50 bytes).  | [optional]
**Salesforce**| [**SalesforceCRMTracking**](SalesforceCRMTracking.md) |   | [optional]
**Capsule**| [**CapsuleCRMTracking1**](CapsuleCRMTracking1.md) |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

