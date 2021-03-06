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

// InlineObject50 struct for InlineObject50
type InlineObject50 struct {
	Ids *[]string `json:"ids,omitempty"`
}

// NewInlineObject50 instantiates a new InlineObject50 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject50() *InlineObject50 {
	this := InlineObject50{}
	return &this
}

// NewInlineObject50WithDefaults instantiates a new InlineObject50 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject50WithDefaults() *InlineObject50 {
	this := InlineObject50{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *InlineObject50) GetIds() []string {
	if o == nil || o.Ids == nil {
		var ret []string
		return ret
	}
	return *o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject50) GetIdsOk() (*[]string, bool) {
	if o == nil || o.Ids == nil {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *InlineObject50) HasIds() bool {
	if o != nil && o.Ids != nil {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *InlineObject50) SetIds(v []string) {
	o.Ids = &v
}

func (o InlineObject50) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject50 struct {
	value *InlineObject50
	isSet bool
}

func (v NullableInlineObject50) Get() *InlineObject50 {
	return v.value
}

func (v *NullableInlineObject50) Set(val *InlineObject50) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject50) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject50) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject50(val *InlineObject50) *NullableInlineObject50 {
	return &NullableInlineObject50{value: val, isSet: true}
}

func (v NullableInlineObject50) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject50) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


