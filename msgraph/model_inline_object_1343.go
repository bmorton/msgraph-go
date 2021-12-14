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

// InlineObject1343 struct for InlineObject1343
type InlineObject1343 struct {
	Rate AnyOfobject `json:"rate,omitempty"`
	Nper AnyOfobject `json:"nper,omitempty"`
	Pmt AnyOfobject `json:"pmt,omitempty"`
	Pv AnyOfobject `json:"pv,omitempty"`
	Type AnyOfobject `json:"type,omitempty"`
}

// NewInlineObject1343 instantiates a new InlineObject1343 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1343() *InlineObject1343 {
	this := InlineObject1343{}
	return &this
}

// NewInlineObject1343WithDefaults instantiates a new InlineObject1343 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1343WithDefaults() *InlineObject1343 {
	this := InlineObject1343{}
	return &this
}

// GetRate returns the Rate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1343) GetRate() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1343) GetRateOk() (*AnyOfobject, bool) {
	if o == nil || o.Rate == nil {
		return nil, false
	}
	return &o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *InlineObject1343) HasRate() bool {
	if o != nil && o.Rate != nil {
		return true
	}

	return false
}

// SetRate gets a reference to the given AnyOfobject and assigns it to the Rate field.
func (o *InlineObject1343) SetRate(v AnyOfobject) {
	o.Rate = v
}

// GetNper returns the Nper field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1343) GetNper() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Nper
}

// GetNperOk returns a tuple with the Nper field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1343) GetNperOk() (*AnyOfobject, bool) {
	if o == nil || o.Nper == nil {
		return nil, false
	}
	return &o.Nper, true
}

// HasNper returns a boolean if a field has been set.
func (o *InlineObject1343) HasNper() bool {
	if o != nil && o.Nper != nil {
		return true
	}

	return false
}

// SetNper gets a reference to the given AnyOfobject and assigns it to the Nper field.
func (o *InlineObject1343) SetNper(v AnyOfobject) {
	o.Nper = v
}

// GetPmt returns the Pmt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1343) GetPmt() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Pmt
}

// GetPmtOk returns a tuple with the Pmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1343) GetPmtOk() (*AnyOfobject, bool) {
	if o == nil || o.Pmt == nil {
		return nil, false
	}
	return &o.Pmt, true
}

// HasPmt returns a boolean if a field has been set.
func (o *InlineObject1343) HasPmt() bool {
	if o != nil && o.Pmt != nil {
		return true
	}

	return false
}

// SetPmt gets a reference to the given AnyOfobject and assigns it to the Pmt field.
func (o *InlineObject1343) SetPmt(v AnyOfobject) {
	o.Pmt = v
}

// GetPv returns the Pv field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1343) GetPv() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Pv
}

// GetPvOk returns a tuple with the Pv field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1343) GetPvOk() (*AnyOfobject, bool) {
	if o == nil || o.Pv == nil {
		return nil, false
	}
	return &o.Pv, true
}

// HasPv returns a boolean if a field has been set.
func (o *InlineObject1343) HasPv() bool {
	if o != nil && o.Pv != nil {
		return true
	}

	return false
}

// SetPv gets a reference to the given AnyOfobject and assigns it to the Pv field.
func (o *InlineObject1343) SetPv(v AnyOfobject) {
	o.Pv = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1343) GetType() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1343) GetTypeOk() (*AnyOfobject, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineObject1343) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given AnyOfobject and assigns it to the Type field.
func (o *InlineObject1343) SetType(v AnyOfobject) {
	o.Type = v
}

func (o InlineObject1343) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Rate != nil {
		toSerialize["rate"] = o.Rate
	}
	if o.Nper != nil {
		toSerialize["nper"] = o.Nper
	}
	if o.Pmt != nil {
		toSerialize["pmt"] = o.Pmt
	}
	if o.Pv != nil {
		toSerialize["pv"] = o.Pv
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1343 struct {
	value *InlineObject1343
	isSet bool
}

func (v NullableInlineObject1343) Get() *InlineObject1343 {
	return v.value
}

func (v *NullableInlineObject1343) Set(val *InlineObject1343) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1343) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1343) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1343(val *InlineObject1343) *NullableInlineObject1343 {
	return &NullableInlineObject1343{value: val, isSet: true}
}

func (v NullableInlineObject1343) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1343) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


