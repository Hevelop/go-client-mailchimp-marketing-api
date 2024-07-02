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

// checks if the ClickDetailMembers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClickDetailMembers{}

// ClickDetailMembers A collection of members who clicked on a specific link within a campaign.
type ClickDetailMembers struct {
	// An array of objects, each representing a member who clicked a specific link within a campaign.
	Members []ClickDetailMember `json:"members,omitempty"`
	// The campaign id.
	CampaignId *string `json:"campaign_id,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int32 `json:"total_items,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewClickDetailMembers instantiates a new ClickDetailMembers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClickDetailMembers() *ClickDetailMembers {
	this := ClickDetailMembers{}
	return &this
}

// NewClickDetailMembersWithDefaults instantiates a new ClickDetailMembers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClickDetailMembersWithDefaults() *ClickDetailMembers {
	this := ClickDetailMembers{}
	return &this
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *ClickDetailMembers) GetMembers() []ClickDetailMember {
	if o == nil || IsNil(o.Members) {
		var ret []ClickDetailMember
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickDetailMembers) GetMembersOk() ([]ClickDetailMember, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *ClickDetailMembers) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []ClickDetailMember and assigns it to the Members field.
func (o *ClickDetailMembers) SetMembers(v []ClickDetailMember) {
	o.Members = v
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *ClickDetailMembers) GetCampaignId() string {
	if o == nil || IsNil(o.CampaignId) {
		var ret string
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickDetailMembers) GetCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *ClickDetailMembers) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given string and assigns it to the CampaignId field.
func (o *ClickDetailMembers) SetCampaignId(v string) {
	o.CampaignId = &v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *ClickDetailMembers) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickDetailMembers) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *ClickDetailMembers) HasTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *ClickDetailMembers) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ClickDetailMembers) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickDetailMembers) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ClickDetailMembers) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *ClickDetailMembers) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o ClickDetailMembers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClickDetailMembers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Members) {
		toSerialize["members"] = o.Members
	}
	if !IsNil(o.CampaignId) {
		toSerialize["campaign_id"] = o.CampaignId
	}
	if !IsNil(o.TotalItems) {
		toSerialize["total_items"] = o.TotalItems
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableClickDetailMembers struct {
	value *ClickDetailMembers
	isSet bool
}

func (v NullableClickDetailMembers) Get() *ClickDetailMembers {
	return v.value
}

func (v *NullableClickDetailMembers) Set(val *ClickDetailMembers) {
	v.value = val
	v.isSet = true
}

func (v NullableClickDetailMembers) IsSet() bool {
	return v.isSet
}

func (v *NullableClickDetailMembers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClickDetailMembers(val *ClickDetailMembers) *NullableClickDetailMembers {
	return &NullableClickDetailMembers{value: val, isSet: true}
}

func (v NullableClickDetailMembers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClickDetailMembers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


