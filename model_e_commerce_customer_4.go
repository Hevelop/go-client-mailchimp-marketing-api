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
	"bytes"
	"fmt"
)

// checks if the ECommerceCustomer4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ECommerceCustomer4{}

// ECommerceCustomer4 Information about a specific customer. Orders for existing customers should include only the `id` parameter in the `customer` object body.
type ECommerceCustomer4 struct {
	// A unique identifier for the customer. Limited to 50 characters.
	Id string `json:"id"`
	// The customer's email address.
	EmailAddress string `json:"email_address"`
	// The customer's opt-in status. This value will never overwrite the opt-in status of a pre-existing Mailchimp list member, but will apply to list members that are added through the e-commerce API endpoints. Customers who don't opt in to your Mailchimp list [will be added as `Transactional` members](https://mailchimp.com/developer/marketing/docs/e-commerce/#customers).
	OptInStatus bool `json:"opt_in_status"`
	// The customer's company.
	Company *string `json:"company,omitempty"`
	// The customer's first name.
	FirstName *string `json:"first_name,omitempty"`
	// The customer's last name.
	LastName *string `json:"last_name,omitempty"`
	Address *Address `json:"address,omitempty"`
}

type _ECommerceCustomer4 ECommerceCustomer4

// NewECommerceCustomer4 instantiates a new ECommerceCustomer4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewECommerceCustomer4(id string, emailAddress string, optInStatus bool) *ECommerceCustomer4 {
	this := ECommerceCustomer4{}
	this.Id = id
	this.EmailAddress = emailAddress
	this.OptInStatus = optInStatus
	return &this
}

// NewECommerceCustomer4WithDefaults instantiates a new ECommerceCustomer4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewECommerceCustomer4WithDefaults() *ECommerceCustomer4 {
	this := ECommerceCustomer4{}
	return &this
}

// GetId returns the Id field value
func (o *ECommerceCustomer4) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ECommerceCustomer4) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ECommerceCustomer4) SetId(v string) {
	o.Id = v
}

// GetEmailAddress returns the EmailAddress field value
func (o *ECommerceCustomer4) GetEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value
// and a boolean to check if the value has been set.
func (o *ECommerceCustomer4) GetEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailAddress, true
}

// SetEmailAddress sets field value
func (o *ECommerceCustomer4) SetEmailAddress(v string) {
	o.EmailAddress = v
}

// GetOptInStatus returns the OptInStatus field value
func (o *ECommerceCustomer4) GetOptInStatus() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.OptInStatus
}

// GetOptInStatusOk returns a tuple with the OptInStatus field value
// and a boolean to check if the value has been set.
func (o *ECommerceCustomer4) GetOptInStatusOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OptInStatus, true
}

// SetOptInStatus sets field value
func (o *ECommerceCustomer4) SetOptInStatus(v bool) {
	o.OptInStatus = v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *ECommerceCustomer4) GetCompany() string {
	if o == nil || IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCustomer4) GetCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *ECommerceCustomer4) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *ECommerceCustomer4) SetCompany(v string) {
	o.Company = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *ECommerceCustomer4) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCustomer4) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *ECommerceCustomer4) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *ECommerceCustomer4) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *ECommerceCustomer4) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCustomer4) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *ECommerceCustomer4) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *ECommerceCustomer4) SetLastName(v string) {
	o.LastName = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ECommerceCustomer4) GetAddress() Address {
	if o == nil || IsNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCustomer4) GetAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ECommerceCustomer4) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *ECommerceCustomer4) SetAddress(v Address) {
	o.Address = &v
}

func (o ECommerceCustomer4) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ECommerceCustomer4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["email_address"] = o.EmailAddress
	toSerialize["opt_in_status"] = o.OptInStatus
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

func (o *ECommerceCustomer4) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"email_address",
		"opt_in_status",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varECommerceCustomer4 := _ECommerceCustomer4{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varECommerceCustomer4)

	if err != nil {
		return err
	}

	*o = ECommerceCustomer4(varECommerceCustomer4)

	return err
}

type NullableECommerceCustomer4 struct {
	value *ECommerceCustomer4
	isSet bool
}

func (v NullableECommerceCustomer4) Get() *ECommerceCustomer4 {
	return v.value
}

func (v *NullableECommerceCustomer4) Set(val *ECommerceCustomer4) {
	v.value = val
	v.isSet = true
}

func (v NullableECommerceCustomer4) IsSet() bool {
	return v.isSet
}

func (v *NullableECommerceCustomer4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECommerceCustomer4(val *ECommerceCustomer4) *NullableECommerceCustomer4 {
	return &NullableECommerceCustomer4{value: val, isSet: true}
}

func (v NullableECommerceCustomer4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECommerceCustomer4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

