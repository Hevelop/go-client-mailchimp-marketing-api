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

// checks if the CampaignA type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignA{}

// CampaignA Stats for Campaign A.
type CampaignA struct {
	// Bounces for Campaign A.
	Bounces *int32 `json:"bounces,omitempty"`
	// Abuse reports for Campaign A.
	AbuseReports *int32 `json:"abuse_reports,omitempty"`
	// Unsubscribes for Campaign A.
	Unsubs *int32 `json:"unsubs,omitempty"`
	// Recipient Clicks for Campaign A.
	RecipientClicks *int32 `json:"recipient_clicks,omitempty"`
	// Forwards for Campaign A.
	Forwards *int32 `json:"forwards,omitempty"`
	// Opens from forwards for Campaign A.
	ForwardsOpens *int32 `json:"forwards_opens,omitempty"`
	// Opens for Campaign A.
	Opens *int32 `json:"opens,omitempty"`
	// The last open for Campaign A.
	LastOpen *string `json:"last_open,omitempty"`
	// Unique opens for Campaign A.
	UniqueOpens *int32 `json:"unique_opens,omitempty"`
}

// NewCampaignA instantiates a new CampaignA object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignA() *CampaignA {
	this := CampaignA{}
	return &this
}

// NewCampaignAWithDefaults instantiates a new CampaignA object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignAWithDefaults() *CampaignA {
	this := CampaignA{}
	return &this
}

// GetBounces returns the Bounces field value if set, zero value otherwise.
func (o *CampaignA) GetBounces() int32 {
	if o == nil || IsNil(o.Bounces) {
		var ret int32
		return ret
	}
	return *o.Bounces
}

// GetBouncesOk returns a tuple with the Bounces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignA) GetBouncesOk() (*int32, bool) {
	if o == nil || IsNil(o.Bounces) {
		return nil, false
	}
	return o.Bounces, true
}

// HaveBounces returns a boolean if a field has been set.
func (o *CampaignA) HaveBounces() bool {
	if o != nil && !IsNil(o.Bounces) {
		return true
	}

	return false
}

// SetBounces gets a reference to the given int32 and assigns it to the Bounces field.
func (o *CampaignA) SetBounces(v int32) {
	o.Bounces = &v
}

// GetAbuseReports returns the AbuseReports field value if set, zero value otherwise.
func (o *CampaignA) GetAbuseReports() int32 {
	if o == nil || IsNil(o.AbuseReports) {
		var ret int32
		return ret
	}
	return *o.AbuseReports
}

// GetAbuseReportsOk returns a tuple with the AbuseReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignA) GetAbuseReportsOk() (*int32, bool) {
	if o == nil || IsNil(o.AbuseReports) {
		return nil, false
	}
	return o.AbuseReports, true
}

// HaveAbuseReports returns a boolean if a field has been set.
func (o *CampaignA) HaveAbuseReports() bool {
	if o != nil && !IsNil(o.AbuseReports) {
		return true
	}

	return false
}

// SetAbuseReports gets a reference to the given int32 and assigns it to the AbuseReports field.
func (o *CampaignA) SetAbuseReports(v int32) {
	o.AbuseReports = &v
}

// GetUnsubs returns the Unsubs field value if set, zero value otherwise.
func (o *CampaignA) GetUnsubs() int32 {
	if o == nil || IsNil(o.Unsubs) {
		var ret int32
		return ret
	}
	return *o.Unsubs
}

// GetUnsubsOk returns a tuple with the Unsubs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignA) GetUnsubsOk() (*int32, bool) {
	if o == nil || IsNil(o.Unsubs) {
		return nil, false
	}
	return o.Unsubs, true
}

// HaveUnsubs returns a boolean if a field has been set.
func (o *CampaignA) HaveUnsubs() bool {
	if o != nil && !IsNil(o.Unsubs) {
		return true
	}

	return false
}

// SetUnsubs gets a reference to the given int32 and assigns it to the Unsubs field.
func (o *CampaignA) SetUnsubs(v int32) {
	o.Unsubs = &v
}

// GetRecipientClicks returns the RecipientClicks field value if set, zero value otherwise.
func (o *CampaignA) GetRecipientClicks() int32 {
	if o == nil || IsNil(o.RecipientClicks) {
		var ret int32
		return ret
	}
	return *o.RecipientClicks
}

// GetRecipientClicksOk returns a tuple with the RecipientClicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignA) GetRecipientClicksOk() (*int32, bool) {
	if o == nil || IsNil(o.RecipientClicks) {
		return nil, false
	}
	return o.RecipientClicks, true
}

// HaveRecipientClicks returns a boolean if a field has been set.
func (o *CampaignA) HaveRecipientClicks() bool {
	if o != nil && !IsNil(o.RecipientClicks) {
		return true
	}

	return false
}

// SetRecipientClicks gets a reference to the given int32 and assigns it to the RecipientClicks field.
func (o *CampaignA) SetRecipientClicks(v int32) {
	o.RecipientClicks = &v
}

// GetForwards returns the Forwards field value if set, zero value otherwise.
func (o *CampaignA) GetForwards() int32 {
	if o == nil || IsNil(o.Forwards) {
		var ret int32
		return ret
	}
	return *o.Forwards
}

// GetForwardsOk returns a tuple with the Forwards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignA) GetForwardsOk() (*int32, bool) {
	if o == nil || IsNil(o.Forwards) {
		return nil, false
	}
	return o.Forwards, true
}

