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

// checks if the PromosInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromosInner{}

// PromosInner struct for PromosInner
type PromosInner struct {
	// The Promo Code
	Code *string `json:"code,omitempty"`
	// The amount of discount applied on the total price. For example if the total cost was $100 and the customer paid $95.5, amount_discounted will be 4.5 For free shipping set amount_discounted to 0
	AmountDiscounted *float32 `json:"amount_discounted,omitempty"`
	// Type of discount. For free shipping set type to fixed
	Type *string `json:"type,omitempty"`
}

// NewPromosInner instantiates a new PromosInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromosInner() *PromosInner {
	this := PromosInner{}
	return &this
}

// NewPromosInnerWithDefaults instantiates a new PromosInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromosInnerWithDefaults() *PromosInner {
	this := PromosInner{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *PromosInner) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromosInner) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *PromosInner) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *PromosInner) SetCode(v string) {
	o.Code = &v
}

// GetAmountDiscounted returns the AmountDiscounted field value if set, zero value otherwise.
func (o *PromosInner) GetAmountDiscounted() float32 {
	if o == nil || IsNil(o.AmountDiscounted) {
		var ret float32
		return ret
	}
	return *o.AmountDiscounted
}

// GetAmountDiscountedOk returns a tuple with the AmountDiscounted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromosInner) GetAmountDiscountedOk() (*float32, bool) {
	if o == nil || IsNil(o.AmountDiscounted) {
		return nil, false
	}
	return o.AmountDiscounted, true
}

// HasAmountDiscounted returns a boolean if a field has been set.
func (o *PromosInner) HasAmountDiscounted() bool {
	if o != nil && !IsNil(o.AmountDiscounted) {
		return true
	}

	return false
}

// SetAmountDiscounted gets a reference to the given float32 and assigns it to the AmountDiscounted field.
func (o *PromosInner) SetAmountDiscounted(v float32) {
	o.AmountDiscounted = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PromosInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromosInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PromosInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PromosInner) SetType(v string) {
	o.Type = &v
}

func (o PromosInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromosInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.AmountDiscounted) {
		toSerialize["amount_discounted"] = o.AmountDiscounted
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullablePromosInner struct {
	value *PromosInner
	isSet bool
}

func (v NullablePromosInner) Get() *PromosInner {
	return v.value
}

func (v *NullablePromosInner) Set(val *PromosInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePromosInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePromosInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromosInner(val *PromosInner) *NullablePromosInner {
	return &NullablePromosInner{value: val, isSet: true}
}

func (v NullablePromosInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromosInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


