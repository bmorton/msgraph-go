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

// InlineObject49 struct for InlineObject49
type InlineObject49 struct {
	GroupIds *[]string `json:"groupIds,omitempty"`
}

// NewInlineObject49 instantiates a new InlineObject49 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject49() *InlineObject49 {
	this := InlineObject49{}
	return &this
}

// NewInlineObject49WithDefaults instantiates a new InlineObject49 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject49WithDefaults() *InlineObject49 {
	this := InlineObject49{}
	return &this
}

// GetGroupIds returns the GroupIds field value if set, zero value otherwise.
func (o *InlineObject49) GetGroupIds() []string {
	if o == nil || o.GroupIds == nil {
		var ret []string
		return ret
	}
	return *o.GroupIds
}

// GetGroupIdsOk returns a tuple with the GroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject49) GetGroupIdsOk() (*[]string, bool) {
	if o == nil || o.GroupIds == nil {
		return nil, false
	}
	return o.GroupIds, true
}

// HasGroupIds returns a boolean if a field has been set.
func (o *InlineObject49) HasGroupIds() bool {
	if o != nil && o.GroupIds != nil {
		return true
	}

	return false
}

// SetGroupIds gets a reference to the given []string and assigns it to the GroupIds field.
func (o *InlineObject49) SetGroupIds(v []string) {
	o.GroupIds = &v
}

func (o InlineObject49) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GroupIds != nil {
		toSerialize["groupIds"] = o.GroupIds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject49 struct {
	value *InlineObject49
	isSet bool
}

func (v NullableInlineObject49) Get() *InlineObject49 {
	return v.value
}

func (v *NullableInlineObject49) Set(val *InlineObject49) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject49) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject49) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject49(val *InlineObject49) *NullableInlineObject49 {
	return &NullableInlineObject49{value: val, isSet: true}
}

func (v NullableInlineObject49) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject49) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


