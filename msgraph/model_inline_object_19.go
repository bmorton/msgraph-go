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

// InlineObject19 struct for InlineObject19
type InlineObject19 struct {
	Values *[]*AnyOfmicrosoftGraphConversationMember `json:"values,omitempty"`
}

// NewInlineObject19 instantiates a new InlineObject19 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject19() *InlineObject19 {
	this := InlineObject19{}
	return &this
}

// NewInlineObject19WithDefaults instantiates a new InlineObject19 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject19WithDefaults() *InlineObject19 {
	this := InlineObject19{}
	return &this
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *InlineObject19) GetValues() []*AnyOfmicrosoftGraphConversationMember {
	if o == nil || o.Values == nil {
		var ret []*AnyOfmicrosoftGraphConversationMember
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject19) GetValuesOk() (*[]*AnyOfmicrosoftGraphConversationMember, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *InlineObject19) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []*AnyOfmicrosoftGraphConversationMember and assigns it to the Values field.
func (o *InlineObject19) SetValues(v []*AnyOfmicrosoftGraphConversationMember) {
	o.Values = &v
}

func (o InlineObject19) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject19 struct {
	value *InlineObject19
	isSet bool
}

func (v NullableInlineObject19) Get() *InlineObject19 {
	return v.value
}

func (v *NullableInlineObject19) Set(val *InlineObject19) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject19) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject19) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject19(val *InlineObject19) *NullableInlineObject19 {
	return &NullableInlineObject19{value: val, isSet: true}
}

func (v NullableInlineObject19) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject19) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


