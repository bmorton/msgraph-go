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

// InlineObject1597 struct for InlineObject1597
type InlineObject1597 struct {
	Percent *int32 `json:"percent,omitempty"`
}

// NewInlineObject1597 instantiates a new InlineObject1597 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1597() *InlineObject1597 {
	this := InlineObject1597{}
	return &this
}

// NewInlineObject1597WithDefaults instantiates a new InlineObject1597 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1597WithDefaults() *InlineObject1597 {
	this := InlineObject1597{}
	return &this
}

// GetPercent returns the Percent field value if set, zero value otherwise.
func (o *InlineObject1597) GetPercent() int32 {
	if o == nil || o.Percent == nil {
		var ret int32
		return ret
	}
	return *o.Percent
}

// GetPercentOk returns a tuple with the Percent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1597) GetPercentOk() (*int32, bool) {
	if o == nil || o.Percent == nil {
		return nil, false
	}
	return o.Percent, true
}

// HasPercent returns a boolean if a field has been set.
func (o *InlineObject1597) HasPercent() bool {
	if o != nil && o.Percent != nil {
		return true
	}

	return false
}

// SetPercent gets a reference to the given int32 and assigns it to the Percent field.
func (o *InlineObject1597) SetPercent(v int32) {
	o.Percent = &v
}

func (o InlineObject1597) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Percent != nil {
		toSerialize["percent"] = o.Percent
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1597 struct {
	value *InlineObject1597
	isSet bool
}

func (v NullableInlineObject1597) Get() *InlineObject1597 {
	return v.value
}

func (v *NullableInlineObject1597) Set(val *InlineObject1597) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1597) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1597) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1597(val *InlineObject1597) *NullableInlineObject1597 {
	return &NullableInlineObject1597{value: val, isSet: true}
}

func (v NullableInlineObject1597) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1597) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


