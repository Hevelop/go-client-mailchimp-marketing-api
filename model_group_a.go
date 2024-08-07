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

// checks if the GroupA type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupA{}

// GroupA Stats for Group A.
type GroupA struct {
	// The total number of clicks for Group A.
	TotalClicksA *int32 `json:"total_clicks_a,omitempty"`
	// The percentage of total clicks for Group A.
	ClickPercentageA *float32 `json:"click_percentage_a,omitempty"`
	// The number of unique clicks for Group A.
	UniqueClicksA *int32 `json:"unique_clicks_a,omitempty"`
	// The percentage of unique clicks for Group A.
	UniqueClickPercentageA *float32 `json:"unique_click_percentage_a,omitempty"`
}

// NewGroupA instantiates a new GroupA object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupA() *GroupA {
	this := GroupA{}
	return &this
}

// NewGroupAWithDefaults instantiates a new GroupA object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupAWithDefaults() *GroupA {
	this := GroupA{}
	return &this
}

// GetTotalClicksA returns the TotalClicksA field value if set, zero value otherwise.
func (o *GroupA) GetTotalClicksA() int32 {
	if o == nil || IsNil(o.TotalClicksA) {
		var ret int32
		return ret
	}
	return *o.TotalClicksA
}

// GetTotalClicksAOk returns a tuple with the TotalClicksA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupA) GetTotalClicksAOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalClicksA) {
		return nil, false
	}
	return o.TotalClicksA, true
}

// HasTotalClicksA returns a boolean if a field has been set.
func (o *GroupA) HasTotalClicksA() bool {
	if o != nil && !IsNil(o.TotalClicksA) {
		return true
	}

	return false
}

// SetTotalClicksA gets a reference to the given int32 and assigns it to the TotalClicksA field.
func (o *GroupA) SetTotalClicksA(v int32) {
	o.TotalClicksA = &v
}

// GetClickPercentageA returns the ClickPercentageA field value if set, zero value otherwise.
func (o *GroupA) GetClickPercentageA() float32 {
	if o == nil || IsNil(o.ClickPercentageA) {
		var ret float32
		return ret
	}
	return *o.ClickPercentageA
}

// GetClickPercentageAOk returns a tuple with the ClickPercentageA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupA) GetClickPercentageAOk() (*float32, bool) {
	if o == nil || IsNil(o.ClickPercentageA) {
		return nil, false
	}
	return o.ClickPercentageA, true
}

// HasClickPercentageA returns a boolean if a field has been set.
func (o *GroupA) HasClickPercentageA() bool {
	if o != nil && !IsNil(o.ClickPercentageA) {
		return true
	}

	return false
}

// SetClickPercentageA gets a reference to the given float32 and assigns it to the ClickPercentageA field.
func (o *GroupA) SetClickPercentageA(v float32) {
	o.ClickPercentageA = &v
}

// GetUniqueClicksA returns the UniqueClicksA field value if set, zero value otherwise.
func (o *GroupA) GetUniqueClicksA() int32 {
	if o == nil || IsNil(o.UniqueClicksA) {
		var ret int32
		return ret
	}
	return *o.UniqueClicksA
}

// GetUniqueClicksAOk returns a tuple with the UniqueClicksA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupA) GetUniqueClicksAOk() (*int32, bool) {
	if o == nil || IsNil(o.UniqueClicksA) {
		return nil, false
	}
	return o.UniqueClicksA, true
}

// HasUniqueClicksA returns a boolean if a field has been set.
func (o *GroupA) HasUniqueClicksA() bool {
	if o != nil && !IsNil(o.UniqueClicksA) {
		return true
	}

	return false
}

// SetUniqueClicksA gets a reference to the given int32 and assigns it to the UniqueClicksA field.
func (o *GroupA) SetUniqueClicksA(v int32) {
	o.UniqueClicksA = &v
}

// GetUniqueClickPercentageA returns the UniqueClickPercentageA field value if set, zero value otherwise.
func (o *GroupA) GetUniqueClickPercentageA() float32 {
	if o == nil || IsNil(o.UniqueClickPercentageA) {
		var ret float32
		return ret
	}
	return *o.UniqueClickPercentageA
}

// GetUniqueClickPercentageAOk returns a tuple with the UniqueClickPercentageA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupA) GetUniqueClickPercentageAOk() (*float32, bool) {
	if o == nil || IsNil(o.UniqueClickPercentageA) {
		return nil, false
	}
	return o.UniqueClickPercentageA, true
}

// HasUniqueClickPercentageA returns a boolean if a field has been set.
func (o *GroupA) HasUniqueClickPercentageA() bool {
	if o != nil && !IsNil(o.UniqueClickPercentageA) {
		return true
	}

	return false
}

// SetUniqueClickPercentageA gets a reference to the given float32 and assigns it to the UniqueClickPercentageA field.
func (o *GroupA) SetUniqueClickPercentageA(v float32) {
	o.UniqueClickPercentageA = &v
}

func (o GroupA) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupA) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalClicksA) {
		toSerialize["total_clicks_a"] = o.TotalClicksA
	}
	if !IsNil(o.ClickPercentageA) {
		toSerialize["click_percentage_a"] = o.ClickPercentageA
	}
	if !IsNil(o.UniqueClicksA) {
		toSerialize["unique_clicks_a"] = o.UniqueClicksA
	}
	if !IsNil(o.UniqueClickPercentageA) {
		toSerialize["unique_click_percentage_a"] = o.UniqueClickPercentageA
	}
	return toSerialize, nil
}

type NullableGroupA struct {
	value *GroupA
	isSet bool
}

func (v NullableGroupA) Get() *GroupA {
	return v.value
}

func (v *NullableGroupA) Set(val *GroupA) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupA) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupA) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupA(val *GroupA) *NullableGroupA {
	return &NullableGroupA{value: val, isSet: true}
}

func (v NullableGroupA) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupA) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


