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

// checks if the GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner{}

// GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner struct for GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner
type GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner struct {
	Date *string `json:"date,omitempty"`
	Revenue *float32 `json:"revenue,omitempty"`
}

// NewGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner instantiates a new GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner() *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner {
	this := GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner{}
	return &this
}

// NewGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInnerWithDefaults instantiates a new GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInnerWithDefaults() *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner {
	this := GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner) SetDate(v string) {
	o.Date = &v
}

// GetRevenue returns the Revenue field value if set, zero value otherwise.
func (o *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner) GetRevenue() float32 {
	if o == nil || IsNil(o.Revenue) {
		var ret float32
		return ret
	}
	return *o.Revenue
}

// GetRevenueOk returns a tuple with the Revenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner) GetRevenueOk() (*float32, bool) {
	if o == nil || IsNil(o.Revenue) {
		return nil, false
	}
	return o.Revenue, true
}

// HasRevenue returns a boolean if a field has been set.
func (o *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner) HasRevenue() bool {
	if o != nil && !IsNil(o.Revenue) {
		return true
	}

	return false
}

// SetRevenue gets a reference to the given float32 and assigns it to the Revenue field.
func (o *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner) SetRevenue(v float32) {
	o.Revenue = &v
}

func (o GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.Revenue) {
		toSerialize["revenue"] = o.Revenue
	}
	return toSerialize, nil
}

type NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner struct {
	value *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner
	isSet bool
}

func (v NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner) Get() *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner {
	return v.value
}

func (v *NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner) Set(val *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner(val *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner) *NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner {
	return &NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner{value: val, isSet: true}
}

func (v NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityRevenueInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


