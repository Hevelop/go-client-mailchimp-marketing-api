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

// checks if the CampaignSettings5 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignSettings5{}

// CampaignSettings5 The settings for your campaign, including subject, from name, reply-to address, and more.
type CampaignSettings5 struct {
	// The subject line for the campaign.
	SubjectLine *string `json:"subject_line,omitempty"`
	// The preview text for the campaign.
	PreviewText *string `json:"preview_text,omitempty"`
	// The title of the campaign.
	Title *string `json:"title,omitempty"`
	// The 'from' name on the campaign (not an email address).
	FromName *string `json:"from_name,omitempty"`
	// The reply-to email address for the campaign.
	ReplyTo *string `json:"reply_to,omitempty"`
	// Use Mailchimp Conversation feature to manage replies.
	UseConversation *bool `json:"use_conversation,omitempty"`
	// The campaign's custom 'To' name. Typically the first name [audience field](https://mailchimp.com/help/getting-started-with-merge-tags/).
	ToName *string `json:"to_name,omitempty"`
	// If the campaign is listed in a folder, the id for that folder.
	FolderId *string `json:"folder_id,omitempty"`
	// Whether Mailchimp [authenticated](https://mailchimp.com/help/about-email-authentication/) the campaign. Defaults to `true`.
	Authenticate *bool `json:"authenticate,omitempty"`
	// Automatically append Mailchimp's [default footer](https://mailchimp.com/help/about-campaign-footers/) to the campaign.
	AutoFooter *bool `json:"auto_footer,omitempty"`
	// Automatically inline the CSS included with the campaign content.
	InlineCss *bool `json:"inline_css,omitempty"`
	// Automatically tweet a link to the [campaign archive](https://mailchimp.com/help/about-email-campaign-archives-and-pages/) page when the campaign is sent.
	AutoTweet *bool `json:"auto_tweet,omitempty"`
	// An array of [Facebook](https://mailchimp.com/help/connect-or-disconnect-the-facebook-integration/) page ids to auto-post to.
	AutoFbPost []string `json:"auto_fb_post,omitempty"`
	// Allows Facebook comments on the campaign (also force-enables the Campaign Archive toolbar). Defaults to `true`.
	FbComments *bool `json:"fb_comments,omitempty"`
	// Send this campaign using [Timewarp](https://mailchimp.com/help/use-timewarp/).
	Timewarp *bool `json:"timewarp,omitempty"`
	// The id for the template used in this campaign.
	TemplateId *int32 `json:"template_id,omitempty"`
	// Whether the campaign uses the drag-and-drop editor.
	DragAndDrop *bool `json:"drag_and_drop,omitempty"`
}

// NewCampaignSettings5 instantiates a new CampaignSettings5 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignSettings5() *CampaignSettings5 {
	this := CampaignSettings5{}
	return &this
}

// NewCampaignSettings5WithDefaults instantiates a new CampaignSettings5 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignSettings5WithDefaults() *CampaignSettings5 {
	this := CampaignSettings5{}
	return &this
}

// GetSubjectLine returns the SubjectLine field value if set, zero value otherwise.
func (o *CampaignSettings5) GetSubjectLine() string {
	if o == nil || IsNil(o.SubjectLine) {
		var ret string
		return ret
	}
	return *o.SubjectLine
}

// GetSubjectLineOk returns a tuple with the SubjectLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignSettings5) GetSubjectLineOk() (*string, bool) {
	if o == nil || IsNil(o.SubjectLine) {
		return nil, false
	}
	return o.SubjectLine, true
}

// HasSubjectLine returns a boolean if a field has been set.
func (o *CampaignSettings5) HasSubjectLine() bool {
	if o != nil && !IsNil(o.SubjectLine) {
		return true
	}

	return false
}

// SetSubjectLine gets a reference to the given string and assigns it to the SubjectLine field.
func (o *CampaignSettings5) SetSubjectLine(v string) {
	o.SubjectLine = &v
}

// GetPreviewText returns the PreviewText field value if set, zero value otherwise.
func (o *CampaignSettings5) GetPreviewText() string {
	if o == nil || IsNil(o.PreviewText) {
		var ret string
		return ret
	}
	return *o.PreviewText
}

// GetPreviewTextOk returns a tuple with the PreviewText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignSettings5) GetPreviewTextOk() (*string, bool) {
	if o == nil || IsNil(o.PreviewText) {
		return nil, false
	}
	return o.PreviewText, true
}

// HasPreviewText returns a boolean if a field has been set.
func (o *CampaignSettings5) HasPreviewText() bool {
	if o != nil && !IsNil(o.PreviewText) {
		return true
	}

	return false
}

