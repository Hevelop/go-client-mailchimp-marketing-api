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

// checks if the ECommerceOrder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ECommerceOrder{}

// ECommerceOrder Information about a specific order.
type ECommerceOrder struct {
	// A unique identifier for the order.
	Id *string `json:"id,omitempty"`
	Customer *ECommerceCustomer `json:"customer,omitempty"`
	// The unique identifier for the store.
	StoreId *string `json:"store_id,omitempty"`
	// A string that uniquely identifies the campaign associated with an order.
	CampaignId *string `json:"campaign_id,omitempty"`
	// The URL for the page where the buyer landed when entering the shop.
	LandingSite *string `json:"landing_site,omitempty"`
	// The order status. Use this parameter to trigger [Order Notifications](https://mailchimp.com/developer/marketing/docs/e-commerce/#order-notifications).
	FinancialStatus *string `json:"financial_status,omitempty"`
	// The fulfillment status for the order. Use this parameter to trigger [Order Notifications](https://mailchimp.com/developer/marketing/docs/e-commerce/#order-notifications).
	FulfillmentStatus *string `json:"fulfillment_status,omitempty"`
	// The three-letter ISO 4217 code for the currency that the store accepts.
	CurrencyCode *string `json:"currency_code,omitempty"`
	// The order total associated with an order.
	OrderTotal *float32 `json:"order_total,omitempty"`
	// The URL for the order.
	OrderUrl *string `json:"order_url,omitempty"`
	// The total amount of the discounts to be applied to the price of the order.
	DiscountTotal *float32 `json:"discount_total,omitempty"`
	// The tax total associated with an order.
	TaxTotal *float32 `json:"tax_total,omitempty"`
	// The shipping total for the order.
	ShippingTotal *float32 `json:"shipping_total,omitempty"`
	// The Mailchimp tracking code for the order. Uses the 'mc_tc' parameter in E-Commerce tracking URLs.
	TrackingCode *string `json:"tracking_code,omitempty"`
	// The date and time the order was processed in ISO 8601 format.
	ProcessedAtForeign *time.Time `json:"processed_at_foreign,omitempty"`
	// The date and time the order was cancelled in ISO 8601 format.
	CancelledAtForeign *time.Time `json:"cancelled_at_foreign,omitempty"`
	// The date and time the order was updated in ISO 8601 format.
	UpdatedAtForeign *time.Time `json:"updated_at_foreign,omitempty"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
	BillingAddress *BillingAddress `json:"billing_address,omitempty"`
	// The promo codes applied on the order
	Promos []PromosInner `json:"promos,omitempty"`
	// An array of the order's line items.
	Lines []ECommerceOrderLineItem `json:"lines,omitempty"`
	Outreach *Outreach `json:"outreach,omitempty"`
	// The tracking number associated with the order.
	TrackingNumber *string `json:"tracking_number,omitempty"`
	// The tracking carrier associated with the order.
	TrackingCarrier *string `json:"tracking_carrier,omitempty"`
	// The tracking URL associated with the order.
	TrackingUrl *string `json:"tracking_url,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewECommerceOrder instantiates a new ECommerceOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewECommerceOrder() *ECommerceOrder {
	this := ECommerceOrder{}
	return &this
}

