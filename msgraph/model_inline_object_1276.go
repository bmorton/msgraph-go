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

// InlineObject1276 struct for InlineObject1276
type InlineObject1276 struct {
	Settlement AnyOfobject `json:"settlement,omitempty"`
	Maturity AnyOfobject `json:"maturity,omitempty"`
	Frequency AnyOfobject `json:"frequency,omitempty"`
	Basis AnyOfobject `json:"basis,omitempty"`
}

// NewInlineObject1276 instantiates a new InlineObject1276 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1276() *InlineObject1276 {
	this := InlineObject1276{}
	return &this
}

// NewInlineObject1276WithDefaults instantiates a new InlineObject1276 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1276WithDefaults() *InlineObject1276 {
	this := InlineObject1276{}
	return &this
}

// GetSettlement returns the Settlement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1276) GetSettlement() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Settlement
}

// GetSettlementOk returns a tuple with the Settlement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1276) GetSettlementOk() (*AnyOfobject, bool) {
	if o == nil || o.Settlement == nil {
		return nil, false
	}
	return &o.Settlement, true
}

// HasSettlement returns a boolean if a field has been set.
func (o *InlineObject1276) HasSettlement() bool {
	if o != nil && o.Settlement != nil {
		return true
	}

	return false
}

// SetSettlement gets a reference to the given AnyOfobject and assigns it to the Settlement field.
func (o *InlineObject1276) SetSettlement(v AnyOfobject) {
	o.Settlement = v
}

// GetMaturity returns the Maturity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1276) GetMaturity() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Maturity
}

// GetMaturityOk returns a tuple with the Maturity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1276) GetMaturityOk() (*AnyOfobject, bool) {
	if o == nil || o.Maturity == nil {
		return nil, false
	}
	return &o.Maturity, true
}

// HasMaturity returns a boolean if a field has been set.
func (o *InlineObject1276) HasMaturity() bool {
	if o != nil && o.Maturity != nil {
		return true
	}

	return false
}

// SetMaturity gets a reference to the given AnyOfobject and assigns it to the Maturity field.
func (o *InlineObject1276) SetMaturity(v AnyOfobject) {
	o.Maturity = v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1276) GetFrequency() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1276) GetFrequencyOk() (*AnyOfobject, bool) {
	if o == nil || o.Frequency == nil {
		return nil, false
	}
	return &o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *InlineObject1276) HasFrequency() bool {
	if o != nil && o.Frequency != nil {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given AnyOfobject and assigns it to the Frequency field.
func (o *InlineObject1276) SetFrequency(v AnyOfobject) {
	o.Frequency = v
}

// GetBasis returns the Basis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1276) GetBasis() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Basis
}

// GetBasisOk returns a tuple with the Basis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1276) GetBasisOk() (*AnyOfobject, bool) {
	if o == nil || o.Basis == nil {
		return nil, false
	}
	return &o.Basis, true
}

// HasBasis returns a boolean if a field has been set.
func (o *InlineObject1276) HasBasis() bool {
	if o != nil && o.Basis != nil {
		return true
	}

	return false
}

// SetBasis gets a reference to the given AnyOfobject and assigns it to the Basis field.
func (o *InlineObject1276) SetBasis(v AnyOfobject) {
	o.Basis = v
}

func (o InlineObject1276) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Settlement != nil {
		toSerialize["settlement"] = o.Settlement
	}
	if o.Maturity != nil {
		toSerialize["maturity"] = o.Maturity
	}
	if o.Frequency != nil {
		toSerialize["frequency"] = o.Frequency
	}
	if o.Basis != nil {
		toSerialize["basis"] = o.Basis
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1276 struct {
	value *InlineObject1276
	isSet bool
}

func (v NullableInlineObject1276) Get() *InlineObject1276 {
	return v.value
}

func (v *NullableInlineObject1276) Set(val *InlineObject1276) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1276) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1276) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1276(val *InlineObject1276) *NullableInlineObject1276 {
	return &NullableInlineObject1276{value: val, isSet: true}
}

func (v NullableInlineObject1276) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1276) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