// HaveForwards returns a boolean if a field has been set.
func (o *CampaignA) HaveForwards() bool {
	if o != nil && !IsNil(o.Forwards) {
		return true
	}

	return false
}

// SetForwards gets a reference to the given int32 and assigns it to the Forwards field.
func (o *CampaignA) SetForwards(v int32) {
	o.Forwards = &v
}

// GetForwardsOpens returns the ForwardsOpens field value if set, zero value otherwise.
func (o *CampaignA) GetForwardsOpens() int32 {
	if o == nil || IsNil(o.ForwardsOpens) {
		var ret int32
		return ret
	}
	return *o.ForwardsOpens
}

// GetForwardsOpensOk returns a tuple with the ForwardsOpens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignA) GetForwardsOpensOk() (*int32, bool) {
	if o == nil || IsNil(o.ForwardsOpens) {
		return nil, false
	}
	return o.ForwardsOpens, true
}

// HaveForwardsOpens returns a boolean if a field has been set.
func (o *CampaignA) HaveForwardsOpens() bool {
	if o != nil && !IsNil(o.ForwardsOpens) {
		return true
	}

	return false
}

// SetForwardsOpens gets a reference to the given int32 and assigns it to the ForwardsOpens field.
func (o *CampaignA) SetForwardsOpens(v int32) {
	o.ForwardsOpens = &v
}

// GetOpens returns the Opens field value if set, zero value otherwise.
func (o *CampaignA) GetOpens() int32 {
	if o == nil || IsNil(o.Opens) {
		var ret int32
		return ret
	}
	return *o.Opens
}

// GetOpensOk returns a tuple with the Opens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignA) GetOpensOk() (*int32, bool) {
	if o == nil || IsNil(o.Opens) {
		return nil, false
	}
	return o.Opens, true
}

// HaveOpens returns a boolean if a field has been set.
func (o *CampaignA) HaveOpens() bool {
	if o != nil && !IsNil(o.Opens) {
		return true
	}

	return false
}

// SetOpens gets a reference to the given int32 and assigns it to the Opens field.
func (o *CampaignA) SetOpens(v int32) {
	o.Opens = &v
}

// GetLastOpen returns the LastOpen field value if set, zero value otherwise.
func (o *CampaignA) GetLastOpen() string {
	if o == nil || IsNil(o.LastOpen) {
		var ret string
		return ret
	}
	return *o.LastOpen
}

// GetLastOpenOk returns a tuple with the LastOpen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignA) GetLastOpenOk() (*string, bool) {
	if o == nil || IsNil(o.LastOpen) {
		return nil, false
	}
	return o.LastOpen, true
}

// HaveLastOpen returns a boolean if a field has been set.
func (o *CampaignA) HaveLastOpen() bool {
	if o != nil && !IsNil(o.LastOpen) {
		return true
	}

	return false
}

// SetLastOpen gets a reference to the given string and assigns it to the LastOpen field.
func (o *CampaignA) SetLastOpen(v string) {
	o.LastOpen = &v
}

// GetUniqueOpens returns the UniqueOpens field value if set, zero value otherwise.
func (o *CampaignA) GetUniqueOpens() int32 {
	if o == nil || IsNil(o.UniqueOpens) {
		var ret int32
		return ret
	}
	return *o.UniqueOpens
}

// GetUniqueOpensOk returns a tuple with the UniqueOpens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignA) GetUniqueOpensOk() (*int32, bool) {
	if o == nil || IsNil(o.UniqueOpens) {
		return nil, false
	}
	return o.UniqueOpens, true
}

// HaveUniqueOpens returns a boolean if a field has been set.
func (o *CampaignA) HaveUniqueOpens() bool {
	if o != nil && !IsNil(o.UniqueOpens) {
		return true
	}

	return false
}

// SetUniqueOpens gets a reference to the given int32 and assigns it to the UniqueOpens field.
func (o *CampaignA) SetUniqueOpens(v int32) {
	o.UniqueOpens = &v
}

func (o CampaignA) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignA) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bounces) {
		toSerialize["bounces"] = o.Bounces
	}
	if !IsNil(o.AbuseReports) {
		toSerialize["abuse_reports"] = o.AbuseReports
	}
	if !IsNil(o.Unsubs) {
		toSerialize["unsubs"] = o.Unsubs
	}
	if !IsNil(o.RecipientClicks) {
		toSerialize["recipient_clicks"] = o.RecipientClicks
	}
	if !IsNil(o.Forwards) {
		toSerialize["forwards"] = o.Forwards
	}
	if !IsNil(o.ForwardsOpens) {
		toSerialize["forwards_opens"] = o.ForwardsOpens
	}
	if !IsNil(o.Opens) {
		toSerialize["opens"] = o.Opens
	}
	if !IsNil(o.LastOpen) {
		toSerialize["last_open"] = o.LastOpen
	}
	if !IsNil(o.UniqueOpens) {
		toSerialize["unique_opens"] = o.UniqueOpens
	}
	return toSerialize, nil
}

type NullableCampaignA struct {
	value *CampaignA
	isSet bool
}

func (v NullableCampaignA) Get() *CampaignA {
	return v.value
}

func (v *NullableCampaignA) Set(val *CampaignA) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignA) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignA) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignA(val *CampaignA) *NullableCampaignA {
	return &NullableCampaignA{value: val, isSet: true}
}

func (v NullableCampaignA) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignA) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