// NewECommerceOrderWithDefaults instantiates a new ECommerceOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewECommerceOrderWithDefaults() *ECommerceOrder {
	this := ECommerceOrder{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ECommerceOrder) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ECommerceOrder) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ECommerceOrder) SetId(v string) {
	o.Id = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *ECommerceOrder) GetCustomer() ECommerceCustomer {
	if o == nil || IsNil(o.Customer) {
		var ret ECommerceCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder) GetCustomerOk() (*ECommerceCustomer, bool) {
	if o == nil || IsNil(o.Customer) {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *ECommerceOrder) HasCustomer() bool {
	if o != nil && !IsNil(o.Customer) {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given ECommerceCustomer and assigns it to the Customer field.
func (o *ECommerceOrder) SetCustomer(v ECommerceCustomer) {
	o.Customer = &v
}

// GetStoreId returns the StoreId field value if set, zero value otherwise.
func (o *ECommerceOrder) GetStoreId() string {
	if o == nil || IsNil(o.StoreId) {
		var ret string
		return ret
	}
	return *o.StoreId
}

// GetStoreIdOk returns a tuple with the StoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder) GetStoreIdOk() (*string, bool) {
	if o == nil || IsNil(o.StoreId) {
		return nil, false
	}
	return o.StoreId, true
}

// HasStoreId returns a boolean if a field has been set.
func (o *ECommerceOrder) HasStoreId() bool {
	if o != nil && !IsNil(o.StoreId) {
		return true
	}

	return false
}

// SetStoreId gets a reference to the given string and assigns it to the StoreId field.
func (o *ECommerceOrder) SetStoreId(v string) {
	o.StoreId = &v
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *ECommerceOrder) GetCampaignId() string {
	if o == nil || IsNil(o.CampaignId) {
		var ret string
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder) GetCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *ECommerceOrder) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given string and assigns it to the CampaignId field.
func (o *ECommerceOrder) SetCampaignId(v string) {
	o.CampaignId = &v
}

// GetLandingSite returns the LandingSite field value if set, zero value otherwise.
func (o *ECommerceOrder) GetLandingSite() string {
	if o == nil || IsNil(o.LandingSite) {
		var ret string
		return ret
	}
	return *o.LandingSite
}

// GetLandingSiteOk returns a tuple with the LandingSite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder) GetLandingSiteOk() (*string, bool) {
	if o == nil || IsNil(o.LandingSite) {
		return nil, false
	}
	return o.LandingSite, true
}

// HasLandingSite returns a boolean if a field has been set.
func (o *ECommerceOrder) HasLandingSite() bool {
	if o != nil && !IsNil(o.LandingSite) {
		return true
	}

	return false
}

// SetLandingSite gets a reference to the given string and assigns it to the LandingSite field.
func (o *ECommerceOrder) SetLandingSite(v string) {
	o.LandingSite = &v
}

// GetFinancialStatus returns the FinancialStatus field value if set, zero value otherwise.
func (o *ECommerceOrder) GetFinancialStatus() string {
	if o == nil || IsNil(o.FinancialStatus) {
		var ret string
		return ret
	}
	return *o.FinancialStatus
}

// GetFinancialStatusOk returns a tuple with the FinancialStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder) GetFinancialStatusOk() (*string, bool) {
	if o == nil || IsNil(o.FinancialStatus) {
		return nil, false
	}
	return o.FinancialStatus, true
}

// HasFinancialStatus returns a boolean if a field has been set.
func (o *ECommerceOrder) HasFinancialStatus() bool {
	if o != nil && !IsNil(o.FinancialStatus) {
		return true
	}

	return false
}

// SetFinancialStatus gets a reference to the given string and assigns it to the FinancialStatus field.
func (o *ECommerceOrder) SetFinancialStatus(v string) {
	o.FinancialStatus = &v
}

// GetFulfillmentStatus returns the FulfillmentStatus field value if set, zero value otherwise.
func (o *ECommerceOrder) GetFulfillmentStatus() string {
	if o == nil || IsNil(o.FulfillmentStatus) {
		var ret string
		return ret
	}
	return *o.FulfillmentStatus
}

// GetFulfillmentStatusOk returns a tuple with the FulfillmentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder) GetFulfillmentStatusOk() (*string, bool) {
	if o == nil || IsNil(o.FulfillmentStatus) {
		return nil, false
	}
	return o.FulfillmentStatus, true
}

// HasFulfillmentStatus returns a boolean if a field has been set.
func (o *ECommerceOrder) HasFulfillmentStatus() bool {
	if o != nil && !IsNil(o.FulfillmentStatus) {
		return true
	}

	return false
}

// SetFulfillmentStatus gets a reference to the given string and assigns it to the FulfillmentStatus field.
func (o *ECommerceOrder) SetFulfillmentStatus(v string) {
	o.FulfillmentStatus = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *ECommerceOrder) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *ECommerceOrder) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *ECommerceOrder) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetOrderTotal returns the OrderTotal field value if set, zero value otherwise.
func (o *ECommerceOrder) GetOrderTotal() float32 {
	if o == nil || IsNil(o.OrderTotal) {
		var ret float32
		return ret
	}
	return *o.OrderTotal
}

// GetOrderTotalOk returns a tuple with the OrderTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder) GetOrderTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.OrderTotal) {
		return nil, false
	}
	return o.OrderTotal, true
}

// HasOrderTotal returns a boolean if a field has been set.
func (o *ECommerceOrder) HasOrderTotal() bool {
	if o != nil && !IsNil(o.OrderTotal) {
		return true
	}

	return false
}

// SetOrderTotal gets a reference to the given float32 and assigns it to the OrderTotal field.
func (o *ECommerceOrder) SetOrderTotal(v float32) {
	o.OrderTotal = &v
}

