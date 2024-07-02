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

// checks if the ClickDetailMember type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClickDetailMember{}

// ClickDetailMember A subscriber who clicked a specific URL in a specific campaign.
type ClickDetailMember struct {
	// The MD5 hash of the lowercase version of the list member's email address.
	EmailId *string `json:"email_id,omitempty"`
	// Email address for a subscriber.
	EmailAddress *string `json:"email_address,omitempty"`
	// A dictionary of merge fields where the keys are the merge tags. See the [Merge Fields documentation](https://mailchimp.com/developer/marketing/docs/merge-fields/#structure) for more about the structure.
	MergeFields map[string]map[string]interface{} `json:"merge_fields,omitempty"`
	// [VIP status](https://mailchimp.com/help/designate-and-send-to-vip-contacts/) for subscriber.
	Vip *bool `json:"vip,omitempty"`
	// The total number of times the subscriber clicked on the link.
	Clicks *int32 `json:"clicks,omitempty"`
	// The campaign id.
	CampaignId *string `json:"campaign_id,omitempty"`
	// The id for the tracked URL in the campaign.
	UrlId *string `json:"url_id,omitempty"`
	// The list id.
	ListId *string `json:"list_id,omitempty"`
	// The status of the list used, namely if it's deleted or disabled.
	ListIsActive *bool `json:"list_is_active,omitempty"`
	// The status of the member, namely if they are subscribed, unsubscribed, deleted, non-subscribed, transactional, pending, or need reconfirmation.
	ContactStatus *string `json:"contact_status,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewClickDetailMember instantiates a new ClickDetailMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClickDetailMember() *ClickDetailMember {
	this := ClickDetailMember{}
	return &this
}

// NewClickDetailMemberWithDefaults instantiates a new ClickDetailMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClickDetailMemberWithDefaults() *ClickDetailMember {
	this := ClickDetailMember{}
	return &this
}

// GetEmailId returns the EmailId field value if set, zero value otherwise.
func (o *ClickDetailMember) GetEmailId() string {
	if o == nil || IsNil(o.EmailId) {
		var ret string
		return ret
	}
	return *o.EmailId
}

// GetEmailIdOk returns a tuple with the EmailId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickDetailMember) GetEmailIdOk() (*string, bool) {
	if o == nil || IsNil(o.EmailId) {
		return nil, false
	}
	return o.EmailId, true
}

// HasEmailId returns a boolean if a field has been set.
func (o *ClickDetailMember) HasEmailId() bool {
	if o != nil && !IsNil(o.EmailId) {
		return true
	}

	return false
}

// SetEmailId gets a reference to the given string and assigns it to the EmailId field.
func (o *ClickDetailMember) SetEmailId(v string) {
	o.EmailId = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *ClickDetailMember) GetEmailAddress() string {
	if o == nil || IsNil(o.EmailAddress) {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickDetailMember) GetEmailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EmailAddress) {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *ClickDetailMember) HasEmailAddress() bool {
	if o != nil && !IsNil(o.EmailAddress) {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *ClickDetailMember) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetMergeFields returns the MergeFields field value if set, zero value otherwise.
func (o *ClickDetailMember) GetMergeFields() map[string]map[string]interface{} {
	if o == nil || IsNil(o.MergeFields) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.MergeFields
}

// GetMergeFieldsOk returns a tuple with the MergeFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickDetailMember) GetMergeFieldsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.MergeFields) {
		return map[string]map[string]interface{}{}, false
	}
	return o.MergeFields, true
}

// HasMergeFields returns a boolean if a field has been set.
func (o *ClickDetailMember) HasMergeFields() bool {
	if o != nil && !IsNil(o.MergeFields) {
		return true
	}

	return false
}

// SetMergeFields gets a reference to the given map[string]map[string]interface{} and assigns it to the MergeFields field.
func (o *ClickDetailMember) SetMergeFields(v map[string]map[string]interface{}) {
	o.MergeFields = v
}

// GetVip returns the Vip field value if set, zero value otherwise.
func (o *ClickDetailMember) GetVip() bool {
	if o == nil || IsNil(o.Vip) {
		var ret bool
		return ret
	}
	return *o.Vip
}

// GetVipOk returns a tuple with the Vip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickDetailMember) GetVipOk() (*bool, bool) {
	if o == nil || IsNil(o.Vip) {
		return nil, false
	}
	return o.Vip, true
}

// HasVip returns a boolean if a field has been set.
func (o *ClickDetailMember) HasVip() bool {
	if o != nil && !IsNil(o.Vip) {
		return true
	}

	return false
}

// SetVip gets a reference to the given bool and assigns it to the Vip field.
func (o *ClickDetailMember) SetVip(v bool) {
	o.Vip = &v
}

// GetClicks returns the Clicks field value if set, zero value otherwise.
func (o *ClickDetailMember) GetClicks() int32 {
	if o == nil || IsNil(o.Clicks) {
		var ret int32
		return ret
	}
	return *o.Clicks
}

// GetClicksOk returns a tuple with the Clicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickDetailMember) GetClicksOk() (*int32, bool) {
	if o == nil || IsNil(o.Clicks) {
		return nil, false
	}
	return o.Clicks, true
}

// HasClicks returns a boolean if a field has been set.
func (o *ClickDetailMember) HasClicks() bool {
	if o != nil && !IsNil(o.Clicks) {
		return true
	}

	return false
}

// SetClicks gets a reference to the given int32 and assigns it to the Clicks field.
func (o *ClickDetailMember) SetClicks(v int32) {
	o.Clicks = &v
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *ClickDetailMember) GetCampaignId() string {
	if o == nil || IsNil(o.CampaignId) {
		var ret string
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickDetailMember) GetCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *ClickDetailMember) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given string and assigns it to the CampaignId field.
func (o *ClickDetailMember) SetCampaignId(v string) {
	o.CampaignId = &v
}

// GetUrlId returns the UrlId field value if set, zero value otherwise.
func (o *ClickDetailMember) GetUrlId() string {
	if o == nil || IsNil(o.UrlId) {
		var ret string
		return ret
	}
	return *o.UrlId
}

// GetUrlIdOk returns a tuple with the UrlId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickDetailMember) GetUrlIdOk() (*string, bool) {
	if o == nil || IsNil(o.UrlId) {
		return nil, false
	}
	return o.UrlId, true
}

// HasUrlId returns a boolean if a field has been set.
func (o *ClickDetailMember) HasUrlId() bool {
	if o != nil && !IsNil(o.UrlId) {
		return true
	}

	return false
}

// SetUrlId gets a reference to the given string and assigns it to the UrlId field.
func (o *ClickDetailMember) SetUrlId(v string) {
	o.UrlId = &v
}

// GetListId returns the ListId field value if set, zero value otherwise.
func (o *ClickDetailMember) GetListId() string {
	if o == nil || IsNil(o.ListId) {
		var ret string
		return ret
	}
	return *o.ListId
}

// GetListIdOk returns a tuple with the ListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickDetailMember) GetListIdOk() (*string, bool) {
	if o == nil || IsNil(o.ListId) {
		return nil, false
	}
	return o.ListId, true
}

// HasListId returns a boolean if a field has been set.
func (o *ClickDetailMember) HasListId() bool {
	if o != nil && !IsNil(o.ListId) {
		return true
	}

	return false
}

// SetListId gets a reference to the given string and assigns it to the ListId field.
func (o *ClickDetailMember) SetListId(v string) {
	o.ListId = &v
}

// GetListIsActive returns the ListIsActive field value if set, zero value otherwise.
func (o *ClickDetailMember) GetListIsActive() bool {
	if o == nil || IsNil(o.ListIsActive) {
		var ret bool
		return ret
	}
	return *o.ListIsActive
}

// GetListIsActiveOk returns a tuple with the ListIsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickDetailMember) GetListIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.ListIsActive) {
		return nil, false
	}
	return o.ListIsActive, true
}

// HasListIsActive returns a boolean if a field has been set.
func (o *ClickDetailMember) HasListIsActive() bool {
	if o != nil && !IsNil(o.ListIsActive) {
		return true
	}

	return false
}

// SetListIsActive gets a reference to the given bool and assigns it to the ListIsActive field.
func (o *ClickDetailMember) SetListIsActive(v bool) {
	o.ListIsActive = &v
}

// GetContactStatus returns the ContactStatus field value if set, zero value otherwise.
func (o *ClickDetailMember) GetContactStatus() string {
	if o == nil || IsNil(o.ContactStatus) {
		var ret string
		return ret
	}
	return *o.ContactStatus
}

// GetContactStatusOk returns a tuple with the ContactStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickDetailMember) GetContactStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ContactStatus) {
		return nil, false
	}
	return o.ContactStatus, true
}

// HasContactStatus returns a boolean if a field has been set.
func (o *ClickDetailMember) HasContactStatus() bool {
	if o != nil && !IsNil(o.ContactStatus) {
		return true
	}

	return false
}

// SetContactStatus gets a reference to the given string and assigns it to the ContactStatus field.
func (o *ClickDetailMember) SetContactStatus(v string) {
	o.ContactStatus = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ClickDetailMember) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickDetailMember) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ClickDetailMember) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *ClickDetailMember) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o ClickDetailMember) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClickDetailMember) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmailId) {
		toSerialize["email_id"] = o.EmailId
	}
	if !IsNil(o.EmailAddress) {
		toSerialize["email_address"] = o.EmailAddress
	}
	if !IsNil(o.MergeFields) {
		toSerialize["merge_fields"] = o.MergeFields
	}
	if !IsNil(o.Vip) {
		toSerialize["vip"] = o.Vip
	}
	if !IsNil(o.Clicks) {
		toSerialize["clicks"] = o.Clicks
	}
	if !IsNil(o.CampaignId) {
		toSerialize["campaign_id"] = o.CampaignId
	}
	if !IsNil(o.UrlId) {
		toSerialize["url_id"] = o.UrlId
	}
	if !IsNil(o.ListId) {
		toSerialize["list_id"] = o.ListId
	}
	if !IsNil(o.ListIsActive) {
		toSerialize["list_is_active"] = o.ListIsActive
	}
	if !IsNil(o.ContactStatus) {
		toSerialize["contact_status"] = o.ContactStatus
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableClickDetailMember struct {
	value *ClickDetailMember
	isSet bool
}

func (v NullableClickDetailMember) Get() *ClickDetailMember {
	return v.value
}

func (v *NullableClickDetailMember) Set(val *ClickDetailMember) {
	v.value = val
	v.isSet = true
}

func (v NullableClickDetailMember) IsSet() bool {
	return v.isSet
}

func (v *NullableClickDetailMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClickDetailMember(val *ClickDetailMember) *NullableClickDetailMember {
	return &NullableClickDetailMember{value: val, isSet: true}
}

func (v NullableClickDetailMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClickDetailMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


