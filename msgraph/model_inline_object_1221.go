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

// InlineObject1221 struct for InlineObject1221
type InlineObject1221 struct {
	Number AnyOfobject `json:"number,omitempty"`
}

// NewInlineObject1221 instantiates a new InlineObject1221 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1221() *InlineObject1221 {
	this := InlineObject1221{}
	return &this
}

// NewInlineObject1221WithDefaults instantiates a new InlineObject1221 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1221WithDefaults() *InlineObject1221 {
	this := InlineObject1221{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1221) GetNumber() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1221) GetNumberOk() (*AnyOfobject, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return &o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *InlineObject1221) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given AnyOfobject and assigns it to the Number field.
func (o *InlineObject1221) SetNumber(v AnyOfobject) {
	o.Number = v
}

func (o InlineObject1221) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1221 struct {
	value *InlineObject1221
	isSet bool
}

func (v NullableInlineObject1221) Get() *InlineObject1221 {
	return v.value
}

func (v *NullableInlineObject1221) Set(val *InlineObject1221) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1221) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1221) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1221(val *InlineObject1221) *NullableInlineObject1221 {
	return &NullableInlineObject1221{value: val, isSet: true}
}

func (v NullableInlineObject1221) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1221) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