// GetOrderUrl returns the OrderUrl field value if set, zero value otherwise.
func (o *ECommerceOrder) GetOrderUrl() string {
	if o == nil || IsNil(o.OrderUrl) {
		var ret string
		return ret
	}
	return *o.OrderUrl
}

// GetOrderUrlOk returns a tuple with the OrderUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder) GetOrderUrlOk() (*string, bool) {
	if o == nil || IsNil(o.OrderUrl) {
		return nil, false
	}
	return o.OrderUrl, true
}

// HasOrderUrl returns a boolean if a field has been set.
func (o *ECommerceOrder) HasOrderUrl() bool {
	if o != nil && !IsNil(o.OrderUrl) {
		return true
	}

	return false
}

// SetOrderUrl gets a reference to the given string and assigns it to the OrderUrl field.
func (o *ECommerceOrder) SetOrderUrl(v string) {
	o.OrderUrl = &v
}

// GetDiscountTotal returns the DiscountTotal field value if set, zero value otherwise.
func (o *ECommerceOrder) GetDiscountTotal() float32 {
	if o == nil || IsNil(o.DiscountTotal) {
		var ret float32
		return ret
	}
	return *o.DiscountTotal
}

// GetDiscountTotalOk returns a tuple with the DiscountTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder) GetDiscountTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.DiscountTotal) {
		return nil, false
	}
	return o.DiscountTotal, true
}

// HasDiscountTotal returns a boolean if a field has been set.
func (o *ECommerceOrder) HasDiscountTotal() bool {
	if o != nil && !IsNil(o.DiscountTotal) {
		return true
	}

	return false
}

// SetDiscountTotal gets a reference to the given float32 and assigns it to the DiscountTotal field.
func (o *ECommerceOrder) SetDiscountTotal(v float32) {
	o.DiscountTotal = &v
}

// GetTaxTotal returns the TaxTotal field value if set, zero value otherwise.
func (o *ECommerceOrder) GetTaxTotal() float32 {
	if o == nil || IsNil(o.TaxTotal) {
		var ret float32
		return ret
	}
	return *o.TaxTotal
}

// GetTaxTotalOk returns a tuple with the TaxTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder) GetTaxTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.TaxTotal) {
		return nil, false
	}
	return o.TaxTotal, true
}

// HasTaxTotal returns a boolean if a field has been set.
func (o *ECommerceOrder) HasTaxTotal() bool {
	if o != nil && !IsNil(o.TaxTotal) {
		return true
	}

	return false
}

// SetTaxTotal gets a reference to the given float32 and assigns it to the TaxTotal field.
func (o *ECommerceOrder) SetTaxTotal(v float32) {
	o.TaxTotal = &v
}

// GetShippingTotal returns the ShippingTotal field value if set, zero value otherwise.
func (o *ECommerceOrder) GetShippingTotal() float32 {
	if o == nil || IsNil(o.ShippingTotal) {
		var ret float32
		return ret
	}
	return *o.ShippingTotal
}

// GetShippingTotalOk returns a tuple with the ShippingTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder) GetShippingTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.ShippingTotal) {
		return nil, false
	}
	return o.ShippingTotal, true
}

// HasShippingTotal returns a boolean if a field has been set.
func (o *ECommerceOrder) HasShippingTotal() bool {
	if o != nil && !IsNil(o.ShippingTotal) {
		return true
	}

	return false
}

// SetShippingTotal gets a reference to the given float32 and assigns it to the ShippingTotal field.
func (o *ECommerceOrder) SetShippingTotal(v float32) {
	o.ShippingTotal = &v
}

// GetTrackingCode returns the TrackingCode field value if set, zero value otherwise.
func (o *ECommerceOrder) GetTrackingCode() string {
	if o == nil || IsNil(o.TrackingCode) {
		var ret string
		return ret
	}
	return *o.TrackingCode
}

// GetTrackingCodeOk returns a tuple with the TrackingCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder) GetTrackingCodeOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingCode) {
		return nil, false
	}
	return o.TrackingCode, true
}

// HasTrackingCode returns a boolean if a field has been set.
func (o *ECommerceOrder) HasTrackingCode() bool {
	if o != nil && !IsNil(o.TrackingCode) {
		return true
	}

	return false
}

// SetTrackingCode gets a reference to the given string and assigns it to the TrackingCode field.
func (o *ECommerceOrder) SetTrackingCode(v string) {
	o.TrackingCode = &v
}

