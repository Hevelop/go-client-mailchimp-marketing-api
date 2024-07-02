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

// checks if the AutomationWorkflow1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutomationWorkflow1{}

// AutomationWorkflow1 A summary of an individual Automation workflow's settings and content.
type AutomationWorkflow1 struct {
	Recipients List1 `json:"recipients"`
	Settings *AutomationCampaignSettings1 `json:"settings,omitempty"`
	TriggerSettings AutomationTrigger1 `json:"trigger_settings"`
}

type _AutomationWorkflow1 AutomationWorkflow1

// NewAutomationWorkflow1 instantiates a new AutomationWorkflow1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutomationWorkflow1(recipients List1, triggerSettings AutomationTrigger1) *AutomationWorkflow1 {
	this := AutomationWorkflow1{}
	this.Recipients = recipients
	this.TriggerSettings = triggerSettings
	return &this
}

// NewAutomationWorkflow1WithDefaults instantiates a new AutomationWorkflow1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutomationWorkflow1WithDefaults() *AutomationWorkflow1 {
	this := AutomationWorkflow1{}
	return &this
}

// GetRecipients returns the Recipients field value
func (o *AutomationWorkflow1) GetRecipients() List1 {
	if o == nil {
		var ret List1
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *AutomationWorkflow1) GetRecipientsOk() (*List1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recipients, true
}

// SetRecipients sets field value
func (o *AutomationWorkflow1) SetRecipients(v List1) {
	o.Recipients = v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *AutomationWorkflow1) GetSettings() AutomationCampaignSettings1 {
	if o == nil || IsNil(o.Settings) {
		var ret AutomationCampaignSettings1
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationWorkflow1) GetSettingsOk() (*AutomationCampaignSettings1, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *AutomationWorkflow1) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given AutomationCampaignSettings1 and assigns it to the Settings field.
func (o *AutomationWorkflow1) SetSettings(v AutomationCampaignSettings1) {
	o.Settings = &v
}

// GetTriggerSettings returns the TriggerSettings field value
func (o *AutomationWorkflow1) GetTriggerSettings() AutomationTrigger1 {
	if o == nil {
		var ret AutomationTrigger1
		return ret
	}

	return o.TriggerSettings
}

// GetTriggerSettingsOk returns a tuple with the TriggerSettings field value
// and a boolean to check if the value has been set.
func (o *AutomationWorkflow1) GetTriggerSettingsOk() (*AutomationTrigger1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerSettings, true
}

// SetTriggerSettings sets field value
func (o *AutomationWorkflow1) SetTriggerSettings(v AutomationTrigger1) {
	o.TriggerSettings = v
}

func (o AutomationWorkflow1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutomationWorkflow1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["recipients"] = o.Recipients
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	toSerialize["trigger_settings"] = o.TriggerSettings
	return toSerialize, nil
}

func (o *AutomationWorkflow1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"recipients",
		"trigger_settings",
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

	varAutomationWorkflow1 := _AutomationWorkflow1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAutomationWorkflow1)

	if err != nil {
		return err
	}

	*o = AutomationWorkflow1(varAutomationWorkflow1)

	return err
}

type NullableAutomationWorkflow1 struct {
	value *AutomationWorkflow1
	isSet bool
}

func (v NullableAutomationWorkflow1) Get() *AutomationWorkflow1 {
	return v.value
}

func (v *NullableAutomationWorkflow1) Set(val *AutomationWorkflow1) {
	v.value = val
	v.isSet = true
}

func (v NullableAutomationWorkflow1) IsSet() bool {
	return v.isSet
}

func (v *NullableAutomationWorkflow1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutomationWorkflow1(val *AutomationWorkflow1) *NullableAutomationWorkflow1 {
	return &NullableAutomationWorkflow1{value: val, isSet: true}
}

func (v NullableAutomationWorkflow1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutomationWorkflow1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


