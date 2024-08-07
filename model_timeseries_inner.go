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

// checks if the TimeseriesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeseriesInner{}

// TimeseriesInner struct for TimeseriesInner
type TimeseriesInner struct {
	// The date and time for the series in ISO 8601 format.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// The number of emails sent in the timeseries.
	EmailsSent *int32 `json:"emails_sent,omitempty"`
	// The number of unique opens in the timeseries.
	UniqueOpens *int32 `json:"unique_opens,omitempty"`
	// The number of clicks in the timeseries.
	RecipientsClicks *int32 `json:"recipients_clicks,omitempty"`
}

// NewTimeseriesInner instantiates a new TimeseriesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeseriesInner() *TimeseriesInner {
	this := TimeseriesInner{}
	return &this
}

// NewTimeseriesInnerWithDefaults instantiates a new TimeseriesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeseriesInnerWithDefaults() *TimeseriesInner {
	this := TimeseriesInner{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *TimeseriesInner) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesInner) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *TimeseriesInner) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *TimeseriesInner) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetEmailsSent returns the EmailsSent field value if set, zero value otherwise.
func (o *TimeseriesInner) GetEmailsSent() int32 {
	if o == nil || IsNil(o.EmailsSent) {
		var ret int32
		return ret
	}
	return *o.EmailsSent
}

// GetEmailsSentOk returns a tuple with the EmailsSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesInner) GetEmailsSentOk() (*int32, bool) {
	if o == nil || IsNil(o.EmailsSent) {
		return nil, false
	}
	return o.EmailsSent, true
}

// HasEmailsSent returns a boolean if a field has been set.
func (o *TimeseriesInner) HasEmailsSent() bool {
	if o != nil && !IsNil(o.EmailsSent) {
		return true
	}

	return false
}

// SetEmailsSent gets a reference to the given int32 and assigns it to the EmailsSent field.
func (o *TimeseriesInner) SetEmailsSent(v int32) {
	o.EmailsSent = &v
}

// GetUniqueOpens returns the UniqueOpens field value if set, zero value otherwise.
func (o *TimeseriesInner) GetUniqueOpens() int32 {
	if o == nil || IsNil(o.UniqueOpens) {
		var ret int32
		return ret
	}
	return *o.UniqueOpens
}

// GetUniqueOpensOk returns a tuple with the UniqueOpens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesInner) GetUniqueOpensOk() (*int32, bool) {
	if o == nil || IsNil(o.UniqueOpens) {
		return nil, false
	}
	return o.UniqueOpens, true
}

// HasUniqueOpens returns a boolean if a field has been set.
func (o *TimeseriesInner) HasUniqueOpens() bool {
	if o != nil && !IsNil(o.UniqueOpens) {
		return true
	}

	return false
}

// SetUniqueOpens gets a reference to the given int32 and assigns it to the UniqueOpens field.
func (o *TimeseriesInner) SetUniqueOpens(v int32) {
	o.UniqueOpens = &v
}

// GetRecipientsClicks returns the RecipientsClicks field value if set, zero value otherwise.
func (o *TimeseriesInner) GetRecipientsClicks() int32 {
	if o == nil || IsNil(o.RecipientsClicks) {
		var ret int32
		return ret
	}
	return *o.RecipientsClicks
}

// GetRecipientsClicksOk returns a tuple with the RecipientsClicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesInner) GetRecipientsClicksOk() (*int32, bool) {
	if o == nil || IsNil(o.RecipientsClicks) {
		return nil, false
	}
	return o.RecipientsClicks, true
}

// HasRecipientsClicks returns a boolean if a field has been set.
func (o *TimeseriesInner) HasRecipientsClicks() bool {
	if o != nil && !IsNil(o.RecipientsClicks) {
		return true
	}

	return false
}

// SetRecipientsClicks gets a reference to the given int32 and assigns it to the RecipientsClicks field.
func (o *TimeseriesInner) SetRecipientsClicks(v int32) {
	o.RecipientsClicks = &v
}

func (o TimeseriesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeseriesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.EmailsSent) {
		toSerialize["emails_sent"] = o.EmailsSent
	}
	if !IsNil(o.UniqueOpens) {
		toSerialize["unique_opens"] = o.UniqueOpens
	}
	if !IsNil(o.RecipientsClicks) {
		toSerialize["recipients_clicks"] = o.RecipientsClicks
	}
	return toSerialize, nil
}

type NullableTimeseriesInner struct {
	value *TimeseriesInner
	isSet bool
}

func (v NullableTimeseriesInner) Get() *TimeseriesInner {
	return v.value
}

func (v *NullableTimeseriesInner) Set(val *TimeseriesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeseriesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeseriesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeseriesInner(val *TimeseriesInner) *NullableTimeseriesInner {
	return &NullableTimeseriesInner{value: val, isSet: true}
}

func (v NullableTimeseriesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeseriesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


