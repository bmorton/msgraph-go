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

// InlineObject707 struct for InlineObject707
type InlineObject707 struct {
	Ids *[]string `json:"ids,omitempty"`
}

// NewInlineObject707 instantiates a new InlineObject707 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject707() *InlineObject707 {
	this := InlineObject707{}
	return &this
}

// NewInlineObject707WithDefaults instantiates a new InlineObject707 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject707WithDefaults() *InlineObject707 {
	this := InlineObject707{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *InlineObject707) GetIds() []string {
	if o == nil || o.Ids == nil {
		var ret []string
		return ret
	}
	return *o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject707) GetIdsOk() (*[]string, bool) {
	if o == nil || o.Ids == nil {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *InlineObject707) HasIds() bool {
	if o != nil && o.Ids != nil {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *InlineObject707) SetIds(v []string) {
	o.Ids = &v
}

func (o InlineObject707) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject707 struct {
	value *InlineObject707
	isSet bool
}

func (v NullableInlineObject707) Get() *InlineObject707 {
	return v.value
}

func (v *NullableInlineObject707) Set(val *InlineObject707) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject707) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject707) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject707(val *InlineObject707) *NullableInlineObject707 {
	return &NullableInlineObject707{value: val, isSet: true}
}

func (v NullableInlineObject707) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject707) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


