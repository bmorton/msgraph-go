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

// InlineObject120 struct for InlineObject120
type InlineObject120 struct {
	GroupIds *[]string `json:"groupIds,omitempty"`
}

// NewInlineObject120 instantiates a new InlineObject120 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject120() *InlineObject120 {
	this := InlineObject120{}
	return &this
}

// NewInlineObject120WithDefaults instantiates a new InlineObject120 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject120WithDefaults() *InlineObject120 {
	this := InlineObject120{}
	return &this
}

// GetGroupIds returns the GroupIds field value if set, zero value otherwise.
func (o *InlineObject120) GetGroupIds() []string {
	if o == nil || o.GroupIds == nil {
		var ret []string
		return ret
	}
	return *o.GroupIds
}

// GetGroupIdsOk returns a tuple with the GroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject120) GetGroupIdsOk() (*[]string, bool) {
	if o == nil || o.GroupIds == nil {
		return nil, false
	}
	return o.GroupIds, true
}

// HasGroupIds returns a boolean if a field has been set.
func (o *InlineObject120) HasGroupIds() bool {
	if o != nil && o.GroupIds != nil {
		return true
	}

	return false
}

// SetGroupIds gets a reference to the given []string and assigns it to the GroupIds field.
func (o *InlineObject120) SetGroupIds(v []string) {
	o.GroupIds = &v
}

func (o InlineObject120) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GroupIds != nil {
		toSerialize["groupIds"] = o.GroupIds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject120 struct {
	value *InlineObject120
	isSet bool
}

func (v NullableInlineObject120) Get() *InlineObject120 {
	return v.value
}

func (v *NullableInlineObject120) Set(val *InlineObject120) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject120) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject120) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject120(val *InlineObject120) *NullableInlineObject120 {
	return &NullableInlineObject120{value: val, isSet: true}
}

func (v NullableInlineObject120) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject120) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


