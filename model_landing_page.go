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

// checks if the LandingPage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LandingPage{}

// LandingPage A summary of an individual landing page's settings and content.
type LandingPage struct {
	// A string that uniquely identifies this landing page.
	Id *string `json:"id,omitempty"`
	// The name of this landing page.
	Name *string `json:"name,omitempty"`
	// The title of this landing page seen in the browser's title bar.
	Title *string `json:"title,omitempty"`
	// The description of this landing page.
	Description *string `json:"description,omitempty"`
	// The template_id of this landing page.
	TemplateId *int32 `json:"template_id,omitempty"`
	// The status of this landing page.
	Status *string `json:"status,omitempty"`
	// The list's ID associated with this landing page.
	ListId *string `json:"list_id,omitempty"`
	// The ID of the store associated with this landing page.
	StoreId *string `json:"store_id,omitempty"`
	// The ID used in the Mailchimp web application.
	WebId *int32 `json:"web_id,omitempty"`
	// Created by mobile or web
	CreatedBySource *string `json:"created_by_source,omitempty"`
	// The url of the published landing page.
	Url *string `json:"url,omitempty"`
	// The time this landing page was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The time this landing page was published.
	PublishedAt *time.Time `json:"published_at,omitempty"`
	// The time this landing page was unpublished.
	UnpublishedAt *time.Time `json:"unpublished_at,omitempty"`
	// The time this landing page was updated at.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Tracking *TrackingSettings `json:"tracking,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewLandingPage instantiates a new LandingPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLandingPage() *LandingPage {
	this := LandingPage{}
	return &this
}

// NewLandingPageWithDefaults instantiates a new LandingPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLandingPageWithDefaults() *LandingPage {
	this := LandingPage{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LandingPage) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPage) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LandingPage) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LandingPage) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LandingPage) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPage) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LandingPage) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LandingPage) SetName(v string) {
	o.Name = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *LandingPage) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPage) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *LandingPage) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *LandingPage) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LandingPage) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPage) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LandingPage) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LandingPage) SetDescription(v string) {
	o.Description = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *LandingPage) GetTemplateId() int32 {
	if o == nil || IsNil(o.TemplateId) {
		var ret int32
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPage) GetTemplateIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *LandingPage) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given int32 and assigns it to the TemplateId field.
func (o *LandingPage) SetTemplateId(v int32) {
	o.TemplateId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LandingPage) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPage) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LandingPage) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *LandingPage) SetStatus(v string) {
	o.Status = &v
}

// GetListId returns the ListId field value if set, zero value otherwise.
func (o *LandingPage) GetListId() string {
	if o == nil || IsNil(o.ListId) {
		var ret string
		return ret
	}
	return *o.ListId
}

// GetListIdOk returns a tuple with the ListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPage) GetListIdOk() (*string, bool) {
	if o == nil || IsNil(o.ListId) {
		return nil, false
	}
	return o.ListId, true
}

// HasListId returns a boolean if a field has been set.
func (o *LandingPage) HasListId() bool {
	if o != nil && !IsNil(o.ListId) {
		return true
	}

	return false
}

// SetListId gets a reference to the given string and assigns it to the ListId field.
func (o *LandingPage) SetListId(v string) {
	o.ListId = &v
}

// GetStoreId returns the StoreId field value if set, zero value otherwise.
func (o *LandingPage) GetStoreId() string {
	if o == nil || IsNil(o.StoreId) {
		var ret string
		return ret
	}
	return *o.StoreId
}

// GetStoreIdOk returns a tuple with the StoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPage) GetStoreIdOk() (*string, bool) {
	if o == nil || IsNil(o.StoreId) {
		return nil, false
	}
	return o.StoreId, true
}

// HasStoreId returns a boolean if a field has been set.
func (o *LandingPage) HasStoreId() bool {
	if o != nil && !IsNil(o.StoreId) {
		return true
	}

	return false
}

// SetStoreId gets a reference to the given string and assigns it to the StoreId field.
func (o *LandingPage) SetStoreId(v string) {
	o.StoreId = &v
}

// GetWebId returns the WebId field value if set, zero value otherwise.
func (o *LandingPage) GetWebId() int32 {
	if o == nil || IsNil(o.WebId) {
		var ret int32
		return ret
	}
	return *o.WebId
}

// GetWebIdOk returns a tuple with the WebId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPage) GetWebIdOk() (*int32, bool) {
	if o == nil || IsNil(o.WebId) {
		return nil, false
	}
	return o.WebId, true
}

// HasWebId returns a boolean if a field has been set.
func (o *LandingPage) HasWebId() bool {
	if o != nil && !IsNil(o.WebId) {
		return true
	}

	return false
}

// SetWebId gets a reference to the given int32 and assigns it to the WebId field.
func (o *LandingPage) SetWebId(v int32) {
	o.WebId = &v
}

// GetCreatedBySource returns the CreatedBySource field value if set, zero value otherwise.
func (o *LandingPage) GetCreatedBySource() string {
	if o == nil || IsNil(o.CreatedBySource) {
		var ret string
		return ret
	}
	return *o.CreatedBySource
}

