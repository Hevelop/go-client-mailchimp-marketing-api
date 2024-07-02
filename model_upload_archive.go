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
	"bytes"
	"fmt"
)

// checks if the UploadArchive type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadArchive{}

// UploadArchive Available when uploading an archive to create campaign content. The archive should include all campaign content and images. [Learn more](https://mailchimp.com/help/import-a-custom-html-template/).
type UploadArchive struct {
	// The base64-encoded representation of the archive file.
	ArchiveContent string `json:"archive_content"`
	// The type of encoded file. Defaults to zip.
	ArchiveType *string `json:"archive_type,omitempty"`
}

type _UploadArchive UploadArchive

// NewUploadArchive instantiates a new UploadArchive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadArchive(archiveContent string) *UploadArchive {
	this := UploadArchive{}
	this.ArchiveContent = archiveContent
	return &this
}

// NewUploadArchiveWithDefaults instantiates a new UploadArchive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadArchiveWithDefaults() *UploadArchive {
	this := UploadArchive{}
	return &this
}

// GetArchiveContent returns the ArchiveContent field value
func (o *UploadArchive) GetArchiveContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ArchiveContent
}

// GetArchiveContentOk returns a tuple with the ArchiveContent field value
// and a boolean to check if the value has been set.
func (o *UploadArchive) GetArchiveContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArchiveContent, true
}

// SetArchiveContent sets field value
func (o *UploadArchive) SetArchiveContent(v string) {
	o.ArchiveContent = v
}

// GetArchiveType returns the ArchiveType field value if set, zero value otherwise.
func (o *UploadArchive) GetArchiveType() string {
	if o == nil || IsNil(o.ArchiveType) {
		var ret string
		return ret
	}
	return *o.ArchiveType
}

// GetArchiveTypeOk returns a tuple with the ArchiveType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadArchive) GetArchiveTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ArchiveType) {
		return nil, false
	}
	return o.ArchiveType, true
}

// HasArchiveType returns a boolean if a field has been set.
func (o *UploadArchive) HasArchiveType() bool {
	if o != nil && !IsNil(o.ArchiveType) {
		return true
	}

	return false
}

// SetArchiveType gets a reference to the given string and assigns it to the ArchiveType field.
func (o *UploadArchive) SetArchiveType(v string) {
	o.ArchiveType = &v
}

func (o UploadArchive) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadArchive) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["archive_content"] = o.ArchiveContent
	if !IsNil(o.ArchiveType) {
		toSerialize["archive_type"] = o.ArchiveType
	}
	return toSerialize, nil
}

func (o *UploadArchive) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"archive_content",
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

	varUploadArchive := _UploadArchive{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUploadArchive)

	if err != nil {
		return err
	}

	*o = UploadArchive(varUploadArchive)

	return err
}

type NullableUploadArchive struct {
	value *UploadArchive
	isSet bool
}

func (v NullableUploadArchive) Get() *UploadArchive {
	return v.value
}

func (v *NullableUploadArchive) Set(val *UploadArchive) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadArchive) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadArchive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadArchive(val *UploadArchive) *NullableUploadArchive {
	return &NullableUploadArchive{value: val, isSet: true}
}

func (v NullableUploadArchive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadArchive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


