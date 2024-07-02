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

// checks if the Conversation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Conversation{}

// Conversation Details about an individual conversation. Conversation tracking is a feature available to paid accounts that lets you view replies to your campaigns in your Mailchimp account.
type Conversation struct {
	// A string that uniquely identifies this conversation.
	Id *string `json:"id,omitempty"`
	// The total number of messages in this conversation.
	MessageCount *int32 `json:"message_count,omitempty"`
	// The unique identifier of the campaign for this conversation.
	CampaignId *string `json:"campaign_id,omitempty"`
	// The unique identifier of the list for this conversation.
	ListId *string `json:"list_id,omitempty"`
	// The number of unread messages in this conversation.
	UnreadMessages *int32 `json:"unread_messages,omitempty"`
	// A label representing the sender of this message.
	FromLabel *string `json:"from_label,omitempty"`
	// A label representing the email of the sender of this message.
	FromEmail *string `json:"from_email,omitempty"`
	// The subject of the message.
	Subject *string `json:"subject,omitempty"`
	LastMessage *LastMessage `json:"last_message,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewConversation instantiates a new Conversation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConversation() *Conversation {
	this := Conversation{}
	return &this
}

// NewConversationWithDefaults instantiates a new Conversation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConversationWithDefaults() *Conversation {
	this := Conversation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Conversation) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Conversation) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Conversation) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Conversation) SetId(v string) {
	o.Id = &v
}

// GetMessageCount returns the MessageCount field value if set, zero value otherwise.
func (o *Conversation) GetMessageCount() int32 {
	if o == nil || IsNil(o.MessageCount) {
		var ret int32
		return ret
	}
	return *o.MessageCount
}

// GetMessageCountOk returns a tuple with the MessageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Conversation) GetMessageCountOk() (*int32, bool) {
	if o == nil || IsNil(o.MessageCount) {
		return nil, false
	}
	return o.MessageCount, true
}

// HasMessageCount returns a boolean if a field has been set.
func (o *Conversation) HasMessageCount() bool {
	if o != nil && !IsNil(o.MessageCount) {
		return true
	}

	return false
}

// SetMessageCount gets a reference to the given int32 and assigns it to the MessageCount field.
func (o *Conversation) SetMessageCount(v int32) {
	o.MessageCount = &v
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *Conversation) GetCampaignId() string {
	if o == nil || IsNil(o.CampaignId) {
		var ret string
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Conversation) GetCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *Conversation) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given string and assigns it to the CampaignId field.
func (o *Conversation) SetCampaignId(v string) {
	o.CampaignId = &v
}

// GetListId returns the ListId field value if set, zero value otherwise.
func (o *Conversation) GetListId() string {
	if o == nil || IsNil(o.ListId) {
		var ret string
		return ret
	}
	return *o.ListId
}

// GetListIdOk returns a tuple with the ListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Conversation) GetListIdOk() (*string, bool) {
	if o == nil || IsNil(o.ListId) {
		return nil, false
	}
	return o.ListId, true
}

// HasListId returns a boolean if a field has been set.
func (o *Conversation) HasListId() bool {
	if o != nil && !IsNil(o.ListId) {
		return true
	}

	return false
}

// SetListId gets a reference to the given string and assigns it to the ListId field.
func (o *Conversation) SetListId(v string) {
	o.ListId = &v
}

// GetUnreadMessages returns the UnreadMessages field value if set, zero value otherwise.
func (o *Conversation) GetUnreadMessages() int32 {
	if o == nil || IsNil(o.UnreadMessages) {
		var ret int32
		return ret
	}
	return *o.UnreadMessages
}

// GetUnreadMessagesOk returns a tuple with the UnreadMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Conversation) GetUnreadMessagesOk() (*int32, bool) {
	if o == nil || IsNil(o.UnreadMessages) {
		return nil, false
	}
	return o.UnreadMessages, true
}

// HasUnreadMessages returns a boolean if a field has been set.
func (o *Conversation) HasUnreadMessages() bool {
	if o != nil && !IsNil(o.UnreadMessages) {
		return true
	}

	return false
}

// SetUnreadMessages gets a reference to the given int32 and assigns it to the UnreadMessages field.
func (o *Conversation) SetUnreadMessages(v int32) {
	o.UnreadMessages = &v
}

// GetFromLabel returns the FromLabel field value if set, zero value otherwise.
func (o *Conversation) GetFromLabel() string {
	if o == nil || IsNil(o.FromLabel) {
		var ret string
		return ret
	}
	return *o.FromLabel
}

// GetFromLabelOk returns a tuple with the FromLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Conversation) GetFromLabelOk() (*string, bool) {
	if o == nil || IsNil(o.FromLabel) {
		return nil, false
	}
	return o.FromLabel, true
}

// HasFromLabel returns a boolean if a field has been set.
func (o *Conversation) HasFromLabel() bool {
	if o != nil && !IsNil(o.FromLabel) {
		return true
	}

	return false
}

// SetFromLabel gets a reference to the given string and assigns it to the FromLabel field.
func (o *Conversation) SetFromLabel(v string) {
	o.FromLabel = &v
}

// GetFromEmail returns the FromEmail field value if set, zero value otherwise.
func (o *Conversation) GetFromEmail() string {
	if o == nil || IsNil(o.FromEmail) {
		var ret string
		return ret
	}
	return *o.FromEmail
}

// GetFromEmailOk returns a tuple with the FromEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Conversation) GetFromEmailOk() (*string, bool) {
	if o == nil || IsNil(o.FromEmail) {
		return nil, false
	}
	return o.FromEmail, true
}

// HasFromEmail returns a boolean if a field has been set.
func (o *Conversation) HasFromEmail() bool {
	if o != nil && !IsNil(o.FromEmail) {
		return true
	}

	return false
}

// SetFromEmail gets a reference to the given string and assigns it to the FromEmail field.
func (o *Conversation) SetFromEmail(v string) {
	o.FromEmail = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *Conversation) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Conversation) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *Conversation) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *Conversation) SetSubject(v string) {
	o.Subject = &v
}

// GetLastMessage returns the LastMessage field value if set, zero value otherwise.
func (o *Conversation) GetLastMessage() LastMessage {
	if o == nil || IsNil(o.LastMessage) {
		var ret LastMessage
		return ret
	}
	return *o.LastMessage
}

// GetLastMessageOk returns a tuple with the LastMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Conversation) GetLastMessageOk() (*LastMessage, bool) {
	if o == nil || IsNil(o.LastMessage) {
		return nil, false
	}
	return o.LastMessage, true
}

// HasLastMessage returns a boolean if a field has been set.
func (o *Conversation) HasLastMessage() bool {
	if o != nil && !IsNil(o.LastMessage) {
		return true
	}

	return false
}

// SetLastMessage gets a reference to the given LastMessage and assigns it to the LastMessage field.
func (o *Conversation) SetLastMessage(v LastMessage) {
	o.LastMessage = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Conversation) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Conversation) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Conversation) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *Conversation) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o Conversation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Conversation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MessageCount) {
		toSerialize["message_count"] = o.MessageCount
	}
	if !IsNil(o.CampaignId) {
		toSerialize["campaign_id"] = o.CampaignId
	}
	if !IsNil(o.ListId) {
		toSerialize["list_id"] = o.ListId
	}
	if !IsNil(o.UnreadMessages) {
		toSerialize["unread_messages"] = o.UnreadMessages
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
	if !IsNil(o.LastMessage) {
		toSerialize["last_message"] = o.LastMessage
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableConversation struct {
	value *Conversation
	isSet bool
}

func (v NullableConversation) Get() *Conversation {
	return v.value
}

func (v *NullableConversation) Set(val *Conversation) {
	v.value = val
	v.isSet = true
}

func (v NullableConversation) IsSet() bool {
	return v.isSet
}

func (v *NullableConversation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConversation(val *Conversation) *NullableConversation {
	return &NullableConversation{value: val, isSet: true}
}

func (v NullableConversation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConversation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


