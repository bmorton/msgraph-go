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

// MicrosoftGraphFreeBusyError struct for MicrosoftGraphFreeBusyError
type MicrosoftGraphFreeBusyError struct {
	// Describes the error.
	Message NullableString `json:"message,omitempty"`
	// The response code from querying for the availability of the user, distribution list, or resource.
	ResponseCode NullableString `json:"responseCode,omitempty"`
}

// NewMicrosoftGraphFreeBusyError instantiates a new MicrosoftGraphFreeBusyError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphFreeBusyError() *MicrosoftGraphFreeBusyError {
	this := MicrosoftGraphFreeBusyError{}
	return &this
}

// NewMicrosoftGraphFreeBusyErrorWithDefaults instantiates a new MicrosoftGraphFreeBusyError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphFreeBusyErrorWithDefaults() *MicrosoftGraphFreeBusyError {
	this := MicrosoftGraphFreeBusyError{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphFreeBusyError) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphFreeBusyError) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *MicrosoftGraphFreeBusyError) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *MicrosoftGraphFreeBusyError) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *MicrosoftGraphFreeBusyError) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *MicrosoftGraphFreeBusyError) UnsetMessage() {
	o.Message.Unset()
}

// GetResponseCode returns the ResponseCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphFreeBusyError) GetResponseCode() string {
	if o == nil || o.ResponseCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResponseCode.Get()
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphFreeBusyError) GetResponseCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResponseCode.Get(), o.ResponseCode.IsSet()
}

// HasResponseCode returns a boolean if a field has been set.
func (o *MicrosoftGraphFreeBusyError) HasResponseCode() bool {
	if o != nil && o.ResponseCode.IsSet() {
		return true
	}

	return false
}

// SetResponseCode gets a reference to the given NullableString and assigns it to the ResponseCode field.
func (o *MicrosoftGraphFreeBusyError) SetResponseCode(v string) {
	o.ResponseCode.Set(&v)
}
// SetResponseCodeNil sets the value for ResponseCode to be an explicit nil
func (o *MicrosoftGraphFreeBusyError) SetResponseCodeNil() {
	o.ResponseCode.Set(nil)
}

// UnsetResponseCode ensures that no value is present for ResponseCode, not even an explicit nil
func (o *MicrosoftGraphFreeBusyError) UnsetResponseCode() {
	o.ResponseCode.Unset()
}

func (o MicrosoftGraphFreeBusyError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if o.ResponseCode.IsSet() {
		toSerialize["responseCode"] = o.ResponseCode.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphFreeBusyError struct {
	value *MicrosoftGraphFreeBusyError
	isSet bool
}

func (v NullableMicrosoftGraphFreeBusyError) Get() *MicrosoftGraphFreeBusyError {
	return v.value
}

func (v *NullableMicrosoftGraphFreeBusyError) Set(val *MicrosoftGraphFreeBusyError) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphFreeBusyError) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphFreeBusyError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphFreeBusyError(val *MicrosoftGraphFreeBusyError) *NullableMicrosoftGraphFreeBusyError {
	return &NullableMicrosoftGraphFreeBusyError{value: val, isSet: true}
}

func (v NullableMicrosoftGraphFreeBusyError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphFreeBusyError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

