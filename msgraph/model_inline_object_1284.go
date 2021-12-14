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

// InlineObject1284 struct for InlineObject1284
type InlineObject1284 struct {
	Rate AnyOfobject `json:"rate,omitempty"`
	Nper AnyOfobject `json:"nper,omitempty"`
	Pv AnyOfobject `json:"pv,omitempty"`
	StartPeriod AnyOfobject `json:"startPeriod,omitempty"`
	EndPeriod AnyOfobject `json:"endPeriod,omitempty"`
	Type AnyOfobject `json:"type,omitempty"`
}

// NewInlineObject1284 instantiates a new InlineObject1284 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1284() *InlineObject1284 {
	this := InlineObject1284{}
	return &this
}

// NewInlineObject1284WithDefaults instantiates a new InlineObject1284 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1284WithDefaults() *InlineObject1284 {
	this := InlineObject1284{}
	return &this
}

// GetRate returns the Rate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1284) GetRate() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1284) GetRateOk() (*AnyOfobject, bool) {
	if o == nil || o.Rate == nil {
		return nil, false
	}
	return &o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *InlineObject1284) HasRate() bool {
	if o != nil && o.Rate != nil {
		return true
	}

	return false
}

// SetRate gets a reference to the given AnyOfobject and assigns it to the Rate field.
func (o *InlineObject1284) SetRate(v AnyOfobject) {
	o.Rate = v
}

// GetNper returns the Nper field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1284) GetNper() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Nper
}

// GetNperOk returns a tuple with the Nper field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1284) GetNperOk() (*AnyOfobject, bool) {
	if o == nil || o.Nper == nil {
		return nil, false
	}
	return &o.Nper, true
}

// HasNper returns a boolean if a field has been set.
func (o *InlineObject1284) HasNper() bool {
	if o != nil && o.Nper != nil {
		return true
	}

	return false
}

// SetNper gets a reference to the given AnyOfobject and assigns it to the Nper field.
func (o *InlineObject1284) SetNper(v AnyOfobject) {
	o.Nper = v
}

// GetPv returns the Pv field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1284) GetPv() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Pv
}

// GetPvOk returns a tuple with the Pv field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1284) GetPvOk() (*AnyOfobject, bool) {
	if o == nil || o.Pv == nil {
		return nil, false
	}
	return &o.Pv, true
}

// HasPv returns a boolean if a field has been set.
func (o *InlineObject1284) HasPv() bool {
	if o != nil && o.Pv != nil {
		return true
	}

	return false
}

// SetPv gets a reference to the given AnyOfobject and assigns it to the Pv field.
func (o *InlineObject1284) SetPv(v AnyOfobject) {
	o.Pv = v
}

// GetStartPeriod returns the StartPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1284) GetStartPeriod() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.StartPeriod
}

// GetStartPeriodOk returns a tuple with the StartPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1284) GetStartPeriodOk() (*AnyOfobject, bool) {
	if o == nil || o.StartPeriod == nil {
		return nil, false
	}
	return &o.StartPeriod, true
}

// HasStartPeriod returns a boolean if a field has been set.
func (o *InlineObject1284) HasStartPeriod() bool {
	if o != nil && o.StartPeriod != nil {
		return true
	}

	return false
}

// SetStartPeriod gets a reference to the given AnyOfobject and assigns it to the StartPeriod field.
func (o *InlineObject1284) SetStartPeriod(v AnyOfobject) {
	o.StartPeriod = v
}

// GetEndPeriod returns the EndPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1284) GetEndPeriod() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.EndPeriod
}

// GetEndPeriodOk returns a tuple with the EndPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1284) GetEndPeriodOk() (*AnyOfobject, bool) {
	if o == nil || o.EndPeriod == nil {
		return nil, false
	}
	return &o.EndPeriod, true
}

// HasEndPeriod returns a boolean if a field has been set.
func (o *InlineObject1284) HasEndPeriod() bool {
	if o != nil && o.EndPeriod != nil {
		return true
	}

	return false
}

// SetEndPeriod gets a reference to the given AnyOfobject and assigns it to the EndPeriod field.
func (o *InlineObject1284) SetEndPeriod(v AnyOfobject) {
	o.EndPeriod = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1284) GetType() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1284) GetTypeOk() (*AnyOfobject, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineObject1284) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given AnyOfobject and assigns it to the Type field.
func (o *InlineObject1284) SetType(v AnyOfobject) {
	o.Type = v
}

func (o InlineObject1284) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Rate != nil {
		toSerialize["rate"] = o.Rate
	}
	if o.Nper != nil {
		toSerialize["nper"] = o.Nper
	}
	if o.Pv != nil {
		toSerialize["pv"] = o.Pv
	}
	if o.StartPeriod != nil {
		toSerialize["startPeriod"] = o.StartPeriod
	}
	if o.EndPeriod != nil {
		toSerialize["endPeriod"] = o.EndPeriod
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1284 struct {
	value *InlineObject1284
	isSet bool
}

func (v NullableInlineObject1284) Get() *InlineObject1284 {
	return v.value
}

func (v *NullableInlineObject1284) Set(val *InlineObject1284) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1284) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1284) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1284(val *InlineObject1284) *NullableInlineObject1284 {
	return &NullableInlineObject1284{value: val, isSet: true}
}

func (v NullableInlineObject1284) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1284) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


