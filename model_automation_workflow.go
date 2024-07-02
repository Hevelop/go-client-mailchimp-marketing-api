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
)

// checks if the AutomationWorkflow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutomationWorkflow{}

// AutomationWorkflow A summary of an individual Automation workflow's settings and content.
type AutomationWorkflow struct {
	// A string that identifies the Automation.
	Id *string `json:"id,omitempty"`
	// The date and time the Automation was created in ISO 8601 format.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// The date and time the Automation was started in ISO 8601 format.
	StartTime *time.Time `json:"start_time,omitempty"`
	// The current status of the Automation.
	Status *string `json:"status,omitempty"`
	// The total number of emails sent for the Automation.
	EmailsSent *int32 `json:"emails_sent,omitempty"`
	Recipients *List `json:"recipients,omitempty"`
	Settings *AutomationCampaignSettings `json:"settings,omitempty"`
	Tracking *AutomationTrackingOptions `json:"tracking,omitempty"`
	TriggerSettings *AutomationTrigger `json:"trigger_settings,omitempty"`
	ReportSummary *CampaignReportSummary `json:"report_summary,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewAutomationWorkflow instantiates a new AutomationWorkflow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutomationWorkflow() *AutomationWorkflow {
	this := AutomationWorkflow{}
	return &this
}

// NewAutomationWorkflowWithDefaults instantiates a new AutomationWorkflow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutomationWorkflowWithDefaults() *AutomationWorkflow {
	this := AutomationWorkflow{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AutomationWorkflow) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationWorkflow) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AutomationWorkflow) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AutomationWorkflow) SetId(v string) {
	o.Id = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *AutomationWorkflow) GetCreateTime() time.Time {
	if o == nil || IsNil(o.CreateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationWorkflow) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *AutomationWorkflow) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *AutomationWorkflow) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *AutomationWorkflow) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationWorkflow) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *AutomationWorkflow) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *AutomationWorkflow) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AutomationWorkflow) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationWorkflow) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AutomationWorkflow) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AutomationWorkflow) SetStatus(v string) {
	o.Status = &v
}

// GetEmailsSent returns the EmailsSent field value if set, zero value otherwise.
func (o *AutomationWorkflow) GetEmailsSent() int32 {
	if o == nil || IsNil(o.EmailsSent) {
		var ret int32
		return ret
	}
	return *o.EmailsSent
}

// GetEmailsSentOk returns a tuple with the EmailsSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationWorkflow) GetEmailsSentOk() (*int32, bool) {
	if o == nil || IsNil(o.EmailsSent) {
		return nil, false
	}
	return o.EmailsSent, true
}

// HasEmailsSent returns a boolean if a field has been set.
func (o *AutomationWorkflow) HasEmailsSent() bool {
	if o != nil && !IsNil(o.EmailsSent) {
		return true
	}

	return false
}

// SetEmailsSent gets a reference to the given int32 and assigns it to the EmailsSent field.
func (o *AutomationWorkflow) SetEmailsSent(v int32) {
	o.EmailsSent = &v
}

// GetRecipients returns the Recipients field value if set, zero value otherwise.
func (o *AutomationWorkflow) GetRecipients() List {
	if o == nil || IsNil(o.Recipients) {
		var ret List
		return ret
	}
	return *o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationWorkflow) GetRecipientsOk() (*List, bool) {
	if o == nil || IsNil(o.Recipients) {
		return nil, false
	}
	return o.Recipients, true
}

// HasRecipients returns a boolean if a field has been set.
func (o *AutomationWorkflow) HasRecipients() bool {
	if o != nil && !IsNil(o.Recipients) {
		return true
	}

	return false
}

// SetRecipients gets a reference to the given List and assigns it to the Recipients field.
func (o *AutomationWorkflow) SetRecipients(v List) {
	o.Recipients = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *AutomationWorkflow) GetSettings() AutomationCampaignSettings {
	if o == nil || IsNil(o.Settings) {
		var ret AutomationCampaignSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationWorkflow) GetSettingsOk() (*AutomationCampaignSettings, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *AutomationWorkflow) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given AutomationCampaignSettings and assigns it to the Settings field.
func (o *AutomationWorkflow) SetSettings(v AutomationCampaignSettings) {
	o.Settings = &v
}

// GetTracking returns the Tracking field value if set, zero value otherwise.
func (o *AutomationWorkflow) GetTracking() AutomationTrackingOptions {
	if o == nil || IsNil(o.Tracking) {
		var ret AutomationTrackingOptions
		return ret
	}
	return *o.Tracking
}

// GetTrackingOk returns a tuple with the Tracking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationWorkflow) GetTrackingOk() (*AutomationTrackingOptions, bool) {
	if o == nil || IsNil(o.Tracking) {
		return nil, false
	}
	return o.Tracking, true
}

// HasTracking returns a boolean if a field has been set.
func (o *AutomationWorkflow) HasTracking() bool {
	if o != nil && !IsNil(o.Tracking) {
		return true
	}

	return false
}

// SetTracking gets a reference to the given AutomationTrackingOptions and assigns it to the Tracking field.
func (o *AutomationWorkflow) SetTracking(v AutomationTrackingOptions) {
	o.Tracking = &v
}

// GetTriggerSettings returns the TriggerSettings field value if set, zero value otherwise.
func (o *AutomationWorkflow) GetTriggerSettings() AutomationTrigger {
	if o == nil || IsNil(o.TriggerSettings) {
		var ret AutomationTrigger
		return ret
	}
	return *o.TriggerSettings
}

// GetTriggerSettingsOk returns a tuple with the TriggerSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationWorkflow) GetTriggerSettingsOk() (*AutomationTrigger, bool) {
	if o == nil || IsNil(o.TriggerSettings) {
		return nil, false
	}
	return o.TriggerSettings, true
}

// HasTriggerSettings returns a boolean if a field has been set.
func (o *AutomationWorkflow) HasTriggerSettings() bool {
	if o != nil && !IsNil(o.TriggerSettings) {
		return true
	}

	return false
}

// SetTriggerSettings gets a reference to the given AutomationTrigger and assigns it to the TriggerSettings field.
func (o *AutomationWorkflow) SetTriggerSettings(v AutomationTrigger) {
	o.TriggerSettings = &v
}

// GetReportSummary returns the ReportSummary field value if set, zero value otherwise.
func (o *AutomationWorkflow) GetReportSummary() CampaignReportSummary {
	if o == nil || IsNil(o.ReportSummary) {
		var ret CampaignReportSummary
		return ret
	}
	return *o.ReportSummary
}

// GetReportSummaryOk returns a tuple with the ReportSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationWorkflow) GetReportSummaryOk() (*CampaignReportSummary, bool) {
	if o == nil || IsNil(o.ReportSummary) {
		return nil, false
	}
	return o.ReportSummary, true
}

// HasReportSummary returns a boolean if a field has been set.
func (o *AutomationWorkflow) HasReportSummary() bool {
	if o != nil && !IsNil(o.ReportSummary) {
		return true
	}

	return false
}

// SetReportSummary gets a reference to the given CampaignReportSummary and assigns it to the ReportSummary field.
func (o *AutomationWorkflow) SetReportSummary(v CampaignReportSummary) {
	o.ReportSummary = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AutomationWorkflow) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationWorkflow) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AutomationWorkflow) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *AutomationWorkflow) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o AutomationWorkflow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutomationWorkflow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CreateTime) {
		toSerialize["create_time"] = o.CreateTime
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.EmailsSent) {
		toSerialize["emails_sent"] = o.EmailsSent
	}
	if !IsNil(o.Recipients) {
		toSerialize["recipients"] = o.Recipients
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !IsNil(o.Tracking) {
		toSerialize["tracking"] = o.Tracking
	}
	if !IsNil(o.TriggerSettings) {
		toSerialize["trigger_settings"] = o.TriggerSettings
	}
	if !IsNil(o.ReportSummary) {
		toSerialize["report_summary"] = o.ReportSummary
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableAutomationWorkflow struct {
	value *AutomationWorkflow
	isSet bool
}

func (v NullableAutomationWorkflow) Get() *AutomationWorkflow {
	return v.value
}

func (v *NullableAutomationWorkflow) Set(val *AutomationWorkflow) {
	v.value = val
	v.isSet = true
}

func (v NullableAutomationWorkflow) IsSet() bool {
	return v.isSet
}

func (v *NullableAutomationWorkflow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutomationWorkflow(val *AutomationWorkflow) *NullableAutomationWorkflow {
	return &NullableAutomationWorkflow{value: val, isSet: true}
}

func (v NullableAutomationWorkflow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutomationWorkflow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


