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

// InlineObject593 struct for InlineObject593
type InlineObject593 struct {
	SecurityEnabledOnly NullableBool `json:"securityEnabledOnly,omitempty"`
}

// NewInlineObject593 instantiates a new InlineObject593 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject593() *InlineObject593 {
	this := InlineObject593{}
	var securityEnabledOnly bool = false
	this.SecurityEnabledOnly = *NewNullableBool(&securityEnabledOnly)
	return &this
}

// NewInlineObject593WithDefaults instantiates a new InlineObject593 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject593WithDefaults() *InlineObject593 {
	this := InlineObject593{}
	var securityEnabledOnly bool = false
	this.SecurityEnabledOnly = *NewNullableBool(&securityEnabledOnly)
	return &this
}

// GetSecurityEnabledOnly returns the SecurityEnabledOnly field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject593) GetSecurityEnabledOnly() bool {
	if o == nil || o.SecurityEnabledOnly.Get() == nil {
		var ret bool
		return ret
	}
	return *o.SecurityEnabledOnly.Get()
}

// GetSecurityEnabledOnlyOk returns a tuple with the SecurityEnabledOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject593) GetSecurityEnabledOnlyOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SecurityEnabledOnly.Get(), o.SecurityEnabledOnly.IsSet()
}

// HasSecurityEnabledOnly returns a boolean if a field has been set.
func (o *InlineObject593) HasSecurityEnabledOnly() bool {
	if o != nil && o.SecurityEnabledOnly.IsSet() {
		return true
	}

	return false
}

// SetSecurityEnabledOnly gets a reference to the given NullableBool and assigns it to the SecurityEnabledOnly field.
func (o *InlineObject593) SetSecurityEnabledOnly(v bool) {
	o.SecurityEnabledOnly.Set(&v)
}
// SetSecurityEnabledOnlyNil sets the value for SecurityEnabledOnly to be an explicit nil
func (o *InlineObject593) SetSecurityEnabledOnlyNil() {
	o.SecurityEnabledOnly.Set(nil)
}

// UnsetSecurityEnabledOnly ensures that no value is present for SecurityEnabledOnly, not even an explicit nil
func (o *InlineObject593) UnsetSecurityEnabledOnly() {
	o.SecurityEnabledOnly.Unset()
}

func (o InlineObject593) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SecurityEnabledOnly.IsSet() {
		toSerialize["securityEnabledOnly"] = o.SecurityEnabledOnly.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject593 struct {
	value *InlineObject593
	isSet bool
}

func (v NullableInlineObject593) Get() *InlineObject593 {
	return v.value
}

func (v *NullableInlineObject593) Set(val *InlineObject593) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject593) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject593) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject593(val *InlineObject593) *NullableInlineObject593 {
	return &NullableInlineObject593{value: val, isSet: true}
}

func (v NullableInlineObject593) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject593) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


