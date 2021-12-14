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

// MicrosoftGraphMailTipsError struct for MicrosoftGraphMailTipsError
type MicrosoftGraphMailTipsError struct {
	// The error code.
	Code NullableString `json:"code,omitempty"`
	// The error message.
	Message NullableString `json:"message,omitempty"`
}

// NewMicrosoftGraphMailTipsError instantiates a new MicrosoftGraphMailTipsError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphMailTipsError() *MicrosoftGraphMailTipsError {
	this := MicrosoftGraphMailTipsError{}
	return &this
}

// NewMicrosoftGraphMailTipsErrorWithDefaults instantiates a new MicrosoftGraphMailTipsError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphMailTipsErrorWithDefaults() *MicrosoftGraphMailTipsError {
	this := MicrosoftGraphMailTipsError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphMailTipsError) GetCode() string {
	if o == nil || o.Code.Get() == nil {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphMailTipsError) GetCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *MicrosoftGraphMailTipsError) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *MicrosoftGraphMailTipsError) SetCode(v string) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *MicrosoftGraphMailTipsError) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *MicrosoftGraphMailTipsError) UnsetCode() {
	o.Code.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphMailTipsError) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphMailTipsError) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *MicrosoftGraphMailTipsError) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *MicrosoftGraphMailTipsError) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *MicrosoftGraphMailTipsError) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *MicrosoftGraphMailTipsError) UnsetMessage() {
	o.Message.Unset()
}

func (o MicrosoftGraphMailTipsError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphMailTipsError struct {
	value *MicrosoftGraphMailTipsError
	isSet bool
}

func (v NullableMicrosoftGraphMailTipsError) Get() *MicrosoftGraphMailTipsError {
	return v.value
}

func (v *NullableMicrosoftGraphMailTipsError) Set(val *MicrosoftGraphMailTipsError) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphMailTipsError) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphMailTipsError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphMailTipsError(val *MicrosoftGraphMailTipsError) *NullableMicrosoftGraphMailTipsError {
	return &NullableMicrosoftGraphMailTipsError{value: val, isSet: true}
}

func (v NullableMicrosoftGraphMailTipsError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphMailTipsError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

