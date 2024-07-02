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

// checks if the ListWebhooks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListWebhooks{}

// ListWebhooks Manage webhooks for a specific list.
type ListWebhooks struct {
	// An array of objects, each representing a specific list member.
	Webhooks []ListWebhooks `json:"webhooks,omitempty"`
	// The list id.
	ListId *string `json:"list_id,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int32 `json:"total_items,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewListWebhooks instantiates a new ListWebhooks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWebhooks() *ListWebhooks {
	this := ListWebhooks{}
	return &this
}

// NewListWebhooksWithDefaults instantiates a new ListWebhooks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWebhooksWithDefaults() *ListWebhooks {
	this := ListWebhooks{}
	return &this
}

// GetWebhooks returns the Webhooks field value if set, zero value otherwise.
func (o *ListWebhooks) GetWebhooks() []ListWebhooks {
	if o == nil || IsNil(o.Webhooks) {
		var ret []ListWebhooks
		return ret
	}
	return o.Webhooks
}

// GetWebhooksOk returns a tuple with the Webhooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWebhooks) GetWebhooksOk() ([]ListWebhooks, bool) {
	if o == nil || IsNil(o.Webhooks) {
		return nil, false
	}
	return o.Webhooks, true
}

// ContainsWebhooks returns a boolean if a field has been set.
func (o *ListWebhooks) ContainsWebhooks() bool {
	if o != nil && !IsNil(o.Webhooks) {
		return true
	}

	return false
}

// SetWebhooks gets a reference to the given []ListWebhooks and assigns it to the Webhooks field.
func (o *ListWebhooks) SetWebhooks(v []ListWebhooks) {
	o.Webhooks = v
}

// GetListId returns the ListId field value if set, zero value otherwise.
func (o *ListWebhooks) GetListId() string {
	if o == nil || IsNil(o.ListId) {
		var ret string
		return ret
	}
	return *o.ListId
}

// GetListIdOk returns a tuple with the ListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWebhooks) GetListIdOk() (*string, bool) {
	if o == nil || IsNil(o.ListId) {
		return nil, false
	}
	return o.ListId, true
}

// ContainsListId returns a boolean if a field has been set.
func (o *ListWebhooks) ContainsListId() bool {
	if o != nil && !IsNil(o.ListId) {
		return true
	}

	return false
}

// SetListId gets a reference to the given string and assigns it to the ListId field.
func (o *ListWebhooks) SetListId(v string) {
	o.ListId = &v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *ListWebhooks) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWebhooks) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// ContainsTotalItems returns a boolean if a field has been set.
func (o *ListWebhooks) ContainsTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *ListWebhooks) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ListWebhooks) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWebhooks) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// ContainsLinks returns a boolean if a field has been set.
func (o *ListWebhooks) ContainsLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *ListWebhooks) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o ListWebhooks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListWebhooks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Webhooks) {
		toSerialize["webhooks"] = o.Webhooks
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

type NullableListWebhooks struct {
	value *ListWebhooks
	isSet bool
}

func (v NullableListWebhooks) Get() *ListWebhooks {
	return v.value
}

func (v *NullableListWebhooks) Set(val *ListWebhooks) {
	v.value = val
	v.isSet = true
}

func (v NullableListWebhooks) IsSet() bool {
	return v.isSet
}

func (v *NullableListWebhooks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListWebhooks(val *ListWebhooks) *NullableListWebhooks {
	return &NullableListWebhooks{value: val, isSet: true}
}

func (v NullableListWebhooks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListWebhooks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


