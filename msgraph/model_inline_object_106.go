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

// InlineObject106 struct for InlineObject106
type InlineObject106 struct {
	GroupIds *[]string `json:"groupIds,omitempty"`
}

// NewInlineObject106 instantiates a new InlineObject106 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject106() *InlineObject106 {
	this := InlineObject106{}
	return &this
}

// NewInlineObject106WithDefaults instantiates a new InlineObject106 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject106WithDefaults() *InlineObject106 {
	this := InlineObject106{}
	return &this
}

// GetGroupIds returns the GroupIds field value if set, zero value otherwise.
func (o *InlineObject106) GetGroupIds() []string {
	if o == nil || o.GroupIds == nil {
		var ret []string
		return ret
	}
	return *o.GroupIds
}

// GetGroupIdsOk returns a tuple with the GroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject106) GetGroupIdsOk() (*[]string, bool) {
	if o == nil || o.GroupIds == nil {
		return nil, false
	}
	return o.GroupIds, true
}

// HasGroupIds returns a boolean if a field has been set.
func (o *InlineObject106) HasGroupIds() bool {
	if o != nil && o.GroupIds != nil {
		return true
	}

	return false
}

// SetGroupIds gets a reference to the given []string and assigns it to the GroupIds field.
func (o *InlineObject106) SetGroupIds(v []string) {
	o.GroupIds = &v
}

func (o InlineObject106) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GroupIds != nil {
		toSerialize["groupIds"] = o.GroupIds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject106 struct {
	value *InlineObject106
	isSet bool
}

func (v NullableInlineObject106) Get() *InlineObject106 {
	return v.value
}

func (v *NullableInlineObject106) Set(val *InlineObject106) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject106) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject106) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject106(val *InlineObject106) *NullableInlineObject106 {
	return &NullableInlineObject106{value: val, isSet: true}
}

func (v NullableInlineObject106) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject106) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


