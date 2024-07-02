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

// checks if the Location2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Location2{}

// Location2 Subscriber location information.
type Location2 struct {
	// The location latitude.
	Latitude *float32 `json:"latitude,omitempty"`
	// The location longitude.
	Longitude *float32 `json:"longitude,omitempty"`
	// The time difference in hours from GMT.
	Gmtoff *int32 `json:"gmtoff,omitempty"`
	// The offset for timezones where daylight saving time is observed.
	Dstoff *int32 `json:"dstoff,omitempty"`
	// The unique code for the location country.
	CountryCode *string `json:"country_code,omitempty"`
	// The timezone for the location.
	Timezone *string `json:"timezone,omitempty"`
	// The region for the location.
	Region *string `json:"region,omitempty"`
}

// NewLocation2 instantiates a new Location2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocation2() *Location2 {
	this := Location2{}
	return &this
}

// NewLocation2WithDefaults instantiates a new Location2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocation2WithDefaults() *Location2 {
	this := Location2{}
	return &this
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *Location2) GetLatitude() float32 {
	if o == nil || IsNil(o.Latitude) {
		var ret float32
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location2) GetLatitudeOk() (*float32, bool) {
	if o == nil || IsNil(o.Latitude) {
		return nil, false
	}
	return o.Latitude, true
}

// HaveLatitude returns a boolean if a field has been set.
func (o *Location2) HaveLatitude() bool {
	if o != nil && !IsNil(o.Latitude) {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given float32 and assigns it to the Latitude field.
func (o *Location2) SetLatitude(v float32) {
	o.Latitude = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *Location2) GetLongitude() float32 {
	if o == nil || IsNil(o.Longitude) {
		var ret float32
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location2) GetLongitudeOk() (*float32, bool) {
	if o == nil || IsNil(o.Longitude) {
		return nil, false
	}
	return o.Longitude, true
}

// HaveLongitude returns a boolean if a field has been set.
func (o *Location2) HaveLongitude() bool {
	if o != nil && !IsNil(o.Longitude) {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given float32 and assigns it to the Longitude field.
func (o *Location2) SetLongitude(v float32) {
	o.Longitude = &v
}

// GetGmtoff returns the Gmtoff field value if set, zero value otherwise.
func (o *Location2) GetGmtoff() int32 {
	if o == nil || IsNil(o.Gmtoff) {
		var ret int32
		return ret
	}
	return *o.Gmtoff
}

// GetGmtoffOk returns a tuple with the Gmtoff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location2) GetGmtoffOk() (*int32, bool) {
	if o == nil || IsNil(o.Gmtoff) {
		return nil, false
	}
	return o.Gmtoff, true
}

// HaveGmtoff returns a boolean if a field has been set.
func (o *Location2) HaveGmtoff() bool {
	if o != nil && !IsNil(o.Gmtoff) {
		return true
	}

	return false
}

// SetGmtoff gets a reference to the given int32 and assigns it to the Gmtoff field.
func (o *Location2) SetGmtoff(v int32) {
	o.Gmtoff = &v
}

// GetDstoff returns the Dstoff field value if set, zero value otherwise.
func (o *Location2) GetDstoff() int32 {
	if o == nil || IsNil(o.Dstoff) {
		var ret int32
		return ret
	}
	return *o.Dstoff
}

// GetDstoffOk returns a tuple with the Dstoff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location2) GetDstoffOk() (*int32, bool) {
	if o == nil || IsNil(o.Dstoff) {
		return nil, false
	}
	return o.Dstoff, true
}

// HaveDstoff returns a boolean if a field has been set.
func (o *Location2) HaveDstoff() bool {
	if o != nil && !IsNil(o.Dstoff) {
		return true
	}

	return false
}

// SetDstoff gets a reference to the given int32 and assigns it to the Dstoff field.
func (o *Location2) SetDstoff(v int32) {
	o.Dstoff = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *Location2) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location2) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HaveCountryCode returns a boolean if a field has been set.
func (o *Location2) HaveCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *Location2) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *Location2) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location2) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HaveTimezone returns a boolean if a field has been set.
func (o *Location2) HaveTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *Location2) SetTimezone(v string) {
	o.Timezone = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *Location2) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location2) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HaveRegion returns a boolean if a field has been set.
func (o *Location2) HaveRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *Location2) SetRegion(v string) {
	o.Region = &v
}

func (o Location2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Location2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Latitude) {
		toSerialize["latitude"] = o.Latitude
	}
	if !IsNil(o.Longitude) {
		toSerialize["longitude"] = o.Longitude
	}
	if !IsNil(o.Gmtoff) {
		toSerialize["gmtoff"] = o.Gmtoff
	}
	if !IsNil(o.Dstoff) {
		toSerialize["dstoff"] = o.Dstoff
	}
	if !IsNil(o.CountryCode) {
		toSerialize["country_code"] = o.CountryCode
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	return toSerialize, nil
}

type NullableLocation2 struct {
	value *Location2
	isSet bool
}

func (v NullableLocation2) Get() *Location2 {
	return v.value
}

func (v *NullableLocation2) Set(val *Location2) {
	v.value = val
	v.isSet = true
}

func (v NullableLocation2) IsSet() bool {
	return v.isSet
}

func (v *NullableLocation2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocation2(val *Location2) *NullableLocation2 {
	return &NullableLocation2{value: val, isSet: true}
}

func (v NullableLocation2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocation2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