// SetPreviewText gets a reference to the given string and assigns it to the PreviewText field.
func (o *CampaignSettings5) SetPreviewText(v string) {
	o.PreviewText = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *CampaignSettings5) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignSettings5) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *CampaignSettings5) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *CampaignSettings5) SetTitle(v string) {
	o.Title = &v
}

// GetFromName returns the FromName field value if set, zero value otherwise.
func (o *CampaignSettings5) GetFromName() string {
	if o == nil || IsNil(o.FromName) {
		var ret string
		return ret
	}
	return *o.FromName
}

// GetFromNameOk returns a tuple with the FromName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignSettings5) GetFromNameOk() (*string, bool) {
	if o == nil || IsNil(o.FromName) {
		return nil, false
	}
	return o.FromName, true
}

// HasFromName returns a boolean if a field has been set.
func (o *CampaignSettings5) HasFromName() bool {
	if o != nil && !IsNil(o.FromName) {
		return true
	}

	return false
}

// SetFromName gets a reference to the given string and assigns it to the FromName field.
func (o *CampaignSettings5) SetFromName(v string) {
	o.FromName = &v
}

// GetReplyTo returns the ReplyTo field value if set, zero value otherwise.
func (o *CampaignSettings5) GetReplyTo() string {
	if o == nil || IsNil(o.ReplyTo) {
		var ret string
		return ret
	}
	return *o.ReplyTo
}

// GetReplyToOk returns a tuple with the ReplyTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignSettings5) GetReplyToOk() (*string, bool) {
	if o == nil || IsNil(o.ReplyTo) {
		return nil, false
	}
	return o.ReplyTo, true
}

// HasReplyTo returns a boolean if a field has been set.
func (o *CampaignSettings5) HasReplyTo() bool {
	if o != nil && !IsNil(o.ReplyTo) {
		return true
	}

	return false
}

// SetReplyTo gets a reference to the given string and assigns it to the ReplyTo field.
func (o *CampaignSettings5) SetReplyTo(v string) {
	o.ReplyTo = &v
}

// GetUseConversation returns the UseConversation field value if set, zero value otherwise.
func (o *CampaignSettings5) GetUseConversation() bool {
	if o == nil || IsNil(o.UseConversation) {
		var ret bool
		return ret
	}
	return *o.UseConversation
}

// GetUseConversationOk returns a tuple with the UseConversation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignSettings5) GetUseConversationOk() (*bool, bool) {
	if o == nil || IsNil(o.UseConversation) {
		return nil, false
	}
	return o.UseConversation, true
}

// HasUseConversation returns a boolean if a field has been set.
func (o *CampaignSettings5) HasUseConversation() bool {
	if o != nil && !IsNil(o.UseConversation) {
		return true
	}

	return false
}

// SetUseConversation gets a reference to the given bool and assigns it to the UseConversation field.
func (o *CampaignSettings5) SetUseConversation(v bool) {
	o.UseConversation = &v
}

// GetToName returns the ToName field value if set, zero value otherwise.
func (o *CampaignSettings5) GetToName() string {
	if o == nil || IsNil(o.ToName) {
		var ret string
		return ret
	}
	return *o.ToName
}

// GetToNameOk returns a tuple with the ToName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignSettings5) GetToNameOk() (*string, bool) {
	if o == nil || IsNil(o.ToName) {
		return nil, false
	}
	return o.ToName, true
}

// HasToName returns a boolean if a field has been set.
func (o *CampaignSettings5) HasToName() bool {
	if o != nil && !IsNil(o.ToName) {
		return true
	}

	return false
}

