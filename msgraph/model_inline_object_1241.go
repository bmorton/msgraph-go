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

// InlineObject1241 struct for InlineObject1241
type InlineObject1241 struct {
	Trials AnyOfobject `json:"trials,omitempty"`
	ProbabilityS AnyOfobject `json:"probabilityS,omitempty"`
	NumberS AnyOfobject `json:"numberS,omitempty"`
	NumberS2 AnyOfobject `json:"numberS2,omitempty"`
}

// NewInlineObject1241 instantiates a new InlineObject1241 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1241() *InlineObject1241 {
	this := InlineObject1241{}
	return &this
}

// NewInlineObject1241WithDefaults instantiates a new InlineObject1241 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1241WithDefaults() *InlineObject1241 {
	this := InlineObject1241{}
	return &this
}

// GetTrials returns the Trials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1241) GetTrials() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Trials
}

// GetTrialsOk returns a tuple with the Trials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1241) GetTrialsOk() (*AnyOfobject, bool) {
	if o == nil || o.Trials == nil {
		return nil, false
	}
	return &o.Trials, true
}

// HasTrials returns a boolean if a field has been set.
func (o *InlineObject1241) HasTrials() bool {
	if o != nil && o.Trials != nil {
		return true
	}

	return false
}

// SetTrials gets a reference to the given AnyOfobject and assigns it to the Trials field.
func (o *InlineObject1241) SetTrials(v AnyOfobject) {
	o.Trials = v
}

// GetProbabilityS returns the ProbabilityS field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1241) GetProbabilityS() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.ProbabilityS
}

// GetProbabilitySOk returns a tuple with the ProbabilityS field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1241) GetProbabilitySOk() (*AnyOfobject, bool) {
	if o == nil || o.ProbabilityS == nil {
		return nil, false
	}
	return &o.ProbabilityS, true
}

// HasProbabilityS returns a boolean if a field has been set.
func (o *InlineObject1241) HasProbabilityS() bool {
	if o != nil && o.ProbabilityS != nil {
		return true
	}

	return false
}

// SetProbabilityS gets a reference to the given AnyOfobject and assigns it to the ProbabilityS field.
func (o *InlineObject1241) SetProbabilityS(v AnyOfobject) {
	o.ProbabilityS = v
}

// GetNumberS returns the NumberS field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1241) GetNumberS() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.NumberS
}

// GetNumberSOk returns a tuple with the NumberS field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1241) GetNumberSOk() (*AnyOfobject, bool) {
	if o == nil || o.NumberS == nil {
		return nil, false
	}
	return &o.NumberS, true
}

// HasNumberS returns a boolean if a field has been set.
func (o *InlineObject1241) HasNumberS() bool {
	if o != nil && o.NumberS != nil {
		return true
	}

	return false
}

// SetNumberS gets a reference to the given AnyOfobject and assigns it to the NumberS field.
func (o *InlineObject1241) SetNumberS(v AnyOfobject) {
	o.NumberS = v
}

// GetNumberS2 returns the NumberS2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1241) GetNumberS2() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.NumberS2
}

// GetNumberS2Ok returns a tuple with the NumberS2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1241) GetNumberS2Ok() (*AnyOfobject, bool) {
	if o == nil || o.NumberS2 == nil {
		return nil, false
	}
	return &o.NumberS2, true
}

// HasNumberS2 returns a boolean if a field has been set.
func (o *InlineObject1241) HasNumberS2() bool {
	if o != nil && o.NumberS2 != nil {
		return true
	}

	return false
}

// SetNumberS2 gets a reference to the given AnyOfobject and assigns it to the NumberS2 field.
func (o *InlineObject1241) SetNumberS2(v AnyOfobject) {
	o.NumberS2 = v
}

func (o InlineObject1241) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Trials != nil {
		toSerialize["trials"] = o.Trials
	}
	if o.ProbabilityS != nil {
		toSerialize["probabilityS"] = o.ProbabilityS
	}
	if o.NumberS != nil {
		toSerialize["numberS"] = o.NumberS
	}
	if o.NumberS2 != nil {
		toSerialize["numberS2"] = o.NumberS2
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1241 struct {
	value *InlineObject1241
	isSet bool
}

func (v NullableInlineObject1241) Get() *InlineObject1241 {
	return v.value
}

func (v *NullableInlineObject1241) Set(val *InlineObject1241) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1241) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1241) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1241(val *InlineObject1241) *NullableInlineObject1241 {
	return &NullableInlineObject1241{value: val, isSet: true}
}

func (v NullableInlineObject1241) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1241) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

