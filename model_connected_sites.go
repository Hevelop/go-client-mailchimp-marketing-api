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

// checks if the ConnectedSites type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectedSites{}

// ConnectedSites A collection of connected sites in the account.
type ConnectedSites struct {
	// An array of objects, each representing a connected site.
	Sites []ConnectedSite `json:"sites,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int32 `json:"total_items,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewConnectedSites instantiates a new ConnectedSites object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectedSites() *ConnectedSites {
	this := ConnectedSites{}
	return &this
}

// NewConnectedSitesWithDefaults instantiates a new ConnectedSites object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectedSitesWithDefaults() *ConnectedSites {
	this := ConnectedSites{}
	return &this
}

// GetSites returns the Sites field value if set, zero value otherwise.
func (o *ConnectedSites) GetSites() []ConnectedSite {
	if o == nil || IsNil(o.Sites) {
		var ret []ConnectedSite
		return ret
	}
	return o.Sites
}

// GetSitesOk returns a tuple with the Sites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedSites) GetSitesOk() ([]ConnectedSite, bool) {
	if o == nil || IsNil(o.Sites) {
		return nil, false
	}
	return o.Sites, true
}

// HaveSites returns a boolean if a field has been set.
func (o *ConnectedSites) HaveSites() bool {
	if o != nil && !IsNil(o.Sites) {
		return true
	}

	return false
}

// SetSites gets a reference to the given []ConnectedSite and assigns it to the Sites field.
func (o *ConnectedSites) SetSites(v []ConnectedSite) {
	o.Sites = v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *ConnectedSites) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedSites) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HaveTotalItems returns a boolean if a field has been set.
func (o *ConnectedSites) HaveTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *ConnectedSites) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ConnectedSites) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedSites) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HaveLinks returns a boolean if a field has been set.
func (o *ConnectedSites) HaveLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *ConnectedSites) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o ConnectedSites) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectedSites) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sites) {
		toSerialize["sites"] = o.Sites
	}
	if !IsNil(o.TotalItems) {
		toSerialize["total_items"] = o.TotalItems
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableConnectedSites struct {
	value *ConnectedSites
	isSet bool
}

func (v NullableConnectedSites) Get() *ConnectedSites {
	return v.value
}

func (v *NullableConnectedSites) Set(val *ConnectedSites) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectedSites) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectedSites) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectedSites(val *ConnectedSites) *NullableConnectedSites {
	return &NullableConnectedSites{value: val, isSet: true}
}

func (v NullableConnectedSites) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectedSites) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


