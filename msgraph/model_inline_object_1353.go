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

// InlineObject1353 struct for InlineObject1353
type InlineObject1353 struct {
	Number AnyOfobject `json:"number,omitempty"`
	Step AnyOfobject `json:"step,omitempty"`
}

// NewInlineObject1353 instantiates a new InlineObject1353 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1353() *InlineObject1353 {
	this := InlineObject1353{}
	return &this
}

// NewInlineObject1353WithDefaults instantiates a new InlineObject1353 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1353WithDefaults() *InlineObject1353 {
	this := InlineObject1353{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1353) GetNumber() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1353) GetNumberOk() (*AnyOfobject, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return &o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *InlineObject1353) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given AnyOfobject and assigns it to the Number field.
func (o *InlineObject1353) SetNumber(v AnyOfobject) {
	o.Number = v
}

// GetStep returns the Step field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1353) GetStep() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Step
}

// GetStepOk returns a tuple with the Step field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1353) GetStepOk() (*AnyOfobject, bool) {
	if o == nil || o.Step == nil {
		return nil, false
	}
	return &o.Step, true
}

// HasStep returns a boolean if a field has been set.
func (o *InlineObject1353) HasStep() bool {
	if o != nil && o.Step != nil {
		return true
	}

	return false
}

// SetStep gets a reference to the given AnyOfobject and assigns it to the Step field.
func (o *InlineObject1353) SetStep(v AnyOfobject) {
	o.Step = v
}

func (o InlineObject1353) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.Step != nil {
		toSerialize["step"] = o.Step
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1353 struct {
	value *InlineObject1353
	isSet bool
}

func (v NullableInlineObject1353) Get() *InlineObject1353 {
	return v.value
}

func (v *NullableInlineObject1353) Set(val *InlineObject1353) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1353) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1353) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1353(val *InlineObject1353) *NullableInlineObject1353 {
	return &NullableInlineObject1353{value: val, isSet: true}
}

func (v NullableInlineObject1353) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1353) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


