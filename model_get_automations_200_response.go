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

// checks if the GetAutomations200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAutomations200Response{}

// GetAutomations200Response An array of objects, each representing an Automation workflow.
type GetAutomations200Response struct {
	// An array of objects, each representing an Automation workflow.
	Automations []AutomationWorkflow `json:"automations,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int32 `json:"total_items,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewGetAutomations200Response instantiates a new GetAutomations200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAutomations200Response() *GetAutomations200Response {
	this := GetAutomations200Response{}
	return &this
}

// NewGetAutomations200ResponseWithDefaults instantiates a new GetAutomations200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAutomations200ResponseWithDefaults() *GetAutomations200Response {
	this := GetAutomations200Response{}
	return &this
}

// GetAutomations returns the Automations field value if set, zero value otherwise.
func (o *GetAutomations200Response) GetAutomations() []AutomationWorkflow {
	if o == nil || IsNil(o.Automations) {
		var ret []AutomationWorkflow
		return ret
	}
	return o.Automations
}

// GetAutomationsOk returns a tuple with the Automations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAutomations200Response) GetAutomationsOk() ([]AutomationWorkflow, bool) {
	if o == nil || IsNil(o.Automations) {
		return nil, false
	}
	return o.Automations, true
}

// HasAutomations returns a boolean if a field has been set.
func (o *GetAutomations200Response) HasAutomations() bool {
	if o != nil && !IsNil(o.Automations) {
		return true
	}

	return false
}

// SetAutomations gets a reference to the given []AutomationWorkflow and assigns it to the Automations field.
func (o *GetAutomations200Response) SetAutomations(v []AutomationWorkflow) {
	o.Automations = v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *GetAutomations200Response) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAutomations200Response) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *GetAutomations200Response) HasTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *GetAutomations200Response) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GetAutomations200Response) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAutomations200Response) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GetAutomations200Response) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *GetAutomations200Response) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o GetAutomations200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAutomations200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Automations) {
		toSerialize["automations"] = o.Automations
	}
	if !IsNil(o.TotalItems) {
		toSerialize["total_items"] = o.TotalItems
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableGetAutomations200Response struct {
	value *GetAutomations200Response
	isSet bool
}

func (v NullableGetAutomations200Response) Get() *GetAutomations200Response {
	return v.value
}

func (v *NullableGetAutomations200Response) Set(val *GetAutomations200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAutomations200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAutomations200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAutomations200Response(val *GetAutomations200Response) *NullableGetAutomations200Response {
	return &NullableGetAutomations200Response{value: val, isSet: true}
}

func (v NullableGetAutomations200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAutomations200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


