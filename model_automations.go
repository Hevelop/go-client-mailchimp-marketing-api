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

// checks if the Automations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Automations{}

// Automations Details for the automations attached to this store.
type Automations struct {
	AbandonedCart *AbandonedCartAutomation `json:"abandoned_cart,omitempty"`
	AbandonedBrowse *AbandonedBrowseAutomation `json:"abandoned_browse,omitempty"`
}

// NewAutomations instantiates a new Automations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutomations() *Automations {
	this := Automations{}
	return &this
}

// NewAutomationsWithDefaults instantiates a new Automations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutomationsWithDefaults() *Automations {
	this := Automations{}
	return &this
}

// GetAbandonedCart returns the AbandonedCart field value if set, zero value otherwise.
func (o *Automations) GetAbandonedCart() AbandonedCartAutomation {
	if o == nil || IsNil(o.AbandonedCart) {
		var ret AbandonedCartAutomation
		return ret
	}
	return *o.AbandonedCart
}

// GetAbandonedCartOk returns a tuple with the AbandonedCart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Automations) GetAbandonedCartOk() (*AbandonedCartAutomation, bool) {
	if o == nil || IsNil(o.AbandonedCart) {
		return nil, false
	}
	return o.AbandonedCart, true
}

// HaveAbandonedCart returns a boolean if a field has been set.
func (o *Automations) HaveAbandonedCart() bool {
	if o != nil && !IsNil(o.AbandonedCart) {
		return true
	}

	return false
}

// SetAbandonedCart gets a reference to the given AbandonedCartAutomation and assigns it to the AbandonedCart field.
func (o *Automations) SetAbandonedCart(v AbandonedCartAutomation) {
	o.AbandonedCart = &v
}

// GetAbandonedBrowse returns the AbandonedBrowse field value if set, zero value otherwise.
func (o *Automations) GetAbandonedBrowse() AbandonedBrowseAutomation {
	if o == nil || IsNil(o.AbandonedBrowse) {
		var ret AbandonedBrowseAutomation
		return ret
	}
	return *o.AbandonedBrowse
}

// GetAbandonedBrowseOk returns a tuple with the AbandonedBrowse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Automations) GetAbandonedBrowseOk() (*AbandonedBrowseAutomation, bool) {
	if o == nil || IsNil(o.AbandonedBrowse) {
		return nil, false
	}
	return o.AbandonedBrowse, true
}

// HaveAbandonedBrowse returns a boolean if a field has been set.
func (o *Automations) HaveAbandonedBrowse() bool {
	if o != nil && !IsNil(o.AbandonedBrowse) {
		return true
	}

	return false
}

// SetAbandonedBrowse gets a reference to the given AbandonedBrowseAutomation and assigns it to the AbandonedBrowse field.
func (o *Automations) SetAbandonedBrowse(v AbandonedBrowseAutomation) {
	o.AbandonedBrowse = &v
}

func (o Automations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Automations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AbandonedCart) {
		toSerialize["abandoned_cart"] = o.AbandonedCart
	}
	if !IsNil(o.AbandonedBrowse) {
		toSerialize["abandoned_browse"] = o.AbandonedBrowse
	}
	return toSerialize, nil
}

type NullableAutomations struct {
	value *Automations
	isSet bool
}

func (v NullableAutomations) Get() *Automations {
	return v.value
}

func (v *NullableAutomations) Set(val *Automations) {
	v.value = val
	v.isSet = true
}

func (v NullableAutomations) IsSet() bool {
	return v.isSet
}

func (v *NullableAutomations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutomations(val *Automations) *NullableAutomations {
	return &NullableAutomations{value: val, isSet: true}
}

func (v NullableAutomations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutomations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


