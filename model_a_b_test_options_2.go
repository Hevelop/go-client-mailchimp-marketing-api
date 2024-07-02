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

// checks if the ABTestOptions2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ABTestOptions2{}

// ABTestOptions2 The settings specific to A/B test campaigns.
type ABTestOptions2 struct {
	// ID for the winning combination.
	WinningCombinationId *string `json:"winning_combination_id,omitempty"`
	// ID of the campaign that was sent to the remaining recipients based on the winning combination.
	WinningCampaignId *string `json:"winning_campaign_id,omitempty"`
	// The combination that performs the best. This may be determined automatically by click rate, open rate, or total revenue -- or you may choose manually based on the reporting data you find the most valuable. For Multivariate Campaigns testing send_time, winner_criteria is ignored. For Multivariate Campaigns with 'manual' as the winner_criteria, the winner must be chosen in the Mailchimp web application.
	WinnerCriteria string `json:"winner_criteria"`
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

type _ABTestOptions2 ABTestOptions2

// NewABTestOptions2 instantiates a new ABTestOptions2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewABTestOptions2(winnerCriteria string) *ABTestOptions2 {
	this := ABTestOptions2{}
	this.WinnerCriteria = winnerCriteria
	return &this
}

// NewABTestOptions2WithDefaults instantiates a new ABTestOptions2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewABTestOptions2WithDefaults() *ABTestOptions2 {
	this := ABTestOptions2{}
	return &this
}

// GetWinningCombinationId returns the WinningCombinationId field value if set, zero value otherwise.
func (o *ABTestOptions2) GetWinningCombinationId() string {
	if o == nil || IsNil(o.WinningCombinationId) {
		var ret string
		return ret
	}
	return *o.WinningCombinationId
}

// GetWinningCombinationIdOk returns a tuple with the WinningCombinationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ABTestOptions2) GetWinningCombinationIdOk() (*string, bool) {
	if o == nil || IsNil(o.WinningCombinationId) {
		return nil, false
	}
	return o.WinningCombinationId, true
}

// HasWinningCombinationId returns a boolean if a field has been set.
func (o *ABTestOptions2) HasWinningCombinationId() bool {
	if o != nil && !IsNil(o.WinningCombinationId) {
		return true
	}

	return false
}

// SetWinningCombinationId gets a reference to the given string and assigns it to the WinningCombinationId field.
func (o *ABTestOptions2) SetWinningCombinationId(v string) {
	o.WinningCombinationId = &v
}

// GetWinningCampaignId returns the WinningCampaignId field value if set, zero value otherwise.
func (o *ABTestOptions2) GetWinningCampaignId() string {
	if o == nil || IsNil(o.WinningCampaignId) {
		var ret string
		return ret
	}
	return *o.WinningCampaignId
}

// GetWinningCampaignIdOk returns a tuple with the WinningCampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ABTestOptions2) GetWinningCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.WinningCampaignId) {
		return nil, false
	}
	return o.WinningCampaignId, true
}

// HasWinningCampaignId returns a boolean if a field has been set.
func (o *ABTestOptions2) HasWinningCampaignId() bool {
	if o != nil && !IsNil(o.WinningCampaignId) {
		return true
	}

	return false
}

// SetWinningCampaignId gets a reference to the given string and assigns it to the WinningCampaignId field.
func (o *ABTestOptions2) SetWinningCampaignId(v string) {
	o.WinningCampaignId = &v
}

// GetWinnerCriteria returns the WinnerCriteria field value
func (o *ABTestOptions2) GetWinnerCriteria() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WinnerCriteria
}

// GetWinnerCriteriaOk returns a tuple with the WinnerCriteria field value
// and a boolean to check if the value has been set.
func (o *ABTestOptions2) GetWinnerCriteriaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WinnerCriteria, true
}

// SetWinnerCriteria sets field value
func (o *ABTestOptions2) SetWinnerCriteria(v string) {
	o.WinnerCriteria = v
}

