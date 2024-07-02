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

// checks if the CollectionOfConversationMessages type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionOfConversationMessages{}

// CollectionOfConversationMessages Messages from a specific conversation.
type CollectionOfConversationMessages struct {
	// An array of objects, each representing a conversation messages resources.
	ConversationMessages []ConversationMessage `json:"conversation_messages,omitempty"`
	// A string that identifies this conversation.
	ConversationId *string `json:"conversation_id,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int32 `json:"total_items,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewCollectionOfConversationMessages instantiates a new CollectionOfConversationMessages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfConversationMessages() *CollectionOfConversationMessages {
	this := CollectionOfConversationMessages{}
	return &this
}

// NewCollectionOfConversationMessagesWithDefaults instantiates a new CollectionOfConversationMessages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfConversationMessagesWithDefaults() *CollectionOfConversationMessages {
	this := CollectionOfConversationMessages{}
	return &this
}

// GetConversationMessages returns the ConversationMessages field value if set, zero value otherwise.
func (o *CollectionOfConversationMessages) GetConversationMessages() []ConversationMessage {
	if o == nil || IsNil(o.ConversationMessages) {
		var ret []ConversationMessage
		return ret
	}
	return o.ConversationMessages
}

// GetConversationMessagesOk returns a tuple with the ConversationMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfConversationMessages) GetConversationMessagesOk() ([]ConversationMessage, bool) {
	if o == nil || IsNil(o.ConversationMessages) {
		return nil, false
	}
	return o.ConversationMessages, true
}

// HaveConversationMessages returns a boolean if a field has been set.
func (o *CollectionOfConversationMessages) HaveConversationMessages() bool {
	if o != nil && !IsNil(o.ConversationMessages) {
		return true
	}

	return false
}

// SetConversationMessages gets a reference to the given []ConversationMessage and assigns it to the ConversationMessages field.
func (o *CollectionOfConversationMessages) SetConversationMessages(v []ConversationMessage) {
	o.ConversationMessages = v
}

// GetConversationId returns the ConversationId field value if set, zero value otherwise.
func (o *CollectionOfConversationMessages) GetConversationId() string {
	if o == nil || IsNil(o.ConversationId) {
		var ret string
		return ret
	}
	return *o.ConversationId
}

// GetConversationIdOk returns a tuple with the ConversationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfConversationMessages) GetConversationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConversationId) {
		return nil, false
	}
	return o.ConversationId, true
}

// HaveConversationId returns a boolean if a field has been set.
func (o *CollectionOfConversationMessages) HaveConversationId() bool {
	if o != nil && !IsNil(o.ConversationId) {
		return true
	}

	return false
}

// SetConversationId gets a reference to the given string and assigns it to the ConversationId field.
func (o *CollectionOfConversationMessages) SetConversationId(v string) {
	o.ConversationId = &v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *CollectionOfConversationMessages) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfConversationMessages) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HaveTotalItems returns a boolean if a field has been set.
func (o *CollectionOfConversationMessages) HaveTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *CollectionOfConversationMessages) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CollectionOfConversationMessages) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfConversationMessages) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HaveLinks returns a boolean if a field has been set.
func (o *CollectionOfConversationMessages) HaveLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *CollectionOfConversationMessages) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o CollectionOfConversationMessages) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionOfConversationMessages) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConversationMessages) {
		toSerialize["conversation_messages"] = o.ConversationMessages
	}
	if !IsNil(o.ConversationId) {
		toSerialize["conversation_id"] = o.ConversationId
	}
	if !IsNil(o.TotalItems) {
		toSerialize["total_items"] = o.TotalItems
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableCollectionOfConversationMessages struct {
	value *CollectionOfConversationMessages
	isSet bool
}

func (v NullableCollectionOfConversationMessages) Get() *CollectionOfConversationMessages {
	return v.value
}

func (v *NullableCollectionOfConversationMessages) Set(val *CollectionOfConversationMessages) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfConversationMessages) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfConversationMessages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfConversationMessages(val *CollectionOfConversationMessages) *NullableCollectionOfConversationMessages {
	return &NullableCollectionOfConversationMessages{value: val, isSet: true}
}

func (v NullableCollectionOfConversationMessages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfConversationMessages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


