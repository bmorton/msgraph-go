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

// InlineObject803 struct for InlineObject803
type InlineObject803 struct {
	Commands *[]*AnyOfmicrosoftGraphOnenotePatchContentCommand `json:"commands,omitempty"`
}

// NewInlineObject803 instantiates a new InlineObject803 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject803() *InlineObject803 {
	this := InlineObject803{}
	return &this
}

// NewInlineObject803WithDefaults instantiates a new InlineObject803 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject803WithDefaults() *InlineObject803 {
	this := InlineObject803{}
	return &this
}

// GetCommands returns the Commands field value if set, zero value otherwise.
func (o *InlineObject803) GetCommands() []*AnyOfmicrosoftGraphOnenotePatchContentCommand {
	if o == nil || o.Commands == nil {
		var ret []*AnyOfmicrosoftGraphOnenotePatchContentCommand
		return ret
	}
	return *o.Commands
}

// GetCommandsOk returns a tuple with the Commands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject803) GetCommandsOk() (*[]*AnyOfmicrosoftGraphOnenotePatchContentCommand, bool) {
	if o == nil || o.Commands == nil {
		return nil, false
	}
	return o.Commands, true
}

// HasCommands returns a boolean if a field has been set.
func (o *InlineObject803) HasCommands() bool {
	if o != nil && o.Commands != nil {
		return true
	}

	return false
}

// SetCommands gets a reference to the given []*AnyOfmicrosoftGraphOnenotePatchContentCommand and assigns it to the Commands field.
func (o *InlineObject803) SetCommands(v []*AnyOfmicrosoftGraphOnenotePatchContentCommand) {
	o.Commands = &v
}

func (o InlineObject803) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Commands != nil {
		toSerialize["commands"] = o.Commands
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject803 struct {
	value *InlineObject803
	isSet bool
}

func (v NullableInlineObject803) Get() *InlineObject803 {
	return v.value
}

func (v *NullableInlineObject803) Set(val *InlineObject803) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject803) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject803) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject803(val *InlineObject803) *NullableInlineObject803 {
	return &NullableInlineObject803{value: val, isSet: true}
}

func (v NullableInlineObject803) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject803) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

