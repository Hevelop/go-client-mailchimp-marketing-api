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

// checks if the MembersToSubscribeUnsubscribeToFromAListInBatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MembersToSubscribeUnsubscribeToFromAListInBatch{}

// MembersToSubscribeUnsubscribeToFromAListInBatch Members to subscribe to or unsubscribe from a list.
type MembersToSubscribeUnsubscribeToFromAListInBatch struct {
	// An array of objects, each representing an email address and the subscription status for a specific list. Up to 500 members may be added or updated with each API call.
	Members []AddListMembers `json:"members"`
	// Whether this batch operation will replace all existing tags with tags in request.
	SyncTags *bool `json:"sync_tags,omitempty"`
	// Whether this batch operation will change existing members' subscription status.
	UpdateExisting *bool `json:"update_existing,omitempty"`
}

type _MembersToSubscribeUnsubscribeToFromAListInBatch MembersToSubscribeUnsubscribeToFromAListInBatch

// NewMembersToSubscribeUnsubscribeToFromAListInBatch instantiates a new MembersToSubscribeUnsubscribeToFromAListInBatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMembersToSubscribeUnsubscribeToFromAListInBatch(members []AddListMembers) *MembersToSubscribeUnsubscribeToFromAListInBatch {
	this := MembersToSubscribeUnsubscribeToFromAListInBatch{}
	this.Members = members
	return &this
}

// NewMembersToSubscribeUnsubscribeToFromAListInBatchWithDefaults instantiates a new MembersToSubscribeUnsubscribeToFromAListInBatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMembersToSubscribeUnsubscribeToFromAListInBatchWithDefaults() *MembersToSubscribeUnsubscribeToFromAListInBatch {
	this := MembersToSubscribeUnsubscribeToFromAListInBatch{}
	return &this
}

// GetMembers returns the Members field value
func (o *MembersToSubscribeUnsubscribeToFromAListInBatch) GetMembers() []AddListMembers {
	if o == nil {
		var ret []AddListMembers
		return ret
	}

	return o.Members
}

// GetMembersOk returns a tuple with the Members field value
// and a boolean to check if the value has been set.
func (o *MembersToSubscribeUnsubscribeToFromAListInBatch) GetMembersOk() ([]AddListMembers, bool) {
	if o == nil {
		return nil, false
	}
	return o.Members, true
}

// SetMembers sets field value
func (o *MembersToSubscribeUnsubscribeToFromAListInBatch) SetMembers(v []AddListMembers) {
	o.Members = v
}

// GetSyncTags returns the SyncTags field value if set, zero value otherwise.
func (o *MembersToSubscribeUnsubscribeToFromAListInBatch) GetSyncTags() bool {
	if o == nil || IsNil(o.SyncTags) {
		var ret bool
		return ret
	}
	return *o.SyncTags
}

// GetSyncTagsOk returns a tuple with the SyncTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembersToSubscribeUnsubscribeToFromAListInBatch) GetSyncTagsOk() (*bool, bool) {
	if o == nil || IsNil(o.SyncTags) {
		return nil, false
	}
	return o.SyncTags, true
}

// HasSyncTags returns a boolean if a field has been set.
func (o *MembersToSubscribeUnsubscribeToFromAListInBatch) HasSyncTags() bool {
	if o != nil && !IsNil(o.SyncTags) {
		return true
	}

	return false
}

// SetSyncTags gets a reference to the given bool and assigns it to the SyncTags field.
func (o *MembersToSubscribeUnsubscribeToFromAListInBatch) SetSyncTags(v bool) {
	o.SyncTags = &v
}

// GetUpdateExisting returns the UpdateExisting field value if set, zero value otherwise.
func (o *MembersToSubscribeUnsubscribeToFromAListInBatch) GetUpdateExisting() bool {
	if o == nil || IsNil(o.UpdateExisting) {
		var ret bool
		return ret
	}
	return *o.UpdateExisting
}

// GetUpdateExistingOk returns a tuple with the UpdateExisting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembersToSubscribeUnsubscribeToFromAListInBatch) GetUpdateExistingOk() (*bool, bool) {
	if o == nil || IsNil(o.UpdateExisting) {
		return nil, false
	}
	return o.UpdateExisting, true
}

// HasUpdateExisting returns a boolean if a field has been set.
func (o *MembersToSubscribeUnsubscribeToFromAListInBatch) HasUpdateExisting() bool {
	if o != nil && !IsNil(o.UpdateExisting) {
		return true
	}

	return false
}

// SetUpdateExisting gets a reference to the given bool and assigns it to the UpdateExisting field.
func (o *MembersToSubscribeUnsubscribeToFromAListInBatch) SetUpdateExisting(v bool) {
	o.UpdateExisting = &v
}

func (o MembersToSubscribeUnsubscribeToFromAListInBatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MembersToSubscribeUnsubscribeToFromAListInBatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["members"] = o.Members
	if !IsNil(o.SyncTags) {
		toSerialize["sync_tags"] = o.SyncTags
	}
	if !IsNil(o.UpdateExisting) {
		toSerialize["update_existing"] = o.UpdateExisting
	}
	return toSerialize, nil
}

func (o *MembersToSubscribeUnsubscribeToFromAListInBatch) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"members",
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

	varMembersToSubscribeUnsubscribeToFromAListInBatch := _MembersToSubscribeUnsubscribeToFromAListInBatch{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMembersToSubscribeUnsubscribeToFromAListInBatch)

	if err != nil {
		return err
	}

	*o = MembersToSubscribeUnsubscribeToFromAListInBatch(varMembersToSubscribeUnsubscribeToFromAListInBatch)

	return err
}

type NullableMembersToSubscribeUnsubscribeToFromAListInBatch struct {
	value *MembersToSubscribeUnsubscribeToFromAListInBatch
	isSet bool
}

func (v NullableMembersToSubscribeUnsubscribeToFromAListInBatch) Get() *MembersToSubscribeUnsubscribeToFromAListInBatch {
	return v.value
}

func (v *NullableMembersToSubscribeUnsubscribeToFromAListInBatch) Set(val *MembersToSubscribeUnsubscribeToFromAListInBatch) {
	v.value = val
	v.isSet = true
}

func (v NullableMembersToSubscribeUnsubscribeToFromAListInBatch) IsSet() bool {
	return v.isSet
}

func (v *NullableMembersToSubscribeUnsubscribeToFromAListInBatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMembersToSubscribeUnsubscribeToFromAListInBatch(val *MembersToSubscribeUnsubscribeToFromAListInBatch) *NullableMembersToSubscribeUnsubscribeToFromAListInBatch {
	return &NullableMembersToSubscribeUnsubscribeToFromAListInBatch{value: val, isSet: true}
}

func (v NullableMembersToSubscribeUnsubscribeToFromAListInBatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMembersToSubscribeUnsubscribeToFromAListInBatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


