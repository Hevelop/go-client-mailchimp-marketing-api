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

// checks if the ListSignupForms type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSignupForms{}

// ListSignupForms List Signup Forms.
type ListSignupForms struct {
	// List signup form.
	SignupForms []SignupForm `json:"signup_forms,omitempty"`
	// The list id.
	ListId *string `json:"list_id,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int32 `json:"total_items,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewListSignupForms instantiates a new ListSignupForms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSignupForms() *ListSignupForms {
	this := ListSignupForms{}
	return &this
}

// NewListSignupFormsWithDefaults instantiates a new ListSignupForms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSignupFormsWithDefaults() *ListSignupForms {
	this := ListSignupForms{}
	return &this
}

// GetSignupForms returns the SignupForms field value if set, zero value otherwise.
func (o *ListSignupForms) GetSignupForms() []SignupForm {
	if o == nil || IsNil(o.SignupForms) {
		var ret []SignupForm
		return ret
	}
	return o.SignupForms
}

// GetSignupFormsOk returns a tuple with the SignupForms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSignupForms) GetSignupFormsOk() ([]SignupForm, bool) {
	if o == nil || IsNil(o.SignupForms) {
		return nil, false
	}
	return o.SignupForms, true
}

// HasSignupForms returns a boolean if a field has been set.
func (o *ListSignupForms) HasSignupForms() bool {
	if o != nil && !IsNil(o.SignupForms) {
		return true
	}

	return false
}

// SetSignupForms gets a reference to the given []SignupForm and assigns it to the SignupForms field.
func (o *ListSignupForms) SetSignupForms(v []SignupForm) {
	o.SignupForms = v
}

// GetListId returns the ListId field value if set, zero value otherwise.
func (o *ListSignupForms) GetListId() string {
	if o == nil || IsNil(o.ListId) {
		var ret string
		return ret
	}
	return *o.ListId
}

// GetListIdOk returns a tuple with the ListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSignupForms) GetListIdOk() (*string, bool) {
	if o == nil || IsNil(o.ListId) {
		return nil, false
	}
	return o.ListId, true
}

// HasListId returns a boolean if a field has been set.
func (o *ListSignupForms) HasListId() bool {
	if o != nil && !IsNil(o.ListId) {
		return true
	}

	return false
}

// SetListId gets a reference to the given string and assigns it to the ListId field.
func (o *ListSignupForms) SetListId(v string) {
	o.ListId = &v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *ListSignupForms) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSignupForms) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *ListSignupForms) HasTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *ListSignupForms) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ListSignupForms) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSignupForms) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ListSignupForms) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *ListSignupForms) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o ListSignupForms) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSignupForms) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SignupForms) {
		toSerialize["signup_forms"] = o.SignupForms
	}
	if !IsNil(o.ListId) {
		toSerialize["list_id"] = o.ListId
	}
	if !IsNil(o.TotalItems) {
		toSerialize["total_items"] = o.TotalItems
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableListSignupForms struct {
	value *ListSignupForms
	isSet bool
}

func (v NullableListSignupForms) Get() *ListSignupForms {
	return v.value
}

func (v *NullableListSignupForms) Set(val *ListSignupForms) {
	v.value = val
	v.isSet = true
}

func (v NullableListSignupForms) IsSet() bool {
	return v.isSet
}

func (v *NullableListSignupForms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSignupForms(val *ListSignupForms) *NullableListSignupForms {
	return &NullableListSignupForms{value: val, isSet: true}
}

func (v NullableListSignupForms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSignupForms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


