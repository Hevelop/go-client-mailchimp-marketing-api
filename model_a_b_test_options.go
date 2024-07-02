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

// checks if the ABTestOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ABTestOptions{}

// ABTestOptions The settings specific to A/B test campaigns.
type ABTestOptions struct {
	// ID for the winning combination.
	WinningCombinationId *string `json:"winning_combination_id,omitempty"`
	// ID of the campaign that was sent to the remaining recipients based on the winning combination.
	WinningCampaignId *string `json:"winning_campaign_id,omitempty"`
	// The combination that performs the best. This may be determined automatically by click rate, open rate, or total revenue -- or you may choose manually based on the reporting data you find the most valuable. For Multivariate Campaigns testing send_time, winner_criteria is ignored. For Multivariate Campaigns with 'manual' as the winner_criteria, the winner must be chosen in the Mailchimp web application.
	WinnerCriteria *string `json:"winner_criteria,omitempty"`
	// The number of minutes to wait before choosing the winning campaign. The value of wait_time must be greater than 0 and in whole hours, specified in minutes.
	WaitTime *int32 `json:"wait_time,omitempty"`
	// The percentage of recipients to send the test combinations to, must be a value between 10 and 100.
	TestSize *int32 `json:"test_size,omitempty"`
	// The possible subject lines to test. If no subject lines are provided, settings.subject_line will be used.
	SubjectLines []string `json:"subject_lines,omitempty"`
	// The possible send times to test. The times provided should be in the format YYYY-MM-DD HH:MM:SS. If send_times are provided to test, the test_size will be set to 100% and winner_criteria will be ignored.
	SendTimes []time.Time `json:"send_times,omitempty"`
	// The possible from names. The number of from_names provided must match the number of reply_to_addresses. If no from_names are provided, settings.from_name will be used.
	FromNames []string `json:"from_names,omitempty"`
	// The possible reply-to addresses. The number of reply_to_addresses provided must match the number of from_names. If no reply_to_addresses are provided, settings.reply_to will be used.
	ReplyToAddresses []string `json:"reply_to_addresses,omitempty"`
	// Descriptions of possible email contents. To set campaign contents, make a PUT request to /campaigns/{campaign_id}/content with the field 'variate_contents'.
	Contents []string `json:"contents,omitempty"`
	// Combinations of possible variables used to build emails.
	Combinations []CombinationsInner `json:"combinations,omitempty"`
}

// NewABTestOptions instantiates a new ABTestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewABTestOptions() *ABTestOptions {
	this := ABTestOptions{}
	return &this
}

// NewABTestOptionsWithDefaults instantiates a new ABTestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewABTestOptionsWithDefaults() *ABTestOptions {
	this := ABTestOptions{}
	return &this
}

// GetWinningCombinationId returns the WinningCombinationId field value if set, zero value otherwise.
func (o *ABTestOptions) GetWinningCombinationId() string {
	if o == nil || IsNil(o.WinningCombinationId) {
		var ret string
		return ret
	}
	return *o.WinningCombinationId
}

// GetWinningCombinationIdOk returns a tuple with the WinningCombinationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ABTestOptions) GetWinningCombinationIdOk() (*string, bool) {
	if o == nil || IsNil(o.WinningCombinationId) {
		return nil, false
	}
	return o.WinningCombinationId, true
}

// HaveWinningCombinationId returns a boolean if a field has been set.
func (o *ABTestOptions) HaveWinningCombinationId() bool {
	if o != nil && !IsNil(o.WinningCombinationId) {
		return true
	}

	return false
}

// SetWinningCombinationId gets a reference to the given string and assigns it to the WinningCombinationId field.
func (o *ABTestOptions) SetWinningCombinationId(v string) {
	o.WinningCombinationId = &v
}

// GetWinningCampaignId returns the WinningCampaignId field value if set, zero value otherwise.
func (o *ABTestOptions) GetWinningCampaignId() string {
	if o == nil || IsNil(o.WinningCampaignId) {
		var ret string
		return ret
	}
	return *o.WinningCampaignId
}

// GetWinningCampaignIdOk returns a tuple with the WinningCampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ABTestOptions) GetWinningCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.WinningCampaignId) {
		return nil, false
	}
	return o.WinningCampaignId, true
}

// HaveWinningCampaignId returns a boolean if a field has been set.
func (o *ABTestOptions) HaveWinningCampaignId() bool {
	if o != nil && !IsNil(o.WinningCampaignId) {
		return true
	}

	return false
}

// SetWinningCampaignId gets a reference to the given string and assigns it to the WinningCampaignId field.
func (o *ABTestOptions) SetWinningCampaignId(v string) {
	o.WinningCampaignId = &v
}

