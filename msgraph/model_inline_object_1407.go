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

// InlineObject1407 struct for InlineObject1407
type InlineObject1407 struct {
	Array AnyOfobject `json:"array,omitempty"`
	K AnyOfobject `json:"k,omitempty"`
}

// NewInlineObject1407 instantiates a new InlineObject1407 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1407() *InlineObject1407 {
	this := InlineObject1407{}
	return &this
}

// NewInlineObject1407WithDefaults instantiates a new InlineObject1407 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1407WithDefaults() *InlineObject1407 {
	this := InlineObject1407{}
	return &this
}

// GetArray returns the Array field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1407) GetArray() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1407) GetArrayOk() (*AnyOfobject, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return &o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *InlineObject1407) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given AnyOfobject and assigns it to the Array field.
func (o *InlineObject1407) SetArray(v AnyOfobject) {
	o.Array = v
}

// GetK returns the K field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1407) GetK() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.K
}

// GetKOk returns a tuple with the K field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1407) GetKOk() (*AnyOfobject, bool) {
	if o == nil || o.K == nil {
		return nil, false
	}
	return &o.K, true
}

// HasK returns a boolean if a field has been set.
func (o *InlineObject1407) HasK() bool {
	if o != nil && o.K != nil {
		return true
	}

	return false
}

// SetK gets a reference to the given AnyOfobject and assigns it to the K field.
func (o *InlineObject1407) SetK(v AnyOfobject) {
	o.K = v
}

func (o InlineObject1407) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Array != nil {
		toSerialize["array"] = o.Array
	}
	if o.K != nil {
		toSerialize["k"] = o.K
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1407 struct {
	value *InlineObject1407
	isSet bool
}

func (v NullableInlineObject1407) Get() *InlineObject1407 {
	return v.value
}

func (v *NullableInlineObject1407) Set(val *InlineObject1407) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1407) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1407) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1407(val *InlineObject1407) *NullableInlineObject1407 {
	return &NullableInlineObject1407{value: val, isSet: true}
}

func (v NullableInlineObject1407) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1407) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


