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

// checks if the DailyClicksAndVisitsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DailyClicksAndVisitsData{}

// DailyClicksAndVisitsData The clicks and visits data from the last seven days.
type DailyClicksAndVisitsData struct {
	Clicks []Clicks1Inner `json:"clicks,omitempty"`
	Visits []VisitsInner `json:"visits,omitempty"`
	UniqueVisits []UniqueVisitsInner `json:"unique_visits,omitempty"`
}

// NewDailyClicksAndVisitsData instantiates a new DailyClicksAndVisitsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDailyClicksAndVisitsData() *DailyClicksAndVisitsData {
	this := DailyClicksAndVisitsData{}
	return &this
}

// NewDailyClicksAndVisitsDataWithDefaults instantiates a new DailyClicksAndVisitsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDailyClicksAndVisitsDataWithDefaults() *DailyClicksAndVisitsData {
	this := DailyClicksAndVisitsData{}
	return &this
}

// GetClicks returns the Clicks field value if set, zero value otherwise.
func (o *DailyClicksAndVisitsData) GetClicks() []Clicks1Inner {
	if o == nil || IsNil(o.Clicks) {
		var ret []Clicks1Inner
		return ret
	}
	return o.Clicks
}

// GetClicksOk returns a tuple with the Clicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyClicksAndVisitsData) GetClicksOk() ([]Clicks1Inner, bool) {
	if o == nil || IsNil(o.Clicks) {
		return nil, false
	}
	return o.Clicks, true
}

// HasClicks returns a boolean if a field has been set.
func (o *DailyClicksAndVisitsData) HasClicks() bool {
	if o != nil && !IsNil(o.Clicks) {
		return true
	}

	return false
}

// SetClicks gets a reference to the given []Clicks1Inner and assigns it to the Clicks field.
func (o *DailyClicksAndVisitsData) SetClicks(v []Clicks1Inner) {
	o.Clicks = v
}

// GetVisits returns the Visits field value if set, zero value otherwise.
func (o *DailyClicksAndVisitsData) GetVisits() []VisitsInner {
	if o == nil || IsNil(o.Visits) {
		var ret []VisitsInner
		return ret
	}
	return o.Visits
}

// GetVisitsOk returns a tuple with the Visits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyClicksAndVisitsData) GetVisitsOk() ([]VisitsInner, bool) {
	if o == nil || IsNil(o.Visits) {
		return nil, false
	}
	return o.Visits, true
}

// HasVisits returns a boolean if a field has been set.
func (o *DailyClicksAndVisitsData) HasVisits() bool {
	if o != nil && !IsNil(o.Visits) {
		return true
	}

	return false
}

// SetVisits gets a reference to the given []VisitsInner and assigns it to the Visits field.
func (o *DailyClicksAndVisitsData) SetVisits(v []VisitsInner) {
	o.Visits = v
}

// GetUniqueVisits returns the UniqueVisits field value if set, zero value otherwise.
func (o *DailyClicksAndVisitsData) GetUniqueVisits() []UniqueVisitsInner {
	if o == nil || IsNil(o.UniqueVisits) {
		var ret []UniqueVisitsInner
		return ret
	}
	return o.UniqueVisits
}

// GetUniqueVisitsOk returns a tuple with the UniqueVisits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyClicksAndVisitsData) GetUniqueVisitsOk() ([]UniqueVisitsInner, bool) {
	if o == nil || IsNil(o.UniqueVisits) {
		return nil, false
	}
	return o.UniqueVisits, true
}

// HasUniqueVisits returns a boolean if a field has been set.
func (o *DailyClicksAndVisitsData) HasUniqueVisits() bool {
	if o != nil && !IsNil(o.UniqueVisits) {
		return true
	}

	return false
}

// SetUniqueVisits gets a reference to the given []UniqueVisitsInner and assigns it to the UniqueVisits field.
func (o *DailyClicksAndVisitsData) SetUniqueVisits(v []UniqueVisitsInner) {
	o.UniqueVisits = v
}

func (o DailyClicksAndVisitsData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DailyClicksAndVisitsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Clicks) {
		toSerialize["clicks"] = o.Clicks
	}
	if !IsNil(o.Visits) {
		toSerialize["visits"] = o.Visits
	}
	if !IsNil(o.UniqueVisits) {
		toSerialize["unique_visits"] = o.UniqueVisits
	}
	return toSerialize, nil
}

type NullableDailyClicksAndVisitsData struct {
	value *DailyClicksAndVisitsData
	isSet bool
}

func (v NullableDailyClicksAndVisitsData) Get() *DailyClicksAndVisitsData {
	return v.value
}

func (v *NullableDailyClicksAndVisitsData) Set(val *DailyClicksAndVisitsData) {
	v.value = val
	v.isSet = true
}

func (v NullableDailyClicksAndVisitsData) IsSet() bool {
	return v.isSet
}

func (v *NullableDailyClicksAndVisitsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDailyClicksAndVisitsData(val *DailyClicksAndVisitsData) *NullableDailyClicksAndVisitsData {
	return &NullableDailyClicksAndVisitsData{value: val, isSet: true}
}

func (v NullableDailyClicksAndVisitsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDailyClicksAndVisitsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

