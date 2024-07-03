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

// checks if the OpenLocations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenLocations{}

// OpenLocations Top open locations for a specific campaign.
type OpenLocations struct {
	// An array of objects, each representing a top location for opens.
	Locations []ListsInner `json:"locations,omitempty"`
	// The campaign id.
	CampaignId *string `json:"campaign_id,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int32 `json:"total_items,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewOpenLocations instantiates a new OpenLocations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenLocations() *OpenLocations {
	this := OpenLocations{}
	return &this
}

// NewOpenLocationsWithDefaults instantiates a new OpenLocations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenLocationsWithDefaults() *OpenLocations {
	this := OpenLocations{}
	return &this
}

// GetLocations returns the Locations field value if set, zero value otherwise.
func (o *OpenLocations) GetLocations() []ListsInner {
	if o == nil || IsNil(o.Locations) {
		var ret []ListsInner
		return ret
	}
	return o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenLocations) GetLocationsOk() ([]ListsInner, bool) {
	if o == nil || IsNil(o.Locations) {
		return nil, false
	}
	return o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *OpenLocations) HasLocations() bool {
	if o != nil && !IsNil(o.Locations) {
		return true
	}

	return false
}

// SetLocations gets a reference to the given []ListsInner and assigns it to the Locations field.
func (o *OpenLocations) SetLocations(v []ListsInner) {
	o.Locations = v
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *OpenLocations) GetCampaignId() string {
	if o == nil || IsNil(o.CampaignId) {
		var ret string
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenLocations) GetCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *OpenLocations) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given string and assigns it to the CampaignId field.
func (o *OpenLocations) SetCampaignId(v string) {
	o.CampaignId = &v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *OpenLocations) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenLocations) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *OpenLocations) HasTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *OpenLocations) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OpenLocations) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenLocations) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OpenLocations) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *OpenLocations) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o OpenLocations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenLocations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Locations) {
		toSerialize["locations"] = o.Locations
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

type NullableOpenLocations struct {
	value *OpenLocations
	isSet bool
}

func (v NullableOpenLocations) Get() *OpenLocations {
	return v.value
}

func (v *NullableOpenLocations) Set(val *OpenLocations) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenLocations) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenLocations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenLocations(val *OpenLocations) *NullableOpenLocations {
	return &NullableOpenLocations{value: val, isSet: true}
}

func (v NullableOpenLocations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenLocations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