// SetToName gets a reference to the given string and assigns it to the ToName field.
func (o *CampaignSettings5) SetToName(v string) {
	o.ToName = &v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *CampaignSettings5) GetFolderId() string {
	if o == nil || IsNil(o.FolderId) {
		var ret string
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignSettings5) GetFolderIdOk() (*string, bool) {
	if o == nil || IsNil(o.FolderId) {
		return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *CampaignSettings5) HasFolderId() bool {
	if o != nil && !IsNil(o.FolderId) {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given string and assigns it to the FolderId field.
func (o *CampaignSettings5) SetFolderId(v string) {
	o.FolderId = &v
}

// GetAuthenticate returns the Authenticate field value if set, zero value otherwise.
func (o *CampaignSettings5) GetAuthenticate() bool {
	if o == nil || IsNil(o.Authenticate) {
		var ret bool
		return ret
	}
	return *o.Authenticate
}

// GetAuthenticateOk returns a tuple with the Authenticate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignSettings5) GetAuthenticateOk() (*bool, bool) {
	if o == nil || IsNil(o.Authenticate) {
		return nil, false
	}
	return o.Authenticate, true
}

// HasAuthenticate returns a boolean if a field has been set.
func (o *CampaignSettings5) HasAuthenticate() bool {
	if o != nil && !IsNil(o.Authenticate) {
		return true
	}

	return false
}

// SetAuthenticate gets a reference to the given bool and assigns it to the Authenticate field.
func (o *CampaignSettings5) SetAuthenticate(v bool) {
	o.Authenticate = &v
}

// GetAutoFooter returns the AutoFooter field value if set, zero value otherwise.
func (o *CampaignSettings5) GetAutoFooter() bool {
	if o == nil || IsNil(o.AutoFooter) {
		var ret bool
		return ret
	}
	return *o.AutoFooter
}

// GetAutoFooterOk returns a tuple with the AutoFooter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignSettings5) GetAutoFooterOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoFooter) {
		return nil, false
	}
	return o.AutoFooter, true
}

// HasAutoFooter returns a boolean if a field has been set.
func (o *CampaignSettings5) HasAutoFooter() bool {
	if o != nil && !IsNil(o.AutoFooter) {
		return true
	}

	return false
}

// SetAutoFooter gets a reference to the given bool and assigns it to the AutoFooter field.
func (o *CampaignSettings5) SetAutoFooter(v bool) {
	o.AutoFooter = &v
}

// GetInlineCss returns the InlineCss field value if set, zero value otherwise.
func (o *CampaignSettings5) GetInlineCss() bool {
	if o == nil || IsNil(o.InlineCss) {
		var ret bool
		return ret
	}
	return *o.InlineCss
}

// GetInlineCssOk returns a tuple with the InlineCss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignSettings5) GetInlineCssOk() (*bool, bool) {
	if o == nil || IsNil(o.InlineCss) {
		return nil, false
	}
	return o.InlineCss, true
}

// HasInlineCss returns a boolean if a field has been set.
func (o *CampaignSettings5) HasInlineCss() bool {
	if o != nil && !IsNil(o.InlineCss) {
		return true
	}

	return false
}

// SetInlineCss gets a reference to the given bool and assigns it to the InlineCss field.
func (o *CampaignSettings5) SetInlineCss(v bool) {
	o.InlineCss = &v
}

// GetAutoTweet returns the AutoTweet field value if set, zero value otherwise.
func (o *CampaignSettings5) GetAutoTweet() bool {
	if o == nil || IsNil(o.AutoTweet) {
		var ret bool
		return ret
	}
	return *o.AutoTweet
}

// GetAutoTweetOk returns a tuple with the AutoTweet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignSettings5) GetAutoTweetOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoTweet) {
		return nil, false
	}
	return o.AutoTweet, true
}

// HasAutoTweet returns a boolean if a field has been set.
func (o *CampaignSettings5) HasAutoTweet() bool {
	if o != nil && !IsNil(o.AutoTweet) {
		return true
	}

	return false
}

// SetAutoTweet gets a reference to the given bool and assigns it to the AutoTweet field.
func (o *CampaignSettings5) SetAutoTweet(v bool) {
	o.AutoTweet = &v
}

// GetAutoFbPost returns the AutoFbPost field value if set, zero value otherwise.
func (o *CampaignSettings5) GetAutoFbPost() []string {
	if o == nil || IsNil(o.AutoFbPost) {
		var ret []string
		return ret
	}
	return o.AutoFbPost
}

// GetAutoFbPostOk returns a tuple with the AutoFbPost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignSettings5) GetAutoFbPostOk() ([]string, bool) {
	if o == nil || IsNil(o.AutoFbPost) {
		return nil, false
	}
	return o.AutoFbPost, true
}

// HasAutoFbPost returns a boolean if a field has been set.
func (o *CampaignSettings5) HasAutoFbPost() bool {
	if o != nil && !IsNil(o.AutoFbPost) {
		return true
	}

	return false
}

// SetAutoFbPost gets a reference to the given []string and assigns it to the AutoFbPost field.
func (o *CampaignSettings5) SetAutoFbPost(v []string) {
	o.AutoFbPost = v
}

// GetFbComments returns the FbComments field value if set, zero value otherwise.
func (o *CampaignSettings5) GetFbComments() bool {
	if o == nil || IsNil(o.FbComments) {
		var ret bool
		return ret
	}
	return *o.FbComments
}

// GetFbCommentsOk returns a tuple with the FbComments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignSettings5) GetFbCommentsOk() (*bool, bool) {
	if o == nil || IsNil(o.FbComments) {
		return nil, false
	}
	return o.FbComments, true
}

// HasFbComments returns a boolean if a field has been set.
func (o *CampaignSettings5) HasFbComments() bool {
	if o != nil && !IsNil(o.FbComments) {
		return true
	}

	return false
}

