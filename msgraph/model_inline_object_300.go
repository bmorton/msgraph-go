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

// InlineObject300 struct for InlineObject300
type InlineObject300 struct {
	Commands *[]*AnyOfmicrosoftGraphOnenotePatchContentCommand `json:"commands,omitempty"`
}

// NewInlineObject300 instantiates a new InlineObject300 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject300() *InlineObject300 {
	this := InlineObject300{}
	return &this
}

// NewInlineObject300WithDefaults instantiates a new InlineObject300 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject300WithDefaults() *InlineObject300 {
	this := InlineObject300{}
	return &this
}

// GetCommands returns the Commands field value if set, zero value otherwise.
func (o *InlineObject300) GetCommands() []*AnyOfmicrosoftGraphOnenotePatchContentCommand {
	if o == nil || o.Commands == nil {
		var ret []*AnyOfmicrosoftGraphOnenotePatchContentCommand
		return ret
	}
	return *o.Commands
}

// GetCommandsOk returns a tuple with the Commands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject300) GetCommandsOk() (*[]*AnyOfmicrosoftGraphOnenotePatchContentCommand, bool) {
	if o == nil || o.Commands == nil {
		return nil, false
	}
	return o.Commands, true
}

// HasCommands returns a boolean if a field has been set.
func (o *InlineObject300) HasCommands() bool {
	if o != nil && o.Commands != nil {
		return true
	}

	return false
}

// SetCommands gets a reference to the given []*AnyOfmicrosoftGraphOnenotePatchContentCommand and assigns it to the Commands field.
func (o *InlineObject300) SetCommands(v []*AnyOfmicrosoftGraphOnenotePatchContentCommand) {
	o.Commands = &v
}

func (o InlineObject300) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Commands != nil {
		toSerialize["commands"] = o.Commands
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject300 struct {
	value *InlineObject300
	isSet bool
}

func (v NullableInlineObject300) Get() *InlineObject300 {
	return v.value
}

func (v *NullableInlineObject300) Set(val *InlineObject300) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject300) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject300) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject300(val *InlineObject300) *NullableInlineObject300 {
	return &NullableInlineObject300{value: val, isSet: true}
}

func (v NullableInlineObject300) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject300) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


