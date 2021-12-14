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

// InlineObject1215 struct for InlineObject1215
type InlineObject1215 struct {
	Values AnyOfobject `json:"values,omitempty"`
}

// NewInlineObject1215 instantiates a new InlineObject1215 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1215() *InlineObject1215 {
	this := InlineObject1215{}
	return &this
}

// NewInlineObject1215WithDefaults instantiates a new InlineObject1215 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1215WithDefaults() *InlineObject1215 {
	this := InlineObject1215{}
	return &this
}

// GetValues returns the Values field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1215) GetValues() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1215) GetValuesOk() (*AnyOfobject, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *InlineObject1215) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given AnyOfobject and assigns it to the Values field.
func (o *InlineObject1215) SetValues(v AnyOfobject) {
	o.Values = v
}

func (o InlineObject1215) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1215 struct {
	value *InlineObject1215
	isSet bool
}

func (v NullableInlineObject1215) Get() *InlineObject1215 {
	return v.value
}

func (v *NullableInlineObject1215) Set(val *InlineObject1215) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1215) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1215) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1215(val *InlineObject1215) *NullableInlineObject1215 {
	return &NullableInlineObject1215{value: val, isSet: true}
}

func (v NullableInlineObject1215) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1215) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


