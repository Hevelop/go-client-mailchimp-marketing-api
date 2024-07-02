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

// checks if the ECommerceStore2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ECommerceStore2{}

// ECommerceStore2 An individual store in an account.
type ECommerceStore2 struct {
	// The name of the store.
	Name *string `json:"name,omitempty"`
	// The e-commerce platform of the store.
	Platform *string `json:"platform,omitempty"`
	// The store domain.
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
}

// NewECommerceStore2 instantiates a new ECommerceStore2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewECommerceStore2() *ECommerceStore2 {
	this := ECommerceStore2{}
	return &this
}

// NewECommerceStore2WithDefaults instantiates a new ECommerceStore2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewECommerceStore2WithDefaults() *ECommerceStore2 {
	this := ECommerceStore2{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ECommerceStore2) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceStore2) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HaveName returns a boolean if a field has been set.
func (o *ECommerceStore2) HaveName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ECommerceStore2) SetName(v string) {
	o.Name = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *ECommerceStore2) GetPlatform() string {
	if o == nil || IsNil(o.Platform) {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceStore2) GetPlatformOk() (*string, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HavePlatform returns a boolean if a field has been set.
func (o *ECommerceStore2) HavePlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *ECommerceStore2) SetPlatform(v string) {
	o.Platform = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *ECommerceStore2) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceStore2) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HaveDomain returns a boolean if a field has been set.
func (o *ECommerceStore2) HaveDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *ECommerceStore2) SetDomain(v string) {
	o.Domain = &v
}

// GetIsSyncing returns the IsSyncing field value if set, zero value otherwise.
func (o *ECommerceStore2) GetIsSyncing() bool {
	if o == nil || IsNil(o.IsSyncing) {
		var ret bool
		return ret
	}
	return *o.IsSyncing
}

// GetIsSyncingOk returns a tuple with the IsSyncing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceStore2) GetIsSyncingOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSyncing) {
		return nil, false
	}
	return o.IsSyncing, true
}

// HaveIsSyncing returns a boolean if a field has been set.
func (o *ECommerceStore2) HaveIsSyncing() bool {
	if o != nil && !IsNil(o.IsSyncing) {
		return true
	}

	return false
}

// SetIsSyncing gets a reference to the given bool and assigns it to the IsSyncing field.
func (o *ECommerceStore2) SetIsSyncing(v bool) {
	o.IsSyncing = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *ECommerceStore2) GetEmailAddress() string {
	if o == nil || IsNil(o.EmailAddress) {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceStore2) GetEmailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EmailAddress) {
		return nil, false
	}
	return o.EmailAddress, true
}

// HaveEmailAddress returns a boolean if a field has been set.
func (o *ECommerceStore2) HaveEmailAddress() bool {
	if o != nil && !IsNil(o.EmailAddress) {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *ECommerceStore2) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *ECommerceStore2) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceStore2) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HaveCurrencyCode returns a boolean if a field has been set.
func (o *ECommerceStore2) HaveCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *ECommerceStore2) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetMoneyFormat returns the MoneyFormat field value if set, zero value otherwise.
func (o *ECommerceStore2) GetMoneyFormat() string {
	if o == nil || IsNil(o.MoneyFormat) {
		var ret string
		return ret
	}
	return *o.MoneyFormat
}

// GetMoneyFormatOk returns a tuple with the MoneyFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceStore2) GetMoneyFormatOk() (*string, bool) {
	if o == nil || IsNil(o.MoneyFormat) {
		return nil, false
	}
	return o.MoneyFormat, true
}

// HaveMoneyFormat returns a boolean if a field has been set.
func (o *ECommerceStore2) HaveMoneyFormat() bool {
	if o != nil && !IsNil(o.MoneyFormat) {
		return true
	}

	return false
}

// SetMoneyFormat gets a reference to the given string and assigns it to the MoneyFormat field.
func (o *ECommerceStore2) SetMoneyFormat(v string) {
	o.MoneyFormat = &v
}

// GetPrimaryLocale returns the PrimaryLocale field value if set, zero value otherwise.
func (o *ECommerceStore2) GetPrimaryLocale() string {
	if o == nil || IsNil(o.PrimaryLocale) {
		var ret string
		return ret
	}
	return *o.PrimaryLocale
}

// GetPrimaryLocaleOk returns a tuple with the PrimaryLocale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceStore2) GetPrimaryLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryLocale) {
		return nil, false
	}
	return o.PrimaryLocale, true
}

// HavePrimaryLocale returns a boolean if a field has been set.
func (o *ECommerceStore2) HavePrimaryLocale() bool {
	if o != nil && !IsNil(o.PrimaryLocale) {
		return true
	}

	return false
}

// SetPrimaryLocale gets a reference to the given string and assigns it to the PrimaryLocale field.
func (o *ECommerceStore2) SetPrimaryLocale(v string) {
	o.PrimaryLocale = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *ECommerceStore2) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceStore2) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HaveTimezone returns a boolean if a field has been set.
func (o *ECommerceStore2) HaveTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *ECommerceStore2) SetTimezone(v string) {
	o.Timezone = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *ECommerceStore2) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceStore2) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HavePhone returns a boolean if a field has been set.
func (o *ECommerceStore2) HavePhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *ECommerceStore2) SetPhone(v string) {
	o.Phone = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ECommerceStore2) GetAddress() Address1 {
	if o == nil || IsNil(o.Address) {
		var ret Address1
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceStore2) GetAddressOk() (*Address1, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HaveAddress returns a boolean if a field has been set.
func (o *ECommerceStore2) HaveAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address1 and assigns it to the Address field.
func (o *ECommerceStore2) SetAddress(v Address1) {
	o.Address = &v
}

func (o ECommerceStore2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ECommerceStore2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	return toSerialize, nil
}

type NullableECommerceStore2 struct {
	value *ECommerceStore2
	isSet bool
}

func (v NullableECommerceStore2) Get() *ECommerceStore2 {
	return v.value
}

func (v *NullableECommerceStore2) Set(val *ECommerceStore2) {
	v.value = val
	v.isSet = true
}

func (v NullableECommerceStore2) IsSet() bool {
	return v.isSet
}

func (v *NullableECommerceStore2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECommerceStore2(val *ECommerceStore2) *NullableECommerceStore2 {
	return &NullableECommerceStore2{value: val, isSet: true}
}

func (v NullableECommerceStore2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECommerceStore2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


