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

// checks if the ECommerceCustomer5 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ECommerceCustomer5{}

// ECommerceCustomer5 Information about a specific customer. Orders for existing customers should include only the `id` parameter in the `customer` object body.
type ECommerceCustomer5 struct {
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

// NewECommerceCustomer5 instantiates a new ECommerceCustomer5 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewECommerceCustomer5() *ECommerceCustomer5 {
	this := ECommerceCustomer5{}
	return &this
}

// NewECommerceCustomer5WithDefaults instantiates a new ECommerceCustomer5 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewECommerceCustomer5WithDefaults() *ECommerceCustomer5 {
	this := ECommerceCustomer5{}
	return &this
}

// GetOptInStatus returns the OptInStatus field value if set, zero value otherwise.
func (o *ECommerceCustomer5) GetOptInStatus() bool {
	if o == nil || IsNil(o.OptInStatus) {
		var ret bool
		return ret
	}
	return *o.OptInStatus
}

// GetOptInStatusOk returns a tuple with the OptInStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCustomer5) GetOptInStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.OptInStatus) {
		return nil, false
	}
	return o.OptInStatus, true
}

// HasOptInStatus returns a boolean if a field has been set.
func (o *ECommerceCustomer5) HasOptInStatus() bool {
	if o != nil && !IsNil(o.OptInStatus) {
		return true
	}

	return false
}

// SetOptInStatus gets a reference to the given bool and assigns it to the OptInStatus field.
func (o *ECommerceCustomer5) SetOptInStatus(v bool) {
	o.OptInStatus = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *ECommerceCustomer5) GetCompany() string {
	if o == nil || IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCustomer5) GetCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *ECommerceCustomer5) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *ECommerceCustomer5) SetCompany(v string) {
	o.Company = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *ECommerceCustomer5) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCustomer5) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *ECommerceCustomer5) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *ECommerceCustomer5) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *ECommerceCustomer5) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCustomer5) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *ECommerceCustomer5) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *ECommerceCustomer5) SetLastName(v string) {
	o.LastName = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ECommerceCustomer5) GetAddress() Address {
	if o == nil || IsNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCustomer5) GetAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ECommerceCustomer5) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *ECommerceCustomer5) SetAddress(v Address) {
	o.Address = &v
}

func (o ECommerceCustomer5) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ECommerceCustomer5) ToMap() (map[string]interface{}, error) {
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

type NullableECommerceCustomer5 struct {
	value *ECommerceCustomer5
	isSet bool
}

func (v NullableECommerceCustomer5) Get() *ECommerceCustomer5 {
	return v.value
}

func (v *NullableECommerceCustomer5) Set(val *ECommerceCustomer5) {
	v.value = val
	v.isSet = true
}

func (v NullableECommerceCustomer5) IsSet() bool {
	return v.isSet
}

func (v *NullableECommerceCustomer5) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECommerceCustomer5(val *ECommerceCustomer5) *NullableECommerceCustomer5 {
	return &NullableECommerceCustomer5{value: val, isSet: true}
}

func (v NullableECommerceCustomer5) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECommerceCustomer5) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