// GetProcessedAtForeign returns the ProcessedAtForeign field value if set, zero value otherwise.
func (o *ECommerceOrder) GetProcessedAtForeign() time.Time {
	if o == nil || IsNil(o.ProcessedAtForeign) {
		var ret time.Time
		return ret
	}
	return *o.ProcessedAtForeign
}

// GetProcessedAtForeignOk returns a tuple with the ProcessedAtForeign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder) GetProcessedAtForeignOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ProcessedAtForeign) {
		return nil, false
	}
	return o.ProcessedAtForeign, true
}

// HasProcessedAtForeign returns a boolean if a field has been set.
func (o *ECommerceOrder) HasProcessedAtForeign() bool {
	if o != nil && !IsNil(o.ProcessedAtForeign) {
		return true
	}

	return false
}

// SetProcessedAtForeign gets a reference to the given time.Time and assigns it to the ProcessedAtForeign field.
func (o *ECommerceOrder) SetProcessedAtForeign(v time.Time) {
	o.ProcessedAtForeign = &v
}

// GetCancelledAtForeign returns the CancelledAtForeign field value if set, zero value otherwise.
func (o *ECommerceOrder) GetCancelledAtForeign() time.Time {
	if o == nil || IsNil(o.CancelledAtForeign) {
		var ret time.Time
		return ret
	}
	return *o.CancelledAtForeign
}

// GetCancelledAtForeignOk returns a tuple with the CancelledAtForeign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder) GetCancelledAtForeignOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CancelledAtForeign) {
		return nil, false
	}
	return o.CancelledAtForeign, true
}

// HasCancelledAtForeign returns a boolean if a field has been set.
func (o *ECommerceOrder) HasCancelledAtForeign() bool {
	if o != nil && !IsNil(o.CancelledAtForeign) {
		return true
	}

	return false
}

// SetCancelledAtForeign gets a reference to the given time.Time and assigns it to the CancelledAtForeign field.
func (o *ECommerceOrder) SetCancelledAtForeign(v time.Time) {
	o.CancelledAtForeign = &v
}

// GetUpdatedAtForeign returns the UpdatedAtForeign field value if set, zero value otherwise.
func (o *ECommerceOrder) GetUpdatedAtForeign() time.Time {
	if o == nil || IsNil(o.UpdatedAtForeign) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAtForeign
}

// GetUpdatedAtForeignOk returns a tuple with the UpdatedAtForeign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder) GetUpdatedAtForeignOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAtForeign) {
		return nil, false
	}
	return o.UpdatedAtForeign, true
}

// HasUpdatedAtForeign returns a boolean if a field has been set.
func (o *ECommerceOrder) HasUpdatedAtForeign() bool {
	if o != nil && !IsNil(o.UpdatedAtForeign) {
		return true
	}

	return false
}

// SetUpdatedAtForeign gets a reference to the given time.Time and assigns it to the UpdatedAtForeign field.
func (o *ECommerceOrder) SetUpdatedAtForeign(v time.Time) {
	o.UpdatedAtForeign = &v
}

// GetShippingAddress returns the ShippingAddress field value if set, zero value otherwise.
func (o *ECommerceOrder) GetShippingAddress() ShippingAddress {
	if o == nil || IsNil(o.ShippingAddress) {
		var ret ShippingAddress
		return ret
	}
	return *o.ShippingAddress
}

// GetShippingAddressOk returns a tuple with the ShippingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder) GetShippingAddressOk() (*ShippingAddress, bool) {
	if o == nil || IsNil(o.ShippingAddress) {
		return nil, false
	}
	return o.ShippingAddress, true
}

// HasShippingAddress returns a boolean if a field has been set.
func (o *ECommerceOrder) HasShippingAddress() bool {
	if o != nil && !IsNil(o.ShippingAddress) {
		return true
	}

	return false
}

// SetShippingAddress gets a reference to the given ShippingAddress and assigns it to the ShippingAddress field.
func (o *ECommerceOrder) SetShippingAddress(v ShippingAddress) {
	o.ShippingAddress = &v
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *ECommerceOrder) GetBillingAddress() BillingAddress {
	if o == nil || IsNil(o.BillingAddress) {
		var ret BillingAddress
		return ret
	}
	return *o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder) GetBillingAddressOk() (*BillingAddress, bool) {
	if o == nil || IsNil(o.BillingAddress) {
		return nil, false
	}
	return o.BillingAddress, true
}

