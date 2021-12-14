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

// InlineObject1585 struct for InlineObject1585
type InlineObject1585 struct {
	Icon AnyOfmicrosoftGraphWorkbookIcon `json:"icon,omitempty"`
}

// NewInlineObject1585 instantiates a new InlineObject1585 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1585() *InlineObject1585 {
	this := InlineObject1585{}
	return &this
}

// NewInlineObject1585WithDefaults instantiates a new InlineObject1585 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1585WithDefaults() *InlineObject1585 {
	this := InlineObject1585{}
	return &this
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1585) GetIcon() AnyOfmicrosoftGraphWorkbookIcon {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookIcon
		return ret
	}
	return o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1585) GetIconOk() (*AnyOfmicrosoftGraphWorkbookIcon, bool) {
	if o == nil || o.Icon == nil {
		return nil, false
	}
	return &o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *InlineObject1585) HasIcon() bool {
	if o != nil && o.Icon != nil {
		return true
	}

	return false
}

// SetIcon gets a reference to the given AnyOfmicrosoftGraphWorkbookIcon and assigns it to the Icon field.
func (o *InlineObject1585) SetIcon(v AnyOfmicrosoftGraphWorkbookIcon) {
	o.Icon = v
}

func (o InlineObject1585) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Icon != nil {
		toSerialize["icon"] = o.Icon
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1585 struct {
	value *InlineObject1585
	isSet bool
}

func (v NullableInlineObject1585) Get() *InlineObject1585 {
	return v.value
}

func (v *NullableInlineObject1585) Set(val *InlineObject1585) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1585) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1585) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1585(val *InlineObject1585) *NullableInlineObject1585 {
	return &NullableInlineObject1585{value: val, isSet: true}
}

func (v NullableInlineObject1585) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1585) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

