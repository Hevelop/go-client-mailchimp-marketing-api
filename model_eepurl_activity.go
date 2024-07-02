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

// checks if the EepurlActivity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EepurlActivity{}

// EepurlActivity A summary of social activity for the campaign, tracked by EepURL.
type EepurlActivity struct {
	Twitter *TwitterStats `json:"twitter,omitempty"`
	Clicks *ClickSummary `json:"clicks,omitempty"`
	// A summary of the top referrers for the campaign.
	Referrers []Referrer `json:"referrers,omitempty"`
	// The shortened link used for tracking.
	Eepurl *string `json:"eepurl,omitempty"`
	// The unique id for the campaign.
	CampaignId *string `json:"campaign_id,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int32 `json:"total_items,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewEepurlActivity instantiates a new EepurlActivity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEepurlActivity() *EepurlActivity {
	this := EepurlActivity{}
	return &this
}

// NewEepurlActivityWithDefaults instantiates a new EepurlActivity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEepurlActivityWithDefaults() *EepurlActivity {
	this := EepurlActivity{}
	return &this
}

// GetTwitter returns the Twitter field value if set, zero value otherwise.
func (o *EepurlActivity) GetTwitter() TwitterStats {
	if o == nil || IsNil(o.Twitter) {
		var ret TwitterStats
		return ret
	}
	return *o.Twitter
}

// GetTwitterOk returns a tuple with the Twitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EepurlActivity) GetTwitterOk() (*TwitterStats, bool) {
	if o == nil || IsNil(o.Twitter) {
		return nil, false
	}
	return o.Twitter, true
}

// HasTwitter returns a boolean if a field has been set.
func (o *EepurlActivity) HasTwitter() bool {
	if o != nil && !IsNil(o.Twitter) {
		return true
	}

	return false
}

// SetTwitter gets a reference to the given TwitterStats and assigns it to the Twitter field.
func (o *EepurlActivity) SetTwitter(v TwitterStats) {
	o.Twitter = &v
}

// GetClicks returns the Clicks field value if set, zero value otherwise.
func (o *EepurlActivity) GetClicks() ClickSummary {
	if o == nil || IsNil(o.Clicks) {
		var ret ClickSummary
		return ret
	}
	return *o.Clicks
}

// GetClicksOk returns a tuple with the Clicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EepurlActivity) GetClicksOk() (*ClickSummary, bool) {
	if o == nil || IsNil(o.Clicks) {
		return nil, false
	}
	return o.Clicks, true
}

// HasClicks returns a boolean if a field has been set.
func (o *EepurlActivity) HasClicks() bool {
	if o != nil && !IsNil(o.Clicks) {
		return true
	}

	return false
}

// SetClicks gets a reference to the given ClickSummary and assigns it to the Clicks field.
func (o *EepurlActivity) SetClicks(v ClickSummary) {
	o.Clicks = &v
}

// GetReferrers returns the Referrers field value if set, zero value otherwise.
func (o *EepurlActivity) GetReferrers() []Referrer {
	if o == nil || IsNil(o.Referrers) {
		var ret []Referrer
		return ret
	}
	return o.Referrers
}

// GetReferrersOk returns a tuple with the Referrers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EepurlActivity) GetReferrersOk() ([]Referrer, bool) {
	if o == nil || IsNil(o.Referrers) {
		return nil, false
	}
	return o.Referrers, true
}

// HasReferrers returns a boolean if a field has been set.
func (o *EepurlActivity) HasReferrers() bool {
	if o != nil && !IsNil(o.Referrers) {
		return true
	}

	return false
}

// SetReferrers gets a reference to the given []Referrer and assigns it to the Referrers field.
func (o *EepurlActivity) SetReferrers(v []Referrer) {
	o.Referrers = v
}

// GetEepurl returns the Eepurl field value if set, zero value otherwise.
func (o *EepurlActivity) GetEepurl() string {
	if o == nil || IsNil(o.Eepurl) {
		var ret string
		return ret
	}
	return *o.Eepurl
}

// GetEepurlOk returns a tuple with the Eepurl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EepurlActivity) GetEepurlOk() (*string, bool) {
	if o == nil || IsNil(o.Eepurl) {
		return nil, false
	}
	return o.Eepurl, true
}

// HasEepurl returns a boolean if a field has been set.
func (o *EepurlActivity) HasEepurl() bool {
	if o != nil && !IsNil(o.Eepurl) {
		return true
	}

	return false
}

// SetEepurl gets a reference to the given string and assigns it to the Eepurl field.
func (o *EepurlActivity) SetEepurl(v string) {
	o.Eepurl = &v
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *EepurlActivity) GetCampaignId() string {
	if o == nil || IsNil(o.CampaignId) {
		var ret string
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EepurlActivity) GetCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *EepurlActivity) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given string and assigns it to the CampaignId field.
func (o *EepurlActivity) SetCampaignId(v string) {
	o.CampaignId = &v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *EepurlActivity) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EepurlActivity) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *EepurlActivity) HasTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *EepurlActivity) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *EepurlActivity) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EepurlActivity) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EepurlActivity) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *EepurlActivity) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o EepurlActivity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EepurlActivity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Twitter) {
		toSerialize["twitter"] = o.Twitter
	}
	if !IsNil(o.Clicks) {
		toSerialize["clicks"] = o.Clicks
	}
	if !IsNil(o.Referrers) {
		toSerialize["referrers"] = o.Referrers
	}
	if !IsNil(o.Eepurl) {
		toSerialize["eepurl"] = o.Eepurl
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

type NullableEepurlActivity struct {
	value *EepurlActivity
	isSet bool
}

func (v NullableEepurlActivity) Get() *EepurlActivity {
	return v.value
}

func (v *NullableEepurlActivity) Set(val *EepurlActivity) {
	v.value = val
	v.isSet = true
}

func (v NullableEepurlActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableEepurlActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEepurlActivity(val *EepurlActivity) *NullableEepurlActivity {
	return &NullableEepurlActivity{value: val, isSet: true}
}

func (v NullableEepurlActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEepurlActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


