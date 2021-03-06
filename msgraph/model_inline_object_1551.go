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

// InlineObject1551 struct for InlineObject1551
type InlineObject1551 struct {
	SerialNumber AnyOfobject `json:"serialNumber,omitempty"`
	ReturnType AnyOfobject `json:"returnType,omitempty"`
}

// NewInlineObject1551 instantiates a new InlineObject1551 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1551() *InlineObject1551 {
	this := InlineObject1551{}
	return &this
}

// NewInlineObject1551WithDefaults instantiates a new InlineObject1551 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1551WithDefaults() *InlineObject1551 {
	this := InlineObject1551{}
	return &this
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1551) GetSerialNumber() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1551) GetSerialNumberOk() (*AnyOfobject, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return &o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *InlineObject1551) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given AnyOfobject and assigns it to the SerialNumber field.
func (o *InlineObject1551) SetSerialNumber(v AnyOfobject) {
	o.SerialNumber = v
}

// GetReturnType returns the ReturnType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1551) GetReturnType() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.ReturnType
}

// GetReturnTypeOk returns a tuple with the ReturnType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1551) GetReturnTypeOk() (*AnyOfobject, bool) {
	if o == nil || o.ReturnType == nil {
		return nil, false
	}
	return &o.ReturnType, true
}

// HasReturnType returns a boolean if a field has been set.
func (o *InlineObject1551) HasReturnType() bool {
	if o != nil && o.ReturnType != nil {
		return true
	}

	return false
}

// SetReturnType gets a reference to the given AnyOfobject and assigns it to the ReturnType field.
func (o *InlineObject1551) SetReturnType(v AnyOfobject) {
	o.ReturnType = v
}

func (o InlineObject1551) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SerialNumber != nil {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	if o.ReturnType != nil {
		toSerialize["returnType"] = o.ReturnType
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1551 struct {
	value *InlineObject1551
	isSet bool
}

func (v NullableInlineObject1551) Get() *InlineObject1551 {
	return v.value
}

func (v *NullableInlineObject1551) Set(val *InlineObject1551) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1551) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1551) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1551(val *InlineObject1551) *NullableInlineObject1551 {
	return &NullableInlineObject1551{value: val, isSet: true}
}

func (v NullableInlineObject1551) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1551) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


