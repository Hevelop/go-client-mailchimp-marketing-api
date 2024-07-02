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

// checks if the MemberActivityEvents type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MemberActivityEvents{}

// MemberActivityEvents The last 50 member events for a list.
type MemberActivityEvents struct {
	// An array of objects, each representing a member event.
	Activity []MemberActivity `json:"activity,omitempty"`
	// The MD5 hash of the lowercase version of the list member's email address.
	EmailId *string `json:"email_id,omitempty"`
	// As Mailchimp evolves beyond email, you may eventually have contacts without email addresses. While the `email_id` is the MD5 hash of their email address, this `contact_id` is agnostic of contact’s inclusion of an email address.
	ContactId *string `json:"contact_id,omitempty"`
	// The list id.
	ListId *string `json:"list_id,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int32 `json:"total_items,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewMemberActivityEvents instantiates a new MemberActivityEvents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberActivityEvents() *MemberActivityEvents {
	this := MemberActivityEvents{}
	return &this
}

// NewMemberActivityEventsWithDefaults instantiates a new MemberActivityEvents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberActivityEventsWithDefaults() *MemberActivityEvents {
	this := MemberActivityEvents{}
	return &this
}

// GetActivity returns the Activity field value if set, zero value otherwise.
func (o *MemberActivityEvents) GetActivity() []MemberActivity {
	if o == nil || IsNil(o.Activity) {
		var ret []MemberActivity
		return ret
	}
	return o.Activity
}

// GetActivityOk returns a tuple with the Activity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberActivityEvents) GetActivityOk() ([]MemberActivity, bool) {
	if o == nil || IsNil(o.Activity) {
		return nil, false
	}
	return o.Activity, true
}

// HasActivity returns a boolean if a field has been set.
func (o *MemberActivityEvents) HasActivity() bool {
	if o != nil && !IsNil(o.Activity) {
		return true
	}

	return false
}

// SetActivity gets a reference to the given []MemberActivity and assigns it to the Activity field.
func (o *MemberActivityEvents) SetActivity(v []MemberActivity) {
	o.Activity = v
}

// GetEmailId returns the EmailId field value if set, zero value otherwise.
func (o *MemberActivityEvents) GetEmailId() string {
	if o == nil || IsNil(o.EmailId) {
		var ret string
		return ret
	}
	return *o.EmailId
}

// GetEmailIdOk returns a tuple with the EmailId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberActivityEvents) GetEmailIdOk() (*string, bool) {
	if o == nil || IsNil(o.EmailId) {
		return nil, false
	}
	return o.EmailId, true
}

// HasEmailId returns a boolean if a field has been set.
func (o *MemberActivityEvents) HasEmailId() bool {
	if o != nil && !IsNil(o.EmailId) {
		return true
	}

	return false
}

// SetEmailId gets a reference to the given string and assigns it to the EmailId field.
func (o *MemberActivityEvents) SetEmailId(v string) {
	o.EmailId = &v
}

// GetContactId returns the ContactId field value if set, zero value otherwise.
func (o *MemberActivityEvents) GetContactId() string {
	if o == nil || IsNil(o.ContactId) {
		var ret string
		return ret
	}
	return *o.ContactId
}

// GetContactIdOk returns a tuple with the ContactId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberActivityEvents) GetContactIdOk() (*string, bool) {
	if o == nil || IsNil(o.ContactId) {
		return nil, false
	}
	return o.ContactId, true
}

// HasContactId returns a boolean if a field has been set.
func (o *MemberActivityEvents) HasContactId() bool {
	if o != nil && !IsNil(o.ContactId) {
		return true
	}

	return false
}

// SetContactId gets a reference to the given string and assigns it to the ContactId field.
func (o *MemberActivityEvents) SetContactId(v string) {
	o.ContactId = &v
}

// GetListId returns the ListId field value if set, zero value otherwise.
func (o *MemberActivityEvents) GetListId() string {
	if o == nil || IsNil(o.ListId) {
		var ret string
		return ret
	}
	return *o.ListId
}

// GetListIdOk returns a tuple with the ListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberActivityEvents) GetListIdOk() (*string, bool) {
	if o == nil || IsNil(o.ListId) {
		return nil, false
	}
	return o.ListId, true
}

// HasListId returns a boolean if a field has been set.
func (o *MemberActivityEvents) HasListId() bool {
	if o != nil && !IsNil(o.ListId) {
		return true
	}

	return false
}

// SetListId gets a reference to the given string and assigns it to the ListId field.
func (o *MemberActivityEvents) SetListId(v string) {
	o.ListId = &v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *MemberActivityEvents) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberActivityEvents) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *MemberActivityEvents) HasTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *MemberActivityEvents) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *MemberActivityEvents) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberActivityEvents) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *MemberActivityEvents) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *MemberActivityEvents) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o MemberActivityEvents) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemberActivityEvents) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Activity) {
		toSerialize["activity"] = o.Activity
	}
	if !IsNil(o.EmailId) {
		toSerialize["email_id"] = o.EmailId
	}
	if !IsNil(o.ContactId) {
		toSerialize["contact_id"] = o.ContactId
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

type NullableMemberActivityEvents struct {
	value *MemberActivityEvents
	isSet bool
}

func (v NullableMemberActivityEvents) Get() *MemberActivityEvents {
	return v.value
}

func (v *NullableMemberActivityEvents) Set(val *MemberActivityEvents) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberActivityEvents) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberActivityEvents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberActivityEvents(val *MemberActivityEvents) *NullableMemberActivityEvents {
	return &NullableMemberActivityEvents{value: val, isSet: true}
}

func (v NullableMemberActivityEvents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberActivityEvents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


