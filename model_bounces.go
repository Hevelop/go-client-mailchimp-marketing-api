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

// checks if the Bounces type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Bounces{}

// Bounces An object describing the bounce summary for the campaign.
type Bounces struct {
	// The total number of hard bounced email addresses.
	HardBounces *int32 `json:"hard_bounces,omitempty"`
	// The total number of soft bounced email addresses.
	SoftBounces *int32 `json:"soft_bounces,omitempty"`
	// The total number of addresses that were syntax-related bounces.
	SyntaxErrors *int32 `json:"syntax_errors,omitempty"`
}

// NewBounces instantiates a new Bounces object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBounces() *Bounces {
	this := Bounces{}
	return &this
}

// NewBouncesWithDefaults instantiates a new Bounces object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBouncesWithDefaults() *Bounces {
	this := Bounces{}
	return &this
}

// GetHardBounces returns the HardBounces field value if set, zero value otherwise.
func (o *Bounces) GetHardBounces() int32 {
	if o == nil || IsNil(o.HardBounces) {
		var ret int32
		return ret
	}
	return *o.HardBounces
}

// GetHardBouncesOk returns a tuple with the HardBounces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bounces) GetHardBouncesOk() (*int32, bool) {
	if o == nil || IsNil(o.HardBounces) {
		return nil, false
	}
	return o.HardBounces, true
}

// HasHardBounces returns a boolean if a field has been set.
func (o *Bounces) HasHardBounces() bool {
	if o != nil && !IsNil(o.HardBounces) {
		return true
	}

	return false
}

// SetHardBounces gets a reference to the given int32 and assigns it to the HardBounces field.
func (o *Bounces) SetHardBounces(v int32) {
	o.HardBounces = &v
}

// GetSoftBounces returns the SoftBounces field value if set, zero value otherwise.
func (o *Bounces) GetSoftBounces() int32 {
	if o == nil || IsNil(o.SoftBounces) {
		var ret int32
		return ret
	}
	return *o.SoftBounces
}

// GetSoftBouncesOk returns a tuple with the SoftBounces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bounces) GetSoftBouncesOk() (*int32, bool) {
	if o == nil || IsNil(o.SoftBounces) {
		return nil, false
	}
	return o.SoftBounces, true
}

// HasSoftBounces returns a boolean if a field has been set.
func (o *Bounces) HasSoftBounces() bool {
	if o != nil && !IsNil(o.SoftBounces) {
		return true
	}

	return false
}

// SetSoftBounces gets a reference to the given int32 and assigns it to the SoftBounces field.
func (o *Bounces) SetSoftBounces(v int32) {
	o.SoftBounces = &v
}

// GetSyntaxErrors returns the SyntaxErrors field value if set, zero value otherwise.
func (o *Bounces) GetSyntaxErrors() int32 {
	if o == nil || IsNil(o.SyntaxErrors) {
		var ret int32
		return ret
	}
	return *o.SyntaxErrors
}

// GetSyntaxErrorsOk returns a tuple with the SyntaxErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bounces) GetSyntaxErrorsOk() (*int32, bool) {
	if o == nil || IsNil(o.SyntaxErrors) {
		return nil, false
	}
	return o.SyntaxErrors, true
}

// HasSyntaxErrors returns a boolean if a field has been set.
func (o *Bounces) HasSyntaxErrors() bool {
	if o != nil && !IsNil(o.SyntaxErrors) {
		return true
	}

	return false
}

// SetSyntaxErrors gets a reference to the given int32 and assigns it to the SyntaxErrors field.
func (o *Bounces) SetSyntaxErrors(v int32) {
	o.SyntaxErrors = &v
}

func (o Bounces) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Bounces) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HardBounces) {
		toSerialize["hard_bounces"] = o.HardBounces
	}
	if !IsNil(o.SoftBounces) {
		toSerialize["soft_bounces"] = o.SoftBounces
	}
	if !IsNil(o.SyntaxErrors) {
		toSerialize["syntax_errors"] = o.SyntaxErrors
	}
	return toSerialize, nil
}

type NullableBounces struct {
	value *Bounces
	isSet bool
}

func (v NullableBounces) Get() *Bounces {
	return v.value
}

func (v *NullableBounces) Set(val *Bounces) {
	v.value = val
	v.isSet = true
}

func (v NullableBounces) IsSet() bool {
	return v.isSet
}

func (v *NullableBounces) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBounces(val *Bounces) *NullableBounces {
	return &NullableBounces{value: val, isSet: true}
}

func (v NullableBounces) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBounces) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


