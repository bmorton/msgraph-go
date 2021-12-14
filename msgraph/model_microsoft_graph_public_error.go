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

// MicrosoftGraphPublicError struct for MicrosoftGraphPublicError
type MicrosoftGraphPublicError struct {
	// Represents the error code.
	Code NullableString `json:"code,omitempty"`
	// Details of the error.
	Details *[]*AnyOfmicrosoftGraphPublicErrorDetail `json:"details,omitempty"`
	// Details of the inner error.
	InnerError AnyOfmicrosoftGraphPublicInnerError `json:"innerError,omitempty"`
	// A non-localized message for the developer.
	Message NullableString `json:"message,omitempty"`
	// The target of the error.
	Target NullableString `json:"target,omitempty"`
}

// NewMicrosoftGraphPublicError instantiates a new MicrosoftGraphPublicError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphPublicError() *MicrosoftGraphPublicError {
	this := MicrosoftGraphPublicError{}
	return &this
}

// NewMicrosoftGraphPublicErrorWithDefaults instantiates a new MicrosoftGraphPublicError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphPublicErrorWithDefaults() *MicrosoftGraphPublicError {
	this := MicrosoftGraphPublicError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPublicError) GetCode() string {
	if o == nil || o.Code.Get() == nil {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPublicError) GetCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *MicrosoftGraphPublicError) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *MicrosoftGraphPublicError) SetCode(v string) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *MicrosoftGraphPublicError) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *MicrosoftGraphPublicError) UnsetCode() {
	o.Code.Unset()
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *MicrosoftGraphPublicError) GetDetails() []*AnyOfmicrosoftGraphPublicErrorDetail {
	if o == nil || o.Details == nil {
		var ret []*AnyOfmicrosoftGraphPublicErrorDetail
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPublicError) GetDetailsOk() (*[]*AnyOfmicrosoftGraphPublicErrorDetail, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *MicrosoftGraphPublicError) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []*AnyOfmicrosoftGraphPublicErrorDetail and assigns it to the Details field.
func (o *MicrosoftGraphPublicError) SetDetails(v []*AnyOfmicrosoftGraphPublicErrorDetail) {
	o.Details = &v
}

// GetInnerError returns the InnerError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPublicError) GetInnerError() AnyOfmicrosoftGraphPublicInnerError {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPublicInnerError
		return ret
	}
	return o.InnerError
}

// GetInnerErrorOk returns a tuple with the InnerError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPublicError) GetInnerErrorOk() (*AnyOfmicrosoftGraphPublicInnerError, bool) {
	if o == nil || o.InnerError == nil {
		return nil, false
	}
	return &o.InnerError, true
}

// HasInnerError returns a boolean if a field has been set.
func (o *MicrosoftGraphPublicError) HasInnerError() bool {
	if o != nil && o.InnerError != nil {
		return true
	}

	return false
}

// SetInnerError gets a reference to the given AnyOfmicrosoftGraphPublicInnerError and assigns it to the InnerError field.
func (o *MicrosoftGraphPublicError) SetInnerError(v AnyOfmicrosoftGraphPublicInnerError) {
	o.InnerError = v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPublicError) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPublicError) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *MicrosoftGraphPublicError) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *MicrosoftGraphPublicError) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *MicrosoftGraphPublicError) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *MicrosoftGraphPublicError) UnsetMessage() {
	o.Message.Unset()
}

// GetTarget returns the Target field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPublicError) GetTarget() string {
	if o == nil || o.Target.Get() == nil {
		var ret string
		return ret
	}
	return *o.Target.Get()
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPublicError) GetTargetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Target.Get(), o.Target.IsSet()
}

// HasTarget returns a boolean if a field has been set.
func (o *MicrosoftGraphPublicError) HasTarget() bool {
	if o != nil && o.Target.IsSet() {
		return true
	}

	return false
}

// SetTarget gets a reference to the given NullableString and assigns it to the Target field.
func (o *MicrosoftGraphPublicError) SetTarget(v string) {
	o.Target.Set(&v)
}
// SetTargetNil sets the value for Target to be an explicit nil
func (o *MicrosoftGraphPublicError) SetTargetNil() {
	o.Target.Set(nil)
}

// UnsetTarget ensures that no value is present for Target, not even an explicit nil
func (o *MicrosoftGraphPublicError) UnsetTarget() {
	o.Target.Unset()
}

func (o MicrosoftGraphPublicError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.InnerError != nil {
		toSerialize["innerError"] = o.InnerError
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if o.Target.IsSet() {
		toSerialize["target"] = o.Target.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphPublicError struct {
	value *MicrosoftGraphPublicError
	isSet bool
}

func (v NullableMicrosoftGraphPublicError) Get() *MicrosoftGraphPublicError {
	return v.value
}

func (v *NullableMicrosoftGraphPublicError) Set(val *MicrosoftGraphPublicError) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphPublicError) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphPublicError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphPublicError(val *MicrosoftGraphPublicError) *NullableMicrosoftGraphPublicError {
	return &NullableMicrosoftGraphPublicError{value: val, isSet: true}
}

func (v NullableMicrosoftGraphPublicError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphPublicError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

