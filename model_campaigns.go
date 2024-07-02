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

// checks if the Campaigns type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Campaigns{}

// Campaigns Campaigns and Snippets found for given search term.
type Campaigns struct {
	// An array of matching campaigns and snippets.
	Results []ResultsInner `json:"results,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int32 `json:"total_items,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewCampaigns instantiates a new Campaigns object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaigns() *Campaigns {
	this := Campaigns{}
	return &this
}

// NewCampaignsWithDefaults instantiates a new Campaigns object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignsWithDefaults() *Campaigns {
	this := Campaigns{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *Campaigns) GetResults() []ResultsInner {
	if o == nil || IsNil(o.Results) {
		var ret []ResultsInner
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaigns) GetResultsOk() ([]ResultsInner, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HaveResults returns a boolean if a field has been set.
func (o *Campaigns) HaveResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ResultsInner and assigns it to the Results field.
func (o *Campaigns) SetResults(v []ResultsInner) {
	o.Results = v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *Campaigns) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaigns) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HaveTotalItems returns a boolean if a field has been set.
func (o *Campaigns) HaveTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *Campaigns) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Campaigns) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaigns) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HaveLinks returns a boolean if a field has been set.
func (o *Campaigns) HaveLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *Campaigns) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o Campaigns) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Campaigns) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.TotalItems) {
		toSerialize["total_items"] = o.TotalItems
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableCampaigns struct {
	value *Campaigns
	isSet bool
}

func (v NullableCampaigns) Get() *Campaigns {
	return v.value
}

func (v *NullableCampaigns) Set(val *Campaigns) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaigns) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaigns) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaigns(val *Campaigns) *NullableCampaigns {
	return &NullableCampaigns{value: val, isSet: true}
}

func (v NullableCampaigns) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaigns) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


