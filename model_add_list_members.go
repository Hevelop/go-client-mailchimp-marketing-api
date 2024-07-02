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

// checks if the AddListMembers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddListMembers{}

// AddListMembers Individuals who are currently or have been previously subscribed to this list, including members who have bounced or unsubscribed.
type AddListMembers struct {
	// Email address for a subscriber.
	EmailAddress *string `json:"email_address,omitempty"`
	// Type of email this member asked to get ('html' or 'text').
	EmailType *string `json:"email_type,omitempty"`
	// Subscriber's current status.
	Status *string `json:"status,omitempty"`
	// A dictionary of merge fields where the keys are the merge tags. See the [Merge Fields documentation](https://mailchimp.com/developer/marketing/docs/merge-fields/#structure) for more about the structure.
	MergeFields map[string]map[string]interface{} `json:"merge_fields,omitempty"`
	// The key of this object's properties is the ID of the interest in question.
	Interests *map[string]bool `json:"interests,omitempty"`
	// If set/detected, the [subscriber's language](https://mailchimp.com/help/view-and-edit-contact-languages/).
	Language *string `json:"language,omitempty"`
	// [VIP status](https://mailchimp.com/help/designate-and-send-to-vip-contacts/) for subscriber.
	Vip *bool `json:"vip,omitempty"`
	Location *Location `json:"location,omitempty"`
	// IP address the subscriber signed up from.
	IpSignup *string `json:"ip_signup,omitempty"`
	// The date and time the subscriber signed up for the list in ISO 8601 format.
	TimestampSignup *time.Time `json:"timestamp_signup,omitempty"`
	// The IP address the subscriber used to confirm their opt-in status.
	IpOpt *string `json:"ip_opt,omitempty"`
	// The date and time the subscriber confirmed their opt-in status in ISO 8601 format.
	TimestampOpt *time.Time `json:"timestamp_opt,omitempty"`
}

// NewAddListMembers instantiates a new AddListMembers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddListMembers() *AddListMembers {
	this := AddListMembers{}
	return &this
}

// NewAddListMembersWithDefaults instantiates a new AddListMembers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddListMembersWithDefaults() *AddListMembers {
	this := AddListMembers{}
	return &this
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *AddListMembers) GetEmailAddress() string {
	if o == nil || IsNil(o.EmailAddress) {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddListMembers) GetEmailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EmailAddress) {
		return nil, false
	}
	return o.EmailAddress, true
}

// HaveEmailAddress returns a boolean if a field has been set.
func (o *AddListMembers) HaveEmailAddress() bool {
	if o != nil && !IsNil(o.EmailAddress) {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *AddListMembers) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetEmailType returns the EmailType field value if set, zero value otherwise.
func (o *AddListMembers) GetEmailType() string {
	if o == nil || IsNil(o.EmailType) {
		var ret string
		return ret
	}
	return *o.EmailType
}

// GetEmailTypeOk returns a tuple with the EmailType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddListMembers) GetEmailTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EmailType) {
		return nil, false
	}
	return o.EmailType, true
}

// HaveEmailType returns a boolean if a field has been set.
func (o *AddListMembers) HaveEmailType() bool {
	if o != nil && !IsNil(o.EmailType) {
		return true
	}

	return false
}

// SetEmailType gets a reference to the given string and assigns it to the EmailType field.
func (o *AddListMembers) SetEmailType(v string) {
	o.EmailType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AddListMembers) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddListMembers) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HaveStatus returns a boolean if a field has been set.
func (o *AddListMembers) HaveStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AddListMembers) SetStatus(v string) {
	o.Status = &v
}

// GetMergeFields returns the MergeFields field value if set, zero value otherwise.
func (o *AddListMembers) GetMergeFields() map[string]map[string]interface{} {
	if o == nil || IsNil(o.MergeFields) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.MergeFields
}

// GetMergeFieldsOk returns a tuple with the MergeFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddListMembers) GetMergeFieldsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.MergeFields) {
		return map[string]map[string]interface{}{}, false
	}
	return o.MergeFields, true
}

// HaveMergeFields returns a boolean if a field has been set.
func (o *AddListMembers) HaveMergeFields() bool {
	if o != nil && !IsNil(o.MergeFields) {
		return true
	}

	return false
}

// SetMergeFields gets a reference to the given map[string]map[string]interface{} and assigns it to the MergeFields field.
func (o *AddListMembers) SetMergeFields(v map[string]map[string]interface{}) {
	o.MergeFields = v
}

// GetInterests returns the Interests field value if set, zero value otherwise.
func (o *AddListMembers) GetInterests() map[string]bool {
	if o == nil || IsNil(o.Interests) {
		var ret map[string]bool
		return ret
	}
	return *o.Interests
}

// GetInterestsOk returns a tuple with the Interests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddListMembers) GetInterestsOk() (*map[string]bool, bool) {
	if o == nil || IsNil(o.Interests) {
		return nil, false
	}
	return o.Interests, true
}

// HaveInterests returns a boolean if a field has been set.
func (o *AddListMembers) HaveInterests() bool {
	if o != nil && !IsNil(o.Interests) {
		return true
	}

	return false
}

// SetInterests gets a reference to the given map[string]bool and assigns it to the Interests field.
func (o *AddListMembers) SetInterests(v map[string]bool) {
	o.Interests = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *AddListMembers) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddListMembers) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HaveLanguage returns a boolean if a field has been set.
func (o *AddListMembers) HaveLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *AddListMembers) SetLanguage(v string) {
	o.Language = &v
}

