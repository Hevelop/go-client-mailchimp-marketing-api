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

// checks if the GetCampaigns200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCampaigns200Response{}

// GetCampaigns200Response An array of campaigns.
type GetCampaigns200Response struct {
	// An array of campaigns.
	Campaigns []Campaign `json:"campaigns,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int32 `json:"total_items,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewGetCampaigns200Response instantiates a new GetCampaigns200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCampaigns200Response() *GetCampaigns200Response {
	this := GetCampaigns200Response{}
	return &this
}

// NewGetCampaigns200ResponseWithDefaults instantiates a new GetCampaigns200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCampaigns200ResponseWithDefaults() *GetCampaigns200Response {
	this := GetCampaigns200Response{}
	return &this
}

// GetCampaigns returns the Campaigns field value if set, zero value otherwise.
func (o *GetCampaigns200Response) GetCampaigns() []Campaign {
	if o == nil || IsNil(o.Campaigns) {
		var ret []Campaign
		return ret
	}
	return o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCampaigns200Response) GetCampaignsOk() ([]Campaign, bool) {
	if o == nil || IsNil(o.Campaigns) {
		return nil, false
	}
	return o.Campaigns, true
}

// HasCampaigns returns a boolean if a field has been set.
func (o *GetCampaigns200Response) HasCampaigns() bool {
	if o != nil && !IsNil(o.Campaigns) {
		return true
	}

	return false
}

// SetCampaigns gets a reference to the given []Campaign and assigns it to the Campaigns field.
func (o *GetCampaigns200Response) SetCampaigns(v []Campaign) {
	o.Campaigns = v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *GetCampaigns200Response) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCampaigns200Response) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *GetCampaigns200Response) HasTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *GetCampaigns200Response) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GetCampaigns200Response) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCampaigns200Response) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GetCampaigns200Response) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *GetCampaigns200Response) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o GetCampaigns200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCampaigns200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Campaigns) {
		toSerialize["campaigns"] = o.Campaigns
	}
	if !IsNil(o.TotalItems) {
		toSerialize["total_items"] = o.TotalItems
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableGetCampaigns200Response struct {
	value *GetCampaigns200Response
	isSet bool
}

func (v NullableGetCampaigns200Response) Get() *GetCampaigns200Response {
	return v.value
}

func (v *NullableGetCampaigns200Response) Set(val *GetCampaigns200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCampaigns200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCampaigns200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCampaigns200Response(val *GetCampaigns200Response) *NullableGetCampaigns200Response {
	return &NullableGetCampaigns200Response{value: val, isSet: true}
}

func (v NullableGetCampaigns200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCampaigns200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


