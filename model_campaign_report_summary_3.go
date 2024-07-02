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

// checks if the CampaignReportSummary3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignReportSummary3{}

// CampaignReportSummary3 For sent campaigns, a summary of opens and clicks.
type CampaignReportSummary3 struct {
	// The total number of opens for a campaign.
	Opens *int32 `json:"opens,omitempty"`
	// The number of unique opens.
	UniqueOpens *int32 `json:"unique_opens,omitempty"`
	// The number of unique opens divided by the total number of successful deliveries.
	OpenRate *float32 `json:"open_rate,omitempty"`
	// The total number of clicks for an campaign.
	Clicks *int32 `json:"clicks,omitempty"`
	// The number of unique clicks.
	SubscriberClicks *int32 `json:"subscriber_clicks,omitempty"`
	// The number of unique clicks divided by the total number of successful deliveries.
	ClickRate *float32 `json:"click_rate,omitempty"`
	Ecommerce *ECommerceReport `json:"ecommerce,omitempty"`
}

// NewCampaignReportSummary3 instantiates a new CampaignReportSummary3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignReportSummary3() *CampaignReportSummary3 {
	this := CampaignReportSummary3{}
	return &this
}

// NewCampaignReportSummary3WithDefaults instantiates a new CampaignReportSummary3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignReportSummary3WithDefaults() *CampaignReportSummary3 {
	this := CampaignReportSummary3{}
	return &this
}

// GetOpens returns the Opens field value if set, zero value otherwise.
func (o *CampaignReportSummary3) GetOpens() int32 {
	if o == nil || IsNil(o.Opens) {
		var ret int32
		return ret
	}
	return *o.Opens
}

// GetOpensOk returns a tuple with the Opens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignReportSummary3) GetOpensOk() (*int32, bool) {
	if o == nil || IsNil(o.Opens) {
		return nil, false
	}
	return o.Opens, true
}

// HaveOpens returns a boolean if a field has been set.
func (o *CampaignReportSummary3) HaveOpens() bool {
	if o != nil && !IsNil(o.Opens) {
		return true
	}

	return false
}

// SetOpens gets a reference to the given int32 and assigns it to the Opens field.
func (o *CampaignReportSummary3) SetOpens(v int32) {
	o.Opens = &v
}

// GetUniqueOpens returns the UniqueOpens field value if set, zero value otherwise.
func (o *CampaignReportSummary3) GetUniqueOpens() int32 {
	if o == nil || IsNil(o.UniqueOpens) {
		var ret int32
		return ret
	}
	return *o.UniqueOpens
}

// GetUniqueOpensOk returns a tuple with the UniqueOpens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignReportSummary3) GetUniqueOpensOk() (*int32, bool) {
	if o == nil || IsNil(o.UniqueOpens) {
		return nil, false
	}
	return o.UniqueOpens, true
}

// HaveUniqueOpens returns a boolean if a field has been set.
func (o *CampaignReportSummary3) HaveUniqueOpens() bool {
	if o != nil && !IsNil(o.UniqueOpens) {
		return true
	}

	return false
}

// SetUniqueOpens gets a reference to the given int32 and assigns it to the UniqueOpens field.
func (o *CampaignReportSummary3) SetUniqueOpens(v int32) {
	o.UniqueOpens = &v
}

// GetOpenRate returns the OpenRate field value if set, zero value otherwise.
func (o *CampaignReportSummary3) GetOpenRate() float32 {
	if o == nil || IsNil(o.OpenRate) {
		var ret float32
		return ret
	}
	return *o.OpenRate
}

// GetOpenRateOk returns a tuple with the OpenRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignReportSummary3) GetOpenRateOk() (*float32, bool) {
	if o == nil || IsNil(o.OpenRate) {
		return nil, false
	}
	return o.OpenRate, true
}

// HaveOpenRate returns a boolean if a field has been set.
func (o *CampaignReportSummary3) HaveOpenRate() bool {
	if o != nil && !IsNil(o.OpenRate) {
		return true
	}

	return false
}

// SetOpenRate gets a reference to the given float32 and assigns it to the OpenRate field.
func (o *CampaignReportSummary3) SetOpenRate(v float32) {
	o.OpenRate = &v
}

// GetClicks returns the Clicks field value if set, zero value otherwise.
func (o *CampaignReportSummary3) GetClicks() int32 {
	if o == nil || IsNil(o.Clicks) {
		var ret int32
		return ret
	}
	return *o.Clicks
}

// GetClicksOk returns a tuple with the Clicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignReportSummary3) GetClicksOk() (*int32, bool) {
	if o == nil || IsNil(o.Clicks) {
		return nil, false
	}
	return o.Clicks, true
}