// GetWaitTime returns the WaitTime field value if set, zero value otherwise.
func (o *ABTestOptions2) GetWaitTime() int32 {
	if o == nil || IsNil(o.WaitTime) {
		var ret int32
		return ret
	}
	return *o.WaitTime
}

// GetWaitTimeOk returns a tuple with the WaitTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ABTestOptions2) GetWaitTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.WaitTime) {
		return nil, false
	}
	return o.WaitTime, true
}

// HasWaitTime returns a boolean if a field has been set.
func (o *ABTestOptions2) HasWaitTime() bool {
	if o != nil && !IsNil(o.WaitTime) {
		return true
	}

	return false
}

// SetWaitTime gets a reference to the given int32 and assigns it to the WaitTime field.
func (o *ABTestOptions2) SetWaitTime(v int32) {
	o.WaitTime = &v
}

// GetTestSize returns the TestSize field value if set, zero value otherwise.
func (o *ABTestOptions2) GetTestSize() int32 {
	if o == nil || IsNil(o.TestSize) {
		var ret int32
		return ret
	}
	return *o.TestSize
}

// GetTestSizeOk returns a tuple with the TestSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ABTestOptions2) GetTestSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.TestSize) {
		return nil, false
	}
	return o.TestSize, true
}

// HasTestSize returns a boolean if a field has been set.
func (o *ABTestOptions2) HasTestSize() bool {
	if o != nil && !IsNil(o.TestSize) {
		return true
	}

	return false
}

// SetTestSize gets a reference to the given int32 and assigns it to the TestSize field.
func (o *ABTestOptions2) SetTestSize(v int32) {
	o.TestSize = &v
}

// GetSubjectLines returns the SubjectLines field value if set, zero value otherwise.
func (o *ABTestOptions2) GetSubjectLines() []string {
	if o == nil || IsNil(o.SubjectLines) {
		var ret []string
		return ret
	}
	return o.SubjectLines
}

// GetSubjectLinesOk returns a tuple with the SubjectLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ABTestOptions2) GetSubjectLinesOk() ([]string, bool) {
	if o == nil || IsNil(o.SubjectLines) {
		return nil, false
	}
	return o.SubjectLines, true
}

// HasSubjectLines returns a boolean if a field has been set.
func (o *ABTestOptions2) HasSubjectLines() bool {
	if o != nil && !IsNil(o.SubjectLines) {
		return true
	}

	return false
}

// SetSubjectLines gets a reference to the given []string and assigns it to the SubjectLines field.
func (o *ABTestOptions2) SetSubjectLines(v []string) {
	o.SubjectLines = v
}

// GetSendTimes returns the SendTimes field value if set, zero value otherwise.
func (o *ABTestOptions2) GetSendTimes() []time.Time {
	if o == nil || IsNil(o.SendTimes) {
		var ret []time.Time
		return ret
	}
	return o.SendTimes
}

// GetSendTimesOk returns a tuple with the SendTimes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ABTestOptions2) GetSendTimesOk() ([]time.Time, bool) {
	if o == nil || IsNil(o.SendTimes) {
		return nil, false
	}
	return o.SendTimes, true
}

// HasSendTimes returns a boolean if a field has been set.
func (o *ABTestOptions2) HasSendTimes() bool {
	if o != nil && !IsNil(o.SendTimes) {
		return true
	}

	return false
}

// SetSendTimes gets a reference to the given []time.Time and assigns it to the SendTimes field.
func (o *ABTestOptions2) SetSendTimes(v []time.Time) {
	o.SendTimes = v
}

// GetFromNames returns the FromNames field value if set, zero value otherwise.
func (o *ABTestOptions2) GetFromNames() []string {
	if o == nil || IsNil(o.FromNames) {
		var ret []string
		return ret
	}
	return o.FromNames
}

// GetFromNamesOk returns a tuple with the FromNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ABTestOptions2) GetFromNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.FromNames) {
		return nil, false
	}
	return o.FromNames, true
}

