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

// InlineObject1279 struct for InlineObject1279
type InlineObject1279 struct {
	Settlement AnyOfobject `json:"settlement,omitempty"`
	Maturity AnyOfobject `json:"maturity,omitempty"`
	Frequency AnyOfobject `json:"frequency,omitempty"`
	Basis AnyOfobject `json:"basis,omitempty"`
}

// NewInlineObject1279 instantiates a new InlineObject1279 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1279() *InlineObject1279 {
	this := InlineObject1279{}
	return &this
}

// NewInlineObject1279WithDefaults instantiates a new InlineObject1279 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1279WithDefaults() *InlineObject1279 {
	this := InlineObject1279{}
	return &this
}

// GetSettlement returns the Settlement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1279) GetSettlement() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Settlement
}

// GetSettlementOk returns a tuple with the Settlement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1279) GetSettlementOk() (*AnyOfobject, bool) {
	if o == nil || o.Settlement == nil {
		return nil, false
	}
	return &o.Settlement, true
}

// HasSettlement returns a boolean if a field has been set.
func (o *InlineObject1279) HasSettlement() bool {
	if o != nil && o.Settlement != nil {
		return true
	}

	return false
}

// SetSettlement gets a reference to the given AnyOfobject and assigns it to the Settlement field.
func (o *InlineObject1279) SetSettlement(v AnyOfobject) {
	o.Settlement = v
}

// GetMaturity returns the Maturity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1279) GetMaturity() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Maturity
}

// GetMaturityOk returns a tuple with the Maturity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1279) GetMaturityOk() (*AnyOfobject, bool) {
	if o == nil || o.Maturity == nil {
		return nil, false
	}
	return &o.Maturity, true
}

// HasMaturity returns a boolean if a field has been set.
func (o *InlineObject1279) HasMaturity() bool {
	if o != nil && o.Maturity != nil {
		return true
	}

	return false
}

// SetMaturity gets a reference to the given AnyOfobject and assigns it to the Maturity field.
func (o *InlineObject1279) SetMaturity(v AnyOfobject) {
	o.Maturity = v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1279) GetFrequency() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1279) GetFrequencyOk() (*AnyOfobject, bool) {
	if o == nil || o.Frequency == nil {
		return nil, false
	}
	return &o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *InlineObject1279) HasFrequency() bool {
	if o != nil && o.Frequency != nil {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given AnyOfobject and assigns it to the Frequency field.
func (o *InlineObject1279) SetFrequency(v AnyOfobject) {
	o.Frequency = v
}

// GetBasis returns the Basis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1279) GetBasis() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Basis
}

// GetBasisOk returns a tuple with the Basis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1279) GetBasisOk() (*AnyOfobject, bool) {
	if o == nil || o.Basis == nil {
		return nil, false
	}
	return &o.Basis, true
}

// HasBasis returns a boolean if a field has been set.
func (o *InlineObject1279) HasBasis() bool {
	if o != nil && o.Basis != nil {
		return true
	}

	return false
}

// SetBasis gets a reference to the given AnyOfobject and assigns it to the Basis field.
func (o *InlineObject1279) SetBasis(v AnyOfobject) {
	o.Basis = v
}

func (o InlineObject1279) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject1279 struct {
	value *InlineObject1279
	isSet bool
}

func (v NullableInlineObject1279) Get() *InlineObject1279 {
	return v.value
}

func (v *NullableInlineObject1279) Set(val *InlineObject1279) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1279) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1279) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1279(val *InlineObject1279) *NullableInlineObject1279 {
	return &NullableInlineObject1279{value: val, isSet: true}
}

func (v NullableInlineObject1279) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1279) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

