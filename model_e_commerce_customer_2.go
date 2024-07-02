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

// checks if the ECommerceCustomer2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ECommerceCustomer2{}

// ECommerceCustomer2 Information about a specific customer. Orders for existing customers should include only the `id` parameter in the `customer` object body.
type ECommerceCustomer2 struct {
	// The customer's opt-in status. This value will never overwrite the opt-in status of a pre-existing Mailchimp list member, but will apply to list members that are added through the e-commerce API endpoints. Customers who don't opt in to your Mailchimp list [will be added as `Transactional` members](https://mailchimp.com/developer/marketing/docs/e-commerce/#customers).
	OptInStatus *bool `json:"opt_in_status,omitempty"`
	// The customer's company.
	Company *string `json:"company,omitempty"`
	// The customer's first name.
	FirstName *string `json:"first_name,omitempty"`
	// The customer's last name.
	LastName *string `json:"last_name,omitempty"`
	Address *Address `json:"address,omitempty"`
}

// NewECommerceCustomer2 instantiates a new ECommerceCustomer2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewECommerceCustomer2() *ECommerceCustomer2 {
	this := ECommerceCustomer2{}
	return &this
}

// NewECommerceCustomer2WithDefaults instantiates a new ECommerceCustomer2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewECommerceCustomer2WithDefaults() *ECommerceCustomer2 {
	this := ECommerceCustomer2{}
	return &this
}

// GetOptInStatus returns the OptInStatus field value if set, zero value otherwise.
func (o *ECommerceCustomer2) GetOptInStatus() bool {
	if o == nil || IsNil(o.OptInStatus) {
		var ret bool
		return ret
	}
	return *o.OptInStatus
}

// GetOptInStatusOk returns a tuple with the OptInStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCustomer2) GetOptInStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.OptInStatus) {
		return nil, false
	}
	return o.OptInStatus, true
}

// HasOptInStatus returns a boolean if a field has been set.
func (o *ECommerceCustomer2) HasOptInStatus() bool {
	if o != nil && !IsNil(o.OptInStatus) {
		return true
	}

	return false
}

// SetOptInStatus gets a reference to the given bool and assigns it to the OptInStatus field.
func (o *ECommerceCustomer2) SetOptInStatus(v bool) {
	o.OptInStatus = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *ECommerceCustomer2) GetCompany() string {
	if o == nil || IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCustomer2) GetCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *ECommerceCustomer2) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *ECommerceCustomer2) SetCompany(v string) {
	o.Company = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *ECommerceCustomer2) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCustomer2) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *ECommerceCustomer2) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *ECommerceCustomer2) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *ECommerceCustomer2) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCustomer2) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *ECommerceCustomer2) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *ECommerceCustomer2) SetLastName(v string) {
	o.LastName = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ECommerceCustomer2) GetAddress() Address {
	if o == nil || IsNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCustomer2) GetAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ECommerceCustomer2) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *ECommerceCustomer2) SetAddress(v Address) {
	o.Address = &v
}

func (o ECommerceCustomer2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ECommerceCustomer2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OptInStatus) {
		toSerialize["opt_in_status"] = o.OptInStatus
	}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	return toSerialize, nil
}

type NullableECommerceCustomer2 struct {
	value *ECommerceCustomer2
	isSet bool
}

func (v NullableECommerceCustomer2) Get() *ECommerceCustomer2 {
	return v.value
}

func (v *NullableECommerceCustomer2) Set(val *ECommerceCustomer2) {
	v.value = val
	v.isSet = true
}

func (v NullableECommerceCustomer2) IsSet() bool {
	return v.isSet
}

func (v *NullableECommerceCustomer2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECommerceCustomer2(val *ECommerceCustomer2) *NullableECommerceCustomer2 {
	return &NullableECommerceCustomer2{value: val, isSet: true}
}

func (v NullableECommerceCustomer2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECommerceCustomer2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


