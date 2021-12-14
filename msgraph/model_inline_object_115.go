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

// InlineObject115 struct for InlineObject115
type InlineObject115 struct {
	SecurityEnabledOnly NullableBool `json:"securityEnabledOnly,omitempty"`
}

// NewInlineObject115 instantiates a new InlineObject115 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject115() *InlineObject115 {
	this := InlineObject115{}
	var securityEnabledOnly bool = false
	this.SecurityEnabledOnly = *NewNullableBool(&securityEnabledOnly)
	return &this
}

// NewInlineObject115WithDefaults instantiates a new InlineObject115 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject115WithDefaults() *InlineObject115 {
	this := InlineObject115{}
	var securityEnabledOnly bool = false
	this.SecurityEnabledOnly = *NewNullableBool(&securityEnabledOnly)
	return &this
}

// GetSecurityEnabledOnly returns the SecurityEnabledOnly field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject115) GetSecurityEnabledOnly() bool {
	if o == nil || o.SecurityEnabledOnly.Get() == nil {
		var ret bool
		return ret
	}
	return *o.SecurityEnabledOnly.Get()
}

// GetSecurityEnabledOnlyOk returns a tuple with the SecurityEnabledOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject115) GetSecurityEnabledOnlyOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SecurityEnabledOnly.Get(), o.SecurityEnabledOnly.IsSet()
}

// HasSecurityEnabledOnly returns a boolean if a field has been set.
func (o *InlineObject115) HasSecurityEnabledOnly() bool {
	if o != nil && o.SecurityEnabledOnly.IsSet() {
		return true
	}

	return false
}

// SetSecurityEnabledOnly gets a reference to the given NullableBool and assigns it to the SecurityEnabledOnly field.
func (o *InlineObject115) SetSecurityEnabledOnly(v bool) {
	o.SecurityEnabledOnly.Set(&v)
}
// SetSecurityEnabledOnlyNil sets the value for SecurityEnabledOnly to be an explicit nil
func (o *InlineObject115) SetSecurityEnabledOnlyNil() {
	o.SecurityEnabledOnly.Set(nil)
}

// UnsetSecurityEnabledOnly ensures that no value is present for SecurityEnabledOnly, not even an explicit nil
func (o *InlineObject115) UnsetSecurityEnabledOnly() {
	o.SecurityEnabledOnly.Unset()
}

func (o InlineObject115) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SecurityEnabledOnly.IsSet() {
		toSerialize["securityEnabledOnly"] = o.SecurityEnabledOnly.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject115 struct {
	value *InlineObject115
	isSet bool
}

func (v NullableInlineObject115) Get() *InlineObject115 {
	return v.value
}

func (v *NullableInlineObject115) Set(val *InlineObject115) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject115) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject115) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject115(val *InlineObject115) *NullableInlineObject115 {
	return &NullableInlineObject115{value: val, isSet: true}
}

func (v NullableInlineObject115) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject115) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

