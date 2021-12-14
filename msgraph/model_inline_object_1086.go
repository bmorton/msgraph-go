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

// InlineObject1086 struct for InlineObject1086
type InlineObject1086 struct {
	SecurityEnabledOnly NullableBool `json:"securityEnabledOnly,omitempty"`
}

// NewInlineObject1086 instantiates a new InlineObject1086 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1086() *InlineObject1086 {
	this := InlineObject1086{}
	var securityEnabledOnly bool = false
	this.SecurityEnabledOnly = *NewNullableBool(&securityEnabledOnly)
	return &this
}

// NewInlineObject1086WithDefaults instantiates a new InlineObject1086 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1086WithDefaults() *InlineObject1086 {
	this := InlineObject1086{}
	var securityEnabledOnly bool = false
	this.SecurityEnabledOnly = *NewNullableBool(&securityEnabledOnly)
	return &this
}

// GetSecurityEnabledOnly returns the SecurityEnabledOnly field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1086) GetSecurityEnabledOnly() bool {
	if o == nil || o.SecurityEnabledOnly.Get() == nil {
		var ret bool
		return ret
	}
	return *o.SecurityEnabledOnly.Get()
}

// GetSecurityEnabledOnlyOk returns a tuple with the SecurityEnabledOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1086) GetSecurityEnabledOnlyOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SecurityEnabledOnly.Get(), o.SecurityEnabledOnly.IsSet()
}

// HasSecurityEnabledOnly returns a boolean if a field has been set.
func (o *InlineObject1086) HasSecurityEnabledOnly() bool {
	if o != nil && o.SecurityEnabledOnly.IsSet() {
		return true
	}

	return false
}

// SetSecurityEnabledOnly gets a reference to the given NullableBool and assigns it to the SecurityEnabledOnly field.
func (o *InlineObject1086) SetSecurityEnabledOnly(v bool) {
	o.SecurityEnabledOnly.Set(&v)
}
// SetSecurityEnabledOnlyNil sets the value for SecurityEnabledOnly to be an explicit nil
func (o *InlineObject1086) SetSecurityEnabledOnlyNil() {
	o.SecurityEnabledOnly.Set(nil)
}

// UnsetSecurityEnabledOnly ensures that no value is present for SecurityEnabledOnly, not even an explicit nil
func (o *InlineObject1086) UnsetSecurityEnabledOnly() {
	o.SecurityEnabledOnly.Unset()
}

func (o InlineObject1086) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SecurityEnabledOnly.IsSet() {
		toSerialize["securityEnabledOnly"] = o.SecurityEnabledOnly.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1086 struct {
	value *InlineObject1086
	isSet bool
}

func (v NullableInlineObject1086) Get() *InlineObject1086 {
	return v.value
}

func (v *NullableInlineObject1086) Set(val *InlineObject1086) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1086) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1086) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1086(val *InlineObject1086) *NullableInlineObject1086 {
	return &NullableInlineObject1086{value: val, isSet: true}
}

func (v NullableInlineObject1086) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1086) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


