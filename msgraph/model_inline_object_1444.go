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

// InlineObject1444 struct for InlineObject1444
type InlineObject1444 struct {
	Logical AnyOfobject `json:"logical,omitempty"`
}

// NewInlineObject1444 instantiates a new InlineObject1444 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1444() *InlineObject1444 {
	this := InlineObject1444{}
	return &this
}

// NewInlineObject1444WithDefaults instantiates a new InlineObject1444 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1444WithDefaults() *InlineObject1444 {
	this := InlineObject1444{}
	return &this
}

// GetLogical returns the Logical field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1444) GetLogical() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Logical
}

// GetLogicalOk returns a tuple with the Logical field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1444) GetLogicalOk() (*AnyOfobject, bool) {
	if o == nil || o.Logical == nil {
		return nil, false
	}
	return &o.Logical, true
}

// HasLogical returns a boolean if a field has been set.
func (o *InlineObject1444) HasLogical() bool {
	if o != nil && o.Logical != nil {
		return true
	}

	return false
}

// SetLogical gets a reference to the given AnyOfobject and assigns it to the Logical field.
func (o *InlineObject1444) SetLogical(v AnyOfobject) {
	o.Logical = v
}

func (o InlineObject1444) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Logical != nil {
		toSerialize["logical"] = o.Logical
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1444 struct {
	value *InlineObject1444
	isSet bool
}

func (v NullableInlineObject1444) Get() *InlineObject1444 {
	return v.value
}

func (v *NullableInlineObject1444) Set(val *InlineObject1444) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1444) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1444) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1444(val *InlineObject1444) *NullableInlineObject1444 {
	return &NullableInlineObject1444{value: val, isSet: true}
}

func (v NullableInlineObject1444) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1444) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

