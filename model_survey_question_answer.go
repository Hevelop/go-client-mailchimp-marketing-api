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

// checks if the SurveyQuestionAnswer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SurveyQuestionAnswer{}

// SurveyQuestionAnswer The details of a survey question's answer.
type SurveyQuestionAnswer struct {
	// The ID of the answer.
	Id *string `json:"id,omitempty"`
	// The raw text answer.
	Value *string `json:"value,omitempty"`
	// The ID of the survey response.
	ResponseId *string `json:"response_id,omitempty"`
	// The date and time when the survey response was submitted in ISO 8601 format.
	SubmittedAt *time.Time `json:"submitted_at,omitempty"`
	Contact *Contact `json:"contact,omitempty"`
	// If this contact was added to the Mailchimp audience via this survey.
	IsNewContact *bool `json:"is_new_contact,omitempty"`
}

// NewSurveyQuestionAnswer instantiates a new SurveyQuestionAnswer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSurveyQuestionAnswer() *SurveyQuestionAnswer {
	this := SurveyQuestionAnswer{}
	return &this
}

// NewSurveyQuestionAnswerWithDefaults instantiates a new SurveyQuestionAnswer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSurveyQuestionAnswerWithDefaults() *SurveyQuestionAnswer {
	this := SurveyQuestionAnswer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SurveyQuestionAnswer) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SurveyQuestionAnswer) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SurveyQuestionAnswer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SurveyQuestionAnswer) SetId(v string) {
	o.Id = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SurveyQuestionAnswer) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SurveyQuestionAnswer) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SurveyQuestionAnswer) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SurveyQuestionAnswer) SetValue(v string) {
	o.Value = &v
}

// GetResponseId returns the ResponseId field value if set, zero value otherwise.
func (o *SurveyQuestionAnswer) GetResponseId() string {
	if o == nil || IsNil(o.ResponseId) {
		var ret string
		return ret
	}
	return *o.ResponseId
}

// GetResponseIdOk returns a tuple with the ResponseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SurveyQuestionAnswer) GetResponseIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseId) {
		return nil, false
	}
	return o.ResponseId, true
}

// HasResponseId returns a boolean if a field has been set.
func (o *SurveyQuestionAnswer) HasResponseId() bool {
	if o != nil && !IsNil(o.ResponseId) {
		return true
	}

	return false
}

// SetResponseId gets a reference to the given string and assigns it to the ResponseId field.
func (o *SurveyQuestionAnswer) SetResponseId(v string) {
	o.ResponseId = &v
}

// GetSubmittedAt returns the SubmittedAt field value if set, zero value otherwise.
func (o *SurveyQuestionAnswer) GetSubmittedAt() time.Time {
	if o == nil || IsNil(o.SubmittedAt) {
		var ret time.Time
		return ret
	}
	return *o.SubmittedAt
}

// GetSubmittedAtOk returns a tuple with the SubmittedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SurveyQuestionAnswer) GetSubmittedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SubmittedAt) {
		return nil, false
	}
	return o.SubmittedAt, true
}

// HasSubmittedAt returns a boolean if a field has been set.
func (o *SurveyQuestionAnswer) HasSubmittedAt() bool {
	if o != nil && !IsNil(o.SubmittedAt) {
		return true
	}

	return false
}

// SetSubmittedAt gets a reference to the given time.Time and assigns it to the SubmittedAt field.
func (o *SurveyQuestionAnswer) SetSubmittedAt(v time.Time) {
	o.SubmittedAt = &v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *SurveyQuestionAnswer) GetContact() Contact {
	if o == nil || IsNil(o.Contact) {
		var ret Contact
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SurveyQuestionAnswer) GetContactOk() (*Contact, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *SurveyQuestionAnswer) HasContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given Contact and assigns it to the Contact field.
func (o *SurveyQuestionAnswer) SetContact(v Contact) {
	o.Contact = &v
}

// GetIsNewContact returns the IsNewContact field value if set, zero value otherwise.
func (o *SurveyQuestionAnswer) GetIsNewContact() bool {
	if o == nil || IsNil(o.IsNewContact) {
		var ret bool
		return ret
	}
	return *o.IsNewContact
}

// GetIsNewContactOk returns a tuple with the IsNewContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SurveyQuestionAnswer) GetIsNewContactOk() (*bool, bool) {
	if o == nil || IsNil(o.IsNewContact) {
		return nil, false
	}
	return o.IsNewContact, true
}

// HasIsNewContact returns a boolean if a field has been set.
func (o *SurveyQuestionAnswer) HasIsNewContact() bool {
	if o != nil && !IsNil(o.IsNewContact) {
		return true
	}

	return false
}

// SetIsNewContact gets a reference to the given bool and assigns it to the IsNewContact field.
func (o *SurveyQuestionAnswer) SetIsNewContact(v bool) {
	o.IsNewContact = &v
}

func (o SurveyQuestionAnswer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SurveyQuestionAnswer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.ResponseId) {
		toSerialize["response_id"] = o.ResponseId
	}
	if !IsNil(o.SubmittedAt) {
		toSerialize["submitted_at"] = o.SubmittedAt
	}
	if !IsNil(o.Contact) {
		toSerialize["contact"] = o.Contact
	}
	if !IsNil(o.IsNewContact) {
		toSerialize["is_new_contact"] = o.IsNewContact
	}
	return toSerialize, nil
}

type NullableSurveyQuestionAnswer struct {
	value *SurveyQuestionAnswer
	isSet bool
}

func (v NullableSurveyQuestionAnswer) Get() *SurveyQuestionAnswer {
	return v.value
}

func (v *NullableSurveyQuestionAnswer) Set(val *SurveyQuestionAnswer) {
	v.value = val
	v.isSet = true
}

func (v NullableSurveyQuestionAnswer) IsSet() bool {
	return v.isSet
}

func (v *NullableSurveyQuestionAnswer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSurveyQuestionAnswer(val *SurveyQuestionAnswer) *NullableSurveyQuestionAnswer {
	return &NullableSurveyQuestionAnswer{value: val, isSet: true}
}

func (v NullableSurveyQuestionAnswer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSurveyQuestionAnswer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


