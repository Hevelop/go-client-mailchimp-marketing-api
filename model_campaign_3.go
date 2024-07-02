/*
Mailchimp Marketing API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.0.55
Contact: apihelp@mailchimp.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mailchimpmarketingapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the Campaign3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Campaign3{}

// Campaign3 A summary of an individual campaign's settings and content.
type Campaign3 struct {
	// A string that uniquely identifies this campaign.
	Id *string `json:"id,omitempty"`
	// The ID used in the Mailchimp web application. View this campaign in your Mailchimp account at `https://{dc}.admin.mailchimp.com/campaigns/show/?id={web_id}`.
	WebId *int32 `json:"web_id,omitempty"`
	// If this campaign is the child of another campaign, this identifies the parent campaign. For Example, for RSS or Automation children.
	ParentCampaignId *string `json:"parent_campaign_id,omitempty"`
	// There are four types of [campaigns](https://mailchimp.com/help/getting-started-with-campaigns/) you can create in Mailchimp. A/B Split campaigns have been deprecated and variate campaigns should be used instead.
	Type string `json:"type"`
	// The date and time the campaign was created in ISO 8601 format.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// The link to the campaign's archive version.
	ArchiveUrl *string `json:"archive_url,omitempty"`
	// The original link to the campaign's archive version.
	LongArchiveUrl *string `json:"long_archive_url,omitempty"`
	// The current status of the campaign.
	Status *string `json:"status,omitempty"`
	// The total number of emails sent for this campaign.
	EmailsSent *int32 `json:"emails_sent,omitempty"`
	// The date and time a campaign was sent in ISO 8601 format.
	SendTime *time.Time `json:"send_time,omitempty"`
	// How the campaign's content is put together ('template', 'drag_and_drop', 'html', 'url').
	ContentType *string `json:"content_type,omitempty"`
	// Determines if the campaign needs its blocks refreshed by opening the web-based campaign editor. Deprecated and will always return false.
	NeedsBlockRefresh *bool `json:"needs_block_refresh,omitempty"`
	// Determines if the campaign qualifies to be resent to non-openers.
	Resendable *bool `json:"resendable,omitempty"`
	Recipients *List6 `json:"recipients,omitempty"`
	Settings *CampaignSettings5 `json:"settings,omitempty"`
	VariateSettings *ABTestOptions2 `json:"variate_settings,omitempty"`
	Tracking *CampaignTrackingOptions1 `json:"tracking,omitempty"`
	RssOpts *RSSOptions3 `json:"rss_opts,omitempty"`
	AbSplitOpts *ABTestingOptions `json:"ab_split_opts,omitempty"`
	SocialCard *CampaignSocialCard `json:"social_card,omitempty"`
	ReportSummary *CampaignReportSummary3 `json:"report_summary,omitempty"`
	DeliveryStatus *CampaignDeliveryStatus `json:"delivery_status,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

type _Campaign3 Campaign3

// NewCampaign3 instantiates a new Campaign3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaign3(type_ string) *Campaign3 {
	this := Campaign3{}
	this.Type = type_
	return &this
}

// NewCampaign3WithDefaults instantiates a new Campaign3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaign3WithDefaults() *Campaign3 {
	this := Campaign3{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Campaign3) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign3) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Campaign3) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Campaign3) SetId(v string) {
	o.Id = &v
}

// GetWebId returns the WebId field value if set, zero value otherwise.
func (o *Campaign3) GetWebId() int32 {
	if o == nil || IsNil(o.WebId) {
		var ret int32
		return ret
	}
	return *o.WebId
}

// GetWebIdOk returns a tuple with the WebId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign3) GetWebIdOk() (*int32, bool) {
	if o == nil || IsNil(o.WebId) {
		return nil, false
	}
	return o.WebId, true
}

// HasWebId returns a boolean if a field has been set.
func (o *Campaign3) HasWebId() bool {
	if o != nil && !IsNil(o.WebId) {
		return true
	}

	return false
}

// SetWebId gets a reference to the given int32 and assigns it to the WebId field.
func (o *Campaign3) SetWebId(v int32) {
	o.WebId = &v
}

// GetParentCampaignId returns the ParentCampaignId field value if set, zero value otherwise.
func (o *Campaign3) GetParentCampaignId() string {
	if o == nil || IsNil(o.ParentCampaignId) {
		var ret string
		return ret
	}
	return *o.ParentCampaignId
}

// GetParentCampaignIdOk returns a tuple with the ParentCampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign3) GetParentCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentCampaignId) {
		return nil, false
	}
	return o.ParentCampaignId, true
}

// HasParentCampaignId returns a boolean if a field has been set.
func (o *Campaign3) HasParentCampaignId() bool {
	if o != nil && !IsNil(o.ParentCampaignId) {
		return true
	}

	return false
}

// SetParentCampaignId gets a reference to the given string and assigns it to the ParentCampaignId field.
func (o *Campaign3) SetParentCampaignId(v string) {
	o.ParentCampaignId = &v
}

// GetType returns the Type field value
func (o *Campaign3) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Campaign3) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Campaign3) SetType(v string) {
	o.Type = v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *Campaign3) GetCreateTime() time.Time {
	if o == nil || IsNil(o.CreateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign3) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *Campaign3) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *Campaign3) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetArchiveUrl returns the ArchiveUrl field value if set, zero value otherwise.
func (o *Campaign3) GetArchiveUrl() string {
	if o == nil || IsNil(o.ArchiveUrl) {
		var ret string
		return ret
	}
	return *o.ArchiveUrl
}

// GetArchiveUrlOk returns a tuple with the ArchiveUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign3) GetArchiveUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ArchiveUrl) {
		return nil, false
	}
	return o.ArchiveUrl, true
}

// HasArchiveUrl returns a boolean if a field has been set.
func (o *Campaign3) HasArchiveUrl() bool {
	if o != nil && !IsNil(o.ArchiveUrl) {
		return true
	}

	return false
}

// SetArchiveUrl gets a reference to the given string and assigns it to the ArchiveUrl field.
func (o *Campaign3) SetArchiveUrl(v string) {
	o.ArchiveUrl = &v
}

// GetLongArchiveUrl returns the LongArchiveUrl field value if set, zero value otherwise.
func (o *Campaign3) GetLongArchiveUrl() string {
	if o == nil || IsNil(o.LongArchiveUrl) {
		var ret string
		return ret
	}
	return *o.LongArchiveUrl
}

// GetLongArchiveUrlOk returns a tuple with the LongArchiveUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign3) GetLongArchiveUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LongArchiveUrl) {
		return nil, false
	}
	return o.LongArchiveUrl, true
}

// HasLongArchiveUrl returns a boolean if a field has been set.
func (o *Campaign3) HasLongArchiveUrl() bool {
	if o != nil && !IsNil(o.LongArchiveUrl) {
		return true
	}

	return false
}

// SetLongArchiveUrl gets a reference to the given string and assigns it to the LongArchiveUrl field.
func (o *Campaign3) SetLongArchiveUrl(v string) {
	o.LongArchiveUrl = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Campaign3) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign3) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Campaign3) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Campaign3) SetStatus(v string) {
	o.Status = &v
}

// GetEmailsSent returns the EmailsSent field value if set, zero value otherwise.
func (o *Campaign3) GetEmailsSent() int32 {
	if o == nil || IsNil(o.EmailsSent) {
		var ret int32
		return ret
	}
	return *o.EmailsSent
}

// GetEmailsSentOk returns a tuple with the EmailsSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign3) GetEmailsSentOk() (*int32, bool) {
	if o == nil || IsNil(o.EmailsSent) {
		return nil, false
	}
	return o.EmailsSent, true
}

// HasEmailsSent returns a boolean if a field has been set.
func (o *Campaign3) HasEmailsSent() bool {
	if o != nil && !IsNil(o.EmailsSent) {
		return true
	}

	return false
}

// SetEmailsSent gets a reference to the given int32 and assigns it to the EmailsSent field.
func (o *Campaign3) SetEmailsSent(v int32) {
	o.EmailsSent = &v
}

// GetSendTime returns the SendTime field value if set, zero value otherwise.
func (o *Campaign3) GetSendTime() time.Time {
	if o == nil || IsNil(o.SendTime) {
		var ret time.Time
		return ret
	}
	return *o.SendTime
}

// GetSendTimeOk returns a tuple with the SendTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign3) GetSendTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SendTime) {
		return nil, false
	}
	return o.SendTime, true
}

// HasSendTime returns a boolean if a field has been set.
func (o *Campaign3) HasSendTime() bool {
	if o != nil && !IsNil(o.SendTime) {
		return true
	}

	return false
}

// SetSendTime gets a reference to the given time.Time and assigns it to the SendTime field.
func (o *Campaign3) SetSendTime(v time.Time) {
	o.SendTime = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *Campaign3) GetContentType() string {
	if o == nil || IsNil(o.ContentType) {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign3) GetContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *Campaign3) HasContentType() bool {
	if o != nil && !IsNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *Campaign3) SetContentType(v string) {
	o.ContentType = &v
}

// GetNeedsBlockRefresh returns the NeedsBlockRefresh field value if set, zero value otherwise.
func (o *Campaign3) GetNeedsBlockRefresh() bool {
	if o == nil || IsNil(o.NeedsBlockRefresh) {
		var ret bool
		return ret
	}
	return *o.NeedsBlockRefresh
}

// GetNeedsBlockRefreshOk returns a tuple with the NeedsBlockRefresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign3) GetNeedsBlockRefreshOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedsBlockRefresh) {
		return nil, false
	}
	return o.NeedsBlockRefresh, true
}

// HasNeedsBlockRefresh returns a boolean if a field has been set.
func (o *Campaign3) HasNeedsBlockRefresh() bool {
	if o != nil && !IsNil(o.NeedsBlockRefresh) {
		return true
	}

	return false
}

// SetNeedsBlockRefresh gets a reference to the given bool and assigns it to the NeedsBlockRefresh field.
func (o *Campaign3) SetNeedsBlockRefresh(v bool) {
	o.NeedsBlockRefresh = &v
}

// GetResendable returns the Resendable field value if set, zero value otherwise.
func (o *Campaign3) GetResendable() bool {
	if o == nil || IsNil(o.Resendable) {
		var ret bool
		return ret
	}
	return *o.Resendable
}

// GetResendableOk returns a tuple with the Resendable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign3) GetResendableOk() (*bool, bool) {
	if o == nil || IsNil(o.Resendable) {
		return nil, false
	}
	return o.Resendable, true
}

// HasResendable returns a boolean if a field has been set.
func (o *Campaign3) HasResendable() bool {
	if o != nil && !IsNil(o.Resendable) {
		return true
	}

	return false
}

// SetResendable gets a reference to the given bool and assigns it to the Resendable field.
func (o *Campaign3) SetResendable(v bool) {
	o.Resendable = &v
}

// GetRecipients returns the Recipients field value if set, zero value otherwise.
func (o *Campaign3) GetRecipients() List6 {
	if o == nil || IsNil(o.Recipients) {
		var ret List6
		return ret
	}
	return *o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign3) GetRecipientsOk() (*List6, bool) {
	if o == nil || IsNil(o.Recipients) {
		return nil, false
	}
	return o.Recipients, true
}

// HasRecipients returns a boolean if a field has been set.
func (o *Campaign3) HasRecipients() bool {
	if o != nil && !IsNil(o.Recipients) {
		return true
	}

	return false
}

// SetRecipients gets a reference to the given List6 and assigns it to the Recipients field.
func (o *Campaign3) SetRecipients(v List6) {
	o.Recipients = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *Campaign3) GetSettings() CampaignSettings5 {
	if o == nil || IsNil(o.Settings) {
		var ret CampaignSettings5
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign3) GetSettingsOk() (*CampaignSettings5, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *Campaign3) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given CampaignSettings5 and assigns it to the Settings field.
func (o *Campaign3) SetSettings(v CampaignSettings5) {
	o.Settings = &v
}

// GetVariateSettings returns the VariateSettings field value if set, zero value otherwise.
func (o *Campaign3) GetVariateSettings() ABTestOptions2 {
	if o == nil || IsNil(o.VariateSettings) {
		var ret ABTestOptions2
		return ret
	}
	return *o.VariateSettings
}

// GetVariateSettingsOk returns a tuple with the VariateSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign3) GetVariateSettingsOk() (*ABTestOptions2, bool) {
	if o == nil || IsNil(o.VariateSettings) {
		return nil, false
	}
	return o.VariateSettings, true
}

// HasVariateSettings returns a boolean if a field has been set.
func (o *Campaign3) HasVariateSettings() bool {
	if o != nil && !IsNil(o.VariateSettings) {
		return true
	}

	return false
}

// SetVariateSettings gets a reference to the given ABTestOptions2 and assigns it to the VariateSettings field.
func (o *Campaign3) SetVariateSettings(v ABTestOptions2) {
	o.VariateSettings = &v
}

// GetTracking returns the Tracking field value if set, zero value otherwise.
func (o *Campaign3) GetTracking() CampaignTrackingOptions1 {
	if o == nil || IsNil(o.Tracking) {
		var ret CampaignTrackingOptions1
		return ret
	}
	return *o.Tracking
}

// GetTrackingOk returns a tuple with the Tracking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign3) GetTrackingOk() (*CampaignTrackingOptions1, bool) {
	if o == nil || IsNil(o.Tracking) {
		return nil, false
	}
	return o.Tracking, true
}

// HasTracking returns a boolean if a field has been set.
func (o *Campaign3) HasTracking() bool {
	if o != nil && !IsNil(o.Tracking) {
		return true
	}

	return false
}

// SetTracking gets a reference to the given CampaignTrackingOptions1 and assigns it to the Tracking field.
func (o *Campaign3) SetTracking(v CampaignTrackingOptions1) {
	o.Tracking = &v
}

// GetRssOpts returns the RssOpts field value if set, zero value otherwise.
func (o *Campaign3) GetRssOpts() RSSOptions3 {
	if o == nil || IsNil(o.RssOpts) {
		var ret RSSOptions3
		return ret
	}
	return *o.RssOpts
}

// GetRssOptsOk returns a tuple with the RssOpts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign3) GetRssOptsOk() (*RSSOptions3, bool) {
	if o == nil || IsNil(o.RssOpts) {
		return nil, false
	}
	return o.RssOpts, true
}

// HasRssOpts returns a boolean if a field has been set.
func (o *Campaign3) HasRssOpts() bool {
	if o != nil && !IsNil(o.RssOpts) {
		return true
	}

	return false
}

// SetRssOpts gets a reference to the given RSSOptions3 and assigns it to the RssOpts field.
func (o *Campaign3) SetRssOpts(v RSSOptions3) {
	o.RssOpts = &v
}

// GetAbSplitOpts returns the AbSplitOpts field value if set, zero value otherwise.
func (o *Campaign3) GetAbSplitOpts() ABTestingOptions {
	if o == nil || IsNil(o.AbSplitOpts) {
		var ret ABTestingOptions
		return ret
	}
	return *o.AbSplitOpts
}

// GetAbSplitOptsOk returns a tuple with the AbSplitOpts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign3) GetAbSplitOptsOk() (*ABTestingOptions, bool) {
	if o == nil || IsNil(o.AbSplitOpts) {
		return nil, false
	}
	return o.AbSplitOpts, true
}

// HasAbSplitOpts returns a boolean if a field has been set.
func (o *Campaign3) HasAbSplitOpts() bool {
	if o != nil && !IsNil(o.AbSplitOpts) {
		return true
	}

	return false
}

// SetAbSplitOpts gets a reference to the given ABTestingOptions and assigns it to the AbSplitOpts field.
func (o *Campaign3) SetAbSplitOpts(v ABTestingOptions) {
	o.AbSplitOpts = &v
}

// GetSocialCard returns the SocialCard field value if set, zero value otherwise.
func (o *Campaign3) GetSocialCard() CampaignSocialCard {
	if o == nil || IsNil(o.SocialCard) {
		var ret CampaignSocialCard
		return ret
	}
	return *o.SocialCard
}

// GetSocialCardOk returns a tuple with the SocialCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign3) GetSocialCardOk() (*CampaignSocialCard, bool) {
	if o == nil || IsNil(o.SocialCard) {
		return nil, false
	}
	return o.SocialCard, true
}

// HasSocialCard returns a boolean if a field has been set.
func (o *Campaign3) HasSocialCard() bool {
	if o != nil && !IsNil(o.SocialCard) {
		return true
	}

	return false
}

// SetSocialCard gets a reference to the given CampaignSocialCard and assigns it to the SocialCard field.
func (o *Campaign3) SetSocialCard(v CampaignSocialCard) {
	o.SocialCard = &v
}

// GetReportSummary returns the ReportSummary field value if set, zero value otherwise.
func (o *Campaign3) GetReportSummary() CampaignReportSummary3 {
	if o == nil || IsNil(o.ReportSummary) {
		var ret CampaignReportSummary3
		return ret
	}
	return *o.ReportSummary
}

// GetReportSummaryOk returns a tuple with the ReportSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign3) GetReportSummaryOk() (*CampaignReportSummary3, bool) {
	if o == nil || IsNil(o.ReportSummary) {
		return nil, false
	}
	return o.ReportSummary, true
}

// HasReportSummary returns a boolean if a field has been set.
func (o *Campaign3) HasReportSummary() bool {
	if o != nil && !IsNil(o.ReportSummary) {
		return true
	}

	return false
}

// SetReportSummary gets a reference to the given CampaignReportSummary3 and assigns it to the ReportSummary field.
func (o *Campaign3) SetReportSummary(v CampaignReportSummary3) {
	o.ReportSummary = &v
}

// GetDeliveryStatus returns the DeliveryStatus field value if set, zero value otherwise.
func (o *Campaign3) GetDeliveryStatus() CampaignDeliveryStatus {
	if o == nil || IsNil(o.DeliveryStatus) {
		var ret CampaignDeliveryStatus
		return ret
	}
	return *o.DeliveryStatus
}

// GetDeliveryStatusOk returns a tuple with the DeliveryStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign3) GetDeliveryStatusOk() (*CampaignDeliveryStatus, bool) {
	if o == nil || IsNil(o.DeliveryStatus) {
		return nil, false
	}
	return o.DeliveryStatus, true
}

// HasDeliveryStatus returns a boolean if a field has been set.
func (o *Campaign3) HasDeliveryStatus() bool {
	if o != nil && !IsNil(o.DeliveryStatus) {
		return true
	}

	return false
}

// SetDeliveryStatus gets a reference to the given CampaignDeliveryStatus and assigns it to the DeliveryStatus field.
func (o *Campaign3) SetDeliveryStatus(v CampaignDeliveryStatus) {
	o.DeliveryStatus = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Campaign3) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign3) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Campaign3) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *Campaign3) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o Campaign3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Campaign3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.WebId) {
		toSerialize["web_id"] = o.WebId
	}
	if !IsNil(o.ParentCampaignId) {
		toSerialize["parent_campaign_id"] = o.ParentCampaignId
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.CreateTime) {
		toSerialize["create_time"] = o.CreateTime
	}
	if !IsNil(o.ArchiveUrl) {
		toSerialize["archive_url"] = o.ArchiveUrl
	}
	if !IsNil(o.LongArchiveUrl) {
		toSerialize["long_archive_url"] = o.LongArchiveUrl
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.EmailsSent) {
		toSerialize["emails_sent"] = o.EmailsSent
	}
	if !IsNil(o.SendTime) {
		toSerialize["send_time"] = o.SendTime
	}
	if !IsNil(o.ContentType) {
		toSerialize["content_type"] = o.ContentType
	}
	if !IsNil(o.NeedsBlockRefresh) {
		toSerialize["needs_block_refresh"] = o.NeedsBlockRefresh
	}
	if !IsNil(o.Resendable) {
		toSerialize["resendable"] = o.Resendable
	}
	if !IsNil(o.Recipients) {
		toSerialize["recipients"] = o.Recipients
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !IsNil(o.VariateSettings) {
		toSerialize["variate_settings"] = o.VariateSettings
	}
	if !IsNil(o.Tracking) {
		toSerialize["tracking"] = o.Tracking
	}
	if !IsNil(o.RssOpts) {
		toSerialize["rss_opts"] = o.RssOpts
	}
	if !IsNil(o.AbSplitOpts) {
		toSerialize["ab_split_opts"] = o.AbSplitOpts
	}
	if !IsNil(o.SocialCard) {
		toSerialize["social_card"] = o.SocialCard
	}
	if !IsNil(o.ReportSummary) {
		toSerialize["report_summary"] = o.ReportSummary
	}
	if !IsNil(o.DeliveryStatus) {
		toSerialize["delivery_status"] = o.DeliveryStatus
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

func (o *Campaign3) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCampaign3 := _Campaign3{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCampaign3)

	if err != nil {
		return err
	}

	*o = Campaign3(varCampaign3)

	return err
}

type NullableCampaign3 struct {
	value *Campaign3
	isSet bool
}

func (v NullableCampaign3) Get() *Campaign3 {
	return v.value
}

func (v *NullableCampaign3) Set(val *Campaign3) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaign3) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaign3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaign3(val *Campaign3) *NullableCampaign3 {
	return &NullableCampaign3{value: val, isSet: true}
}

func (v NullableCampaign3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaign3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

