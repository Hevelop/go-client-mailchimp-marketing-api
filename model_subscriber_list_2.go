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
	"bytes"
	"fmt"
)

// checks if the SubscriberList2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriberList2{}

// SubscriberList2 Information about a specific list.
type SubscriberList2 struct {
	// The name of the list.
	Name string `json:"name"`
	Contact ListContact2 `json:"contact"`
	// The [permission reminder](https://mailchimp.com/help/edit-the-permission-reminder/) for the list.
	PermissionReminder string `json:"permission_reminder"`
	// Whether campaigns for this list use the [Archive Bar](https://mailchimp.com/help/about-email-campaign-archives-and-pages/) in archives by default.
	UseArchiveBar *bool `json:"use_archive_bar,omitempty"`
	CampaignDefaults CampaignDefaults1 `json:"campaign_defaults"`
	// The email address to send [subscribe notifications](https://mailchimp.com/help/change-subscribe-and-unsubscribe-notifications/) to.
	NotifyOnSubscribe *string `json:"notify_on_subscribe,omitempty"`
	// The email address to send [unsubscribe notifications](https://mailchimp.com/help/change-subscribe-and-unsubscribe-notifications/) to.
	NotifyOnUnsubscribe *string `json:"notify_on_unsubscribe,omitempty"`
	// Whether the list supports [multiple formats for emails](https://mailchimp.com/help/change-audience-name-defaults/). When set to `true`, subscribers can choose whether they want to receive HTML or plain-text emails. When set to `false`, subscribers will receive HTML emails, with a plain-text alternative backup.
	EmailTypeOption bool `json:"email_type_option"`
	// Whether or not to require the subscriber to confirm subscription via email.
	DoubleOptin *bool `json:"double_optin,omitempty"`
	// Whether or not the list has marketing permissions (eg. GDPR) enabled.
	MarketingPermissions *bool `json:"marketing_permissions,omitempty"`
}

type _SubscriberList2 SubscriberList2

// NewSubscriberList2 instantiates a new SubscriberList2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriberList2(name string, contact ListContact2, permissionReminder string, campaignDefaults CampaignDefaults1, emailTypeOption bool) *SubscriberList2 {
	this := SubscriberList2{}
	this.Name = name
	this.Contact = contact
	this.PermissionReminder = permissionReminder
	var useArchiveBar bool = false
	this.UseArchiveBar = &useArchiveBar
	this.CampaignDefaults = campaignDefaults
	var notifyOnSubscribe string = "false"
	this.NotifyOnSubscribe = &notifyOnSubscribe
	var notifyOnUnsubscribe string = "false"
	this.NotifyOnUnsubscribe = &notifyOnUnsubscribe
	this.EmailTypeOption = emailTypeOption
	var doubleOptin bool = false
	this.DoubleOptin = &doubleOptin
	var marketingPermissions bool = false
	this.MarketingPermissions = &marketingPermissions
	return &this
}

// NewSubscriberList2WithDefaults instantiates a new SubscriberList2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriberList2WithDefaults() *SubscriberList2 {
	this := SubscriberList2{}
	var useArchiveBar bool = false
	this.UseArchiveBar = &useArchiveBar
	var notifyOnSubscribe string = "false"
	this.NotifyOnSubscribe = &notifyOnSubscribe
	var notifyOnUnsubscribe string = "false"
	this.NotifyOnUnsubscribe = &notifyOnUnsubscribe
	var doubleOptin bool = false
	this.DoubleOptin = &doubleOptin
	var marketingPermissions bool = false
	this.MarketingPermissions = &marketingPermissions
	return &this
}

// GetName returns the Name field value
func (o *SubscriberList2) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SubscriberList2) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SubscriberList2) SetName(v string) {
	o.Name = v
}

// GetContact returns the Contact field value
func (o *SubscriberList2) GetContact() ListContact2 {
	if o == nil {
		var ret ListContact2
		return ret
	}

	return o.Contact
}

// GetContactOk returns a tuple with the Contact field value
// and a boolean to check if the value has been set.
func (o *SubscriberList2) GetContactOk() (*ListContact2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contact, true
}

// SetContact sets field value
func (o *SubscriberList2) SetContact(v ListContact2) {
	o.Contact = v
}

// GetPermissionReminder returns the PermissionReminder field value
func (o *SubscriberList2) GetPermissionReminder() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PermissionReminder
}

