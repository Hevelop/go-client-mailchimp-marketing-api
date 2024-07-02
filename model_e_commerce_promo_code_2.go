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

// checks if the ECommercePromoCode2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ECommercePromoCode2{}

// ECommercePromoCode2 Information about an Ecommerce Store's specific Promo Code.
type ECommercePromoCode2 struct {
	// The discount code. Restricted to UTF-8 characters with max length 50.
	Code *string `json:"code,omitempty"`
	// The url that should be used in the promotion campaign restricted to UTF-8 characters with max length 2000.
	RedemptionUrl *string `json:"redemption_url,omitempty"`
	// Number of times promo code has been used.
	UsageCount *int32 `json:"usage_count,omitempty"`
	// Whether the promo code is currently enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The date and time the promotion was created in ISO 8601 format.
	CreatedAtForeign *time.Time `json:"created_at_foreign,omitempty"`
	// The date and time the promotion was updated in ISO 8601 format.
	UpdatedAtForeign *time.Time `json:"updated_at_foreign,omitempty"`
}

// NewECommercePromoCode2 instantiates a new ECommercePromoCode2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewECommercePromoCode2() *ECommercePromoCode2 {
	this := ECommercePromoCode2{}
	return &this
}

// NewECommercePromoCode2WithDefaults instantiates a new ECommercePromoCode2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewECommercePromoCode2WithDefaults() *ECommercePromoCode2 {
	this := ECommercePromoCode2{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ECommercePromoCode2) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommercePromoCode2) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ECommercePromoCode2) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ECommercePromoCode2) SetCode(v string) {
	o.Code = &v
}

// GetRedemptionUrl returns the RedemptionUrl field value if set, zero value otherwise.
func (o *ECommercePromoCode2) GetRedemptionUrl() string {
	if o == nil || IsNil(o.RedemptionUrl) {
		var ret string
		return ret
	}
	return *o.RedemptionUrl
}

// GetRedemptionUrlOk returns a tuple with the RedemptionUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommercePromoCode2) GetRedemptionUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RedemptionUrl) {
		return nil, false
	}
	return o.RedemptionUrl, true
}

// HasRedemptionUrl returns a boolean if a field has been set.
func (o *ECommercePromoCode2) HasRedemptionUrl() bool {
	if o != nil && !IsNil(o.RedemptionUrl) {
		return true
	}

	return false
}

// SetRedemptionUrl gets a reference to the given string and assigns it to the RedemptionUrl field.
func (o *ECommercePromoCode2) SetRedemptionUrl(v string) {
	o.RedemptionUrl = &v
}

// GetUsageCount returns the UsageCount field value if set, zero value otherwise.
func (o *ECommercePromoCode2) GetUsageCount() int32 {
	if o == nil || IsNil(o.UsageCount) {
		var ret int32
		return ret
	}
	return *o.UsageCount
}

// GetUsageCountOk returns a tuple with the UsageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommercePromoCode2) GetUsageCountOk() (*int32, bool) {
	if o == nil || IsNil(o.UsageCount) {
		return nil, false
	}
	return o.UsageCount, true
}

// HasUsageCount returns a boolean if a field has been set.
func (o *ECommercePromoCode2) HasUsageCount() bool {
	if o != nil && !IsNil(o.UsageCount) {
		return true
	}

	return false
}

// SetUsageCount gets a reference to the given int32 and assigns it to the UsageCount field.
func (o *ECommercePromoCode2) SetUsageCount(v int32) {
	o.UsageCount = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ECommercePromoCode2) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommercePromoCode2) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ECommercePromoCode2) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ECommercePromoCode2) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetCreatedAtForeign returns the CreatedAtForeign field value if set, zero value otherwise.
func (o *ECommercePromoCode2) GetCreatedAtForeign() time.Time {
	if o == nil || IsNil(o.CreatedAtForeign) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAtForeign
}

// GetCreatedAtForeignOk returns a tuple with the CreatedAtForeign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommercePromoCode2) GetCreatedAtForeignOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAtForeign) {
		return nil, false
	}
	return o.CreatedAtForeign, true
}

// HasCreatedAtForeign returns a boolean if a field has been set.
func (o *ECommercePromoCode2) HasCreatedAtForeign() bool {
	if o != nil && !IsNil(o.CreatedAtForeign) {
		return true
	}

	return false
}

// SetCreatedAtForeign gets a reference to the given time.Time and assigns it to the CreatedAtForeign field.
func (o *ECommercePromoCode2) SetCreatedAtForeign(v time.Time) {
	o.CreatedAtForeign = &v
}

// GetUpdatedAtForeign returns the UpdatedAtForeign field value if set, zero value otherwise.
func (o *ECommercePromoCode2) GetUpdatedAtForeign() time.Time {
	if o == nil || IsNil(o.UpdatedAtForeign) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAtForeign
}

// GetUpdatedAtForeignOk returns a tuple with the UpdatedAtForeign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommercePromoCode2) GetUpdatedAtForeignOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAtForeign) {
		return nil, false
	}
	return o.UpdatedAtForeign, true
}

// HasUpdatedAtForeign returns a boolean if a field has been set.
func (o *ECommercePromoCode2) HasUpdatedAtForeign() bool {
	if o != nil && !IsNil(o.UpdatedAtForeign) {
		return true
	}

	return false
}

// SetUpdatedAtForeign gets a reference to the given time.Time and assigns it to the UpdatedAtForeign field.
func (o *ECommercePromoCode2) SetUpdatedAtForeign(v time.Time) {
	o.UpdatedAtForeign = &v
}

func (o ECommercePromoCode2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ECommercePromoCode2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.RedemptionUrl) {
		toSerialize["redemption_url"] = o.RedemptionUrl
	}
	if !IsNil(o.UsageCount) {
		toSerialize["usage_count"] = o.UsageCount
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.CreatedAtForeign) {
		toSerialize["created_at_foreign"] = o.CreatedAtForeign
	}
	if !IsNil(o.UpdatedAtForeign) {
		toSerialize["updated_at_foreign"] = o.UpdatedAtForeign
	}
	return toSerialize, nil
}

type NullableECommercePromoCode2 struct {
	value *ECommercePromoCode2
	isSet bool
}

func (v NullableECommercePromoCode2) Get() *ECommercePromoCode2 {
	return v.value
}

func (v *NullableECommercePromoCode2) Set(val *ECommercePromoCode2) {
	v.value = val
	v.isSet = true
}

func (v NullableECommercePromoCode2) IsSet() bool {
	return v.isSet
}

func (v *NullableECommercePromoCode2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECommercePromoCode2(val *ECommercePromoCode2) *NullableECommercePromoCode2 {
	return &NullableECommercePromoCode2{value: val, isSet: true}
}

func (v NullableECommercePromoCode2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECommercePromoCode2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


