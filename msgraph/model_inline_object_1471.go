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

// InlineObject1471 struct for InlineObject1471
type InlineObject1471 struct {
	Settlement AnyOfobject `json:"settlement,omitempty"`
	Maturity AnyOfobject `json:"maturity,omitempty"`
	Issue AnyOfobject `json:"issue,omitempty"`
	Rate AnyOfobject `json:"rate,omitempty"`
	Yld AnyOfobject `json:"yld,omitempty"`
	Basis AnyOfobject `json:"basis,omitempty"`
}

// NewInlineObject1471 instantiates a new InlineObject1471 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1471() *InlineObject1471 {
	this := InlineObject1471{}
	return &this
}

// NewInlineObject1471WithDefaults instantiates a new InlineObject1471 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1471WithDefaults() *InlineObject1471 {
	this := InlineObject1471{}
	return &this
}

// GetSettlement returns the Settlement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1471) GetSettlement() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Settlement
}

// GetSettlementOk returns a tuple with the Settlement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1471) GetSettlementOk() (*AnyOfobject, bool) {
	if o == nil || o.Settlement == nil {
		return nil, false
	}
	return &o.Settlement, true
}

// HasSettlement returns a boolean if a field has been set.
func (o *InlineObject1471) HasSettlement() bool {
	if o != nil && o.Settlement != nil {
		return true
	}

	return false
}

// SetSettlement gets a reference to the given AnyOfobject and assigns it to the Settlement field.
func (o *InlineObject1471) SetSettlement(v AnyOfobject) {
	o.Settlement = v
}

// GetMaturity returns the Maturity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1471) GetMaturity() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Maturity
}

// GetMaturityOk returns a tuple with the Maturity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1471) GetMaturityOk() (*AnyOfobject, bool) {
	if o == nil || o.Maturity == nil {
		return nil, false
	}
	return &o.Maturity, true
}

// HasMaturity returns a boolean if a field has been set.
func (o *InlineObject1471) HasMaturity() bool {
	if o != nil && o.Maturity != nil {
		return true
	}

	return false
}

// SetMaturity gets a reference to the given AnyOfobject and assigns it to the Maturity field.
func (o *InlineObject1471) SetMaturity(v AnyOfobject) {
	o.Maturity = v
}

// GetIssue returns the Issue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1471) GetIssue() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Issue
}

// GetIssueOk returns a tuple with the Issue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1471) GetIssueOk() (*AnyOfobject, bool) {
	if o == nil || o.Issue == nil {
		return nil, false
	}
	return &o.Issue, true
}

// HasIssue returns a boolean if a field has been set.
func (o *InlineObject1471) HasIssue() bool {
	if o != nil && o.Issue != nil {
		return true
	}

	return false
}

// SetIssue gets a reference to the given AnyOfobject and assigns it to the Issue field.
func (o *InlineObject1471) SetIssue(v AnyOfobject) {
	o.Issue = v
}

// GetRate returns the Rate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1471) GetRate() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1471) GetRateOk() (*AnyOfobject, bool) {
	if o == nil || o.Rate == nil {
		return nil, false
	}
	return &o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *InlineObject1471) HasRate() bool {
	if o != nil && o.Rate != nil {
		return true
	}

	return false
}

// SetRate gets a reference to the given AnyOfobject and assigns it to the Rate field.
func (o *InlineObject1471) SetRate(v AnyOfobject) {
	o.Rate = v
}

// GetYld returns the Yld field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1471) GetYld() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Yld
}

// GetYldOk returns a tuple with the Yld field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1471) GetYldOk() (*AnyOfobject, bool) {
	if o == nil || o.Yld == nil {
		return nil, false
	}
	return &o.Yld, true
}

// HasYld returns a boolean if a field has been set.
func (o *InlineObject1471) HasYld() bool {
	if o != nil && o.Yld != nil {
		return true
	}

	return false
}

// SetYld gets a reference to the given AnyOfobject and assigns it to the Yld field.
func (o *InlineObject1471) SetYld(v AnyOfobject) {
	o.Yld = v
}

// GetBasis returns the Basis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1471) GetBasis() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Basis
}

// GetBasisOk returns a tuple with the Basis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1471) GetBasisOk() (*AnyOfobject, bool) {
	if o == nil || o.Basis == nil {
		return nil, false
	}
	return &o.Basis, true
}

// HasBasis returns a boolean if a field has been set.
func (o *InlineObject1471) HasBasis() bool {
	if o != nil && o.Basis != nil {
		return true
	}

	return false
}

// SetBasis gets a reference to the given AnyOfobject and assigns it to the Basis field.
func (o *InlineObject1471) SetBasis(v AnyOfobject) {
	o.Basis = v
}

func (o InlineObject1471) MarshalJSON() ([]byte, error) {
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
	if o.Rate != nil {
		toSerialize["rate"] = o.Rate
	}
	if o.Yld != nil {
		toSerialize["yld"] = o.Yld
	}
	if o.Basis != nil {
		toSerialize["basis"] = o.Basis
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1471 struct {
	value *InlineObject1471
	isSet bool
}

func (v NullableInlineObject1471) Get() *InlineObject1471 {
	return v.value
}

func (v *NullableInlineObject1471) Set(val *InlineObject1471) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1471) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1471) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1471(val *InlineObject1471) *NullableInlineObject1471 {
	return &NullableInlineObject1471{value: val, isSet: true}
}

func (v NullableInlineObject1471) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1471) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


