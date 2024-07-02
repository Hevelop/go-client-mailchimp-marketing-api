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

// checks if the OpenActivity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenActivity{}

// OpenActivity A list of a member's opens activity in a specific campaign.
type OpenActivity struct {
	// The unique id for the campaign.
	CampaignId *string `json:"campaign_id,omitempty"`
	// The unique id for the list.
	ListId *string `json:"list_id,omitempty"`
	// The status of the list used, namely if it's deleted or disabled.
	ListIsActive *bool `json:"list_is_active,omitempty"`
	// The status of the member, namely if they are subscribed, unsubscribed, deleted, non-subscribed, transactional, pending, or need reconfirmation.
	ContactStatus *string `json:"contact_status,omitempty"`
	// The MD5 hash of the lowercase version of the list member's email address.
	EmailId *string `json:"email_id,omitempty"`
	// Email address for a subscriber.
	EmailAddress *string `json:"email_address,omitempty"`
	// A dictionary of merge fields where the keys are the merge tags. See the [Merge Fields documentation](https://mailchimp.com/developer/marketing/docs/merge-fields/#structure) for more about the structure.
	MergeFields map[string]map[string]interface{} `json:"merge_fields,omitempty"`
	// [VIP status](https://mailchimp.com/help/designate-and-send-to-vip-contacts/) for subscriber.
	Vip *bool `json:"vip,omitempty"`
	// The total number of times the this campaign was opened by the list member.
	OpensCount *int32 `json:"opens_count,omitempty"`
	// An array of timestamps for each time a list member opened the campaign. If a list member opens an email multiple times, this will return a separate timestamp for each open event.
	Opens []MemberActivity1 `json:"opens,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewOpenActivity instantiates a new OpenActivity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenActivity() *OpenActivity {
	this := OpenActivity{}
	return &this
}

// NewOpenActivityWithDefaults instantiates a new OpenActivity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenActivityWithDefaults() *OpenActivity {
	this := OpenActivity{}
	return &this
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *OpenActivity) GetCampaignId() string {
	if o == nil || IsNil(o.CampaignId) {
		var ret string
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenActivity) GetCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *OpenActivity) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given string and assigns it to the CampaignId field.
func (o *OpenActivity) SetCampaignId(v string) {
	o.CampaignId = &v
}

// GetListId returns the ListId field value if set, zero value otherwise.
func (o *OpenActivity) GetListId() string {
	if o == nil || IsNil(o.ListId) {
		var ret string
		return ret
	}
	return *o.ListId
}

// GetListIdOk returns a tuple with the ListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenActivity) GetListIdOk() (*string, bool) {
	if o == nil || IsNil(o.ListId) {
		return nil, false
	}
	return o.ListId, true
}

// HasListId returns a boolean if a field has been set.
func (o *OpenActivity) HasListId() bool {
	if o != nil && !IsNil(o.ListId) {
		return true
	}

	return false
}

// SetListId gets a reference to the given string and assigns it to the ListId field.
func (o *OpenActivity) SetListId(v string) {
	o.ListId = &v
}

// GetListIsActive returns the ListIsActive field value if set, zero value otherwise.
func (o *OpenActivity) GetListIsActive() bool {
	if o == nil || IsNil(o.ListIsActive) {
		var ret bool
		return ret
	}
	return *o.ListIsActive
}

// GetListIsActiveOk returns a tuple with the ListIsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenActivity) GetListIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.ListIsActive) {
		return nil, false
	}
	return o.ListIsActive, true
}

// HasListIsActive returns a boolean if a field has been set.
func (o *OpenActivity) HasListIsActive() bool {
	if o != nil && !IsNil(o.ListIsActive) {
		return true
	}

	return false
}

// SetListIsActive gets a reference to the given bool and assigns it to the ListIsActive field.
func (o *OpenActivity) SetListIsActive(v bool) {
	o.ListIsActive = &v
}

// GetContactStatus returns the ContactStatus field value if set, zero value otherwise.
func (o *OpenActivity) GetContactStatus() string {
	if o == nil || IsNil(o.ContactStatus) {
		var ret string
		return ret
	}
	return *o.ContactStatus
}

// GetContactStatusOk returns a tuple with the ContactStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenActivity) GetContactStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ContactStatus) {
		return nil, false
	}
	return o.ContactStatus, true
}

// HasContactStatus returns a boolean if a field has been set.
func (o *OpenActivity) HasContactStatus() bool {
	if o != nil && !IsNil(o.ContactStatus) {
		return true
	}

	return false
}

// SetContactStatus gets a reference to the given string and assigns it to the ContactStatus field.
func (o *OpenActivity) SetContactStatus(v string) {
	o.ContactStatus = &v
}

// GetEmailId returns the EmailId field value if set, zero value otherwise.
func (o *OpenActivity) GetEmailId() string {
	if o == nil || IsNil(o.EmailId) {
		var ret string
		return ret
	}
	return *o.EmailId
}

// GetEmailIdOk returns a tuple with the EmailId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenActivity) GetEmailIdOk() (*string, bool) {
	if o == nil || IsNil(o.EmailId) {
		return nil, false
	}
	return o.EmailId, true
}

// HasEmailId returns a boolean if a field has been set.
func (o *OpenActivity) HasEmailId() bool {
	if o != nil && !IsNil(o.EmailId) {
		return true
	}

	return false
}

// SetEmailId gets a reference to the given string and assigns it to the EmailId field.
func (o *OpenActivity) SetEmailId(v string) {
	o.EmailId = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *OpenActivity) GetEmailAddress() string {
	if o == nil || IsNil(o.EmailAddress) {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenActivity) GetEmailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EmailAddress) {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *OpenActivity) HasEmailAddress() bool {
	if o != nil && !IsNil(o.EmailAddress) {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *OpenActivity) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetMergeFields returns the MergeFields field value if set, zero value otherwise.
func (o *OpenActivity) GetMergeFields() map[string]map[string]interface{} {
	if o == nil || IsNil(o.MergeFields) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.MergeFields
}

// GetMergeFieldsOk returns a tuple with the MergeFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenActivity) GetMergeFieldsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.MergeFields) {
		return map[string]map[string]interface{}{}, false
	}
	return o.MergeFields, true
}

// HasMergeFields returns a boolean if a field has been set.
func (o *OpenActivity) HasMergeFields() bool {
	if o != nil && !IsNil(o.MergeFields) {
		return true
	}

	return false
}

// SetMergeFields gets a reference to the given map[string]map[string]interface{} and assigns it to the MergeFields field.
func (o *OpenActivity) SetMergeFields(v map[string]map[string]interface{}) {
	o.MergeFields = v
}

// GetVip returns the Vip field value if set, zero value otherwise.
func (o *OpenActivity) GetVip() bool {
	if o == nil || IsNil(o.Vip) {
		var ret bool
		return ret
	}
	return *o.Vip
}

// GetVipOk returns a tuple with the Vip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenActivity) GetVipOk() (*bool, bool) {
	if o == nil || IsNil(o.Vip) {
		return nil, false
	}
	return o.Vip, true
}

// HasVip returns a boolean if a field has been set.
func (o *OpenActivity) HasVip() bool {
	if o != nil && !IsNil(o.Vip) {
		return true
	}

	return false
}

// SetVip gets a reference to the given bool and assigns it to the Vip field.
func (o *OpenActivity) SetVip(v bool) {
	o.Vip = &v
}

// GetOpensCount returns the OpensCount field value if set, zero value otherwise.
func (o *OpenActivity) GetOpensCount() int32 {
	if o == nil || IsNil(o.OpensCount) {
		var ret int32
		return ret
	}
	return *o.OpensCount
}

// GetOpensCountOk returns a tuple with the OpensCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenActivity) GetOpensCountOk() (*int32, bool) {
	if o == nil || IsNil(o.OpensCount) {
		return nil, false
	}
	return o.OpensCount, true
}

// HasOpensCount returns a boolean if a field has been set.
func (o *OpenActivity) HasOpensCount() bool {
	if o != nil && !IsNil(o.OpensCount) {
		return true
	}

	return false
}

// SetOpensCount gets a reference to the given int32 and assigns it to the OpensCount field.
func (o *OpenActivity) SetOpensCount(v int32) {
	o.OpensCount = &v
}

// GetOpens returns the Opens field value if set, zero value otherwise.
func (o *OpenActivity) GetOpens() []MemberActivity1 {
	if o == nil || IsNil(o.Opens) {
		var ret []MemberActivity1
		return ret
	}
	return o.Opens
}

// GetOpensOk returns a tuple with the Opens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenActivity) GetOpensOk() ([]MemberActivity1, bool) {
	if o == nil || IsNil(o.Opens) {
		return nil, false
	}
	return o.Opens, true
}

// HasOpens returns a boolean if a field has been set.
func (o *OpenActivity) HasOpens() bool {
	if o != nil && !IsNil(o.Opens) {
		return true
	}

	return false
}

// SetOpens gets a reference to the given []MemberActivity1 and assigns it to the Opens field.
func (o *OpenActivity) SetOpens(v []MemberActivity1) {
	o.Opens = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OpenActivity) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenActivity) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OpenActivity) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *OpenActivity) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o OpenActivity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenActivity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignId) {
		toSerialize["campaign_id"] = o.CampaignId
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
	if !IsNil(o.OpensCount) {
		toSerialize["opens_count"] = o.OpensCount
	}
	if !IsNil(o.Opens) {
		toSerialize["opens"] = o.Opens
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableOpenActivity struct {
	value *OpenActivity
	isSet bool
}

func (v NullableOpenActivity) Get() *OpenActivity {
	return v.value
}

func (v *NullableOpenActivity) Set(val *OpenActivity) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenActivity(val *OpenActivity) *NullableOpenActivity {
	return &NullableOpenActivity{value: val, isSet: true}
}

func (v NullableOpenActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


