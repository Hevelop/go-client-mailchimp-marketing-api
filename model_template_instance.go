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

// checks if the TemplateInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateInstance{}

// TemplateInstance Information about a specific template.
type TemplateInstance struct {
	// The individual id for the template.
	Id *int32 `json:"id,omitempty"`
	// The type of template (user, base, or gallery).
	Type *string `json:"type,omitempty"`
	// The name of the template.
	Name *string `json:"name,omitempty"`
	// Whether the template uses the drag and drop editor.
	DragAndDrop *bool `json:"drag_and_drop,omitempty"`
	// Whether the template contains media queries to make it responsive.
	Responsive *bool `json:"responsive,omitempty"`
	// If available, the category the template is listed in.
	Category *string `json:"category,omitempty"`
	// The date and time the template was created in ISO 8601 format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time the template was edited in ISO 8601 format.
	DateEdited *time.Time `json:"date_edited,omitempty"`
	// The login name for template's creator.
	CreatedBy *string `json:"created_by,omitempty"`
	// The login name who last edited the template.
	EditedBy *string `json:"edited_by,omitempty"`
	// User templates are not 'deleted,' but rather marked as 'inactive.' Returns whether the template is still active.
	Active *bool `json:"active,omitempty"`
	// The id of the folder the template is currently in.
	FolderId *string `json:"folder_id,omitempty"`
	// If available, the URL for a thumbnail of the template.
	Thumbnail *string `json:"thumbnail,omitempty"`
	// The URL used for [template sharing](https://mailchimp.com/help/share-a-template/).
	ShareUrl *string `json:"share_url,omitempty"`
	// How the template's content is put together.
	ContentType *string `json:"content_type,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewTemplateInstance instantiates a new TemplateInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateInstance() *TemplateInstance {
	this := TemplateInstance{}
	return &this
}

// NewTemplateInstanceWithDefaults instantiates a new TemplateInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateInstanceWithDefaults() *TemplateInstance {
	this := TemplateInstance{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TemplateInstance) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateInstance) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TemplateInstance) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *TemplateInstance) SetId(v int32) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TemplateInstance) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateInstance) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TemplateInstance) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TemplateInstance) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TemplateInstance) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateInstance) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TemplateInstance) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TemplateInstance) SetName(v string) {
	o.Name = &v
}

// GetDragAndDrop returns the DragAndDrop field value if set, zero value otherwise.
func (o *TemplateInstance) GetDragAndDrop() bool {
	if o == nil || IsNil(o.DragAndDrop) {
		var ret bool
		return ret
	}
	return *o.DragAndDrop
}

// GetDragAndDropOk returns a tuple with the DragAndDrop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateInstance) GetDragAndDropOk() (*bool, bool) {
	if o == nil || IsNil(o.DragAndDrop) {
		return nil, false
	}
	return o.DragAndDrop, true
}

// HasDragAndDrop returns a boolean if a field has been set.
func (o *TemplateInstance) HasDragAndDrop() bool {
	if o != nil && !IsNil(o.DragAndDrop) {
		return true
	}

	return false
}

// SetDragAndDrop gets a reference to the given bool and assigns it to the DragAndDrop field.
func (o *TemplateInstance) SetDragAndDrop(v bool) {
	o.DragAndDrop = &v
}

// GetResponsive returns the Responsive field value if set, zero value otherwise.
func (o *TemplateInstance) GetResponsive() bool {
	if o == nil || IsNil(o.Responsive) {
		var ret bool
		return ret
	}
	return *o.Responsive
}

// GetResponsiveOk returns a tuple with the Responsive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateInstance) GetResponsiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Responsive) {
		return nil, false
	}
	return o.Responsive, true
}

// HasResponsive returns a boolean if a field has been set.
func (o *TemplateInstance) HasResponsive() bool {
	if o != nil && !IsNil(o.Responsive) {
		return true
	}

	return false
}

// SetResponsive gets a reference to the given bool and assigns it to the Responsive field.
func (o *TemplateInstance) SetResponsive(v bool) {
	o.Responsive = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *TemplateInstance) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateInstance) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *TemplateInstance) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *TemplateInstance) SetCategory(v string) {
	o.Category = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *TemplateInstance) GetDateCreated() time.Time {
	if o == nil || IsNil(o.DateCreated) {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateInstance) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *TemplateInstance) HasDateCreated() bool {
	if o != nil && !IsNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *TemplateInstance) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetDateEdited returns the DateEdited field value if set, zero value otherwise.
func (o *TemplateInstance) GetDateEdited() time.Time {
	if o == nil || IsNil(o.DateEdited) {
		var ret time.Time
		return ret
	}
	return *o.DateEdited
}

// GetDateEditedOk returns a tuple with the DateEdited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateInstance) GetDateEditedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateEdited) {
		return nil, false
	}
	return o.DateEdited, true
}

// HasDateEdited returns a boolean if a field has been set.
func (o *TemplateInstance) HasDateEdited() bool {
	if o != nil && !IsNil(o.DateEdited) {
		return true
	}

	return false
}

// SetDateEdited gets a reference to the given time.Time and assigns it to the DateEdited field.
func (o *TemplateInstance) SetDateEdited(v time.Time) {
	o.DateEdited = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *TemplateInstance) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateInstance) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *TemplateInstance) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *TemplateInstance) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetEditedBy returns the EditedBy field value if set, zero value otherwise.
func (o *TemplateInstance) GetEditedBy() string {
	if o == nil || IsNil(o.EditedBy) {
		var ret string
		return ret
	}
	return *o.EditedBy
}

// GetEditedByOk returns a tuple with the EditedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateInstance) GetEditedByOk() (*string, bool) {
	if o == nil || IsNil(o.EditedBy) {
		return nil, false
	}
	return o.EditedBy, true
}

// HasEditedBy returns a boolean if a field has been set.
func (o *TemplateInstance) HasEditedBy() bool {
	if o != nil && !IsNil(o.EditedBy) {
		return true
	}

	return false
}

// SetEditedBy gets a reference to the given string and assigns it to the EditedBy field.
func (o *TemplateInstance) SetEditedBy(v string) {
	o.EditedBy = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *TemplateInstance) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateInstance) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *TemplateInstance) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *TemplateInstance) SetActive(v bool) {
	o.Active = &v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *TemplateInstance) GetFolderId() string {
	if o == nil || IsNil(o.FolderId) {
		var ret string
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateInstance) GetFolderIdOk() (*string, bool) {
	if o == nil || IsNil(o.FolderId) {
		return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *TemplateInstance) HasFolderId() bool {
	if o != nil && !IsNil(o.FolderId) {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given string and assigns it to the FolderId field.
func (o *TemplateInstance) SetFolderId(v string) {
	o.FolderId = &v
}

// GetThumbnail returns the Thumbnail field value if set, zero value otherwise.
func (o *TemplateInstance) GetThumbnail() string {
	if o == nil || IsNil(o.Thumbnail) {
		var ret string
		return ret
	}
	return *o.Thumbnail
}

// GetThumbnailOk returns a tuple with the Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateInstance) GetThumbnailOk() (*string, bool) {
	if o == nil || IsNil(o.Thumbnail) {
		return nil, false
	}
	return o.Thumbnail, true
}

// HasThumbnail returns a boolean if a field has been set.
func (o *TemplateInstance) HasThumbnail() bool {
	if o != nil && !IsNil(o.Thumbnail) {
		return true
	}

	return false
}

// SetThumbnail gets a reference to the given string and assigns it to the Thumbnail field.
func (o *TemplateInstance) SetThumbnail(v string) {
	o.Thumbnail = &v
}

// GetShareUrl returns the ShareUrl field value if set, zero value otherwise.
func (o *TemplateInstance) GetShareUrl() string {
	if o == nil || IsNil(o.ShareUrl) {
		var ret string
		return ret
	}
	return *o.ShareUrl
}

// GetShareUrlOk returns a tuple with the ShareUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateInstance) GetShareUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ShareUrl) {
		return nil, false
	}
	return o.ShareUrl, true
}

// HasShareUrl returns a boolean if a field has been set.
func (o *TemplateInstance) HasShareUrl() bool {
	if o != nil && !IsNil(o.ShareUrl) {
		return true
	}

	return false
}

// SetShareUrl gets a reference to the given string and assigns it to the ShareUrl field.
func (o *TemplateInstance) SetShareUrl(v string) {
	o.ShareUrl = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *TemplateInstance) GetContentType() string {
	if o == nil || IsNil(o.ContentType) {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateInstance) GetContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *TemplateInstance) HasContentType() bool {
	if o != nil && !IsNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *TemplateInstance) SetContentType(v string) {
	o.ContentType = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TemplateInstance) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateInstance) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TemplateInstance) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *TemplateInstance) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o TemplateInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.DragAndDrop) {
		toSerialize["drag_and_drop"] = o.DragAndDrop
	}
	if !IsNil(o.Responsive) {
		toSerialize["responsive"] = o.Responsive
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.DateCreated) {
		toSerialize["date_created"] = o.DateCreated
	}
	if !IsNil(o.DateEdited) {
		toSerialize["date_edited"] = o.DateEdited
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["created_by"] = o.CreatedBy
	}
	if !IsNil(o.EditedBy) {
		toSerialize["edited_by"] = o.EditedBy
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.FolderId) {
		toSerialize["folder_id"] = o.FolderId
	}
	if !IsNil(o.Thumbnail) {
		toSerialize["thumbnail"] = o.Thumbnail
	}
	if !IsNil(o.ShareUrl) {
		toSerialize["share_url"] = o.ShareUrl
	}
	if !IsNil(o.ContentType) {
		toSerialize["content_type"] = o.ContentType
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableTemplateInstance struct {
	value *TemplateInstance
	isSet bool
}

func (v NullableTemplateInstance) Get() *TemplateInstance {
	return v.value
}

func (v *NullableTemplateInstance) Set(val *TemplateInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateInstance(val *TemplateInstance) *NullableTemplateInstance {
	return &NullableTemplateInstance{value: val, isSet: true}
}

func (v NullableTemplateInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


