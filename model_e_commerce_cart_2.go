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

// checks if the ECommerceCart2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ECommerceCart2{}

// ECommerceCart2 Information about a specific cart.
type ECommerceCart2 struct {
	Customer *ECommerceCustomer2 `json:"customer,omitempty"`
	// A string that uniquely identifies the campaign associated with a cart.
	CampaignId *string `json:"campaign_id,omitempty"`
	// The URL for the cart. This parameter is required for [Abandoned Cart](https://mailchimp.com/help/create-an-abandoned-cart-email/) automations.
	CheckoutUrl *string `json:"checkout_url,omitempty"`
	// The three-letter ISO 4217 code for the currency that the cart uses.
	CurrencyCode *string `json:"currency_code,omitempty"`
	// The order total for the cart.
	OrderTotal *float32 `json:"order_total,omitempty"`
	// The total tax for the cart.
	TaxTotal *float32 `json:"tax_total,omitempty"`
	// An array of the cart's line items.
	Lines []ECommerceCartLineItem1 `json:"lines,omitempty"`
}

// NewECommerceCart2 instantiates a new ECommerceCart2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewECommerceCart2() *ECommerceCart2 {
	this := ECommerceCart2{}
	return &this
}

// NewECommerceCart2WithDefaults instantiates a new ECommerceCart2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewECommerceCart2WithDefaults() *ECommerceCart2 {
	this := ECommerceCart2{}
	return &this
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *ECommerceCart2) GetCustomer() ECommerceCustomer2 {
	if o == nil || IsNil(o.Customer) {
		var ret ECommerceCustomer2
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCart2) GetCustomerOk() (*ECommerceCustomer2, bool) {
	if o == nil || IsNil(o.Customer) {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *ECommerceCart2) HasCustomer() bool {
	if o != nil && !IsNil(o.Customer) {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given ECommerceCustomer2 and assigns it to the Customer field.
func (o *ECommerceCart2) SetCustomer(v ECommerceCustomer2) {
	o.Customer = &v
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *ECommerceCart2) GetCampaignId() string {
	if o == nil || IsNil(o.CampaignId) {
		var ret string
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCart2) GetCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *ECommerceCart2) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given string and assigns it to the CampaignId field.
func (o *ECommerceCart2) SetCampaignId(v string) {
	o.CampaignId = &v
}

// GetCheckoutUrl returns the CheckoutUrl field value if set, zero value otherwise.
func (o *ECommerceCart2) GetCheckoutUrl() string {
	if o == nil || IsNil(o.CheckoutUrl) {
		var ret string
		return ret
	}
	return *o.CheckoutUrl
}

// GetCheckoutUrlOk returns a tuple with the CheckoutUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCart2) GetCheckoutUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CheckoutUrl) {
		return nil, false
	}
	return o.CheckoutUrl, true
}

// HasCheckoutUrl returns a boolean if a field has been set.
func (o *ECommerceCart2) HasCheckoutUrl() bool {
	if o != nil && !IsNil(o.CheckoutUrl) {
		return true
	}

	return false
}

// SetCheckoutUrl gets a reference to the given string and assigns it to the CheckoutUrl field.
func (o *ECommerceCart2) SetCheckoutUrl(v string) {
	o.CheckoutUrl = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *ECommerceCart2) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCart2) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *ECommerceCart2) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *ECommerceCart2) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetOrderTotal returns the OrderTotal field value if set, zero value otherwise.
func (o *ECommerceCart2) GetOrderTotal() float32 {
	if o == nil || IsNil(o.OrderTotal) {
		var ret float32
		return ret
	}
	return *o.OrderTotal
}

// GetOrderTotalOk returns a tuple with the OrderTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCart2) GetOrderTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.OrderTotal) {
		return nil, false
	}
	return o.OrderTotal, true
}

// HasOrderTotal returns a boolean if a field has been set.
func (o *ECommerceCart2) HasOrderTotal() bool {
	if o != nil && !IsNil(o.OrderTotal) {
		return true
	}

	return false
}

// SetOrderTotal gets a reference to the given float32 and assigns it to the OrderTotal field.
func (o *ECommerceCart2) SetOrderTotal(v float32) {
	o.OrderTotal = &v
}

// GetTaxTotal returns the TaxTotal field value if set, zero value otherwise.
func (o *ECommerceCart2) GetTaxTotal() float32 {
	if o == nil || IsNil(o.TaxTotal) {
		var ret float32
		return ret
	}
	return *o.TaxTotal
}

// GetTaxTotalOk returns a tuple with the TaxTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCart2) GetTaxTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.TaxTotal) {
		return nil, false
	}
	return o.TaxTotal, true
}

// HasTaxTotal returns a boolean if a field has been set.
func (o *ECommerceCart2) HasTaxTotal() bool {
	if o != nil && !IsNil(o.TaxTotal) {
		return true
	}

	return false
}

// SetTaxTotal gets a reference to the given float32 and assigns it to the TaxTotal field.
func (o *ECommerceCart2) SetTaxTotal(v float32) {
	o.TaxTotal = &v
}

// GetLines returns the Lines field value if set, zero value otherwise.
func (o *ECommerceCart2) GetLines() []ECommerceCartLineItem1 {
	if o == nil || IsNil(o.Lines) {
		var ret []ECommerceCartLineItem1
		return ret
	}
	return o.Lines
}

// GetLinesOk returns a tuple with the Lines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCart2) GetLinesOk() ([]ECommerceCartLineItem1, bool) {
	if o == nil || IsNil(o.Lines) {
		return nil, false
	}
	return o.Lines, true
}

// HasLines returns a boolean if a field has been set.
func (o *ECommerceCart2) HasLines() bool {
	if o != nil && !IsNil(o.Lines) {
		return true
	}

	return false
}

// SetLines gets a reference to the given []ECommerceCartLineItem1 and assigns it to the Lines field.
func (o *ECommerceCart2) SetLines(v []ECommerceCartLineItem1) {
	o.Lines = v
}

func (o ECommerceCart2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ECommerceCart2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Customer) {
		toSerialize["customer"] = o.Customer
	}
	if !IsNil(o.CampaignId) {
		toSerialize["campaign_id"] = o.CampaignId
	}
	if !IsNil(o.CheckoutUrl) {
		toSerialize["checkout_url"] = o.CheckoutUrl
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if !IsNil(o.OrderTotal) {
		toSerialize["order_total"] = o.OrderTotal
	}
	if !IsNil(o.TaxTotal) {
		toSerialize["tax_total"] = o.TaxTotal
	}
	if !IsNil(o.Lines) {
		toSerialize["lines"] = o.Lines
	}
	return toSerialize, nil
}

type NullableECommerceCart2 struct {
	value *ECommerceCart2
	isSet bool
}

func (v NullableECommerceCart2) Get() *ECommerceCart2 {
	return v.value
}

func (v *NullableECommerceCart2) Set(val *ECommerceCart2) {
	v.value = val
	v.isSet = true
}

func (v NullableECommerceCart2) IsSet() bool {
	return v.isSet
}

func (v *NullableECommerceCart2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECommerceCart2(val *ECommerceCart2) *NullableECommerceCart2 {
	return &NullableECommerceCart2{value: val, isSet: true}
}

func (v NullableECommerceCart2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECommerceCart2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