// GetPermissionReminderOk returns a tuple with the PermissionReminder field value
// and a boolean to check if the value has been set.
func (o *SubscriberList2) GetPermissionReminderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PermissionReminder, true
}

// SetPermissionReminder sets field value
func (o *SubscriberList2) SetPermissionReminder(v string) {
	o.PermissionReminder = v
}

// GetUseArchiveBar returns the UseArchiveBar field value if set, zero value otherwise.
func (o *SubscriberList2) GetUseArchiveBar() bool {
	if o == nil || IsNil(o.UseArchiveBar) {
		var ret bool
		return ret
	}
	return *o.UseArchiveBar
}

// GetUseArchiveBarOk returns a tuple with the UseArchiveBar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberList2) GetUseArchiveBarOk() (*bool, bool) {
	if o == nil || IsNil(o.UseArchiveBar) {
		return nil, false
	}
	return o.UseArchiveBar, true
}

// HasUseArchiveBar returns a boolean if a field has been set.
func (o *SubscriberList2) HasUseArchiveBar() bool {
	if o != nil && !IsNil(o.UseArchiveBar) {
		return true
	}

	return false
}

// SetUseArchiveBar gets a reference to the given bool and assigns it to the UseArchiveBar field.
func (o *SubscriberList2) SetUseArchiveBar(v bool) {
	o.UseArchiveBar = &v
}

// GetCampaignDefaults returns the CampaignDefaults field value
func (o *SubscriberList2) GetCampaignDefaults() CampaignDefaults1 {
	if o == nil {
		var ret CampaignDefaults1
		return ret
	}

	return o.CampaignDefaults
}

// GetCampaignDefaultsOk returns a tuple with the CampaignDefaults field value
// and a boolean to check if the value has been set.
func (o *SubscriberList2) GetCampaignDefaultsOk() (*CampaignDefaults1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignDefaults, true
}

// SetCampaignDefaults sets field value
func (o *SubscriberList2) SetCampaignDefaults(v CampaignDefaults1) {
	o.CampaignDefaults = v
}

// GetNotifyOnSubscribe returns the NotifyOnSubscribe field value if set, zero value otherwise.
func (o *SubscriberList2) GetNotifyOnSubscribe() string {
	if o == nil || IsNil(o.NotifyOnSubscribe) {
		var ret string
		return ret
	}
	return *o.NotifyOnSubscribe
}

// GetNotifyOnSubscribeOk returns a tuple with the NotifyOnSubscribe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberList2) GetNotifyOnSubscribeOk() (*string, bool) {
	if o == nil || IsNil(o.NotifyOnSubscribe) {
		return nil, false
	}
	return o.NotifyOnSubscribe, true
}

// HasNotifyOnSubscribe returns a boolean if a field has been set.
func (o *SubscriberList2) HasNotifyOnSubscribe() bool {
	if o != nil && !IsNil(o.NotifyOnSubscribe) {
		return true
	}

	return false
}

// SetNotifyOnSubscribe gets a reference to the given string and assigns it to the NotifyOnSubscribe field.
func (o *SubscriberList2) SetNotifyOnSubscribe(v string) {
	o.NotifyOnSubscribe = &v
}

// GetNotifyOnUnsubscribe returns the NotifyOnUnsubscribe field value if set, zero value otherwise.
func (o *SubscriberList2) GetNotifyOnUnsubscribe() string {
	if o == nil || IsNil(o.NotifyOnUnsubscribe) {
		var ret string
		return ret
	}
	return *o.NotifyOnUnsubscribe
}

// GetNotifyOnUnsubscribeOk returns a tuple with the NotifyOnUnsubscribe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberList2) GetNotifyOnUnsubscribeOk() (*string, bool) {
	if o == nil || IsNil(o.NotifyOnUnsubscribe) {
		return nil, false
	}
	return o.NotifyOnUnsubscribe, true
}

// HasNotifyOnUnsubscribe returns a boolean if a field has been set.
func (o *SubscriberList2) HasNotifyOnUnsubscribe() bool {
	if o != nil && !IsNil(o.NotifyOnUnsubscribe) {
		return true
	}

	return false
}

// SetNotifyOnUnsubscribe gets a reference to the given string and assigns it to the NotifyOnUnsubscribe field.
func (o *SubscriberList2) SetNotifyOnUnsubscribe(v string) {
	o.NotifyOnUnsubscribe = &v
}

