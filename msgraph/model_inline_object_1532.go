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

// InlineObject1532 struct for InlineObject1532
type InlineObject1532 struct {
	Settlement AnyOfobject `json:"settlement,omitempty"`
	Maturity AnyOfobject `json:"maturity,omitempty"`
	Pr AnyOfobject `json:"pr,omitempty"`
}

// NewInlineObject1532 instantiates a new InlineObject1532 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1532() *InlineObject1532 {
	this := InlineObject1532{}
	return &this
}

// NewInlineObject1532WithDefaults instantiates a new InlineObject1532 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1532WithDefaults() *InlineObject1532 {
	this := InlineObject1532{}
	return &this
}

// GetSettlement returns the Settlement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1532) GetSettlement() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Settlement
}

// GetSettlementOk returns a tuple with the Settlement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1532) GetSettlementOk() (*AnyOfobject, bool) {
	if o == nil || o.Settlement == nil {
		return nil, false
	}
	return &o.Settlement, true
}

// HasSettlement returns a boolean if a field has been set.
func (o *InlineObject1532) HasSettlement() bool {
	if o != nil && o.Settlement != nil {
		return true
	}

	return false
}

// SetSettlement gets a reference to the given AnyOfobject and assigns it to the Settlement field.
func (o *InlineObject1532) SetSettlement(v AnyOfobject) {
	o.Settlement = v
}

// GetMaturity returns the Maturity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1532) GetMaturity() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Maturity
}

// GetMaturityOk returns a tuple with the Maturity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1532) GetMaturityOk() (*AnyOfobject, bool) {
	if o == nil || o.Maturity == nil {
		return nil, false
	}
	return &o.Maturity, true
}

// HasMaturity returns a boolean if a field has been set.
func (o *InlineObject1532) HasMaturity() bool {
	if o != nil && o.Maturity != nil {
		return true
	}

	return false
}

// SetMaturity gets a reference to the given AnyOfobject and assigns it to the Maturity field.
func (o *InlineObject1532) SetMaturity(v AnyOfobject) {
	o.Maturity = v
}

// GetPr returns the Pr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1532) GetPr() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Pr
}

// GetPrOk returns a tuple with the Pr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1532) GetPrOk() (*AnyOfobject, bool) {
	if o == nil || o.Pr == nil {
		return nil, false
	}
	return &o.Pr, true
}

// HasPr returns a boolean if a field has been set.
func (o *InlineObject1532) HasPr() bool {
	if o != nil && o.Pr != nil {
		return true
	}

	return false
}

// SetPr gets a reference to the given AnyOfobject and assigns it to the Pr field.
func (o *InlineObject1532) SetPr(v AnyOfobject) {
	o.Pr = v
}

func (o InlineObject1532) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Settlement != nil {
		toSerialize["settlement"] = o.Settlement
	}
	if o.Maturity != nil {
		toSerialize["maturity"] = o.Maturity
	}
	if o.Pr != nil {
		toSerialize["pr"] = o.Pr
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1532 struct {
	value *InlineObject1532
	isSet bool
}

func (v NullableInlineObject1532) Get() *InlineObject1532 {
	return v.value
}

func (v *NullableInlineObject1532) Set(val *InlineObject1532) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1532) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1532) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1532(val *InlineObject1532) *NullableInlineObject1532 {
	return &NullableInlineObject1532{value: val, isSet: true}
}

func (v NullableInlineObject1532) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1532) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


