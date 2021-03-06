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

// InlineObject1362 struct for InlineObject1362
type InlineObject1362 struct {
	LogicalTest AnyOfobject `json:"logicalTest,omitempty"`
	ValueIfTrue AnyOfobject `json:"valueIfTrue,omitempty"`
	ValueIfFalse AnyOfobject `json:"valueIfFalse,omitempty"`
}

// NewInlineObject1362 instantiates a new InlineObject1362 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1362() *InlineObject1362 {
	this := InlineObject1362{}
	return &this
}

// NewInlineObject1362WithDefaults instantiates a new InlineObject1362 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1362WithDefaults() *InlineObject1362 {
	this := InlineObject1362{}
	return &this
}

// GetLogicalTest returns the LogicalTest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1362) GetLogicalTest() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.LogicalTest
}

// GetLogicalTestOk returns a tuple with the LogicalTest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1362) GetLogicalTestOk() (*AnyOfobject, bool) {
	if o == nil || o.LogicalTest == nil {
		return nil, false
	}
	return &o.LogicalTest, true
}

// HasLogicalTest returns a boolean if a field has been set.
func (o *InlineObject1362) HasLogicalTest() bool {
	if o != nil && o.LogicalTest != nil {
		return true
	}

	return false
}

// SetLogicalTest gets a reference to the given AnyOfobject and assigns it to the LogicalTest field.
func (o *InlineObject1362) SetLogicalTest(v AnyOfobject) {
	o.LogicalTest = v
}

// GetValueIfTrue returns the ValueIfTrue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1362) GetValueIfTrue() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.ValueIfTrue
}

// GetValueIfTrueOk returns a tuple with the ValueIfTrue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1362) GetValueIfTrueOk() (*AnyOfobject, bool) {
	if o == nil || o.ValueIfTrue == nil {
		return nil, false
	}
	return &o.ValueIfTrue, true
}

// HasValueIfTrue returns a boolean if a field has been set.
func (o *InlineObject1362) HasValueIfTrue() bool {
	if o != nil && o.ValueIfTrue != nil {
		return true
	}

	return false
}

// SetValueIfTrue gets a reference to the given AnyOfobject and assigns it to the ValueIfTrue field.
func (o *InlineObject1362) SetValueIfTrue(v AnyOfobject) {
	o.ValueIfTrue = v
}

// GetValueIfFalse returns the ValueIfFalse field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1362) GetValueIfFalse() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.ValueIfFalse
}

// GetValueIfFalseOk returns a tuple with the ValueIfFalse field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1362) GetValueIfFalseOk() (*AnyOfobject, bool) {
	if o == nil || o.ValueIfFalse == nil {
		return nil, false
	}
	return &o.ValueIfFalse, true
}

// HasValueIfFalse returns a boolean if a field has been set.
func (o *InlineObject1362) HasValueIfFalse() bool {
	if o != nil && o.ValueIfFalse != nil {
		return true
	}

	return false
}

// SetValueIfFalse gets a reference to the given AnyOfobject and assigns it to the ValueIfFalse field.
func (o *InlineObject1362) SetValueIfFalse(v AnyOfobject) {
	o.ValueIfFalse = v
}

func (o InlineObject1362) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LogicalTest != nil {
		toSerialize["logicalTest"] = o.LogicalTest
	}
	if o.ValueIfTrue != nil {
		toSerialize["valueIfTrue"] = o.ValueIfTrue
	}
	if o.ValueIfFalse != nil {
		toSerialize["valueIfFalse"] = o.ValueIfFalse
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1362 struct {
	value *InlineObject1362
	isSet bool
}

func (v NullableInlineObject1362) Get() *InlineObject1362 {
	return v.value
}

func (v *NullableInlineObject1362) Set(val *InlineObject1362) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1362) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1362) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1362(val *InlineObject1362) *NullableInlineObject1362 {
	return &NullableInlineObject1362{value: val, isSet: true}
}

func (v NullableInlineObject1362) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1362) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


