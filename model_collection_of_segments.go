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

// checks if the CollectionOfSegments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionOfSegments{}

// CollectionOfSegments A list of available segments.
type CollectionOfSegments struct {
	// An array of objects, each representing a list segment.
	Segments []List7 `json:"segments,omitempty"`
	// The list id.
	ListId *string `json:"list_id,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int32 `json:"total_items,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewCollectionOfSegments instantiates a new CollectionOfSegments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfSegments() *CollectionOfSegments {
	this := CollectionOfSegments{}
	return &this
}

// NewCollectionOfSegmentsWithDefaults instantiates a new CollectionOfSegments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfSegmentsWithDefaults() *CollectionOfSegments {
	this := CollectionOfSegments{}
	return &this
}

// GetSegments returns the Segments field value if set, zero value otherwise.
func (o *CollectionOfSegments) GetSegments() []List7 {
	if o == nil || IsNil(o.Segments) {
		var ret []List7
		return ret
	}
	return o.Segments
}

// GetSegmentsOk returns a tuple with the Segments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfSegments) GetSegmentsOk() ([]List7, bool) {
	if o == nil || IsNil(o.Segments) {
		return nil, false
	}
	return o.Segments, true
}

// HasSegments returns a boolean if a field has been set.
func (o *CollectionOfSegments) HasSegments() bool {
	if o != nil && !IsNil(o.Segments) {
		return true
	}

	return false
}

// SetSegments gets a reference to the given []List7 and assigns it to the Segments field.
func (o *CollectionOfSegments) SetSegments(v []List7) {
	o.Segments = v
}

// GetListId returns the ListId field value if set, zero value otherwise.
func (o *CollectionOfSegments) GetListId() string {
	if o == nil || IsNil(o.ListId) {
		var ret string
		return ret
	}
	return *o.ListId
}

// GetListIdOk returns a tuple with the ListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfSegments) GetListIdOk() (*string, bool) {
	if o == nil || IsNil(o.ListId) {
		return nil, false
	}
	return o.ListId, true
}

// HasListId returns a boolean if a field has been set.
func (o *CollectionOfSegments) HasListId() bool {
	if o != nil && !IsNil(o.ListId) {
		return true
	}

	return false
}

// SetListId gets a reference to the given string and assigns it to the ListId field.
func (o *CollectionOfSegments) SetListId(v string) {
	o.ListId = &v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *CollectionOfSegments) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfSegments) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *CollectionOfSegments) HasTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *CollectionOfSegments) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CollectionOfSegments) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfSegments) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CollectionOfSegments) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *CollectionOfSegments) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o CollectionOfSegments) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionOfSegments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Segments) {
		toSerialize["segments"] = o.Segments
	}
	if !IsNil(o.ListId) {
		toSerialize["list_id"] = o.ListId
	}
	if !IsNil(o.TotalItems) {
		toSerialize["total_items"] = o.TotalItems
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableCollectionOfSegments struct {
	value *CollectionOfSegments
	isSet bool
}

func (v NullableCollectionOfSegments) Get() *CollectionOfSegments {
	return v.value
}

func (v *NullableCollectionOfSegments) Set(val *CollectionOfSegments) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfSegments) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfSegments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfSegments(val *CollectionOfSegments) *NullableCollectionOfSegments {
	return &NullableCollectionOfSegments{value: val, isSet: true}
}

func (v NullableCollectionOfSegments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfSegments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


