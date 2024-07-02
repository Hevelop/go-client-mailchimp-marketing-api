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

// checks if the ToNonClickers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToNonClickers{}

// ToNonClickers Determines if the campaign qualifies to be resent to non-clickers.
type ToNonClickers struct {
	// Determines if the campaign qualifies to be resent to this segment.
	IsEligible *bool `json:"is_eligible,omitempty"`
	// The reason the campaign is not eligible to be resent to this segment.
	Reason *string `json:"reason,omitempty"`
}

// NewToNonClickers instantiates a new ToNonClickers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToNonClickers() *ToNonClickers {
	this := ToNonClickers{}
	return &this
}

// NewToNonClickersWithDefaults instantiates a new ToNonClickers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToNonClickersWithDefaults() *ToNonClickers {
	this := ToNonClickers{}
	return &this
}

// GetIsEligible returns the IsEligible field value if set, zero value otherwise.
func (o *ToNonClickers) GetIsEligible() bool {
	if o == nil || IsNil(o.IsEligible) {
		var ret bool
		return ret
	}
	return *o.IsEligible
}

// GetIsEligibleOk returns a tuple with the IsEligible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToNonClickers) GetIsEligibleOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEligible) {
		return nil, false
	}
	return o.IsEligible, true
}

// ContainsIsEligible returns a boolean if a field has been set.
func (o *ToNonClickers) ContainsIsEligible() bool {
	if o != nil && !IsNil(o.IsEligible) {
		return true
	}

	return false
}

// SetIsEligible gets a reference to the given bool and assigns it to the IsEligible field.
func (o *ToNonClickers) SetIsEligible(v bool) {
	o.IsEligible = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ToNonClickers) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToNonClickers) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// ContainsReason returns a boolean if a field has been set.
func (o *ToNonClickers) ContainsReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ToNonClickers) SetReason(v string) {
	o.Reason = &v
}

func (o ToNonClickers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToNonClickers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsEligible) {
		toSerialize["is_eligible"] = o.IsEligible
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

type NullableToNonClickers struct {
	value *ToNonClickers
	isSet bool
}

func (v NullableToNonClickers) Get() *ToNonClickers {
	return v.value
}

func (v *NullableToNonClickers) Set(val *ToNonClickers) {
	v.value = val
	v.isSet = true
}

func (v NullableToNonClickers) IsSet() bool {
	return v.isSet
}

func (v *NullableToNonClickers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToNonClickers(val *ToNonClickers) *NullableToNonClickers {
	return &NullableToNonClickers{value: val, isSet: true}
}

func (v NullableToNonClickers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToNonClickers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


