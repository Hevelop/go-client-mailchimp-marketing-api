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

// checks if the SignupForm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignupForm{}

// SignupForm List signup form.
type SignupForm struct {
	Header *SignupFormHeaderOptions `json:"header,omitempty"`
	// The signup form body content.
	Contents []CollectionOfContentForListSignupForms `json:"contents,omitempty"`
	// An array of objects, each representing an element style for the signup form.
	Styles []CollectionOfElementStyleForListSignupForms `json:"styles,omitempty"`
	// Signup form URL.
	SignupFormUrl *string `json:"signup_form_url,omitempty"`
	// The signup form's list id.
	ListId *string `json:"list_id,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewSignupForm instantiates a new SignupForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignupForm() *SignupForm {
	this := SignupForm{}
	return &this
}

// NewSignupFormWithDefaults instantiates a new SignupForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignupFormWithDefaults() *SignupForm {
	this := SignupForm{}
	return &this
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *SignupForm) GetHeader() SignupFormHeaderOptions {
	if o == nil || IsNil(o.Header) {
		var ret SignupFormHeaderOptions
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignupForm) GetHeaderOk() (*SignupFormHeaderOptions, bool) {
	if o == nil || IsNil(o.Header) {
		return nil, false
	}
	return o.Header, true
}

// HaveHeader returns a boolean if a field has been set.
func (o *SignupForm) HaveHeader() bool {
	if o != nil && !IsNil(o.Header) {
		return true
	}

	return false
}

// SetHeader gets a reference to the given SignupFormHeaderOptions and assigns it to the Header field.
func (o *SignupForm) SetHeader(v SignupFormHeaderOptions) {
	o.Header = &v
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *SignupForm) GetContents() []CollectionOfContentForListSignupForms {
	if o == nil || IsNil(o.Contents) {
		var ret []CollectionOfContentForListSignupForms
		return ret
	}
	return o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignupForm) GetContentsOk() ([]CollectionOfContentForListSignupForms, bool) {
	if o == nil || IsNil(o.Contents) {
		return nil, false
	}
	return o.Contents, true
}

// HaveContents returns a boolean if a field has been set.
func (o *SignupForm) HaveContents() bool {
	if o != nil && !IsNil(o.Contents) {
		return true
	}

	return false
}

// SetContents gets a reference to the given []CollectionOfContentForListSignupForms and assigns it to the Contents field.
func (o *SignupForm) SetContents(v []CollectionOfContentForListSignupForms) {
	o.Contents = v
}

// GetStyles returns the Styles field value if set, zero value otherwise.
func (o *SignupForm) GetStyles() []CollectionOfElementStyleForListSignupForms {
	if o == nil || IsNil(o.Styles) {
		var ret []CollectionOfElementStyleForListSignupForms
		return ret
	}
	return o.Styles
}

// GetStylesOk returns a tuple with the Styles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignupForm) GetStylesOk() ([]CollectionOfElementStyleForListSignupForms, bool) {
	if o == nil || IsNil(o.Styles) {
		return nil, false
	}
	return o.Styles, true
}

// HaveStyles returns a boolean if a field has been set.
func (o *SignupForm) HaveStyles() bool {
	if o != nil && !IsNil(o.Styles) {
		return true
	}

	return false
}

// SetStyles gets a reference to the given []CollectionOfElementStyleForListSignupForms and assigns it to the Styles field.
func (o *SignupForm) SetStyles(v []CollectionOfElementStyleForListSignupForms) {
	o.Styles = v
}

// GetSignupFormUrl returns the SignupFormUrl field value if set, zero value otherwise.
func (o *SignupForm) GetSignupFormUrl() string {
	if o == nil || IsNil(o.SignupFormUrl) {
		var ret string
		return ret
	}
	return *o.SignupFormUrl
}

// GetSignupFormUrlOk returns a tuple with the SignupFormUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignupForm) GetSignupFormUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SignupFormUrl) {
		return nil, false
	}
	return o.SignupFormUrl, true
}

// HaveSignupFormUrl returns a boolean if a field has been set.
func (o *SignupForm) HaveSignupFormUrl() bool {
	if o != nil && !IsNil(o.SignupFormUrl) {
		return true
	}

	return false
}

// SetSignupFormUrl gets a reference to the given string and assigns it to the SignupFormUrl field.
func (o *SignupForm) SetSignupFormUrl(v string) {
	o.SignupFormUrl = &v
}

// GetListId returns the ListId field value if set, zero value otherwise.
func (o *SignupForm) GetListId() string {
	if o == nil || IsNil(o.ListId) {
		var ret string
		return ret
	}
	return *o.ListId
}

// GetListIdOk returns a tuple with the ListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignupForm) GetListIdOk() (*string, bool) {
	if o == nil || IsNil(o.ListId) {
		return nil, false
	}
	return o.ListId, true
}

// HaveListId returns a boolean if a field has been set.
func (o *SignupForm) HaveListId() bool {
	if o != nil && !IsNil(o.ListId) {
		return true
	}

	return false
}

// SetListId gets a reference to the given string and assigns it to the ListId field.
func (o *SignupForm) SetListId(v string) {
	o.ListId = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SignupForm) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignupForm) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HaveLinks returns a boolean if a field has been set.
func (o *SignupForm) HaveLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *SignupForm) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o SignupForm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignupForm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Header) {
		toSerialize["header"] = o.Header
	}
	if !IsNil(o.Contents) {
		toSerialize["contents"] = o.Contents
	}
	if !IsNil(o.Styles) {
		toSerialize["styles"] = o.Styles
	}
	if !IsNil(o.SignupFormUrl) {
		toSerialize["signup_form_url"] = o.SignupFormUrl
	}
	if !IsNil(o.ListId) {
		toSerialize["list_id"] = o.ListId
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableSignupForm struct {
	value *SignupForm
	isSet bool
}

func (v NullableSignupForm) Get() *SignupForm {
	return v.value
}

func (v *NullableSignupForm) Set(val *SignupForm) {
	v.value = val
	v.isSet = true
}

func (v NullableSignupForm) IsSet() bool {
	return v.isSet
}

func (v *NullableSignupForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignupForm(val *SignupForm) *NullableSignupForm {
	return &NullableSignupForm{value: val, isSet: true}
}

func (v NullableSignupForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignupForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


