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

// InlineObject1066 struct for InlineObject1066
type InlineObject1066 struct {
	UserPrincipalName NullableString `json:"userPrincipalName,omitempty"`
}

// NewInlineObject1066 instantiates a new InlineObject1066 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1066() *InlineObject1066 {
	this := InlineObject1066{}
	return &this
}

// NewInlineObject1066WithDefaults instantiates a new InlineObject1066 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1066WithDefaults() *InlineObject1066 {
	this := InlineObject1066{}
	return &this
}

// GetUserPrincipalName returns the UserPrincipalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1066) GetUserPrincipalName() string {
	if o == nil || o.UserPrincipalName.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserPrincipalName.Get()
}

// GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1066) GetUserPrincipalNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserPrincipalName.Get(), o.UserPrincipalName.IsSet()
}

// HasUserPrincipalName returns a boolean if a field has been set.
func (o *InlineObject1066) HasUserPrincipalName() bool {
	if o != nil && o.UserPrincipalName.IsSet() {
		return true
	}

	return false
}

// SetUserPrincipalName gets a reference to the given NullableString and assigns it to the UserPrincipalName field.
func (o *InlineObject1066) SetUserPrincipalName(v string) {
	o.UserPrincipalName.Set(&v)
}
// SetUserPrincipalNameNil sets the value for UserPrincipalName to be an explicit nil
func (o *InlineObject1066) SetUserPrincipalNameNil() {
	o.UserPrincipalName.Set(nil)
}

// UnsetUserPrincipalName ensures that no value is present for UserPrincipalName, not even an explicit nil
func (o *InlineObject1066) UnsetUserPrincipalName() {
	o.UserPrincipalName.Unset()
}

func (o InlineObject1066) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserPrincipalName.IsSet() {
		toSerialize["userPrincipalName"] = o.UserPrincipalName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1066 struct {
	value *InlineObject1066
	isSet bool
}

func (v NullableInlineObject1066) Get() *InlineObject1066 {
	return v.value
}

func (v *NullableInlineObject1066) Set(val *InlineObject1066) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1066) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1066) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1066(val *InlineObject1066) *NullableInlineObject1066 {
	return &NullableInlineObject1066{value: val, isSet: true}
}

func (v NullableInlineObject1066) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1066) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


