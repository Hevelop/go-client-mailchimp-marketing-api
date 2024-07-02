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

// checks if the AccountExportsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountExportsInner{}

// AccountExportsInner An account export.
type AccountExportsInner struct {
	// The ID for the export.
	ExportId *int32 `json:"export_id,omitempty"`
	// Start time for the export.
	Started *time.Time `json:"started,omitempty"`
	// If finished, the finish time for the export.
	Finished *time.Time `json:"finished,omitempty"`
	// The size of the uncompressed export in bytes.
	SizeInBytes *int32 `json:"size_in_bytes,omitempty"`
	// If the export is finished, the download URL for an export. URLs are only valid for 90 days after the export completes.
	DownloadUrl *string `json:"download_url,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewAccountExportsInner instantiates a new AccountExportsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountExportsInner() *AccountExportsInner {
	this := AccountExportsInner{}
	return &this
}

// NewAccountExportsInnerWithDefaults instantiates a new AccountExportsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountExportsInnerWithDefaults() *AccountExportsInner {
	this := AccountExportsInner{}
	return &this
}

// GetExportId returns the ExportId field value if set, zero value otherwise.
func (o *AccountExportsInner) GetExportId() int32 {
	if o == nil || IsNil(o.ExportId) {
		var ret int32
		return ret
	}
	return *o.ExportId
}

// GetExportIdOk returns a tuple with the ExportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountExportsInner) GetExportIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ExportId) {
		return nil, false
	}
	return o.ExportId, true
}

// HaveExportId returns a boolean if a field has been set.
func (o *AccountExportsInner) HaveExportId() bool {
	if o != nil && !IsNil(o.ExportId) {
		return true
	}

	return false
}

// SetExportId gets a reference to the given int32 and assigns it to the ExportId field.
func (o *AccountExportsInner) SetExportId(v int32) {
	o.ExportId = &v
}

// GetStarted returns the Started field value if set, zero value otherwise.
func (o *AccountExportsInner) GetStarted() time.Time {
	if o == nil || IsNil(o.Started) {
		var ret time.Time
		return ret
	}
	return *o.Started
}

// GetStartedOk returns a tuple with the Started field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountExportsInner) GetStartedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Started) {
		return nil, false
	}
	return o.Started, true
}

// HaveStarted returns a boolean if a field has been set.
func (o *AccountExportsInner) HaveStarted() bool {
	if o != nil && !IsNil(o.Started) {
		return true
	}

	return false
}

// SetStarted gets a reference to the given time.Time and assigns it to the Started field.
func (o *AccountExportsInner) SetStarted(v time.Time) {
	o.Started = &v
}

// GetFinished returns the Finished field value if set, zero value otherwise.
func (o *AccountExportsInner) GetFinished() time.Time {
	if o == nil || IsNil(o.Finished) {
		var ret time.Time
		return ret
	}
	return *o.Finished
}

// GetFinishedOk returns a tuple with the Finished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountExportsInner) GetFinishedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Finished) {
		return nil, false
	}
	return o.Finished, true
}

// HaveFinished returns a boolean if a field has been set.
func (o *AccountExportsInner) HaveFinished() bool {
	if o != nil && !IsNil(o.Finished) {
		return true
	}

	return false
}

// SetFinished gets a reference to the given time.Time and assigns it to the Finished field.
func (o *AccountExportsInner) SetFinished(v time.Time) {
	o.Finished = &v
}

// GetSizeInBytes returns the SizeInBytes field value if set, zero value otherwise.
func (o *AccountExportsInner) GetSizeInBytes() int32 {
	if o == nil || IsNil(o.SizeInBytes) {
		var ret int32
		return ret
	}
	return *o.SizeInBytes
}

// GetSizeInBytesOk returns a tuple with the SizeInBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountExportsInner) GetSizeInBytesOk() (*int32, bool) {
	if o == nil || IsNil(o.SizeInBytes) {
		return nil, false
	}
	return o.SizeInBytes, true
}

// HaveSizeInBytes returns a boolean if a field has been set.
func (o *AccountExportsInner) HaveSizeInBytes() bool {
	if o != nil && !IsNil(o.SizeInBytes) {
		return true
	}

	return false
}

// SetSizeInBytes gets a reference to the given int32 and assigns it to the SizeInBytes field.
func (o *AccountExportsInner) SetSizeInBytes(v int32) {
	o.SizeInBytes = &v
}

// GetDownloadUrl returns the DownloadUrl field value if set, zero value otherwise.
func (o *AccountExportsInner) GetDownloadUrl() string {
	if o == nil || IsNil(o.DownloadUrl) {
		var ret string
		return ret
	}
	return *o.DownloadUrl
}

// GetDownloadUrlOk returns a tuple with the DownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountExportsInner) GetDownloadUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DownloadUrl) {
		return nil, false
	}
	return o.DownloadUrl, true
}

// HaveDownloadUrl returns a boolean if a field has been set.
func (o *AccountExportsInner) HaveDownloadUrl() bool {
	if o != nil && !IsNil(o.DownloadUrl) {
		return true
	}

	return false
}

// SetDownloadUrl gets a reference to the given string and assigns it to the DownloadUrl field.
func (o *AccountExportsInner) SetDownloadUrl(v string) {
	o.DownloadUrl = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AccountExportsInner) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountExportsInner) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HaveLinks returns a boolean if a field has been set.
func (o *AccountExportsInner) HaveLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *AccountExportsInner) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o AccountExportsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountExportsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExportId) {
		toSerialize["export_id"] = o.ExportId
	}
	if !IsNil(o.Started) {
		toSerialize["started"] = o.Started
	}
	if !IsNil(o.Finished) {
		toSerialize["finished"] = o.Finished
	}
	if !IsNil(o.SizeInBytes) {
		toSerialize["size_in_bytes"] = o.SizeInBytes
	}
	if !IsNil(o.DownloadUrl) {
		toSerialize["download_url"] = o.DownloadUrl
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableAccountExportsInner struct {
	value *AccountExportsInner
	isSet bool
}

func (v NullableAccountExportsInner) Get() *AccountExportsInner {
	return v.value
}

func (v *NullableAccountExportsInner) Set(val *AccountExportsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountExportsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountExportsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountExportsInner(val *AccountExportsInner) *NullableAccountExportsInner {
	return &NullableAccountExportsInner{value: val, isSet: true}
}

func (v NullableAccountExportsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountExportsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


