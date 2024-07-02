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

// checks if the Sources1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Sources1{}

// Sources1 The possible sources of any events that can trigger the webhook and whether they are enabled.
type Sources1 struct {
	// Whether the webhook is triggered by subscriber-initiated actions.
	User *bool `json:"user,omitempty"`
	// Whether the webhook is triggered by admin-initiated actions in the web interface.
	Admin *bool `json:"admin,omitempty"`
	// Whether the webhook is triggered by actions initiated via the API.
	Api *bool `json:"api,omitempty"`
}

// NewSources1 instantiates a new Sources1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSources1() *Sources1 {
	this := Sources1{}
	return &this
}

// NewSources1WithDefaults instantiates a new Sources1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSources1WithDefaults() *Sources1 {
	this := Sources1{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *Sources1) GetUser() bool {
	if o == nil || IsNil(o.User) {
		var ret bool
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sources1) GetUserOk() (*bool, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *Sources1) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given bool and assigns it to the User field.
func (o *Sources1) SetUser(v bool) {
	o.User = &v
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *Sources1) GetAdmin() bool {
	if o == nil || IsNil(o.Admin) {
		var ret bool
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sources1) GetAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.Admin) {
		return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *Sources1) HasAdmin() bool {
	if o != nil && !IsNil(o.Admin) {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given bool and assigns it to the Admin field.
func (o *Sources1) SetAdmin(v bool) {
	o.Admin = &v
}

// GetApi returns the Api field value if set, zero value otherwise.
func (o *Sources1) GetApi() bool {
	if o == nil || IsNil(o.Api) {
		var ret bool
		return ret
	}
	return *o.Api
}

// GetApiOk returns a tuple with the Api field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sources1) GetApiOk() (*bool, bool) {
	if o == nil || IsNil(o.Api) {
		return nil, false
	}
	return o.Api, true
}

// HasApi returns a boolean if a field has been set.
func (o *Sources1) HasApi() bool {
	if o != nil && !IsNil(o.Api) {
		return true
	}

	return false
}

// SetApi gets a reference to the given bool and assigns it to the Api field.
func (o *Sources1) SetApi(v bool) {
	o.Api = &v
}

func (o Sources1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Sources1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.Admin) {
		toSerialize["admin"] = o.Admin
	}
	if !IsNil(o.Api) {
		toSerialize["api"] = o.Api
	}
	return toSerialize, nil
}

type NullableSources1 struct {
	value *Sources1
	isSet bool
}

func (v NullableSources1) Get() *Sources1 {
	return v.value
}

func (v *NullableSources1) Set(val *Sources1) {
	v.value = val
	v.isSet = true
}

func (v NullableSources1) IsSet() bool {
	return v.isSet
}

func (v *NullableSources1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSources1(val *Sources1) *NullableSources1 {
	return &NullableSources1{value: val, isSet: true}
}

func (v NullableSources1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSources1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


