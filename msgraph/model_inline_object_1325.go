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

// InlineObject1325 struct for InlineObject1325
type InlineObject1325 struct {
	ErrorVal AnyOfobject `json:"errorVal,omitempty"`
}

// NewInlineObject1325 instantiates a new InlineObject1325 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1325() *InlineObject1325 {
	this := InlineObject1325{}
	return &this
}

// NewInlineObject1325WithDefaults instantiates a new InlineObject1325 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1325WithDefaults() *InlineObject1325 {
	this := InlineObject1325{}
	return &this
}

// GetErrorVal returns the ErrorVal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1325) GetErrorVal() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.ErrorVal
}

// GetErrorValOk returns a tuple with the ErrorVal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1325) GetErrorValOk() (*AnyOfobject, bool) {
	if o == nil || o.ErrorVal == nil {
		return nil, false
	}
	return &o.ErrorVal, true
}

// HasErrorVal returns a boolean if a field has been set.
func (o *InlineObject1325) HasErrorVal() bool {
	if o != nil && o.ErrorVal != nil {
		return true
	}

	return false
}

// SetErrorVal gets a reference to the given AnyOfobject and assigns it to the ErrorVal field.
func (o *InlineObject1325) SetErrorVal(v AnyOfobject) {
	o.ErrorVal = v
}

func (o InlineObject1325) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorVal != nil {
		toSerialize["errorVal"] = o.ErrorVal
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1325 struct {
	value *InlineObject1325
	isSet bool
}

func (v NullableInlineObject1325) Get() *InlineObject1325 {
	return v.value
}

func (v *NullableInlineObject1325) Set(val *InlineObject1325) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1325) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1325) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1325(val *InlineObject1325) *NullableInlineObject1325 {
	return &NullableInlineObject1325{value: val, isSet: true}
}

func (v NullableInlineObject1325) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1325) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

