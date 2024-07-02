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
	"bytes"
	"fmt"
)

// checks if the CreateAnAccountExport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAnAccountExport{}

// CreateAnAccountExport Creates an account export with the given parameters.
type CreateAnAccountExport struct {
	// The stages of an account export to include.
	IncludeStages []string `json:"include_stages"`
	// An ISO 8601 date that will limit the export to only records created after a given time. For instance, the reports stage will contain any campaign sent after the given timestamp. Audiences, however, are excluded from this limit.
	SinceTimestamp *time.Time `json:"since_timestamp,omitempty"`
}

type _CreateAnAccountExport CreateAnAccountExport

// NewCreateAnAccountExport instantiates a new CreateAnAccountExport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAnAccountExport(includeStages []string) *CreateAnAccountExport {
	this := CreateAnAccountExport{}
	this.IncludeStages = includeStages
	return &this
}

// NewCreateAnAccountExportWithDefaults instantiates a new CreateAnAccountExport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAnAccountExportWithDefaults() *CreateAnAccountExport {
	this := CreateAnAccountExport{}
	return &this
}

// GetIncludeStages returns the IncludeStages field value
func (o *CreateAnAccountExport) GetIncludeStages() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.IncludeStages
}

// GetIncludeStagesOk returns a tuple with the IncludeStages field value
// and a boolean to check if the value has been set.
func (o *CreateAnAccountExport) GetIncludeStagesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IncludeStages, true
}

// SetIncludeStages sets field value
func (o *CreateAnAccountExport) SetIncludeStages(v []string) {
	o.IncludeStages = v
}

// GetSinceTimestamp returns the SinceTimestamp field value if set, zero value otherwise.
func (o *CreateAnAccountExport) GetSinceTimestamp() time.Time {
	if o == nil || IsNil(o.SinceTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.SinceTimestamp
}

// GetSinceTimestampOk returns a tuple with the SinceTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAnAccountExport) GetSinceTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SinceTimestamp) {
		return nil, false
	}
	return o.SinceTimestamp, true
}

// ContainsSinceTimestamp returns a boolean if a field has been set.
func (o *CreateAnAccountExport) ContainsSinceTimestamp() bool {
	if o != nil && !IsNil(o.SinceTimestamp) {
		return true
	}

	return false
}

// SetSinceTimestamp gets a reference to the given time.Time and assigns it to the SinceTimestamp field.
func (o *CreateAnAccountExport) SetSinceTimestamp(v time.Time) {
	o.SinceTimestamp = &v
}

func (o CreateAnAccountExport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAnAccountExport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["include_stages"] = o.IncludeStages
	if !IsNil(o.SinceTimestamp) {
		toSerialize["since_timestamp"] = o.SinceTimestamp
	}
	return toSerialize, nil
}

func (o *CreateAnAccountExport) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"include_stages",
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

	varCreateAnAccountExport := _CreateAnAccountExport{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateAnAccountExport)

	if err != nil {
		return err
	}

	*o = CreateAnAccountExport(varCreateAnAccountExport)

	return err
}

type NullableCreateAnAccountExport struct {
	value *CreateAnAccountExport
	isSet bool
}

func (v NullableCreateAnAccountExport) Get() *CreateAnAccountExport {
	return v.value
}

func (v *NullableCreateAnAccountExport) Set(val *CreateAnAccountExport) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAnAccountExport) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAnAccountExport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAnAccountExport(val *CreateAnAccountExport) *NullableCreateAnAccountExport {
	return &NullableCreateAnAccountExport{value: val, isSet: true}
}

func (v NullableCreateAnAccountExport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAnAccountExport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


