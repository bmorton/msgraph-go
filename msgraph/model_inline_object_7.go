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

// InlineObject7 struct for InlineObject7
type InlineObject7 struct {
	PasswordCredential AnyOfmicrosoftGraphPasswordCredential `json:"passwordCredential,omitempty"`
}

// NewInlineObject7 instantiates a new InlineObject7 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject7() *InlineObject7 {
	this := InlineObject7{}
	return &this
}

// NewInlineObject7WithDefaults instantiates a new InlineObject7 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject7WithDefaults() *InlineObject7 {
	this := InlineObject7{}
	return &this
}

// GetPasswordCredential returns the PasswordCredential field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject7) GetPasswordCredential() AnyOfmicrosoftGraphPasswordCredential {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPasswordCredential
		return ret
	}
	return o.PasswordCredential
}

// GetPasswordCredentialOk returns a tuple with the PasswordCredential field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject7) GetPasswordCredentialOk() (*AnyOfmicrosoftGraphPasswordCredential, bool) {
	if o == nil || o.PasswordCredential == nil {
		return nil, false
	}
	return &o.PasswordCredential, true
}

// HasPasswordCredential returns a boolean if a field has been set.
func (o *InlineObject7) HasPasswordCredential() bool {
	if o != nil && o.PasswordCredential != nil {
		return true
	}

	return false
}

// SetPasswordCredential gets a reference to the given AnyOfmicrosoftGraphPasswordCredential and assigns it to the PasswordCredential field.
func (o *InlineObject7) SetPasswordCredential(v AnyOfmicrosoftGraphPasswordCredential) {
	o.PasswordCredential = v
}

func (o InlineObject7) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PasswordCredential != nil {
		toSerialize["passwordCredential"] = o.PasswordCredential
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject7 struct {
	value *InlineObject7
	isSet bool
}

func (v NullableInlineObject7) Get() *InlineObject7 {
	return v.value
}

func (v *NullableInlineObject7) Set(val *InlineObject7) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject7) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject7) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject7(val *InlineObject7) *NullableInlineObject7 {
	return &NullableInlineObject7{value: val, isSet: true}
}

func (v NullableInlineObject7) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject7) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


