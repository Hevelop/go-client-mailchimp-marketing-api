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

// checks if the ItemsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemsInner{}

// ItemsInner struct for ItemsInner
type ItemsInner struct {
	// The item type.
	Type *string `json:"type,omitempty"`
	// The ID for the specific item.
	Id *int32 `json:"id,omitempty"`
	// The heading for the specific item.
	Heading *string `json:"heading,omitempty"`
	// Details about the specific feedback item.
	Details *string `json:"details,omitempty"`
}

// NewItemsInner instantiates a new ItemsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemsInner() *ItemsInner {
	this := ItemsInner{}
	return &this
}

// NewItemsInnerWithDefaults instantiates a new ItemsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemsInnerWithDefaults() *ItemsInner {
	this := ItemsInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ItemsInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemsInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ItemsInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ItemsInner) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ItemsInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemsInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ItemsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ItemsInner) SetId(v int32) {
	o.Id = &v
}

// GetHeading returns the Heading field value if set, zero value otherwise.
func (o *ItemsInner) GetHeading() string {
	if o == nil || IsNil(o.Heading) {
		var ret string
		return ret
	}
	return *o.Heading
}

// GetHeadingOk returns a tuple with the Heading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemsInner) GetHeadingOk() (*string, bool) {
	if o == nil || IsNil(o.Heading) {
		return nil, false
	}
	return o.Heading, true
}

// HasHeading returns a boolean if a field has been set.
func (o *ItemsInner) HasHeading() bool {
	if o != nil && !IsNil(o.Heading) {
		return true
	}

	return false
}

// SetHeading gets a reference to the given string and assigns it to the Heading field.
func (o *ItemsInner) SetHeading(v string) {
	o.Heading = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ItemsInner) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemsInner) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ItemsInner) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *ItemsInner) SetDetails(v string) {
	o.Details = &v
}

func (o ItemsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ItemsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Heading) {
		toSerialize["heading"] = o.Heading
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableItemsInner struct {
	value *ItemsInner
	isSet bool
}

func (v NullableItemsInner) Get() *ItemsInner {
	return v.value
}

func (v *NullableItemsInner) Set(val *ItemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableItemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableItemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemsInner(val *ItemsInner) *NullableItemsInner {
	return &NullableItemsInner{value: val, isSet: true}
}

func (v NullableItemsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


