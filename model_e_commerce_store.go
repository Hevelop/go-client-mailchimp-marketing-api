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

// checks if the ECommerceStore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ECommerceStore{}

// ECommerceStore An individual store in an account.
type ECommerceStore struct {
	// The unique identifier for the store.
	Id *string `json:"id,omitempty"`
	// The unique identifier for the list that's associated with the store. The `list_id` for a specific store can't change.
	ListId *string `json:"list_id,omitempty"`
	// The name of the store.
	Name *string `json:"name,omitempty"`
	// The e-commerce platform of the store.
	Platform *string `json:"platform,omitempty"`
	// The store domain.  The store domain must be unique within a user account.
	Domain *string `json:"domain,omitempty"`
	// Whether to disable automations because the store is currently [syncing](https://mailchimp.com/developer/marketing/docs/e-commerce/#pausing-store-automations).
	IsSyncing *bool `json:"is_syncing,omitempty"`
	// The email address for the store.
	EmailAddress *string `json:"email_address,omitempty"`
	// The three-letter ISO 4217 code for the currency that the store accepts.
	CurrencyCode *string `json:"currency_code,omitempty"`
	// The currency format for the store. For example: `$`, `£`, etc.
	MoneyFormat *string `json:"money_format,omitempty"`
	// The primary locale for the store. For example: `en`, `de`, etc.
	PrimaryLocale *string `json:"primary_locale,omitempty"`
	// The timezone for the store.
	Timezone *string `json:"timezone,omitempty"`
	// The store phone number.
	Phone *string `json:"phone,omitempty"`
	Address *Address1 `json:"address,omitempty"`
	ConnectedSite *ConnectedSite2 `json:"connected_site,omitempty"`
	Automations *Automations `json:"automations,omitempty"`
	// The status of the list connected to the store, namely if it's deleted or disabled.
	ListIsActive *bool `json:"list_is_active,omitempty"`
	// The date and time the store was created in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The date and time the store was last updated in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewECommerceStore instantiates a new ECommerceStore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewECommerceStore() *ECommerceStore {
	this := ECommerceStore{}
	return &this
}

// NewECommerceStoreWithDefaults instantiates a new ECommerceStore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewECommerceStoreWithDefaults() *ECommerceStore {
	this := ECommerceStore{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ECommerceStore) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceStore) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// ContainsId returns a boolean if a field has been set.
func (o *ECommerceStore) ContainsId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ECommerceStore) SetId(v string) {
	o.Id = &v
}

// GetListId returns the ListId field value if set, zero value otherwise.
func (o *ECommerceStore) GetListId() string {
	if o == nil || IsNil(o.ListId) {
		var ret string
		return ret
	}
	return *o.ListId
}

// GetListIdOk returns a tuple with the ListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceStore) GetListIdOk() (*string, bool) {
	if o == nil || IsNil(o.ListId) {
		return nil, false
	}
	return o.ListId, true
}

// ContainsListId returns a boolean if a field has been set.
func (o *ECommerceStore) ContainsListId() bool {
	if o != nil && !IsNil(o.ListId) {
		return true
	}

	return false
}

// SetListId gets a reference to the given string and assigns it to the ListId field.
func (o *ECommerceStore) SetListId(v string) {
	o.ListId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ECommerceStore) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceStore) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// ContainsName returns a boolean if a field has been set.
func (o *ECommerceStore) ContainsName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ECommerceStore) SetName(v string) {
	o.Name = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *ECommerceStore) GetPlatform() string {
	if o == nil || IsNil(o.Platform) {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceStore) GetPlatformOk() (*string, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// ContainsPlatform returns a boolean if a field has been set.
func (o *ECommerceStore) ContainsPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *ECommerceStore) SetPlatform(v string) {
	o.Platform = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *ECommerceStore) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceStore) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// ContainsDomain returns a boolean if a field has been set.
func (o *ECommerceStore) ContainsDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *ECommerceStore) SetDomain(v string) {
	o.Domain = &v
}

// GetIsSyncing returns the IsSyncing field value if set, zero value otherwise.
func (o *ECommerceStore) GetIsSyncing() bool {
	if o == nil || IsNil(o.IsSyncing) {
		var ret bool
		return ret
	}
	return *o.IsSyncing
}

