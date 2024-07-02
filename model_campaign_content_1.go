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

// checks if the CampaignContent1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignContent1{}

// CampaignContent1 The HTML and plain-text content for a campaign
type CampaignContent1 struct {
	// The plain-text portion of the campaign. If left unspecified, we'll generate this automatically.
	PlainText *string `json:"plain_text,omitempty"`
	// The raw HTML for the campaign.
	Html *string `json:"html,omitempty"`
	// When importing a campaign, the URL where the HTML lives.
	Url *string `json:"url,omitempty"`
	Template *TemplateContent `json:"template,omitempty"`
	Archive *UploadArchive `json:"archive,omitempty"`
	// Content options for [Multivariate Campaigns](https://mailchimp.com/help/about-multivariate-campaigns/). Each content option must provide HTML content and may optionally provide plain text. For campaigns not testing content, only one object should be provided.
	VariateContents []VariateContentsInner1 `json:"variate_contents,omitempty"`
}

// NewCampaignContent1 instantiates a new CampaignContent1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignContent1() *CampaignContent1 {
	this := CampaignContent1{}
	return &this
}

// NewCampaignContent1WithDefaults instantiates a new CampaignContent1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignContent1WithDefaults() *CampaignContent1 {
	this := CampaignContent1{}
	return &this
}

// GetPlainText returns the PlainText field value if set, zero value otherwise.
func (o *CampaignContent1) GetPlainText() string {
	if o == nil || IsNil(o.PlainText) {
		var ret string
		return ret
	}
	return *o.PlainText
}

// GetPlainTextOk returns a tuple with the PlainText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignContent1) GetPlainTextOk() (*string, bool) {
	if o == nil || IsNil(o.PlainText) {
		return nil, false
	}
	return o.PlainText, true
}

// HasPlainText returns a boolean if a field has been set.
func (o *CampaignContent1) HasPlainText() bool {
	if o != nil && !IsNil(o.PlainText) {
		return true
	}

	return false
}

// SetPlainText gets a reference to the given string and assigns it to the PlainText field.
func (o *CampaignContent1) SetPlainText(v string) {
	o.PlainText = &v
}

// GetHtml returns the Html field value if set, zero value otherwise.
func (o *CampaignContent1) GetHtml() string {
	if o == nil || IsNil(o.Html) {
		var ret string
		return ret
	}
	return *o.Html
}

// GetHtmlOk returns a tuple with the Html field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignContent1) GetHtmlOk() (*string, bool) {
	if o == nil || IsNil(o.Html) {
		return nil, false
	}
	return o.Html, true
}

// HasHtml returns a boolean if a field has been set.
func (o *CampaignContent1) HasHtml() bool {
	if o != nil && !IsNil(o.Html) {
		return true
	}

	return false
}

// SetHtml gets a reference to the given string and assigns it to the Html field.
func (o *CampaignContent1) SetHtml(v string) {
	o.Html = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CampaignContent1) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignContent1) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CampaignContent1) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CampaignContent1) SetUrl(v string) {
	o.Url = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *CampaignContent1) GetTemplate() TemplateContent {
	if o == nil || IsNil(o.Template) {
		var ret TemplateContent
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignContent1) GetTemplateOk() (*TemplateContent, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *CampaignContent1) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given TemplateContent and assigns it to the Template field.
func (o *CampaignContent1) SetTemplate(v TemplateContent) {
	o.Template = &v
}

// GetArchive returns the Archive field value if set, zero value otherwise.
func (o *CampaignContent1) GetArchive() UploadArchive {
	if o == nil || IsNil(o.Archive) {
		var ret UploadArchive
		return ret
	}
	return *o.Archive
}

// GetArchiveOk returns a tuple with the Archive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignContent1) GetArchiveOk() (*UploadArchive, bool) {
	if o == nil || IsNil(o.Archive) {
		return nil, false
	}
	return o.Archive, true
}

// HasArchive returns a boolean if a field has been set.
func (o *CampaignContent1) HasArchive() bool {
	if o != nil && !IsNil(o.Archive) {
		return true
	}

	return false
}

// SetArchive gets a reference to the given UploadArchive and assigns it to the Archive field.
func (o *CampaignContent1) SetArchive(v UploadArchive) {
	o.Archive = &v
}

// GetVariateContents returns the VariateContents field value if set, zero value otherwise.
func (o *CampaignContent1) GetVariateContents() []VariateContentsInner1 {
	if o == nil || IsNil(o.VariateContents) {
		var ret []VariateContentsInner1
		return ret
	}
	return o.VariateContents
}

// GetVariateContentsOk returns a tuple with the VariateContents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignContent1) GetVariateContentsOk() ([]VariateContentsInner1, bool) {
	if o == nil || IsNil(o.VariateContents) {
		return nil, false
	}
	return o.VariateContents, true
}

// HasVariateContents returns a boolean if a field has been set.
func (o *CampaignContent1) HasVariateContents() bool {
	if o != nil && !IsNil(o.VariateContents) {
		return true
	}

	return false
}

// SetVariateContents gets a reference to the given []VariateContentsInner1 and assigns it to the VariateContents field.
func (o *CampaignContent1) SetVariateContents(v []VariateContentsInner1) {
	o.VariateContents = v
}

func (o CampaignContent1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignContent1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PlainText) {
		toSerialize["plain_text"] = o.PlainText
	}
	if !IsNil(o.Html) {
		toSerialize["html"] = o.Html
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	if !IsNil(o.Archive) {
		toSerialize["archive"] = o.Archive
	}
	if !IsNil(o.VariateContents) {
		toSerialize["variate_contents"] = o.VariateContents
	}
	return toSerialize, nil
}

type NullableCampaignContent1 struct {
	value *CampaignContent1
	isSet bool
}

func (v NullableCampaignContent1) Get() *CampaignContent1 {
	return v.value
}

func (v *NullableCampaignContent1) Set(val *CampaignContent1) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignContent1) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignContent1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignContent1(val *CampaignContent1) *NullableCampaignContent1 {
	return &NullableCampaignContent1{value: val, isSet: true}
}

func (v NullableCampaignContent1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignContent1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


