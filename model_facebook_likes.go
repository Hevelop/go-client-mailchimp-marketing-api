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

// checks if the FacebookLikes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FacebookLikes{}

// FacebookLikes An object describing campaign engagement on Facebook.
type FacebookLikes struct {
	// The number of recipients who liked the campaign on Facebook.
	RecipientLikes *int32 `json:"recipient_likes,omitempty"`
	// The number of unique likes.
	UniqueLikes *int32 `json:"unique_likes,omitempty"`
	// The number of Facebook likes for the campaign.
	FacebookLikes *int32 `json:"facebook_likes,omitempty"`
}

// NewFacebookLikes instantiates a new FacebookLikes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFacebookLikes() *FacebookLikes {
	this := FacebookLikes{}
	return &this
}

// NewFacebookLikesWithDefaults instantiates a new FacebookLikes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFacebookLikesWithDefaults() *FacebookLikes {
	this := FacebookLikes{}
	return &this
}

// GetRecipientLikes returns the RecipientLikes field value if set, zero value otherwise.
func (o *FacebookLikes) GetRecipientLikes() int32 {
	if o == nil || IsNil(o.RecipientLikes) {
		var ret int32
		return ret
	}
	return *o.RecipientLikes
}

// GetRecipientLikesOk returns a tuple with the RecipientLikes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacebookLikes) GetRecipientLikesOk() (*int32, bool) {
	if o == nil || IsNil(o.RecipientLikes) {
		return nil, false
	}
	return o.RecipientLikes, true
}

// HasRecipientLikes returns a boolean if a field has been set.
func (o *FacebookLikes) HasRecipientLikes() bool {
	if o != nil && !IsNil(o.RecipientLikes) {
		return true
	}

	return false
}

// SetRecipientLikes gets a reference to the given int32 and assigns it to the RecipientLikes field.
func (o *FacebookLikes) SetRecipientLikes(v int32) {
	o.RecipientLikes = &v
}

// GetUniqueLikes returns the UniqueLikes field value if set, zero value otherwise.
func (o *FacebookLikes) GetUniqueLikes() int32 {
	if o == nil || IsNil(o.UniqueLikes) {
		var ret int32
		return ret
	}
	return *o.UniqueLikes
}

// GetUniqueLikesOk returns a tuple with the UniqueLikes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacebookLikes) GetUniqueLikesOk() (*int32, bool) {
	if o == nil || IsNil(o.UniqueLikes) {
		return nil, false
	}
	return o.UniqueLikes, true
}

// HasUniqueLikes returns a boolean if a field has been set.
func (o *FacebookLikes) HasUniqueLikes() bool {
	if o != nil && !IsNil(o.UniqueLikes) {
		return true
	}

	return false
}

// SetUniqueLikes gets a reference to the given int32 and assigns it to the UniqueLikes field.
func (o *FacebookLikes) SetUniqueLikes(v int32) {
	o.UniqueLikes = &v
}

// GetFacebookLikes returns the FacebookLikes field value if set, zero value otherwise.
func (o *FacebookLikes) GetFacebookLikes() int32 {
	if o == nil || IsNil(o.FacebookLikes) {
		var ret int32
		return ret
	}
	return *o.FacebookLikes
}

// GetFacebookLikesOk returns a tuple with the FacebookLikes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacebookLikes) GetFacebookLikesOk() (*int32, bool) {
	if o == nil || IsNil(o.FacebookLikes) {
		return nil, false
	}
	return o.FacebookLikes, true
}

// HasFacebookLikes returns a boolean if a field has been set.
func (o *FacebookLikes) HasFacebookLikes() bool {
	if o != nil && !IsNil(o.FacebookLikes) {
		return true
	}

	return false
}

// SetFacebookLikes gets a reference to the given int32 and assigns it to the FacebookLikes field.
func (o *FacebookLikes) SetFacebookLikes(v int32) {
	o.FacebookLikes = &v
}

func (o FacebookLikes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FacebookLikes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RecipientLikes) {
		toSerialize["recipient_likes"] = o.RecipientLikes
	}
	if !IsNil(o.UniqueLikes) {
		toSerialize["unique_likes"] = o.UniqueLikes
	}
	if !IsNil(o.FacebookLikes) {
		toSerialize["facebook_likes"] = o.FacebookLikes
	}
	return toSerialize, nil
}

type NullableFacebookLikes struct {
	value *FacebookLikes
	isSet bool
}

func (v NullableFacebookLikes) Get() *FacebookLikes {
	return v.value
}

func (v *NullableFacebookLikes) Set(val *FacebookLikes) {
	v.value = val
	v.isSet = true
}

func (v NullableFacebookLikes) IsSet() bool {
	return v.isSet
}

func (v *NullableFacebookLikes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFacebookLikes(val *FacebookLikes) *NullableFacebookLikes {
	return &NullableFacebookLikes{value: val, isSet: true}
}

func (v NullableFacebookLikes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFacebookLikes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


