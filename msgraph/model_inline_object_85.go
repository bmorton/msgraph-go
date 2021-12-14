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

// InlineObject85 struct for InlineObject85
type InlineObject85 struct {
	UpdateWindowsDeviceAccountActionParameter AnyOfmicrosoftGraphUpdateWindowsDeviceAccountActionParameter `json:"updateWindowsDeviceAccountActionParameter,omitempty"`
}

// NewInlineObject85 instantiates a new InlineObject85 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject85() *InlineObject85 {
	this := InlineObject85{}
	return &this
}

// NewInlineObject85WithDefaults instantiates a new InlineObject85 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject85WithDefaults() *InlineObject85 {
	this := InlineObject85{}
	return &this
}

// GetUpdateWindowsDeviceAccountActionParameter returns the UpdateWindowsDeviceAccountActionParameter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject85) GetUpdateWindowsDeviceAccountActionParameter() AnyOfmicrosoftGraphUpdateWindowsDeviceAccountActionParameter {
	if o == nil  {
		var ret AnyOfmicrosoftGraphUpdateWindowsDeviceAccountActionParameter
		return ret
	}
	return o.UpdateWindowsDeviceAccountActionParameter
}

// GetUpdateWindowsDeviceAccountActionParameterOk returns a tuple with the UpdateWindowsDeviceAccountActionParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject85) GetUpdateWindowsDeviceAccountActionParameterOk() (*AnyOfmicrosoftGraphUpdateWindowsDeviceAccountActionParameter, bool) {
	if o == nil || o.UpdateWindowsDeviceAccountActionParameter == nil {
		return nil, false
	}
	return &o.UpdateWindowsDeviceAccountActionParameter, true
}

// HasUpdateWindowsDeviceAccountActionParameter returns a boolean if a field has been set.
func (o *InlineObject85) HasUpdateWindowsDeviceAccountActionParameter() bool {
	if o != nil && o.UpdateWindowsDeviceAccountActionParameter != nil {
		return true
	}

	return false
}

// SetUpdateWindowsDeviceAccountActionParameter gets a reference to the given AnyOfmicrosoftGraphUpdateWindowsDeviceAccountActionParameter and assigns it to the UpdateWindowsDeviceAccountActionParameter field.
func (o *InlineObject85) SetUpdateWindowsDeviceAccountActionParameter(v AnyOfmicrosoftGraphUpdateWindowsDeviceAccountActionParameter) {
	o.UpdateWindowsDeviceAccountActionParameter = v
}

func (o InlineObject85) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UpdateWindowsDeviceAccountActionParameter != nil {
		toSerialize["updateWindowsDeviceAccountActionParameter"] = o.UpdateWindowsDeviceAccountActionParameter
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject85 struct {
	value *InlineObject85
	isSet bool
}

func (v NullableInlineObject85) Get() *InlineObject85 {
	return v.value
}

func (v *NullableInlineObject85) Set(val *InlineObject85) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject85) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject85) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject85(val *InlineObject85) *NullableInlineObject85 {
	return &NullableInlineObject85{value: val, isSet: true}
}

func (v NullableInlineObject85) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject85) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


