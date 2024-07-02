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

// checks if the InterestGroupings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InterestGroupings{}

// InterestGroupings Information about this list's interest categories.
type InterestGroupings struct {
	// The ID for the list that this category belongs to.
	ListId *string `json:"list_id,omitempty"`
	// This array contains individual interest categories.
	Categories []InterestCategory `json:"categories,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int32 `json:"total_items,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewInterestGroupings instantiates a new InterestGroupings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterestGroupings() *InterestGroupings {
	this := InterestGroupings{}
	return &this
}

// NewInterestGroupingsWithDefaults instantiates a new InterestGroupings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterestGroupingsWithDefaults() *InterestGroupings {
	this := InterestGroupings{}
	return &this
}

// GetListId returns the ListId field value if set, zero value otherwise.
func (o *InterestGroupings) GetListId() string {
	if o == nil || IsNil(o.ListId) {
		var ret string
		return ret
	}
	return *o.ListId
}

// GetListIdOk returns a tuple with the ListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterestGroupings) GetListIdOk() (*string, bool) {
	if o == nil || IsNil(o.ListId) {
		return nil, false
	}
	return o.ListId, true
}

// HasListId returns a boolean if a field has been set.
func (o *InterestGroupings) HasListId() bool {
	if o != nil && !IsNil(o.ListId) {
		return true
	}

	return false
}

// SetListId gets a reference to the given string and assigns it to the ListId field.
func (o *InterestGroupings) SetListId(v string) {
	o.ListId = &v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *InterestGroupings) GetCategories() []InterestCategory {
	if o == nil || IsNil(o.Categories) {
		var ret []InterestCategory
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterestGroupings) GetCategoriesOk() ([]InterestCategory, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *InterestGroupings) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []InterestCategory and assigns it to the Categories field.
func (o *InterestGroupings) SetCategories(v []InterestCategory) {
	o.Categories = v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *InterestGroupings) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterestGroupings) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *InterestGroupings) HasTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *InterestGroupings) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *InterestGroupings) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterestGroupings) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *InterestGroupings) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *InterestGroupings) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o InterestGroupings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterestGroupings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ListId) {
		toSerialize["list_id"] = o.ListId
	}
	if !IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	if !IsNil(o.TotalItems) {
		toSerialize["total_items"] = o.TotalItems
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableInterestGroupings struct {
	value *InterestGroupings
	isSet bool
}

func (v NullableInterestGroupings) Get() *InterestGroupings {
	return v.value
}

func (v *NullableInterestGroupings) Set(val *InterestGroupings) {
	v.value = val
	v.isSet = true
}

func (v NullableInterestGroupings) IsSet() bool {
	return v.isSet
}

func (v *NullableInterestGroupings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterestGroupings(val *InterestGroupings) *NullableInterestGroupings {
	return &NullableInterestGroupings{value: val, isSet: true}
}

func (v NullableInterestGroupings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterestGroupings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

