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

// checks if the EmailClients type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailClients{}

// EmailClients The top email clients based on user-agent strings.
type EmailClients struct {
	// An array of top email clients.
	Clients []EmailClient `json:"clients,omitempty"`
	// The list id.
	ListId *string `json:"list_id,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int32 `json:"total_items,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewEmailClients instantiates a new EmailClients object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailClients() *EmailClients {
	this := EmailClients{}
	return &this
}

// NewEmailClientsWithDefaults instantiates a new EmailClients object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailClientsWithDefaults() *EmailClients {
	this := EmailClients{}
	return &this
}

// GetClients returns the Clients field value if set, zero value otherwise.
func (o *EmailClients) GetClients() []EmailClient {
	if o == nil || IsNil(o.Clients) {
		var ret []EmailClient
		return ret
	}
	return o.Clients
}

// GetClientsOk returns a tuple with the Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailClients) GetClientsOk() ([]EmailClient, bool) {
	if o == nil || IsNil(o.Clients) {
		return nil, false
	}
	return o.Clients, true
}

// HasClients returns a boolean if a field has been set.
func (o *EmailClients) HasClients() bool {
	if o != nil && !IsNil(o.Clients) {
		return true
	}

	return false
}

// SetClients gets a reference to the given []EmailClient and assigns it to the Clients field.
func (o *EmailClients) SetClients(v []EmailClient) {
	o.Clients = v
}

// GetListId returns the ListId field value if set, zero value otherwise.
func (o *EmailClients) GetListId() string {
	if o == nil || IsNil(o.ListId) {
		var ret string
		return ret
	}
	return *o.ListId
}

// GetListIdOk returns a tuple with the ListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailClients) GetListIdOk() (*string, bool) {
	if o == nil || IsNil(o.ListId) {
		return nil, false
	}
	return o.ListId, true
}

// HasListId returns a boolean if a field has been set.
func (o *EmailClients) HasListId() bool {
	if o != nil && !IsNil(o.ListId) {
		return true
	}

	return false
}

// SetListId gets a reference to the given string and assigns it to the ListId field.
func (o *EmailClients) SetListId(v string) {
	o.ListId = &v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *EmailClients) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailClients) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *EmailClients) HasTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *EmailClients) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *EmailClients) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailClients) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EmailClients) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *EmailClients) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o EmailClients) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailClients) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Clients) {
		toSerialize["clients"] = o.Clients
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

type NullableEmailClients struct {
	value *EmailClients
	isSet bool
}

func (v NullableEmailClients) Get() *EmailClients {
	return v.value
}

func (v *NullableEmailClients) Set(val *EmailClients) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailClients) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailClients) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailClients(val *EmailClients) *NullableEmailClients {
	return &NullableEmailClients{value: val, isSet: true}
}

func (v NullableEmailClients) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailClients) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

