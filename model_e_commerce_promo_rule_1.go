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

// checks if the ECommercePromoRule1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ECommercePromoRule1{}

// ECommercePromoRule1 Information about an Ecommerce Store's specific Promo Rule.
type ECommercePromoRule1 struct {
	// A unique identifier for the promo rule. If Ecommerce platform does not support promo rule, use promo code id as promo rule id. Restricted to UTF-8 characters with max length 50.
	Id string `json:"id"`
	// The title that will show up in promotion campaign. Restricted to UTF-8 characters with max length of 100 bytes.
	Title *string `json:"title,omitempty"`
	// The description of a promotion restricted to UTF-8 characters with max length 255.
	Description string `json:"description"`
	// The date and time when the promotion is in effect in ISO 8601 format.
	StartsAt *time.Time `json:"starts_at,omitempty"`
	// The date and time when the promotion ends. Must be after starts_at and in ISO 8601 format.
	EndsAt *string `json:"ends_at,omitempty"`
	// The amount of the promo code discount. If 'type' is 'fixed', the amount is treated as a monetary value. If 'type' is 'percentage', amount must be a decimal value between 0.0 and 1.0, inclusive.
	Amount float32 `json:"amount"`
	// Type of discount. For free shipping set type to fixed.
	Type string `json:"type"`
	// The target that the discount applies to.
	Target string `json:"target"`
	// Whether the promo rule is currently enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The date and time the promotion was created in ISO 8601 format.
	CreatedAtForeign *time.Time `json:"created_at_foreign,omitempty"`
	// The date and time the promotion was updated in ISO 8601 format.
	UpdatedAtForeign *time.Time `json:"updated_at_foreign,omitempty"`
}

type _ECommercePromoRule1 ECommercePromoRule1

// NewECommercePromoRule1 instantiates a new ECommercePromoRule1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewECommercePromoRule1(id string, description string, amount float32, type_ string, target string) *ECommercePromoRule1 {
	this := ECommercePromoRule1{}
	this.Id = id
	this.Description = description
	this.Amount = amount
	this.Type = type_
	this.Target = target
	return &this
}

// NewECommercePromoRule1WithDefaults instantiates a new ECommercePromoRule1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewECommercePromoRule1WithDefaults() *ECommercePromoRule1 {
	this := ECommercePromoRule1{}
	return &this
}

// GetId returns the Id field value
func (o *ECommercePromoRule1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ECommercePromoRule1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ECommercePromoRule1) SetId(v string) {
	o.Id = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ECommercePromoRule1) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommercePromoRule1) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HaveTitle returns a boolean if a field has been set.
func (o *ECommercePromoRule1) HaveTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ECommercePromoRule1) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value
func (o *ECommercePromoRule1) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ECommercePromoRule1) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ECommercePromoRule1) SetDescription(v string) {
	o.Description = v
}

// GetStartsAt returns the StartsAt field value if set, zero value otherwise.
func (o *ECommercePromoRule1) GetStartsAt() time.Time {
	if o == nil || IsNil(o.StartsAt) {
		var ret time.Time
		return ret
	}
	return *o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommercePromoRule1) GetStartsAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartsAt) {
		return nil, false
	}
	return o.StartsAt, true
}

// HaveStartsAt returns a boolean if a field has been set.
func (o *ECommercePromoRule1) HaveStartsAt() bool {
	if o != nil && !IsNil(o.StartsAt) {
		return true
	}

	return false
}

// SetStartsAt gets a reference to the given time.Time and assigns it to the StartsAt field.
func (o *ECommercePromoRule1) SetStartsAt(v time.Time) {
	o.StartsAt = &v
}

// GetEndsAt returns the EndsAt field value if set, zero value otherwise.
func (o *ECommercePromoRule1) GetEndsAt() string {
	if o == nil || IsNil(o.EndsAt) {
		var ret string
		return ret
	}
	return *o.EndsAt
}

// GetEndsAtOk returns a tuple with the EndsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommercePromoRule1) GetEndsAtOk() (*string, bool) {
	if o == nil || IsNil(o.EndsAt) {
		return nil, false
	}
	return o.EndsAt, true
}

// HaveEndsAt returns a boolean if a field has been set.
func (o *ECommercePromoRule1) HaveEndsAt() bool {
	if o != nil && !IsNil(o.EndsAt) {
		return true
	}

	return false
}