// GetCreatedBySourceOk returns a tuple with the CreatedBySource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPage) GetCreatedBySourceOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBySource) {
		return nil, false
	}
	return o.CreatedBySource, true
}

// HasCreatedBySource returns a boolean if a field has been set.
func (o *LandingPage) HasCreatedBySource() bool {
	if o != nil && !IsNil(o.CreatedBySource) {
		return true
	}

	return false
}

// SetCreatedBySource gets a reference to the given string and assigns it to the CreatedBySource field.
func (o *LandingPage) SetCreatedBySource(v string) {
	o.CreatedBySource = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *LandingPage) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPage) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *LandingPage) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *LandingPage) SetUrl(v string) {
	o.Url = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *LandingPage) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPage) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *LandingPage) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *LandingPage) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetPublishedAt returns the PublishedAt field value if set, zero value otherwise.
func (o *LandingPage) GetPublishedAt() time.Time {
	if o == nil || IsNil(o.PublishedAt) {
		var ret time.Time
		return ret
	}
	return *o.PublishedAt
}

// GetPublishedAtOk returns a tuple with the PublishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPage) GetPublishedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PublishedAt) {
		return nil, false
	}
	return o.PublishedAt, true
}

// HasPublishedAt returns a boolean if a field has been set.
func (o *LandingPage) HasPublishedAt() bool {
	if o != nil && !IsNil(o.PublishedAt) {
		return true
	}

	return false
}

// SetPublishedAt gets a reference to the given time.Time and assigns it to the PublishedAt field.
func (o *LandingPage) SetPublishedAt(v time.Time) {
	o.PublishedAt = &v
}

// GetUnpublishedAt returns the UnpublishedAt field value if set, zero value otherwise.
func (o *LandingPage) GetUnpublishedAt() time.Time {
	if o == nil || IsNil(o.UnpublishedAt) {
		var ret time.Time
		return ret
	}
	return *o.UnpublishedAt
}

// GetUnpublishedAtOk returns a tuple with the UnpublishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPage) GetUnpublishedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UnpublishedAt) {
		return nil, false
	}
	return o.UnpublishedAt, true
}

// HasUnpublishedAt returns a boolean if a field has been set.
func (o *LandingPage) HasUnpublishedAt() bool {
	if o != nil && !IsNil(o.UnpublishedAt) {
		return true
	}

	return false
}

// SetUnpublishedAt gets a reference to the given time.Time and assigns it to the UnpublishedAt field.
func (o *LandingPage) SetUnpublishedAt(v time.Time) {
	o.UnpublishedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *LandingPage) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPage) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *LandingPage) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *LandingPage) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetTracking returns the Tracking field value if set, zero value otherwise.
func (o *LandingPage) GetTracking() TrackingSettings {
	if o == nil || IsNil(o.Tracking) {
		var ret TrackingSettings
		return ret
	}
	return *o.Tracking
}

// GetTrackingOk returns a tuple with the Tracking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPage) GetTrackingOk() (*TrackingSettings, bool) {
	if o == nil || IsNil(o.Tracking) {
		return nil, false
	}
	return o.Tracking, true
}

// HasTracking returns a boolean if a field has been set.
func (o *LandingPage) HasTracking() bool {
	if o != nil && !IsNil(o.Tracking) {
		return true
	}

	return false
}

// SetTracking gets a reference to the given TrackingSettings and assigns it to the Tracking field.
func (o *LandingPage) SetTracking(v TrackingSettings) {
	o.Tracking = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *LandingPage) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LandingPage) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *LandingPage) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *LandingPage) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o LandingPage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LandingPage) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.TemplateId) {
		toSerialize["template_id"] = o.TemplateId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.ListId) {
		toSerialize["list_id"] = o.ListId
	}
	if !IsNil(o.StoreId) {
		toSerialize["store_id"] = o.StoreId
	}
	if !IsNil(o.WebId) {
		toSerialize["web_id"] = o.WebId
	}
	if !IsNil(o.CreatedBySource) {
		toSerialize["created_by_source"] = o.CreatedBySource
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.PublishedAt) {
		toSerialize["published_at"] = o.PublishedAt
	}
	if !IsNil(o.UnpublishedAt) {
		toSerialize["unpublished_at"] = o.UnpublishedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Tracking) {
		toSerialize["tracking"] = o.Tracking
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableLandingPage struct {
	value *LandingPage
	isSet bool
}

func (v NullableLandingPage) Get() *LandingPage {
	return v.value
}

func (v *NullableLandingPage) Set(val *LandingPage) {
	v.value = val
	v.isSet = true
}

func (v NullableLandingPage) IsSet() bool {
	return v.isSet
}

func (v *NullableLandingPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLandingPage(val *LandingPage) *NullableLandingPage {
	return &NullableLandingPage{value: val, isSet: true}
}

func (v NullableLandingPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLandingPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


