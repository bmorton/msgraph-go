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

// InlineObject538 struct for InlineObject538
type InlineObject538 struct {
	ApplyTo *string `json:"applyTo,omitempty"`
}

// NewInlineObject538 instantiates a new InlineObject538 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject538() *InlineObject538 {
	this := InlineObject538{}
	return &this
}

// NewInlineObject538WithDefaults instantiates a new InlineObject538 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject538WithDefaults() *InlineObject538 {
	this := InlineObject538{}
	return &this
}

// GetApplyTo returns the ApplyTo field value if set, zero value otherwise.
func (o *InlineObject538) GetApplyTo() string {
	if o == nil || o.ApplyTo == nil {
		var ret string
		return ret
	}
	return *o.ApplyTo
}

// GetApplyToOk returns a tuple with the ApplyTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject538) GetApplyToOk() (*string, bool) {
	if o == nil || o.ApplyTo == nil {
		return nil, false
	}
	return o.ApplyTo, true
}

// HasApplyTo returns a boolean if a field has been set.
func (o *InlineObject538) HasApplyTo() bool {
	if o != nil && o.ApplyTo != nil {
		return true
	}

	return false
}

// SetApplyTo gets a reference to the given string and assigns it to the ApplyTo field.
func (o *InlineObject538) SetApplyTo(v string) {
	o.ApplyTo = &v
}

func (o InlineObject538) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApplyTo != nil {
		toSerialize["applyTo"] = o.ApplyTo
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject538 struct {
	value *InlineObject538
	isSet bool
}

func (v NullableInlineObject538) Get() *InlineObject538 {
	return v.value
}

func (v *NullableInlineObject538) Set(val *InlineObject538) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject538) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject538) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject538(val *InlineObject538) *NullableInlineObject538 {
	return &NullableInlineObject538{value: val, isSet: true}
}

func (v NullableInlineObject538) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject538) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


