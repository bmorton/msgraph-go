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

// InlineObject1373 struct for InlineObject1373
type InlineObject1373 struct {
	Inumber AnyOfobject `json:"inumber,omitempty"`
}

// NewInlineObject1373 instantiates a new InlineObject1373 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1373() *InlineObject1373 {
	this := InlineObject1373{}
	return &this
}

// NewInlineObject1373WithDefaults instantiates a new InlineObject1373 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1373WithDefaults() *InlineObject1373 {
	this := InlineObject1373{}
	return &this
}

// GetInumber returns the Inumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1373) GetInumber() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Inumber
}

// GetInumberOk returns a tuple with the Inumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1373) GetInumberOk() (*AnyOfobject, bool) {
	if o == nil || o.Inumber == nil {
		return nil, false
	}
	return &o.Inumber, true
}

// HasInumber returns a boolean if a field has been set.
func (o *InlineObject1373) HasInumber() bool {
	if o != nil && o.Inumber != nil {
		return true
	}

	return false
}

// SetInumber gets a reference to the given AnyOfobject and assigns it to the Inumber field.
func (o *InlineObject1373) SetInumber(v AnyOfobject) {
	o.Inumber = v
}

func (o InlineObject1373) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Inumber != nil {
		toSerialize["inumber"] = o.Inumber
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1373 struct {
	value *InlineObject1373
	isSet bool
}

func (v NullableInlineObject1373) Get() *InlineObject1373 {
	return v.value
}

func (v *NullableInlineObject1373) Set(val *InlineObject1373) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1373) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1373) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1373(val *InlineObject1373) *NullableInlineObject1373 {
	return &NullableInlineObject1373{value: val, isSet: true}
}

func (v NullableInlineObject1373) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1373) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