// GetIsSyncingOk returns a tuple with the IsSyncing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceStore) GetIsSyncingOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSyncing) {
		return nil, false
	}
	return o.IsSyncing, true
}

// ContainsIsSyncing returns a boolean if a field has been set.
func (o *ECommerceStore) ContainsIsSyncing() bool {
	if o != nil && !IsNil(o.IsSyncing) {
		return true
	}

	return false
}

// SetIsSyncing gets a reference to the given bool and assigns it to the IsSyncing field.
func (o *ECommerceStore) SetIsSyncing(v bool) {
	o.IsSyncing = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *ECommerceStore) GetEmailAddress() string {
	if o == nil || IsNil(o.EmailAddress) {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceStore) GetEmailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EmailAddress) {
		return nil, false
	}
	return o.EmailAddress, true
}

// ContainsEmailAddress returns a boolean if a field has been set.
func (o *ECommerceStore) ContainsEmailAddress() bool {
	if o != nil && !IsNil(o.EmailAddress) {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *ECommerceStore) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *ECommerceStore) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceStore) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// ContainsCurrencyCode returns a boolean if a field has been set.
func (o *ECommerceStore) ContainsCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *ECommerceStore) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetMoneyFormat returns the MoneyFormat field value if set, zero value otherwise.
func (o *ECommerceStore) GetMoneyFormat() string {
	if o == nil || IsNil(o.MoneyFormat) {
		var ret string
		return ret
	}
	return *o.MoneyFormat
}

// GetMoneyFormatOk returns a tuple with the MoneyFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceStore) GetMoneyFormatOk() (*string, bool) {
	if o == nil || IsNil(o.MoneyFormat) {
		return nil, false
	}
	return o.MoneyFormat, true
}

// ContainsMoneyFormat returns a boolean if a field has been set.
func (o *ECommerceStore) ContainsMoneyFormat() bool {
	if o != nil && !IsNil(o.MoneyFormat) {
		return true
	}

	return false
}

// SetMoneyFormat gets a reference to the given string and assigns it to the MoneyFormat field.
func (o *ECommerceStore) SetMoneyFormat(v string) {
	o.MoneyFormat = &v
}

// GetPrimaryLocale returns the PrimaryLocale field value if set, zero value otherwise.
func (o *ECommerceStore) GetPrimaryLocale() string {
	if o == nil || IsNil(o.PrimaryLocale) {
		var ret string
		return ret
	}
	return *o.PrimaryLocale
}

// GetPrimaryLocaleOk returns a tuple with the PrimaryLocale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceStore) GetPrimaryLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryLocale) {
		return nil, false
	}
	return o.PrimaryLocale, true
}

// ContainsPrimaryLocale returns a boolean if a field has been set.
func (o *ECommerceStore) ContainsPrimaryLocale() bool {
	if o != nil && !IsNil(o.PrimaryLocale) {
		return true
	}

	return false
}

// SetPrimaryLocale gets a reference to the given string and assigns it to the PrimaryLocale field.
func (o *ECommerceStore) SetPrimaryLocale(v string) {
	o.PrimaryLocale = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *ECommerceStore) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceStore) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// ContainsTimezone returns a boolean if a field has been set.
func (o *ECommerceStore) ContainsTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *ECommerceStore) SetTimezone(v string) {
	o.Timezone = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *ECommerceStore) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceStore) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// ContainsPhone returns a boolean if a field has been set.
