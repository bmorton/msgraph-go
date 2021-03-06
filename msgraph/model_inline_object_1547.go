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

// InlineObject1547 struct for InlineObject1547
type InlineObject1547 struct {
	Values AnyOfobject `json:"values,omitempty"`
}

// NewInlineObject1547 instantiates a new InlineObject1547 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1547() *InlineObject1547 {
	this := InlineObject1547{}
	return &this
}

// NewInlineObject1547WithDefaults instantiates a new InlineObject1547 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1547WithDefaults() *InlineObject1547 {
	this := InlineObject1547{}
	return &this
}

// GetValues returns the Values field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1547) GetValues() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1547) GetValuesOk() (*AnyOfobject, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *InlineObject1547) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given AnyOfobject and assigns it to the Values field.
func (o *InlineObject1547) SetValues(v AnyOfobject) {
	o.Values = v
}

func (o InlineObject1547) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1547 struct {
	value *InlineObject1547
	isSet bool
}

func (v NullableInlineObject1547) Get() *InlineObject1547 {
	return v.value
}

func (v *NullableInlineObject1547) Set(val *InlineObject1547) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1547) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1547) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1547(val *InlineObject1547) *NullableInlineObject1547 {
	return &NullableInlineObject1547{value: val, isSet: true}
}

func (v NullableInlineObject1547) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1547) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


