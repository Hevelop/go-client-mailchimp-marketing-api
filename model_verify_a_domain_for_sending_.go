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

// checks if the VerifyADomainForSending type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerifyADomainForSending{}

// VerifyADomainForSending Submit a response to the verification challenge and verify a domain for sending.
type VerifyADomainForSending struct {
	// The code that was sent to the email address provided when adding a new domain to verify.
	Code string `json:"code"`
}

type _VerifyADomainForSending VerifyADomainForSending

// NewVerifyADomainForSending instantiates a new VerifyADomainForSending object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyADomainForSending(code string) *VerifyADomainForSending {
	this := VerifyADomainForSending{}
	this.Code = code
	return &this
}

// NewVerifyADomainForSendingWithDefaults instantiates a new VerifyADomainForSending object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyADomainForSendingWithDefaults() *VerifyADomainForSending {
	this := VerifyADomainForSending{}
	return &this
}

// GetCode returns the Code field value
func (o *VerifyADomainForSending) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *VerifyADomainForSending) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *VerifyADomainForSending) SetCode(v string) {
	o.Code = v
}

func (o VerifyADomainForSending) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerifyADomainForSending) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	return toSerialize, nil
}

func (o *VerifyADomainForSending) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
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

	varVerifyADomainForSending := _VerifyADomainForSending{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVerifyADomainForSending)

	if err != nil {
		return err
	}

	*o = VerifyADomainForSending(varVerifyADomainForSending)

	return err
}

type NullableVerifyADomainForSending struct {
	value *VerifyADomainForSending
	isSet bool
}

func (v NullableVerifyADomainForSending) Get() *VerifyADomainForSending {
	return v.value
}

func (v *NullableVerifyADomainForSending) Set(val *VerifyADomainForSending) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyADomainForSending) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyADomainForSending) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyADomainForSending(val *VerifyADomainForSending) *NullableVerifyADomainForSending {
	return &NullableVerifyADomainForSending{value: val, isSet: true}
}

func (v NullableVerifyADomainForSending) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyADomainForSending) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


