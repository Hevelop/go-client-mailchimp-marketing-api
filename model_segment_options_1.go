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

// checks if the SegmentOptions1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SegmentOptions1{}

// SegmentOptions1 An object representing all segmentation options. This object should contain a `saved_segment_id` to use an existing segment, or you can create a new segment by including both `match` and `conditions` options.
type SegmentOptions1 struct {
	// The id for an existing saved segment.
	SavedSegmentId *int32 `json:"saved_segment_id,omitempty"`
	// The prebuilt segment id, if a prebuilt segment has been designated for this campaign.
	PrebuiltSegmentId *string `json:"prebuilt_segment_id,omitempty"`
	// Segment match type.
	Match *string `json:"match,omitempty"`
	// Segment match conditions. There are multiple possible types, see the [condition types documentation](https://mailchimp.com/developer/marketing/docs/alternative-schemas/#segment-condition-schemas).
	Conditions []map[string]interface{} `json:"conditions,omitempty"`
}

// NewSegmentOptions1 instantiates a new SegmentOptions1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSegmentOptions1() *SegmentOptions1 {
	this := SegmentOptions1{}
	return &this
}

// NewSegmentOptions1WithDefaults instantiates a new SegmentOptions1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSegmentOptions1WithDefaults() *SegmentOptions1 {
	this := SegmentOptions1{}
	return &this
}

// GetSavedSegmentId returns the SavedSegmentId field value if set, zero value otherwise.
func (o *SegmentOptions1) GetSavedSegmentId() int32 {
	if o == nil || IsNil(o.SavedSegmentId) {
		var ret int32
		return ret
	}
	return *o.SavedSegmentId
}

// GetSavedSegmentIdOk returns a tuple with the SavedSegmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentOptions1) GetSavedSegmentIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SavedSegmentId) {
		return nil, false
	}
	return o.SavedSegmentId, true
}

// HaveSavedSegmentId returns a boolean if a field has been set.
func (o *SegmentOptions1) HaveSavedSegmentId() bool {
	if o != nil && !IsNil(o.SavedSegmentId) {
		return true
	}

	return false
}

// SetSavedSegmentId gets a reference to the given int32 and assigns it to the SavedSegmentId field.
func (o *SegmentOptions1) SetSavedSegmentId(v int32) {
	o.SavedSegmentId = &v
}

// GetPrebuiltSegmentId returns the PrebuiltSegmentId field value if set, zero value otherwise.
func (o *SegmentOptions1) GetPrebuiltSegmentId() string {
	if o == nil || IsNil(o.PrebuiltSegmentId) {
		var ret string
		return ret
	}
	return *o.PrebuiltSegmentId
}

// GetPrebuiltSegmentIdOk returns a tuple with the PrebuiltSegmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentOptions1) GetPrebuiltSegmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.PrebuiltSegmentId) {
		return nil, false
	}
	return o.PrebuiltSegmentId, true
}

// HavePrebuiltSegmentId returns a boolean if a field has been set.
func (o *SegmentOptions1) HavePrebuiltSegmentId() bool {
	if o != nil && !IsNil(o.PrebuiltSegmentId) {
		return true
	}

	return false
}

// SetPrebuiltSegmentId gets a reference to the given string and assigns it to the PrebuiltSegmentId field.
func (o *SegmentOptions1) SetPrebuiltSegmentId(v string) {
	o.PrebuiltSegmentId = &v
}

// GetMatch returns the Match field value if set, zero value otherwise.
func (o *SegmentOptions1) GetMatch() string {
	if o == nil || IsNil(o.Match) {
		var ret string
		return ret
	}
	return *o.Match
}

// GetMatchOk returns a tuple with the Match field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentOptions1) GetMatchOk() (*string, bool) {
	if o == nil || IsNil(o.Match) {
		return nil, false
	}
	return o.Match, true
}

// HaveMatch returns a boolean if a field has been set.
func (o *SegmentOptions1) HaveMatch() bool {
	if o != nil && !IsNil(o.Match) {
		return true
	}

	return false
}

// SetMatch gets a reference to the given string and assigns it to the Match field.
func (o *SegmentOptions1) SetMatch(v string) {
	o.Match = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *SegmentOptions1) GetConditions() []map[string]interface{} {
	if o == nil || IsNil(o.Conditions) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentOptions1) GetConditionsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HaveConditions returns a boolean if a field has been set.
func (o *SegmentOptions1) HaveConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []map[string]interface{} and assigns it to the Conditions field.
func (o *SegmentOptions1) SetConditions(v []map[string]interface{}) {
	o.Conditions = v
}

func (o SegmentOptions1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SegmentOptions1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SavedSegmentId) {
		toSerialize["saved_segment_id"] = o.SavedSegmentId
	}
	if !IsNil(o.PrebuiltSegmentId) {
		toSerialize["prebuilt_segment_id"] = o.PrebuiltSegmentId
	}
	if !IsNil(o.Match) {
		toSerialize["match"] = o.Match
	}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	return toSerialize, nil
}

type NullableSegmentOptions1 struct {
	value *SegmentOptions1
	isSet bool
}

func (v NullableSegmentOptions1) Get() *SegmentOptions1 {
	return v.value
}

func (v *NullableSegmentOptions1) Set(val *SegmentOptions1) {
	v.value = val
	v.isSet = true
}

func (v NullableSegmentOptions1) IsSet() bool {
	return v.isSet
}

func (v *NullableSegmentOptions1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSegmentOptions1(val *SegmentOptions1) *NullableSegmentOptions1 {
	return &NullableSegmentOptions1{value: val, isSet: true}
}

func (v NullableSegmentOptions1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSegmentOptions1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


