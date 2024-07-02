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

// checks if the ECommerceOrderLineItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ECommerceOrderLineItem{}

// ECommerceOrderLineItem Information about a specific order line.
type ECommerceOrderLineItem struct {
	// A unique identifier for an order line item.
	Id *string `json:"id,omitempty"`
	// A unique identifier for the product associated with an order line item.
	ProductId *string `json:"product_id,omitempty"`
	// The name of the product for an order line item.
	ProductTitle *string `json:"product_title,omitempty"`
	// A unique identifier for the product variant associated with an order line item.
	ProductVariantId *string `json:"product_variant_id,omitempty"`
	// The name of the product variant for an order line item.
	ProductVariantTitle *string `json:"product_variant_title,omitempty"`
	// The image URL for a product.
	ImageUrl *string `json:"image_url,omitempty"`
	// The order line item quantity.
	Quantity *int32 `json:"quantity,omitempty"`
	// The order line item price.
	Price *float32 `json:"price,omitempty"`
	// The total discount amount applied to a line item.
	Discount *float32 `json:"discount,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewECommerceOrderLineItem instantiates a new ECommerceOrderLineItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewECommerceOrderLineItem() *ECommerceOrderLineItem {
	this := ECommerceOrderLineItem{}
	return &this
}

// NewECommerceOrderLineItemWithDefaults instantiates a new ECommerceOrderLineItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewECommerceOrderLineItemWithDefaults() *ECommerceOrderLineItem {
	this := ECommerceOrderLineItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ECommerceOrderLineItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrderLineItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HaveId returns a boolean if a field has been set.
func (o *ECommerceOrderLineItem) HaveId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ECommerceOrderLineItem) SetId(v string) {
	o.Id = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *ECommerceOrderLineItem) GetProductId() string {
	if o == nil || IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrderLineItem) GetProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HaveProductId returns a boolean if a field has been set.
func (o *ECommerceOrderLineItem) HaveProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *ECommerceOrderLineItem) SetProductId(v string) {
	o.ProductId = &v
}

// GetProductTitle returns the ProductTitle field value if set, zero value otherwise.
func (o *ECommerceOrderLineItem) GetProductTitle() string {
	if o == nil || IsNil(o.ProductTitle) {
		var ret string
		return ret
	}
	return *o.ProductTitle
}

// GetProductTitleOk returns a tuple with the ProductTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrderLineItem) GetProductTitleOk() (*string, bool) {
	if o == nil || IsNil(o.ProductTitle) {
		return nil, false
	}
	return o.ProductTitle, true
}

// HaveProductTitle returns a boolean if a field has been set.
func (o *ECommerceOrderLineItem) HaveProductTitle() bool {
	if o != nil && !IsNil(o.ProductTitle) {
		return true
	}

	return false
}

// SetProductTitle gets a reference to the given string and assigns it to the ProductTitle field.
func (o *ECommerceOrderLineItem) SetProductTitle(v string) {
	o.ProductTitle = &v
}

// GetProductVariantId returns the ProductVariantId field value if set, zero value otherwise.
func (o *ECommerceOrderLineItem) GetProductVariantId() string {
	if o == nil || IsNil(o.ProductVariantId) {
		var ret string
		return ret
	}
	return *o.ProductVariantId
}

// GetProductVariantIdOk returns a tuple with the ProductVariantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrderLineItem) GetProductVariantIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductVariantId) {
		return nil, false
	}
	return o.ProductVariantId, true
}

// HaveProductVariantId returns a boolean if a field has been set.
func (o *ECommerceOrderLineItem) HaveProductVariantId() bool {
	if o != nil && !IsNil(o.ProductVariantId) {
		return true
	}

	return false
}

// SetProductVariantId gets a reference to the given string and assigns it to the ProductVariantId field.
func (o *ECommerceOrderLineItem) SetProductVariantId(v string) {
	o.ProductVariantId = &v
}

// GetProductVariantTitle returns the ProductVariantTitle field value if set, zero value otherwise.
func (o *ECommerceOrderLineItem) GetProductVariantTitle() string {
	if o == nil || IsNil(o.ProductVariantTitle) {
		var ret string
		return ret
	}
	return *o.ProductVariantTitle
}

// GetProductVariantTitleOk returns a tuple with the ProductVariantTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrderLineItem) GetProductVariantTitleOk() (*string, bool) {
	if o == nil || IsNil(o.ProductVariantTitle) {
		return nil, false
	}
	return o.ProductVariantTitle, true
}

// HaveProductVariantTitle returns a boolean if a field has been set.
func (o *ECommerceOrderLineItem) HaveProductVariantTitle() bool {
	if o != nil && !IsNil(o.ProductVariantTitle) {
		return true
	}

	return false
}

// SetProductVariantTitle gets a reference to the given string and assigns it to the ProductVariantTitle field.
func (o *ECommerceOrderLineItem) SetProductVariantTitle(v string) {
	o.ProductVariantTitle = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *ECommerceOrderLineItem) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrderLineItem) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HaveImageUrl returns a boolean if a field has been set.
func (o *ECommerceOrderLineItem) HaveImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *ECommerceOrderLineItem) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *ECommerceOrderLineItem) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrderLineItem) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HaveQuantity returns a boolean if a field has been set.
func (o *ECommerceOrderLineItem) HaveQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *ECommerceOrderLineItem) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ECommerceOrderLineItem) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrderLineItem) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HavePrice returns a boolean if a field has been set.
func (o *ECommerceOrderLineItem) HavePrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *ECommerceOrderLineItem) SetPrice(v float32) {
	o.Price = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *ECommerceOrderLineItem) GetDiscount() float32 {
	if o == nil || IsNil(o.Discount) {
		var ret float32
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrderLineItem) GetDiscountOk() (*float32, bool) {
	if o == nil || IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HaveDiscount returns a boolean if a field has been set.
func (o *ECommerceOrderLineItem) HaveDiscount() bool {
	if o != nil && !IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given float32 and assigns it to the Discount field.
func (o *ECommerceOrderLineItem) SetDiscount(v float32) {
	o.Discount = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ECommerceOrderLineItem) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrderLineItem) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HaveLinks returns a boolean if a field has been set.
func (o *ECommerceOrderLineItem) HaveLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *ECommerceOrderLineItem) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o ECommerceOrderLineItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ECommerceOrderLineItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ProductId) {
		toSerialize["product_id"] = o.ProductId
	}
	if !IsNil(o.ProductTitle) {
		toSerialize["product_title"] = o.ProductTitle
	}
	if !IsNil(o.ProductVariantId) {
		toSerialize["product_variant_id"] = o.ProductVariantId
	}
	if !IsNil(o.ProductVariantTitle) {
		toSerialize["product_variant_title"] = o.ProductVariantTitle
	}
	if !IsNil(o.ImageUrl) {
		toSerialize["image_url"] = o.ImageUrl
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Discount) {
		toSerialize["discount"] = o.Discount
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableECommerceOrderLineItem struct {
	value *ECommerceOrderLineItem
	isSet bool
}

func (v NullableECommerceOrderLineItem) Get() *ECommerceOrderLineItem {
	return v.value
}

func (v *NullableECommerceOrderLineItem) Set(val *ECommerceOrderLineItem) {
	v.value = val
	v.isSet = true
}

func (v NullableECommerceOrderLineItem) IsSet() bool {
	return v.isSet
}

func (v *NullableECommerceOrderLineItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECommerceOrderLineItem(val *ECommerceOrderLineItem) *NullableECommerceOrderLineItem {
	return &NullableECommerceOrderLineItem{value: val, isSet: true}
}

func (v NullableECommerceOrderLineItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECommerceOrderLineItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


