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

// InlineObject1482 struct for InlineObject1482
type InlineObject1482 struct {
	Nper AnyOfobject `json:"nper,omitempty"`
	Pmt AnyOfobject `json:"pmt,omitempty"`
	Pv AnyOfobject `json:"pv,omitempty"`
	Fv AnyOfobject `json:"fv,omitempty"`
	Type AnyOfobject `json:"type,omitempty"`
	Guess AnyOfobject `json:"guess,omitempty"`
}

// NewInlineObject1482 instantiates a new InlineObject1482 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1482() *InlineObject1482 {
	this := InlineObject1482{}
	return &this
}

// NewInlineObject1482WithDefaults instantiates a new InlineObject1482 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1482WithDefaults() *InlineObject1482 {
	this := InlineObject1482{}
	return &this
}

// GetNper returns the Nper field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1482) GetNper() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Nper
}

// GetNperOk returns a tuple with the Nper field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1482) GetNperOk() (*AnyOfobject, bool) {
	if o == nil || o.Nper == nil {
		return nil, false
	}
	return &o.Nper, true
}

// HasNper returns a boolean if a field has been set.
func (o *InlineObject1482) HasNper() bool {
	if o != nil && o.Nper != nil {
		return true
	}

	return false
}

// SetNper gets a reference to the given AnyOfobject and assigns it to the Nper field.
func (o *InlineObject1482) SetNper(v AnyOfobject) {
	o.Nper = v
}

// GetPmt returns the Pmt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1482) GetPmt() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Pmt
}

// GetPmtOk returns a tuple with the Pmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1482) GetPmtOk() (*AnyOfobject, bool) {
	if o == nil || o.Pmt == nil {
		return nil, false
	}
	return &o.Pmt, true
}

// HasPmt returns a boolean if a field has been set.
func (o *InlineObject1482) HasPmt() bool {
	if o != nil && o.Pmt != nil {
		return true
	}

	return false
}

// SetPmt gets a reference to the given AnyOfobject and assigns it to the Pmt field.
func (o *InlineObject1482) SetPmt(v AnyOfobject) {
	o.Pmt = v
}

// GetPv returns the Pv field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1482) GetPv() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Pv
}

// GetPvOk returns a tuple with the Pv field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1482) GetPvOk() (*AnyOfobject, bool) {
	if o == nil || o.Pv == nil {
		return nil, false
	}
	return &o.Pv, true
}

// HasPv returns a boolean if a field has been set.
func (o *InlineObject1482) HasPv() bool {
	if o != nil && o.Pv != nil {
		return true
	}

	return false
}

// SetPv gets a reference to the given AnyOfobject and assigns it to the Pv field.
func (o *InlineObject1482) SetPv(v AnyOfobject) {
	o.Pv = v
}

// GetFv returns the Fv field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1482) GetFv() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Fv
}

// GetFvOk returns a tuple with the Fv field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1482) GetFvOk() (*AnyOfobject, bool) {
	if o == nil || o.Fv == nil {
		return nil, false
	}
	return &o.Fv, true
}

// HasFv returns a boolean if a field has been set.
func (o *InlineObject1482) HasFv() bool {
	if o != nil && o.Fv != nil {
		return true
	}

	return false
}

// SetFv gets a reference to the given AnyOfobject and assigns it to the Fv field.
func (o *InlineObject1482) SetFv(v AnyOfobject) {
	o.Fv = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1482) GetType() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1482) GetTypeOk() (*AnyOfobject, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineObject1482) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given AnyOfobject and assigns it to the Type field.
func (o *InlineObject1482) SetType(v AnyOfobject) {
	o.Type = v
}

// GetGuess returns the Guess field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1482) GetGuess() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Guess
}

// GetGuessOk returns a tuple with the Guess field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1482) GetGuessOk() (*AnyOfobject, bool) {
	if o == nil || o.Guess == nil {
		return nil, false
	}
	return &o.Guess, true
}

// HasGuess returns a boolean if a field has been set.
func (o *InlineObject1482) HasGuess() bool {
	if o != nil && o.Guess != nil {
		return true
	}

	return false
}

// SetGuess gets a reference to the given AnyOfobject and assigns it to the Guess field.
func (o *InlineObject1482) SetGuess(v AnyOfobject) {
	o.Guess = v
}

func (o InlineObject1482) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Nper != nil {
		toSerialize["nper"] = o.Nper
	}
	if o.Pmt != nil {
		toSerialize["pmt"] = o.Pmt
	}
	if o.Pv != nil {
		toSerialize["pv"] = o.Pv
	}
	if o.Fv != nil {
		toSerialize["fv"] = o.Fv
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Guess != nil {
		toSerialize["guess"] = o.Guess
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1482 struct {
	value *InlineObject1482
	isSet bool
}

func (v NullableInlineObject1482) Get() *InlineObject1482 {
	return v.value
}

func (v *NullableInlineObject1482) Set(val *InlineObject1482) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1482) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1482) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1482(val *InlineObject1482) *NullableInlineObject1482 {
	return &NullableInlineObject1482{value: val, isSet: true}
}

func (v NullableInlineObject1482) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1482) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

