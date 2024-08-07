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

// checks if the GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner{}

// GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner struct for GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner
type GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner struct {
	Date *string `json:"date,omitempty"`
	Impressions *int32 `json:"impressions,omitempty"`
}

// NewGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner instantiates a new GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner() *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner {
	this := GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner{}
	return &this
}

// NewGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInnerWithDefaults instantiates a new GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInnerWithDefaults() *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner {
	this := GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner) SetDate(v string) {
	o.Date = &v
}

// GetImpressions returns the Impressions field value if set, zero value otherwise.
func (o *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner) GetImpressions() int32 {
	if o == nil || IsNil(o.Impressions) {
		var ret int32
		return ret
	}
	return *o.Impressions
}

// GetImpressionsOk returns a tuple with the Impressions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner) GetImpressionsOk() (*int32, bool) {
	if o == nil || IsNil(o.Impressions) {
		return nil, false
	}
	return o.Impressions, true
}

// HasImpressions returns a boolean if a field has been set.
func (o *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner) HasImpressions() bool {
	if o != nil && !IsNil(o.Impressions) {
		return true
	}

	return false
}

// SetImpressions gets a reference to the given int32 and assigns it to the Impressions field.
func (o *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner) SetImpressions(v int32) {
	o.Impressions = &v
}

func (o GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.Impressions) {
		toSerialize["impressions"] = o.Impressions
	}
	return toSerialize, nil
}

type NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner struct {
	value *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner
	isSet bool
}

func (v NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner) Get() *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner {
	return v.value
}

func (v *NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner) Set(val *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner(val *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner) *NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner {
	return &NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner{value: val, isSet: true}
}

func (v NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityImpressionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


