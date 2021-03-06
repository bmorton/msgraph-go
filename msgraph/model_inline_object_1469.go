/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
)

// InlineObject1469 struct for InlineObject1469
type InlineObject1469 struct {
	Settlement AnyOfobject `json:"settlement,omitempty"`
	Maturity AnyOfobject `json:"maturity,omitempty"`
	Rate AnyOfobject `json:"rate,omitempty"`
	Yld AnyOfobject `json:"yld,omitempty"`
	Redemption AnyOfobject `json:"redemption,omitempty"`
	Frequency AnyOfobject `json:"frequency,omitempty"`
	Basis AnyOfobject `json:"basis,omitempty"`
}

// NewInlineObject1469 instantiates a new InlineObject1469 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1469() *InlineObject1469 {
	this := InlineObject1469{}
	return &this
}

// NewInlineObject1469WithDefaults instantiates a new InlineObject1469 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1469WithDefaults() *InlineObject1469 {
	this := InlineObject1469{}
	return &this
}

// GetSettlement returns the Settlement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1469) GetSettlement() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Settlement
}

// GetSettlementOk returns a tuple with the Settlement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1469) GetSettlementOk() (*AnyOfobject, bool) {
	if o == nil || o.Settlement == nil {
		return nil, false
	}
	return &o.Settlement, true
}

// HasSettlement returns a boolean if a field has been set.
func (o *InlineObject1469) HasSettlement() bool {
	if o != nil && o.Settlement != nil {
		return true
	}

	return false
}

// SetSettlement gets a reference to the given AnyOfobject and assigns it to the Settlement field.
func (o *InlineObject1469) SetSettlement(v AnyOfobject) {
	o.Settlement = v
}

// GetMaturity returns the Maturity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1469) GetMaturity() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Maturity
}

// GetMaturityOk returns a tuple with the Maturity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1469) GetMaturityOk() (*AnyOfobject, bool) {
	if o == nil || o.Maturity == nil {
		return nil, false
	}
	return &o.Maturity, true
}

// HasMaturity returns a boolean if a field has been set.
func (o *InlineObject1469) HasMaturity() bool {
	if o != nil && o.Maturity != nil {
		return true
	}

	return false
}

// SetMaturity gets a reference to the given AnyOfobject and assigns it to the Maturity field.
func (o *InlineObject1469) SetMaturity(v AnyOfobject) {
	o.Maturity = v
}

// GetRate returns the Rate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1469) GetRate() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1469) GetRateOk() (*AnyOfobject, bool) {
	if o == nil || o.Rate == nil {
		return nil, false
	}
	return &o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *InlineObject1469) HasRate() bool {
	if o != nil && o.Rate != nil {
		return true
	}

	return false
}

// SetRate gets a reference to the given AnyOfobject and assigns it to the Rate field.
func (o *InlineObject1469) SetRate(v AnyOfobject) {
	o.Rate = v
}

// GetYld returns the Yld field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1469) GetYld() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Yld
}

// GetYldOk returns a tuple with the Yld field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1469) GetYldOk() (*AnyOfobject, bool) {
	if o == nil || o.Yld == nil {
		return nil, false
	}
	return &o.Yld, true
}

// HasYld returns a boolean if a field has been set.
func (o *InlineObject1469) HasYld() bool {
	if o != nil && o.Yld != nil {
		return true
	}

	return false
}

// SetYld gets a reference to the given AnyOfobject and assigns it to the Yld field.
func (o *InlineObject1469) SetYld(v AnyOfobject) {
	o.Yld = v
}

// GetRedemption returns the Redemption field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1469) GetRedemption() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Redemption
}

// GetRedemptionOk returns a tuple with the Redemption field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1469) GetRedemptionOk() (*AnyOfobject, bool) {
	if o == nil || o.Redemption == nil {
		return nil, false
	}
	return &o.Redemption, true
}

// HasRedemption returns a boolean if a field has been set.
func (o *InlineObject1469) HasRedemption() bool {
	if o != nil && o.Redemption != nil {
		return true
	}

	return false
}

// SetRedemption gets a reference to the given AnyOfobject and assigns it to the Redemption field.
func (o *InlineObject1469) SetRedemption(v AnyOfobject) {
	o.Redemption = v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1469) GetFrequency() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1469) GetFrequencyOk() (*AnyOfobject, bool) {
	if o == nil || o.Frequency == nil {
		return nil, false
	}
	return &o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *InlineObject1469) HasFrequency() bool {
	if o != nil && o.Frequency != nil {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given AnyOfobject and assigns it to the Frequency field.
func (o *InlineObject1469) SetFrequency(v AnyOfobject) {
	o.Frequency = v
}

// GetBasis returns the Basis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1469) GetBasis() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Basis
}

// GetBasisOk returns a tuple with the Basis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1469) GetBasisOk() (*AnyOfobject, bool) {
	if o == nil || o.Basis == nil {
		return nil, false
	}
	return &o.Basis, true
}

// HasBasis returns a boolean if a field has been set.
func (o *InlineObject1469) HasBasis() bool {
	if o != nil && o.Basis != nil {
		return true
	}

	return false
}

// SetBasis gets a reference to the given AnyOfobject and assigns it to the Basis field.
func (o *InlineObject1469) SetBasis(v AnyOfobject) {
	o.Basis = v
}

func (o InlineObject1469) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Settlement != nil {
		toSerialize["settlement"] = o.Settlement
	}
	if o.Maturity != nil {
		toSerialize["maturity"] = o.Maturity
	}
	if o.Rate != nil {
		toSerialize["rate"] = o.Rate
	}
	if o.Yld != nil {
		toSerialize["yld"] = o.Yld
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

type NullableInlineObject1469 struct {
	value *InlineObject1469
	isSet bool
}

func (v NullableInlineObject1469) Get() *InlineObject1469 {
	return v.value
}

func (v *NullableInlineObject1469) Set(val *InlineObject1469) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1469) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1469) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1469(val *InlineObject1469) *NullableInlineObject1469 {
	return &NullableInlineObject1469{value: val, isSet: true}
}

func (v NullableInlineObject1469) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1469) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


