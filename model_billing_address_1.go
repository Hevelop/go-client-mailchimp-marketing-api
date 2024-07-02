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

// checks if the BillingAddress1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingAddress1{}

// BillingAddress1 The billing address for the order.
type BillingAddress1 struct {
	// The name associated with the billing address.
	Name *string `json:"name,omitempty"`
	// The billing address for the order.
	Address1 *string `json:"address1,omitempty"`
	// An additional field for the billing address.
	Address2 *string `json:"address2,omitempty"`
	// The city in the billing address.
	City *string `json:"city,omitempty"`
	// The state or normalized province in the billing address.
	Province *string `json:"province,omitempty"`
	// The two-letter code for the province in the billing address.
	ProvinceCode *string `json:"province_code,omitempty"`
	// The postal or zip code in the billing address.
	PostalCode *string `json:"postal_code,omitempty"`
	// The country in the billing address.
	Country *string `json:"country,omitempty"`
	// The two-letter code for the country in the billing address.
	CountryCode *string `json:"country_code,omitempty"`
	// The longitude for the billing address location.
	Longitude *float32 `json:"longitude,omitempty"`
	// The latitude for the billing address location.
	Latitude *float32 `json:"latitude,omitempty"`
	// The phone number for the billing address
	Phone *string `json:"phone,omitempty"`
	// The company associated with the billing address.
	Company *string `json:"company,omitempty"`
}

// NewBillingAddress1 instantiates a new BillingAddress1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingAddress1() *BillingAddress1 {
	this := BillingAddress1{}
	return &this
}

// NewBillingAddress1WithDefaults instantiates a new BillingAddress1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingAddress1WithDefaults() *BillingAddress1 {
	this := BillingAddress1{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BillingAddress1) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAddress1) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HaveName returns a boolean if a field has been set.
func (o *BillingAddress1) HaveName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BillingAddress1) SetName(v string) {
	o.Name = &v
}

// GetAddress1 returns the Address1 field value if set, zero value otherwise.
func (o *BillingAddress1) GetAddress1() string {
	if o == nil || IsNil(o.Address1) {
		var ret string
		return ret
	}
	return *o.Address1
}

// GetAddress1Ok returns a tuple with the Address1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAddress1) GetAddress1Ok() (*string, bool) {
	if o == nil || IsNil(o.Address1) {
		return nil, false
	}
	return o.Address1, true
}

// HaveAddress1 returns a boolean if a field has been set.
func (o *BillingAddress1) HaveAddress1() bool {
	if o != nil && !IsNil(o.Address1) {
		return true
	}

	return false
}

// SetAddress1 gets a reference to the given string and assigns it to the Address1 field.
func (o *BillingAddress1) SetAddress1(v string) {
	o.Address1 = &v
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise.
func (o *BillingAddress1) GetAddress2() string {
	if o == nil || IsNil(o.Address2) {
		var ret string
		return ret
	}
	return *o.Address2
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAddress1) GetAddress2Ok() (*string, bool) {
	if o == nil || IsNil(o.Address2) {
		return nil, false
	}
	return o.Address2, true
}

// HaveAddress2 returns a boolean if a field has been set.
func (o *BillingAddress1) HaveAddress2() bool {
	if o != nil && !IsNil(o.Address2) {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given string and assigns it to the Address2 field.
func (o *BillingAddress1) SetAddress2(v string) {
	o.Address2 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *BillingAddress1) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAddress1) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HaveCity returns a boolean if a field has been set.
func (o *BillingAddress1) HaveCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *BillingAddress1) SetCity(v string) {
	o.City = &v
}

// GetProvince returns the Province field value if set, zero value otherwise.
func (o *BillingAddress1) GetProvince() string {
	if o == nil || IsNil(o.Province) {
		var ret string
		return ret
	}
	return *o.Province
}

// GetProvinceOk returns a tuple with the Province field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAddress1) GetProvinceOk() (*string, bool) {
	if o == nil || IsNil(o.Province) {
		return nil, false
	}
	return o.Province, true
}

// HaveProvince returns a boolean if a field has been set.
func (o *BillingAddress1) HaveProvince() bool {
	if o != nil && !IsNil(o.Province) {
		return true
	}

	return false
}

// SetProvince gets a reference to the given string and assigns it to the Province field.
func (o *BillingAddress1) SetProvince(v string) {
	o.Province = &v
}

