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

// checks if the ABSplitStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ABSplitStats{}

// ABSplitStats General stats about different groups of an A/B Split campaign. Does not return information about Multivariate Campaigns.
type ABSplitStats struct {
	A *CampaignA `json:"a,omitempty"`
	B *CampaignB `json:"b,omitempty"`
}

// NewABSplitStats instantiates a new ABSplitStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewABSplitStats() *ABSplitStats {
	this := ABSplitStats{}
	return &this
}

// NewABSplitStatsWithDefaults instantiates a new ABSplitStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewABSplitStatsWithDefaults() *ABSplitStats {
	this := ABSplitStats{}
	return &this
}

// GetA returns the A field value if set, zero value otherwise.
func (o *ABSplitStats) GetA() CampaignA {
	if o == nil || IsNil(o.A) {
		var ret CampaignA
		return ret
	}
	return *o.A
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ABSplitStats) GetAOk() (*CampaignA, bool) {
	if o == nil || IsNil(o.A) {
		return nil, false
	}
	return o.A, true
}

// ContainsA returns a boolean if a field has been set.
func (o *ABSplitStats) ContainsA() bool {
	if o != nil && !IsNil(o.A) {
		return true
	}

	return false
}

// SetA gets a reference to the given CampaignA and assigns it to the A field.
func (o *ABSplitStats) SetA(v CampaignA) {
	o.A = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *ABSplitStats) GetB() CampaignB {
	if o == nil || IsNil(o.B) {
		var ret CampaignB
		return ret
	}
	return *o.B
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ABSplitStats) GetBOk() (*CampaignB, bool) {
	if o == nil || IsNil(o.B) {
		return nil, false
	}
	return o.B, true
}

// ContainsB returns a boolean if a field has been set.
func (o *ABSplitStats) ContainsB() bool {
	if o != nil && !IsNil(o.B) {
		return true
	}

	return false
}

// SetB gets a reference to the given CampaignB and assigns it to the B field.
func (o *ABSplitStats) SetB(v CampaignB) {
	o.B = &v
}

func (o ABSplitStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ABSplitStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.A) {
		toSerialize["a"] = o.A
	}
	if !IsNil(o.B) {
		toSerialize["b"] = o.B
	}
	return toSerialize, nil
}

type NullableABSplitStats struct {
	value *ABSplitStats
	isSet bool
}

func (v NullableABSplitStats) Get() *ABSplitStats {
	return v.value
}

func (v *NullableABSplitStats) Set(val *ABSplitStats) {
	v.value = val
	v.isSet = true
}

func (v NullableABSplitStats) IsSet() bool {
	return v.isSet
}

func (v *NullableABSplitStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableABSplitStats(val *ABSplitStats) *NullableABSplitStats {
	return &NullableABSplitStats{value: val, isSet: true}
}

func (v NullableABSplitStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableABSplitStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


