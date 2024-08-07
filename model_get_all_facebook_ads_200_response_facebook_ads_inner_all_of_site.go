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

// checks if the GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite{}

// GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite Connected Site
type GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite struct {
	// The ID of this connected site.
	Id *int32 `json:"id,omitempty"`
	// The name of the connected site
	Name *string `json:"name,omitempty"`
	// The URL for this connected site.
	Url *string `json:"url,omitempty"`
}

// NewGetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite instantiates a new GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite() *GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite {
	this := GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite{}
	return &this
}

// NewGetAllFacebookAds200ResponseFacebookAdsInnerAllOfSiteWithDefaults instantiates a new GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllFacebookAds200ResponseFacebookAdsInnerAllOfSiteWithDefaults() *GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite {
	this := GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite) SetUrl(v string) {
	o.Url = &v
}

func (o GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableGetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite struct {
	value *GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite
	isSet bool
}

func (v NullableGetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite) Get() *GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite {
	return v.value
}

func (v *NullableGetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite) Set(val *GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite(val *GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite) *NullableGetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite {
	return &NullableGetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite{value: val, isSet: true}
}

func (v NullableGetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


