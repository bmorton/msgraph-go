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

// InlineObject1429 struct for InlineObject1429
type InlineObject1429 struct {
	SerialNumber AnyOfobject `json:"serialNumber,omitempty"`
}

// NewInlineObject1429 instantiates a new InlineObject1429 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1429() *InlineObject1429 {
	this := InlineObject1429{}
	return &this
}

// NewInlineObject1429WithDefaults instantiates a new InlineObject1429 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1429WithDefaults() *InlineObject1429 {
	this := InlineObject1429{}
	return &this
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1429) GetSerialNumber() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1429) GetSerialNumberOk() (*AnyOfobject, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return &o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *InlineObject1429) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given AnyOfobject and assigns it to the SerialNumber field.
func (o *InlineObject1429) SetSerialNumber(v AnyOfobject) {
	o.SerialNumber = v
}

func (o InlineObject1429) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SerialNumber != nil {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1429 struct {
	value *InlineObject1429
	isSet bool
}

func (v NullableInlineObject1429) Get() *InlineObject1429 {
	return v.value
}

func (v *NullableInlineObject1429) Set(val *InlineObject1429) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1429) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1429) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1429(val *InlineObject1429) *NullableInlineObject1429 {
	return &NullableInlineObject1429{value: val, isSet: true}
}

func (v NullableInlineObject1429) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1429) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