// HasFromNames returns a boolean if a field has been set.
func (o *ABTestOptions2) HasFromNames() bool {
	if o != nil && !IsNil(o.FromNames) {
		return true
	}

	return false
}

// SetFromNames gets a reference to the given []string and assigns it to the FromNames field.
func (o *ABTestOptions2) SetFromNames(v []string) {
	o.FromNames = v
}

// GetReplyToAddresses returns the ReplyToAddresses field value if set, zero value otherwise.
func (o *ABTestOptions2) GetReplyToAddresses() []string {
	if o == nil || IsNil(o.ReplyToAddresses) {
		var ret []string
		return ret
	}
	return o.ReplyToAddresses
}

// GetReplyToAddressesOk returns a tuple with the ReplyToAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ABTestOptions2) GetReplyToAddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.ReplyToAddresses) {
		return nil, false
	}
	return o.ReplyToAddresses, true
}

// HasReplyToAddresses returns a boolean if a field has been set.
func (o *ABTestOptions2) HasReplyToAddresses() bool {
	if o != nil && !IsNil(o.ReplyToAddresses) {
		return true
	}

	return false
}

// SetReplyToAddresses gets a reference to the given []string and assigns it to the ReplyToAddresses field.
func (o *ABTestOptions2) SetReplyToAddresses(v []string) {
	o.ReplyToAddresses = v
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *ABTestOptions2) GetContents() []string {
	if o == nil || IsNil(o.Contents) {
		var ret []string
		return ret
	}
	return o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ABTestOptions2) GetContentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Contents) {
		return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *ABTestOptions2) HasContents() bool {
	if o != nil && !IsNil(o.Contents) {
		return true
	}

	return false
}

// SetContents gets a reference to the given []string and assigns it to the Contents field.
func (o *ABTestOptions2) SetContents(v []string) {
	o.Contents = v
}

// GetCombinations returns the Combinations field value if set, zero value otherwise.
func (o *ABTestOptions2) GetCombinations() []CombinationsInner {
	if o == nil || IsNil(o.Combinations) {
		var ret []CombinationsInner
		return ret
	}
	return o.Combinations
}

// GetCombinationsOk returns a tuple with the Combinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ABTestOptions2) GetCombinationsOk() ([]CombinationsInner, bool) {
	if o == nil || IsNil(o.Combinations) {
		return nil, false
	}
	return o.Combinations, true
}

// HasCombinations returns a boolean if a field has been set.
func (o *ABTestOptions2) HasCombinations() bool {
	if o != nil && !IsNil(o.Combinations) {
		return true
	}

	return false
}

// SetCombinations gets a reference to the given []CombinationsInner and assigns it to the Combinations field.
func (o *ABTestOptions2) SetCombinations(v []CombinationsInner) {
	o.Combinations = v
}

func (o ABTestOptions2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ABTestOptions2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WinningCombinationId) {
		toSerialize["winning_combination_id"] = o.WinningCombinationId
	}
	if !IsNil(o.WinningCampaignId) {
		toSerialize["winning_campaign_id"] = o.WinningCampaignId
	}
	toSerialize["winner_criteria"] = o.WinnerCriteria
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

func (o *ABTestOptions2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"winner_criteria",
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

	varABTestOptions2 := _ABTestOptions2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varABTestOptions2)

	if err != nil {
		return err
	}

	*o = ABTestOptions2(varABTestOptions2)

	return err
}

type NullableABTestOptions2 struct {
	value *ABTestOptions2
	isSet bool
}

func (v NullableABTestOptions2) Get() *ABTestOptions2 {
	return v.value
}

func (v *NullableABTestOptions2) Set(val *ABTestOptions2) {
	v.value = val
	v.isSet = true
}

func (v NullableABTestOptions2) IsSet() bool {
	return v.isSet
}

func (v *NullableABTestOptions2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableABTestOptions2(val *ABTestOptions2) *NullableABTestOptions2 {
	return &NullableABTestOptions2{value: val, isSet: true}
}

func (v NullableABTestOptions2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableABTestOptions2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