// SetFbComments gets a reference to the given bool and assigns it to the FbComments field.
func (o *CampaignSettings5) SetFbComments(v bool) {
	o.FbComments = &v
}

// GetTimewarp returns the Timewarp field value if set, zero value otherwise.
func (o *CampaignSettings5) GetTimewarp() bool {
	if o == nil || IsNil(o.Timewarp) {
		var ret bool
		return ret
	}
	return *o.Timewarp
}

// GetTimewarpOk returns a tuple with the Timewarp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignSettings5) GetTimewarpOk() (*bool, bool) {
	if o == nil || IsNil(o.Timewarp) {
		return nil, false
	}
	return o.Timewarp, true
}

// HasTimewarp returns a boolean if a field has been set.
func (o *CampaignSettings5) HasTimewarp() bool {
	if o != nil && !IsNil(o.Timewarp) {
		return true
	}

	return false
}

// SetTimewarp gets a reference to the given bool and assigns it to the Timewarp field.
func (o *CampaignSettings5) SetTimewarp(v bool) {
	o.Timewarp = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *CampaignSettings5) GetTemplateId() int32 {
	if o == nil || IsNil(o.TemplateId) {
		var ret int32
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignSettings5) GetTemplateIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *CampaignSettings5) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given int32 and assigns it to the TemplateId field.
func (o *CampaignSettings5) SetTemplateId(v int32) {
	o.TemplateId = &v
}

// GetDragAndDrop returns the DragAndDrop field value if set, zero value otherwise.
func (o *CampaignSettings5) GetDragAndDrop() bool {
	if o == nil || IsNil(o.DragAndDrop) {
		var ret bool
		return ret
	}
	return *o.DragAndDrop
}

// GetDragAndDropOk returns a tuple with the DragAndDrop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignSettings5) GetDragAndDropOk() (*bool, bool) {
	if o == nil || IsNil(o.DragAndDrop) {
		return nil, false
	}
	return o.DragAndDrop, true
}

// HasDragAndDrop returns a boolean if a field has been set.
func (o *CampaignSettings5) HasDragAndDrop() bool {
	if o != nil && !IsNil(o.DragAndDrop) {
		return true
	}

	return false
}

// SetDragAndDrop gets a reference to the given bool and assigns it to the DragAndDrop field.
func (o *CampaignSettings5) SetDragAndDrop(v bool) {
	o.DragAndDrop = &v
}

func (o CampaignSettings5) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignSettings5) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubjectLine) {
		toSerialize["subject_line"] = o.SubjectLine
	}
	if !IsNil(o.PreviewText) {
		toSerialize["preview_text"] = o.PreviewText
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.FromName) {
		toSerialize["from_name"] = o.FromName
	}
	if !IsNil(o.ReplyTo) {
		toSerialize["reply_to"] = o.ReplyTo
	}
	if !IsNil(o.UseConversation) {
		toSerialize["use_conversation"] = o.UseConversation
	}
	if !IsNil(o.ToName) {
		toSerialize["to_name"] = o.ToName
	}
	if !IsNil(o.FolderId) {
		toSerialize["folder_id"] = o.FolderId
	}
	if !IsNil(o.Authenticate) {
		toSerialize["authenticate"] = o.Authenticate
	}
	if !IsNil(o.AutoFooter) {
		toSerialize["auto_footer"] = o.AutoFooter
	}
	if !IsNil(o.InlineCss) {
		toSerialize["inline_css"] = o.InlineCss
	}
	if !IsNil(o.AutoTweet) {
		toSerialize["auto_tweet"] = o.AutoTweet
	}
	if !IsNil(o.AutoFbPost) {
		toSerialize["auto_fb_post"] = o.AutoFbPost
	}
	if !IsNil(o.FbComments) {
		toSerialize["fb_comments"] = o.FbComments
	}
	if !IsNil(o.Timewarp) {
		toSerialize["timewarp"] = o.Timewarp
	}
	if !IsNil(o.TemplateId) {
		toSerialize["template_id"] = o.TemplateId
	}
	if !IsNil(o.DragAndDrop) {
		toSerialize["drag_and_drop"] = o.DragAndDrop
	}
	return toSerialize, nil
}

type NullableCampaignSettings5 struct {
	value *CampaignSettings5
	isSet bool
}

func (v NullableCampaignSettings5) Get() *CampaignSettings5 {
	return v.value
}

func (v *NullableCampaignSettings5) Set(val *CampaignSettings5) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignSettings5) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignSettings5) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignSettings5(val *CampaignSettings5) *NullableCampaignSettings5 {
	return &NullableCampaignSettings5{value: val, isSet: true}
}

func (v NullableCampaignSettings5) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignSettings5) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


