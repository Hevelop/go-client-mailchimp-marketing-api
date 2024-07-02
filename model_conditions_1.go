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

// checks if the Conditions1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Conditions1{}

// Conditions1 The [conditions of the segment](https://mailchimp.com/help/save-and-manage-segments/). Static and fuzzy segments don't have conditions.
type Conditions1 struct {
	// Match type.
	Match *string `json:"match,omitempty"`
	// Segment match conditions. There are multiple possible types, see the [condition types documentation](https://mailchimp.com/developer/marketing/docs/alternative-schemas/#segment-condition-schemas).
	Conditions []map[string]interface{} `json:"conditions,omitempty"`
}

// NewConditions1 instantiates a new Conditions1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditions1() *Conditions1 {
	this := Conditions1{}
	return &this
}

// NewConditions1WithDefaults instantiates a new Conditions1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditions1WithDefaults() *Conditions1 {
	this := Conditions1{}
	return &this
}

// GetMatch returns the Match field value if set, zero value otherwise.
func (o *Conditions1) GetMatch() string {
	if o == nil || IsNil(o.Match) {
		var ret string
		return ret
	}
	return *o.Match
}

// GetMatchOk returns a tuple with the Match field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Conditions1) GetMatchOk() (*string, bool) {
	if o == nil || IsNil(o.Match) {
		return nil, false
	}
	return o.Match, true
}

// HaveMatch returns a boolean if a field has been set.
func (o *Conditions1) HaveMatch() bool {
	if o != nil && !IsNil(o.Match) {
		return true
	}

	return false
}

// SetMatch gets a reference to the given string and assigns it to the Match field.
func (o *Conditions1) SetMatch(v string) {
	o.Match = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *Conditions1) GetConditions() []map[string]interface{} {
	if o == nil || IsNil(o.Conditions) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Conditions1) GetConditionsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HaveConditions returns a boolean if a field has been set.
func (o *Conditions1) HaveConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []map[string]interface{} and assigns it to the Conditions field.
func (o *Conditions1) SetConditions(v []map[string]interface{}) {
	o.Conditions = v
}

func (o Conditions1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Conditions1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Match) {
		toSerialize["match"] = o.Match
	}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	return toSerialize, nil
}

type NullableConditions1 struct {
	value *Conditions1
	isSet bool
}

func (v NullableConditions1) Get() *Conditions1 {
	return v.value
}

func (v *NullableConditions1) Set(val *Conditions1) {
	v.value = val
	v.isSet = true
}

func (v NullableConditions1) IsSet() bool {
	return v.isSet
}

func (v *NullableConditions1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditions1(val *Conditions1) *NullableConditions1 {
	return &NullableConditions1{value: val, isSet: true}
}

func (v NullableConditions1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditions1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


