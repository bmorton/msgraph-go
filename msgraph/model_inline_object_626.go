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

// InlineObject626 struct for InlineObject626
type InlineObject626 struct {
	Commands *[]*AnyOfmicrosoftGraphOnenotePatchContentCommand `json:"commands,omitempty"`
}

// NewInlineObject626 instantiates a new InlineObject626 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject626() *InlineObject626 {
	this := InlineObject626{}
	return &this
}

// NewInlineObject626WithDefaults instantiates a new InlineObject626 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject626WithDefaults() *InlineObject626 {
	this := InlineObject626{}
	return &this
}

// GetCommands returns the Commands field value if set, zero value otherwise.
func (o *InlineObject626) GetCommands() []*AnyOfmicrosoftGraphOnenotePatchContentCommand {
	if o == nil || o.Commands == nil {
		var ret []*AnyOfmicrosoftGraphOnenotePatchContentCommand
		return ret
	}
	return *o.Commands
}

// GetCommandsOk returns a tuple with the Commands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject626) GetCommandsOk() (*[]*AnyOfmicrosoftGraphOnenotePatchContentCommand, bool) {
	if o == nil || o.Commands == nil {
		return nil, false
	}
	return o.Commands, true
}

// HasCommands returns a boolean if a field has been set.
func (o *InlineObject626) HasCommands() bool {
	if o != nil && o.Commands != nil {
		return true
	}

	return false
}

// SetCommands gets a reference to the given []*AnyOfmicrosoftGraphOnenotePatchContentCommand and assigns it to the Commands field.
func (o *InlineObject626) SetCommands(v []*AnyOfmicrosoftGraphOnenotePatchContentCommand) {
	o.Commands = &v
}

func (o InlineObject626) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Commands != nil {
		toSerialize["commands"] = o.Commands
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject626 struct {
	value *InlineObject626
	isSet bool
}

func (v NullableInlineObject626) Get() *InlineObject626 {
	return v.value
}

func (v *NullableInlineObject626) Set(val *InlineObject626) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject626) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject626) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject626(val *InlineObject626) *NullableInlineObject626 {
	return &NullableInlineObject626{value: val, isSet: true}
}

func (v NullableInlineObject626) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject626) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