// GetWinnerCriteria returns the WinnerCriteria field value if set, zero value otherwise.
func (o *ABTestOptions) GetWinnerCriteria() string {
	if o == nil || IsNil(o.WinnerCriteria) {
		var ret string
		return ret
	}
	return *o.WinnerCriteria
}

// GetWinnerCriteriaOk returns a tuple with the WinnerCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ABTestOptions) GetWinnerCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.WinnerCriteria) {
		return nil, false
	}
	return o.WinnerCriteria, true
}

// HaveWinnerCriteria returns a boolean if a field has been set.
func (o *ABTestOptions) HaveWinnerCriteria() bool {
	if o != nil && !IsNil(o.WinnerCriteria) {
		return true
	}

	return false
}

// SetWinnerCriteria gets a reference to the given string and assigns it to the WinnerCriteria field.
func (o *ABTestOptions) SetWinnerCriteria(v string) {
	o.WinnerCriteria = &v
}

// GetWaitTime returns the WaitTime field value if set, zero value otherwise.
func (o *ABTestOptions) GetWaitTime() int32 {
	if o == nil || IsNil(o.WaitTime) {
		var ret int32
		return ret
	}
	return *o.WaitTime
}

// GetWaitTimeOk returns a tuple with the WaitTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ABTestOptions) GetWaitTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.WaitTime) {
		return nil, false
	}
	return o.WaitTime, true
}

// HaveWaitTime returns a boolean if a field has been set.
func (o *ABTestOptions) HaveWaitTime() bool {
	if o != nil && !IsNil(o.WaitTime) {
		return true
	}

	return false
}

// SetWaitTime gets a reference to the given int32 and assigns it to the WaitTime field.
func (o *ABTestOptions) SetWaitTime(v int32) {
	o.WaitTime = &v
}

// GetTestSize returns the TestSize field value if set, zero value otherwise.
func (o *ABTestOptions) GetTestSize() int32 {
	if o == nil || IsNil(o.TestSize) {
		var ret int32
		return ret
	}
	return *o.TestSize
}

// GetTestSizeOk returns a tuple with the TestSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ABTestOptions) GetTestSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.TestSize) {
		return nil, false
	}
	return o.TestSize, true
}

// HaveTestSize returns a boolean if a field has been set.
func (o *ABTestOptions) HaveTestSize() bool {
	if o != nil && !IsNil(o.TestSize) {
		return true
	}

	return false
}

// SetTestSize gets a reference to the given int32 and assigns it to the TestSize field.
func (o *ABTestOptions) SetTestSize(v int32) {
	o.TestSize = &v
}

// GetSubjectLines returns the SubjectLines field value if set, zero value otherwise.
func (o *ABTestOptions) GetSubjectLines() []string {
	if o == nil || IsNil(o.SubjectLines) {
		var ret []string
		return ret
	}
	return o.SubjectLines
}

// GetSubjectLinesOk returns a tuple with the SubjectLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ABTestOptions) GetSubjectLinesOk() ([]string, bool) {
	if o == nil || IsNil(o.SubjectLines) {
		return nil, false
	}
	return o.SubjectLines, true
}

// HaveSubjectLines returns a boolean if a field has been set.
func (o *ABTestOptions) HaveSubjectLines() bool {
	if o != nil && !IsNil(o.SubjectLines) {
		return true
	}

	return false
}

// SetSubjectLines gets a reference to the given []string and assigns it to the SubjectLines field.
func (o *ABTestOptions) SetSubjectLines(v []string) {
	o.SubjectLines = v
}

// GetSendTimes returns the SendTimes field value if set, zero value otherwise.
func (o *ABTestOptions) GetSendTimes() []time.Time {
	if o == nil || IsNil(o.SendTimes) {
		var ret []time.Time
		return ret
	}
	return o.SendTimes
}

// GetSendTimesOk returns a tuple with the SendTimes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ABTestOptions) GetSendTimesOk() ([]time.Time, bool) {
	if o == nil || IsNil(o.SendTimes) {
		return nil, false
	}
	return o.SendTimes, true
}

// HaveSendTimes returns a boolean if a field has been set.
func (o *ABTestOptions) HaveSendTimes() bool {
	if o != nil && !IsNil(o.SendTimes) {
		return true
	}

	return false
}

// SetSendTimes gets a reference to the given []time.Time and assigns it to the SendTimes field.
func (o *ABTestOptions) SetSendTimes(v []time.Time) {
	o.SendTimes = v
}

// GetFromNames returns the FromNames field value if set, zero value otherwise.
func (o *ABTestOptions) GetFromNames() []string {
	if o == nil || IsNil(o.FromNames) {
		var ret []string
		return ret
	}
	return o.FromNames
}