// HasBillingAddress returns a boolean if a field has been set.
func (o *ECommerceOrder) HasBillingAddress() bool {
	if o != nil && !IsNil(o.BillingAddress) {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given BillingAddress and assigns it to the BillingAddress field.
func (o *ECommerceOrder) SetBillingAddress(v BillingAddress) {
	o.BillingAddress = &v
}

// GetPromos returns the Promos field value if set, zero value otherwise.
func (o *ECommerceOrder) GetPromos() []PromosInner {
	if o == nil || IsNil(o.Promos) {
		var ret []PromosInner
		return ret
	}
	return o.Promos
}

// GetPromosOk returns a tuple with the Promos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder) GetPromosOk() ([]PromosInner, bool) {
	if o == nil || IsNil(o.Promos) {
		return nil, false
	}
	return o.Promos, true
}

// HasPromos returns a boolean if a field has been set.
func (o *ECommerceOrder) HasPromos() bool {
	if o != nil && !IsNil(o.Promos) {
		return true
	}

	return false
}

// SetPromos gets a reference to the given []PromosInner and assigns it to the Promos field.
func (o *ECommerceOrder) SetPromos(v []PromosInner) {
	o.Promos = v
}

// GetLines returns the Lines field value if set, zero value otherwise.
func (o *ECommerceOrder) GetLines() []ECommerceOrderLineItem {
	if o == nil || IsNil(o.Lines) {
		var ret []ECommerceOrderLineItem
		return ret
	}
	return o.Lines
}

// GetLinesOk returns a tuple with the Lines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder) GetLinesOk() ([]ECommerceOrderLineItem, bool) {
	if o == nil || IsNil(o.Lines) {
		return nil, false
	}
	return o.Lines, true
}

// HasLines returns a boolean if a field has been set.
func (o *ECommerceOrder) HasLines() bool {
	if o != nil && !IsNil(o.Lines) {
		return true
	}

	return false
}

// SetLines gets a reference to the given []ECommerceOrderLineItem and assigns it to the Lines field.
func (o *ECommerceOrder) SetLines(v []ECommerceOrderLineItem) {
	o.Lines = v
}

// GetOutreach returns the Outreach field value if set, zero value otherwise.
func (o *ECommerceOrder) GetOutreach() Outreach {
	if o == nil || IsNil(o.Outreach) {
		var ret Outreach
		return ret
	}
	return *o.Outreach
}

// GetOutreachOk returns a tuple with the Outreach field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder) GetOutreachOk() (*Outreach, bool) {
	if o == nil || IsNil(o.Outreach) {
		return nil, false
	}
	return o.Outreach, true
}

// HasOutreach returns a boolean if a field has been set.
func (o *ECommerceOrder) HasOutreach() bool {
	if o != nil && !IsNil(o.Outreach) {
		return true
	}

	return false
}

// SetOutreach gets a reference to the given Outreach and assigns it to the Outreach field.
func (o *ECommerceOrder) SetOutreach(v Outreach) {
	o.Outreach = &v
}

// GetTrackingNumber returns the TrackingNumber field value if set, zero value otherwise.
func (o *ECommerceOrder) GetTrackingNumber() string {
	if o == nil || IsNil(o.TrackingNumber) {
		var ret string
		return ret
	}
	return *o.TrackingNumber
}

// GetTrackingNumberOk returns a tuple with the TrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder) GetTrackingNumberOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingNumber) {
		return nil, false
	}
	return o.TrackingNumber, true
}

// HasTrackingNumber returns a boolean if a field has been set.
func (o *ECommerceOrder) HasTrackingNumber() bool {
	if o != nil && !IsNil(o.TrackingNumber) {
		return true
	}

	return false
}

// SetTrackingNumber gets a reference to the given string and assigns it to the TrackingNumber field.
func (o *ECommerceOrder) SetTrackingNumber(v string) {
	o.TrackingNumber = &v
}

// GetTrackingCarrier returns the TrackingCarrier field value if set, zero value otherwise.
func (o *ECommerceOrder) GetTrackingCarrier() string {
	if o == nil || IsNil(o.TrackingCarrier) {
		var ret string
		return ret
	}
	return *o.TrackingCarrier
}

// GetTrackingCarrierOk returns a tuple with the TrackingCarrier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder) GetTrackingCarrierOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingCarrier) {
		return nil, false
	}
	return o.TrackingCarrier, true
}

