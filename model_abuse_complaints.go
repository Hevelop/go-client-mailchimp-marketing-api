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

// checks if the AbuseComplaints type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AbuseComplaints{}

// AbuseComplaints A collection of abuse complaints for a specific list. An abuse complaint occurs when your recipient clicks to 'report spam' in their email program.
type AbuseComplaints struct {
	// An array of objects, each representing an abuse report resource.
	AbuseReports []AbuseComplaint `json:"abuse_reports,omitempty"`
	// The list id for the abuse report.
	ListId *string `json:"list_id,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int32 `json:"total_items,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewAbuseComplaints instantiates a new AbuseComplaints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAbuseComplaints() *AbuseComplaints {
	this := AbuseComplaints{}
	return &this
}

// NewAbuseComplaintsWithDefaults instantiates a new AbuseComplaints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAbuseComplaintsWithDefaults() *AbuseComplaints {
	this := AbuseComplaints{}
	return &this
}

// GetAbuseReports returns the AbuseReports field value if set, zero value otherwise.
func (o *AbuseComplaints) GetAbuseReports() []AbuseComplaint {
	if o == nil || IsNil(o.AbuseReports) {
		var ret []AbuseComplaint
		return ret
	}
	return o.AbuseReports
}

// GetAbuseReportsOk returns a tuple with the AbuseReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbuseComplaints) GetAbuseReportsOk() ([]AbuseComplaint, bool) {
	if o == nil || IsNil(o.AbuseReports) {
		return nil, false
	}
	return o.AbuseReports, true
}

// HasAbuseReports returns a boolean if a field has been set.
func (o *AbuseComplaints) HasAbuseReports() bool {
	if o != nil && !IsNil(o.AbuseReports) {
		return true
	}

	return false
}

// SetAbuseReports gets a reference to the given []AbuseComplaint and assigns it to the AbuseReports field.
func (o *AbuseComplaints) SetAbuseReports(v []AbuseComplaint) {
	o.AbuseReports = v
}

// GetListId returns the ListId field value if set, zero value otherwise.
func (o *AbuseComplaints) GetListId() string {
	if o == nil || IsNil(o.ListId) {
		var ret string
		return ret
	}
	return *o.ListId
}

// GetListIdOk returns a tuple with the ListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbuseComplaints) GetListIdOk() (*string, bool) {
	if o == nil || IsNil(o.ListId) {
		return nil, false
	}
	return o.ListId, true
}

// HasListId returns a boolean if a field has been set.
func (o *AbuseComplaints) HasListId() bool {
	if o != nil && !IsNil(o.ListId) {
		return true
	}

	return false
}

// SetListId gets a reference to the given string and assigns it to the ListId field.
func (o *AbuseComplaints) SetListId(v string) {
	o.ListId = &v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *AbuseComplaints) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbuseComplaints) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *AbuseComplaints) HasTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *AbuseComplaints) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AbuseComplaints) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbuseComplaints) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AbuseComplaints) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *AbuseComplaints) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o AbuseComplaints) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AbuseComplaints) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AbuseReports) {
		toSerialize["abuse_reports"] = o.AbuseReports
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

type NullableAbuseComplaints struct {
	value *AbuseComplaints
	isSet bool
}

func (v NullableAbuseComplaints) Get() *AbuseComplaints {
	return v.value
}

func (v *NullableAbuseComplaints) Set(val *AbuseComplaints) {
	v.value = val
	v.isSet = true
}

func (v NullableAbuseComplaints) IsSet() bool {
	return v.isSet
}

func (v *NullableAbuseComplaints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAbuseComplaints(val *AbuseComplaints) *NullableAbuseComplaints {
	return &NullableAbuseComplaints{value: val, isSet: true}
}

func (v NullableAbuseComplaints) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAbuseComplaints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


