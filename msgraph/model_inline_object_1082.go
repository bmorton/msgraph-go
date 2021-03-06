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

// InlineObject1082 struct for InlineObject1082
type InlineObject1082 struct {
	Ids *[]string `json:"ids,omitempty"`
}

// NewInlineObject1082 instantiates a new InlineObject1082 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1082() *InlineObject1082 {
	this := InlineObject1082{}
	return &this
}

// NewInlineObject1082WithDefaults instantiates a new InlineObject1082 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1082WithDefaults() *InlineObject1082 {
	this := InlineObject1082{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *InlineObject1082) GetIds() []string {
	if o == nil || o.Ids == nil {
		var ret []string
		return ret
	}
	return *o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1082) GetIdsOk() (*[]string, bool) {
	if o == nil || o.Ids == nil {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *InlineObject1082) HasIds() bool {
	if o != nil && o.Ids != nil {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *InlineObject1082) SetIds(v []string) {
	o.Ids = &v
}

func (o InlineObject1082) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1082 struct {
	value *InlineObject1082
	isSet bool
}

func (v NullableInlineObject1082) Get() *InlineObject1082 {
	return v.value
}

func (v *NullableInlineObject1082) Set(val *InlineObject1082) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1082) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1082) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1082(val *InlineObject1082) *NullableInlineObject1082 {
	return &NullableInlineObject1082{value: val, isSet: true}
}

func (v NullableInlineObject1082) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1082) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


