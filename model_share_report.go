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

// checks if the ShareReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShareReport{}

// ShareReport The url and password for the [VIP report](https://mailchimp.com/help/share-a-campaign-report/).
type ShareReport struct {
	// The URL for the VIP report.
	ShareUrl *string `json:"share_url,omitempty"`
	// If password protected, the password for the VIP report.
	SharePassword *string `json:"share_password,omitempty"`
}

// NewShareReport instantiates a new ShareReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShareReport() *ShareReport {
	this := ShareReport{}
	return &this
}

// NewShareReportWithDefaults instantiates a new ShareReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShareReportWithDefaults() *ShareReport {
	this := ShareReport{}
	return &this
}

// GetShareUrl returns the ShareUrl field value if set, zero value otherwise.
func (o *ShareReport) GetShareUrl() string {
	if o == nil || IsNil(o.ShareUrl) {
		var ret string
		return ret
	}
	return *o.ShareUrl
}

// GetShareUrlOk returns a tuple with the ShareUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareReport) GetShareUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ShareUrl) {
		return nil, false
	}
	return o.ShareUrl, true
}

// HasShareUrl returns a boolean if a field has been set.
func (o *ShareReport) HasShareUrl() bool {
	if o != nil && !IsNil(o.ShareUrl) {
		return true
	}

	return false
}

// SetShareUrl gets a reference to the given string and assigns it to the ShareUrl field.
func (o *ShareReport) SetShareUrl(v string) {
	o.ShareUrl = &v
}

// GetSharePassword returns the SharePassword field value if set, zero value otherwise.
func (o *ShareReport) GetSharePassword() string {
	if o == nil || IsNil(o.SharePassword) {
		var ret string
		return ret
	}
	return *o.SharePassword
}

// GetSharePasswordOk returns a tuple with the SharePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareReport) GetSharePasswordOk() (*string, bool) {
	if o == nil || IsNil(o.SharePassword) {
		return nil, false
	}
	return o.SharePassword, true
}

// HasSharePassword returns a boolean if a field has been set.
func (o *ShareReport) HasSharePassword() bool {
	if o != nil && !IsNil(o.SharePassword) {
		return true
	}

	return false
}

// SetSharePassword gets a reference to the given string and assigns it to the SharePassword field.
func (o *ShareReport) SetSharePassword(v string) {
	o.SharePassword = &v
}

func (o ShareReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShareReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShareUrl) {
		toSerialize["share_url"] = o.ShareUrl
	}
	if !IsNil(o.SharePassword) {
		toSerialize["share_password"] = o.SharePassword
	}
	return toSerialize, nil
}

type NullableShareReport struct {
	value *ShareReport
	isSet bool
}

func (v NullableShareReport) Get() *ShareReport {
	return v.value
}

func (v *NullableShareReport) Set(val *ShareReport) {
	v.value = val
	v.isSet = true
}

func (v NullableShareReport) IsSet() bool {
	return v.isSet
}

func (v *NullableShareReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShareReport(val *ShareReport) *NullableShareReport {
	return &NullableShareReport{value: val, isSet: true}
}

func (v NullableShareReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShareReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


