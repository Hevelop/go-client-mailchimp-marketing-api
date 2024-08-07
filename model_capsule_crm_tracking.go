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

// checks if the CapsuleCRMTracking type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CapsuleCRMTracking{}

// CapsuleCRMTracking Deprecated
type CapsuleCRMTracking struct {
	// Update contact notes for a campaign based on a subscriber's email addresses.
	Notes *bool `json:"notes,omitempty"`
}

// NewCapsuleCRMTracking instantiates a new CapsuleCRMTracking object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapsuleCRMTracking() *CapsuleCRMTracking {
	this := CapsuleCRMTracking{}
	return &this
}

// NewCapsuleCRMTrackingWithDefaults instantiates a new CapsuleCRMTracking object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapsuleCRMTrackingWithDefaults() *CapsuleCRMTracking {
	this := CapsuleCRMTracking{}
	return &this
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *CapsuleCRMTracking) GetNotes() bool {
	if o == nil || IsNil(o.Notes) {
		var ret bool
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapsuleCRMTracking) GetNotesOk() (*bool, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *CapsuleCRMTracking) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given bool and assigns it to the Notes field.
func (o *CapsuleCRMTracking) SetNotes(v bool) {
	o.Notes = &v
}

func (o CapsuleCRMTracking) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapsuleCRMTracking) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	return toSerialize, nil
}

type NullableCapsuleCRMTracking struct {
	value *CapsuleCRMTracking
	isSet bool
}

func (v NullableCapsuleCRMTracking) Get() *CapsuleCRMTracking {
	return v.value
}

func (v *NullableCapsuleCRMTracking) Set(val *CapsuleCRMTracking) {
	v.value = val
	v.isSet = true
}

func (v NullableCapsuleCRMTracking) IsSet() bool {
	return v.isSet
}

func (v *NullableCapsuleCRMTracking) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapsuleCRMTracking(val *CapsuleCRMTracking) *NullableCapsuleCRMTracking {
	return &NullableCapsuleCRMTracking{value: val, isSet: true}
}

func (v NullableCapsuleCRMTracking) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapsuleCRMTracking) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


