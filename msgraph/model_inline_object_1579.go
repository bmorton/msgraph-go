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

// InlineObject1579 struct for InlineObject1579
type InlineObject1579 struct {
	Count *int32 `json:"count,omitempty"`
}

// NewInlineObject1579 instantiates a new InlineObject1579 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1579() *InlineObject1579 {
	this := InlineObject1579{}
	return &this
}

// NewInlineObject1579WithDefaults instantiates a new InlineObject1579 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1579WithDefaults() *InlineObject1579 {
	this := InlineObject1579{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *InlineObject1579) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1579) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *InlineObject1579) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *InlineObject1579) SetCount(v int32) {
	o.Count = &v
}

func (o InlineObject1579) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1579 struct {
	value *InlineObject1579
	isSet bool
}

func (v NullableInlineObject1579) Get() *InlineObject1579 {
	return v.value
}

func (v *NullableInlineObject1579) Set(val *InlineObject1579) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1579) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1579) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1579(val *InlineObject1579) *NullableInlineObject1579 {
	return &NullableInlineObject1579{value: val, isSet: true}
}

func (v NullableInlineObject1579) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1579) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


