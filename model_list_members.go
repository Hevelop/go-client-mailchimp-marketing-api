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
	"time"
)

// checks if the ListMembers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListMembers{}

// ListMembers Individuals who are currently or have been previously subscribed to this list, including members who have bounced or unsubscribed.
type ListMembers struct {
	// The MD5 hash of the lowercase version of the list member's email address.
	Id *string `json:"id,omitempty"`
	// Email address for a subscriber.
	EmailAddress *string `json:"email_address,omitempty"`
	// An identifier for the address across all of Mailchimp.
	UniqueEmailId *string `json:"unique_email_id,omitempty"`
	// Type of email this member asked to get ('html' or 'text').
	EmailType *string `json:"email_type,omitempty"`
	// Subscriber's current status.
	Status *string `json:"status,omitempty"`
	// An individual merge var and value for a member.
	MergeFields map[string]map[string]interface{} `json:"merge_fields,omitempty"`
	// The key of this object's properties is the ID of the interest in question.
	Interests *map[string]bool `json:"interests,omitempty"`
	Stats *SubscriberStats `json:"stats,omitempty"`
	// IP address the subscriber signed up from.
	IpSignup *string `json:"ip_signup,omitempty"`
	// The date and time the subscriber signed up for the list in ISO 8601 format.
	TimestampSignup *time.Time `json:"timestamp_signup,omitempty"`
	// The IP address the subscriber used to confirm their opt-in status.
	IpOpt *string `json:"ip_opt,omitempty"`
	// The date and time the subscriber confirmed their opt-in status in ISO 8601 format.
	TimestampOpt *time.Time `json:"timestamp_opt,omitempty"`
	// Star rating for this member, between 1 and 5.
	MemberRating *int32 `json:"member_rating,omitempty"`
	// The date and time the member's info was last changed in ISO 8601 format.
	LastChanged *time.Time `json:"last_changed,omitempty"`
	// If set/detected, the [subscriber's language](https://mailchimp.com/help/view-and-edit-contact-languages/).
	Language *string `json:"language,omitempty"`
	// [VIP status](https://mailchimp.com/help/designate-and-send-to-vip-contacts/) for subscriber.
	Vip *bool `json:"vip,omitempty"`
	// The list member's email client.
	EmailClient *string `json:"email_client,omitempty"`
	Location *Location1 `json:"location,omitempty"`
	LastNote *Notes `json:"last_note,omitempty"`
	// The number of tags applied to this member.
	TagsCount *int32 `json:"tags_count,omitempty"`
	// The tags applied to this member.
	Tags []TagsInner `json:"tags,omitempty"`
	// The list id.
	ListId *string `json:"list_id,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewListMembers instantiates a new ListMembers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListMembers() *ListMembers {
	this := ListMembers{}
	return &this
}

// NewListMembersWithDefaults instantiates a new ListMembers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListMembersWithDefaults() *ListMembers {
	this := ListMembers{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListMembers) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMembers) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ListMembers) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ListMembers) SetId(v string) {
	o.Id = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *ListMembers) GetEmailAddress() string {
	if o == nil || IsNil(o.EmailAddress) {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMembers) GetEmailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EmailAddress) {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *ListMembers) HasEmailAddress() bool {
	if o != nil && !IsNil(o.EmailAddress) {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *ListMembers) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetUniqueEmailId returns the UniqueEmailId field value if set, zero value otherwise.
func (o *ListMembers) GetUniqueEmailId() string {
	if o == nil || IsNil(o.UniqueEmailId) {
		var ret string
		return ret
	}
	return *o.UniqueEmailId
}

// GetUniqueEmailIdOk returns a tuple with the UniqueEmailId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMembers) GetUniqueEmailIdOk() (*string, bool) {
	if o == nil || IsNil(o.UniqueEmailId) {
		return nil, false
	}
	return o.UniqueEmailId, true
}

// HasUniqueEmailId returns a boolean if a field has been set.
func (o *ListMembers) HasUniqueEmailId() bool {
	if o != nil && !IsNil(o.UniqueEmailId) {
		return true
	}

	return false
}

// SetUniqueEmailId gets a reference to the given string and assigns it to the UniqueEmailId field.
func (o *ListMembers) SetUniqueEmailId(v string) {
	o.UniqueEmailId = &v
}

// GetEmailType returns the EmailType field value if set, zero value otherwise.
func (o *ListMembers) GetEmailType() string {
	if o == nil || IsNil(o.EmailType) {
		var ret string
		return ret
	}
	return *o.EmailType
}

// GetEmailTypeOk returns a tuple with the EmailType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMembers) GetEmailTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EmailType) {
		return nil, false
	}
	return o.EmailType, true
}

// HasEmailType returns a boolean if a field has been set.
func (o *ListMembers) HasEmailType() bool {
	if o != nil && !IsNil(o.EmailType) {
		return true
	}

	return false
}

// SetEmailType gets a reference to the given string and assigns it to the EmailType field.
func (o *ListMembers) SetEmailType(v string) {
	o.EmailType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ListMembers) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMembers) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ListMembers) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ListMembers) SetStatus(v string) {
	o.Status = &v
}

// GetMergeFields returns the MergeFields field value if set, zero value otherwise.
func (o *ListMembers) GetMergeFields() map[string]map[string]interface{} {
	if o == nil || IsNil(o.MergeFields) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.MergeFields
}

// GetMergeFieldsOk returns a tuple with the MergeFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMembers) GetMergeFieldsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.MergeFields) {
		return map[string]map[string]interface{}{}, false
	}
	return o.MergeFields, true
}

// HasMergeFields returns a boolean if a field has been set.
func (o *ListMembers) HasMergeFields() bool {
	if o != nil && !IsNil(o.MergeFields) {
		return true
	}

	return false
}

// SetMergeFields gets a reference to the given map[string]map[string]interface{} and assigns it to the MergeFields field.
func (o *ListMembers) SetMergeFields(v map[string]map[string]interface{}) {
	o.MergeFields = v
}

// GetInterests returns the Interests field value if set, zero value otherwise.
func (o *ListMembers) GetInterests() map[string]bool {
	if o == nil || IsNil(o.Interests) {
		var ret map[string]bool
		return ret
	}
	return *o.Interests
}

// GetInterestsOk returns a tuple with the Interests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMembers) GetInterestsOk() (*map[string]bool, bool) {
	if o == nil || IsNil(o.Interests) {
		return nil, false
	}
	return o.Interests, true
}

// HasInterests returns a boolean if a field has been set.
func (o *ListMembers) HasInterests() bool {
	if o != nil && !IsNil(o.Interests) {
		return true
	}

	return false
}

// SetInterests gets a reference to the given map[string]bool and assigns it to the Interests field.
func (o *ListMembers) SetInterests(v map[string]bool) {
	o.Interests = &v
}

// GetStats returns the Stats field value if set, zero value otherwise.
func (o *ListMembers) GetStats() SubscriberStats {
	if o == nil || IsNil(o.Stats) {
		var ret SubscriberStats
		return ret
	}
	return *o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMembers) GetStatsOk() (*SubscriberStats, bool) {
	if o == nil || IsNil(o.Stats) {
		return nil, false
	}
	return o.Stats, true
}

// HasStats returns a boolean if a field has been set.
func (o *ListMembers) HasStats() bool {
	if o != nil && !IsNil(o.Stats) {
		return true
	}

	return false
}

// SetStats gets a reference to the given SubscriberStats and assigns it to the Stats field.
func (o *ListMembers) SetStats(v SubscriberStats) {
	o.Stats = &v
}

// GetIpSignup returns the IpSignup field value if set, zero value otherwise.
func (o *ListMembers) GetIpSignup() string {
	if o == nil || IsNil(o.IpSignup) {
		var ret string
		return ret
	}
	return *o.IpSignup
}

// GetIpSignupOk returns a tuple with the IpSignup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMembers) GetIpSignupOk() (*string, bool) {
	if o == nil || IsNil(o.IpSignup) {
		return nil, false
	}
	return o.IpSignup, true
}

// HasIpSignup returns a boolean if a field has been set.
func (o *ListMembers) HasIpSignup() bool {
	if o != nil && !IsNil(o.IpSignup) {
		return true
	}

	return false
}

// SetIpSignup gets a reference to the given string and assigns it to the IpSignup field.
func (o *ListMembers) SetIpSignup(v string) {
	o.IpSignup = &v
}

// GetTimestampSignup returns the TimestampSignup field value if set, zero value otherwise.
func (o *ListMembers) GetTimestampSignup() time.Time {
	if o == nil || IsNil(o.TimestampSignup) {
		var ret time.Time
		return ret
	}
	return *o.TimestampSignup
}

// GetTimestampSignupOk returns a tuple with the TimestampSignup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMembers) GetTimestampSignupOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TimestampSignup) {
		return nil, false
	}
	return o.TimestampSignup, true
}

// HasTimestampSignup returns a boolean if a field has been set.
func (o *ListMembers) HasTimestampSignup() bool {
	if o != nil && !IsNil(o.TimestampSignup) {
		return true
	}

	return false
}

// SetTimestampSignup gets a reference to the given time.Time and assigns it to the TimestampSignup field.
func (o *ListMembers) SetTimestampSignup(v time.Time) {
	o.TimestampSignup = &v
}

// GetIpOpt returns the IpOpt field value if set, zero value otherwise.
func (o *ListMembers) GetIpOpt() string {
	if o == nil || IsNil(o.IpOpt) {
		var ret string
		return ret
	}
	return *o.IpOpt
}

// GetIpOptOk returns a tuple with the IpOpt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMembers) GetIpOptOk() (*string, bool) {
	if o == nil || IsNil(o.IpOpt) {
		return nil, false
	}
	return o.IpOpt, true
}

// HasIpOpt returns a boolean if a field has been set.
func (o *ListMembers) HasIpOpt() bool {
	if o != nil && !IsNil(o.IpOpt) {
		return true
	}

	return false
}

// SetIpOpt gets a reference to the given string and assigns it to the IpOpt field.
func (o *ListMembers) SetIpOpt(v string) {
	o.IpOpt = &v
}

// GetTimestampOpt returns the TimestampOpt field value if set, zero value otherwise.
func (o *ListMembers) GetTimestampOpt() time.Time {
	if o == nil || IsNil(o.TimestampOpt) {
		var ret time.Time
		return ret
	}
	return *o.TimestampOpt
}

// GetTimestampOptOk returns a tuple with the TimestampOpt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMembers) GetTimestampOptOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TimestampOpt) {
		return nil, false
	}
	return o.TimestampOpt, true
}

// HasTimestampOpt returns a boolean if a field has been set.
func (o *ListMembers) HasTimestampOpt() bool {
	if o != nil && !IsNil(o.TimestampOpt) {
		return true
	}

	return false
}

// SetTimestampOpt gets a reference to the given time.Time and assigns it to the TimestampOpt field.
func (o *ListMembers) SetTimestampOpt(v time.Time) {
	o.TimestampOpt = &v
}

// GetMemberRating returns the MemberRating field value if set, zero value otherwise.
func (o *ListMembers) GetMemberRating() int32 {
	if o == nil || IsNil(o.MemberRating) {
		var ret int32
		return ret
	}
	return *o.MemberRating
}

// GetMemberRatingOk returns a tuple with the MemberRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMembers) GetMemberRatingOk() (*int32, bool) {
	if o == nil || IsNil(o.MemberRating) {
		return nil, false
	}
	return o.MemberRating, true
}

// HasMemberRating returns a boolean if a field has been set.
func (o *ListMembers) HasMemberRating() bool {
	if o != nil && !IsNil(o.MemberRating) {
		return true
	}

	return false
}

// SetMemberRating gets a reference to the given int32 and assigns it to the MemberRating field.
func (o *ListMembers) SetMemberRating(v int32) {
	o.MemberRating = &v
}

// GetLastChanged returns the LastChanged field value if set, zero value otherwise.
func (o *ListMembers) GetLastChanged() time.Time {
	if o == nil || IsNil(o.LastChanged) {
		var ret time.Time
		return ret
	}
	return *o.LastChanged
}

// GetLastChangedOk returns a tuple with the LastChanged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMembers) GetLastChangedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastChanged) {
		return nil, false
	}
	return o.LastChanged, true
}

// HasLastChanged returns a boolean if a field has been set.
func (o *ListMembers) HasLastChanged() bool {
	if o != nil && !IsNil(o.LastChanged) {
		return true
	}

	return false
}

// SetLastChanged gets a reference to the given time.Time and assigns it to the LastChanged field.
func (o *ListMembers) SetLastChanged(v time.Time) {
	o.LastChanged = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *ListMembers) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMembers) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *ListMembers) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *ListMembers) SetLanguage(v string) {
	o.Language = &v
}

// GetVip returns the Vip field value if set, zero value otherwise.
func (o *ListMembers) GetVip() bool {
	if o == nil || IsNil(o.Vip) {
		var ret bool
		return ret
	}
	return *o.Vip
}

// GetVipOk returns a tuple with the Vip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMembers) GetVipOk() (*bool, bool) {
	if o == nil || IsNil(o.Vip) {
		return nil, false
	}
	return o.Vip, true
}

// HasVip returns a boolean if a field has been set.
func (o *ListMembers) HasVip() bool {
	if o != nil && !IsNil(o.Vip) {
		return true
	}

	return false
}

// SetVip gets a reference to the given bool and assigns it to the Vip field.
func (o *ListMembers) SetVip(v bool) {
	o.Vip = &v
}

// GetEmailClient returns the EmailClient field value if set, zero value otherwise.
func (o *ListMembers) GetEmailClient() string {
	if o == nil || IsNil(o.EmailClient) {
		var ret string
		return ret
	}
	return *o.EmailClient
}

// GetEmailClientOk returns a tuple with the EmailClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMembers) GetEmailClientOk() (*string, bool) {
	if o == nil || IsNil(o.EmailClient) {
		return nil, false
	}
	return o.EmailClient, true
}

// HasEmailClient returns a boolean if a field has been set.
func (o *ListMembers) HasEmailClient() bool {
	if o != nil && !IsNil(o.EmailClient) {
		return true
	}

	return false
}

// SetEmailClient gets a reference to the given string and assigns it to the EmailClient field.
func (o *ListMembers) SetEmailClient(v string) {
	o.EmailClient = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ListMembers) GetLocation() Location1 {
	if o == nil || IsNil(o.Location) {
		var ret Location1
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMembers) GetLocationOk() (*Location1, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ListMembers) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given Location1 and assigns it to the Location field.
func (o *ListMembers) SetLocation(v Location1) {
	o.Location = &v
}

// GetLastNote returns the LastNote field value if set, zero value otherwise.
func (o *ListMembers) GetLastNote() Notes {
	if o == nil || IsNil(o.LastNote) {
		var ret Notes
		return ret
	}
	return *o.LastNote
}

// GetLastNoteOk returns a tuple with the LastNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMembers) GetLastNoteOk() (*Notes, bool) {
	if o == nil || IsNil(o.LastNote) {
		return nil, false
	}
	return o.LastNote, true
}

// HasLastNote returns a boolean if a field has been set.
func (o *ListMembers) HasLastNote() bool {
	if o != nil && !IsNil(o.LastNote) {
		return true
	}

	return false
}

// SetLastNote gets a reference to the given Notes and assigns it to the LastNote field.
func (o *ListMembers) SetLastNote(v Notes) {
	o.LastNote = &v
}

// GetTagsCount returns the TagsCount field value if set, zero value otherwise.
func (o *ListMembers) GetTagsCount() int32 {
	if o == nil || IsNil(o.TagsCount) {
		var ret int32
		return ret
	}
	return *o.TagsCount
}

// GetTagsCountOk returns a tuple with the TagsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMembers) GetTagsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TagsCount) {
		return nil, false
	}
	return o.TagsCount, true
}

// HasTagsCount returns a boolean if a field has been set.
func (o *ListMembers) HasTagsCount() bool {
	if o != nil && !IsNil(o.TagsCount) {
		return true
	}

	return false
}

// SetTagsCount gets a reference to the given int32 and assigns it to the TagsCount field.
func (o *ListMembers) SetTagsCount(v int32) {
	o.TagsCount = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ListMembers) GetTags() []TagsInner {
	if o == nil || IsNil(o.Tags) {
		var ret []TagsInner
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMembers) GetTagsOk() ([]TagsInner, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ListMembers) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagsInner and assigns it to the Tags field.
func (o *ListMembers) SetTags(v []TagsInner) {
	o.Tags = v
}

// GetListId returns the ListId field value if set, zero value otherwise.
func (o *ListMembers) GetListId() string {
	if o == nil || IsNil(o.ListId) {
		var ret string
		return ret
	}
	return *o.ListId
}

// GetListIdOk returns a tuple with the ListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMembers) GetListIdOk() (*string, bool) {
	if o == nil || IsNil(o.ListId) {
		return nil, false
	}
	return o.ListId, true
}

// HasListId returns a boolean if a field has been set.
func (o *ListMembers) HasListId() bool {
	if o != nil && !IsNil(o.ListId) {
		return true
	}

	return false
}

// SetListId gets a reference to the given string and assigns it to the ListId field.
func (o *ListMembers) SetListId(v string) {
	o.ListId = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ListMembers) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMembers) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ListMembers) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *ListMembers) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o ListMembers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListMembers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.EmailAddress) {
		toSerialize["email_address"] = o.EmailAddress
	}
	if !IsNil(o.UniqueEmailId) {
		toSerialize["unique_email_id"] = o.UniqueEmailId
	}
	if !IsNil(o.EmailType) {
		toSerialize["email_type"] = o.EmailType
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.MergeFields) {
		toSerialize["merge_fields"] = o.MergeFields
	}
	if !IsNil(o.Interests) {
		toSerialize["interests"] = o.Interests
	}
	if !IsNil(o.Stats) {
		toSerialize["stats"] = o.Stats
	}
	if !IsNil(o.IpSignup) {
		toSerialize["ip_signup"] = o.IpSignup
	}
	if !IsNil(o.TimestampSignup) {
		toSerialize["timestamp_signup"] = o.TimestampSignup
	}
	if !IsNil(o.IpOpt) {
		toSerialize["ip_opt"] = o.IpOpt
	}
	if !IsNil(o.TimestampOpt) {
		toSerialize["timestamp_opt"] = o.TimestampOpt
	}
	if !IsNil(o.MemberRating) {
		toSerialize["member_rating"] = o.MemberRating
	}
	if !IsNil(o.LastChanged) {
		toSerialize["last_changed"] = o.LastChanged
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.Vip) {
		toSerialize["vip"] = o.Vip
	}
	if !IsNil(o.EmailClient) {
		toSerialize["email_client"] = o.EmailClient
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.LastNote) {
		toSerialize["last_note"] = o.LastNote
	}
	if !IsNil(o.TagsCount) {
		toSerialize["tags_count"] = o.TagsCount
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.ListId) {
		toSerialize["list_id"] = o.ListId
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableListMembers struct {
	value *ListMembers
	isSet bool
}

func (v NullableListMembers) Get() *ListMembers {
	return v.value
}

func (v *NullableListMembers) Set(val *ListMembers) {
	v.value = val
	v.isSet = true
}

func (v NullableListMembers) IsSet() bool {
	return v.isSet
}

func (v *NullableListMembers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListMembers(val *ListMembers) *NullableListMembers {
	return &NullableListMembers{value: val, isSet: true}
}

func (v NullableListMembers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListMembers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


