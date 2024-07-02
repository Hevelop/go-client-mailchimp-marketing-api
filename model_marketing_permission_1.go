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

// checks if the MarketingPermission1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarketingPermission1{}

// MarketingPermission1 A single marketing permission a subscriber has either opted-in to or opted-out of.
type MarketingPermission1 struct {
	// The id for the marketing permission on the list
	MarketingPermissionId *string `json:"marketing_permission_id,omitempty"`
	// If the subscriber has opted-in to the marketing permission.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewMarketingPermission1 instantiates a new MarketingPermission1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketingPermission1() *MarketingPermission1 {
	this := MarketingPermission1{}
	return &this
}

// NewMarketingPermission1WithDefaults instantiates a new MarketingPermission1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketingPermission1WithDefaults() *MarketingPermission1 {
	this := MarketingPermission1{}
	return &this
}

// GetMarketingPermissionId returns the MarketingPermissionId field value if set, zero value otherwise.
func (o *MarketingPermission1) GetMarketingPermissionId() string {
	if o == nil || IsNil(o.MarketingPermissionId) {
		var ret string
		return ret
	}
	return *o.MarketingPermissionId
}

// GetMarketingPermissionIdOk returns a tuple with the MarketingPermissionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingPermission1) GetMarketingPermissionIdOk() (*string, bool) {
	if o == nil || IsNil(o.MarketingPermissionId) {
		return nil, false
	}
	return o.MarketingPermissionId, true
}

// ContainsMarketingPermissionId returns a boolean if a field has been set.
func (o *MarketingPermission1) ContainsMarketingPermissionId() bool {
	if o != nil && !IsNil(o.MarketingPermissionId) {
		return true
	}

	return false
}

// SetMarketingPermissionId gets a reference to the given string and assigns it to the MarketingPermissionId field.
func (o *MarketingPermission1) SetMarketingPermissionId(v string) {
	o.MarketingPermissionId = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *MarketingPermission1) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingPermission1) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// ContainsEnabled returns a boolean if a field has been set.
func (o *MarketingPermission1) ContainsEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *MarketingPermission1) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o MarketingPermission1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarketingPermission1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MarketingPermissionId) {
		toSerialize["marketing_permission_id"] = o.MarketingPermissionId
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableMarketingPermission1 struct {
	value *MarketingPermission1
	isSet bool
}

func (v NullableMarketingPermission1) Get() *MarketingPermission1 {
	return v.value
}

func (v *NullableMarketingPermission1) Set(val *MarketingPermission1) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketingPermission1) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketingPermission1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketingPermission1(val *MarketingPermission1) *NullableMarketingPermission1 {
	return &NullableMarketingPermission1{value: val, isSet: true}
}

func (v NullableMarketingPermission1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketingPermission1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


