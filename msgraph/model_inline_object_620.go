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

// InlineObject620 struct for InlineObject620
type InlineObject620 struct {
	Commands *[]*AnyOfmicrosoftGraphOnenotePatchContentCommand `json:"commands,omitempty"`
}

// NewInlineObject620 instantiates a new InlineObject620 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject620() *InlineObject620 {
	this := InlineObject620{}
	return &this
}

// NewInlineObject620WithDefaults instantiates a new InlineObject620 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject620WithDefaults() *InlineObject620 {
	this := InlineObject620{}
	return &this
}

// GetCommands returns the Commands field value if set, zero value otherwise.
func (o *InlineObject620) GetCommands() []*AnyOfmicrosoftGraphOnenotePatchContentCommand {
	if o == nil || o.Commands == nil {
		var ret []*AnyOfmicrosoftGraphOnenotePatchContentCommand
		return ret
	}
	return *o.Commands
}

// GetCommandsOk returns a tuple with the Commands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject620) GetCommandsOk() (*[]*AnyOfmicrosoftGraphOnenotePatchContentCommand, bool) {
	if o == nil || o.Commands == nil {
		return nil, false
	}
	return o.Commands, true
}

// HasCommands returns a boolean if a field has been set.
func (o *InlineObject620) HasCommands() bool {
	if o != nil && o.Commands != nil {
		return true
	}

	return false
}

// SetCommands gets a reference to the given []*AnyOfmicrosoftGraphOnenotePatchContentCommand and assigns it to the Commands field.
func (o *InlineObject620) SetCommands(v []*AnyOfmicrosoftGraphOnenotePatchContentCommand) {
	o.Commands = &v
}

func (o InlineObject620) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Commands != nil {
		toSerialize["commands"] = o.Commands
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject620 struct {
	value *InlineObject620
	isSet bool
}

func (v NullableInlineObject620) Get() *InlineObject620 {
	return v.value
}

func (v *NullableInlineObject620) Set(val *InlineObject620) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject620) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject620) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject620(val *InlineObject620) *NullableInlineObject620 {
	return &NullableInlineObject620{value: val, isSet: true}
}

func (v NullableInlineObject620) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject620) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

