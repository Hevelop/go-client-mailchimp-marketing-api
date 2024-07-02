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

// checks if the CampaignAdvice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignAdvice{}

// CampaignAdvice Campaign feedback details.
type CampaignAdvice struct {
	// The sentiment type for a feedback message.
	Type *string `json:"type,omitempty"`
	// The advice message.
	Message *string `json:"message,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewCampaignAdvice instantiates a new CampaignAdvice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignAdvice() *CampaignAdvice {
	this := CampaignAdvice{}
	return &this
}

// NewCampaignAdviceWithDefaults instantiates a new CampaignAdvice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignAdviceWithDefaults() *CampaignAdvice {
	this := CampaignAdvice{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CampaignAdvice) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAdvice) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CampaignAdvice) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CampaignAdvice) SetType(v string) {
	o.Type = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CampaignAdvice) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAdvice) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CampaignAdvice) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CampaignAdvice) SetMessage(v string) {
	o.Message = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CampaignAdvice) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignAdvice) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CampaignAdvice) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *CampaignAdvice) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o CampaignAdvice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignAdvice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableCampaignAdvice struct {
	value *CampaignAdvice
	isSet bool
}

func (v NullableCampaignAdvice) Get() *CampaignAdvice {
	return v.value
}

func (v *NullableCampaignAdvice) Set(val *CampaignAdvice) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignAdvice) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignAdvice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignAdvice(val *CampaignAdvice) *NullableCampaignAdvice {
	return &NullableCampaignAdvice{value: val, isSet: true}
}

func (v NullableCampaignAdvice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignAdvice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


