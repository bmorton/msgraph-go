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

// InlineObject1322 struct for InlineObject1322
type InlineObject1322 struct {
	X AnyOfobject `json:"X,omitempty"`
}

// NewInlineObject1322 instantiates a new InlineObject1322 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1322() *InlineObject1322 {
	this := InlineObject1322{}
	return &this
}

// NewInlineObject1322WithDefaults instantiates a new InlineObject1322 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1322WithDefaults() *InlineObject1322 {
	this := InlineObject1322{}
	return &this
}

// GetX returns the X field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1322) GetX() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1322) GetXOk() (*AnyOfobject, bool) {
	if o == nil || o.X == nil {
		return nil, false
	}
	return &o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *InlineObject1322) HasX() bool {
	if o != nil && o.X != nil {
		return true
	}

	return false
}

// SetX gets a reference to the given AnyOfobject and assigns it to the X field.
func (o *InlineObject1322) SetX(v AnyOfobject) {
	o.X = v
}

func (o InlineObject1322) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.X != nil {
		toSerialize["X"] = o.X
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1322 struct {
	value *InlineObject1322
	isSet bool
}

func (v NullableInlineObject1322) Get() *InlineObject1322 {
	return v.value
}

func (v *NullableInlineObject1322) Set(val *InlineObject1322) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1322) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1322) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1322(val *InlineObject1322) *NullableInlineObject1322 {
	return &NullableInlineObject1322{value: val, isSet: true}
}

func (v NullableInlineObject1322) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1322) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


