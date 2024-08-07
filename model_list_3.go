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

// checks if the List3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &List3{}

// List3 List settings for the campaign.
type List3 struct {
	// The unique list id.
	ListId *string `json:"list_id,omitempty"`
	// The status of the list used, namely if it's deleted or disabled.
	ListIsActive *bool `json:"list_is_active,omitempty"`
	// The name of the list.
	ListName *string `json:"list_name,omitempty"`
	// A description of the [segment](https://mailchimp.com/help/create-and-send-to-a-segment/) used for the campaign. Formatted as a string marked up with HTML.
	SegmentText *string `json:"segment_text,omitempty"`
	// Count of the recipients on the associated list. Formatted as an integer.
	RecipientCount *int32 `json:"recipient_count,omitempty"`
	SegmentOpts *SegmentOptions1 `json:"segment_opts,omitempty"`
}

// NewList3 instantiates a new List3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewList3() *List3 {
	this := List3{}
	return &this
}

// NewList3WithDefaults instantiates a new List3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewList3WithDefaults() *List3 {
	this := List3{}
	return &this
}

// GetListId returns the ListId field value if set, zero value otherwise.
func (o *List3) GetListId() string {
	if o == nil || IsNil(o.ListId) {
		var ret string
		return ret
	}
	return *o.ListId
}

// GetListIdOk returns a tuple with the ListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *List3) GetListIdOk() (*string, bool) {
	if o == nil || IsNil(o.ListId) {
		return nil, false
	}
	return o.ListId, true
}

// HasListId returns a boolean if a field has been set.
func (o *List3) HasListId() bool {
	if o != nil && !IsNil(o.ListId) {
		return true
	}

	return false
}

// SetListId gets a reference to the given string and assigns it to the ListId field.
func (o *List3) SetListId(v string) {
	o.ListId = &v
}

// GetListIsActive returns the ListIsActive field value if set, zero value otherwise.
func (o *List3) GetListIsActive() bool {
	if o == nil || IsNil(o.ListIsActive) {
		var ret bool
		return ret
	}
	return *o.ListIsActive
}

// GetListIsActiveOk returns a tuple with the ListIsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *List3) GetListIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.ListIsActive) {
		return nil, false
	}
	return o.ListIsActive, true
}

// HasListIsActive returns a boolean if a field has been set.
func (o *List3) HasListIsActive() bool {
	if o != nil && !IsNil(o.ListIsActive) {
		return true
	}

	return false
}

// SetListIsActive gets a reference to the given bool and assigns it to the ListIsActive field.
func (o *List3) SetListIsActive(v bool) {
	o.ListIsActive = &v
}

// GetListName returns the ListName field value if set, zero value otherwise.
func (o *List3) GetListName() string {
	if o == nil || IsNil(o.ListName) {
		var ret string
		return ret
	}
	return *o.ListName
}

// GetListNameOk returns a tuple with the ListName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *List3) GetListNameOk() (*string, bool) {
	if o == nil || IsNil(o.ListName) {
		return nil, false
	}
	return o.ListName, true
}

// HasListName returns a boolean if a field has been set.
func (o *List3) HasListName() bool {
	if o != nil && !IsNil(o.ListName) {
		return true
	}

	return false
}

// SetListName gets a reference to the given string and assigns it to the ListName field.
func (o *List3) SetListName(v string) {
	o.ListName = &v
}

// GetSegmentText returns the SegmentText field value if set, zero value otherwise.
func (o *List3) GetSegmentText() string {
	if o == nil || IsNil(o.SegmentText) {
		var ret string
		return ret
	}
	return *o.SegmentText
}

// GetSegmentTextOk returns a tuple with the SegmentText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *List3) GetSegmentTextOk() (*string, bool) {
	if o == nil || IsNil(o.SegmentText) {
		return nil, false
	}
	return o.SegmentText, true
}

// HasSegmentText returns a boolean if a field has been set.
func (o *List3) HasSegmentText() bool {
	if o != nil && !IsNil(o.SegmentText) {
		return true
	}

	return false
}

// SetSegmentText gets a reference to the given string and assigns it to the SegmentText field.
func (o *List3) SetSegmentText(v string) {
	o.SegmentText = &v
}

// GetRecipientCount returns the RecipientCount field value if set, zero value otherwise.
func (o *List3) GetRecipientCount() int32 {
	if o == nil || IsNil(o.RecipientCount) {
		var ret int32
		return ret
	}
	return *o.RecipientCount
}

// GetRecipientCountOk returns a tuple with the RecipientCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *List3) GetRecipientCountOk() (*int32, bool) {
	if o == nil || IsNil(o.RecipientCount) {
		return nil, false
	}
	return o.RecipientCount, true
}

// HasRecipientCount returns a boolean if a field has been set.
func (o *List3) HasRecipientCount() bool {
	if o != nil && !IsNil(o.RecipientCount) {
		return true
	}

	return false
}

// SetRecipientCount gets a reference to the given int32 and assigns it to the RecipientCount field.
func (o *List3) SetRecipientCount(v int32) {
	o.RecipientCount = &v
}

// GetSegmentOpts returns the SegmentOpts field value if set, zero value otherwise.
func (o *List3) GetSegmentOpts() SegmentOptions1 {
	if o == nil || IsNil(o.SegmentOpts) {
		var ret SegmentOptions1
		return ret
	}
	return *o.SegmentOpts
}

// GetSegmentOptsOk returns a tuple with the SegmentOpts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *List3) GetSegmentOptsOk() (*SegmentOptions1, bool) {
	if o == nil || IsNil(o.SegmentOpts) {
		return nil, false
	}
	return o.SegmentOpts, true
}

// HasSegmentOpts returns a boolean if a field has been set.
func (o *List3) HasSegmentOpts() bool {
	if o != nil && !IsNil(o.SegmentOpts) {
		return true
	}

	return false
}

// SetSegmentOpts gets a reference to the given SegmentOptions1 and assigns it to the SegmentOpts field.
func (o *List3) SetSegmentOpts(v SegmentOptions1) {
	o.SegmentOpts = &v
}

func (o List3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o List3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ListId) {
		toSerialize["list_id"] = o.ListId
	}
	if !IsNil(o.ListIsActive) {
		toSerialize["list_is_active"] = o.ListIsActive
	}
	if !IsNil(o.ListName) {
		toSerialize["list_name"] = o.ListName
	}
	if !IsNil(o.SegmentText) {
		toSerialize["segment_text"] = o.SegmentText
	}
	if !IsNil(o.RecipientCount) {
		toSerialize["recipient_count"] = o.RecipientCount
	}
	if !IsNil(o.SegmentOpts) {
		toSerialize["segment_opts"] = o.SegmentOpts
	}
	return toSerialize, nil
}

type NullableList3 struct {
	value *List3
	isSet bool
}

func (v NullableList3) Get() *List3 {
	return v.value
}

func (v *NullableList3) Set(val *List3) {
	v.value = val
	v.isSet = true
}

func (v NullableList3) IsSet() bool {
	return v.isSet
}

func (v *NullableList3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableList3(val *List3) *NullableList3 {
	return &NullableList3{value: val, isSet: true}
}

func (v NullableList3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableList3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


