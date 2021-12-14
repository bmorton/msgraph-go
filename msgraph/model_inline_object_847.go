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

// InlineObject847 struct for InlineObject847
type InlineObject847 struct {
	Values *[]*AnyOfmicrosoftGraphConversationMember `json:"values,omitempty"`
}

// NewInlineObject847 instantiates a new InlineObject847 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject847() *InlineObject847 {
	this := InlineObject847{}
	return &this
}

// NewInlineObject847WithDefaults instantiates a new InlineObject847 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject847WithDefaults() *InlineObject847 {
	this := InlineObject847{}
	return &this
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *InlineObject847) GetValues() []*AnyOfmicrosoftGraphConversationMember {
	if o == nil || o.Values == nil {
		var ret []*AnyOfmicrosoftGraphConversationMember
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject847) GetValuesOk() (*[]*AnyOfmicrosoftGraphConversationMember, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *InlineObject847) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []*AnyOfmicrosoftGraphConversationMember and assigns it to the Values field.
func (o *InlineObject847) SetValues(v []*AnyOfmicrosoftGraphConversationMember) {
	o.Values = &v
}

func (o InlineObject847) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject847 struct {
	value *InlineObject847
	isSet bool
}

func (v NullableInlineObject847) Get() *InlineObject847 {
	return v.value
}

func (v *NullableInlineObject847) Set(val *InlineObject847) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject847) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject847) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject847(val *InlineObject847) *NullableInlineObject847 {
	return &NullableInlineObject847{value: val, isSet: true}
}

func (v NullableInlineObject847) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject847) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


