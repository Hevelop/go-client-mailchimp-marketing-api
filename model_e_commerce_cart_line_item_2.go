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

// checks if the ECommerceCartLineItem2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ECommerceCartLineItem2{}

// ECommerceCartLineItem2 Information about a specific cart line item.
type ECommerceCartLineItem2 struct {
	// A unique identifier for the product associated with the cart line item.
	ProductId *string `json:"product_id,omitempty"`
	// A unique identifier for the product variant associated with the cart line item.
	ProductVariantId *string `json:"product_variant_id,omitempty"`
	// The quantity of a cart line item.
	Quantity *int32 `json:"quantity,omitempty"`
	// The price of a cart line item.
	Price *float32 `json:"price,omitempty"`
}

// NewECommerceCartLineItem2 instantiates a new ECommerceCartLineItem2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewECommerceCartLineItem2() *ECommerceCartLineItem2 {
	this := ECommerceCartLineItem2{}
	return &this
}

// NewECommerceCartLineItem2WithDefaults instantiates a new ECommerceCartLineItem2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewECommerceCartLineItem2WithDefaults() *ECommerceCartLineItem2 {
	this := ECommerceCartLineItem2{}
	return &this
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *ECommerceCartLineItem2) GetProductId() string {
	if o == nil || IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCartLineItem2) GetProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// ContainsProductId returns a boolean if a field has been set.
func (o *ECommerceCartLineItem2) ContainsProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *ECommerceCartLineItem2) SetProductId(v string) {
	o.ProductId = &v
}

// GetProductVariantId returns the ProductVariantId field value if set, zero value otherwise.
func (o *ECommerceCartLineItem2) GetProductVariantId() string {
	if o == nil || IsNil(o.ProductVariantId) {
		var ret string
		return ret
	}
	return *o.ProductVariantId
}

// GetProductVariantIdOk returns a tuple with the ProductVariantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCartLineItem2) GetProductVariantIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductVariantId) {
		return nil, false
	}
	return o.ProductVariantId, true
}

// ContainsProductVariantId returns a boolean if a field has been set.
func (o *ECommerceCartLineItem2) ContainsProductVariantId() bool {
	if o != nil && !IsNil(o.ProductVariantId) {
		return true
	}

	return false
}

// SetProductVariantId gets a reference to the given string and assigns it to the ProductVariantId field.
func (o *ECommerceCartLineItem2) SetProductVariantId(v string) {
	o.ProductVariantId = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *ECommerceCartLineItem2) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCartLineItem2) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// ContainsQuantity returns a boolean if a field has been set.
func (o *ECommerceCartLineItem2) ContainsQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *ECommerceCartLineItem2) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ECommerceCartLineItem2) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCartLineItem2) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// ContainsPrice returns a boolean if a field has been set.
func (o *ECommerceCartLineItem2) ContainsPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *ECommerceCartLineItem2) SetPrice(v float32) {
	o.Price = &v
}

func (o ECommerceCartLineItem2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ECommerceCartLineItem2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductId) {
		toSerialize["product_id"] = o.ProductId
	}
	if !IsNil(o.ProductVariantId) {
		toSerialize["product_variant_id"] = o.ProductVariantId
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	return toSerialize, nil
}

type NullableECommerceCartLineItem2 struct {
	value *ECommerceCartLineItem2
	isSet bool
}

func (v NullableECommerceCartLineItem2) Get() *ECommerceCartLineItem2 {
	return v.value
}

func (v *NullableECommerceCartLineItem2) Set(val *ECommerceCartLineItem2) {
	v.value = val
	v.isSet = true
}

func (v NullableECommerceCartLineItem2) IsSet() bool {
	return v.isSet
}

func (v *NullableECommerceCartLineItem2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECommerceCartLineItem2(val *ECommerceCartLineItem2) *NullableECommerceCartLineItem2 {
	return &NullableECommerceCartLineItem2{value: val, isSet: true}
}

func (v NullableECommerceCartLineItem2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECommerceCartLineItem2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


