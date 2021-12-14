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

// InlineObject354 struct for InlineObject354
type InlineObject354 struct {
	Ids *[]string `json:"ids,omitempty"`
	Types *[]*string `json:"types,omitempty"`
}

// NewInlineObject354 instantiates a new InlineObject354 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject354() *InlineObject354 {
	this := InlineObject354{}
	return &this
}

// NewInlineObject354WithDefaults instantiates a new InlineObject354 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject354WithDefaults() *InlineObject354 {
	this := InlineObject354{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *InlineObject354) GetIds() []string {
	if o == nil || o.Ids == nil {
		var ret []string
		return ret
	}
	return *o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject354) GetIdsOk() (*[]string, bool) {
	if o == nil || o.Ids == nil {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *InlineObject354) HasIds() bool {
	if o != nil && o.Ids != nil {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *InlineObject354) SetIds(v []string) {
	o.Ids = &v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *InlineObject354) GetTypes() []*string {
	if o == nil || o.Types == nil {
		var ret []*string
		return ret
	}
	return *o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject354) GetTypesOk() (*[]*string, bool) {
	if o == nil || o.Types == nil {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *InlineObject354) HasTypes() bool {
	if o != nil && o.Types != nil {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []*string and assigns it to the Types field.
func (o *InlineObject354) SetTypes(v []*string) {
	o.Types = &v
}

func (o InlineObject354) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	if o.Types != nil {
		toSerialize["types"] = o.Types
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject354 struct {
	value *InlineObject354
	isSet bool
}

func (v NullableInlineObject354) Get() *InlineObject354 {
	return v.value
}

func (v *NullableInlineObject354) Set(val *InlineObject354) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject354) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject354) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject354(val *InlineObject354) *NullableInlineObject354 {
	return &NullableInlineObject354{value: val, isSet: true}
}

func (v NullableInlineObject354) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject354) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


