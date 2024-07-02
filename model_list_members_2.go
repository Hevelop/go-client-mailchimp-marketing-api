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

// checks if the ListMembers2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListMembers2{}

// ListMembers2 Manage members of a specific Mailchimp list, including currently subscribed, unsubscribed, and bounced members.
type ListMembers2 struct {
	// An array of objects, each representing a specific list member.
	Members []ListMembers2 `json:"members,omitempty"`
	// The list id.
	ListId *string `json:"list_id,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int32 `json:"total_items,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewListMembers2 instantiates a new ListMembers2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListMembers2() *ListMembers2 {
	this := ListMembers2{}
	return &this
}

// NewListMembers2WithDefaults instantiates a new ListMembers2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListMembers2WithDefaults() *ListMembers2 {
	this := ListMembers2{}
	return &this
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *ListMembers2) GetMembers() []ListMembers2 {
	if o == nil || IsNil(o.Members) {
		var ret []ListMembers2
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMembers2) GetMembersOk() ([]ListMembers2, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *ListMembers2) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []ListMembers2 and assigns it to the Members field.
func (o *ListMembers2) SetMembers(v []ListMembers2) {
	o.Members = v
}

// GetListId returns the ListId field value if set, zero value otherwise.
func (o *ListMembers2) GetListId() string {
	if o == nil || IsNil(o.ListId) {
		var ret string
		return ret
	}
	return *o.ListId
}

// GetListIdOk returns a tuple with the ListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMembers2) GetListIdOk() (*string, bool) {
	if o == nil || IsNil(o.ListId) {
		return nil, false
	}
	return o.ListId, true
}

// HasListId returns a boolean if a field has been set.
func (o *ListMembers2) HasListId() bool {
	if o != nil && !IsNil(o.ListId) {
		return true
	}

	return false
}

// SetListId gets a reference to the given string and assigns it to the ListId field.
func (o *ListMembers2) SetListId(v string) {
	o.ListId = &v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *ListMembers2) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMembers2) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *ListMembers2) HasTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *ListMembers2) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ListMembers2) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMembers2) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ListMembers2) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *ListMembers2) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o ListMembers2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListMembers2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Members) {
		toSerialize["members"] = o.Members
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

type NullableListMembers2 struct {
	value *ListMembers2
	isSet bool
}

func (v NullableListMembers2) Get() *ListMembers2 {
	return v.value
}

func (v *NullableListMembers2) Set(val *ListMembers2) {
	v.value = val
	v.isSet = true
}

func (v NullableListMembers2) IsSet() bool {
	return v.isSet
}

func (v *NullableListMembers2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListMembers2(val *ListMembers2) *NullableListMembers2 {
	return &NullableListMembers2{value: val, isSet: true}
}

func (v NullableListMembers2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListMembers2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


