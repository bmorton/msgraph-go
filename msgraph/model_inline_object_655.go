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

// InlineObject655 struct for InlineObject655
type InlineObject655 struct {
	Commands *[]*AnyOfmicrosoftGraphOnenotePatchContentCommand `json:"commands,omitempty"`
}

// NewInlineObject655 instantiates a new InlineObject655 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject655() *InlineObject655 {
	this := InlineObject655{}
	return &this
}

// NewInlineObject655WithDefaults instantiates a new InlineObject655 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject655WithDefaults() *InlineObject655 {
	this := InlineObject655{}
	return &this
}

// GetCommands returns the Commands field value if set, zero value otherwise.
func (o *InlineObject655) GetCommands() []*AnyOfmicrosoftGraphOnenotePatchContentCommand {
	if o == nil || o.Commands == nil {
		var ret []*AnyOfmicrosoftGraphOnenotePatchContentCommand
		return ret
	}
	return *o.Commands
}

// GetCommandsOk returns a tuple with the Commands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject655) GetCommandsOk() (*[]*AnyOfmicrosoftGraphOnenotePatchContentCommand, bool) {
	if o == nil || o.Commands == nil {
		return nil, false
	}
	return o.Commands, true
}

// HasCommands returns a boolean if a field has been set.
func (o *InlineObject655) HasCommands() bool {
	if o != nil && o.Commands != nil {
		return true
	}

	return false
}

// SetCommands gets a reference to the given []*AnyOfmicrosoftGraphOnenotePatchContentCommand and assigns it to the Commands field.
func (o *InlineObject655) SetCommands(v []*AnyOfmicrosoftGraphOnenotePatchContentCommand) {
	o.Commands = &v
}

func (o InlineObject655) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Commands != nil {
		toSerialize["commands"] = o.Commands
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject655 struct {
	value *InlineObject655
	isSet bool
}

func (v NullableInlineObject655) Get() *InlineObject655 {
	return v.value
}

func (v *NullableInlineObject655) Set(val *InlineObject655) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject655) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject655) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject655(val *InlineObject655) *NullableInlineObject655 {
	return &NullableInlineObject655{value: val, isSet: true}
}

func (v NullableInlineObject655) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject655) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


