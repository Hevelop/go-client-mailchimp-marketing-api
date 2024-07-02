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

// checks if the LandingPageReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LandingPageReport{}

// LandingPageReport A summary of an individual landing page's settings and content.
type LandingPageReport struct {
	// A string that uniquely identifies this landing page.
	Id *string `json:"id,omitempty"`
	// The name of this landing page the user will see.
	Name *string `json:"name,omitempty"`
	// The name of the landing page the user's customers will see.
	Title *string `json:"title,omitempty"`
	// The landing page url.
	Url *string `json:"url,omitempty"`
	// The time this landing page was published.
	PublishedAt *time.Time `json:"published_at,omitempty"`
	// The time this landing page was unpublished.
	UnpublishedAt *time.Time `json:"unpublished_at,omitempty"`
	// The status of the landing page.
	Status *string `json:"status,omitempty"`
	// The list id connected to this landing page.
	ListId *string `json:"list_id,omitempty"`
	// The number of visits to this landing pages.
	Visits *int32 `json:"visits,omitempty"`
	// The number of unique visits to this landing pages.
	UniqueVisits *int32 `json:"unique_visits,omitempty"`
	// The number of subscribes to this landing pages.
	Subscribes *int32 `json:"subscribes,omitempty"`
	// The number of clicks to this landing pages.
	Clicks *int32 `json:"clicks,omitempty"`
	// The percentage of people who visited your landing page and were added to your list.
	ConversionRate *float32 `json:"conversion_rate,omitempty"`
	Timeseries *LandingPageReportTimeseries `json:"timeseries,omitempty"`
	Ecommerce *LandingPageReportEcommerce `json:"ecommerce,omitempty"`
	// The ID used in the Mailchimp web application.
	WebId *int32 `json:"web_id,omitempty"`
	// List Name
	ListName *string `json:"list_name,omitempty"`
	// A list of tags associated to the landing page.
	SignupTags []Tag `json:"signup_tags,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewLandingPageReport instantiates a new LandingPageReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLandingPageReport() *LandingPageReport {
	this := LandingPageReport{}
	return &this
}

// NewLandingPageReportWithDefaults instantiates a new LandingPageReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLandingPageReportWithDefaults() *LandingPageReport {
	this := LandingPageReport{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LandingPageReport) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPageReport) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LandingPageReport) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LandingPageReport) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LandingPageReport) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPageReport) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LandingPageReport) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LandingPageReport) SetName(v string) {
	o.Name = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *LandingPageReport) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPageReport) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *LandingPageReport) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *LandingPageReport) SetTitle(v string) {
	o.Title = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *LandingPageReport) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPageReport) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *LandingPageReport) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *LandingPageReport) SetUrl(v string) {
	o.Url = &v
}

// GetPublishedAt returns the PublishedAt field value if set, zero value otherwise.
func (o *LandingPageReport) GetPublishedAt() time.Time {
	if o == nil || IsNil(o.PublishedAt) {
		var ret time.Time
		return ret
	}
	return *o.PublishedAt
}

// GetPublishedAtOk returns a tuple with the PublishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPageReport) GetPublishedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PublishedAt) {
		return nil, false
	}
	return o.PublishedAt, true
}

// HasPublishedAt returns a boolean if a field has been set.
func (o *LandingPageReport) HasPublishedAt() bool {
	if o != nil && !IsNil(o.PublishedAt) {
		return true
	}

	return false
}

// SetPublishedAt gets a reference to the given time.Time and assigns it to the PublishedAt field.
func (o *LandingPageReport) SetPublishedAt(v time.Time) {
	o.PublishedAt = &v
}

// GetUnpublishedAt returns the UnpublishedAt field value if set, zero value otherwise.
func (o *LandingPageReport) GetUnpublishedAt() time.Time {
	if o == nil || IsNil(o.UnpublishedAt) {
		var ret time.Time
		return ret
	}
	return *o.UnpublishedAt
}

// GetUnpublishedAtOk returns a tuple with the UnpublishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPageReport) GetUnpublishedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UnpublishedAt) {
		return nil, false
	}
	return o.UnpublishedAt, true
}

// HasUnpublishedAt returns a boolean if a field has been set.
func (o *LandingPageReport) HasUnpublishedAt() bool {
	if o != nil && !IsNil(o.UnpublishedAt) {
		return true
	}

	return false
}

// SetUnpublishedAt gets a reference to the given time.Time and assigns it to the UnpublishedAt field.
func (o *LandingPageReport) SetUnpublishedAt(v time.Time) {
	o.UnpublishedAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LandingPageReport) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPageReport) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LandingPageReport) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *LandingPageReport) SetStatus(v string) {
	o.Status = &v
}

// GetListId returns the ListId field value if set, zero value otherwise.
func (o *LandingPageReport) GetListId() string {
	if o == nil || IsNil(o.ListId) {
		var ret string
		return ret
	}
	return *o.ListId
}

// GetListIdOk returns a tuple with the ListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPageReport) GetListIdOk() (*string, bool) {
	if o == nil || IsNil(o.ListId) {
		return nil, false
	}
	return o.ListId, true
}

// HasListId returns a boolean if a field has been set.
func (o *LandingPageReport) HasListId() bool {
	if o != nil && !IsNil(o.ListId) {
		return true
	}

	return false
}

// SetListId gets a reference to the given string and assigns it to the ListId field.
func (o *LandingPageReport) SetListId(v string) {
	o.ListId = &v
}

// GetVisits returns the Visits field value if set, zero value otherwise.
func (o *LandingPageReport) GetVisits() int32 {
	if o == nil || IsNil(o.Visits) {
		var ret int32
		return ret
	}
	return *o.Visits
}

// GetVisitsOk returns a tuple with the Visits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPageReport) GetVisitsOk() (*int32, bool) {
	if o == nil || IsNil(o.Visits) {
		return nil, false
	}
	return o.Visits, true
}

// HasVisits returns a boolean if a field has been set.
func (o *LandingPageReport) HasVisits() bool {
	if o != nil && !IsNil(o.Visits) {
		return true
	}

	return false
}

// SetVisits gets a reference to the given int32 and assigns it to the Visits field.
func (o *LandingPageReport) SetVisits(v int32) {
	o.Visits = &v
}

// GetUniqueVisits returns the UniqueVisits field value if set, zero value otherwise.
func (o *LandingPageReport) GetUniqueVisits() int32 {
	if o == nil || IsNil(o.UniqueVisits) {
		var ret int32
		return ret
	}
	return *o.UniqueVisits
}

// GetUniqueVisitsOk returns a tuple with the UniqueVisits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPageReport) GetUniqueVisitsOk() (*int32, bool) {
	if o == nil || IsNil(o.UniqueVisits) {
		return nil, false
	}
	return o.UniqueVisits, true
}

// HasUniqueVisits returns a boolean if a field has been set.
func (o *LandingPageReport) HasUniqueVisits() bool {
	if o != nil && !IsNil(o.UniqueVisits) {
		return true
	}

	return false
}

// SetUniqueVisits gets a reference to the given int32 and assigns it to the UniqueVisits field.
func (o *LandingPageReport) SetUniqueVisits(v int32) {
	o.UniqueVisits = &v
}

// GetSubscribes returns the Subscribes field value if set, zero value otherwise.
func (o *LandingPageReport) GetSubscribes() int32 {
	if o == nil || IsNil(o.Subscribes) {
		var ret int32
		return ret
	}
	return *o.Subscribes
}

// GetSubscribesOk returns a tuple with the Subscribes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPageReport) GetSubscribesOk() (*int32, bool) {
	if o == nil || IsNil(o.Subscribes) {
		return nil, false
	}
	return o.Subscribes, true
}

// HasSubscribes returns a boolean if a field has been set.
func (o *LandingPageReport) HasSubscribes() bool {
	if o != nil && !IsNil(o.Subscribes) {
		return true
	}

	return false
}

// SetSubscribes gets a reference to the given int32 and assigns it to the Subscribes field.
func (o *LandingPageReport) SetSubscribes(v int32) {
	o.Subscribes = &v
}

// GetClicks returns the Clicks field value if set, zero value otherwise.
func (o *LandingPageReport) GetClicks() int32 {
	if o == nil || IsNil(o.Clicks) {
		var ret int32
		return ret
	}
	return *o.Clicks
}

// GetClicksOk returns a tuple with the Clicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPageReport) GetClicksOk() (*int32, bool) {
	if o == nil || IsNil(o.Clicks) {
		return nil, false
	}
	return o.Clicks, true
}

// HasClicks returns a boolean if a field has been set.
func (o *LandingPageReport) HasClicks() bool {
	if o != nil && !IsNil(o.Clicks) {
		return true
	}

	return false
}

// SetClicks gets a reference to the given int32 and assigns it to the Clicks field.
func (o *LandingPageReport) SetClicks(v int32) {
	o.Clicks = &v
}

// GetConversionRate returns the ConversionRate field value if set, zero value otherwise.
func (o *LandingPageReport) GetConversionRate() float32 {
	if o == nil || IsNil(o.ConversionRate) {
		var ret float32
		return ret
	}
	return *o.ConversionRate
}

// GetConversionRateOk returns a tuple with the ConversionRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPageReport) GetConversionRateOk() (*float32, bool) {
	if o == nil || IsNil(o.ConversionRate) {
		return nil, false
	}
	return o.ConversionRate, true
}

// HasConversionRate returns a boolean if a field has been set.
func (o *LandingPageReport) HasConversionRate() bool {
	if o != nil && !IsNil(o.ConversionRate) {
		return true
	}

	return false
}

// SetConversionRate gets a reference to the given float32 and assigns it to the ConversionRate field.
func (o *LandingPageReport) SetConversionRate(v float32) {
	o.ConversionRate = &v
}

// GetTimeseries returns the Timeseries field value if set, zero value otherwise.
func (o *LandingPageReport) GetTimeseries() LandingPageReportTimeseries {
	if o == nil || IsNil(o.Timeseries) {
		var ret LandingPageReportTimeseries
		return ret
	}
	return *o.Timeseries
}

// GetTimeseriesOk returns a tuple with the Timeseries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPageReport) GetTimeseriesOk() (*LandingPageReportTimeseries, bool) {
	if o == nil || IsNil(o.Timeseries) {
		return nil, false
	}
	return o.Timeseries, true
}

// HasTimeseries returns a boolean if a field has been set.
func (o *LandingPageReport) HasTimeseries() bool {
	if o != nil && !IsNil(o.Timeseries) {
		return true
	}

	return false
}

// SetTimeseries gets a reference to the given LandingPageReportTimeseries and assigns it to the Timeseries field.
func (o *LandingPageReport) SetTimeseries(v LandingPageReportTimeseries) {
	o.Timeseries = &v
}

// GetEcommerce returns the Ecommerce field value if set, zero value otherwise.
func (o *LandingPageReport) GetEcommerce() LandingPageReportEcommerce {
	if o == nil || IsNil(o.Ecommerce) {
		var ret LandingPageReportEcommerce
		return ret
	}
	return *o.Ecommerce
}

// GetEcommerceOk returns a tuple with the Ecommerce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPageReport) GetEcommerceOk() (*LandingPageReportEcommerce, bool) {
	if o == nil || IsNil(o.Ecommerce) {
		return nil, false
	}
	return o.Ecommerce, true
}

// HasEcommerce returns a boolean if a field has been set.
func (o *LandingPageReport) HasEcommerce() bool {
	if o != nil && !IsNil(o.Ecommerce) {
		return true
	}

	return false
}

// SetEcommerce gets a reference to the given LandingPageReportEcommerce and assigns it to the Ecommerce field.
func (o *LandingPageReport) SetEcommerce(v LandingPageReportEcommerce) {
	o.Ecommerce = &v
}

// GetWebId returns the WebId field value if set, zero value otherwise.
func (o *LandingPageReport) GetWebId() int32 {
	if o == nil || IsNil(o.WebId) {
		var ret int32
		return ret
	}
	return *o.WebId
}

// GetWebIdOk returns a tuple with the WebId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPageReport) GetWebIdOk() (*int32, bool) {
	if o == nil || IsNil(o.WebId) {
		return nil, false
	}
	return o.WebId, true
}

// HasWebId returns a boolean if a field has been set.
func (o *LandingPageReport) HasWebId() bool {
	if o != nil && !IsNil(o.WebId) {
		return true
	}

	return false
}

// SetWebId gets a reference to the given int32 and assigns it to the WebId field.
func (o *LandingPageReport) SetWebId(v int32) {
	o.WebId = &v
}

// GetListName returns the ListName field value if set, zero value otherwise.
func (o *LandingPageReport) GetListName() string {
	if o == nil || IsNil(o.ListName) {
		var ret string
		return ret
	}
	return *o.ListName
}

// GetListNameOk returns a tuple with the ListName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPageReport) GetListNameOk() (*string, bool) {
	if o == nil || IsNil(o.ListName) {
		return nil, false
	}
	return o.ListName, true
}

// HasListName returns a boolean if a field has been set.
func (o *LandingPageReport) HasListName() bool {
	if o != nil && !IsNil(o.ListName) {
		return true
	}

	return false
}

// SetListName gets a reference to the given string and assigns it to the ListName field.
func (o *LandingPageReport) SetListName(v string) {
	o.ListName = &v
}

// GetSignupTags returns the SignupTags field value if set, zero value otherwise.
func (o *LandingPageReport) GetSignupTags() []Tag {
	if o == nil || IsNil(o.SignupTags) {
		var ret []Tag
		return ret
	}
	return o.SignupTags
}

// GetSignupTagsOk returns a tuple with the SignupTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPageReport) GetSignupTagsOk() ([]Tag, bool) {
	if o == nil || IsNil(o.SignupTags) {
		return nil, false
	}
	return o.SignupTags, true
}

// HasSignupTags returns a boolean if a field has been set.
func (o *LandingPageReport) HasSignupTags() bool {
	if o != nil && !IsNil(o.SignupTags) {
		return true
	}

	return false
}

// SetSignupTags gets a reference to the given []Tag and assigns it to the SignupTags field.
func (o *LandingPageReport) SetSignupTags(v []Tag) {
	o.SignupTags = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *LandingPageReport) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPageReport) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *LandingPageReport) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *LandingPageReport) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o LandingPageReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LandingPageReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.PublishedAt) {
		toSerialize["published_at"] = o.PublishedAt
	}
	if !IsNil(o.UnpublishedAt) {
		toSerialize["unpublished_at"] = o.UnpublishedAt
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.ListId) {
		toSerialize["list_id"] = o.ListId
	}
	if !IsNil(o.Visits) {
		toSerialize["visits"] = o.Visits
	}
	if !IsNil(o.UniqueVisits) {
		toSerialize["unique_visits"] = o.UniqueVisits
	}
	if !IsNil(o.Subscribes) {
		toSerialize["subscribes"] = o.Subscribes
	}
	if !IsNil(o.Clicks) {
		toSerialize["clicks"] = o.Clicks
	}
	if !IsNil(o.ConversionRate) {
		toSerialize["conversion_rate"] = o.ConversionRate
	}
	if !IsNil(o.Timeseries) {
		toSerialize["timeseries"] = o.Timeseries
	}
	if !IsNil(o.Ecommerce) {
		toSerialize["ecommerce"] = o.Ecommerce
	}
	if !IsNil(o.WebId) {
		toSerialize["web_id"] = o.WebId
	}
	if !IsNil(o.ListName) {
		toSerialize["list_name"] = o.ListName
	}
	if !IsNil(o.SignupTags) {
		toSerialize["signup_tags"] = o.SignupTags
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableLandingPageReport struct {
	value *LandingPageReport
	isSet bool
}

func (v NullableLandingPageReport) Get() *LandingPageReport {
	return v.value
}

func (v *NullableLandingPageReport) Set(val *LandingPageReport) {
	v.value = val
	v.isSet = true
}

func (v NullableLandingPageReport) IsSet() bool {
	return v.isSet
}

func (v *NullableLandingPageReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLandingPageReport(val *LandingPageReport) *NullableLandingPageReport {
	return &NullableLandingPageReport{value: val, isSet: true}
}

func (v NullableLandingPageReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLandingPageReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


