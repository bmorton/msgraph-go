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

// InlineObject1533 struct for InlineObject1533
type InlineObject1533 struct {
	Value AnyOfobject `json:"value,omitempty"`
	FormatText AnyOfobject `json:"formatText,omitempty"`
}

// NewInlineObject1533 instantiates a new InlineObject1533 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1533() *InlineObject1533 {
	this := InlineObject1533{}
	return &this
}

// NewInlineObject1533WithDefaults instantiates a new InlineObject1533 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1533WithDefaults() *InlineObject1533 {
	this := InlineObject1533{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1533) GetValue() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1533) GetValueOk() (*AnyOfobject, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InlineObject1533) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given AnyOfobject and assigns it to the Value field.
func (o *InlineObject1533) SetValue(v AnyOfobject) {
	o.Value = v
}

// GetFormatText returns the FormatText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1533) GetFormatText() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.FormatText
}

// GetFormatTextOk returns a tuple with the FormatText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1533) GetFormatTextOk() (*AnyOfobject, bool) {
	if o == nil || o.FormatText == nil {
		return nil, false
	}
	return &o.FormatText, true
}

// HasFormatText returns a boolean if a field has been set.
func (o *InlineObject1533) HasFormatText() bool {
	if o != nil && o.FormatText != nil {
		return true
	}

	return false
}

// SetFormatText gets a reference to the given AnyOfobject and assigns it to the FormatText field.
func (o *InlineObject1533) SetFormatText(v AnyOfobject) {
	o.FormatText = v
}

func (o InlineObject1533) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.FormatText != nil {
		toSerialize["formatText"] = o.FormatText
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1533 struct {
	value *InlineObject1533
	isSet bool
}

func (v NullableInlineObject1533) Get() *InlineObject1533 {
	return v.value
}

func (v *NullableInlineObject1533) Set(val *InlineObject1533) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1533) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1533) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1533(val *InlineObject1533) *NullableInlineObject1533 {
	return &NullableInlineObject1533{value: val, isSet: true}
}

func (v NullableInlineObject1533) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1533) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


