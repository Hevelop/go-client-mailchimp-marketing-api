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
	"bytes"
	"fmt"
)

// checks if the ECommercePromoCode1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ECommercePromoCode1{}

// ECommercePromoCode1 Information about an Ecommerce Store's specific Promo Code.
type ECommercePromoCode1 struct {
	// A unique identifier for the promo code. Restricted to UTF-8 characters with max length 50.
	Id string `json:"id"`
	// The discount code. Restricted to UTF-8 characters with max length 50.
	Code string `json:"code"`
	// The url that should be used in the promotion campaign restricted to UTF-8 characters with max length 2000.
	RedemptionUrl string `json:"redemption_url"`
	// Number of times promo code has been used.
	UsageCount *int32 `json:"usage_count,omitempty"`
	// Whether the promo code is currently enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The date and time the promotion was created in ISO 8601 format.
	CreatedAtForeign *time.Time `json:"created_at_foreign,omitempty"`
	// The date and time the promotion was updated in ISO 8601 format.
	UpdatedAtForeign *time.Time `json:"updated_at_foreign,omitempty"`
}

type _ECommercePromoCode1 ECommercePromoCode1

// NewECommercePromoCode1 instantiates a new ECommercePromoCode1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewECommercePromoCode1(id string, code string, redemptionUrl string) *ECommercePromoCode1 {
	this := ECommercePromoCode1{}
	this.Id = id
	this.Code = code
	this.RedemptionUrl = redemptionUrl
	return &this
}

// NewECommercePromoCode1WithDefaults instantiates a new ECommercePromoCode1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewECommercePromoCode1WithDefaults() *ECommercePromoCode1 {
	this := ECommercePromoCode1{}
	return &this
}

// GetId returns the Id field value
func (o *ECommercePromoCode1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ECommercePromoCode1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ECommercePromoCode1) SetId(v string) {
	o.Id = v
}

// GetCode returns the Code field value
func (o *ECommercePromoCode1) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ECommercePromoCode1) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ECommercePromoCode1) SetCode(v string) {
	o.Code = v
}

// GetRedemptionUrl returns the RedemptionUrl field value
func (o *ECommercePromoCode1) GetRedemptionUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RedemptionUrl
}

// GetRedemptionUrlOk returns a tuple with the RedemptionUrl field value
// and a boolean to check if the value has been set.
func (o *ECommercePromoCode1) GetRedemptionUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RedemptionUrl, true
}

// SetRedemptionUrl sets field value
func (o *ECommercePromoCode1) SetRedemptionUrl(v string) {
	o.RedemptionUrl = v
}

// GetUsageCount returns the UsageCount field value if set, zero value otherwise.
func (o *ECommercePromoCode1) GetUsageCount() int32 {
	if o == nil || IsNil(o.UsageCount) {
		var ret int32
		return ret
	}
	return *o.UsageCount
}

// GetUsageCountOk returns a tuple with the UsageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommercePromoCode1) GetUsageCountOk() (*int32, bool) {
	if o == nil || IsNil(o.UsageCount) {
		return nil, false
	}
	return o.UsageCount, true
}

// HaveUsageCount returns a boolean if a field has been set.
func (o *ECommercePromoCode1) HaveUsageCount() bool {
	if o != nil && !IsNil(o.UsageCount) {
		return true
	}

	return false
}

// SetUsageCount gets a reference to the given int32 and assigns it to the UsageCount field.
func (o *ECommercePromoCode1) SetUsageCount(v int32) {
	o.UsageCount = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ECommercePromoCode1) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommercePromoCode1) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HaveEnabled returns a boolean if a field has been set.
func (o *ECommercePromoCode1) HaveEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ECommercePromoCode1) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetCreatedAtForeign returns the CreatedAtForeign field value if set, zero value otherwise.
func (o *ECommercePromoCode1) GetCreatedAtForeign() time.Time {
	if o == nil || IsNil(o.CreatedAtForeign) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAtForeign
}

// GetCreatedAtForeignOk returns a tuple with the CreatedAtForeign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommercePromoCode1) GetCreatedAtForeignOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAtForeign) {
		return nil, false
	}
	return o.CreatedAtForeign, true
}

// HaveCreatedAtForeign returns a boolean if a field has been set.
func (o *ECommercePromoCode1) HaveCreatedAtForeign() bool {
	if o != nil && !IsNil(o.CreatedAtForeign) {
		return true
	}

	return false
}

// SetCreatedAtForeign gets a reference to the given time.Time and assigns it to the CreatedAtForeign field.
func (o *ECommercePromoCode1) SetCreatedAtForeign(v time.Time) {
	o.CreatedAtForeign = &v
}

// GetUpdatedAtForeign returns the UpdatedAtForeign field value if set, zero value otherwise.
func (o *ECommercePromoCode1) GetUpdatedAtForeign() time.Time {
	if o == nil || IsNil(o.UpdatedAtForeign) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAtForeign
}

// GetUpdatedAtForeignOk returns a tuple with the UpdatedAtForeign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommercePromoCode1) GetUpdatedAtForeignOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAtForeign) {
		return nil, false
	}
	return o.UpdatedAtForeign, true
}

// HaveUpdatedAtForeign returns a boolean if a field has been set.
func (o *ECommercePromoCode1) HaveUpdatedAtForeign() bool {
	if o != nil && !IsNil(o.UpdatedAtForeign) {
		return true
	}

	return false
}

// SetUpdatedAtForeign gets a reference to the given time.Time and assigns it to the UpdatedAtForeign field.
func (o *ECommercePromoCode1) SetUpdatedAtForeign(v time.Time) {
	o.UpdatedAtForeign = &v
}

func (o ECommercePromoCode1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ECommercePromoCode1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["code"] = o.Code
	toSerialize["redemption_url"] = o.RedemptionUrl
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

func (o *ECommercePromoCode1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"code",
		"redemption_url",
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

	varECommercePromoCode1 := _ECommercePromoCode1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varECommercePromoCode1)

	if err != nil {
		return err
	}

	*o = ECommercePromoCode1(varECommercePromoCode1)

	return err
}

type NullableECommercePromoCode1 struct {
	value *ECommercePromoCode1
	isSet bool
}

func (v NullableECommercePromoCode1) Get() *ECommercePromoCode1 {
	return v.value
}

func (v *NullableECommercePromoCode1) Set(val *ECommercePromoCode1) {
	v.value = val
	v.isSet = true
}

func (v NullableECommercePromoCode1) IsSet() bool {
	return v.isSet
}

func (v *NullableECommercePromoCode1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECommercePromoCode1(val *ECommercePromoCode1) *NullableECommercePromoCode1 {
	return &NullableECommercePromoCode1{value: val, isSet: true}
}

func (v NullableECommercePromoCode1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECommercePromoCode1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


