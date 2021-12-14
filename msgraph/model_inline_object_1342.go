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

// InlineObject1342 struct for InlineObject1342
type InlineObject1342 struct {
	Number AnyOfobject `json:"number,omitempty"`
	Significance AnyOfobject `json:"significance,omitempty"`
}

// NewInlineObject1342 instantiates a new InlineObject1342 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1342() *InlineObject1342 {
	this := InlineObject1342{}
	return &this
}

// NewInlineObject1342WithDefaults instantiates a new InlineObject1342 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1342WithDefaults() *InlineObject1342 {
	this := InlineObject1342{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1342) GetNumber() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1342) GetNumberOk() (*AnyOfobject, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return &o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *InlineObject1342) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given AnyOfobject and assigns it to the Number field.
func (o *InlineObject1342) SetNumber(v AnyOfobject) {
	o.Number = v
}

// GetSignificance returns the Significance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1342) GetSignificance() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Significance
}

// GetSignificanceOk returns a tuple with the Significance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1342) GetSignificanceOk() (*AnyOfobject, bool) {
	if o == nil || o.Significance == nil {
		return nil, false
	}
	return &o.Significance, true
}

// HasSignificance returns a boolean if a field has been set.
func (o *InlineObject1342) HasSignificance() bool {
	if o != nil && o.Significance != nil {
		return true
	}

	return false
}

// SetSignificance gets a reference to the given AnyOfobject and assigns it to the Significance field.
func (o *InlineObject1342) SetSignificance(v AnyOfobject) {
	o.Significance = v
}

func (o InlineObject1342) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.Significance != nil {
		toSerialize["significance"] = o.Significance
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1342 struct {
	value *InlineObject1342
	isSet bool
}

func (v NullableInlineObject1342) Get() *InlineObject1342 {
	return v.value
}

func (v *NullableInlineObject1342) Set(val *InlineObject1342) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1342) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1342) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1342(val *InlineObject1342) *NullableInlineObject1342 {
	return &NullableInlineObject1342{value: val, isSet: true}
}

func (v NullableInlineObject1342) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1342) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