// HaveClicks returns a boolean if a field has been set.
func (o *CampaignReportSummary3) HaveClicks() bool {
	if o != nil && !IsNil(o.Clicks) {
		return true
	}

	return false
}

// SetClicks gets a reference to the given int32 and assigns it to the Clicks field.
func (o *CampaignReportSummary3) SetClicks(v int32) {
	o.Clicks = &v
}

// GetSubscriberClicks returns the SubscriberClicks field value if set, zero value otherwise.
func (o *CampaignReportSummary3) GetSubscriberClicks() int32 {
	if o == nil || IsNil(o.SubscriberClicks) {
		var ret int32
		return ret
	}
	return *o.SubscriberClicks
}

// GetSubscriberClicksOk returns a tuple with the SubscriberClicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignReportSummary3) GetSubscriberClicksOk() (*int32, bool) {
	if o == nil || IsNil(o.SubscriberClicks) {
		return nil, false
	}
	return o.SubscriberClicks, true
}

// HaveSubscriberClicks returns a boolean if a field has been set.
func (o *CampaignReportSummary3) HaveSubscriberClicks() bool {
	if o != nil && !IsNil(o.SubscriberClicks) {
		return true
	}

	return false
}

// SetSubscriberClicks gets a reference to the given int32 and assigns it to the SubscriberClicks field.
func (o *CampaignReportSummary3) SetSubscriberClicks(v int32) {
	o.SubscriberClicks = &v
}

// GetClickRate returns the ClickRate field value if set, zero value otherwise.
func (o *CampaignReportSummary3) GetClickRate() float32 {
	if o == nil || IsNil(o.ClickRate) {
		var ret float32
		return ret
	}
	return *o.ClickRate
}

// GetClickRateOk returns a tuple with the ClickRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignReportSummary3) GetClickRateOk() (*float32, bool) {
	if o == nil || IsNil(o.ClickRate) {
		return nil, false
	}
	return o.ClickRate, true
}

// HaveClickRate returns a boolean if a field has been set.
func (o *CampaignReportSummary3) HaveClickRate() bool {
	if o != nil && !IsNil(o.ClickRate) {
		return true
	}

	return false
}

// SetClickRate gets a reference to the given float32 and assigns it to the ClickRate field.
func (o *CampaignReportSummary3) SetClickRate(v float32) {
	o.ClickRate = &v
}

// GetEcommerce returns the Ecommerce field value if set, zero value otherwise.
func (o *CampaignReportSummary3) GetEcommerce() ECommerceReport {
	if o == nil || IsNil(o.Ecommerce) {
		var ret ECommerceReport
		return ret
	}
	return *o.Ecommerce
}

// GetEcommerceOk returns a tuple with the Ecommerce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignReportSummary3) GetEcommerceOk() (*ECommerceReport, bool) {
	if o == nil || IsNil(o.Ecommerce) {
		return nil, false
	}
	return o.Ecommerce, true
}

// HaveEcommerce returns a boolean if a field has been set.
func (o *CampaignReportSummary3) HaveEcommerce() bool {
	if o != nil && !IsNil(o.Ecommerce) {
		return true
	}

	return false
}

// SetEcommerce gets a reference to the given ECommerceReport and assigns it to the Ecommerce field.
func (o *CampaignReportSummary3) SetEcommerce(v ECommerceReport) {
	o.Ecommerce = &v
}

func (o CampaignReportSummary3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignReportSummary3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Opens) {
		toSerialize["opens"] = o.Opens
	}
	if !IsNil(o.UniqueOpens) {
		toSerialize["unique_opens"] = o.UniqueOpens
	}
	if !IsNil(o.OpenRate) {
		toSerialize["open_rate"] = o.OpenRate
	}
	if !IsNil(o.Clicks) {
		toSerialize["clicks"] = o.Clicks
	}
	if !IsNil(o.SubscriberClicks) {
		toSerialize["subscriber_clicks"] = o.SubscriberClicks
	}
	if !IsNil(o.ClickRate) {
		toSerialize["click_rate"] = o.ClickRate
	}
	if !IsNil(o.Ecommerce) {
		toSerialize["ecommerce"] = o.Ecommerce
	}
	return toSerialize, nil
}

type NullableCampaignReportSummary3 struct {
	value *CampaignReportSummary3
	isSet bool
}

func (v NullableCampaignReportSummary3) Get() *CampaignReportSummary3 {
	return v.value
}

func (v *NullableCampaignReportSummary3) Set(val *CampaignReportSummary3) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignReportSummary3) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignReportSummary3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignReportSummary3(val *CampaignReportSummary3) *NullableCampaignReportSummary3 {
	return &NullableCampaignReportSummary3{value: val, isSet: true}
}

func (v NullableCampaignReportSummary3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignReportSummary3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


