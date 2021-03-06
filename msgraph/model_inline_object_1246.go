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

// InlineObject1246 struct for InlineObject1246
type InlineObject1246 struct {
	Number AnyOfobject `json:"number,omitempty"`
	ShiftAmount AnyOfobject `json:"shiftAmount,omitempty"`
}

// NewInlineObject1246 instantiates a new InlineObject1246 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1246() *InlineObject1246 {
	this := InlineObject1246{}
	return &this
}

// NewInlineObject1246WithDefaults instantiates a new InlineObject1246 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1246WithDefaults() *InlineObject1246 {
	this := InlineObject1246{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1246) GetNumber() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1246) GetNumberOk() (*AnyOfobject, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return &o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *InlineObject1246) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given AnyOfobject and assigns it to the Number field.
func (o *InlineObject1246) SetNumber(v AnyOfobject) {
	o.Number = v
}

// GetShiftAmount returns the ShiftAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1246) GetShiftAmount() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.ShiftAmount
}

// GetShiftAmountOk returns a tuple with the ShiftAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1246) GetShiftAmountOk() (*AnyOfobject, bool) {
	if o == nil || o.ShiftAmount == nil {
		return nil, false
	}
	return &o.ShiftAmount, true
}

// HasShiftAmount returns a boolean if a field has been set.
func (o *InlineObject1246) HasShiftAmount() bool {
	if o != nil && o.ShiftAmount != nil {
		return true
	}

	return false
}

// SetShiftAmount gets a reference to the given AnyOfobject and assigns it to the ShiftAmount field.
func (o *InlineObject1246) SetShiftAmount(v AnyOfobject) {
	o.ShiftAmount = v
}

func (o InlineObject1246) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.ShiftAmount != nil {
		toSerialize["shiftAmount"] = o.ShiftAmount
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1246 struct {
	value *InlineObject1246
	isSet bool
}

func (v NullableInlineObject1246) Get() *InlineObject1246 {
	return v.value
}

func (v *NullableInlineObject1246) Set(val *InlineObject1246) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1246) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1246) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1246(val *InlineObject1246) *NullableInlineObject1246 {
	return &NullableInlineObject1246{value: val, isSet: true}
}

func (v NullableInlineObject1246) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1246) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


