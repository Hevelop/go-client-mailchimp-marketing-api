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

// checks if the FileManager type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileManager{}

// FileManager A list of available images and files stored in the File Manager for the account.
type FileManager struct {
	//  A list of files and images in an account.
	Files []GalleryFile `json:"files,omitempty"`
	// The total size of all File Manager files in bytes.
	TotalFileSize *float32 `json:"total_file_size,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int32 `json:"total_items,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewFileManager instantiates a new FileManager object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileManager() *FileManager {
	this := FileManager{}
	return &this
}

// NewFileManagerWithDefaults instantiates a new FileManager object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileManagerWithDefaults() *FileManager {
	this := FileManager{}
	return &this
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *FileManager) GetFiles() []GalleryFile {
	if o == nil || IsNil(o.Files) {
		var ret []GalleryFile
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileManager) GetFilesOk() ([]GalleryFile, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *FileManager) HasFiles() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []GalleryFile and assigns it to the Files field.
func (o *FileManager) SetFiles(v []GalleryFile) {
	o.Files = v
}

// GetTotalFileSize returns the TotalFileSize field value if set, zero value otherwise.
func (o *FileManager) GetTotalFileSize() float32 {
	if o == nil || IsNil(o.TotalFileSize) {
		var ret float32
		return ret
	}
	return *o.TotalFileSize
}

// GetTotalFileSizeOk returns a tuple with the TotalFileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileManager) GetTotalFileSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalFileSize) {
		return nil, false
	}
	return o.TotalFileSize, true
}

// HasTotalFileSize returns a boolean if a field has been set.
func (o *FileManager) HasTotalFileSize() bool {
	if o != nil && !IsNil(o.TotalFileSize) {
		return true
	}

	return false
}

// SetTotalFileSize gets a reference to the given float32 and assigns it to the TotalFileSize field.
func (o *FileManager) SetTotalFileSize(v float32) {
	o.TotalFileSize = &v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *FileManager) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileManager) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *FileManager) HasTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *FileManager) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *FileManager) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileManager) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *FileManager) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *FileManager) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o FileManager) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileManager) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Files) {
		toSerialize["files"] = o.Files
	}
	if !IsNil(o.TotalFileSize) {
		toSerialize["total_file_size"] = o.TotalFileSize
	}
	if !IsNil(o.TotalItems) {
		toSerialize["total_items"] = o.TotalItems
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableFileManager struct {
	value *FileManager
	isSet bool
}

func (v NullableFileManager) Get() *FileManager {
	return v.value
}

func (v *NullableFileManager) Set(val *FileManager) {
	v.value = val
	v.isSet = true
}

func (v NullableFileManager) IsSet() bool {
	return v.isSet
}

func (v *NullableFileManager) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileManager(val *FileManager) *NullableFileManager {
	return &NullableFileManager{value: val, isSet: true}
}

func (v NullableFileManager) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileManager) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

