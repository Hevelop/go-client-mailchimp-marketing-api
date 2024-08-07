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

// checks if the GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce{}

// GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce struct for GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce
type GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce struct {
	TotalRevenue *float32 `json:"total_revenue,omitempty"`
	CurrencyCode *string `json:"currency_code,omitempty"`
}

// NewGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce instantiates a new GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce() *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce {
	this := GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce{}
	return &this
}

// NewGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerceWithDefaults instantiates a new GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerceWithDefaults() *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce {
	this := GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce{}
	return &this
}

// GetTotalRevenue returns the TotalRevenue field value if set, zero value otherwise.
func (o *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce) GetTotalRevenue() float32 {
	if o == nil || IsNil(o.TotalRevenue) {
		var ret float32
		return ret
	}
	return *o.TotalRevenue
}

// GetTotalRevenueOk returns a tuple with the TotalRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce) GetTotalRevenueOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalRevenue) {
		return nil, false
	}
	return o.TotalRevenue, true
}

// HasTotalRevenue returns a boolean if a field has been set.
func (o *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce) HasTotalRevenue() bool {
	if o != nil && !IsNil(o.TotalRevenue) {
		return true
	}

	return false
}

// SetTotalRevenue gets a reference to the given float32 and assigns it to the TotalRevenue field.
func (o *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce) SetTotalRevenue(v float32) {
	o.TotalRevenue = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

func (o GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalRevenue) {
		toSerialize["total_revenue"] = o.TotalRevenue
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	return toSerialize, nil
}

type NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce struct {
	value *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce
	isSet bool
}

func (v NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce) Get() *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce {
	return v.value
}

func (v *NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce) Set(val *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce) {
	v.value = val
	v.isSet = true
}

func (v NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce) IsSet() bool {
	return v.isSet
}

func (v *NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce(val *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce) *NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce {
	return &NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce{value: val, isSet: true}
}

func (v NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfReportSummaryEcommerce) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


