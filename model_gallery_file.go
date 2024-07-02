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

// checks if the GalleryFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GalleryFile{}

// GalleryFile An individual file listed in the File Manager.
type GalleryFile struct {
	// The unique id of the file.
	Id *int32 `json:"id,omitempty"`
	// The id of the folder.
	FolderId *int32 `json:"folder_id,omitempty"`
	// The type of file in the File Manager.
	Type *string `json:"type,omitempty"`
	// The name of the file.
	Name *string `json:"name,omitempty"`
	// The url of the full-size file.
	FullSizeUrl *string `json:"full_size_url,omitempty"`
	// The url of the thumbnail preview.
	ThumbnailUrl *string `json:"thumbnail_url,omitempty"`
	// The size of the file in bytes.
	Size *int32 `json:"size,omitempty"`
	// The date and time a file was added to the File Manager in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The username of the profile that uploaded the file.
	CreatedBy *string `json:"created_by,omitempty"`
	// The width of the image.
	Width *int32 `json:"width,omitempty"`
	// The height of an image.
	Height *int32 `json:"height,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewGalleryFile instantiates a new GalleryFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGalleryFile() *GalleryFile {
	this := GalleryFile{}
	return &this
}

// NewGalleryFileWithDefaults instantiates a new GalleryFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGalleryFileWithDefaults() *GalleryFile {
	this := GalleryFile{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GalleryFile) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GalleryFile) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HaveId returns a boolean if a field has been set.
func (o *GalleryFile) HaveId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GalleryFile) SetId(v int32) {
	o.Id = &v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *GalleryFile) GetFolderId() int32 {
	if o == nil || IsNil(o.FolderId) {
		var ret int32
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GalleryFile) GetFolderIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FolderId) {
		return nil, false
	}
	return o.FolderId, true
}

// HaveFolderId returns a boolean if a field has been set.
func (o *GalleryFile) HaveFolderId() bool {
	if o != nil && !IsNil(o.FolderId) {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given int32 and assigns it to the FolderId field.
func (o *GalleryFile) SetFolderId(v int32) {
	o.FolderId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GalleryFile) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GalleryFile) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HaveType returns a boolean if a field has been set.
func (o *GalleryFile) HaveType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GalleryFile) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GalleryFile) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GalleryFile) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HaveName returns a boolean if a field has been set.
func (o *GalleryFile) HaveName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GalleryFile) SetName(v string) {
	o.Name = &v
}

// GetFullSizeUrl returns the FullSizeUrl field value if set, zero value otherwise.
func (o *GalleryFile) GetFullSizeUrl() string {
	if o == nil || IsNil(o.FullSizeUrl) {
		var ret string
		return ret
	}
	return *o.FullSizeUrl
}

// GetFullSizeUrlOk returns a tuple with the FullSizeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GalleryFile) GetFullSizeUrlOk() (*string, bool) {
	if o == nil || IsNil(o.FullSizeUrl) {
		return nil, false
	}
	return o.FullSizeUrl, true
}

// HaveFullSizeUrl returns a boolean if a field has been set.
func (o *GalleryFile) HaveFullSizeUrl() bool {
	if o != nil && !IsNil(o.FullSizeUrl) {
		return true
	}

	return false
}

// SetFullSizeUrl gets a reference to the given string and assigns it to the FullSizeUrl field.
func (o *GalleryFile) SetFullSizeUrl(v string) {
	o.FullSizeUrl = &v
}

// GetThumbnailUrl returns the ThumbnailUrl field value if set, zero value otherwise.
func (o *GalleryFile) GetThumbnailUrl() string {
	if o == nil || IsNil(o.ThumbnailUrl) {
		var ret string
		return ret
	}
	return *o.ThumbnailUrl
}

// GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GalleryFile) GetThumbnailUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ThumbnailUrl) {
		return nil, false
	}
	return o.ThumbnailUrl, true
}

// HaveThumbnailUrl returns a boolean if a field has been set.
func (o *GalleryFile) HaveThumbnailUrl() bool {
	if o != nil && !IsNil(o.ThumbnailUrl) {
		return true
	}

	return false
}

// SetThumbnailUrl gets a reference to the given string and assigns it to the ThumbnailUrl field.
func (o *GalleryFile) SetThumbnailUrl(v string) {
	o.ThumbnailUrl = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *GalleryFile) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GalleryFile) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HaveSize returns a boolean if a field has been set.
func (o *GalleryFile) HaveSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *GalleryFile) SetSize(v int32) {
	o.Size = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GalleryFile) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GalleryFile) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HaveCreatedAt returns a boolean if a field has been set.
func (o *GalleryFile) HaveCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *GalleryFile) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *GalleryFile) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GalleryFile) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HaveCreatedBy returns a boolean if a field has been set.
func (o *GalleryFile) HaveCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *GalleryFile) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *GalleryFile) GetWidth() int32 {
	if o == nil || IsNil(o.Width) {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GalleryFile) GetWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HaveWidth returns a boolean if a field has been set.
func (o *GalleryFile) HaveWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *GalleryFile) SetWidth(v int32) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *GalleryFile) GetHeight() int32 {
	if o == nil || IsNil(o.Height) {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GalleryFile) GetHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HaveHeight returns a boolean if a field has been set.
func (o *GalleryFile) HaveHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *GalleryFile) SetHeight(v int32) {
	o.Height = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GalleryFile) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GalleryFile) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HaveLinks returns a boolean if a field has been set.
func (o *GalleryFile) HaveLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *GalleryFile) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o GalleryFile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GalleryFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.FolderId) {
		toSerialize["folder_id"] = o.FolderId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.FullSizeUrl) {
		toSerialize["full_size_url"] = o.FullSizeUrl
	}
	if !IsNil(o.ThumbnailUrl) {
		toSerialize["thumbnail_url"] = o.ThumbnailUrl
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["created_by"] = o.CreatedBy
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableGalleryFile struct {
	value *GalleryFile
	isSet bool
}

func (v NullableGalleryFile) Get() *GalleryFile {
	return v.value
}

func (v *NullableGalleryFile) Set(val *GalleryFile) {
	v.value = val
	v.isSet = true
}

func (v NullableGalleryFile) IsSet() bool {
	return v.isSet
}

func (v *NullableGalleryFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGalleryFile(val *GalleryFile) *NullableGalleryFile {
	return &NullableGalleryFile{value: val, isSet: true}
}

func (v NullableGalleryFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGalleryFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


