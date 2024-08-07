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

// checks if the CampaignContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignContent{}

// CampaignContent The HTML and plain-text content for a campaign.
type CampaignContent struct {
	// Content options for multivariate campaigns.
	VariateContents []VariateContentsInner `json:"variate_contents,omitempty"`
	// The plain-text portion of the campaign. If left unspecified, we'll generate this automatically.
	PlainText *string `json:"plain_text,omitempty"`
	// The raw HTML for the campaign.
	Html *string `json:"html,omitempty"`
	// The Archive HTML for the campaign.
	ArchiveHtml *string `json:"archive_html,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewCampaignContent instantiates a new CampaignContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignContent() *CampaignContent {
	this := CampaignContent{}
	return &this
}

// NewCampaignContentWithDefaults instantiates a new CampaignContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignContentWithDefaults() *CampaignContent {
	this := CampaignContent{}
	return &this
}

// GetVariateContents returns the VariateContents field value if set, zero value otherwise.
func (o *CampaignContent) GetVariateContents() []VariateContentsInner {
	if o == nil || IsNil(o.VariateContents) {
		var ret []VariateContentsInner
		return ret
	}
	return o.VariateContents
}

// GetVariateContentsOk returns a tuple with the VariateContents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignContent) GetVariateContentsOk() ([]VariateContentsInner, bool) {
	if o == nil || IsNil(o.VariateContents) {
		return nil, false
	}
	return o.VariateContents, true
}

// HasVariateContents returns a boolean if a field has been set.
func (o *CampaignContent) HasVariateContents() bool {
	if o != nil && !IsNil(o.VariateContents) {
		return true
	}

	return false
}

// SetVariateContents gets a reference to the given []VariateContentsInner and assigns it to the VariateContents field.
func (o *CampaignContent) SetVariateContents(v []VariateContentsInner) {
	o.VariateContents = v
}

// GetPlainText returns the PlainText field value if set, zero value otherwise.
func (o *CampaignContent) GetPlainText() string {
	if o == nil || IsNil(o.PlainText) {
		var ret string
		return ret
	}
	return *o.PlainText
}

// GetPlainTextOk returns a tuple with the PlainText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignContent) GetPlainTextOk() (*string, bool) {
	if o == nil || IsNil(o.PlainText) {
		return nil, false
	}
	return o.PlainText, true
}

// HasPlainText returns a boolean if a field has been set.
func (o *CampaignContent) HasPlainText() bool {
	if o != nil && !IsNil(o.PlainText) {
		return true
	}

	return false
}

// SetPlainText gets a reference to the given string and assigns it to the PlainText field.
func (o *CampaignContent) SetPlainText(v string) {
	o.PlainText = &v
}

// GetHtml returns the Html field value if set, zero value otherwise.
func (o *CampaignContent) GetHtml() string {
	if o == nil || IsNil(o.Html) {
		var ret string
		return ret
	}
	return *o.Html
}

// GetHtmlOk returns a tuple with the Html field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignContent) GetHtmlOk() (*string, bool) {
	if o == nil || IsNil(o.Html) {
		return nil, false
	}
	return o.Html, true
}

// HasHtml returns a boolean if a field has been set.
func (o *CampaignContent) HasHtml() bool {
	if o != nil && !IsNil(o.Html) {
		return true
	}

	return false
}

// SetHtml gets a reference to the given string and assigns it to the Html field.
func (o *CampaignContent) SetHtml(v string) {
	o.Html = &v
}

// GetArchiveHtml returns the ArchiveHtml field value if set, zero value otherwise.
func (o *CampaignContent) GetArchiveHtml() string {
	if o == nil || IsNil(o.ArchiveHtml) {
		var ret string
		return ret
	}
	return *o.ArchiveHtml
}

// GetArchiveHtmlOk returns a tuple with the ArchiveHtml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignContent) GetArchiveHtmlOk() (*string, bool) {
	if o == nil || IsNil(o.ArchiveHtml) {
		return nil, false
	}
	return o.ArchiveHtml, true
}

// HasArchiveHtml returns a boolean if a field has been set.
func (o *CampaignContent) HasArchiveHtml() bool {
	if o != nil && !IsNil(o.ArchiveHtml) {
		return true
	}

	return false
}

// SetArchiveHtml gets a reference to the given string and assigns it to the ArchiveHtml field.
func (o *CampaignContent) SetArchiveHtml(v string) {
	o.ArchiveHtml = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CampaignContent) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignContent) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CampaignContent) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *CampaignContent) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o CampaignContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VariateContents) {
		toSerialize["variate_contents"] = o.VariateContents
	}
	if !IsNil(o.PlainText) {
		toSerialize["plain_text"] = o.PlainText
	}
	if !IsNil(o.Html) {
		toSerialize["html"] = o.Html
	}
	if !IsNil(o.ArchiveHtml) {
		toSerialize["archive_html"] = o.ArchiveHtml
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableCampaignContent struct {
	value *CampaignContent
	isSet bool
}

func (v NullableCampaignContent) Get() *CampaignContent {
	return v.value
}

func (v *NullableCampaignContent) Set(val *CampaignContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignContent(val *CampaignContent) *NullableCampaignContent {
	return &NullableCampaignContent{value: val, isSet: true}
}

func (v NullableCampaignContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


