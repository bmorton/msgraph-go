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

// InlineObject1274 struct for InlineObject1274
type InlineObject1274 struct {
	Values AnyOfobject `json:"values,omitempty"`
}

// NewInlineObject1274 instantiates a new InlineObject1274 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1274() *InlineObject1274 {
	this := InlineObject1274{}
	return &this
}

// NewInlineObject1274WithDefaults instantiates a new InlineObject1274 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1274WithDefaults() *InlineObject1274 {
	this := InlineObject1274{}
	return &this
}

// GetValues returns the Values field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1274) GetValues() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1274) GetValuesOk() (*AnyOfobject, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *InlineObject1274) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given AnyOfobject and assigns it to the Values field.
func (o *InlineObject1274) SetValues(v AnyOfobject) {
	o.Values = v
}

func (o InlineObject1274) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1274 struct {
	value *InlineObject1274
	isSet bool
}

func (v NullableInlineObject1274) Get() *InlineObject1274 {
	return v.value
}

func (v *NullableInlineObject1274) Set(val *InlineObject1274) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1274) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1274) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1274(val *InlineObject1274) *NullableInlineObject1274 {
	return &NullableInlineObject1274{value: val, isSet: true}
}

func (v NullableInlineObject1274) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1274) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


