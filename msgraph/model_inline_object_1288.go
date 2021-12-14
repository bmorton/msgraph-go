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

// InlineObject1288 struct for InlineObject1288
type InlineObject1288 struct {
	SerialNumber AnyOfobject `json:"serialNumber,omitempty"`
}

// NewInlineObject1288 instantiates a new InlineObject1288 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1288() *InlineObject1288 {
	this := InlineObject1288{}
	return &this
}

// NewInlineObject1288WithDefaults instantiates a new InlineObject1288 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1288WithDefaults() *InlineObject1288 {
	this := InlineObject1288{}
	return &this
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1288) GetSerialNumber() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1288) GetSerialNumberOk() (*AnyOfobject, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return &o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *InlineObject1288) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given AnyOfobject and assigns it to the SerialNumber field.
func (o *InlineObject1288) SetSerialNumber(v AnyOfobject) {
	o.SerialNumber = v
}

func (o InlineObject1288) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SerialNumber != nil {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1288 struct {
	value *InlineObject1288
	isSet bool
}

func (v NullableInlineObject1288) Get() *InlineObject1288 {
	return v.value
}

func (v *NullableInlineObject1288) Set(val *InlineObject1288) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1288) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1288) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1288(val *InlineObject1288) *NullableInlineObject1288 {
	return &NullableInlineObject1288{value: val, isSet: true}
}

func (v NullableInlineObject1288) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1288) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