// HasTrackingCarrier returns a boolean if a field has been set.
func (o *ECommerceOrder) HasTrackingCarrier() bool {
	if o != nil && !IsNil(o.TrackingCarrier) {
		return true
	}

	return false
}

// SetTrackingCarrier gets a reference to the given string and assigns it to the TrackingCarrier field.
func (o *ECommerceOrder) SetTrackingCarrier(v string) {
	o.TrackingCarrier = &v
}

// GetTrackingUrl returns the TrackingUrl field value if set, zero value otherwise.
func (o *ECommerceOrder) GetTrackingUrl() string {
	if o == nil || IsNil(o.TrackingUrl) {
		var ret string
		return ret
	}
	return *o.TrackingUrl
}

// GetTrackingUrlOk returns a tuple with the TrackingUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder) GetTrackingUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingUrl) {
		return nil, false
	}
	return o.TrackingUrl, true
}

// HasTrackingUrl returns a boolean if a field has been set.
func (o *ECommerceOrder) HasTrackingUrl() bool {
	if o != nil && !IsNil(o.TrackingUrl) {
		return true
	}

	return false
}

// SetTrackingUrl gets a reference to the given string and assigns it to the TrackingUrl field.
func (o *ECommerceOrder) SetTrackingUrl(v string) {
	o.TrackingUrl = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ECommerceOrder) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ECommerceOrder) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *ECommerceOrder) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o ECommerceOrder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ECommerceOrder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Customer) {
		toSerialize["customer"] = o.Customer
	}
	if !IsNil(o.StoreId) {
		toSerialize["store_id"] = o.StoreId
	}
	if !IsNil(o.CampaignId) {
		toSerialize["campaign_id"] = o.CampaignId
	}
	if !IsNil(o.LandingSite) {
		toSerialize["landing_site"] = o.LandingSite
	}
	if !IsNil(o.FinancialStatus) {
		toSerialize["financial_status"] = o.FinancialStatus
	}
	if !IsNil(o.FulfillmentStatus) {
		toSerialize["fulfillment_status"] = o.FulfillmentStatus
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if !IsNil(o.OrderTotal) {
		toSerialize["order_total"] = o.OrderTotal
	}
	if !IsNil(o.OrderUrl) {
		toSerialize["order_url"] = o.OrderUrl
	}
	if !IsNil(o.DiscountTotal) {
		toSerialize["discount_total"] = o.DiscountTotal
	}
	if !IsNil(o.TaxTotal) {
		toSerialize["tax_total"] = o.TaxTotal
	}
	if !IsNil(o.ShippingTotal) {
		toSerialize["shipping_total"] = o.ShippingTotal
	}
	if !IsNil(o.TrackingCode) {
		toSerialize["tracking_code"] = o.TrackingCode
	}
	if !IsNil(o.ProcessedAtForeign) {
		toSerialize["processed_at_foreign"] = o.ProcessedAtForeign
	}
	if !IsNil(o.CancelledAtForeign) {
		toSerialize["cancelled_at_foreign"] = o.CancelledAtForeign
	}
	if !IsNil(o.UpdatedAtForeign) {
		toSerialize["updated_at_foreign"] = o.UpdatedAtForeign
	}
	if !IsNil(o.ShippingAddress) {
		toSerialize["shipping_address"] = o.ShippingAddress
	}
	if !IsNil(o.BillingAddress) {
		toSerialize["billing_address"] = o.BillingAddress
	}
	if !IsNil(o.Promos) {
		toSerialize["promos"] = o.Promos
	}
	if !IsNil(o.Lines) {
		toSerialize["lines"] = o.Lines
	}
	if !IsNil(o.Outreach) {
		toSerialize["outreach"] = o.Outreach
	}
	if !IsNil(o.TrackingNumber) {
		toSerialize["tracking_number"] = o.TrackingNumber
	}
	if !IsNil(o.TrackingCarrier) {
		toSerialize["tracking_carrier"] = o.TrackingCarrier
	}
	if !IsNil(o.TrackingUrl) {
		toSerialize["tracking_url"] = o.TrackingUrl
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableECommerceOrder struct {
	value *ECommerceOrder
	isSet bool
}

func (v NullableECommerceOrder) Get() *ECommerceOrder {
	return v.value
}

func (v *NullableECommerceOrder) Set(val *ECommerceOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableECommerceOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableECommerceOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECommerceOrder(val *ECommerceOrder) *NullableECommerceOrder {
	return &NullableECommerceOrder{value: val, isSet: true}
}

func (v NullableECommerceOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECommerceOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


