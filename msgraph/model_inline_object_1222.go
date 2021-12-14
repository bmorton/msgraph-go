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

// InlineObject1222 struct for InlineObject1222
type InlineObject1222 struct {
	XNum AnyOfobject `json:"xNum,omitempty"`
	YNum AnyOfobject `json:"yNum,omitempty"`
}

// NewInlineObject1222 instantiates a new InlineObject1222 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1222() *InlineObject1222 {
	this := InlineObject1222{}
	return &this
}

// NewInlineObject1222WithDefaults instantiates a new InlineObject1222 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1222WithDefaults() *InlineObject1222 {
	this := InlineObject1222{}
	return &this
}

// GetXNum returns the XNum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1222) GetXNum() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.XNum
}

// GetXNumOk returns a tuple with the XNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1222) GetXNumOk() (*AnyOfobject, bool) {
	if o == nil || o.XNum == nil {
		return nil, false
	}
	return &o.XNum, true
}

// HasXNum returns a boolean if a field has been set.
func (o *InlineObject1222) HasXNum() bool {
	if o != nil && o.XNum != nil {
		return true
	}

	return false
}

// SetXNum gets a reference to the given AnyOfobject and assigns it to the XNum field.
func (o *InlineObject1222) SetXNum(v AnyOfobject) {
	o.XNum = v
}

// GetYNum returns the YNum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1222) GetYNum() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.YNum
}

// GetYNumOk returns a tuple with the YNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1222) GetYNumOk() (*AnyOfobject, bool) {
	if o == nil || o.YNum == nil {
		return nil, false
	}
	return &o.YNum, true
}

// HasYNum returns a boolean if a field has been set.
func (o *InlineObject1222) HasYNum() bool {
	if o != nil && o.YNum != nil {
		return true
	}

	return false
}

// SetYNum gets a reference to the given AnyOfobject and assigns it to the YNum field.
func (o *InlineObject1222) SetYNum(v AnyOfobject) {
	o.YNum = v
}

func (o InlineObject1222) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.XNum != nil {
		toSerialize["xNum"] = o.XNum
	}
	if o.YNum != nil {
		toSerialize["yNum"] = o.YNum
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1222 struct {
	value *InlineObject1222
	isSet bool
}

func (v NullableInlineObject1222) Get() *InlineObject1222 {
	return v.value
}

func (v *NullableInlineObject1222) Set(val *InlineObject1222) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1222) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1222) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1222(val *InlineObject1222) *NullableInlineObject1222 {
	return &NullableInlineObject1222{value: val, isSet: true}
}

func (v NullableInlineObject1222) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1222) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

