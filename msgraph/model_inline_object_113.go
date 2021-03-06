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

// InlineObject113 struct for InlineObject113
type InlineObject113 struct {
	GroupIds *[]string `json:"groupIds,omitempty"`
}

// NewInlineObject113 instantiates a new InlineObject113 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject113() *InlineObject113 {
	this := InlineObject113{}
	return &this
}

// NewInlineObject113WithDefaults instantiates a new InlineObject113 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject113WithDefaults() *InlineObject113 {
	this := InlineObject113{}
	return &this
}

// GetGroupIds returns the GroupIds field value if set, zero value otherwise.
func (o *InlineObject113) GetGroupIds() []string {
	if o == nil || o.GroupIds == nil {
		var ret []string
		return ret
	}
	return *o.GroupIds
}

// GetGroupIdsOk returns a tuple with the GroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject113) GetGroupIdsOk() (*[]string, bool) {
	if o == nil || o.GroupIds == nil {
		return nil, false
	}
	return o.GroupIds, true
}

// HasGroupIds returns a boolean if a field has been set.
func (o *InlineObject113) HasGroupIds() bool {
	if o != nil && o.GroupIds != nil {
		return true
	}

	return false
}

// SetGroupIds gets a reference to the given []string and assigns it to the GroupIds field.
func (o *InlineObject113) SetGroupIds(v []string) {
	o.GroupIds = &v
}

func (o InlineObject113) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GroupIds != nil {
		toSerialize["groupIds"] = o.GroupIds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject113 struct {
	value *InlineObject113
	isSet bool
}

func (v NullableInlineObject113) Get() *InlineObject113 {
	return v.value
}

func (v *NullableInlineObject113) Set(val *InlineObject113) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject113) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject113) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject113(val *InlineObject113) *NullableInlineObject113 {
	return &NullableInlineObject113{value: val, isSet: true}
}

func (v NullableInlineObject113) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject113) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


