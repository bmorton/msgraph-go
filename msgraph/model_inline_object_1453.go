/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// InlineObject1453 struct for InlineObject1453
type InlineObject1453 struct {
	Settlement AnyOfobject `json:"settlement,omitempty"`
	Maturity AnyOfobject `json:"maturity,omitempty"`
	Issue AnyOfobject `json:"issue,omitempty"`
	FirstCoupon AnyOfobject `json:"firstCoupon,omitempty"`
	Rate AnyOfobject `json:"rate,omitempty"`
	Pr AnyOfobject `json:"pr,omitempty"`
	Redemption AnyOfobject `json:"redemption,omitempty"`
	Frequency AnyOfobject `json:"frequency,omitempty"`
	Basis AnyOfobject `json:"basis,omitempty"`
}

// NewInlineObject1453 instantiates a new InlineObject1453 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1453() *InlineObject1453 {
	this := InlineObject1453{}
	return &this
}

// NewInlineObject1453WithDefaults instantiates a new InlineObject1453 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1453WithDefaults() *InlineObject1453 {
	this := InlineObject1453{}
	return &this
}

// GetSettlement returns the Settlement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1453) GetSettlement() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Settlement
}

// GetSettlementOk returns a tuple with the Settlement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1453) GetSettlementOk() (*AnyOfobject, bool) {
	if o == nil || o.Settlement == nil {
		return nil, false
	}
	return &o.Settlement, true
}

// HasSettlement returns a boolean if a field has been set.
func (o *InlineObject1453) HasSettlement() bool {
	if o != nil && o.Settlement != nil {
		return true
	}

	return false
}

// SetSettlement gets a reference to the given AnyOfobject and assigns it to the Settlement field.
func (o *InlineObject1453) SetSettlement(v AnyOfobject) {
	o.Settlement = v
}

// GetMaturity returns the Maturity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1453) GetMaturity() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Maturity
}

// GetMaturityOk returns a tuple with the Maturity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1453) GetMaturityOk() (*AnyOfobject, bool) {
	if o == nil || o.Maturity == nil {
		return nil, false
	}
	return &o.Maturity, true
}

// HasMaturity returns a boolean if a field has been set.
func (o *InlineObject1453) HasMaturity() bool {
	if o != nil && o.Maturity != nil {
		return true
	}

	return false
}

// SetMaturity gets a reference to the given AnyOfobject and assigns it to the Maturity field.
func (o *InlineObject1453) SetMaturity(v AnyOfobject) {
	o.Maturity = v
}

// GetIssue returns the Issue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1453) GetIssue() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Issue
}

// GetIssueOk returns a tuple with the Issue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1453) GetIssueOk() (*AnyOfobject, bool) {
	if o == nil || o.Issue == nil {
		return nil, false
	}
	return &o.Issue, true
}

// HasIssue returns a boolean if a field has been set.
func (o *InlineObject1453) HasIssue() bool {
	if o != nil && o.Issue != nil {
		return true
	}

	return false
}

// SetIssue gets a reference to the given AnyOfobject and assigns it to the Issue field.
func (o *InlineObject1453) SetIssue(v AnyOfobject) {
	o.Issue = v
}

// GetFirstCoupon returns the FirstCoupon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1453) GetFirstCoupon() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.FirstCoupon
}

// GetFirstCouponOk returns a tuple with the FirstCoupon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1453) GetFirstCouponOk() (*AnyOfobject, bool) {
	if o == nil || o.FirstCoupon == nil {
		return nil, false
	}
	return &o.FirstCoupon, true
}

// HasFirstCoupon returns a boolean if a field has been set.
func (o *InlineObject1453) HasFirstCoupon() bool {
	if o != nil && o.FirstCoupon != nil {
		return true
	}

	return false
}

// SetFirstCoupon gets a reference to the given AnyOfobject and assigns it to the FirstCoupon field.
func (o *InlineObject1453) SetFirstCoupon(v AnyOfobject) {
	o.FirstCoupon = v
}

