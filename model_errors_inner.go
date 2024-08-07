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

// checks if the ErrorsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorsInner{}

// ErrorsInner struct for ErrorsInner
type ErrorsInner struct {
	// The email address that could not be added or updated.
	EmailAddress *string `json:"email_address,omitempty"`
	// The error message indicating why the email address could not be added or updated.
	Error *string `json:"error,omitempty"`
	// A unique code that identifies this specifc error.
	ErrorCode *string `json:"error_code,omitempty"`
	// If the error is field-related, information about which field is at issue.
	Field *string `json:"field,omitempty"`
	// Message indicating how to resolve a field-related error.
	FieldMessage *string `json:"field_message,omitempty"`
}

// NewErrorsInner instantiates a new ErrorsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorsInner() *ErrorsInner {
	this := ErrorsInner{}
	return &this
}

// NewErrorsInnerWithDefaults instantiates a new ErrorsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorsInnerWithDefaults() *ErrorsInner {
	this := ErrorsInner{}
	return &this
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *ErrorsInner) GetEmailAddress() string {
	if o == nil || IsNil(o.EmailAddress) {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorsInner) GetEmailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EmailAddress) {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *ErrorsInner) HasEmailAddress() bool {
	if o != nil && !IsNil(o.EmailAddress) {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *ErrorsInner) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ErrorsInner) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorsInner) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ErrorsInner) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ErrorsInner) SetError(v string) {
	o.Error = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *ErrorsInner) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorsInner) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *ErrorsInner) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *ErrorsInner) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *ErrorsInner) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorsInner) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *ErrorsInner) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *ErrorsInner) SetField(v string) {
	o.Field = &v
}

// GetFieldMessage returns the FieldMessage field value if set, zero value otherwise.
func (o *ErrorsInner) GetFieldMessage() string {
	if o == nil || IsNil(o.FieldMessage) {
		var ret string
		return ret
	}
	return *o.FieldMessage
}

// GetFieldMessageOk returns a tuple with the FieldMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorsInner) GetFieldMessageOk() (*string, bool) {
	if o == nil || IsNil(o.FieldMessage) {
		return nil, false
	}
	return o.FieldMessage, true
}

// HasFieldMessage returns a boolean if a field has been set.
func (o *ErrorsInner) HasFieldMessage() bool {
	if o != nil && !IsNil(o.FieldMessage) {
		return true
	}

	return false
}

// SetFieldMessage gets a reference to the given string and assigns it to the FieldMessage field.
func (o *ErrorsInner) SetFieldMessage(v string) {
	o.FieldMessage = &v
}

func (o ErrorsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmailAddress) {
		toSerialize["email_address"] = o.EmailAddress
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.ErrorCode) {
		toSerialize["error_code"] = o.ErrorCode
	}
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	if !IsNil(o.FieldMessage) {
		toSerialize["field_message"] = o.FieldMessage
	}
	return toSerialize, nil
}

type NullableErrorsInner struct {
	value *ErrorsInner
	isSet bool
}

func (v NullableErrorsInner) Get() *ErrorsInner {
	return v.value
}

func (v *NullableErrorsInner) Set(val *ErrorsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorsInner(val *ErrorsInner) *NullableErrorsInner {
	return &NullableErrorsInner{value: val, isSet: true}
}

func (v NullableErrorsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


