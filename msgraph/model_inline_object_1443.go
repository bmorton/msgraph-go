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

// InlineObject1443 struct for InlineObject1443
type InlineObject1443 struct {
	Probability AnyOfobject `json:"probability,omitempty"`
}

// NewInlineObject1443 instantiates a new InlineObject1443 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1443() *InlineObject1443 {
	this := InlineObject1443{}
	return &this
}

// NewInlineObject1443WithDefaults instantiates a new InlineObject1443 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1443WithDefaults() *InlineObject1443 {
	this := InlineObject1443{}
	return &this
}

// GetProbability returns the Probability field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1443) GetProbability() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Probability
}

// GetProbabilityOk returns a tuple with the Probability field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1443) GetProbabilityOk() (*AnyOfobject, bool) {
	if o == nil || o.Probability == nil {
		return nil, false
	}
	return &o.Probability, true
}

// HasProbability returns a boolean if a field has been set.
func (o *InlineObject1443) HasProbability() bool {
	if o != nil && o.Probability != nil {
		return true
	}

	return false
}

// SetProbability gets a reference to the given AnyOfobject and assigns it to the Probability field.
func (o *InlineObject1443) SetProbability(v AnyOfobject) {
	o.Probability = v
}

func (o InlineObject1443) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Probability != nil {
		toSerialize["probability"] = o.Probability
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1443 struct {
	value *InlineObject1443
	isSet bool
}

func (v NullableInlineObject1443) Get() *InlineObject1443 {
	return v.value
}

func (v *NullableInlineObject1443) Set(val *InlineObject1443) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1443) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1443) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1443(val *InlineObject1443) *NullableInlineObject1443 {
	return &NullableInlineObject1443{value: val, isSet: true}
}

func (v NullableInlineObject1443) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1443) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


