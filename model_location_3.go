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

// checks if the Location3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Location3{}

// Location3 An individual click location.
type Location3 struct {
	// The two-digit country code for a recorded click.
	Country *string `json:"country,omitempty"`
	// If available, a specific region where the click was recorded.
	Region *string `json:"region,omitempty"`
}

// NewLocation3 instantiates a new Location3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocation3() *Location3 {
	this := Location3{}
	return &this
}

// NewLocation3WithDefaults instantiates a new Location3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocation3WithDefaults() *Location3 {
	this := Location3{}
	return &this
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *Location3) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location3) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *Location3) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *Location3) SetCountry(v string) {
	o.Country = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *Location3) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location3) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *Location3) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *Location3) SetRegion(v string) {
	o.Region = &v
}

func (o Location3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Location3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	return toSerialize, nil
}

type NullableLocation3 struct {
	value *Location3
	isSet bool
}

func (v NullableLocation3) Get() *Location3 {
	return v.value
}

func (v *NullableLocation3) Set(val *Location3) {
	v.value = val
	v.isSet = true
}

func (v NullableLocation3) IsSet() bool {
	return v.isSet
}

func (v *NullableLocation3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocation3(val *Location3) *NullableLocation3 {
	return &NullableLocation3{value: val, isSet: true}
}

func (v NullableLocation3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocation3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


