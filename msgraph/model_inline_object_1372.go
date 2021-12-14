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

// InlineObject1372 struct for InlineObject1372
type InlineObject1372 struct {
	Inumber1 AnyOfobject `json:"inumber1,omitempty"`
	Inumber2 AnyOfobject `json:"inumber2,omitempty"`
}

// NewInlineObject1372 instantiates a new InlineObject1372 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1372() *InlineObject1372 {
	this := InlineObject1372{}
	return &this
}

// NewInlineObject1372WithDefaults instantiates a new InlineObject1372 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1372WithDefaults() *InlineObject1372 {
	this := InlineObject1372{}
	return &this
}

// GetInumber1 returns the Inumber1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1372) GetInumber1() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Inumber1
}

// GetInumber1Ok returns a tuple with the Inumber1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1372) GetInumber1Ok() (*AnyOfobject, bool) {
	if o == nil || o.Inumber1 == nil {
		return nil, false
	}
	return &o.Inumber1, true
}

// HasInumber1 returns a boolean if a field has been set.
func (o *InlineObject1372) HasInumber1() bool {
	if o != nil && o.Inumber1 != nil {
		return true
	}

	return false
}

// SetInumber1 gets a reference to the given AnyOfobject and assigns it to the Inumber1 field.
func (o *InlineObject1372) SetInumber1(v AnyOfobject) {
	o.Inumber1 = v
}

// GetInumber2 returns the Inumber2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1372) GetInumber2() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Inumber2
}

// GetInumber2Ok returns a tuple with the Inumber2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1372) GetInumber2Ok() (*AnyOfobject, bool) {
	if o == nil || o.Inumber2 == nil {
		return nil, false
	}
	return &o.Inumber2, true
}

// HasInumber2 returns a boolean if a field has been set.
func (o *InlineObject1372) HasInumber2() bool {
	if o != nil && o.Inumber2 != nil {
		return true
	}

	return false
}

// SetInumber2 gets a reference to the given AnyOfobject and assigns it to the Inumber2 field.
func (o *InlineObject1372) SetInumber2(v AnyOfobject) {
	o.Inumber2 = v
}

func (o InlineObject1372) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Inumber1 != nil {
		toSerialize["inumber1"] = o.Inumber1
	}
	if o.Inumber2 != nil {
		toSerialize["inumber2"] = o.Inumber2
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1372 struct {
	value *InlineObject1372
	isSet bool
}

func (v NullableInlineObject1372) Get() *InlineObject1372 {
	return v.value
}

func (v *NullableInlineObject1372) Set(val *InlineObject1372) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1372) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1372) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1372(val *InlineObject1372) *NullableInlineObject1372 {
	return &NullableInlineObject1372{value: val, isSet: true}
}

func (v NullableInlineObject1372) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1372) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


