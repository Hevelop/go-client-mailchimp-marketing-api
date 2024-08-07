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
)

// checks if the TagsInner2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagsInner2{}

// TagsInner2 struct for TagsInner2
type TagsInner2 struct {
	// The unique id for the tag.
	Id *int32 `json:"id,omitempty"`
	// The name of the tag.
	Name *string `json:"name,omitempty"`
	// The date and time the tag was added to the list member in ISO 8601 format.
	DateAdded *time.Time `json:"date_added,omitempty"`
}

// NewTagsInner2 instantiates a new TagsInner2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagsInner2() *TagsInner2 {
	this := TagsInner2{}
	return &this
}

// NewTagsInner2WithDefaults instantiates a new TagsInner2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagsInner2WithDefaults() *TagsInner2 {
	this := TagsInner2{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TagsInner2) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagsInner2) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TagsInner2) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *TagsInner2) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TagsInner2) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagsInner2) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TagsInner2) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TagsInner2) SetName(v string) {
	o.Name = &v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *TagsInner2) GetDateAdded() time.Time {
	if o == nil || IsNil(o.DateAdded) {
		var ret time.Time
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagsInner2) GetDateAddedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *TagsInner2) HasDateAdded() bool {
	if o != nil && !IsNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given time.Time and assigns it to the DateAdded field.
func (o *TagsInner2) SetDateAdded(v time.Time) {
	o.DateAdded = &v
}

func (o TagsInner2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagsInner2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.DateAdded) {
		toSerialize["date_added"] = o.DateAdded
	}
	return toSerialize, nil
}

type NullableTagsInner2 struct {
	value *TagsInner2
	isSet bool
}

func (v NullableTagsInner2) Get() *TagsInner2 {
	return v.value
}

func (v *NullableTagsInner2) Set(val *TagsInner2) {
	v.value = val
	v.isSet = true
}

func (v NullableTagsInner2) IsSet() bool {
	return v.isSet
}

func (v *NullableTagsInner2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagsInner2(val *TagsInner2) *NullableTagsInner2 {
	return &NullableTagsInner2{value: val, isSet: true}
}

func (v NullableTagsInner2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagsInner2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


