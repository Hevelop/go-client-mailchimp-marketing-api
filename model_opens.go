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

// checks if the Opens type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Opens{}

// Opens An object describing the open activity for the campaign.
type Opens struct {
	// The total number of opens for a campaign.
	OpensTotal *int32 `json:"opens_total,omitempty"`
	// The total number of unique opens.
	UniqueOpens *int32 `json:"unique_opens,omitempty"`
	// The number of unique opens divided by the total number of successful deliveries.
	OpenRate *float32 `json:"open_rate,omitempty"`
	// The date and time of the last recorded open in ISO 8601 format.
	LastOpen *time.Time `json:"last_open,omitempty"`
}

// NewOpens instantiates a new Opens object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpens() *Opens {
	this := Opens{}
	return &this
}

// NewOpensWithDefaults instantiates a new Opens object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpensWithDefaults() *Opens {
	this := Opens{}
	return &this
}

// GetOpensTotal returns the OpensTotal field value if set, zero value otherwise.
func (o *Opens) GetOpensTotal() int32 {
	if o == nil || IsNil(o.OpensTotal) {
		var ret int32
		return ret
	}
	return *o.OpensTotal
}

// GetOpensTotalOk returns a tuple with the OpensTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Opens) GetOpensTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.OpensTotal) {
		return nil, false
	}
	return o.OpensTotal, true
}

// HasOpensTotal returns a boolean if a field has been set.
func (o *Opens) HasOpensTotal() bool {
	if o != nil && !IsNil(o.OpensTotal) {
		return true
	}

	return false
}

// SetOpensTotal gets a reference to the given int32 and assigns it to the OpensTotal field.
func (o *Opens) SetOpensTotal(v int32) {
	o.OpensTotal = &v
}

// GetUniqueOpens returns the UniqueOpens field value if set, zero value otherwise.
func (o *Opens) GetUniqueOpens() int32 {
	if o == nil || IsNil(o.UniqueOpens) {
		var ret int32
		return ret
	}
	return *o.UniqueOpens
}

// GetUniqueOpensOk returns a tuple with the UniqueOpens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Opens) GetUniqueOpensOk() (*int32, bool) {
	if o == nil || IsNil(o.UniqueOpens) {
		return nil, false
	}
	return o.UniqueOpens, true
}

// HasUniqueOpens returns a boolean if a field has been set.
func (o *Opens) HasUniqueOpens() bool {
	if o != nil && !IsNil(o.UniqueOpens) {
		return true
	}

	return false
}

// SetUniqueOpens gets a reference to the given int32 and assigns it to the UniqueOpens field.
func (o *Opens) SetUniqueOpens(v int32) {
	o.UniqueOpens = &v
}

// GetOpenRate returns the OpenRate field value if set, zero value otherwise.
func (o *Opens) GetOpenRate() float32 {
	if o == nil || IsNil(o.OpenRate) {
		var ret float32
		return ret
	}
	return *o.OpenRate
}

// GetOpenRateOk returns a tuple with the OpenRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Opens) GetOpenRateOk() (*float32, bool) {
	if o == nil || IsNil(o.OpenRate) {
		return nil, false
	}
	return o.OpenRate, true
}

// HasOpenRate returns a boolean if a field has been set.
func (o *Opens) HasOpenRate() bool {
	if o != nil && !IsNil(o.OpenRate) {
		return true
	}

	return false
}

// SetOpenRate gets a reference to the given float32 and assigns it to the OpenRate field.
func (o *Opens) SetOpenRate(v float32) {
	o.OpenRate = &v
}

// GetLastOpen returns the LastOpen field value if set, zero value otherwise.
func (o *Opens) GetLastOpen() time.Time {
	if o == nil || IsNil(o.LastOpen) {
		var ret time.Time
		return ret
	}
	return *o.LastOpen
}

// GetLastOpenOk returns a tuple with the LastOpen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Opens) GetLastOpenOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastOpen) {
		return nil, false
	}
	return o.LastOpen, true
}

// HasLastOpen returns a boolean if a field has been set.
func (o *Opens) HasLastOpen() bool {
	if o != nil && !IsNil(o.LastOpen) {
		return true
	}

	return false
}

// SetLastOpen gets a reference to the given time.Time and assigns it to the LastOpen field.
func (o *Opens) SetLastOpen(v time.Time) {
	o.LastOpen = &v
}

func (o Opens) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Opens) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OpensTotal) {
		toSerialize["opens_total"] = o.OpensTotal
	}
	if !IsNil(o.UniqueOpens) {
		toSerialize["unique_opens"] = o.UniqueOpens
	}
	if !IsNil(o.OpenRate) {
		toSerialize["open_rate"] = o.OpenRate
	}
	if !IsNil(o.LastOpen) {
		toSerialize["last_open"] = o.LastOpen
	}
	return toSerialize, nil
}

type NullableOpens struct {
	value *Opens
	isSet bool
}

func (v NullableOpens) Get() *Opens {
	return v.value
}

func (v *NullableOpens) Set(val *Opens) {
	v.value = val
	v.isSet = true
}

func (v NullableOpens) IsSet() bool {
	return v.isSet
}

func (v *NullableOpens) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpens(val *Opens) *NullableOpens {
	return &NullableOpens{value: val, isSet: true}
}

func (v NullableOpens) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpens) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


