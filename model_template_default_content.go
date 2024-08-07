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

// checks if the TemplateDefaultContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateDefaultContent{}

// TemplateDefaultContent Default content for a template.
type TemplateDefaultContent struct {
	// The sections that you can edit in the template, including each section's default content.
	Sections map[string]map[string]interface{} `json:"sections,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewTemplateDefaultContent instantiates a new TemplateDefaultContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateDefaultContent() *TemplateDefaultContent {
	this := TemplateDefaultContent{}
	return &this
}

// NewTemplateDefaultContentWithDefaults instantiates a new TemplateDefaultContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateDefaultContentWithDefaults() *TemplateDefaultContent {
	this := TemplateDefaultContent{}
	return &this
}

// GetSections returns the Sections field value if set, zero value otherwise.
func (o *TemplateDefaultContent) GetSections() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Sections) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Sections
}

// GetSectionsOk returns a tuple with the Sections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateDefaultContent) GetSectionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Sections) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Sections, true
}

// HasSections returns a boolean if a field has been set.
func (o *TemplateDefaultContent) HasSections() bool {
	if o != nil && !IsNil(o.Sections) {
		return true
	}

	return false
}

// SetSections gets a reference to the given map[string]map[string]interface{} and assigns it to the Sections field.
func (o *TemplateDefaultContent) SetSections(v map[string]map[string]interface{}) {
	o.Sections = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TemplateDefaultContent) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateDefaultContent) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TemplateDefaultContent) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *TemplateDefaultContent) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o TemplateDefaultContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateDefaultContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sections) {
		toSerialize["sections"] = o.Sections
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableTemplateDefaultContent struct {
	value *TemplateDefaultContent
	isSet bool
}

func (v NullableTemplateDefaultContent) Get() *TemplateDefaultContent {
	return v.value
}

func (v *NullableTemplateDefaultContent) Set(val *TemplateDefaultContent) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateDefaultContent) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateDefaultContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateDefaultContent(val *TemplateDefaultContent) *NullableTemplateDefaultContent {
	return &NullableTemplateDefaultContent{value: val, isSet: true}
}

func (v NullableTemplateDefaultContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateDefaultContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


