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
	"bytes"
	"fmt"
)

// checks if the Campaign2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Campaign2{}

// Campaign2 A summary of an individual campaign's settings and content.
type Campaign2 struct {
	Recipients *List5 `json:"recipients,omitempty"`
	Settings CampaignSettings4 `json:"settings"`
	VariateSettings *ABTestOptions1 `json:"variate_settings,omitempty"`
	Tracking *CampaignTrackingOptions1 `json:"tracking,omitempty"`
	RssOpts *RSSOptions2 `json:"rss_opts,omitempty"`
	SocialCard *CampaignSocialCard `json:"social_card,omitempty"`
}

type _Campaign2 Campaign2

// NewCampaign2 instantiates a new Campaign2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaign2(settings CampaignSettings4) *Campaign2 {
	this := Campaign2{}
	this.Settings = settings
	return &this
}

// NewCampaign2WithDefaults instantiates a new Campaign2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaign2WithDefaults() *Campaign2 {
	this := Campaign2{}
	return &this
}

// GetRecipients returns the Recipients field value if set, zero value otherwise.
func (o *Campaign2) GetRecipients() List5 {
	if o == nil || IsNil(o.Recipients) {
		var ret List5
		return ret
	}
	return *o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign2) GetRecipientsOk() (*List5, bool) {
	if o == nil || IsNil(o.Recipients) {
		return nil, false
	}
	return o.Recipients, true
}

// ContainsRecipients returns a boolean if a field has been set.
func (o *Campaign2) ContainsRecipients() bool {
	if o != nil && !IsNil(o.Recipients) {
		return true
	}

	return false
}

// SetRecipients gets a reference to the given List5 and assigns it to the Recipients field.
func (o *Campaign2) SetRecipients(v List5) {
	o.Recipients = &v
}

// GetSettings returns the Settings field value
func (o *Campaign2) GetSettings() CampaignSettings4 {
	if o == nil {
		var ret CampaignSettings4
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *Campaign2) GetSettingsOk() (*CampaignSettings4, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *Campaign2) SetSettings(v CampaignSettings4) {
	o.Settings = v
}

// GetVariateSettings returns the VariateSettings field value if set, zero value otherwise.
func (o *Campaign2) GetVariateSettings() ABTestOptions1 {
	if o == nil || IsNil(o.VariateSettings) {
		var ret ABTestOptions1
		return ret
	}
	return *o.VariateSettings
}

// GetVariateSettingsOk returns a tuple with the VariateSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign2) GetVariateSettingsOk() (*ABTestOptions1, bool) {
	if o == nil || IsNil(o.VariateSettings) {
		return nil, false
	}
	return o.VariateSettings, true
}

// ContainsVariateSettings returns a boolean if a field has been set.
func (o *Campaign2) ContainsVariateSettings() bool {
	if o != nil && !IsNil(o.VariateSettings) {
		return true
	}

	return false
}

// SetVariateSettings gets a reference to the given ABTestOptions1 and assigns it to the VariateSettings field.
func (o *Campaign2) SetVariateSettings(v ABTestOptions1) {
	o.VariateSettings = &v
}

// GetTracking returns the Tracking field value if set, zero value otherwise.
func (o *Campaign2) GetTracking() CampaignTrackingOptions1 {
	if o == nil || IsNil(o.Tracking) {
		var ret CampaignTrackingOptions1
		return ret
	}
	return *o.Tracking
}

// GetTrackingOk returns a tuple with the Tracking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign2) GetTrackingOk() (*CampaignTrackingOptions1, bool) {
	if o == nil || IsNil(o.Tracking) {
		return nil, false
	}
	return o.Tracking, true
}

// ContainsTracking returns a boolean if a field has been set.
func (o *Campaign2) ContainsTracking() bool {
	if o != nil && !IsNil(o.Tracking) {
		return true
	}

	return false
}

// SetTracking gets a reference to the given CampaignTrackingOptions1 and assigns it to the Tracking field.
func (o *Campaign2) SetTracking(v CampaignTrackingOptions1) {
	o.Tracking = &v
}

// GetRssOpts returns the RssOpts field value if set, zero value otherwise.
func (o *Campaign2) GetRssOpts() RSSOptions2 {
	if o == nil || IsNil(o.RssOpts) {
		var ret RSSOptions2
		return ret
	}
	return *o.RssOpts
}

// GetRssOptsOk returns a tuple with the RssOpts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign2) GetRssOptsOk() (*RSSOptions2, bool) {
	if o == nil || IsNil(o.RssOpts) {
		return nil, false
	}
	return o.RssOpts, true
}

// ContainsRssOpts returns a boolean if a field has been set.
func (o *Campaign2) ContainsRssOpts() bool {
	if o != nil && !IsNil(o.RssOpts) {
		return true
	}

	return false
}

// SetRssOpts gets a reference to the given RSSOptions2 and assigns it to the RssOpts field.
func (o *Campaign2) SetRssOpts(v RSSOptions2) {
	o.RssOpts = &v
}

// GetSocialCard returns the SocialCard field value if set, zero value otherwise.
func (o *Campaign2) GetSocialCard() CampaignSocialCard {
	if o == nil || IsNil(o.SocialCard) {
		var ret CampaignSocialCard
		return ret
	}
	return *o.SocialCard
}

// GetSocialCardOk returns a tuple with the SocialCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign2) GetSocialCardOk() (*CampaignSocialCard, bool) {
	if o == nil || IsNil(o.SocialCard) {
		return nil, false
	}
	return o.SocialCard, true
}

// ContainsSocialCard returns a boolean if a field has been set.
func (o *Campaign2) ContainsSocialCard() bool {
	if o != nil && !IsNil(o.SocialCard) {
		return true
	}

	return false
}

// SetSocialCard gets a reference to the given CampaignSocialCard and assigns it to the SocialCard field.
func (o *Campaign2) SetSocialCard(v CampaignSocialCard) {
	o.SocialCard = &v
}

func (o Campaign2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Campaign2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Recipients) {
		toSerialize["recipients"] = o.Recipients
	}
	toSerialize["settings"] = o.Settings
	if !IsNil(o.VariateSettings) {
		toSerialize["variate_settings"] = o.VariateSettings
	}
	if !IsNil(o.Tracking) {
		toSerialize["tracking"] = o.Tracking
	}
	if !IsNil(o.RssOpts) {
		toSerialize["rss_opts"] = o.RssOpts
	}
	if !IsNil(o.SocialCard) {
		toSerialize["social_card"] = o.SocialCard
	}
	return toSerialize, nil
}

func (o *Campaign2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"settings",
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

	varCampaign2 := _Campaign2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCampaign2)

	if err != nil {
		return err
	}

	*o = Campaign2(varCampaign2)

	return err
}

type NullableCampaign2 struct {
	value *Campaign2
	isSet bool
}

func (v NullableCampaign2) Get() *Campaign2 {
	return v.value
}

func (v *NullableCampaign2) Set(val *Campaign2) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaign2) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaign2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaign2(val *Campaign2) *NullableCampaign2 {
	return &NullableCampaign2{value: val, isSet: true}
}

func (v NullableCampaign2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaign2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


