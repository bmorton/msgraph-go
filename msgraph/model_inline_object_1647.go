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

// InlineObject1647 struct for InlineObject1647
type InlineObject1647 struct {
	Index NullableInt32 `json:"index,omitempty"`
	Values AnyOfobject `json:"values,omitempty"`
}

// NewInlineObject1647 instantiates a new InlineObject1647 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1647() *InlineObject1647 {
	this := InlineObject1647{}
	return &this
}

// NewInlineObject1647WithDefaults instantiates a new InlineObject1647 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1647WithDefaults() *InlineObject1647 {
	this := InlineObject1647{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1647) GetIndex() int32 {
	if o == nil || o.Index.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Index.Get()
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1647) GetIndexOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Index.Get(), o.Index.IsSet()
}

// HasIndex returns a boolean if a field has been set.
func (o *InlineObject1647) HasIndex() bool {
	if o != nil && o.Index.IsSet() {
		return true
	}

	return false
}

// SetIndex gets a reference to the given NullableInt32 and assigns it to the Index field.
func (o *InlineObject1647) SetIndex(v int32) {
	o.Index.Set(&v)
}
// SetIndexNil sets the value for Index to be an explicit nil
func (o *InlineObject1647) SetIndexNil() {
	o.Index.Set(nil)
}

// UnsetIndex ensures that no value is present for Index, not even an explicit nil
func (o *InlineObject1647) UnsetIndex() {
	o.Index.Unset()
}

// GetValues returns the Values field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1647) GetValues() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1647) GetValuesOk() (*AnyOfobject, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *InlineObject1647) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given AnyOfobject and assigns it to the Values field.
func (o *InlineObject1647) SetValues(v AnyOfobject) {
	o.Values = v
}

func (o InlineObject1647) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Index.IsSet() {
		toSerialize["index"] = o.Index.Get()
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1647 struct {
	value *InlineObject1647
	isSet bool
}

func (v NullableInlineObject1647) Get() *InlineObject1647 {
	return v.value
}

func (v *NullableInlineObject1647) Set(val *InlineObject1647) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1647) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1647) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1647(val *InlineObject1647) *NullableInlineObject1647 {
	return &NullableInlineObject1647{value: val, isSet: true}
}

func (v NullableInlineObject1647) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1647) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


