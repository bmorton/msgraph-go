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

// InlineObject1483 struct for InlineObject1483
type InlineObject1483 struct {
	Settlement AnyOfobject `json:"settlement,omitempty"`
	Maturity AnyOfobject `json:"maturity,omitempty"`
	Investment AnyOfobject `json:"investment,omitempty"`
	Discount AnyOfobject `json:"discount,omitempty"`
	Basis AnyOfobject `json:"basis,omitempty"`
}

// NewInlineObject1483 instantiates a new InlineObject1483 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1483() *InlineObject1483 {
	this := InlineObject1483{}
	return &this
}

// NewInlineObject1483WithDefaults instantiates a new InlineObject1483 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1483WithDefaults() *InlineObject1483 {
	this := InlineObject1483{}
	return &this
}

// GetSettlement returns the Settlement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1483) GetSettlement() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Settlement
}

// GetSettlementOk returns a tuple with the Settlement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1483) GetSettlementOk() (*AnyOfobject, bool) {
	if o == nil || o.Settlement == nil {
		return nil, false
	}
	return &o.Settlement, true
}

// HasSettlement returns a boolean if a field has been set.
func (o *InlineObject1483) HasSettlement() bool {
	if o != nil && o.Settlement != nil {
		return true
	}

	return false
}

// SetSettlement gets a reference to the given AnyOfobject and assigns it to the Settlement field.
func (o *InlineObject1483) SetSettlement(v AnyOfobject) {
	o.Settlement = v
}

// GetMaturity returns the Maturity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1483) GetMaturity() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Maturity
}

// GetMaturityOk returns a tuple with the Maturity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1483) GetMaturityOk() (*AnyOfobject, bool) {
	if o == nil || o.Maturity == nil {
		return nil, false
	}
	return &o.Maturity, true
}

// HasMaturity returns a boolean if a field has been set.
func (o *InlineObject1483) HasMaturity() bool {
	if o != nil && o.Maturity != nil {
		return true
	}

	return false
}

// SetMaturity gets a reference to the given AnyOfobject and assigns it to the Maturity field.
func (o *InlineObject1483) SetMaturity(v AnyOfobject) {
	o.Maturity = v
}

// GetInvestment returns the Investment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1483) GetInvestment() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Investment
}

// GetInvestmentOk returns a tuple with the Investment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1483) GetInvestmentOk() (*AnyOfobject, bool) {
	if o == nil || o.Investment == nil {
		return nil, false
	}
	return &o.Investment, true
}

// HasInvestment returns a boolean if a field has been set.
func (o *InlineObject1483) HasInvestment() bool {
	if o != nil && o.Investment != nil {
		return true
	}

	return false
}

// SetInvestment gets a reference to the given AnyOfobject and assigns it to the Investment field.
func (o *InlineObject1483) SetInvestment(v AnyOfobject) {
	o.Investment = v
}

// GetDiscount returns the Discount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1483) GetDiscount() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1483) GetDiscountOk() (*AnyOfobject, bool) {
	if o == nil || o.Discount == nil {
		return nil, false
	}
	return &o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *InlineObject1483) HasDiscount() bool {
	if o != nil && o.Discount != nil {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given AnyOfobject and assigns it to the Discount field.
func (o *InlineObject1483) SetDiscount(v AnyOfobject) {
	o.Discount = v
}

// GetBasis returns the Basis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1483) GetBasis() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Basis
}

// GetBasisOk returns a tuple with the Basis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1483) GetBasisOk() (*AnyOfobject, bool) {
	if o == nil || o.Basis == nil {
		return nil, false
	}
	return &o.Basis, true
}

// HasBasis returns a boolean if a field has been set.
func (o *InlineObject1483) HasBasis() bool {
	if o != nil && o.Basis != nil {
		return true
	}

	return false
}

// SetBasis gets a reference to the given AnyOfobject and assigns it to the Basis field.
func (o *InlineObject1483) SetBasis(v AnyOfobject) {
	o.Basis = v
}

func (o InlineObject1483) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Settlement != nil {
		toSerialize["settlement"] = o.Settlement
	}
	if o.Maturity != nil {
		toSerialize["maturity"] = o.Maturity
	}
	if o.Investment != nil {
		toSerialize["investment"] = o.Investment
	}
	if o.Discount != nil {
		toSerialize["discount"] = o.Discount
	}
	if o.Basis != nil {
		toSerialize["basis"] = o.Basis
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1483 struct {
	value *InlineObject1483
	isSet bool
}

func (v NullableInlineObject1483) Get() *InlineObject1483 {
	return v.value
}

func (v *NullableInlineObject1483) Set(val *InlineObject1483) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1483) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1483) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1483(val *InlineObject1483) *NullableInlineObject1483 {
	return &NullableInlineObject1483{value: val, isSet: true}
}

func (v NullableInlineObject1483) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1483) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


