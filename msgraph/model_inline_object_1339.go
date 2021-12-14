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

// InlineObject1339 struct for InlineObject1339
type InlineObject1339 struct {
	Y AnyOfobject `json:"y,omitempty"`
}

// NewInlineObject1339 instantiates a new InlineObject1339 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1339() *InlineObject1339 {
	this := InlineObject1339{}
	return &this
}

// NewInlineObject1339WithDefaults instantiates a new InlineObject1339 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1339WithDefaults() *InlineObject1339 {
	this := InlineObject1339{}
	return &this
}

// GetY returns the Y field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1339) GetY() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1339) GetYOk() (*AnyOfobject, bool) {
	if o == nil || o.Y == nil {
		return nil, false
	}
	return &o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *InlineObject1339) HasY() bool {
	if o != nil && o.Y != nil {
		return true
	}

	return false
}

// SetY gets a reference to the given AnyOfobject and assigns it to the Y field.
func (o *InlineObject1339) SetY(v AnyOfobject) {
	o.Y = v
}

func (o InlineObject1339) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Y != nil {
		toSerialize["y"] = o.Y
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1339 struct {
	value *InlineObject1339
	isSet bool
}

func (v NullableInlineObject1339) Get() *InlineObject1339 {
	return v.value
}

func (v *NullableInlineObject1339) Set(val *InlineObject1339) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1339) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1339) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1339(val *InlineObject1339) *NullableInlineObject1339 {
	return &NullableInlineObject1339{value: val, isSet: true}
}

func (v NullableInlineObject1339) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1339) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