// GetVip returns the Vip field value if set, zero value otherwise.
func (o *AddListMembers) GetVip() bool {
	if o == nil || IsNil(o.Vip) {
		var ret bool
		return ret
	}
	return *o.Vip
}

// GetVipOk returns a tuple with the Vip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddListMembers) GetVipOk() (*bool, bool) {
	if o == nil || IsNil(o.Vip) {
		return nil, false
	}
	return o.Vip, true
}

// HaveVip returns a boolean if a field has been set.
func (o *AddListMembers) HaveVip() bool {
	if o != nil && !IsNil(o.Vip) {
		return true
	}

	return false
}

// SetVip gets a reference to the given bool and assigns it to the Vip field.
func (o *AddListMembers) SetVip(v bool) {
	o.Vip = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *AddListMembers) GetLocation() Location {
	if o == nil || IsNil(o.Location) {
		var ret Location
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddListMembers) GetLocationOk() (*Location, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HaveLocation returns a boolean if a field has been set.
func (o *AddListMembers) HaveLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given Location and assigns it to the Location field.
func (o *AddListMembers) SetLocation(v Location) {
	o.Location = &v
}

// GetIpSignup returns the IpSignup field value if set, zero value otherwise.
func (o *AddListMembers) GetIpSignup() string {
	if o == nil || IsNil(o.IpSignup) {
		var ret string
		return ret
	}
	return *o.IpSignup
}

// GetIpSignupOk returns a tuple with the IpSignup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddListMembers) GetIpSignupOk() (*string, bool) {
	if o == nil || IsNil(o.IpSignup) {
		return nil, false
	}
	return o.IpSignup, true
}

// HaveIpSignup returns a boolean if a field has been set.
func (o *AddListMembers) HaveIpSignup() bool {
	if o != nil && !IsNil(o.IpSignup) {
		return true
	}

	return false
}

// SetIpSignup gets a reference to the given string and assigns it to the IpSignup field.
func (o *AddListMembers) SetIpSignup(v string) {
	o.IpSignup = &v
}

// GetTimestampSignup returns the TimestampSignup field value if set, zero value otherwise.
func (o *AddListMembers) GetTimestampSignup() time.Time {
	if o == nil || IsNil(o.TimestampSignup) {
		var ret time.Time
		return ret
	}
	return *o.TimestampSignup
}

// GetTimestampSignupOk returns a tuple with the TimestampSignup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddListMembers) GetTimestampSignupOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TimestampSignup) {
		return nil, false
	}
	return o.TimestampSignup, true
}

// HaveTimestampSignup returns a boolean if a field has been set.
func (o *AddListMembers) HaveTimestampSignup() bool {
	if o != nil && !IsNil(o.TimestampSignup) {
		return true
	}

	return false
}

// SetTimestampSignup gets a reference to the given time.Time and assigns it to the TimestampSignup field.
func (o *AddListMembers) SetTimestampSignup(v time.Time) {
	o.TimestampSignup = &v
}

// GetIpOpt returns the IpOpt field value if set, zero value otherwise.
func (o *AddListMembers) GetIpOpt() string {
	if o == nil || IsNil(o.IpOpt) {
		var ret string
		return ret
	}
	return *o.IpOpt
}

// GetIpOptOk returns a tuple with the IpOpt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddListMembers) GetIpOptOk() (*string, bool) {
	if o == nil || IsNil(o.IpOpt) {
		return nil, false
	}
	return o.IpOpt, true
}

// HaveIpOpt returns a boolean if a field has been set.
func (o *AddListMembers) HaveIpOpt() bool {
	if o != nil && !IsNil(o.IpOpt) {
		return true
	}

	return false
}

// SetIpOpt gets a reference to the given string and assigns it to the IpOpt field.
func (o *AddListMembers) SetIpOpt(v string) {
	o.IpOpt = &v
}

// GetTimestampOpt returns the TimestampOpt field value if set, zero value otherwise.
func (o *AddListMembers) GetTimestampOpt() time.Time {
	if o == nil || IsNil(o.TimestampOpt) {
		var ret time.Time
		return ret
	}
	return *o.TimestampOpt
}

// GetTimestampOptOk returns a tuple with the TimestampOpt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddListMembers) GetTimestampOptOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TimestampOpt) {
		return nil, false
	}
	return o.TimestampOpt, true
}

// HaveTimestampOpt returns a boolean if a field has been set.
func (o *AddListMembers) HaveTimestampOpt() bool {
	if o != nil && !IsNil(o.TimestampOpt) {
		return true
	}

	return false
}

// SetTimestampOpt gets a reference to the given time.Time and assigns it to the TimestampOpt field.
func (o *AddListMembers) SetTimestampOpt(v time.Time) {
	o.TimestampOpt = &v
}

func (o AddListMembers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddListMembers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmailAddress) {
		toSerialize["email_address"] = o.EmailAddress
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
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.Vip) {
		toSerialize["vip"] = o.Vip
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
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
	return toSerialize, nil
}

type NullableAddListMembers struct {
	value *AddListMembers
	isSet bool
}

func (v NullableAddListMembers) Get() *AddListMembers {
	return v.value
}

func (v *NullableAddListMembers) Set(val *AddListMembers) {
	v.value = val
	v.isSet = true
}

func (v NullableAddListMembers) IsSet() bool {
	return v.isSet
}

func (v *NullableAddListMembers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddListMembers(val *AddListMembers) *NullableAddListMembers {
	return &NullableAddListMembers{value: val, isSet: true}
}

func (v NullableAddListMembers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddListMembers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


