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

// checks if the GalleryFile2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GalleryFile2{}

// GalleryFile2 An individual file listed in the File Manager.
type GalleryFile2 struct {
	// The id of the folder. Setting `folder_id` to `0` will remove a file from its current folder.
	FolderId *int32 `json:"folder_id,omitempty"`
	// The name of the file.
	Name *string `json:"name,omitempty"`
}

// NewGalleryFile2 instantiates a new GalleryFile2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGalleryFile2() *GalleryFile2 {
	this := GalleryFile2{}
	return &this
}

// NewGalleryFile2WithDefaults instantiates a new GalleryFile2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGalleryFile2WithDefaults() *GalleryFile2 {
	this := GalleryFile2{}
	return &this
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *GalleryFile2) GetFolderId() int32 {
	if o == nil || IsNil(o.FolderId) {
		var ret int32
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GalleryFile2) GetFolderIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FolderId) {
		return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *GalleryFile2) HasFolderId() bool {
	if o != nil && !IsNil(o.FolderId) {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given int32 and assigns it to the FolderId field.
func (o *GalleryFile2) SetFolderId(v int32) {
	o.FolderId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GalleryFile2) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GalleryFile2) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GalleryFile2) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GalleryFile2) SetName(v string) {
	o.Name = &v
}

func (o GalleryFile2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GalleryFile2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FolderId) {
		toSerialize["folder_id"] = o.FolderId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableGalleryFile2 struct {
	value *GalleryFile2
	isSet bool
}

func (v NullableGalleryFile2) Get() *GalleryFile2 {
	return v.value
}

func (v *NullableGalleryFile2) Set(val *GalleryFile2) {
	v.value = val
	v.isSet = true
}

func (v NullableGalleryFile2) IsSet() bool {
	return v.isSet
}

func (v *NullableGalleryFile2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGalleryFile2(val *GalleryFile2) *NullableGalleryFile2 {
	return &NullableGalleryFile2{value: val, isSet: true}
}

func (v NullableGalleryFile2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGalleryFile2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


