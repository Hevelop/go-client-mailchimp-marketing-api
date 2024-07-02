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

// checks if the Goal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Goal{}

// Goal A single instance of a goal activity.
type Goal struct {
	// The id for a Goal event.
	GoalId *int32 `json:"goal_id,omitempty"`
	// The name/type of Goal event triggered.
	Event *string `json:"event,omitempty"`
	// The date and time the user last triggered the Goal event in ISO 8601 format.
	LastVisitedAt *time.Time `json:"last_visited_at,omitempty"`
	// Any extra data passed with the Goal event.
	Data *string `json:"data,omitempty"`
}

// NewGoal instantiates a new Goal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoal() *Goal {
	this := Goal{}
	return &this
}

// NewGoalWithDefaults instantiates a new Goal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoalWithDefaults() *Goal {
	this := Goal{}
	return &this
}

// GetGoalId returns the GoalId field value if set, zero value otherwise.
func (o *Goal) GetGoalId() int32 {
	if o == nil || IsNil(o.GoalId) {
		var ret int32
		return ret
	}
	return *o.GoalId
}

// GetGoalIdOk returns a tuple with the GoalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Goal) GetGoalIdOk() (*int32, bool) {
	if o == nil || IsNil(o.GoalId) {
		return nil, false
	}
	return o.GoalId, true
}

// HasGoalId returns a boolean if a field has been set.
func (o *Goal) HasGoalId() bool {
	if o != nil && !IsNil(o.GoalId) {
		return true
	}

	return false
}

// SetGoalId gets a reference to the given int32 and assigns it to the GoalId field.
func (o *Goal) SetGoalId(v int32) {
	o.GoalId = &v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *Goal) GetEvent() string {
	if o == nil || IsNil(o.Event) {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Goal) GetEventOk() (*string, bool) {
	if o == nil || IsNil(o.Event) {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *Goal) HasEvent() bool {
	if o != nil && !IsNil(o.Event) {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *Goal) SetEvent(v string) {
	o.Event = &v
}

// GetLastVisitedAt returns the LastVisitedAt field value if set, zero value otherwise.
func (o *Goal) GetLastVisitedAt() time.Time {
	if o == nil || IsNil(o.LastVisitedAt) {
		var ret time.Time
		return ret
	}
	return *o.LastVisitedAt
}

// GetLastVisitedAtOk returns a tuple with the LastVisitedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Goal) GetLastVisitedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastVisitedAt) {
		return nil, false
	}
	return o.LastVisitedAt, true
}

// HasLastVisitedAt returns a boolean if a field has been set.
func (o *Goal) HasLastVisitedAt() bool {
	if o != nil && !IsNil(o.LastVisitedAt) {
		return true
	}

	return false
}

// SetLastVisitedAt gets a reference to the given time.Time and assigns it to the LastVisitedAt field.
func (o *Goal) SetLastVisitedAt(v time.Time) {
	o.LastVisitedAt = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Goal) GetData() string {
	if o == nil || IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Goal) GetDataOk() (*string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Goal) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *Goal) SetData(v string) {
	o.Data = &v
}

func (o Goal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Goal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GoalId) {
		toSerialize["goal_id"] = o.GoalId
	}
	if !IsNil(o.Event) {
		toSerialize["event"] = o.Event
	}
	if !IsNil(o.LastVisitedAt) {
		toSerialize["last_visited_at"] = o.LastVisitedAt
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGoal struct {
	value *Goal
	isSet bool
}

func (v NullableGoal) Get() *Goal {
	return v.value
}

func (v *NullableGoal) Set(val *Goal) {
	v.value = val
	v.isSet = true
}

func (v NullableGoal) IsSet() bool {
	return v.isSet
}

func (v *NullableGoal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoal(val *Goal) *NullableGoal {
	return &NullableGoal{value: val, isSet: true}
}

func (v NullableGoal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


