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
)

// checks if the Events1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Events1{}

// Events1 The events that can trigger the webhook and whether they are enabled.
type Events1 struct {
	// Whether the webhook is triggered when a list subscriber is added.
	Subscribe *bool `json:"subscribe,omitempty"`
	// Whether the webhook is triggered when a list member unsubscribes.
	Unsubscribe *bool `json:"unsubscribe,omitempty"`
	// Whether the webhook is triggered when a subscriber's profile is updated.
	Profile *bool `json:"profile,omitempty"`
	// Whether the webhook is triggered when a subscriber's email address is cleaned from the list.
	Cleaned *bool `json:"cleaned,omitempty"`
	// Whether the webhook is triggered when a subscriber's email address is changed.
	Upemail *bool `json:"upemail,omitempty"`
	// Whether the webhook is triggered when a campaign is sent or cancelled.
	Campaign *bool `json:"campaign,omitempty"`
}

// NewEvents1 instantiates a new Events1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvents1() *Events1 {
	this := Events1{}
	return &this
}

// NewEvents1WithDefaults instantiates a new Events1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEvents1WithDefaults() *Events1 {
	this := Events1{}
	return &this
}

// GetSubscribe returns the Subscribe field value if set, zero value otherwise.
func (o *Events1) GetSubscribe() bool {
	if o == nil || IsNil(o.Subscribe) {
		var ret bool
		return ret
	}
	return *o.Subscribe
}

// GetSubscribeOk returns a tuple with the Subscribe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Events1) GetSubscribeOk() (*bool, bool) {
	if o == nil || IsNil(o.Subscribe) {
		return nil, false
	}
	return o.Subscribe, true
}

// HasSubscribe returns a boolean if a field has been set.
func (o *Events1) HasSubscribe() bool {
	if o != nil && !IsNil(o.Subscribe) {
		return true
	}

	return false
}

// SetSubscribe gets a reference to the given bool and assigns it to the Subscribe field.
func (o *Events1) SetSubscribe(v bool) {
	o.Subscribe = &v
}

// GetUnsubscribe returns the Unsubscribe field value if set, zero value otherwise.
func (o *Events1) GetUnsubscribe() bool {
	if o == nil || IsNil(o.Unsubscribe) {
		var ret bool
		return ret
	}
	return *o.Unsubscribe
}

// GetUnsubscribeOk returns a tuple with the Unsubscribe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Events1) GetUnsubscribeOk() (*bool, bool) {
	if o == nil || IsNil(o.Unsubscribe) {
		return nil, false
	}
	return o.Unsubscribe, true
}

// HasUnsubscribe returns a boolean if a field has been set.
func (o *Events1) HasUnsubscribe() bool {
	if o != nil && !IsNil(o.Unsubscribe) {
		return true
	}

	return false
}

// SetUnsubscribe gets a reference to the given bool and assigns it to the Unsubscribe field.
func (o *Events1) SetUnsubscribe(v bool) {
	o.Unsubscribe = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *Events1) GetProfile() bool {
	if o == nil || IsNil(o.Profile) {
		var ret bool
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Events1) GetProfileOk() (*bool, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *Events1) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given bool and assigns it to the Profile field.
func (o *Events1) SetProfile(v bool) {
	o.Profile = &v
}

// GetCleaned returns the Cleaned field value if set, zero value otherwise.
func (o *Events1) GetCleaned() bool {
	if o == nil || IsNil(o.Cleaned) {
		var ret bool
		return ret
	}
	return *o.Cleaned
}

// GetCleanedOk returns a tuple with the Cleaned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Events1) GetCleanedOk() (*bool, bool) {
	if o == nil || IsNil(o.Cleaned) {
		return nil, false
	}
	return o.Cleaned, true
}

// HasCleaned returns a boolean if a field has been set.
func (o *Events1) HasCleaned() bool {
	if o != nil && !IsNil(o.Cleaned) {
		return true
	}

	return false
}

// SetCleaned gets a reference to the given bool and assigns it to the Cleaned field.
func (o *Events1) SetCleaned(v bool) {
	o.Cleaned = &v
}

// GetUpemail returns the Upemail field value if set, zero value otherwise.
func (o *Events1) GetUpemail() bool {
	if o == nil || IsNil(o.Upemail) {
		var ret bool
		return ret
	}
	return *o.Upemail
}

// GetUpemailOk returns a tuple with the Upemail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Events1) GetUpemailOk() (*bool, bool) {
	if o == nil || IsNil(o.Upemail) {
		return nil, false
	}
	return o.Upemail, true
}

// HasUpemail returns a boolean if a field has been set.
func (o *Events1) HasUpemail() bool {
	if o != nil && !IsNil(o.Upemail) {
		return true
	}

	return false
}

// SetUpemail gets a reference to the given bool and assigns it to the Upemail field.
func (o *Events1) SetUpemail(v bool) {
	o.Upemail = &v
}

// GetCampaign returns the Campaign field value if set, zero value otherwise.
func (o *Events1) GetCampaign() bool {
	if o == nil || IsNil(o.Campaign) {
		var ret bool
		return ret
	}
	return *o.Campaign
}

// GetCampaignOk returns a tuple with the Campaign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Events1) GetCampaignOk() (*bool, bool) {
	if o == nil || IsNil(o.Campaign) {
		return nil, false
	}
	return o.Campaign, true
}

// HasCampaign returns a boolean if a field has been set.
func (o *Events1) HasCampaign() bool {
	if o != nil && !IsNil(o.Campaign) {
		return true
	}

	return false
}

// SetCampaign gets a reference to the given bool and assigns it to the Campaign field.
func (o *Events1) SetCampaign(v bool) {
	o.Campaign = &v
}

func (o Events1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Events1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Subscribe) {
		toSerialize["subscribe"] = o.Subscribe
	}
	if !IsNil(o.Unsubscribe) {
		toSerialize["unsubscribe"] = o.Unsubscribe
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !IsNil(o.Cleaned) {
		toSerialize["cleaned"] = o.Cleaned
	}
	if !IsNil(o.Upemail) {
		toSerialize["upemail"] = o.Upemail
	}
	if !IsNil(o.Campaign) {
		toSerialize["campaign"] = o.Campaign
	}
	return toSerialize, nil
}

type NullableEvents1 struct {
	value *Events1
	isSet bool
}

func (v NullableEvents1) Get() *Events1 {
	return v.value
}

func (v *NullableEvents1) Set(val *Events1) {
	v.value = val
	v.isSet = true
}

func (v NullableEvents1) IsSet() bool {
	return v.isSet
}

func (v *NullableEvents1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvents1(val *Events1) *NullableEvents1 {
	return &NullableEvents1{value: val, isSet: true}
}

func (v NullableEvents1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvents1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


