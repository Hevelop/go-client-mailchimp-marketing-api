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

// checks if the TwitterStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TwitterStats{}

// TwitterStats A summary of Twitter activity for a campaign.
type TwitterStats struct {
	// The number of tweets including a link to the campaign.
	Tweets *int32 `json:"tweets,omitempty"`
	// The day and time of the first recorded tweet with a link to the campaign.
	FirstTweet *string `json:"first_tweet,omitempty"`
	// The day and time of the last recorded tweet with a link to the campaign.
	LastTweet *string `json:"last_tweet,omitempty"`
	// The number of retweets that include a link to the campaign.
	Retweets *int32 `json:"retweets,omitempty"`
	// A summary of tweets that include a link to the campaign.
	Statuses []TwitterStatus `json:"statuses,omitempty"`
}

// NewTwitterStats instantiates a new TwitterStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTwitterStats() *TwitterStats {
	this := TwitterStats{}
	return &this
}

// NewTwitterStatsWithDefaults instantiates a new TwitterStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTwitterStatsWithDefaults() *TwitterStats {
	this := TwitterStats{}
	return &this
}

// GetTweets returns the Tweets field value if set, zero value otherwise.
func (o *TwitterStats) GetTweets() int32 {
	if o == nil || IsNil(o.Tweets) {
		var ret int32
		return ret
	}
	return *o.Tweets
}

// GetTweetsOk returns a tuple with the Tweets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwitterStats) GetTweetsOk() (*int32, bool) {
	if o == nil || IsNil(o.Tweets) {
		return nil, false
	}
	return o.Tweets, true
}

// HaveTweets returns a boolean if a field has been set.
func (o *TwitterStats) HaveTweets() bool {
	if o != nil && !IsNil(o.Tweets) {
		return true
	}

	return false
}

// SetTweets gets a reference to the given int32 and assigns it to the Tweets field.
func (o *TwitterStats) SetTweets(v int32) {
	o.Tweets = &v
}

// GetFirstTweet returns the FirstTweet field value if set, zero value otherwise.
func (o *TwitterStats) GetFirstTweet() string {
	if o == nil || IsNil(o.FirstTweet) {
		var ret string
		return ret
	}
	return *o.FirstTweet
}

// GetFirstTweetOk returns a tuple with the FirstTweet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwitterStats) GetFirstTweetOk() (*string, bool) {
	if o == nil || IsNil(o.FirstTweet) {
		return nil, false
	}
	return o.FirstTweet, true
}

// HaveFirstTweet returns a boolean if a field has been set.
func (o *TwitterStats) HaveFirstTweet() bool {
	if o != nil && !IsNil(o.FirstTweet) {
		return true
	}

	return false
}

// SetFirstTweet gets a reference to the given string and assigns it to the FirstTweet field.
func (o *TwitterStats) SetFirstTweet(v string) {
	o.FirstTweet = &v
}

// GetLastTweet returns the LastTweet field value if set, zero value otherwise.
func (o *TwitterStats) GetLastTweet() string {
	if o == nil || IsNil(o.LastTweet) {
		var ret string
		return ret
	}
	return *o.LastTweet
}

// GetLastTweetOk returns a tuple with the LastTweet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwitterStats) GetLastTweetOk() (*string, bool) {
	if o == nil || IsNil(o.LastTweet) {
		return nil, false
	}
	return o.LastTweet, true
}

// HaveLastTweet returns a boolean if a field has been set.
func (o *TwitterStats) HaveLastTweet() bool {
	if o != nil && !IsNil(o.LastTweet) {
		return true
	}

	return false
}

// SetLastTweet gets a reference to the given string and assigns it to the LastTweet field.
func (o *TwitterStats) SetLastTweet(v string) {
	o.LastTweet = &v
}

// GetRetweets returns the Retweets field value if set, zero value otherwise.
func (o *TwitterStats) GetRetweets() int32 {
	if o == nil || IsNil(o.Retweets) {
		var ret int32
		return ret
	}
	return *o.Retweets
}

// GetRetweetsOk returns a tuple with the Retweets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwitterStats) GetRetweetsOk() (*int32, bool) {
	if o == nil || IsNil(o.Retweets) {
		return nil, false
	}
	return o.Retweets, true
}

// HaveRetweets returns a boolean if a field has been set.
func (o *TwitterStats) HaveRetweets() bool {
	if o != nil && !IsNil(o.Retweets) {
		return true
	}

	return false
}

// SetRetweets gets a reference to the given int32 and assigns it to the Retweets field.
func (o *TwitterStats) SetRetweets(v int32) {
	o.Retweets = &v
}

// GetStatuses returns the Statuses field value if set, zero value otherwise.
func (o *TwitterStats) GetStatuses() []TwitterStatus {
	if o == nil || IsNil(o.Statuses) {
		var ret []TwitterStatus
		return ret
	}
	return o.Statuses
}

// GetStatusesOk returns a tuple with the Statuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwitterStats) GetStatusesOk() ([]TwitterStatus, bool) {
	if o == nil || IsNil(o.Statuses) {
		return nil, false
	}
	return o.Statuses, true
}

// HaveStatuses returns a boolean if a field has been set.
func (o *TwitterStats) HaveStatuses() bool {
	if o != nil && !IsNil(o.Statuses) {
		return true
	}

	return false
}

// SetStatuses gets a reference to the given []TwitterStatus and assigns it to the Statuses field.
func (o *TwitterStats) SetStatuses(v []TwitterStatus) {
	o.Statuses = v
}

func (o TwitterStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TwitterStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tweets) {
		toSerialize["tweets"] = o.Tweets
	}
	if !IsNil(o.FirstTweet) {
		toSerialize["first_tweet"] = o.FirstTweet
	}
	if !IsNil(o.LastTweet) {
		toSerialize["last_tweet"] = o.LastTweet
	}
	if !IsNil(o.Retweets) {
		toSerialize["retweets"] = o.Retweets
	}
	if !IsNil(o.Statuses) {
		toSerialize["statuses"] = o.Statuses
	}
	return toSerialize, nil
}

type NullableTwitterStats struct {
	value *TwitterStats
	isSet bool
}

func (v NullableTwitterStats) Get() *TwitterStats {
	return v.value
}

func (v *NullableTwitterStats) Set(val *TwitterStats) {
	v.value = val
	v.isSet = true
}

func (v NullableTwitterStats) IsSet() bool {
	return v.isSet
}

func (v *NullableTwitterStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTwitterStats(val *TwitterStats) *NullableTwitterStats {
	return &NullableTwitterStats{value: val, isSet: true}
}

func (v NullableTwitterStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTwitterStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


