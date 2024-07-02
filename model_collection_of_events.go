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

// checks if the CollectionOfEvents type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionOfEvents{}

// CollectionOfEvents A collection of events for a given contact
type CollectionOfEvents struct {
	// An array of objects, each representing an event.
	Events []Event `json:"events,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int32 `json:"total_items,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewCollectionOfEvents instantiates a new CollectionOfEvents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfEvents() *CollectionOfEvents {
	this := CollectionOfEvents{}
	return &this
}

// NewCollectionOfEventsWithDefaults instantiates a new CollectionOfEvents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfEventsWithDefaults() *CollectionOfEvents {
	this := CollectionOfEvents{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *CollectionOfEvents) GetEvents() []Event {
	if o == nil || IsNil(o.Events) {
		var ret []Event
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfEvents) GetEventsOk() ([]Event, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// ContainsEvents returns a boolean if a field has been set.
func (o *CollectionOfEvents) ContainsEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []Event and assigns it to the Events field.
func (o *CollectionOfEvents) SetEvents(v []Event) {
	o.Events = v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *CollectionOfEvents) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfEvents) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// ContainsTotalItems returns a boolean if a field has been set.
func (o *CollectionOfEvents) ContainsTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *CollectionOfEvents) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CollectionOfEvents) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfEvents) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// ContainsLinks returns a boolean if a field has been set.
func (o *CollectionOfEvents) ContainsLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *CollectionOfEvents) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o CollectionOfEvents) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionOfEvents) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if !IsNil(o.TotalItems) {
		toSerialize["total_items"] = o.TotalItems
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableCollectionOfEvents struct {
	value *CollectionOfEvents
	isSet bool
}

func (v NullableCollectionOfEvents) Get() *CollectionOfEvents {
	return v.value
}

func (v *NullableCollectionOfEvents) Set(val *CollectionOfEvents) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfEvents) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfEvents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfEvents(val *CollectionOfEvents) *NullableCollectionOfEvents {
	return &NullableCollectionOfEvents{value: val, isSet: true}
}

func (v NullableCollectionOfEvents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfEvents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