// GetEmailTypeOption returns the EmailTypeOption field value
func (o *SubscriberList2) GetEmailTypeOption() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EmailTypeOption
}

// GetEmailTypeOptionOk returns a tuple with the EmailTypeOption field value
// and a boolean to check if the value has been set.
func (o *SubscriberList2) GetEmailTypeOptionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailTypeOption, true
}

// SetEmailTypeOption sets field value
func (o *SubscriberList2) SetEmailTypeOption(v bool) {
	o.EmailTypeOption = v
}

// GetDoubleOptin returns the DoubleOptin field value if set, zero value otherwise.
func (o *SubscriberList2) GetDoubleOptin() bool {
	if o == nil || IsNil(o.DoubleOptin) {
		var ret bool
		return ret
	}
	return *o.DoubleOptin
}

// GetDoubleOptinOk returns a tuple with the DoubleOptin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberList2) GetDoubleOptinOk() (*bool, bool) {
	if o == nil || IsNil(o.DoubleOptin) {
		return nil, false
	}
	return o.DoubleOptin, true
}

// HasDoubleOptin returns a boolean if a field has been set.
func (o *SubscriberList2) HasDoubleOptin() bool {
	if o != nil && !IsNil(o.DoubleOptin) {
		return true
	}

	return false
}

// SetDoubleOptin gets a reference to the given bool and assigns it to the DoubleOptin field.
func (o *SubscriberList2) SetDoubleOptin(v bool) {
	o.DoubleOptin = &v
}

// GetMarketingPermissions returns the MarketingPermissions field value if set, zero value otherwise.
func (o *SubscriberList2) GetMarketingPermissions() bool {
	if o == nil || IsNil(o.MarketingPermissions) {
		var ret bool
		return ret
	}
	return *o.MarketingPermissions
}

// GetMarketingPermissionsOk returns a tuple with the MarketingPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberList2) GetMarketingPermissionsOk() (*bool, bool) {
	if o == nil || IsNil(o.MarketingPermissions) {
		return nil, false
	}
	return o.MarketingPermissions, true
}

// HasMarketingPermissions returns a boolean if a field has been set.
func (o *SubscriberList2) HasMarketingPermissions() bool {
	if o != nil && !IsNil(o.MarketingPermissions) {
		return true
	}

	return false
}

// SetMarketingPermissions gets a reference to the given bool and assigns it to the MarketingPermissions field.
func (o *SubscriberList2) SetMarketingPermissions(v bool) {
	o.MarketingPermissions = &v
}

func (o SubscriberList2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriberList2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["contact"] = o.Contact
	toSerialize["permission_reminder"] = o.PermissionReminder
	if !IsNil(o.UseArchiveBar) {
		toSerialize["use_archive_bar"] = o.UseArchiveBar
	}
	toSerialize["campaign_defaults"] = o.CampaignDefaults
	if !IsNil(o.NotifyOnSubscribe) {
		toSerialize["notify_on_subscribe"] = o.NotifyOnSubscribe
	}
	if !IsNil(o.NotifyOnUnsubscribe) {
		toSerialize["notify_on_unsubscribe"] = o.NotifyOnUnsubscribe
	}
	toSerialize["email_type_option"] = o.EmailTypeOption
	if !IsNil(o.DoubleOptin) {
		toSerialize["double_optin"] = o.DoubleOptin
	}
	if !IsNil(o.MarketingPermissions) {
		toSerialize["marketing_permissions"] = o.MarketingPermissions
	}
	return toSerialize, nil
}

func (o *SubscriberList2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"contact",
		"permission_reminder",
		"campaign_defaults",
		"email_type_option",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSubscriberList2 := _SubscriberList2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriberList2)

	if err != nil {
		return err
	}

	*o = SubscriberList2(varSubscriberList2)

	return err
}

type NullableSubscriberList2 struct {
	value *SubscriberList2
	isSet bool
}

func (v NullableSubscriberList2) Get() *SubscriberList2 {
	return v.value
}

func (v *NullableSubscriberList2) Set(val *SubscriberList2) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriberList2) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriberList2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriberList2(val *SubscriberList2) *NullableSubscriberList2 {
	return &NullableSubscriberList2{value: val, isSet: true}
}

func (v NullableSubscriberList2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriberList2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