func (o *ECommerceStore) ContainsPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *ECommerceStore) SetPhone(v string) {
	o.Phone = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ECommerceStore) GetAddress() Address1 {
	if o == nil || IsNil(o.Address) {
		var ret Address1
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceStore) GetAddressOk() (*Address1, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// ContainsAddress returns a boolean if a field has been set.
func (o *ECommerceStore) ContainsAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address1 and assigns it to the Address field.
func (o *ECommerceStore) SetAddress(v Address1) {
	o.Address = &v
}

// GetConnectedSite returns the ConnectedSite field value if set, zero value otherwise.
func (o *ECommerceStore) GetConnectedSite() ConnectedSite2 {
	if o == nil || IsNil(o.ConnectedSite) {
		var ret ConnectedSite2
		return ret
	}
	return *o.ConnectedSite
}

// GetConnectedSiteOk returns a tuple with the ConnectedSite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceStore) GetConnectedSiteOk() (*ConnectedSite2, bool) {
	if o == nil || IsNil(o.ConnectedSite) {
		return nil, false
	}
	return o.ConnectedSite, true
}

// ContainsConnectedSite returns a boolean if a field has been set.
func (o *ECommerceStore) ContainsConnectedSite() bool {
	if o != nil && !IsNil(o.ConnectedSite) {
		return true
	}

	return false
}

// SetConnectedSite gets a reference to the given ConnectedSite2 and assigns it to the ConnectedSite field.
func (o *ECommerceStore) SetConnectedSite(v ConnectedSite2) {
	o.ConnectedSite = &v
}

// GetAutomations returns the Automations field value if set, zero value otherwise.
func (o *ECommerceStore) GetAutomations() Automations {
	if o == nil || IsNil(o.Automations) {
		var ret Automations
		return ret
	}
	return *o.Automations
}

// GetAutomationsOk returns a tuple with the Automations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceStore) GetAutomationsOk() (*Automations, bool) {
	if o == nil || IsNil(o.Automations) {
		return nil, false
	}
	return o.Automations, true
}

// ContainsAutomations returns a boolean if a field has been set.
func (o *ECommerceStore) ContainsAutomations() bool {
	if o != nil && !IsNil(o.Automations) {
		return true
	}

	return false
}

// SetAutomations gets a reference to the given Automations and assigns it to the Automations field.
func (o *ECommerceStore) SetAutomations(v Automations) {
	o.Automations = &v
}

// GetListIsActive returns the ListIsActive field value if set, zero value otherwise.
func (o *ECommerceStore) GetListIsActive() bool {
	if o == nil || IsNil(o.ListIsActive) {
		var ret bool
		return ret
	}
	return *o.ListIsActive
}

// GetListIsActiveOk returns a tuple with the ListIsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceStore) GetListIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.ListIsActive) {
		return nil, false
	}
	return o.ListIsActive, true
}

// ContainsListIsActive returns a boolean if a field has been set.
func (o *ECommerceStore) ContainsListIsActive() bool {
	if o != nil && !IsNil(o.ListIsActive) {
		return true
	}

	return false
}

// SetListIsActive gets a reference to the given bool and assigns it to the ListIsActive field.
func (o *ECommerceStore) SetListIsActive(v bool) {
	o.ListIsActive = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ECommerceStore) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceStore) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// ContainsCreatedAt returns a boolean if a field has been set.
func (o *ECommerceStore) ContainsCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ECommerceStore) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ECommerceStore) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceStore) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// ContainsUpdatedAt returns a boolean if a field has been set.
func (o *ECommerceStore) ContainsUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ECommerceStore) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ECommerceStore) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceStore) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// ContainsLinks returns a boolean if a field has been set.
func (o *ECommerceStore) ContainsLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *ECommerceStore) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o ECommerceStore) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ECommerceStore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ListId) {
		toSerialize["list_id"] = o.ListId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.IsSyncing) {
		toSerialize["is_syncing"] = o.IsSyncing
	}
	if !IsNil(o.EmailAddress) {
		toSerialize["email_address"] = o.EmailAddress
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if !IsNil(o.MoneyFormat) {
		toSerialize["money_format"] = o.MoneyFormat
	}
	if !IsNil(o.PrimaryLocale) {
		toSerialize["primary_locale"] = o.PrimaryLocale
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.ConnectedSite) {
		toSerialize["connected_site"] = o.ConnectedSite
	}
	if !IsNil(o.Automations) {
		toSerialize["automations"] = o.Automations
	}
	if !IsNil(o.ListIsActive) {
		toSerialize["list_is_active"] = o.ListIsActive
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableECommerceStore struct {
	value *ECommerceStore
	isSet bool
}

func (v NullableECommerceStore) Get() *ECommerceStore {
	return v.value
}

func (v *NullableECommerceStore) Set(val *ECommerceStore) {
	v.value = val
	v.isSet = true
}

func (v NullableECommerceStore) IsSet() bool {
	return v.isSet
}

func (v *NullableECommerceStore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECommerceStore(val *ECommerceStore) *NullableECommerceStore {
	return &NullableECommerceStore{value: val, isSet: true}
}

func (v NullableECommerceStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECommerceStore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