// GetProvinceCode returns the ProvinceCode field value if set, zero value otherwise.
func (o *BillingAddress1) GetProvinceCode() string {
	if o == nil || IsNil(o.ProvinceCode) {
		var ret string
		return ret
	}
	return *o.ProvinceCode
}

// GetProvinceCodeOk returns a tuple with the ProvinceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAddress1) GetProvinceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ProvinceCode) {
		return nil, false
	}
	return o.ProvinceCode, true
}

// HaveProvinceCode returns a boolean if a field has been set.
func (o *BillingAddress1) HaveProvinceCode() bool {
	if o != nil && !IsNil(o.ProvinceCode) {
		return true
	}

	return false
}

// SetProvinceCode gets a reference to the given string and assigns it to the ProvinceCode field.
func (o *BillingAddress1) SetProvinceCode(v string) {
	o.ProvinceCode = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *BillingAddress1) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAddress1) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HavePostalCode returns a boolean if a field has been set.
func (o *BillingAddress1) HavePostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *BillingAddress1) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *BillingAddress1) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAddress1) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HaveCountry returns a boolean if a field has been set.
func (o *BillingAddress1) HaveCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *BillingAddress1) SetCountry(v string) {
	o.Country = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *BillingAddress1) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAddress1) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HaveCountryCode returns a boolean if a field has been set.
func (o *BillingAddress1) HaveCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *BillingAddress1) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *BillingAddress1) GetLongitude() float32 {
	if o == nil || IsNil(o.Longitude) {
		var ret float32
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAddress1) GetLongitudeOk() (*float32, bool) {
	if o == nil || IsNil(o.Longitude) {
		return nil, false
	}
	return o.Longitude, true
}

// HaveLongitude returns a boolean if a field has been set.
func (o *BillingAddress1) HaveLongitude() bool {
	if o != nil && !IsNil(o.Longitude) {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given float32 and assigns it to the Longitude field.
func (o *BillingAddress1) SetLongitude(v float32) {
	o.Longitude = &v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *BillingAddress1) GetLatitude() float32 {
	if o == nil || IsNil(o.Latitude) {
		var ret float32
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAddress1) GetLatitudeOk() (*float32, bool) {
	if o == nil || IsNil(o.Latitude) {
		return nil, false
	}
	return o.Latitude, true
}

// HaveLatitude returns a boolean if a field has been set.
func (o *BillingAddress1) HaveLatitude() bool {
	if o != nil && !IsNil(o.Latitude) {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given float32 and assigns it to the Latitude field.
func (o *BillingAddress1) SetLatitude(v float32) {
	o.Latitude = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *BillingAddress1) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAddress1) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HavePhone returns a boolean if a field has been set.
func (o *BillingAddress1) HavePhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *BillingAddress1) SetPhone(v string) {
	o.Phone = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *BillingAddress1) GetCompany() string {
	if o == nil || IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAddress1) GetCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HaveCompany returns a boolean if a field has been set.
func (o *BillingAddress1) HaveCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *BillingAddress1) SetCompany(v string) {
	o.Company = &v
}

func (o BillingAddress1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingAddress1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Address1) {
		toSerialize["address1"] = o.Address1
	}
	if !IsNil(o.Address2) {
		toSerialize["address2"] = o.Address2
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.Province) {
		toSerialize["province"] = o.Province
	}
	if !IsNil(o.ProvinceCode) {
		toSerialize["province_code"] = o.ProvinceCode
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postal_code"] = o.PostalCode
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.CountryCode) {
		toSerialize["country_code"] = o.CountryCode
	}
	if !IsNil(o.Longitude) {
		toSerialize["longitude"] = o.Longitude
	}
	if !IsNil(o.Latitude) {
		toSerialize["latitude"] = o.Latitude
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	return toSerialize, nil
}

type NullableBillingAddress1 struct {
	value *BillingAddress1
	isSet bool
}

func (v NullableBillingAddress1) Get() *BillingAddress1 {
	return v.value
}

func (v *NullableBillingAddress1) Set(val *BillingAddress1) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingAddress1) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingAddress1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingAddress1(val *BillingAddress1) *NullableBillingAddress1 {
	return &NullableBillingAddress1{value: val, isSet: true}
}

func (v NullableBillingAddress1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingAddress1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


