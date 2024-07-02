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

// checks if the ECommerceProductImage2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ECommerceProductImage2{}

// ECommerceProductImage2 Information about a specific product image.
type ECommerceProductImage2 struct {
	// A unique identifier for the product image.
	Id *string `json:"id,omitempty"`
	// The URL for a product image.
	Url *string `json:"url,omitempty"`
	// The list of product variants using the image.
	VariantIds []string `json:"variant_ids,omitempty"`
}

// NewECommerceProductImage2 instantiates a new ECommerceProductImage2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewECommerceProductImage2() *ECommerceProductImage2 {
	this := ECommerceProductImage2{}
	return &this
}

// NewECommerceProductImage2WithDefaults instantiates a new ECommerceProductImage2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewECommerceProductImage2WithDefaults() *ECommerceProductImage2 {
	this := ECommerceProductImage2{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ECommerceProductImage2) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceProductImage2) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ECommerceProductImage2) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ECommerceProductImage2) SetId(v string) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ECommerceProductImage2) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceProductImage2) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ECommerceProductImage2) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ECommerceProductImage2) SetUrl(v string) {
	o.Url = &v
}

// GetVariantIds returns the VariantIds field value if set, zero value otherwise.
func (o *ECommerceProductImage2) GetVariantIds() []string {
	if o == nil || IsNil(o.VariantIds) {
		var ret []string
		return ret
	}
	return o.VariantIds
}

// GetVariantIdsOk returns a tuple with the VariantIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceProductImage2) GetVariantIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.VariantIds) {
		return nil, false
	}
	return o.VariantIds, true
}

// HasVariantIds returns a boolean if a field has been set.
func (o *ECommerceProductImage2) HasVariantIds() bool {
	if o != nil && !IsNil(o.VariantIds) {
		return true
	}

	return false
}

// SetVariantIds gets a reference to the given []string and assigns it to the VariantIds field.
func (o *ECommerceProductImage2) SetVariantIds(v []string) {
	o.VariantIds = v
}

func (o ECommerceProductImage2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ECommerceProductImage2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.VariantIds) {
		toSerialize["variant_ids"] = o.VariantIds
	}
	return toSerialize, nil
}

type NullableECommerceProductImage2 struct {
	value *ECommerceProductImage2
	isSet bool
}

func (v NullableECommerceProductImage2) Get() *ECommerceProductImage2 {
	return v.value
}

func (v *NullableECommerceProductImage2) Set(val *ECommerceProductImage2) {
	v.value = val
	v.isSet = true
}

func (v NullableECommerceProductImage2) IsSet() bool {
	return v.isSet
}

func (v *NullableECommerceProductImage2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECommerceProductImage2(val *ECommerceProductImage2) *NullableECommerceProductImage2 {
	return &NullableECommerceProductImage2{value: val, isSet: true}
}

func (v NullableECommerceProductImage2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECommerceProductImage2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


