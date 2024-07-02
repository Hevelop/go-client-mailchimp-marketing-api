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

// checks if the ConversationMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConversationMessage{}

// ConversationMessage An individual message in a conversation. Conversation tracking is a feature available to paid accounts that lets you view replies to your campaigns in your Mailchimp account.
type ConversationMessage struct {
	// A string that uniquely identifies this message
	Id *string `json:"id,omitempty"`
	// A string that identifies this message's conversation
	ConversationId *string `json:"conversation_id,omitempty"`
	// The list's web ID
	ListId *int32 `json:"list_id,omitempty"`
	// A label representing the sender of this message
	FromLabel *string `json:"from_label,omitempty"`
	// A label representing the email of the sender of this message
	FromEmail *string `json:"from_email,omitempty"`
	// The subject of this message
	Subject *string `json:"subject,omitempty"`
	// The plain-text content of the message
	Message *string `json:"message,omitempty"`
	// Whether this message has been marked as read
	Read *bool `json:"read,omitempty"`
	// The date and time the message was either sent or received in ISO 8601 format.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewConversationMessage instantiates a new ConversationMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConversationMessage() *ConversationMessage {
	this := ConversationMessage{}
	return &this
}

// NewConversationMessageWithDefaults instantiates a new ConversationMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConversationMessageWithDefaults() *ConversationMessage {
	this := ConversationMessage{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConversationMessage) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversationMessage) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// ContainsId returns a boolean if a field has been set.
func (o *ConversationMessage) ContainsId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConversationMessage) SetId(v string) {
	o.Id = &v
}

// GetConversationId returns the ConversationId field value if set, zero value otherwise.
func (o *ConversationMessage) GetConversationId() string {
	if o == nil || IsNil(o.ConversationId) {
		var ret string
		return ret
	}
	return *o.ConversationId
}

// GetConversationIdOk returns a tuple with the ConversationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversationMessage) GetConversationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConversationId) {
		return nil, false
	}
	return o.ConversationId, true
}

// ContainsConversationId returns a boolean if a field has been set.
func (o *ConversationMessage) ContainsConversationId() bool {
	if o != nil && !IsNil(o.ConversationId) {
		return true
	}

	return false
}

// SetConversationId gets a reference to the given string and assigns it to the ConversationId field.
func (o *ConversationMessage) SetConversationId(v string) {
	o.ConversationId = &v
}

// GetListId returns the ListId field value if set, zero value otherwise.
func (o *ConversationMessage) GetListId() int32 {
	if o == nil || IsNil(o.ListId) {
		var ret int32
		return ret
	}
	return *o.ListId
}

// GetListIdOk returns a tuple with the ListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversationMessage) GetListIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ListId) {
		return nil, false
	}
	return o.ListId, true
}

// ContainsListId returns a boolean if a field has been set.
func (o *ConversationMessage) ContainsListId() bool {
	if o != nil && !IsNil(o.ListId) {
		return true
	}

	return false
}

// SetListId gets a reference to the given int32 and assigns it to the ListId field.
func (o *ConversationMessage) SetListId(v int32) {
	o.ListId = &v
}

// GetFromLabel returns the FromLabel field value if set, zero value otherwise.
func (o *ConversationMessage) GetFromLabel() string {
	if o == nil || IsNil(o.FromLabel) {
		var ret string
		return ret
	}
	return *o.FromLabel
}

// GetFromLabelOk returns a tuple with the FromLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversationMessage) GetFromLabelOk() (*string, bool) {
	if o == nil || IsNil(o.FromLabel) {
		return nil, false
	}
	return o.FromLabel, true
}

// ContainsFromLabel returns a boolean if a field has been set.
func (o *ConversationMessage) ContainsFromLabel() bool {
	if o != nil && !IsNil(o.FromLabel) {
		return true
	}

	return false
}

// SetFromLabel gets a reference to the given string and assigns it to the FromLabel field.
func (o *ConversationMessage) SetFromLabel(v string) {
	o.FromLabel = &v
}

// GetFromEmail returns the FromEmail field value if set, zero value otherwise.
func (o *ConversationMessage) GetFromEmail() string {
	if o == nil || IsNil(o.FromEmail) {
		var ret string
		return ret
	}
	return *o.FromEmail
}

// GetFromEmailOk returns a tuple with the FromEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversationMessage) GetFromEmailOk() (*string, bool) {
	if o == nil || IsNil(o.FromEmail) {
		return nil, false
	}
	return o.FromEmail, true
}

// ContainsFromEmail returns a boolean if a field has been set.
func (o *ConversationMessage) ContainsFromEmail() bool {
	if o != nil && !IsNil(o.FromEmail) {
		return true
	}

	return false
}

// SetFromEmail gets a reference to the given string and assigns it to the FromEmail field.
func (o *ConversationMessage) SetFromEmail(v string) {
	o.FromEmail = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *ConversationMessage) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversationMessage) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// ContainsSubject returns a boolean if a field has been set.
func (o *ConversationMessage) ContainsSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *ConversationMessage) SetSubject(v string) {
	o.Subject = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ConversationMessage) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversationMessage) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// ContainsMessage returns a boolean if a field has been set.
func (o *ConversationMessage) ContainsMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ConversationMessage) SetMessage(v string) {
	o.Message = &v
}

// GetRead returns the Read field value if set, zero value otherwise.
func (o *ConversationMessage) GetRead() bool {
	if o == nil || IsNil(o.Read) {
		var ret bool
		return ret
	}
	return *o.Read
}

// GetReadOk returns a tuple with the Read field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversationMessage) GetReadOk() (*bool, bool) {
	if o == nil || IsNil(o.Read) {
		return nil, false
	}
	return o.Read, true
}

// ContainsRead returns a boolean if a field has been set.
func (o *ConversationMessage) ContainsRead() bool {
	if o != nil && !IsNil(o.Read) {
		return true
	}

	return false
}

// SetRead gets a reference to the given bool and assigns it to the Read field.
func (o *ConversationMessage) SetRead(v bool) {
	o.Read = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ConversationMessage) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversationMessage) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// ContainsTimestamp returns a boolean if a field has been set.
func (o *ConversationMessage) ContainsTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *ConversationMessage) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ConversationMessage) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversationMessage) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// ContainsLinks returns a boolean if a field has been set.
func (o *ConversationMessage) ContainsLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *ConversationMessage) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o ConversationMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConversationMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ConversationId) {
		toSerialize["conversation_id"] = o.ConversationId
	}
	if !IsNil(o.ListId) {
		toSerialize["list_id"] = o.ListId
	}
	if !IsNil(o.FromLabel) {
		toSerialize["from_label"] = o.FromLabel
	}
	if !IsNil(o.FromEmail) {
		toSerialize["from_email"] = o.FromEmail
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Read) {
		toSerialize["read"] = o.Read
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableConversationMessage struct {
	value *ConversationMessage
	isSet bool
}

func (v NullableConversationMessage) Get() *ConversationMessage {
	return v.value
}

func (v *NullableConversationMessage) Set(val *ConversationMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableConversationMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableConversationMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConversationMessage(val *ConversationMessage) *NullableConversationMessage {
	return &NullableConversationMessage{value: val, isSet: true}
}

func (v NullableConversationMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConversationMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


