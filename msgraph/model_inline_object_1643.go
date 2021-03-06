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

// InlineObject1643 struct for InlineObject1643
type InlineObject1643 struct {
	Count *int32 `json:"count,omitempty"`
}

// NewInlineObject1643 instantiates a new InlineObject1643 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1643() *InlineObject1643 {
	this := InlineObject1643{}
	return &this
}

// NewInlineObject1643WithDefaults instantiates a new InlineObject1643 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1643WithDefaults() *InlineObject1643 {
	this := InlineObject1643{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *InlineObject1643) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1643) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *InlineObject1643) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *InlineObject1643) SetCount(v int32) {
	o.Count = &v
}

func (o InlineObject1643) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1643 struct {
	value *InlineObject1643
	isSet bool
}

func (v NullableInlineObject1643) Get() *InlineObject1643 {
	return v.value
}

func (v *NullableInlineObject1643) Set(val *InlineObject1643) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1643) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1643) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1643(val *InlineObject1643) *NullableInlineObject1643 {
	return &NullableInlineObject1643{value: val, isSet: true}
}

func (v NullableInlineObject1643) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1643) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


