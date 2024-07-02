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

// checks if the SubscriberStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriberStats{}

// SubscriberStats Open and click rates for this subscriber.
type SubscriberStats struct {
	// A subscriber's average open rate.
	AvgOpenRate *float32 `json:"avg_open_rate,omitempty"`
	// A subscriber's average clickthrough rate.
	AvgClickRate *float32 `json:"avg_click_rate,omitempty"`
}

// NewSubscriberStats instantiates a new SubscriberStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriberStats() *SubscriberStats {
	this := SubscriberStats{}
	return &this
}

// NewSubscriberStatsWithDefaults instantiates a new SubscriberStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriberStatsWithDefaults() *SubscriberStats {
	this := SubscriberStats{}
	return &this
}

// GetAvgOpenRate returns the AvgOpenRate field value if set, zero value otherwise.
func (o *SubscriberStats) GetAvgOpenRate() float32 {
	if o == nil || IsNil(o.AvgOpenRate) {
		var ret float32
		return ret
	}
	return *o.AvgOpenRate
}

// GetAvgOpenRateOk returns a tuple with the AvgOpenRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberStats) GetAvgOpenRateOk() (*float32, bool) {
	if o == nil || IsNil(o.AvgOpenRate) {
		return nil, false
	}
	return o.AvgOpenRate, true
}

// HaveAvgOpenRate returns a boolean if a field has been set.
func (o *SubscriberStats) HaveAvgOpenRate() bool {
	if o != nil && !IsNil(o.AvgOpenRate) {
		return true
	}

	return false
}

// SetAvgOpenRate gets a reference to the given float32 and assigns it to the AvgOpenRate field.
func (o *SubscriberStats) SetAvgOpenRate(v float32) {
	o.AvgOpenRate = &v
}

// GetAvgClickRate returns the AvgClickRate field value if set, zero value otherwise.
func (o *SubscriberStats) GetAvgClickRate() float32 {
	if o == nil || IsNil(o.AvgClickRate) {
		var ret float32
		return ret
	}
	return *o.AvgClickRate
}

// GetAvgClickRateOk returns a tuple with the AvgClickRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberStats) GetAvgClickRateOk() (*float32, bool) {
	if o == nil || IsNil(o.AvgClickRate) {
		return nil, false
	}
	return o.AvgClickRate, true
}

// HaveAvgClickRate returns a boolean if a field has been set.
func (o *SubscriberStats) HaveAvgClickRate() bool {
	if o != nil && !IsNil(o.AvgClickRate) {
		return true
	}

	return false
}

// SetAvgClickRate gets a reference to the given float32 and assigns it to the AvgClickRate field.
func (o *SubscriberStats) SetAvgClickRate(v float32) {
	o.AvgClickRate = &v
}

func (o SubscriberStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriberStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AvgOpenRate) {
		toSerialize["avg_open_rate"] = o.AvgOpenRate
	}
	if !IsNil(o.AvgClickRate) {
		toSerialize["avg_click_rate"] = o.AvgClickRate
	}
	return toSerialize, nil
}

type NullableSubscriberStats struct {
	value *SubscriberStats
	isSet bool
}

func (v NullableSubscriberStats) Get() *SubscriberStats {
	return v.value
}

func (v *NullableSubscriberStats) Set(val *SubscriberStats) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriberStats) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriberStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriberStats(val *SubscriberStats) *NullableSubscriberStats {
	return &NullableSubscriberStats{value: val, isSet: true}
}

func (v NullableSubscriberStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriberStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


