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

// checks if the VerifiedDomains1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerifiedDomains1{}

// VerifiedDomains1 The verified domains currently on the account.
type VerifiedDomains1 struct {
	// The domains on the account
	Domains []VerifiedDomains `json:"domains,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int32 `json:"total_items,omitempty"`
}

// NewVerifiedDomains1 instantiates a new VerifiedDomains1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifiedDomains1() *VerifiedDomains1 {
	this := VerifiedDomains1{}
	return &this
}

// NewVerifiedDomains1WithDefaults instantiates a new VerifiedDomains1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifiedDomains1WithDefaults() *VerifiedDomains1 {
	this := VerifiedDomains1{}
	return &this
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *VerifiedDomains1) GetDomains() []VerifiedDomains {
	if o == nil || IsNil(o.Domains) {
		var ret []VerifiedDomains
		return ret
	}
	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifiedDomains1) GetDomainsOk() ([]VerifiedDomains, bool) {
	if o == nil || IsNil(o.Domains) {
		return nil, false
	}
	return o.Domains, true
}

// ContainsDomains returns a boolean if a field has been set.
func (o *VerifiedDomains1) ContainsDomains() bool {
	if o != nil && !IsNil(o.Domains) {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []VerifiedDomains and assigns it to the Domains field.
func (o *VerifiedDomains1) SetDomains(v []VerifiedDomains) {
	o.Domains = v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *VerifiedDomains1) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifiedDomains1) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// ContainsTotalItems returns a boolean if a field has been set.
func (o *VerifiedDomains1) ContainsTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *VerifiedDomains1) SetTotalItems(v int32) {
	o.TotalItems = &v
}

func (o VerifiedDomains1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerifiedDomains1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Domains) {
		toSerialize["domains"] = o.Domains
	}
	if !IsNil(o.TotalItems) {
		toSerialize["total_items"] = o.TotalItems
	}
	return toSerialize, nil
}

type NullableVerifiedDomains1 struct {
	value *VerifiedDomains1
	isSet bool
}

func (v NullableVerifiedDomains1) Get() *VerifiedDomains1 {
	return v.value
}

func (v *NullableVerifiedDomains1) Set(val *VerifiedDomains1) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifiedDomains1) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifiedDomains1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifiedDomains1(val *VerifiedDomains1) *NullableVerifiedDomains1 {
	return &NullableVerifiedDomains1{value: val, isSet: true}
}

func (v NullableVerifiedDomains1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifiedDomains1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


