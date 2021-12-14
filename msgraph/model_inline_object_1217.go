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

// InlineObject1217 struct for InlineObject1217
type InlineObject1217 struct {
	Reference AnyOfobject `json:"reference,omitempty"`
}

// NewInlineObject1217 instantiates a new InlineObject1217 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1217() *InlineObject1217 {
	this := InlineObject1217{}
	return &this
}

// NewInlineObject1217WithDefaults instantiates a new InlineObject1217 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1217WithDefaults() *InlineObject1217 {
	this := InlineObject1217{}
	return &this
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1217) GetReference() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1217) GetReferenceOk() (*AnyOfobject, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *InlineObject1217) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given AnyOfobject and assigns it to the Reference field.
func (o *InlineObject1217) SetReference(v AnyOfobject) {
	o.Reference = v
}

func (o InlineObject1217) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1217 struct {
	value *InlineObject1217
	isSet bool
}

func (v NullableInlineObject1217) Get() *InlineObject1217 {
	return v.value
}

func (v *NullableInlineObject1217) Set(val *InlineObject1217) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1217) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1217) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1217(val *InlineObject1217) *NullableInlineObject1217 {
	return &NullableInlineObject1217{value: val, isSet: true}
}

func (v NullableInlineObject1217) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1217) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


