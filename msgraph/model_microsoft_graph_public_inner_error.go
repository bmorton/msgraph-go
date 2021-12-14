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

// MicrosoftGraphPublicInnerError struct for MicrosoftGraphPublicInnerError
type MicrosoftGraphPublicInnerError struct {
	// The error code.
	Code NullableString `json:"code,omitempty"`
	// A collection of error details.
	Details *[]*AnyOfmicrosoftGraphPublicErrorDetail `json:"details,omitempty"`
	// The error message.
	Message NullableString `json:"message,omitempty"`
	// The target of the error.
	Target NullableString `json:"target,omitempty"`
}

// NewMicrosoftGraphPublicInnerError instantiates a new MicrosoftGraphPublicInnerError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphPublicInnerError() *MicrosoftGraphPublicInnerError {
	this := MicrosoftGraphPublicInnerError{}
	return &this
}

// NewMicrosoftGraphPublicInnerErrorWithDefaults instantiates a new MicrosoftGraphPublicInnerError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphPublicInnerErrorWithDefaults() *MicrosoftGraphPublicInnerError {
	this := MicrosoftGraphPublicInnerError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPublicInnerError) GetCode() string {
	if o == nil || o.Code.Get() == nil {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPublicInnerError) GetCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *MicrosoftGraphPublicInnerError) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *MicrosoftGraphPublicInnerError) SetCode(v string) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *MicrosoftGraphPublicInnerError) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *MicrosoftGraphPublicInnerError) UnsetCode() {
	o.Code.Unset()
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *MicrosoftGraphPublicInnerError) GetDetails() []*AnyOfmicrosoftGraphPublicErrorDetail {
	if o == nil || o.Details == nil {
		var ret []*AnyOfmicrosoftGraphPublicErrorDetail
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPublicInnerError) GetDetailsOk() (*[]*AnyOfmicrosoftGraphPublicErrorDetail, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *MicrosoftGraphPublicInnerError) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []*AnyOfmicrosoftGraphPublicErrorDetail and assigns it to the Details field.
func (o *MicrosoftGraphPublicInnerError) SetDetails(v []*AnyOfmicrosoftGraphPublicErrorDetail) {
	o.Details = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPublicInnerError) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPublicInnerError) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *MicrosoftGraphPublicInnerError) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *MicrosoftGraphPublicInnerError) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *MicrosoftGraphPublicInnerError) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *MicrosoftGraphPublicInnerError) UnsetMessage() {
	o.Message.Unset()
}

// GetTarget returns the Target field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPublicInnerError) GetTarget() string {
	if o == nil || o.Target.Get() == nil {
		var ret string
		return ret
	}
	return *o.Target.Get()
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPublicInnerError) GetTargetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Target.Get(), o.Target.IsSet()
}

// HasTarget returns a boolean if a field has been set.
func (o *MicrosoftGraphPublicInnerError) HasTarget() bool {
	if o != nil && o.Target.IsSet() {
		return true
	}

	return false
}

// SetTarget gets a reference to the given NullableString and assigns it to the Target field.
func (o *MicrosoftGraphPublicInnerError) SetTarget(v string) {
	o.Target.Set(&v)
}
// SetTargetNil sets the value for Target to be an explicit nil
func (o *MicrosoftGraphPublicInnerError) SetTargetNil() {
	o.Target.Set(nil)
}

// UnsetTarget ensures that no value is present for Target, not even an explicit nil
func (o *MicrosoftGraphPublicInnerError) UnsetTarget() {
	o.Target.Unset()
}

func (o MicrosoftGraphPublicInnerError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if o.Target.IsSet() {
		toSerialize["target"] = o.Target.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphPublicInnerError struct {
	value *MicrosoftGraphPublicInnerError
	isSet bool
}

func (v NullableMicrosoftGraphPublicInnerError) Get() *MicrosoftGraphPublicInnerError {
	return v.value
}

func (v *NullableMicrosoftGraphPublicInnerError) Set(val *MicrosoftGraphPublicInnerError) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphPublicInnerError) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphPublicInnerError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphPublicInnerError(val *MicrosoftGraphPublicInnerError) *NullableMicrosoftGraphPublicInnerError {
	return &NullableMicrosoftGraphPublicInnerError{value: val, isSet: true}
}

func (v NullableMicrosoftGraphPublicInnerError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphPublicInnerError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


