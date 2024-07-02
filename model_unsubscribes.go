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

// checks if the Unsubscribes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Unsubscribes{}

// Unsubscribes A list of members who have unsubscribed from a specific campaign.
type Unsubscribes struct {
	// An array of objects, each representing a member who unsubscribed from a campaign.
	Unsubscribes []Unsubscribes `json:"unsubscribes,omitempty"`
	// The campaign id.
	CampaignId *string `json:"campaign_id,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int32 `json:"total_items,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewUnsubscribes instantiates a new Unsubscribes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnsubscribes() *Unsubscribes {
	this := Unsubscribes{}
	return &this
}

// NewUnsubscribesWithDefaults instantiates a new Unsubscribes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnsubscribesWithDefaults() *Unsubscribes {
	this := Unsubscribes{}
	return &this
}

// GetUnsubscribes returns the Unsubscribes field value if set, zero value otherwise.
func (o *Unsubscribes) GetUnsubscribes() []Unsubscribes {
	if o == nil || IsNil(o.Unsubscribes) {
		var ret []Unsubscribes
		return ret
	}
	return o.Unsubscribes
}

// GetUnsubscribesOk returns a tuple with the Unsubscribes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Unsubscribes) GetUnsubscribesOk() ([]Unsubscribes, bool) {
	if o == nil || IsNil(o.Unsubscribes) {
		return nil, false
	}
	return o.Unsubscribes, true
}

// HaveUnsubscribes returns a boolean if a field has been set.
func (o *Unsubscribes) HaveUnsubscribes() bool {
	if o != nil && !IsNil(o.Unsubscribes) {
		return true
	}

	return false
}

// SetUnsubscribes gets a reference to the given []Unsubscribes and assigns it to the Unsubscribes field.
func (o *Unsubscribes) SetUnsubscribes(v []Unsubscribes) {
	o.Unsubscribes = v
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *Unsubscribes) GetCampaignId() string {
	if o == nil || IsNil(o.CampaignId) {
		var ret string
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Unsubscribes) GetCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HaveCampaignId returns a boolean if a field has been set.
func (o *Unsubscribes) HaveCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given string and assigns it to the CampaignId field.
func (o *Unsubscribes) SetCampaignId(v string) {
	o.CampaignId = &v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *Unsubscribes) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Unsubscribes) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HaveTotalItems returns a boolean if a field has been set.
func (o *Unsubscribes) HaveTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *Unsubscribes) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Unsubscribes) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Unsubscribes) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HaveLinks returns a boolean if a field has been set.
func (o *Unsubscribes) HaveLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *Unsubscribes) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o Unsubscribes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Unsubscribes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Unsubscribes) {
		toSerialize["unsubscribes"] = o.Unsubscribes
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

type NullableUnsubscribes struct {
	value *Unsubscribes
	isSet bool
}

func (v NullableUnsubscribes) Get() *Unsubscribes {
	return v.value
}

func (v *NullableUnsubscribes) Set(val *Unsubscribes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnsubscribes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnsubscribes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnsubscribes(val *Unsubscribes) *NullableUnsubscribes {
	return &NullableUnsubscribes{value: val, isSet: true}
}

func (v NullableUnsubscribes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnsubscribes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