// GetRate returns the Rate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1453) GetRate() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1453) GetRateOk() (*AnyOfobject, bool) {
	if o == nil || o.Rate == nil {
		return nil, false
	}
	return &o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *InlineObject1453) HasRate() bool {
	if o != nil && o.Rate != nil {
		return true
	}

	return false
}

// SetRate gets a reference to the given AnyOfobject and assigns it to the Rate field.
func (o *InlineObject1453) SetRate(v AnyOfobject) {
	o.Rate = v
}

// GetPr returns the Pr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1453) GetPr() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Pr
}

// GetPrOk returns a tuple with the Pr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1453) GetPrOk() (*AnyOfobject, bool) {
	if o == nil || o.Pr == nil {
		return nil, false
	}
	return &o.Pr, true
}

// HasPr returns a boolean if a field has been set.
func (o *InlineObject1453) HasPr() bool {
	if o != nil && o.Pr != nil {
		return true
	}

	return false
}

// SetPr gets a reference to the given AnyOfobject and assigns it to the Pr field.
func (o *InlineObject1453) SetPr(v AnyOfobject) {
	o.Pr = v
}

// GetRedemption returns the Redemption field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1453) GetRedemption() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Redemption
}

// GetRedemptionOk returns a tuple with the Redemption field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1453) GetRedemptionOk() (*AnyOfobject, bool) {
	if o == nil || o.Redemption == nil {
		return nil, false
	}
	return &o.Redemption, true
}

// HasRedemption returns a boolean if a field has been set.
func (o *InlineObject1453) HasRedemption() bool {
	if o != nil && o.Redemption != nil {
		return true
	}

	return false
}

// SetRedemption gets a reference to the given AnyOfobject and assigns it to the Redemption field.
func (o *InlineObject1453) SetRedemption(v AnyOfobject) {
	o.Redemption = v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1453) GetFrequency() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1453) GetFrequencyOk() (*AnyOfobject, bool) {
	if o == nil || o.Frequency == nil {
		return nil, false
	}
	return &o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *InlineObject1453) HasFrequency() bool {
	if o != nil && o.Frequency != nil {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given AnyOfobject and assigns it to the Frequency field.
func (o *InlineObject1453) SetFrequency(v AnyOfobject) {
	o.Frequency = v
}

// GetBasis returns the Basis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1453) GetBasis() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Basis
}

// GetBasisOk returns a tuple with the Basis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1453) GetBasisOk() (*AnyOfobject, bool) {
	if o == nil || o.Basis == nil {
		return nil, false
	}
	return &o.Basis, true
}

// HasBasis returns a boolean if a field has been set.
func (o *InlineObject1453) HasBasis() bool {
	if o != nil && o.Basis != nil {
		return true
	}

	return false
}

// SetBasis gets a reference to the given AnyOfobject and assigns it to the Basis field.
func (o *InlineObject1453) SetBasis(v AnyOfobject) {
	o.Basis = v
}

func (o InlineObject1453) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Settlement != nil {
		toSerialize["settlement"] = o.Settlement
	}
	if o.Maturity != nil {
		toSerialize["maturity"] = o.Maturity
	}
	if o.Issue != nil {
		toSerialize["issue"] = o.Issue
	}
	if o.FirstCoupon != nil {
		toSerialize["firstCoupon"] = o.FirstCoupon
	}
	if o.Rate != nil {
		toSerialize["rate"] = o.Rate
	}
	if o.Pr != nil {
		toSerialize["pr"] = o.Pr
	}
	if o.Redemption != nil {
		toSerialize["redemption"] = o.Redemption
	}
	if o.Frequency != nil {
		toSerialize["frequency"] = o.Frequency
	}
	if o.Basis != nil {
		toSerialize["basis"] = o.Basis
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1453 struct {
	value *InlineObject1453
	isSet bool
}

func (v NullableInlineObject1453) Get() *InlineObject1453 {
	return v.value
}

func (v *NullableInlineObject1453) Set(val *InlineObject1453) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1453) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1453) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1453(val *InlineObject1453) *NullableInlineObject1453 {
	return &NullableInlineObject1453{value: val, isSet: true}
}

func (v NullableInlineObject1453) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1453) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

