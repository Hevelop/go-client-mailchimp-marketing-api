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

// checks if the BatchOperations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchOperations{}

// BatchOperations A summary of batch requests that have been made.
type BatchOperations struct {
	// An array of objects representing batch calls.
	Batches []Batch `json:"batches,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int32 `json:"total_items,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewBatchOperations instantiates a new BatchOperations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchOperations() *BatchOperations {
	this := BatchOperations{}
	return &this
}

// NewBatchOperationsWithDefaults instantiates a new BatchOperations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchOperationsWithDefaults() *BatchOperations {
	this := BatchOperations{}
	return &this
}

// GetBatches returns the Batches field value if set, zero value otherwise.
func (o *BatchOperations) GetBatches() []Batch {
	if o == nil || IsNil(o.Batches) {
		var ret []Batch
		return ret
	}
	return o.Batches
}

// GetBatchesOk returns a tuple with the Batches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchOperations) GetBatchesOk() ([]Batch, bool) {
	if o == nil || IsNil(o.Batches) {
		return nil, false
	}
	return o.Batches, true
}

// HasBatches returns a boolean if a field has been set.
func (o *BatchOperations) HasBatches() bool {
	if o != nil && !IsNil(o.Batches) {
		return true
	}

	return false
}

// SetBatches gets a reference to the given []Batch and assigns it to the Batches field.
func (o *BatchOperations) SetBatches(v []Batch) {
	o.Batches = v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *BatchOperations) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchOperations) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *BatchOperations) HasTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *BatchOperations) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BatchOperations) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchOperations) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BatchOperations) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *BatchOperations) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o BatchOperations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchOperations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Batches) {
		toSerialize["batches"] = o.Batches
	}
	if !IsNil(o.TotalItems) {
		toSerialize["total_items"] = o.TotalItems
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableBatchOperations struct {
	value *BatchOperations
	isSet bool
}

func (v NullableBatchOperations) Get() *BatchOperations {
	return v.value
}

func (v *NullableBatchOperations) Set(val *BatchOperations) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchOperations) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchOperations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchOperations(val *BatchOperations) *NullableBatchOperations {
	return &NullableBatchOperations{value: val, isSet: true}
}

func (v NullableBatchOperations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchOperations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


