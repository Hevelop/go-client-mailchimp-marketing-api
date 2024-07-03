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

// checks if the ListLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListLocation{}

// ListLocation struct for ListLocation
type ListLocation struct {
	// The name of the country.
	Country *string `json:"country,omitempty"`
	// The ISO 3166 2 digit country code.
	Cc *string `json:"cc,omitempty"`
	// The percent of subscribers in the country.
	Percent *float32 `json:"percent,omitempty"`
	// The total number of subscribers in the country.
	Total *int32 `json:"total,omitempty"`
}

// NewListLocation instantiates a new ListLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListLocation() *ListLocation {
	this := ListLocation{}
	return &this
}

// NewListLocationWithDefaults instantiates a new ListLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListLocationWithDefaults() *ListLocation {
	this := ListLocation{}
	return &this
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *ListLocation) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLocation) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *ListLocation) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *ListLocation) SetCountry(v string) {
	o.Country = &v
}

// GetCc returns the Cc field value if set, zero value otherwise.
func (o *ListLocation) GetCc() string {
	if o == nil || IsNil(o.Cc) {
		var ret string
		return ret
	}
	return *o.Cc
}

// GetCcOk returns a tuple with the Cc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLocation) GetCcOk() (*string, bool) {
	if o == nil || IsNil(o.Cc) {
		return nil, false
	}
	return o.Cc, true
}

// HasCc returns a boolean if a field has been set.
func (o *ListLocation) HasCc() bool {
	if o != nil && !IsNil(o.Cc) {
		return true
	}

	return false
}

// SetCc gets a reference to the given string and assigns it to the Cc field.
func (o *ListLocation) SetCc(v string) {
	o.Cc = &v
}

// GetPercent returns the Percent field value if set, zero value otherwise.
func (o *ListLocation) GetPercent() float32 {
	if o == nil || IsNil(o.Percent) {
		var ret float32
		return ret
	}
	return *o.Percent
}

// GetPercentOk returns a tuple with the Percent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLocation) GetPercentOk() (*float32, bool) {
	if o == nil || IsNil(o.Percent) {
		return nil, false
	}
	return o.Percent, true
}

// HasPercent returns a boolean if a field has been set.
func (o *ListLocation) HasPercent() bool {
	if o != nil && !IsNil(o.Percent) {
		return true
	}

	return false
}

// SetPercent gets a reference to the given float32 and assigns it to the Percent field.
func (o *ListLocation) SetPercent(v float32) {
	o.Percent = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ListLocation) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLocation) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ListLocation) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *ListLocation) SetTotal(v int32) {
	o.Total = &v
}

func (o ListLocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.Cc) {
		toSerialize["cc"] = o.Cc
	}
	if !IsNil(o.Percent) {
		toSerialize["percent"] = o.Percent
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableListLocation struct {
	value *ListLocation
	isSet bool
}

func (v NullableListLocation) Get() *ListLocation {
	return v.value
}

func (v *NullableListLocation) Set(val *ListLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableListLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableListLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListLocation(val *ListLocation) *NullableListLocation {
	return &NullableListLocation{value: val, isSet: true}
}

func (v NullableListLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


