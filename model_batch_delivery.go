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

// checks if the BatchDelivery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchDelivery{}

// BatchDelivery Choose whether the campaign should use [Batch Delivery](https://mailchimp.com/help/schedule-batch-delivery/). Cannot be set to `true` for campaigns using [Timewarp](https://mailchimp.com/help/use-timewarp/).
type BatchDelivery struct {
	// The delay, in minutes, between batches.
	BatchDelay int32 `json:"batch_delay"`
	// The number of batches for the campaign send.
	BatchCount int32 `json:"batch_count"`
}

type _BatchDelivery BatchDelivery

// NewBatchDelivery instantiates a new BatchDelivery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchDelivery(batchDelay int32, batchCount int32) *BatchDelivery {
	this := BatchDelivery{}
	this.BatchDelay = batchDelay
	this.BatchCount = batchCount
	return &this
}

// NewBatchDeliveryWithDefaults instantiates a new BatchDelivery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchDeliveryWithDefaults() *BatchDelivery {
	this := BatchDelivery{}
	return &this
}

// GetBatchDelay returns the BatchDelay field value
func (o *BatchDelivery) GetBatchDelay() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BatchDelay
}

// GetBatchDelayOk returns a tuple with the BatchDelay field value
// and a boolean to check if the value has been set.
func (o *BatchDelivery) GetBatchDelayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BatchDelay, true
}

// SetBatchDelay sets field value
func (o *BatchDelivery) SetBatchDelay(v int32) {
	o.BatchDelay = v
}

// GetBatchCount returns the BatchCount field value
func (o *BatchDelivery) GetBatchCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BatchCount
}

// GetBatchCountOk returns a tuple with the BatchCount field value
// and a boolean to check if the value has been set.
func (o *BatchDelivery) GetBatchCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BatchCount, true
}

// SetBatchCount sets field value
func (o *BatchDelivery) SetBatchCount(v int32) {
	o.BatchCount = v
}

func (o BatchDelivery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchDelivery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["batch_delay"] = o.BatchDelay
	toSerialize["batch_count"] = o.BatchCount
	return toSerialize, nil
}

func (o *BatchDelivery) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"batch_delay",
		"batch_count",
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

	varBatchDelivery := _BatchDelivery{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBatchDelivery)

	if err != nil {
		return err
	}

	*o = BatchDelivery(varBatchDelivery)

	return err
}

type NullableBatchDelivery struct {
	value *BatchDelivery
	isSet bool
}

func (v NullableBatchDelivery) Get() *BatchDelivery {
	return v.value
}

func (v *NullableBatchDelivery) Set(val *BatchDelivery) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchDelivery) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchDelivery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchDelivery(val *BatchDelivery) *NullableBatchDelivery {
	return &NullableBatchDelivery{value: val, isSet: true}
}

func (v NullableBatchDelivery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchDelivery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

