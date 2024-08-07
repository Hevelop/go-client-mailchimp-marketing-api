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

// checks if the GetReportsIdEcommerceProductActivity200ResponseProductsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetReportsIdEcommerceProductActivity200ResponseProductsInner{}

// GetReportsIdEcommerceProductActivity200ResponseProductsInner struct for GetReportsIdEcommerceProductActivity200ResponseProductsInner
type GetReportsIdEcommerceProductActivity200ResponseProductsInner struct {
	Title *string `json:"title,omitempty"`
	Sku *string `json:"sku,omitempty"`
	ImageUrl *string `json:"image_url,omitempty"`
	TotalRevenue *float32 `json:"total_revenue,omitempty"`
	TotalPurchased *float32 `json:"total_purchased,omitempty"`
	CurrencyCode *string `json:"currency_code,omitempty"`
	RecommendationTotal *int32 `json:"recommendation_total,omitempty"`
	RecommendationPurchased *int32 `json:"recommendation_purchased,omitempty"`
}

// NewGetReportsIdEcommerceProductActivity200ResponseProductsInner instantiates a new GetReportsIdEcommerceProductActivity200ResponseProductsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetReportsIdEcommerceProductActivity200ResponseProductsInner() *GetReportsIdEcommerceProductActivity200ResponseProductsInner {
	this := GetReportsIdEcommerceProductActivity200ResponseProductsInner{}
	return &this
}

// NewGetReportsIdEcommerceProductActivity200ResponseProductsInnerWithDefaults instantiates a new GetReportsIdEcommerceProductActivity200ResponseProductsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetReportsIdEcommerceProductActivity200ResponseProductsInnerWithDefaults() *GetReportsIdEcommerceProductActivity200ResponseProductsInner {
	this := GetReportsIdEcommerceProductActivity200ResponseProductsInner{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *GetReportsIdEcommerceProductActivity200ResponseProductsInner) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReportsIdEcommerceProductActivity200ResponseProductsInner) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *GetReportsIdEcommerceProductActivity200ResponseProductsInner) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *GetReportsIdEcommerceProductActivity200ResponseProductsInner) SetTitle(v string) {
	o.Title = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *GetReportsIdEcommerceProductActivity200ResponseProductsInner) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReportsIdEcommerceProductActivity200ResponseProductsInner) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *GetReportsIdEcommerceProductActivity200ResponseProductsInner) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *GetReportsIdEcommerceProductActivity200ResponseProductsInner) SetSku(v string) {
	o.Sku = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *GetReportsIdEcommerceProductActivity200ResponseProductsInner) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReportsIdEcommerceProductActivity200ResponseProductsInner) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *GetReportsIdEcommerceProductActivity200ResponseProductsInner) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *GetReportsIdEcommerceProductActivity200ResponseProductsInner) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetTotalRevenue returns the TotalRevenue field value if set, zero value otherwise.
func (o *GetReportsIdEcommerceProductActivity200ResponseProductsInner) GetTotalRevenue() float32 {
	if o == nil || IsNil(o.TotalRevenue) {
		var ret float32
		return ret
	}
	return *o.TotalRevenue
}

// GetTotalRevenueOk returns a tuple with the TotalRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReportsIdEcommerceProductActivity200ResponseProductsInner) GetTotalRevenueOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalRevenue) {
		return nil, false
	}
	return o.TotalRevenue, true
}

// HasTotalRevenue returns a boolean if a field has been set.
func (o *GetReportsIdEcommerceProductActivity200ResponseProductsInner) HasTotalRevenue() bool {
	if o != nil && !IsNil(o.TotalRevenue) {
		return true
	}

	return false
}

// SetTotalRevenue gets a reference to the given float32 and assigns it to the TotalRevenue field.
func (o *GetReportsIdEcommerceProductActivity200ResponseProductsInner) SetTotalRevenue(v float32) {
	o.TotalRevenue = &v
}

// GetTotalPurchased returns the TotalPurchased field value if set, zero value otherwise.
func (o *GetReportsIdEcommerceProductActivity200ResponseProductsInner) GetTotalPurchased() float32 {
	if o == nil || IsNil(o.TotalPurchased) {
		var ret float32
		return ret
	}
	return *o.TotalPurchased
}

// GetTotalPurchasedOk returns a tuple with the TotalPurchased field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReportsIdEcommerceProductActivity200ResponseProductsInner) GetTotalPurchasedOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalPurchased) {
		return nil, false
	}
	return o.TotalPurchased, true
}

// HasTotalPurchased returns a boolean if a field has been set.
func (o *GetReportsIdEcommerceProductActivity200ResponseProductsInner) HasTotalPurchased() bool {
	if o != nil && !IsNil(o.TotalPurchased) {
		return true
	}

	return false
}

