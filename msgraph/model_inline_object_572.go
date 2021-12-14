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

// InlineObject572 struct for InlineObject572
type InlineObject572 struct {
	UserPrincipalName NullableString `json:"userPrincipalName,omitempty"`
}

// NewInlineObject572 instantiates a new InlineObject572 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject572() *InlineObject572 {
	this := InlineObject572{}
	return &this
}

// NewInlineObject572WithDefaults instantiates a new InlineObject572 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject572WithDefaults() *InlineObject572 {
	this := InlineObject572{}
	return &this
}

// GetUserPrincipalName returns the UserPrincipalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject572) GetUserPrincipalName() string {
	if o == nil || o.UserPrincipalName.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserPrincipalName.Get()
}

// GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject572) GetUserPrincipalNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserPrincipalName.Get(), o.UserPrincipalName.IsSet()
}

// HasUserPrincipalName returns a boolean if a field has been set.
func (o *InlineObject572) HasUserPrincipalName() bool {
	if o != nil && o.UserPrincipalName.IsSet() {
		return true
	}

	return false
}

// SetUserPrincipalName gets a reference to the given NullableString and assigns it to the UserPrincipalName field.
func (o *InlineObject572) SetUserPrincipalName(v string) {
	o.UserPrincipalName.Set(&v)
}
// SetUserPrincipalNameNil sets the value for UserPrincipalName to be an explicit nil
func (o *InlineObject572) SetUserPrincipalNameNil() {
	o.UserPrincipalName.Set(nil)
}

// UnsetUserPrincipalName ensures that no value is present for UserPrincipalName, not even an explicit nil
func (o *InlineObject572) UnsetUserPrincipalName() {
	o.UserPrincipalName.Unset()
}

func (o InlineObject572) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserPrincipalName.IsSet() {
		toSerialize["userPrincipalName"] = o.UserPrincipalName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject572 struct {
	value *InlineObject572
	isSet bool
}

func (v NullableInlineObject572) Get() *InlineObject572 {
	return v.value
}

func (v *NullableInlineObject572) Set(val *InlineObject572) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject572) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject572) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject572(val *InlineObject572) *NullableInlineObject572 {
	return &NullableInlineObject572{value: val, isSet: true}
}

func (v NullableInlineObject572) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject572) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