// SetEndsAt gets a reference to the given string and assigns it to the EndsAt field.
func (o *ECommercePromoRule1) SetEndsAt(v string) {
	o.EndsAt = &v
}

// GetAmount returns the Amount field value
func (o *ECommercePromoRule1) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ECommercePromoRule1) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ECommercePromoRule1) SetAmount(v float32) {
	o.Amount = v
}

// GetType returns the Type field value
func (o *ECommercePromoRule1) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ECommercePromoRule1) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ECommercePromoRule1) SetType(v string) {
	o.Type = v
}

// GetTarget returns the Target field value
func (o *ECommercePromoRule1) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *ECommercePromoRule1) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *ECommercePromoRule1) SetTarget(v string) {
	o.Target = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ECommercePromoRule1) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommercePromoRule1) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HaveEnabled returns a boolean if a field has been set.
func (o *ECommercePromoRule1) HaveEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ECommercePromoRule1) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetCreatedAtForeign returns the CreatedAtForeign field value if set, zero value otherwise.
func (o *ECommercePromoRule1) GetCreatedAtForeign() time.Time {
	if o == nil || IsNil(o.CreatedAtForeign) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAtForeign
}

// GetCreatedAtForeignOk returns a tuple with the CreatedAtForeign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommercePromoRule1) GetCreatedAtForeignOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAtForeign) {
		return nil, false
	}
	return o.CreatedAtForeign, true
}

// HaveCreatedAtForeign returns a boolean if a field has been set.
func (o *ECommercePromoRule1) HaveCreatedAtForeign() bool {
	if o != nil && !IsNil(o.CreatedAtForeign) {
		return true
	}

	return false
}

// SetCreatedAtForeign gets a reference to the given time.Time and assigns it to the CreatedAtForeign field.
func (o *ECommercePromoRule1) SetCreatedAtForeign(v time.Time) {
	o.CreatedAtForeign = &v
}

// GetUpdatedAtForeign returns the UpdatedAtForeign field value if set, zero value otherwise.
func (o *ECommercePromoRule1) GetUpdatedAtForeign() time.Time {
	if o == nil || IsNil(o.UpdatedAtForeign) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAtForeign
}

// GetUpdatedAtForeignOk returns a tuple with the UpdatedAtForeign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommercePromoRule1) GetUpdatedAtForeignOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAtForeign) {
		return nil, false
	}
	return o.UpdatedAtForeign, true
}

// HaveUpdatedAtForeign returns a boolean if a field has been set.
func (o *ECommercePromoRule1) HaveUpdatedAtForeign() bool {
	if o != nil && !IsNil(o.UpdatedAtForeign) {
		return true
	}

	return false
}

// SetUpdatedAtForeign gets a reference to the given time.Time and assigns it to the UpdatedAtForeign field.
func (o *ECommercePromoRule1) SetUpdatedAtForeign(v time.Time) {
	o.UpdatedAtForeign = &v
}

func (o ECommercePromoRule1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ECommercePromoRule1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	toSerialize["description"] = o.Description
	if !IsNil(o.StartsAt) {
		toSerialize["starts_at"] = o.StartsAt
	}
	if !IsNil(o.EndsAt) {
		toSerialize["ends_at"] = o.EndsAt
	}
	toSerialize["amount"] = o.Amount
	toSerialize["type"] = o.Type
	toSerialize["target"] = o.Target
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

func (o *ECommercePromoRule1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"description",
		"amount",
		"type",
		"target",
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

	varECommercePromoRule1 := _ECommercePromoRule1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varECommercePromoRule1)

	if err != nil {
		return err
	}

	*o = ECommercePromoRule1(varECommercePromoRule1)

	return err
}

type NullableECommercePromoRule1 struct {
	value *ECommercePromoRule1
	isSet bool
}

func (v NullableECommercePromoRule1) Get() *ECommercePromoRule1 {
	return v.value
}

func (v *NullableECommercePromoRule1) Set(val *ECommercePromoRule1) {
	v.value = val
	v.isSet = true
}

func (v NullableECommercePromoRule1) IsSet() bool {
	return v.isSet
}

func (v *NullableECommercePromoRule1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECommercePromoRule1(val *ECommercePromoRule1) *NullableECommercePromoRule1 {
	return &NullableECommercePromoRule1{value: val, isSet: true}
}

func (v NullableECommercePromoRule1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECommercePromoRule1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


