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

// checks if the CapsuleCRMTracking1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CapsuleCRMTracking1{}

// CapsuleCRMTracking1 Deprecated
type CapsuleCRMTracking1 struct {
	// Update contact notes for a campaign based on a subscriber's email address.
	Notes *bool `json:"notes,omitempty"`
}

// NewCapsuleCRMTracking1 instantiates a new CapsuleCRMTracking1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapsuleCRMTracking1() *CapsuleCRMTracking1 {
	this := CapsuleCRMTracking1{}
	return &this
}

// NewCapsuleCRMTracking1WithDefaults instantiates a new CapsuleCRMTracking1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapsuleCRMTracking1WithDefaults() *CapsuleCRMTracking1 {
	this := CapsuleCRMTracking1{}
	return &this
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *CapsuleCRMTracking1) GetNotes() bool {
	if o == nil || IsNil(o.Notes) {
		var ret bool
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapsuleCRMTracking1) GetNotesOk() (*bool, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// ContainsNotes returns a boolean if a field has been set.
func (o *CapsuleCRMTracking1) ContainsNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given bool and assigns it to the Notes field.
func (o *CapsuleCRMTracking1) SetNotes(v bool) {
	o.Notes = &v
}

func (o CapsuleCRMTracking1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapsuleCRMTracking1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	return toSerialize, nil
}

type NullableCapsuleCRMTracking1 struct {
	value *CapsuleCRMTracking1
	isSet bool
}

func (v NullableCapsuleCRMTracking1) Get() *CapsuleCRMTracking1 {
	return v.value
}

func (v *NullableCapsuleCRMTracking1) Set(val *CapsuleCRMTracking1) {
	v.value = val
	v.isSet = true
}

func (v NullableCapsuleCRMTracking1) IsSet() bool {
	return v.isSet
}

func (v *NullableCapsuleCRMTracking1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapsuleCRMTracking1(val *CapsuleCRMTracking1) *NullableCapsuleCRMTracking1 {
	return &NullableCapsuleCRMTracking1{value: val, isSet: true}
}

func (v NullableCapsuleCRMTracking1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapsuleCRMTracking1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