// GetFromNamesOk returns a tuple with the FromNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ABTestOptions) GetFromNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.FromNames) {
		return nil, false
	}
	return o.FromNames, true
}

// HaveFromNames returns a boolean if a field has been set.
func (o *ABTestOptions) HaveFromNames() bool {
	if o != nil && !IsNil(o.FromNames) {
		return true
	}

	return false
}

// SetFromNames gets a reference to the given []string and assigns it to the FromNames field.
func (o *ABTestOptions) SetFromNames(v []string) {
	o.FromNames = v
}

// GetReplyToAddresses returns the ReplyToAddresses field value if set, zero value otherwise.
func (o *ABTestOptions) GetReplyToAddresses() []string {
	if o == nil || IsNil(o.ReplyToAddresses) {
		var ret []string
		return ret
	}
	return o.ReplyToAddresses
}

// GetReplyToAddressesOk returns a tuple with the ReplyToAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ABTestOptions) GetReplyToAddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.ReplyToAddresses) {
		return nil, false
	}
	return o.ReplyToAddresses, true
}

// HaveReplyToAddresses returns a boolean if a field has been set.
func (o *ABTestOptions) HaveReplyToAddresses() bool {
	if o != nil && !IsNil(o.ReplyToAddresses) {
		return true
	}

	return false
}

// SetReplyToAddresses gets a reference to the given []string and assigns it to the ReplyToAddresses field.
func (o *ABTestOptions) SetReplyToAddresses(v []string) {
	o.ReplyToAddresses = v
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *ABTestOptions) GetContents() []string {
	if o == nil || IsNil(o.Contents) {
		var ret []string
		return ret
	}
	return o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ABTestOptions) GetContentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Contents) {
		return nil, false
	}
	return o.Contents, true
}

// HaveContents returns a boolean if a field has been set.
func (o *ABTestOptions) HaveContents() bool {
	if o != nil && !IsNil(o.Contents) {
		return true
	}

	return false
}

// SetContents gets a reference to the given []string and assigns it to the Contents field.
func (o *ABTestOptions) SetContents(v []string) {
	o.Contents = v
}

// GetCombinations returns the Combinations field value if set, zero value otherwise.
func (o *ABTestOptions) GetCombinations() []CombinationsInner {
	if o == nil || IsNil(o.Combinations) {
		var ret []CombinationsInner
		return ret
	}
	return o.Combinations
}

// GetCombinationsOk returns a tuple with the Combinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ABTestOptions) GetCombinationsOk() ([]CombinationsInner, bool) {
	if o == nil || IsNil(o.Combinations) {
		return nil, false
	}
	return o.Combinations, true
}

// HaveCombinations returns a boolean if a field has been set.
func (o *ABTestOptions) HaveCombinations() bool {
	if o != nil && !IsNil(o.Combinations) {
		return true
	}

	return false
}

// SetCombinations gets a reference to the given []CombinationsInner and assigns it to the Combinations field.
func (o *ABTestOptions) SetCombinations(v []CombinationsInner) {
	o.Combinations = v
}

func (o ABTestOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ABTestOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WinningCombinationId) {
		toSerialize["winning_combination_id"] = o.WinningCombinationId
	}
	if !IsNil(o.WinningCampaignId) {
		toSerialize["winning_campaign_id"] = o.WinningCampaignId
	}
	if !IsNil(o.WinnerCriteria) {
		toSerialize["winner_criteria"] = o.WinnerCriteria
	}
	if !IsNil(o.WaitTime) {
		toSerialize["wait_time"] = o.WaitTime
	}
	if !IsNil(o.TestSize) {
		toSerialize["test_size"] = o.TestSize
	}
	if !IsNil(o.SubjectLines) {
		toSerialize["subject_lines"] = o.SubjectLines
	}
	if !IsNil(o.SendTimes) {
		toSerialize["send_times"] = o.SendTimes
	}
	if !IsNil(o.FromNames) {
		toSerialize["from_names"] = o.FromNames
	}
	if !IsNil(o.ReplyToAddresses) {
		toSerialize["reply_to_addresses"] = o.ReplyToAddresses
	}
	if !IsNil(o.Contents) {
		toSerialize["contents"] = o.Contents
	}
	if !IsNil(o.Combinations) {
		toSerialize["combinations"] = o.Combinations
	}
	return toSerialize, nil
}

type NullableABTestOptions struct {
	value *ABTestOptions
	isSet bool
}

func (v NullableABTestOptions) Get() *ABTestOptions {
	return v.value
}

func (v *NullableABTestOptions) Set(val *ABTestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableABTestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableABTestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableABTestOptions(val *ABTestOptions) *NullableABTestOptions {
	return &NullableABTestOptions{value: val, isSet: true}
}

func (v NullableABTestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableABTestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


