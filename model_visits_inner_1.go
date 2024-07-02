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

// checks if the VisitsInner1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VisitsInner1{}

// VisitsInner1 struct for VisitsInner1
type VisitsInner1 struct {
	Date *string `json:"date,omitempty"`
	Val *int32 `json:"val,omitempty"`
}

// NewVisitsInner1 instantiates a new VisitsInner1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVisitsInner1() *VisitsInner1 {
	this := VisitsInner1{}
	return &this
}

// NewVisitsInner1WithDefaults instantiates a new VisitsInner1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVisitsInner1WithDefaults() *VisitsInner1 {
	this := VisitsInner1{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *VisitsInner1) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VisitsInner1) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *VisitsInner1) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *VisitsInner1) SetDate(v string) {
	o.Date = &v
}

// GetVal returns the Val field value if set, zero value otherwise.
func (o *VisitsInner1) GetVal() int32 {
	if o == nil || IsNil(o.Val) {
		var ret int32
		return ret
	}
	return *o.Val
}

// GetValOk returns a tuple with the Val field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VisitsInner1) GetValOk() (*int32, bool) {
	if o == nil || IsNil(o.Val) {
		return nil, false
	}
	return o.Val, true
}

// HasVal returns a boolean if a field has been set.
func (o *VisitsInner1) HasVal() bool {
	if o != nil && !IsNil(o.Val) {
		return true
	}

	return false
}

// SetVal gets a reference to the given int32 and assigns it to the Val field.
func (o *VisitsInner1) SetVal(v int32) {
	o.Val = &v
}

func (o VisitsInner1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VisitsInner1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.Val) {
		toSerialize["val"] = o.Val
	}
	return toSerialize, nil
}

type NullableVisitsInner1 struct {
	value *VisitsInner1
	isSet bool
}

func (v NullableVisitsInner1) Get() *VisitsInner1 {
	return v.value
}

func (v *NullableVisitsInner1) Set(val *VisitsInner1) {
	v.value = val
	v.isSet = true
}

func (v NullableVisitsInner1) IsSet() bool {
	return v.isSet
}

func (v *NullableVisitsInner1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVisitsInner1(val *VisitsInner1) *NullableVisitsInner1 {
	return &NullableVisitsInner1{value: val, isSet: true}
}

func (v NullableVisitsInner1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVisitsInner1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

