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
	"time"
	"bytes"
	"fmt"
)

// checks if the ECommerceProduct1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ECommerceProduct1{}

// ECommerceProduct1 Information about a specific product.
type ECommerceProduct1 struct {
	// A unique identifier for the product.
	Id string `json:"id"`
	// The title of a product.
	Title string `json:"title"`
	// The handle of a product.
	Handle *string `json:"handle,omitempty"`
	// The URL for a product.
	Url *string `json:"url,omitempty"`
	// The description of a product.
	Description *string `json:"description,omitempty"`
	// The type of product.
	Type *string `json:"type,omitempty"`
	// The vendor for a product.
	Vendor *string `json:"vendor,omitempty"`
	// The image URL for a product.
	ImageUrl *string `json:"image_url,omitempty"`
	// An array of the product's variants. At least one variant is required for each product. A variant can use the same `id` and `title` as the parent product.
	Variants []ECommerceProductVariant1 `json:"variants"`
	// An array of the product's images.
	Images []ECommerceProductImage1 `json:"images,omitempty"`
	// The date and time the product was published.
	PublishedAtForeign *time.Time `json:"published_at_foreign,omitempty"`
}

type _ECommerceProduct1 ECommerceProduct1

// NewECommerceProduct1 instantiates a new ECommerceProduct1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewECommerceProduct1(id string, title string, variants []ECommerceProductVariant1) *ECommerceProduct1 {
	this := ECommerceProduct1{}
	this.Id = id
	this.Title = title
	this.Variants = variants
	return &this
}

// NewECommerceProduct1WithDefaults instantiates a new ECommerceProduct1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewECommerceProduct1WithDefaults() *ECommerceProduct1 {
	this := ECommerceProduct1{}
	return &this
}

// GetId returns the Id field value
func (o *ECommerceProduct1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ECommerceProduct1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ECommerceProduct1) SetId(v string) {
	o.Id = v
}

// GetTitle returns the Title field value
func (o *ECommerceProduct1) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ECommerceProduct1) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ECommerceProduct1) SetTitle(v string) {
	o.Title = v
}

// GetHandle returns the Handle field value if set, zero value otherwise.
func (o *ECommerceProduct1) GetHandle() string {
	if o == nil || IsNil(o.Handle) {
		var ret string
		return ret
	}
	return *o.Handle
}

// GetHandleOk returns a tuple with the Handle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceProduct1) GetHandleOk() (*string, bool) {
	if o == nil || IsNil(o.Handle) {
		return nil, false
	}
	return o.Handle, true
}

// HasHandle returns a boolean if a field has been set.
func (o *ECommerceProduct1) HasHandle() bool {
	if o != nil && !IsNil(o.Handle) {
		return true
	}

	return false
}

// SetHandle gets a reference to the given string and assigns it to the Handle field.
func (o *ECommerceProduct1) SetHandle(v string) {
	o.Handle = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ECommerceProduct1) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceProduct1) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ECommerceProduct1) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ECommerceProduct1) SetUrl(v string) {
	o.Url = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ECommerceProduct1) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceProduct1) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ECommerceProduct1) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ECommerceProduct1) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ECommerceProduct1) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceProduct1) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ECommerceProduct1) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ECommerceProduct1) SetType(v string) {
	o.Type = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *ECommerceProduct1) GetVendor() string {
	if o == nil || IsNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceProduct1) GetVendorOk() (*string, bool) {
	if o == nil || IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *ECommerceProduct1) HasVendor() bool {
	if o != nil && !IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *ECommerceProduct1) SetVendor(v string) {
	o.Vendor = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *ECommerceProduct1) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceProduct1) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *ECommerceProduct1) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *ECommerceProduct1) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetVariants returns the Variants field value
func (o *ECommerceProduct1) GetVariants() []ECommerceProductVariant1 {
	if o == nil {
		var ret []ECommerceProductVariant1
		return ret
	}

	return o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value
// and a boolean to check if the value has been set.
func (o *ECommerceProduct1) GetVariantsOk() ([]ECommerceProductVariant1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Variants, true
}

// SetVariants sets field value
func (o *ECommerceProduct1) SetVariants(v []ECommerceProductVariant1) {
	o.Variants = v
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *ECommerceProduct1) GetImages() []ECommerceProductImage1 {
	if o == nil || IsNil(o.Images) {
		var ret []ECommerceProductImage1
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceProduct1) GetImagesOk() ([]ECommerceProductImage1, bool) {
	if o == nil || IsNil(o.Images) {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *ECommerceProduct1) HasImages() bool {
	if o != nil && !IsNil(o.Images) {
		return true
	}

	return false
}

// SetImages gets a reference to the given []ECommerceProductImage1 and assigns it to the Images field.
func (o *ECommerceProduct1) SetImages(v []ECommerceProductImage1) {
	o.Images = v
}

// GetPublishedAtForeign returns the PublishedAtForeign field value if set, zero value otherwise.
func (o *ECommerceProduct1) GetPublishedAtForeign() time.Time {
	if o == nil || IsNil(o.PublishedAtForeign) {
		var ret time.Time
		return ret
	}
	return *o.PublishedAtForeign
}

// GetPublishedAtForeignOk returns a tuple with the PublishedAtForeign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceProduct1) GetPublishedAtForeignOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PublishedAtForeign) {
		return nil, false
	}
	return o.PublishedAtForeign, true
}

// HasPublishedAtForeign returns a boolean if a field has been set.
func (o *ECommerceProduct1) HasPublishedAtForeign() bool {
	if o != nil && !IsNil(o.PublishedAtForeign) {
		return true
	}

	return false
}

// SetPublishedAtForeign gets a reference to the given time.Time and assigns it to the PublishedAtForeign field.
func (o *ECommerceProduct1) SetPublishedAtForeign(v time.Time) {
	o.PublishedAtForeign = &v
}

func (o ECommerceProduct1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ECommerceProduct1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["title"] = o.Title
	if !IsNil(o.Handle) {
		toSerialize["handle"] = o.Handle
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Vendor) {
		toSerialize["vendor"] = o.Vendor
	}
	if !IsNil(o.ImageUrl) {
		toSerialize["image_url"] = o.ImageUrl
	}
	toSerialize["variants"] = o.Variants
	if !IsNil(o.Images) {
		toSerialize["images"] = o.Images
	}
	if !IsNil(o.PublishedAtForeign) {
		toSerialize["published_at_foreign"] = o.PublishedAtForeign
	}
	return toSerialize, nil
}

func (o *ECommerceProduct1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"title",
		"variants",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varECommerceProduct1 := _ECommerceProduct1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varECommerceProduct1)

	if err != nil {
		return err
	}

	*o = ECommerceProduct1(varECommerceProduct1)

	return err
}

type NullableECommerceProduct1 struct {
	value *ECommerceProduct1
	isSet bool
}

func (v NullableECommerceProduct1) Get() *ECommerceProduct1 {
	return v.value
}

func (v *NullableECommerceProduct1) Set(val *ECommerceProduct1) {
	v.value = val
	v.isSet = true
}

func (v NullableECommerceProduct1) IsSet() bool {
	return v.isSet
}

func (v *NullableECommerceProduct1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECommerceProduct1(val *ECommerceProduct1) *NullableECommerceProduct1 {
	return &NullableECommerceProduct1{value: val, isSet: true}
}

func (v NullableECommerceProduct1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECommerceProduct1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