// SetTotalPurchased gets a reference to the given float32 and assigns it to the TotalPurchased field.
func (o *GetReportsIdEcommerceProductActivity200ResponseProductsInner) SetTotalPurchased(v float32) {
	o.TotalPurchased = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *GetReportsIdEcommerceProductActivity200ResponseProductsInner) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReportsIdEcommerceProductActivity200ResponseProductsInner) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *GetReportsIdEcommerceProductActivity200ResponseProductsInner) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *GetReportsIdEcommerceProductActivity200ResponseProductsInner) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetRecommendationTotal returns the RecommendationTotal field value if set, zero value otherwise.
func (o *GetReportsIdEcommerceProductActivity200ResponseProductsInner) GetRecommendationTotal() int32 {
	if o == nil || IsNil(o.RecommendationTotal) {
		var ret int32
		return ret
	}
	return *o.RecommendationTotal
}

// GetRecommendationTotalOk returns a tuple with the RecommendationTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReportsIdEcommerceProductActivity200ResponseProductsInner) GetRecommendationTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.RecommendationTotal) {
		return nil, false
	}
	return o.RecommendationTotal, true
}

// HasRecommendationTotal returns a boolean if a field has been set.
func (o *GetReportsIdEcommerceProductActivity200ResponseProductsInner) HasRecommendationTotal() bool {
	if o != nil && !IsNil(o.RecommendationTotal) {
		return true
	}

	return false
}

// SetRecommendationTotal gets a reference to the given int32 and assigns it to the RecommendationTotal field.
func (o *GetReportsIdEcommerceProductActivity200ResponseProductsInner) SetRecommendationTotal(v int32) {
	o.RecommendationTotal = &v
}

// GetRecommendationPurchased returns the RecommendationPurchased field value if set, zero value otherwise.
func (o *GetReportsIdEcommerceProductActivity200ResponseProductsInner) GetRecommendationPurchased() int32 {
	if o == nil || IsNil(o.RecommendationPurchased) {
		var ret int32
		return ret
	}
	return *o.RecommendationPurchased
}

// GetRecommendationPurchasedOk returns a tuple with the RecommendationPurchased field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReportsIdEcommerceProductActivity200ResponseProductsInner) GetRecommendationPurchasedOk() (*int32, bool) {
	if o == nil || IsNil(o.RecommendationPurchased) {
		return nil, false
	}
	return o.RecommendationPurchased, true
}

// HasRecommendationPurchased returns a boolean if a field has been set.
func (o *GetReportsIdEcommerceProductActivity200ResponseProductsInner) HasRecommendationPurchased() bool {
	if o != nil && !IsNil(o.RecommendationPurchased) {
		return true
	}

	return false
}

// SetRecommendationPurchased gets a reference to the given int32 and assigns it to the RecommendationPurchased field.
func (o *GetReportsIdEcommerceProductActivity200ResponseProductsInner) SetRecommendationPurchased(v int32) {
	o.RecommendationPurchased = &v
}

func (o GetReportsIdEcommerceProductActivity200ResponseProductsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetReportsIdEcommerceProductActivity200ResponseProductsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.ImageUrl) {
		toSerialize["image_url"] = o.ImageUrl
	}
	if !IsNil(o.TotalRevenue) {
		toSerialize["total_revenue"] = o.TotalRevenue
	}
	if !IsNil(o.TotalPurchased) {
		toSerialize["total_purchased"] = o.TotalPurchased
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if !IsNil(o.RecommendationTotal) {
		toSerialize["recommendation_total"] = o.RecommendationTotal
	}
	if !IsNil(o.RecommendationPurchased) {
		toSerialize["recommendation_purchased"] = o.RecommendationPurchased
	}
	return toSerialize, nil
}

type NullableGetReportsIdEcommerceProductActivity200ResponseProductsInner struct {
	value *GetReportsIdEcommerceProductActivity200ResponseProductsInner
	isSet bool
}

func (v NullableGetReportsIdEcommerceProductActivity200ResponseProductsInner) Get() *GetReportsIdEcommerceProductActivity200ResponseProductsInner {
	return v.value
}

func (v *NullableGetReportsIdEcommerceProductActivity200ResponseProductsInner) Set(val *GetReportsIdEcommerceProductActivity200ResponseProductsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetReportsIdEcommerceProductActivity200ResponseProductsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetReportsIdEcommerceProductActivity200ResponseProductsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetReportsIdEcommerceProductActivity200ResponseProductsInner(val *GetReportsIdEcommerceProductActivity200ResponseProductsInner) *NullableGetReportsIdEcommerceProductActivity200ResponseProductsInner {
	return &NullableGetReportsIdEcommerceProductActivity200ResponseProductsInner{value: val, isSet: true}
}

func (v NullableGetReportsIdEcommerceProductActivity200ResponseProductsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetReportsIdEcommerceProductActivity200ResponseProductsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


